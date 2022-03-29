package snowflake

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/account-usage/sessions.html
func tableSnowflakeSession(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_session",
		Description: "This Account Usage view provides information on the session, including information on the authentication method to Snowflake and the Snowflake login event. Snowflake returns one row for each session created over the last year.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeSession,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "session_id", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "user_name", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "created_on", Require: plugin.Optional, Operators: []string{">", ">=", "=", "<", "<="}},
				{Name: "authentication_method", Require: plugin.Optional, Operators: []string{"="}},
			},
		},
		Columns: snowflakeColumns([]*plugin.Column{
			// Top fields
			{Name: "session_id", Type: proto.ColumnType_INT, Description: "The unique identifier for the current session."},
			{Name: "user_name", Type: proto.ColumnType_STRING, Description: "The user name of the user."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the session was created."},
			{Name: "authentication_method", Type: proto.ColumnType_STRING, Description: "The authentication method used to access Snowflake."},

			// Other fields
			{Name: "client_application_id", Type: proto.ColumnType_STRING, Description: "The identifier for the Snowflake-provided client application used to create the remote session to Snowflake (e.g. JDBC 3.8.7)"},
			{Name: "client_application_version", Type: proto.ColumnType_STRING, Description: "The version number (e.g. 3.8.7) of the Snowflake-provided client application used to create the remote session to Snowflake."},
			{Name: "client_build_id", Type: proto.ColumnType_STRING, Description: "The build number (e.g. 41897) of the third-party client application used to create a remote session to Snowflake, if available. For example, a third-party Java application that uses the JDBC driver to connect to Snowflake."},
			{Name: "client_environment", Type: proto.ColumnType_JSON, Description: "The environment variables (e.g. operating system, OCSP mode) of the client used to create a remote session to Snowflake."},
			{Name: "client_version", Type: proto.ColumnType_STRING, Description: "The version number (e.g. 47154) of the third-party client application that uses a Snowflake-provided client to create a remote session to Snowflake, if available."},
			{Name: "login_event_id", Type: proto.ColumnType_INT, Description: "The unique identifier for the login event."},
		}),
	}
}

type Session struct {
	SessionId                sql.NullInt64  `json:"SESSION_ID"`
	CreatedOn                sql.NullTime   `json:"CREATED_ON"`
	UserName                 sql.NullString `json:"USER_NAME"`
	AuthenticationMethod     sql.NullString `json:"AUTHENTICATION_METHOD"`
	LoginEventId             sql.NullInt64  `json:"LOGIN_EVENT_ID"`
	ClientApplicationVersion sql.NullString `json:"CLIENT_APPLICATION_VERSION"`
	ClientApplicationId      sql.NullString `json:"CLIENT_APPLICATION_ID"`
	ClientEnvironment        sql.NullString `json:"CLIENT_ENVIRONMENT"`
	ClientBuildId            sql.NullString `json:"CLIENT_BUILD_ID"`
	ClientVersion            sql.NullString `json:"CLIENT_VERSION"`
}

// SessionCol returns a reference for a column of a Session
func SessionCol(colname string, item *Session) interface{} {
	switch colname {
	case "SESSION_ID":
		return &item.SessionId
	case "CREATED_ON":
		return &item.CreatedOn
	case "USER_NAME":
		return &item.UserName
	case "AUTHENTICATION_METHOD":
		return &item.AuthenticationMethod
	case "LOGIN_EVENT_ID":
		return &item.LoginEventId
	case "CLIENT_APPLICATION_VERSION":
		return &item.ClientApplicationVersion
	case "CLIENT_APPLICATION_ID":
		return &item.ClientApplicationId
	case "CLIENT_ENVIRONMENT":
		return &item.ClientEnvironment
	case "CLIENT_BUILD_ID":
		return &item.ClientBuildId
	case "CLIENT_VERSION":
		return &item.ClientVersion
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeSession(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_session.listSnowflakeSession", "connnection.error", err)
		return nil, err
	}

	equalQuals := d.KeyColumnQuals
	quals := d.Quals
	var conditions []string = []string{}

	if equalQuals["session_id"] != nil {
		conditions = append(conditions, fmt.Sprintf("session_id %s '%s'", "=", equalQuals["session_id"].GetStringValue()))
	}

	if equalQuals["user_name"] != nil {
		conditions = append(conditions, fmt.Sprintf("user_name %s '%s'", "=", equalQuals["user_name"].GetStringValue()))
	}

	if equalQuals["authentication_method"] != nil {
		conditions = append(conditions, fmt.Sprintf("authentication_method %s '%s'", "=", equalQuals["authentication_method"].GetStringValue()))
	}

	if quals["created_on"] != nil {
		for _, q := range quals["created_on"].Quals {
			tsSecs := q.Value.GetTimestampValue().AsTime().Format("2006-01-02 15:04:05.000")
			conditions = append(conditions, fmt.Sprintf("created_on %s to_timestamp_ltz('%s')", q.Operator, tsSecs))
		}
	}

	condition := strings.Join(conditions, " and ")
	query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.SESSIONS"
	if condition != "" {
		query = fmt.Sprintf("%s where %s;", query, condition)
	}

	logger.Info("listSnowflakeLoginHistory", "query", query)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("snowflake_session.listSnowflakeSession", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_session.listSnowflakeSession", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		session := Session{}
		// make references for the cols with the aid of SessionCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = SessionCol(col, &session)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_session.listSnowflakeSession", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, session)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			session := Session{}
			// make references for the cols with the aid of SessionCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = SessionCol(col, &session)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_session.listSnowflakeSession", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, session)
		}
	}
	return nil, nil
}
