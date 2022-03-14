package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	gosnowflake "github.com/snowflakedb/gosnowflake"
	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/ddl-user-security.html#label-session-policy-ddl
// This command requires the role executing the command to have:
// 	The OWNERSHIP privilege on the session policy or the APPLY on SESSION POLICY privilege.
// 	The USAGE privilege on the schema.
func tableSnowflakeSessionPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_session_policy",
		Description: "A session policy defines the idle session timeout period in minutes and provides the option to override the default idle timeout value of 4 hours.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeSessionPolicies,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Identifier for the session policy."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time of the creation of session policy."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "Name of the database policy belongs."},
			{Name: "schema_name", Type: proto.ColumnType_STRING, Description: "Name of the schema in database policy belongs."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Type of the snowflake policy."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Name of the role that owns the policy."},
			{Name: "session_idle_timeout_mins", Type: proto.ColumnType_INT, Hydrate: DescribeSessionPolicy, Description: "Time period in minutes of inactivity with either the web interface or a programmatic client"},
			{Name: "session_ui_idle_timeout_mins", Type: proto.ColumnType_INT, Hydrate: DescribeSessionPolicy, Description: "Time period in minutes of inactivity with the web interface."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for this policy"},
		},
	}
}

type SessionPolicy Policy

//// LIST FUNCTION

func listSnowflakeSessionPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_row_access_policy.listSnowflakeRowAccessPolicies", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW SESSION POLICIES")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var CreatedOn sql.NullString
		var Name sql.NullString
		var DatabaseName sql.NullString
		var SchemaName sql.NullString
		var Kind sql.NullString
		var Owner sql.NullString
		var Comment sql.NullString

		err = rows.Scan(&CreatedOn, &Name, &DatabaseName, &SchemaName, &Kind, &Owner, &Comment)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, SessionPolicy{CreatedOn, Name, DatabaseName, SchemaName, Kind, Owner, Comment})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var CreatedOn sql.NullString
			var Name sql.NullString
			var DatabaseName sql.NullString
			var SchemaName sql.NullString
			var Kind sql.NullString
			var Owner sql.NullString
			var Comment sql.NullString

			err = rows.Scan(&CreatedOn, &Name, &DatabaseName, &SchemaName, &Kind, &Owner, &Comment)
			if err != nil {
				return nil, err
			}

			d.StreamListItem(ctx, SessionPolicy{CreatedOn, Name, DatabaseName, SchemaName, Kind, Owner, Comment})
		}
	}
	return nil, nil
}

func DescribeSessionPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var policy SessionPolicy
	if h.Item != nil {
		policy = h.Item.(SessionPolicy)
	}

	if !policy.Name.Valid {
		return nil, nil
	}

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("snowflake_session_policy.DescribeUser", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE SESSION POLICY %s.%s.%s", policy.DatabaseName.String, policy.SchemaName.String, policy.Name.String))
	if err != nil {
		if err.(*gosnowflake.SnowflakeError) != nil {
			plugin.Logger(ctx).Info("snowflake_session_policy.DescribeUser", fmt.Sprintf("query_error for session policy %s.%s.%s", policy.DatabaseName.String, policy.SchemaName.String, policy.Name.String), err.(*gosnowflake.SnowflakeError).Error())
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	policyProperties := struct {
		SessionIdleTimeoutMins   sql.NullInt64 `json:"session_idle_timeout_mins"`
		SessionUiIdleTimeoutMins sql.NullInt64 `json:"session_ui_idle_timeout_mins"`
	}{}

	for rows.Next() {
		var created_on sql.NullTime
		var name sql.NullString
		var session_idle_timeout_mins sql.NullInt64
		var session_ui_idle_timeout_mins sql.NullInt64
		var comment sql.NullString

		err = rows.Scan(&created_on, &name, &session_idle_timeout_mins, &session_ui_idle_timeout_mins, &comment)
		if err != nil {
			return nil, err
		}
		policyProperties.SessionIdleTimeoutMins = session_idle_timeout_mins
		policyProperties.SessionUiIdleTimeoutMins = session_ui_idle_timeout_mins
	}
	return policyProperties, nil
}
