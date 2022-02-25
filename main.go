package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-snowflake/snowflake"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: snowflake.Plugin})
}
