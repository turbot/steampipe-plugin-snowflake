package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableRowAccessPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_row_access_policy",
		Description: "Snowflake Row Access Policy",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeRowAccessPolicies,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "database_name", Description: "The database for the row access policy.", Type: proto.ColumnType_STRING},
			{Name: "schema_name", Description: "The schema in database for the row access policy", Type: proto.ColumnType_STRING},
			{Name: "kind", Description: "", Type: proto.ColumnType_STRING},
			{Name: "owner", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "", Type: proto.ColumnType_STRING},
		},
	}
}

type Policy struct {
	CreatedOn    sql.NullString `db:"created_on"`
	Name         sql.NullString `db:"name"`
	DatabaseName sql.NullString `db:"database_name"`
	SchemaName   sql.NullString `db:"schema_name"`
	Kind         sql.NullString `db:"kind"`
	Owner        sql.NullString `db:"owner"`
	Comment      sql.NullString `db:"comment"`
}

type RowAccessPolicy Policy

//// LIST FUNCTION

func listSnowflakeRowAccessPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_row_access_policy.listSnowflakeRowAccessPolicies", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW ROW ACCESS POLICIES")
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

		d.StreamListItem(ctx, RowAccessPolicy{CreatedOn, Name, DatabaseName, SchemaName, Kind, Owner, Comment})
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

			d.StreamListItem(ctx, RowAccessPolicy{CreatedOn, Name, DatabaseName, SchemaName, Kind, Owner, Comment})
		}
	}
	return nil, nil
}
