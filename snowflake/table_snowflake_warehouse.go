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
		Description: "Snowflake Warehouse",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeWarehouses,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "state", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "type", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "size", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "running", Type: proto.ColumnType_INT, Description: ""},
			{Name: "queued", Type: proto.ColumnType_INT, Description: ""},
			{Name: "is_default", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "is_current", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "auto_suspend", Type: proto.ColumnType_INT, Description: ""},
			{Name: "auto_resume", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "available", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "provisioning", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "quiescing", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "other", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "resumed_on", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "updated_on", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "resource_monitor", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "actives", Type: proto.ColumnType_INT, Description: ""},
			{Name: "pendings", Type: proto.ColumnType_INT, Description: ""},
			{Name: "failed", Type: proto.ColumnType_INT, Description: ""},
			{Name: "suspended", Type: proto.ColumnType_INT, Description: ""},
			{Name: "uuid", Type: proto.ColumnType_STRING, Description: ""},
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
	logger.Error("aws_region.listSnowflakeRole", "api.error", "nil")
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("aws_region.listSnowflakeRole", "connnection.error", err)
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
