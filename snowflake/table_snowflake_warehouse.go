package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeWarehouse(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_warehouse",
		Description: "A warehouse, is a cluster of compute resources in Snowflake. Warehouse provides the required resources, such as CPU, memory, and temporary storage, to perform queries.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeWarehouses,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name for warehouse."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "Whether the warehouse is active/running (STARTED), inactive (SUSPENDED), or resizing (RESIZING)."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Warehouse type; STANDARD is the only currently supported type."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "Size of the warehouse (X-Small, Small, Medium, Large, X-Large, etc.)"},
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
		},
	}
}

type Warehouse struct {
	Name            sql.NullString `json:"name"`
	State           sql.NullString `json:"state"`
	Type            sql.NullString `json:"type"`
	Size            sql.NullString `json:"size"`
	Running         sql.NullInt64  `json:"running"`
	Queued          sql.NullInt64  `json:"queued"`
	IsDefault       sql.NullString `json:"is_default"`
	IsCurrent       sql.NullString `json:"is_current"`
	AutoSuspend     sql.NullInt64  `json:"auto_suspend"`
	AutoResume      sql.NullBool   `json:"auto_resume"`
	Available       sql.NullString `json:"available"`
	Provisioning    sql.NullString `json:"provisioning"`
	Quiescing       sql.NullString `json:"quiescing"`
	Other           sql.NullString `json:"other"`
	CreatedOn       sql.NullTime   `json:"created_on"`
	ResumedOn       sql.NullTime   `json:"resumed_on"`
	UpdatedOn       sql.NullTime   `json:"updated_on"`
	Owner           sql.NullString `json:"owner"`
	Comment         sql.NullString `json:"comment"`
	ResourceMonitor sql.NullString `json:"resource_monitor"`
	Actives         sql.NullInt64  `json:"actives"`
	Pendings        sql.NullInt64  `json:"pendings"`
	Failed          sql.NullInt64  `json:"failed"`
	Suspended       sql.NullInt64  `json:"suspended"`
	UUID            sql.NullString `json:"uuid"`
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
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Name sql.NullString
		var State sql.NullString
		var Type sql.NullString
		var Size sql.NullString
		var Running sql.NullInt64
		var Queued sql.NullInt64
		var IsDefault sql.NullString
		var IsCurrent sql.NullString
		var AutoSuspend sql.NullInt64
		var AutoResume sql.NullBool
		var Available sql.NullString
		var Provisioning sql.NullString
		var Quiescing sql.NullString
		var Other sql.NullString
		var CreatedOn sql.NullTime
		var ResumedOn sql.NullTime
		var UpdatedOn sql.NullTime
		var Owner sql.NullString
		var Comment sql.NullString
		var ResourceMonitor sql.NullString
		var Actives sql.NullInt64
		var Pendings sql.NullInt64
		var Failed sql.NullInt64
		var Suspended sql.NullInt64
		var UUID sql.NullString

		err = rows.Scan(&Name, &State, &Type, &Size, &Running, &Queued, &IsDefault, &IsCurrent, &AutoSuspend, &AutoResume, &Available, &Provisioning, &Quiescing, &Other, &CreatedOn, &ResumedOn, &UpdatedOn, &Owner, &Comment, &ResourceMonitor, &Actives, &Pendings, &Failed, &Suspended, &UUID)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, Warehouse{Name, State, Type, Size, Running, Queued, IsDefault, IsCurrent, AutoSuspend, AutoResume, Available, Provisioning, Quiescing, Other, CreatedOn, ResumedOn, UpdatedOn, Owner, Comment, ResourceMonitor, Actives, Pendings, Failed, Suspended, UUID})
	}

	for rows.NextResultSet() {
		for rows.Next() {
			var Name sql.NullString
			var State sql.NullString
			var Type sql.NullString
			var Size sql.NullString
			var Running sql.NullInt64
			var Queued sql.NullInt64
			var IsDefault sql.NullString
			var IsCurrent sql.NullString
			var AutoSuspend sql.NullInt64
			var AutoResume sql.NullBool
			var Available sql.NullString
			var Provisioning sql.NullString
			var Quiescing sql.NullString
			var Other sql.NullString
			var CreatedOn sql.NullTime
			var ResumedOn sql.NullTime
			var UpdatedOn sql.NullTime
			var Owner sql.NullString
			var Comment sql.NullString
			var ResourceMonitor sql.NullString
			var Actives sql.NullInt64
			var Pendings sql.NullInt64
			var Failed sql.NullInt64
			var Suspended sql.NullInt64
			var UUID sql.NullString

			err = rows.Scan(&Name, &State, &Type, &Size, &Running, &Queued, &IsDefault, &IsCurrent, &AutoSuspend, &AutoResume, &Available, &Provisioning, &Quiescing, &Other, &CreatedOn, &ResumedOn, &UpdatedOn, &Owner, &Comment, &ResourceMonitor, &Actives, &Pendings, &Failed, &Suspended, &UUID)
			if err != nil {
				return nil, err
			}

			d.StreamListItem(ctx, Warehouse{Name, State, Type, Size, Running, Queued, IsDefault, IsCurrent, AutoSuspend, AutoResume, Available, Provisioning, Quiescing, Other, CreatedOn, ResumedOn, UpdatedOn, Owner, Comment, ResourceMonitor, Actives, Pendings, Failed, Suspended, UUID})
		}
	}
	return nil, nil
}
