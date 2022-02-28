package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_role",
		Description: "Snowflake Role",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeRole,
		},
		Columns: []*plugin.Column{
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "is_default", Description: "", Type: proto.ColumnType_STRING},
			{Name: "is_current", Description: "", Type: proto.ColumnType_STRING},
			{Name: "is_inherited", Description: "", Type: proto.ColumnType_STRING},
			{Name: "assigned_to_users", Description: "", Type: proto.ColumnType_INT},
			{Name: "owner", Description: "", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "", Type: proto.ColumnType_STRING},
			{Name: "granted_to_roles", Description: "", Type: proto.ColumnType_INT},
			{Name: "granted_roles", Description: "", Type: proto.ColumnType_INT},
		},
	}
}

type Role struct {
	CreatedOn       sql.NullTime   `json:"created_on"`
	Name            sql.NullString `json:"name"`
	IsDefault       sql.NullString `json:"is_default"`
	IsCurrent       sql.NullString `json:"is_current"`
	IsInherited     sql.NullString `json:"is_inherited"`
	AssignedToUsers sql.NullInt64  `json:"assigned_to_users"`
	GrantedToRoles  sql.NullInt64  `json:"granted_to_roles"`
	GrantedRoles    sql.NullInt64  `json:"granted_roles"`
	Owner           sql.NullString `json:"owner"`
	Comment         sql.NullString `json:"comment"`
}

//// LIST FUNCTION

func listSnowflakeRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Error("aws_region.listSnowflakeRole", "api.error", "nil")
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("aws_region.listSnowflakeRole", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW ROLES")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var CreatedOn sql.NullTime
		var Name sql.NullString
		var IsDefault sql.NullString
		var IsCurrent sql.NullString
		var IsInherited sql.NullString
		var AssignedToUsers sql.NullInt64
		var GrantedToRoles sql.NullInt64
		var GrantedRoles sql.NullInt64
		var Owner sql.NullString
		var Comment sql.NullString

		err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &IsInherited, &AssignedToUsers, &GrantedToRoles, &GrantedRoles, &Owner, &Comment)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment})
	}
	defer db.Close()
	return nil, nil
}
