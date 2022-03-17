package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeRole(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_role",
		Description: "An entity to which privileges can be granted. Roles are in turn assigned to users.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeRole,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the role."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the role was created."},
			{Name: "assigned_to_users", Type: proto.ColumnType_INT, Description: "Number of users the role is assigned."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Comment for the role."},
			{Name: "granted_roles", Type: proto.ColumnType_INT, Description: "Number of roles inherited by this role."},
			{Name: "granted_to_roles", Type: proto.ColumnType_INT, Description: "Number of roles that inherit the privileges of this role."},
			{Name: "is_current", Type: proto.ColumnType_STRING, Description: "\"Y\" if is the current role of authenticated user, otherwise \"F\"."},
			{Name: "is_default", Type: proto.ColumnType_STRING, Description: "\"Y\" if is the default role of authenticated user, otherwise \"F\"."},
			{Name: "is_inherited", Type: proto.ColumnType_STRING, Description: "\"Y\" if current role is inherited by authenticated user, otherwise \"F\"."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the role."},
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
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_role.listSnowflakeRole", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW ROLES")
	if err != nil {
		logger.Error("snowflake_role.listSnowflakeRole", "query.error", err)
		return nil, err
	}
	defer rows.Close()

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
			logger.Error("snowflake_role.listSnowflakeRole", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment})
	}

	for rows.NextResultSet() {
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
				logger.Error("snowflake_role.listSnowflakeRole", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, Role{CreatedOn, Name, IsDefault, IsCurrent, IsInherited, AssignedToUsers, GrantedToRoles, GrantedRoles, Owner, Comment})
		}
	}
	return nil, nil
}
