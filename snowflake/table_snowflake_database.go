package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "snowflake_database",
		// https://docs.snowflake.com/en/sql-reference/sql/show-databases.html
		Description: "Snowflake database is a logical grouping of schemas.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeDatabases,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "Name of the database.", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "Creation time of the database.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "is_default", Description: "Name of the default database for authenticating user.", Type: proto.ColumnType_STRING},
			{Name: "is_current", Description: "Name of the current database for authenticating user.", Type: proto.ColumnType_STRING},
			{Name: "origin", Description: "Name of the origin database.", Type: proto.ColumnType_STRING},
			{Name: "owner", Description: "Name of the role that owns the schema.", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "Comment for this database.", Type: proto.ColumnType_STRING},
			{Name: "options", Description: "", Type: proto.ColumnType_STRING},
			{Name: "retention_time", Description: "Number of days that historical data is retained for Time Travel.", Type: proto.ColumnType_STRING},
		},
	}
}

type Database struct {
	CreatedOn     sql.NullString `json:"created_on"`
	Name          sql.NullString `json:"name"`
	IsDefault     sql.NullString `json:"is_default"`
	IsCurrent     sql.NullString `json:"is_current"`
	Origin        sql.NullString `json:"origin"`
	Owner         sql.NullString `json:"owner"`
	Comment       sql.NullString `json:"comment"`
	Options       sql.NullString `json:"options"`
	RetentionTime sql.NullString `json:"retention_time"`
}

//// LIST FUNCTION

func listSnowflakeDatabases(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("aws_region.listAwsRegions", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var CreatedOn sql.NullString
		var Name sql.NullString
		var IsDefault sql.NullString
		var IsCurrent sql.NullString
		var Origin sql.NullString
		var Owner sql.NullString
		var Comment sql.NullString
		var Options sql.NullString
		var RetentionTime sql.NullString

		err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &Origin, &Owner, &Comment, &Options, &RetentionTime)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var CreatedOn sql.NullString
			var Name sql.NullString
			var IsDefault sql.NullString
			var IsCurrent sql.NullString
			var Origin sql.NullString
			var Owner sql.NullString
			var Comment sql.NullString
			var Options sql.NullString
			var RetentionTime sql.NullString

			err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &Origin, &Owner, &Comment, &Options, &RetentionTime)
			if err != nil {
				return nil, err
			}

			d.StreamListItem(ctx, Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime})
		}
	}
	return nil, nil
}
