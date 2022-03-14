package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/sql/show-databases.html
func tableSnowflakeDatabaseGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_database_grant",
		Description: "Lists all privileges that have been granted on the database.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeDatabaseGrants,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "database"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name").Transform(valueFromNullable), Description: "Name of the database."},
			{Name: "privilege", Type: proto.ColumnType_STRING, Description: "A defined level of access to an database."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the access was granted."},
			{Name: "grant_option", Type: proto.ColumnType_BOOL, Description: "If set to TRUE, the recipient role can grant the privilege to other roles."},
			{Name: "granted_by", Type: proto.ColumnType_STRING, Description: "Identifier for the object that granted the privilege."},
			{Name: "granted_on", Type: proto.ColumnType_STRING, Description: "Type of the object."},
			{Name: "granted_to", Type: proto.ColumnType_STRING, Description: "Type of the object role has been granted."},
			{Name: "grantee_name", Type: proto.ColumnType_STRING, Description: "Name of the object role has been granted."},
		},
	}
}

type DatabaseGrant AccountGrant

//// LIST FUNCTION

func listSnowflakeDatabaseGrants(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	database := d.KeyColumnQualString("database")
	if database == "" {
		return nil, nil
	}

	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_database_grant.listSnowflakeDatabaseGrants", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS ON DATABASE %s", database))
	if err != nil {
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
		var grantOption sql.NullString
		var grantedBy sql.NullString

		err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, DatabaseGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var createdOn sql.NullTime
			var privilege sql.NullString
			var grantedOn sql.NullString
			var name sql.NullString
			var grantedTo sql.NullString
			var granteeName sql.NullString
			var grantOption sql.NullString
			var grantedBy sql.NullString

			err = rows.Scan(&createdOn, &privilege, &grantedOn, &name, &grantedTo, &granteeName, &grantOption, &grantedBy)
			if err != nil {
				return nil, err
			}

			d.StreamListItem(ctx, DatabaseGrant{createdOn, privilege, grantedOn, name, grantedTo, granteeName, grantOption, grantedBy})
		}
	}
	return nil, nil
}
