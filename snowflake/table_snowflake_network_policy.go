package snowflake

import (
	"context"
	"database/sql"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
)

//// TABLE DEFINITION

func tableNetworkPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "snowflake_network_policy",
		// Lists all network policies defined in the system. Only returns results for the SECURITYADMIN or ACCOUNTADMIN role.
		Description: "Snowflake Network Policy",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeNetworkPolicies,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "entries_in_allowed_ip_list", Type: proto.ColumnType_INT, Description: ""},
			{Name: "entries_in_blocked_ip_list", Type: proto.ColumnType_INT, Description: ""},
		},
	}
}

type NetworkPolicy struct {
	Name                   sql.NullString `json:"name"`
	CreatedOn              sql.NullTime   `json:"created_on"`
	Comment                sql.NullString `json:"comment"`
	EntriesInAllowedIPList sql.NullInt64  `json:"entries_in_allowed_ip_list"`
	EntriesInBlockedIPList sql.NullInt64  `json:"entries_in_blocked_ip_list"`
}

//// LIST FUNCTION

func listSnowflakeNetworkPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	logger.Error("aws_region.listSnowflakeNetworkPolicies", "api.error", "nil")
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("aws_region.listSnowflakeNetworkPolicies", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW NETWORK POLICIES")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var Name sql.NullString
		var CreatedOn sql.NullTime
		var Comment sql.NullString
		var EntriesInAllowedIPList sql.NullInt64
		var EntriesInBlockedIPList sql.NullInt64

		err = rows.Scan(&CreatedOn, &Name, &Comment, &EntriesInAllowedIPList, &EntriesInBlockedIPList)
		if err != nil {
			return nil, err
		}

		d.StreamListItem(ctx, NetworkPolicy{Name, CreatedOn, Comment, EntriesInAllowedIPList, EntriesInBlockedIPList})
	}
	defer db.Close()
	return nil, nil
}
