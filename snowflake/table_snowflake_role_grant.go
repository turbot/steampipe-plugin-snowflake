package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableSnowflakeRoleGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_role_grant",
		Description: "Lists all privileges and roles granted to the role.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeRoleGrants,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "role"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Name of the role on that access has been granted."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the role was granted to the user/role."},
			{Name: "granted_to", Type: proto.ColumnType_STRING, Description: "Type of the object. Valid values USER and ROLE."},
			{Name: "grantee_name", Type: proto.ColumnType_STRING, Description: "Name of the object role has been granted."},
			{Name: "granted_by", Type: proto.ColumnType_STRING, Description: "Name of the object that granted access on the role."},
		},
	}
}

type RoleGrant struct {
	CreatedOn   sql.NullTime   `json:"created_on"`
	Role        sql.NullString `json:"role"`
	GrantedTo   sql.NullString `json:"granted_to"`
	GranteeName sql.NullString `json:"grantee_name"`
	GrantedBy   sql.NullString `json:"granted_by"`
}

//// LIST FUNCTION

func listSnowflakeRoleGrants(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	role := d.KeyColumnQualString("role")
	if role == "" {
		return nil, nil
	}
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_role_grant.listSnowflakeRoleGrants", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS OF ROLE %s", role))
	if err != nil {
		logger.Error("snowflake_role_grant.listSnowflakeRoleGrants", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var CreatedOn sql.NullTime
		var Role sql.NullString
		var GrantedTo sql.NullString
		var GranteeName sql.NullString
		var GrantedBy sql.NullString

		err = rows.Scan(&CreatedOn, &Role, &GrantedTo, &GranteeName, &GrantedBy)
		if err != nil {
			logger.Error("snowflake_role_grant.listSnowflakeRoleGrants", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, RoleGrant{CreatedOn, Role, GrantedTo, GranteeName, GrantedBy})
	}

	for rows.NextResultSet() {
		var CreatedOn sql.NullTime
		var Role sql.NullString
		var GrantedTo sql.NullString
		var GranteeName sql.NullString
		var GrantedBy sql.NullString

		err = rows.Scan(&CreatedOn, &Role, &GrantedTo, &GranteeName, &GrantedBy)
		if err != nil {
			logger.Error("snowflake_role_grant.listSnowflakeRoleGrants", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, RoleGrant{CreatedOn, Role, GrantedTo, GranteeName, GrantedBy})
	}

	return nil, nil
}
