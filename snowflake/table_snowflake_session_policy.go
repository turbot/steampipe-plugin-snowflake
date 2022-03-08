package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/ddl-user-security.html#label-session-policy-ddl
// This command requires the role executing the command to have:
// 	The OWNERSHIP privilege on the session policy or the APPLY on SESSION POLICY privilege.
// 	The USAGE privilege on the schema.
func tableSessionPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_session_policy",
		Description: "A session policy defines the idle session timeout period in minutes.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeSessionPolicies,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "Identifier for the session policy.", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "database_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "schema_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "kind", Description: "", Type: proto.ColumnType_STRING},
			{Name: "owner", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "", Type: proto.ColumnType_STRING},
			// ADD DESCRIBE SESSION POLICY COLUMNS
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
