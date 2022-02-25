package snowflake

import (
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/schema"
)

type snowflakeConfig struct {
	User     *string `cty:"user"`
	Password *string `cty:"password"`
	Account  *string `cty:"account"`
	Database *string `cty:"database"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"account": {
		Type: schema.TypeString,
	},
	"database": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &snowflakeConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) snowflakeConfig {
	if connection == nil || connection.Config == nil {
		return snowflakeConfig{}
	}
	config, _ := connection.Config.(snowflakeConfig)
	return config
}
