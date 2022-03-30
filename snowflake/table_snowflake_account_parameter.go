package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION
func tableSnowflakeAccountParameter(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_account_parameter",
		Description: "Lists all the account-level parameters.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeAccountParameters,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "key", Type: proto.ColumnType_STRING, Description: "Name of the account parameter."},
			{Name: "value", Type: proto.ColumnType_STRING, Description: "Current value of the parameter."},
			{Name: "default", Type: proto.ColumnType_STRING, Description: "Default value of the parameter."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description for the parameter."},
			{Name: "level", Type: proto.ColumnType_STRING, Description: "Level of the parameter.Can be one of SYSTEM or ACCOUNT."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Data type of the parameter value."},
		}),
	}
}

type Parameter struct {
	Key         sql.NullString `json:"key"`
	Value       sql.NullString `json:"value"`
	Default     sql.NullString `json:"default"`
	Level       sql.NullString `json:"level"`
	Description sql.NullString `json:"description"`
	Type        sql.NullString `json:"type"`
}

// ParameterCol returns a reference for a column of a Parameter
func ParameterCol(colname string, item *Parameter) interface{} {
	switch colname {
	case "key":
		return &item.Key
	case "value":
		return &item.Value
	case "default":
		return &item.Default
	case "level":
		return &item.Level
	case "description":
		return &item.Description
	case "type":
		return &item.Type
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeAccountParameters(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_account_parameters.listSnowflakeAccountParameters", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW PARAMETERS IN ACCOUNT;")
	if err != nil {
		logger.Error("snowflake_account_parameters.listSnowflakeAccountParameters", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_account_parameters.listSnowflakeAccountParameters", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		parameter := Parameter{}
		// make references for the cols with the aid of ParameterCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = ParameterCol(col, &parameter)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_account_parameters.listSnowflakeAccountParameters", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, parameter)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			parameter := Parameter{}
			// make references for the cols with the aid of ParameterCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = ParameterCol(col, &parameter)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_account_parameters.listSnowflakeAccountParameters", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, parameter)
		}
	}
	return nil, nil
}
