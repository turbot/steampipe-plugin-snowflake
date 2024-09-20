package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/user-guide/views-introduction.html
func tableSnowflakeViewGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_view_grant",
		Description: "Lists view-level privileges that have been granted to roles.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeViewGrants,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "view_name"},
				{Name: "database_name"},
				{Name: "schema_name"},
			},
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "view_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("view_name"), Description: "The name of the view."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("database_name"), Description: "The name of the database in which the view exists."},
			{Name: "schema_name", Type: proto.ColumnType_STRING, Transform: transform.FromQual("schema_name"), Description: "The name of the schema in which the view exists."},
			{Name: "privilege", Type: proto.ColumnType_STRING, Description: "A defined level of access to an object."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time privilege was granted."},
			{Name: "grant_option", Type: proto.ColumnType_BOOL, Description: "If set to TRUE, the recipient role can grant the privilege to other roles."},
			{Name: "granted_by", Type: proto.ColumnType_STRING, Description: "Name of the object that granted access on the role."},
			{Name: "granted_on", Type: proto.ColumnType_STRING, Description: "Date and time when the access was granted."},
			{Name: "granted_to", Type: proto.ColumnType_STRING, Description: "Type of the object."},
			{Name: "grantee_name", Type: proto.ColumnType_STRING, Description: "Name of the object role has been granted."},
		}),
	}
}

type ViewGrant AccountGrant

//// LIST FUNCTION

func listSnowflakeViewGrants(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	view := d.EqualsQualString("view_name")
	if view == "" {
		return nil, nil
	}
	database := d.EqualsQualString("database_name")
	if database == "" {
		return nil, nil
	}
	schema := d.EqualsQualString("schema_name")
	if schema == "" {
		return nil, nil
	}
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_view_grant.listSnowflakeViewGrants", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS ON VIEW %s.%s.%s", database, schema, view))
	if err != nil {
		logger.Error("snowflake_view_grant.listSnowflakeViewGrants", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var createdOn sql.NullTime
		var privilege sql.NullString
		var grantedOn sql.NullString
		var name sql.NullString
		var grantedTo sql.NullString
		var granteeName sql.NullString
		var grantOption sql.NullBool
		var grantedBy sql.NullString
		var role sql.NullString

		err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy, &role)
		if err != nil {
			logger.Error("snowflake_view_grant.listSnowflakeViewGrants", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, ViewGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy, role})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var createdOn sql.NullTime
			var privilege sql.NullString
			var grantedOn sql.NullString
			var name sql.NullString
			var grantedTo sql.NullString
			var granteeName sql.NullString
			var grantOption sql.NullBool
			var grantedBy sql.NullString
			var role sql.NullString

			err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy, &role)
			if err != nil {
				logger.Error("snowflake_view_grant.listSnowflakeViewGrants", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, ViewGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy, role})
		}
	}
	return nil, nil
}
