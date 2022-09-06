package snowflake

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/info-schema/schemata.html
func tableSnowflakeLoginHistory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_login_history",
		Description: "This Account Usage view table can be used to query login attempts by Snowflake users within the last 365 days (1 year).",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeLoginHistory,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "event_id", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "event_timestamp", Require: plugin.Optional, Operators: []string{">", ">=", "=", "<", "<="}},
				{Name: "user_name", Require: plugin.Optional, Operators: []string{"="}},
				{Name: "first_authentication_factor", Require: plugin.Optional, Operators: []string{"="}},
			},
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "event_id", Type: proto.ColumnType_INT, Description: "Internal/system-generated identifier for the login attempt."},
			{Name: "event_timestamp", Type: proto.ColumnType_TIMESTAMP, Description: "Time (in the UTC time zone) of the event occurrence."},
			{Name: "event_type", Type: proto.ColumnType_STRING, Description: "Event type, such as LOGIN for authentication events."},
			{Name: "user_name", Type: proto.ColumnType_STRING, Description: "User associated with this event."},
			{Name: "client_ip", Type: proto.ColumnType_STRING, Description: "IP address where the request originated from."},
			{Name: "reported_client_type", Type: proto.ColumnType_STRING, Description: "Reported type of the client software, such as JDBC_DRIVER, ODBC_DRIVER, etc. This information is not authenticated."},
			{Name: "reported_client_version", Type: proto.ColumnType_STRING, Description: "Reported version of the client software. This information is not authenticated."},
			{Name: "first_authentication_factor", Type: proto.ColumnType_STRING, Description: "Method used to authenticate the user (the first factor, if using multi factor authentication)."},
			{Name: "second_authentication_factor", Type: proto.ColumnType_STRING, Description: "The second factor, if using multi factor authentication, or NULL otherwise."},
			{Name: "is_success", Type: proto.ColumnType_STRING, Description: "Whether the user's request was successful or not."},
			{Name: "error_code", Type: proto.ColumnType_INT, Description: "Error code, if the request was not successful."},
			{Name: "error_message", Type: proto.ColumnType_STRING, Description: "Error message returned to the user, if the request was not successful."},
			{Name: "related_event_id", Type: proto.ColumnType_INT, Description: "Reserved for future use."},
		}),
	}
}

type LoginHistory struct {
	EventId                    sql.NullInt64  `json:"EVENT_ID"`
	EventTimestamp             sql.NullTime   `json:"EVENT_TIMESTAMP"`
	EventType                  sql.NullString `json:"EVENT_TYPE"`
	UserName                   sql.NullString `json:"USER_NAME"`
	ClientIp                   sql.NullString `json:"CLIENT_IP"`
	ReportedClientType         sql.NullString `json:"REPORTED_CLIENT_TYPE"`
	ReportedClientVersion      sql.NullString `json:"REPORTED_CLIENT_VERSION"`
	FirstAuthenticationFactor  sql.NullString `json:"FIRST_AUTHENTICATION_FACTOR"`
	SecondAuthenticationFactor sql.NullString `json:"SECOND_AUTHENTICATION_FACTOR"`
	IsSuccess                  sql.NullString `json:"IS_SUCCESS"`
	ErrorCode                  sql.NullInt64  `json:"ERROR_CODE"`
	ErrorMessage               sql.NullString `json:"ERROR_MESSAGE"`
	RelatedEventId             sql.NullInt64  `json:"RELATED_EVENT_ID"`
}

// LoginHistoryCol returns a reference for a column of a LoginHistory
func LoginHistoryCol(colname string, item *LoginHistory) interface{} {
	switch colname {
	case "EVENT_ID":
		return &item.EventId
	case "EVENT_TIMESTAMP":
		return &item.EventTimestamp
	case "EVENT_TYPE":
		return &item.EventType
	case "USER_NAME":
		return &item.UserName
	case "CLIENT_IP":
		return &item.ClientIp
	case "REPORTED_CLIENT_TYPE":
		return &item.ReportedClientType
	case "REPORTED_CLIENT_VERSION":
		return &item.ReportedClientVersion
	case "FIRST_AUTHENTICATION_FACTOR":
		return &item.FirstAuthenticationFactor
	case "SECOND_AUTHENTICATION_FACTOR":
		return &item.SecondAuthenticationFactor
	case "IS_SUCCESS":
		return &item.IsSuccess
	case "ERROR_CODE":
		return &item.ErrorCode
	case "ERROR_MESSAGE":
		return &item.ErrorMessage
	case "RELATED_EVENT_ID":
		return &item.RelatedEventId
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeLoginHistory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_login_history.listSnowflakeLoginHistory", "connnection.error", err)
		return nil, err
	}

	equalQuals := d.KeyColumnQuals
	quals := d.Quals
	var conditions []string = []string{}
	if equalQuals["user_name"] != nil {
		conditions = append(conditions, fmt.Sprintf("user_name %s '%s'", "=", equalQuals["user_name"].GetStringValue()))
	}

	if equalQuals["event_id"] != nil {
		conditions = append(conditions, fmt.Sprintf("event_id %s '%s'", "=", equalQuals["event_id"].GetStringValue()))
	}

	if equalQuals["first_authentication_factor"] != nil {
		conditions = append(conditions, fmt.Sprintf("first_authentication_factor %s '%s'", "=", equalQuals["first_authentication_factor"].GetStringValue()))
	}

	if quals["event_timestamp"] != nil {
		for _, q := range quals["event_timestamp"].Quals {
			tsSecs := q.Value.GetTimestampValue().AsTime().Format("2006-01-02 15:04:05.000")
			conditions = append(conditions, fmt.Sprintf("event_timestamp %s to_timestamp_ltz('%s')", q.Operator, tsSecs))
		}
	}

	condition := strings.Join(conditions, " and ")
	query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.LOGIN_HISTORY"

	if condition != "" {
		query = fmt.Sprintf("%s where %s", query, condition)
	}

	// logger.Info("listSnowflakeLoginHistory", "query", query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("snowflake_login_history.listSnowflakeLoginHistory", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_login_history.listSnowflakeLoginHistory", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		loginHistory := LoginHistory{}
		// make references for the cols with the aid of LoginHistoryCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = LoginHistoryCol(col, &loginHistory)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_login_history.listSnowflakeLoginHistory", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, loginHistory)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			loginHistory := LoginHistory{}
			// make references for the cols with the aid of LoginHistoryCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = LoginHistoryCol(col, &loginHistory)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_login_history.listSnowflakeLoginHistory", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, loginHistory)
		}
	}
	return nil, nil
}
