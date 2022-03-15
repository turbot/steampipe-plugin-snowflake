package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/user-guide/views-introduction.html
func tableSnowflakeView(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_view",
		Description: "Snowflake view is basically a named definition of a query.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeViews,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the view."},
			{Name: "database_name", Type: proto.ColumnType_STRING, Description: "The name of the database in which the view exists."},
			{Name: "schema_name", Type: proto.ColumnType_STRING, Description: "The name of the schema in which the view exists."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp at which the view was created."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "The owner of the view."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Optional comment."},
			{Name: "text", Type: proto.ColumnType_STRING, Description: "The text of the command that created the view, e.g., CREATE VIEW."},
			{Name: "is_secure", Type: proto.ColumnType_BOOL, Description: "True if the view is a secure view; false otherwise."},
			{Name: "is_materialized", Type: proto.ColumnType_BOOL, Description: "True if the view is a materialized view; false otherwise."},
		},
	}
}

type View struct {
	CreatedOn      sql.NullTime   `json:"created_on"`
	Name           sql.NullString `json:"name"`
	Reserved       sql.NullString `json:"reserved"`
	DatabaseName   sql.NullString `json:"database_name"`
	SchemaName     sql.NullString `json:"schema_name"`
	Owner          sql.NullString `json:"owner"`
	Comment        sql.NullString `json:"comment"`
	Text           sql.NullString `json:"text"`
	IsSecure       sql.NullString `json:"is_secure"`
	IsMaterialized sql.NullString `json:"is_materialized"`
}

//// LIST FUNCTION

func listSnowflakeViews(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_view.listSnowflakeViews", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW VIEWS")
	if err != nil {
		logger.Error("snowflake_view.listSnowflakeViews", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var createdOn sql.NullTime
		var name sql.NullString
		var reserved sql.NullString
		var databaseName sql.NullString
		var schemaName sql.NullString
		var owner sql.NullString
		var comment sql.NullString
		var text sql.NullString
		var isSecure sql.NullString
		var isMaterialized sql.NullString

		err = rows.Scan(&createdOn, &name, &reserved, &databaseName, &schemaName, &owner, &comment, &text, &isSecure, &isMaterialized)
		if err != nil {
			logger.Error("snowflake_view.listSnowflakeViews", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, View{createdOn, name, reserved, databaseName, schemaName, owner, comment, text, isSecure, isMaterialized})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var createdOn sql.NullTime
			var name sql.NullString
			var reserved sql.NullString
			var databaseName sql.NullString
			var schemaName sql.NullString
			var owner sql.NullString
			var comment sql.NullString
			var text sql.NullString
			var isSecure sql.NullString
			var isMaterialized sql.NullString
			err = rows.Scan(&createdOn, &name, &reserved, &databaseName, &schemaName, &owner, &comment, &text, &isSecure, &isMaterialized)
			if err != nil {
				logger.Error("snowflake_view.listSnowflakeViews", "query_scan.error", err)
				return nil, err
			}
			d.StreamListItem(ctx, View{createdOn, name, reserved, databaseName, schemaName, owner, comment, text, isSecure, isMaterialized})
		}
	}
	return nil, nil
}
