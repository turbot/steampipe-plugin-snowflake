connection "snowflake" {
  plugin = "snowflake"

  # Your Snowflake Account
  # account = "xy12345"

  # The user of your Snowflake Account
  # user = "steampipe"

  # The region id for Snowflake Account
  # region = "ap-south-1.aws"

  # role = "ACCOUNTADMIN"

  # Authentication to snowflake can be done in three ways
  # 1. Using Password
  # 2. Key Based Authentication
  # 3. OAuth Based Authentication

  # 1.1 Using Password
  # The password for your Snowflake Account
  # password = "~dummy@pass"

  # 1.2 Key Based Authentication
  # https://docs.snowflake.com/en/user-guide/key-pair-auth.html
  # private_key_path       = "/Users/lalitbhardwaj/Turbot/prod/my_sample_codes/Go_Basics/snowflake/rsa_key.p8"
  # private_key_passphrase = "abcde"
  # OR use private key directly
  # private_key = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde"

  # 1.3 OAuth Based Authentication
  # TODO
}
