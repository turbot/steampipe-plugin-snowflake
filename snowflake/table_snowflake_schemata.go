package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

// https://docs.snowflake.com/en/sql-reference/info-schema/schemata.html
func tableSnowflakeSchemata(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_schemata",
		Description: "This Information Schema view displays a row for each schema in the specified (or current) database, including the INFORMATION_SCHEMA schema itself.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeSchemata,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "schema_id", Type: proto.ColumnType_STRING, Description: "ID of the schema."},
			{Name: "schema_name", Type: proto.ColumnType_STRING, Description: "Name of the schema."},
			{Name: "catalog_id", Type: proto.ColumnType_STRING, Description: "ID of the database that the schema belongs to."},
			{Name: "catalog_name", Type: proto.ColumnType_STRING, Description: "Database that the schema belongs to."},
			{Name: "schema_owner", Type: proto.ColumnType_STRING, Description: "Name of the role that owns the schema."},
			{Name: "retention_time", Type: proto.ColumnType_INT, Description: "Number of days that historical data is retained for Time Travel."},
			{Name: "is_transient", Type: proto.ColumnType_STRING, Description: "Whether this is a transient schema."},
			{Name: "is_managed_access", Type: proto.ColumnType_STRING, Description: "Whether the schema is a managed access schema."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for this schema."},
			{Name: "created", Type: proto.ColumnType_TIMESTAMP, Description: "Creation time of the schema."},
			{Name: "last_altered", Type: proto.ColumnType_TIMESTAMP, Description: "Last altered time of the schema."},
			{Name: "deleted", Type: proto.ColumnType_TIMESTAMP, Description: "Deletion time of the schema."},
		}),
	}
}

type Schemata struct {
	SchemaId                   sql.NullInt64  `json:"SCHEMA_ID"`
	SchemaName                 sql.NullString `json:"SCHEMA_NAME"`
	CatalogId                  sql.NullInt64  `json:"CATALOG_ID"`
	CatalogName                sql.NullString `json:"CATALOG_NAME"`
	SchemaOwner                sql.NullString `json:"SCHEMA_OWNER"`
	RetentionTime              sql.NullInt64  `json:"RETENTION_TIME"`
	IsTransient                sql.NullString `json:"IS_TRANSIENT"`
	IsManagedAccess            sql.NullString `json:"IS_MANAGED_ACCESS"`
	DefaultCharacterSetCatalog sql.NullString `json:"DEFAULT_CHARACTER_SET_CATALOG"`
	DefaultCharacterSetSchema  sql.NullString `json:"DEFAULT_CHARACTER_SET_SCHEMA"`
	DefaultCharacterSetName    sql.NullString `json:"DEFAULT_CHARACTER_SET_NAME"`
	SqlPath                    sql.NullString `json:"SQL_PATH"`
	Comment                    sql.NullString `json:"COMMENT"`
	Created                    sql.NullTime   `json:"CREATED"`
	LastAltered                sql.NullTime   `json:"LAST_ALTERED"`
	Deleted                    sql.NullTime   `json:"DELETED"`
}

// SchemataCol returns a reference for a column of a Schemata
func SchemataCol(colname string, item *Schemata) interface{} {
	switch colname {
	case "SCHEMA_ID":
		return &item.SchemaId
	case "SCHEMA_NAME":
		return &item.SchemaName
	case "CATALOG_ID":
		return &item.CatalogId
	case "CATALOG_NAME":
		return &item.CatalogName
	case "SCHEMA_OWNER":
		return &item.SchemaOwner
	case "RETENTION_TIME":
		return &item.RetentionTime
	case "IS_TRANSIENT":
		return &item.IsTransient
	case "IS_MANAGED_ACCESS":
		return &item.IsManagedAccess
	case "DEFAULT_CHARACTER_SET_CATALOG":
		return &item.DefaultCharacterSetCatalog
	case "DEFAULT_CHARACTER_SET_SCHEMA":
		return &item.DefaultCharacterSetSchema
	case "DEFAULT_CHARACTER_SET_NAME":
		return &item.DefaultCharacterSetName
	case "SQL_PATH":
		return &item.SqlPath
	case "COMMENT":
		return &item.Comment
	case "CREATED":
		return &item.Created
	case "LAST_ALTERED":
		return &item.LastAltered
	case "DELETED":
		return &item.Deleted
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeSchemata(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_view_schemata.listSnowflakeSchemata", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "select * from SNOWFLAKE.ACCOUNT_USAGE.schemata;")
	if err != nil {
		logger.Error("snowflake_view_schemata.listSnowflakeSchemata", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_view_schemata.listSnowflakeSchemata", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		schema := Schemata{}
		// make references for the cols with the aid of UserGrantCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = SchemataCol(col, &schema)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_view_schemata.listSnowflakeSchemata", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, schema)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			schema := Schemata{}
			// make references for the cols with the aid of UserGrantCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = SchemataCol(col, &schema)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_view_schemata.listSnowflakeSchemata", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, schema)
		}
	}
	return nil, nil
}
