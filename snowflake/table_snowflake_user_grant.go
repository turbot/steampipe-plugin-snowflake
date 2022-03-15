package snowflake

import (
	"context"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

//// TABLE DEFINITION

func tableSnowflakeUserGrant(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "snowflake_user_grant",
		Description: "List all roles granted to a user.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeUserGrants,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "username"},
			},
		},
		Columns: []*plugin.Column{
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromQual("username"), Description: "Name of the snowflake user."},
			{Name: "role", Type: proto.ColumnType_STRING, Description: "Name of the role that has been granted to user.."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the role was granted to the user/role."},
			{Name: "granted_to", Type: proto.ColumnType_STRING, Description: "Type of the object. Only USER for this table."},
			{Name: "granted_by", Type: proto.ColumnType_STRING, Description: "Name of the object that granted access on the user."},
		},
	}
}

type UserGrant RoleGrant

// UserGrantCol returns a reference for a column of a UserGrant
func UserGrantCol(colname string, sp *UserGrant) interface{} {
	switch colname {
	case "created_on":
		return &sp.CreatedOn
	case "role":
		return &sp.Role
	case "granted_to":
		return &sp.GrantedTo
	case "grantee_name":
		return &sp.GranteeName
	case "granted_by":
		return &sp.GrantedBy
	default:
		panic("unknown column " + colname)
	}
}

//// LIST FUNCTION

func listSnowflakeUserGrants(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	user := d.KeyColumnQualString("username")
	if user == "" {
		return nil, nil
	}
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_user_grant.listSnowflakeUserGrants", "connnection.error", err)
		return nil, err
	}
	// SQL compilation error: qual containing special characters in username leads to sql compilation error.
	// Handle sql compilation error for query by wrapping qual inside double quotes.
	rows, err := db.QueryContext(ctx, fmt.Sprintf("SHOW GRANTS TO USER \"%s\"", user))
	if err != nil {
		logger.Error("snowflake_user_grant.listSnowflakeUserGrants", "query.error", err)
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		logger.Error("snowflake_user_grant.listSnowflakeUserGrants", "get_coloumns.error", err)
		return nil, err
	}

	for rows.Next() {
		userGrant := UserGrant{}
		// make references for the cols with the aid of UserGrantCol
		cols := make([]interface{}, len(columns))

		for i, col := range columns {
			cols[i] = UserGrantCol(col, &userGrant)
		}

		err = rows.Scan(cols...)
		if err != nil {
			logger.Error("snowflake_user_grant.listSnowflakeUserGrants", "query_scan.error", err)
			return nil, err
		}

		d.StreamListItem(ctx, userGrant)
	}

	for rows.NextResultSet() {
		for rows.Next() {
			userGrant := UserGrant{}
			// make references for the cols with the aid of UserGrantCol
			cols := make([]interface{}, len(columns))

			for i, col := range columns {
				cols[i] = UserGrantCol(col, &userGrant)
			}

			err = rows.Scan(cols...)
			if err != nil {
				logger.Error("snowflake_user_grant.listSnowflakeUserGrants", "query_scan.error", err)
				return nil, err
			}

			d.StreamListItem(ctx, userGrant)
		}
	}
	return nil, nil
}
