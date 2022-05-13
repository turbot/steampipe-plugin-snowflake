package snowflake

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeResourceMonitor(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_resource_monitor",
		Description: "Lists all the resource monitors in your account for which you have access privileges.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeResourceMonitors,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name for warehouse."},
			{Name: "credit_quota", Type: proto.ColumnType_DOUBLE, Description: "Monthly credit quota for the resource monitor."},
			{Name: "used_credits", Type: proto.ColumnType_DOUBLE, Description: "Number of credits used in the current monthly billing cycle by all the warehouses associated with the resource monitor."},
			{Name: "remaining_credits", Type: proto.ColumnType_DOUBLE, Description: "Number of credits still available to use in the current monthly billing cycle."},
			{Name: "level", Type: proto.ColumnType_STRING, Description: "Level"},
			{Name: "frequency", Type: proto.ColumnType_STRING, Description: "Daily, Weekly, etc"},
			{Name: "start_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the monitor was started."},
			{Name: "end_time", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the monitor was stopped."},
			{Name: "notify_at", Type: proto.ColumnType_STRING, Description: "Levels to which to alert."},
			{Name: "suspend_at", Type: proto.ColumnType_STRING, Description: "Levels to which to suspend warehouse"},
			{Name: "suspend_immediately_at", Type: proto.ColumnType_STRING, Description: "Levels to which to suspend warehouse"},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the monitor was created."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Role that owns the warehouse."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for the warehouse."},
			{Name: "notify_users", Type: proto.ColumnType_STRING, Description: "Who to notify when alerting."},
		}),
	}
}

type ResourceMonitor struct {
	Name                 sql.NullString `json:"name" db:"name"`
	CreditQuota          sql.NullString `json:"credit_quota" db:"credit_quota"`
	UsedCredits          sql.NullString `json:"used_credits" db:"used_credits"`
	RemainingCredits     sql.NullString `json:"remaining_credits" db:"remaining_credits"`
	Level                sql.NullString `json:"level" db:"level"`
	Frequency            sql.NullString `json:"frequency" db:"frequency"`
	StartTime            sql.NullTime   `json:"start_time" db:"start_time"`
	EndTime              sql.NullTime   `json:"end_time" db:"end_time"`
	NotifyAt             sql.NullString `json:"notify_at" db:"notify_at"`
	SuspendAt            sql.NullString `json:"suspend_at" db:"suspend_at"`
	SuspendImmediatelyAt sql.NullString `json:"suspend_immediately_at" db:"suspend_immediately_at"`
	CreatedOn            sql.NullTime   `json:"created_on" db:"created_on"`
	Owner                sql.NullString `json:"owner" db:"owner"`
	Comment              sql.NullString `json:"comment" db:"comment"`
	NotifyUsers          sql.NullString `json:"notify_users" db:"notify_users"`
}

//// LIST FUNCTION

func listSnowflakeResourceMonitors(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_resource_monitor.listSnowflakeResourceMonitors", "connnection.error", err)
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SHOW RESOURCE MONITORS")
	if err != nil {
		logger.Error("snowflake_resource_monitor.listSnowflakeResourceMonitors", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	resourceMonitors := []ResourceMonitor{}

	err = sqlx.StructScan(rows, &resourceMonitors)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("snowflake_resource_monitor.listSnowflakeResourceMonitors", "no monitors found")
			return nil, nil
		}
		logger.Error("snowflake_resource_monitor.listSnowflakeResourceMonitors", "struct_scan.error", err)
		return nil, err
	}

	for _, resourceMonitor := range resourceMonitors {
		d.StreamListItem(ctx, resourceMonitor)
	}
	return nil, nil
}
