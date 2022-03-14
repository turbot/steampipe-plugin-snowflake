connection "snowflake" {
  plugin = "snowflake"

  # Your Snowflake Account
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#account-identifier-formats-by-cloud-platform-and-region
  # account = "xy12345"

  # The user of your Snowflake Account
  # user = "steampipe"

  # The region id for Snowflake Account
  # https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#snowflake-region-ids
  # region = "ap-south-1.aws"

  # Optional; Specifies the role to use by default for accessing Snowflake objects in the client session.
  # If not mentioned will use the default role for the User
  # role = "ACCOUNTADMIN"

  # Authentication to snowflake can be done in below ways
  # 1. Using Password
  # 2. Key Pair Authentication
  # 3. OAuth Access Token
  # 4. OAuth Refresh Token

  # 1.1 Authentication using password
  # The password for your Snowflake Account
  # password = "~dummy@pass"

  # 1.2 Key Pair Authentication
  # https://docs.snowflake.com/en/user-guide/key-pair-auth.html
  # private_key_path       = "/path/to/snowflake/rsa_key.p8"
  # private_key_passphrase = "abcde"

  # OR use private key directly

  # private_key            = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde"

  # 1.3 OAuth Access Token
  # oauth_access_token = "...."

  # 1.4 OAuth Refresh Token
  # oauth_client_id     = "0oa44dah4cudhAkPU5d7"
  # oauth_client_secret = "wkQYoty7kCRrBzmkqBbubxK-egaJDJ5gT1BH-4b-"
  # oauth_endpoint      = "https://xyz.abc.com/oauth2/auFGTkTZs5d7/v1/token"
  # oauth_redirect_url  = "https://xy1234.ap-south-1.aws.snowflakecomputing.com/"
  # oauth_refresh_token = "0oa44dah4cudhAkPU5d70oa44dah4cudhAkPU5d7"
}
