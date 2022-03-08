package snowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

//// TABLE DEFINITION

func tableNetworkPolicy(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "snowflake_network_policy",
		// Lists all network policies defined in the system. Only returns results for the SECURITYADMIN or ACCOUNTADMIN role.
		// https://docs.snowflake.com/en/user-guide/ui-account.html#network-policies
		Description: "Network policies enable restricting access to your account based on user IP address.",
		List: &plugin.ListConfig{
			Hydrate: listSnowflakeNetworkPolicies,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Identifier for the network policy."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Description: "Date and time when the policy was created."},
			{Name: "comment", Type: proto.ColumnType_STRING, Description: "Specifies a comment for the network policy."},
			{Name: "entries_in_allowed_ip_list", Type: proto.ColumnType_INT, Description: "No of entries in the allowed IP list."},
			{Name: "entries_in_blocked_ip_list", Type: proto.ColumnType_INT, Description: "No of entries in the blocked IP list."},
			{Name: "allowed_ip_list", Type: proto.ColumnType_STRING, Hydrate: DescribeNetworkPolicy, Transform: transform.FromField("ALLOWED_IP_LIST"), Description: "Comma-separated list of one or more IPv4 addresses that are allowed access to your Snowflake account."},
			{Name: "blocked_ip_list", Type: proto.ColumnType_STRING, Hydrate: DescribeNetworkPolicy, Transform: transform.FromField("BLOCKED_IP_LIST"), Description: "Comma-separated list of one or more IPv4 addresses that are denied access to your Snowflake account."},
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
	db, err := connect(ctx, d)
	if err != nil {
		logger.Error("snowflake_network_policy.listSnowflakeNetworkPolicies", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, "SHOW NETWORK POLICIES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	for rows.NextResultSet() {
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
	}

	return nil, nil
}

func DescribeNetworkPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	var policyName string
	if h.Item != nil {
		policyName = h.Item.(NetworkPolicy).Name.String
	} else {
		policyName = d.KeyColumnQualString("name")
	}

	if policyName == "" {
		return nil, nil
	}

	db, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("snowflake_network_policy.DescribeNetworkPolicy", "connnection.error", err)
		return nil, err
	}
	rows, err := db.QueryContext(ctx, fmt.Sprintf("DESCRIBE NETWORK POLICY %s", policyName))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	networkIPlist := map[string]string{}
	for rows.Next() {
		var name sql.NullString
		var value sql.NullString

		err = rows.Scan(&name, &value)
		if err != nil {
			return nil, err
		}
		networkIPlist[name.String] = value.String
	}
	return networkIPlist, nil
}
