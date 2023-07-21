 package snowflake

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
				{Name: "warehouse_id", Require: plugin.Optional},
				{Name: "warehouse_name", Require: plugin.Optional},
			},
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "The beginning of the hour in which this warehouse usage took place."},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "The end of the hour in which this warehouse usage took place."},
			{Name: "warehouse_name", Type: proto.ColumnType_STRING, Description: "Name of the warehouse."},
			{Name: "warehouse_id", Type: proto.ColumnType_STRING, Description: "Internal/system-generated identifier for the warehouse."},
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
	WarehouseId              sql.NullString `json:"WAREHOUSE_Id"`
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

	conditions := []string{}
	if d.EqualsQualString("warehouse_name") != "" {
		conditions = append(conditions, fmt.Sprintf("warehouse_name %s '%s'", "=", d.EqualsQualString("warehouse_name")))
	}
	if d.EqualsQualString("warehouse_id") != "" {
		conditions = append(conditions, fmt.Sprintf("warehouse_id %s '%s'", "=", d.EqualsQualString("warehouse_id")))
	}

	condition := strings.Join(conditions, " and ")
	query := "SELECT * FROM SNOWFLAKE.ACCOUNT_USAGE.WAREHOUSE_METERING_HISTORY"
	if condition != "" {
		query = fmt.Sprintf("%s where %s", query, condition)
	}

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
	case "WAREHOUSE_ID":
		return &item.WarehouseId
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
