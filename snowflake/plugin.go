/*
Package aws implements a steampipe plugin for aws.

This plugin provides data that Steampipe uses to present foreign
tables that represent Amazon AWS resources.
*/
package snowflake

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

const pluginName = "steampipe-plugin-aws"

// Plugin creates this (aws) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo().Transform(valueFromNullable),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"snowflake_account_grant":  tableAccountGrant(ctx),
			"snowflake_database":       tableDatabase(ctx),
			"snowflake_database_grant": tableDatabaseGrant(ctx),
			"snowflake_network_policy": tableNetworkPolicy(ctx),
			"snowflake_role":           tableRole(ctx),
			"snowflake_role_grant":     tableRoleGrant(ctx),
			"snowflake_session_policy": tableSessionPolicy(ctx),
			"snowflake_user":           tableUser(ctx),
			"snowflake_warehouse":      tableSnowflakeWarehouse(ctx),
			// "snowflake_row_access_policy": tableRowAccessPolicy(ctx),
		},
	}

	return p
}
