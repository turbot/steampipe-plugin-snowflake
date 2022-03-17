connection "snowflake" {
  plugin = "snowflake"

  # Your Snowflake Account
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#account-identifier-formats-by-cloud-platform-and-region
  # account = "xy12345"

  # The user of your Snowflake Account
  # user = "steampipe"

  # The region id for Snowflake Account
  # "us-west-2" is the default region. If the region id  is not mentioned steampipe will assume it to be "us-west-2.aws"
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#snowflake-region-ids
  # region = "us-west-2.aws"

  # Optional; Specifies the role to use by default for accessing Snowflake objects in the client session.
  # If not mentioned will use the default role for the User
  # role = "ACCOUNTADMIN"

  # Authentication to snowflake can be done in below ways
  # 1. Using Password
  # 2. Key Pair Authentication
  # 3. OAuth Access Token
  # 4. OAuth Refresh Token

  # 1 Authentication using password
  # The password for your Snowflake Account
  # password = "~dummy@pass"

  # 2 Key Pair Authentication
  # https://docs.snowflake.com/en/user-guide/key-pair-auth.html
  # private_key_path       = "/path/to/snowflake/rsa_key.p8"
  # private_key_passphrase = "abcde"

  # OR use private key directly

  # private_key            = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde"

  # 3 OAuth Access Token
  # https://docs.snowflake.com/en/user-guide/oauth-custom.html
  # oauth_access_token = "eyJraWQiOiJLWjN....jwqt1uCG8Z94ZYZp_LK3YhQbWLkWA"

  # 4 OAuth Refresh Token
  # https://developer.okta.com/docs/guides/refresh-tokens/main/
  # oauth_client_id     = "0oa44dah4cudhAkPU5d7"
  # oauth_client_secret = "wkQYoty7kCRrBzmkqBbubxK-egaJDJ5gT1BH-4b-"
  # oauth_endpoint      = "https://xyz.abc.com/oauth2/auFGTkTZs5d7/v1/token"
  # oauth_redirect_url  = "https://xy1234.ap-south-1.aws.snowflakecomputing.com/"
  # oauth_refresh_token = "0oa44dah4cudhAkPU5d70oa44dah4cudhAkPU5d7"
}

