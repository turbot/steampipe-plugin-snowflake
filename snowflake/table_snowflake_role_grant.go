package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableRoleGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_role_grant",
		Description: "Snowflake Role Grant",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeRoleGrants,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "role"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "role", Description: "", Type: proto.ColumnType_STRING},
			{Name: "created_on", Description: "", Type: proto.ColumnType_TIMESTAMP},
			{Name: "granted_to", Description: "", Type: proto.ColumnType_STRING},
			{Name: "grantee_name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "granted_by", Description: "", Type: proto.ColumnType_STRING},
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
		return nil, err
	}

	for rows.Next() {
		var CreatedOn sql.NullTime
		var Role sql.NullString
		var GrantedTo sql.NullString
		var GranteeName sql.NullString
		var GrantedBy sql.NullString

		err = rows.Scan(&CreatedOn, &Role, &GrantedTo, &GranteeName, &GrantedBy)
		if err != nil {
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
			return nil, err
		}

		d.StreamListItem(ctx, RoleGrant{CreatedOn, Role, GrantedTo, GranteeName, GrantedBy})
	}

	defer db.Close()
	return nil, nil
}
