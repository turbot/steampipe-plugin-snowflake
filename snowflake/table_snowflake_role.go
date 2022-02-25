package snowflake

import (
	"context"
	"time"

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
			{Name: "name", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "The name of the region", Type: proto.ColumnType_TIMESTAMP},
			{Name: "is_default", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "is_current", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "is_inherited", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "assigned_to_users", Description: "The name of the region", Type: proto.ColumnType_INT},
			{Name: "owner", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "comment", Description: "The name of the region", Type: proto.ColumnType_STRING},
			{Name: "granted_to_roles", Description: "The name of the region", Type: proto.ColumnType_INT},
			{Name: "granted_roles", Description: "The name of the region", Type: proto.ColumnType_INT},
		},
	}
}

type Role struct {
	CreatedOn       time.Time `json:"created_on"`
	Name            string    `json:"name"`
	IsDefault       string    `json:"is_default"`
	IsCurrent       string    `json:"is_current"`
	IsInherited     string    `json:"is_inherited"`
	AssignedToUsers int64     `json:"assigned_to_users"`
	GrantedToRoles  int64     `json:"granted_to_roles"`
	GrantedRoles    int64     `json:"granted_roles"`
	Owner           string    `json:"owner"`
	Comment         string    `json:"comment"`
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
		var CreatedOn time.Time
		var Name string
		var IsDefault string
		var IsCurrent string
		var IsInherited string
		var AssignedToUsers int64
		var GrantedToRoles int64
		var GrantedRoles int64
		var Owner string
		var Comment string

		err = rows.Scan(&CreatedOn, &Name, &IsDefault, &IsCurrent, &IsInherited, &AssignedToUsers, &GrantedToRoles, &GrantedRoles, &Owner, &Comment)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment})
	}
	defer db.Close()
	return nil, nil
}
