/*
Package snowflake implements a steampipe plugin for Snowflake.

This plugin provides data that Steampipe uses to present foreign
tables that represent Snowflake resources.
*/
package snowflake

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

const pluginName = "steampipe-plugin-snowflake"

// Plugin creates this (snowflake) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             pluginName,
		DefaultTransform: transform.FromGo().Transform(valueFromNullable),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"snowflake_account_grant":  tableSnowflakeAccountGrant(ctx),
			"snowflake_database":       tableSnowflakeDatabase(ctx),
			"snowflake_database_grant": tableSnowflakeDatabaseGrant(ctx),
			"snowflake_network_policy": tableSnowflakeNetworkPolicy(ctx),
			"snowflake_role":           tableSnowflakeRole(ctx),
			"snowflake_role_grant":     tableSnowflakeRoleGrant(ctx),
			"snowflake_session_policy": tableSnowflakeSessionPolicy(ctx),
			"snowflake_user":           tableSnowflakeUser(ctx),
			"snowflake_user_grant":     tableSnowflakeUserGrant(ctx),
			"snowflake_view":           tableSnowflakeView(ctx),
			"snowflake_view_grant":     tableSnowflakeViewGrant(ctx),
			"snowflake_warehouse":      tableSnowflakeWarehouse(ctx)},
	}

	return p
}
