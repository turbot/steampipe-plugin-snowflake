package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeWarehouseMeteringHistory(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_warehouse_metering_history",
		Description: "This Account Usage view can be used to return the hourly credit usage for a single warehouse (or all the warehouses in your account) within the last 365 days (1 year).",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeWarehouseMeteringHistory,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "start_time", Require: plugin.Optional, Operators: []string{">="}},
				{Name: "end_time", Require: plugin.Optional, Operators: []string{"<="}},
			},
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "The beginning of the hour in which this warehouse usage took place."},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "The end of the hour in which this warehouse usage took place."},
			{Name: "warehouse_name", Type: proto.ColumnType_STRING, Description: "Name of the warehouse."},
			{Name: "credits_used", Type: proto.ColumnType_DOUBLE, Description: "Number of credits billed for this warehouse in this hour."},
			{Name: "credits_used_compute", Type: proto.ColumnType_DOUBLE, Description: "Number of credits used for the warehouse in the hour."},
			{Name: "credits_used_cloud_services", Type: proto.ColumnType_DOUBLE, Description: "Number of credits used for cloud services in the hour."},
		}),
	}
}

type WarehouseMeteringHistory struct {
	StartTime                sql.NullTime   `json:"START_TIME"`
	EndTime                  sql.NullTime   `json:"END_TIME"`
	WarehouseName            sql.NullString `json:"WAREHOUSE_NAME"`
	CreditsUsed              sql.NullString `json:"CREDITS_USED"`
	CreditsUsedCompute       sql.NullString `json:"CREDITS_USED_COMPUTE"`
	CreditsUsedCloudServices sql.NullString `json:"CREDITS_USED_CLOUD_SERVICES"`
}

//// LIST FUNCTION

func listSnowflakeWarehouseMeteringHistory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_warehouse_metering_history.listSnowflakeWarehouseMeteringHistory", "connnection.error", err)
		return nil, err
	}

	// var st, et string

	// if d.EqualsQuals["start_time"] != nil {
	// 	st = d.EqualsQuals["start_time"].GetTimestampValue().AsTime().Format("2006-01-02 15:04:05")
	// }
	// if d.EqualsQuals["end_time"] != nil {
	// 	et = d.EqualsQuals["end_time"].GetTimestampValue().AsTime().Format("2006-01-02 15:04:05.000")
	// }

	// query := fmt.Sprintf("SELECT * FROM TABLE ( SNOWFLAKE.INFORMATION_SCHEMA.WAREHOUSE_METERING_HISTORY('%s','%s'));", st, et)
	query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.WAREHOUSE_METERING_HISTORY"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		logger.Error("snowflake_warehouse_metering_history.listSnowflakeWarehouseMeteringHistory", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_warehouse_metering_history.listSnowflakeWarehouseMeteringHistory", "get_columns.error", err)
		return nil, err
	}

	for rows.Next() {
		WarehouseMeteringHistory := WarehouseMeteringHistory{}
		// make references for the cols with the aid of WarehouseMeteringHistoryCol
		cols := make([]interface{}, len(columns))
		for i, col := range columns {
			cols[i] = WarehouseMeteringHistoryCol(col, &WarehouseMeteringHistory)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_warehouse_metering_history.listSnowflakeWarehouseMeteringHistory", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, WarehouseMeteringHistory)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			WarehouseMeteringHistory := WarehouseMeteringHistory{}
			// make references for the cols with the aid of WarehouseMeteringHistoryCol
			cols := make([]interface{}, len(columns))
			for i, col := range columns {
				cols[i] = WarehouseMeteringHistoryCol(col, &WarehouseMeteringHistory)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_warehouse_metering_history.listSnowflakeWarehouseMeteringHistory", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, WarehouseMeteringHistory)
		}
	}
	return nil, nil
}

// WarehouseMeteringHistoryCol returns a reference for a column of a WarehouseMeteringHistory
func WarehouseMeteringHistoryCol(colname string, item *WarehouseMeteringHistory) interface{} {
	switch colname {
	case "START_TIME":
		return &item.StartTime
	case "END_TIME":
		return &item.EndTime
	case "WAREHOUSE_NAME":
		return &item.WarehouseName
	case "CREDITS_USED":
		return &item.CreditsUsed
	case "CREDITS_USED_COMPUTE":
		return &item.CreditsUsedCompute
	case "CREDITS_USED_CLOUD_SERVICES":
		return &item.CreditsUsedCloudServices
	default:
		panic("unknown column " + colname)
	}
}
