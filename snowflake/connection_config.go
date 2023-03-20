package snowflake

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type snowflakeConfig struct {
	Account              *string `cty:"account"`
	User                 *string `cty:"user"`
	Region               *string `cty:"region"`
	Role                 *string `cty:"role"`
	Password             *string `cty:"password"`
	BrowserAuth          *bool   `cty:"browser_auth"`
	PrivateKeyPath       *string `cty:"private_key_path"`
	PrivateKey           *string `cty:"private_key"`
	PrivateKeyPassphrase *string `cty:"private_key_passphrase"`
	OAuthAccessToken     *string `cty:"oauth_access_token"`
	OAuthClientID        *string `cty:"oauth_client_id"`
	OAuthClientSecret    *string `cty:"oauth_client_secret"`
	OAuthEndpoint        *string `cty:"oauth_endpoint"`
	OAuthRedirectURL     *string `cty:"oauth_redirect_url"`
	OAuthRefreshToken    *string `cty:"oauth_refresh_token"`
	Warehouse            *string `cty:"warehouse"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"account": {
		Type: schema.TypeString,
	},
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"warehouse": {
		Type: schema.TypeString,
	},
	"browser_auth": {
		Type: schema.TypeBool,
	},
	"private_key_path": {
		Type: schema.TypeString,
	},
	"private_key": {
		Type: schema.TypeString,
	},
	"private_key_passphrase": {
		Type: schema.TypeString,
	},
	"oauth_access_token": {
		Type: schema.TypeString,
	},
	"region": {
		Type: schema.TypeString,
	},
	"role": {
		Type: schema.TypeString,
	},
	"oauth_refresh_token": {
		Type: schema.TypeString,
	},
	"oauth_client_id": {
		Type: schema.TypeString,
	},
	"oauth_client_secret": {
		Type: schema.TypeString,
	},
	"oauth_endpoint": {
		Type: schema.TypeString,
	},
	"oauth_redirect_url": {
		Type: schema.TypeString,
	},
	// TODO - Add when generic table support is added
	// "database": {
	// 	Type: schema.TypeString,
	// },
	// "schema": {
	// 	Type: schema.TypeString,
	// },
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
