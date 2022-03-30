connection "snowflake" {
  plugin = "snowflake"

  # Snowflake account ID
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#account-identifier-formats-by-cloud-platform-and-region
  # account = "xy12345"

  # Snowflake username
  # user = "steampipe"

  # Snowflake account region ID, defaults to "us-west-2.aws"
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#snowflake-region-ids
  # region = "us-west-2.aws"

  # Specifies the role to use for accessing Snowflake objects in the client session
  # If not specified, the default role for the user will be used
  # role = "ACCOUNTADMIN"

  # Specifies the sSnowflake warehouse to use for executing snowflake queries (Required for schemata, session and login_history tables)
  # If not specified, the default warehouse for the user will be used
  # warehouse = "COMPUTE_WH"

  # You can connect to Snowflake using one of the following methods:

  # 1. Password
  # The password for your Snowflake Account
  # password = "~dummy@pass"

  # 2. Key pair authentication
  # https://docs.snowflake.com/en/user-guide/key-pair-auth.html
  # private_key_path       = "/path/to/snowflake/rsa_key.p8"
  # private_key_passphrase = "abcde"

  # OR use the private key directly:

  # private_key            = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde"

  # 3. OAuth access token
  # https://docs.snowflake.com/en/user-guide/oauth-custom.html
  # oauth_access_token = "eyJraWQiOiJLWjN....jwqt1uCG8Z94ZYZp_LK3YhQbWLkWA"

  # 4. OAuth refresh token
  # https://developer.okta.com/docs/guides/refresh-tokens/main/
  # oauth_client_id     = "0oa44dah4cudhAkPU5b1"
  # oauth_client_secret = "wkQYoty7kCRrBzmkqBbubxK-egaJDJ5gT1BH-4b-"
  # oauth_endpoint      = "https://xyz.abc.com/oauth2/auFGTkTZs5d7/v1/token"
  # oauth_redirect_url  = "https://xy1234.ap-south-1.aws.snowflakecomputing.com/"
  # oauth_refresh_token = "0oa44dah4cudhAkPU5d70oa44dah4cudhAkPU5e2"
}

