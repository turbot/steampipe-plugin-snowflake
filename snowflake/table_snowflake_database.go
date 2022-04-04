package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/sql/show-databases.html
func tableSnowflakeDatabase(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_database",
		Description: "Snowflake database is a logical grouping of schemas.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeDatabases,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the database."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Creation time of the database."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for this database."},
			{Name: "is_current", Type: proto.ColumnType_STRING, Description: "Name of the current database for authenticating user."},
			{Name: "is_default", Type: proto.ColumnType_STRING, Description: "Name of the default database for authenticating user."},
			{Name: "options", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "origin", Type: proto.ColumnType_STRING, Description: "Name of the origin database."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Name of the role that owns the schema."},
			{Name: "retention_time", Type: proto.ColumnType_INT, Description: "Number of days that historical data is retained for Time Travel."},
		}),
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
		logger.Error("snowflake_database.listSnowflakeDatabases", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW DATABASES")
	if err != nil {
		logger.Error("snowflake_database.listSnowflakeDatabases", "query.error", err)
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
			logger.Error("snowflake_database.listSnowflakeDatabases", "query.error", err)
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
				logger.Error("snowflake_database.listSnowflakeDatabases", "query.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, Database{CreatedOn, Name, IsDefault, IsCurrent, Origin, Owner, Comment, Options, RetentionTime})
		}
	}
	return nil, nil
}
