package snowflake

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type snowflakeConfig struct {
	Account              *string `hcl:"account"`
	User                 *string `hcl:"user"`
	Region               *string `hcl:"region"`
	Role                 *string `hcl:"role"`
	Password             *string `hcl:"password"`
	BrowserAuth          *bool   `hcl:"browser_auth"`
	PrivateKeyPath       *string `hcl:"private_key_path"`
	PrivateKey           *string `hcl:"private_key"`
	PrivateKeyPassphrase *string `hcl:"private_key_passphrase"`
	OAuthAccessToken     *string `hcl:"oauth_access_token"`
	OAuthClientID        *string `hcl:"oauth_client_id"`
	OAuthClientSecret    *string `hcl:"oauth_client_secret"`
	OAuthEndpoint        *string `hcl:"oauth_endpoint"`
	OAuthRedirectURL     *string `hcl:"oauth_redirect_url"`
	OAuthRefreshToken    *string `hcl:"oauth_refresh_token"`
	Warehouse            *string `hcl:"warehouse"`
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
