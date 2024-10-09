package snowflake

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeWarehouse(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_warehouse",
		Description: "A warehouse, is a cluster of compute resources in Snowflake. Warehouse provides the required resources, such as CPU, memory, and temporary storage, to perform queries.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeWarehouses,
		},
		Columns: snowflakeColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name for warehouse."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "Whether the warehouse is active/running (STARTED), inactive (SUSPENDED), or resizing (RESIZING)."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Warehouse type; STANDARD is the only currently supported type."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "Size of the warehouse (X-Small, Small, Medium, Large, X-Large, etc.)"},
			{Name: "min_cluster_count", Type: proto.ColumnType_INT, Description: "Minimum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses)."},
			{Name: "max_cluster_count", Type: proto.ColumnType_INT, Description: "Maximum number of warehouses for the (multi-cluster) warehouse (always 1 for single warehouses)."},
			{Name: "started_clusters", Type: proto.ColumnType_INT, Description: "Number of warehouses currently started."},
			{Name: "running", Type: proto.ColumnType_INT, Description: "Number of SQL statements that are being executed by the warehouse."},
			{Name: "queued", Type: proto.ColumnType_INT, Description: "Number of SQL statements that are queued for the warehouse."},
			{Name: "is_default", Type: proto.ColumnType_STRING, Description: "Whether the warehouse is the default for the current user."},
			{Name: "is_current", Type: proto.ColumnType_STRING, Description: "Whether the warehouse is in use for the session."},
			{Name: "auto_suspend", Type: proto.ColumnType_INT, Description: "Specifies the number of seconds of inactivity after which a warehouse is automatically suspended."},
			{Name: "auto_resume", Type: proto.ColumnType_BOOL, Description: "Specifies whether to automatically resume a warehouse when a SQL statement (e.g. query) is submitted to it."},
			{Name: "available", Type: proto.ColumnType_STRING, Description: "Percentage of the warehouse compute resources that are provisioned and available."},
			{Name: "provisioning", Type: proto.ColumnType_STRING, Description: "Percentage of the warehouse compute resources that are in the process of provisioning."},
			{Name: "quiescing", Type: proto.ColumnType_STRING, Description: "Percentage of the warehouse compute resources that are executing SQL statements, but will be shut down once the queries complete."},
			{Name: "other", Type: proto.ColumnType_STRING, Description: "Percentage of the warehouse compute resources that are in a state other than available, provisioning, or quiescing."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the warehouse was created."},
			{Name: "resumed_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the warehouse was last started or restarted."},
			{Name: "updated_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the warehouse was last updated, which includes changing any of the properties of the warehouse or changing the state (STARTED, SUSPENDED, RESIZING) of the warehouse."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Role that owns the warehouse."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for the warehouse."},
			{Name: "resource_monitor", Type: proto.ColumnType_STRING, Description: "ID of resource monitor explicitly assigned to the warehouse; controls the monthly credit usage for the warehouse."},
			{Name: "scaling_policy", Type: proto.ColumnType_STRING, Description: "Policy that determines when additional warehouses (in a multi-cluster warehouse) are automatically started and shut down."},
		}),
	}
}

type Warehouse struct {
	Name            sql.NullString `json:"name" db:"name"`
	State           sql.NullString `json:"state" db:"state"`
	Type            sql.NullString `json:"type" db:"type"`
	Size            sql.NullString `json:"size" db:"size"`
	MinClusterCount sql.NullInt64  `json:"min_cluster_count" db:"min_cluster_count"`
	MaxClusterCount sql.NullInt64  `json:"max_cluster_count" db:"max_cluster_count"`
	StartedClusters sql.NullInt64  `json:"started_clusters" db:"started_clusters"`
	Running         sql.NullInt64  `json:"running" db:"running"`
	Queued          sql.NullInt64  `json:"queued" db:"queued"`
	IsDefault       sql.NullString `json:"is_default" db:"is_default"`
	IsCurrent       sql.NullString `json:"is_current" db:"is_current"`
	AutoSuspend     sql.NullInt64  `json:"auto_suspend" db:"auto_suspend"`
	AutoResume      sql.NullBool   `json:"auto_resume" db:"auto_resume"`
	Available       sql.NullString `json:"available" db:"available"`
	Provisioning    sql.NullString `json:"provisioning" db:"provisioning"`
	Quiescing       sql.NullString `json:"quiescing" db:"quiescing"`
	Other           sql.NullString `json:"other" db:"other"`
	CreatedOn       sql.NullTime   `json:"created_on" db:"created_on"`
	ResumedOn       sql.NullTime   `json:"resumed_on" db:"resumed_on"`
	UpdatedOn       sql.NullTime   `json:"updated_on" db:"updated_on"`
	Owner           sql.NullString `json:"owner" db:"owner"`
	Comment         sql.NullString `json:"comment" db:"comment"`
	ResourceMonitor sql.NullString `json:"resource_monitor" db:"resource_monitor"`
	Actives         sql.NullInt64  `json:"actives" db:"actives"`
	Pendings        sql.NullInt64  `json:"pendings" db:"pendings"`
	Failed          sql.NullInt64  `json:"failed" db:"failed"`
	Suspended       sql.NullInt64  `json:"suspended" db:"suspended"`
	UUID            sql.NullString `json:"uuid" db:"uuid"`
	ScalingPolicy   sql.NullString `json:"scaling_policy" db:"scaling_policy"`
	Budget          sql.NullString `json:"budget" db:"budget"`
	OwnerRoleType   sql.NullString `json:"owner_role_type" db:"owner_role_type"`
}

//// LIST FUNCTION

func listSnowflakeWarehouses(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_warehouse.listSnowflakeWarehouses", "connnection.error", err)
		return nil, err
	}

	rows, err := db.QueryContext(ctx, "SHOW WAREHOUSES")
	if err != nil {
		logger.Error("snowflake_warehouse.listSnowflakeWarehouses", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	dbs := []Warehouse{}

	err = sqlx.StructScan(rows, &dbs)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("snowflake_warehouse.listSnowflakeWarehouses", "no warehouses found")
			return nil, nil
		}
		logger.Error("snowflake_warehouse.listSnowflakeWarehouses", "struct_scan.error", err)
		return nil, err
	}

	for _, warehouse := range dbs {
		d.StreamListItem(ctx, warehouse)
	}
	return nil, nil
}

func ScanWarehouse(row *sqlx.Row) (*Warehouse, error) {
	w := &Warehouse{}
	err := row.StructScan(w)
	return w, err
}
