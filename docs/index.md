---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/snowflake.svg"
brand_color: "#A0E3F6"
display_name: "Snowflake"
name: "snowflake"
description: "Steampipe plugin for querying roles, databases, and more from Snowflake."
og_description: "Query Snowflake with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/snowflake-social-graphic.png"
---

# Snowflake + Steampipe

[Snowflake](https://app.snowflake.com/) enables data storage, processing, and analytic solutions that are faster, easier to use, and far more flexible than traditional offerings.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example: List inactive users

```sql
select
  name,
  email,
  disabled,
  default_role,
  has_password,
  has_rsa_public_key
from
  snowflake_user
where
  (last_success_login > now() - interval '30 days') and
  last_success_login is not null;
```

```
+-----------+------------------+----------+--------------+--------------+--------------------+
| name      | email            | disabled | default_role | has_password | has_rsa_public_key |
+-----------+------------------+----------+--------------+--------------+--------------------+
| ROHIT     | rohit@xyz.com    | false    | ACCOUNTADMIN | true         | false              |
| TEST1     | test1@xyz.com    | false    | PUBLIC       | true         | true               |
+-----------+------------------+----------+--------------+--------------+--------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/snowflake/tables)**

## Get started

### Install

Download and install the latest Snowflake plugin:

```bash
steampipe plugin install snowflake
```

### Credentials

The Snowflake provider supports multiple ways to authenticate:

- Password
- Keypair Authentication
- OAuth Access Token
- OAuth Refresh Token
- Browser Auth

In all authentication methods account, user, and region is required.

### Using username and password

- [Create and manage user with Web Interface](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-the-web-interface)
- [Create and manage user using SQL](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-sql)

```hcl
connection "snowflake" {
  plugin   = "snowflake"
  account  = "xy12345"
  user     = "steampipe"
  region   = "ap-south-1.aws"
  role     = "ACCOUNTADMIN"
  password = "~dummy@pass"
}
```

#### Keypair Authentication

You can [generate the public and private keys](https://docs.snowflake.com/en/user-guide/key-pair-auth.html) here.

```hcl
connection "snowflake" {
  plugin   = "snowflake"
  account  = "xy12345"
  user     = "steampipe"
  region   = "ap-south-1.aws"
  role     = "ACCOUNTADMIN"
  private_key_path = "/Users/lalitbhardwaj/Turbot/prod/my_sample_codes/Go_Basics/snowflake/rsa_key.p8"
  # or private_key = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde" # If key supports passphrase authentication
}
```

#### OAuth Access Token

If you have an OAuth access token:

```hcl
connection "snowflake" {
  plugin   = "snowflake"
  account  = "xy12345"
  user     = "steampipe"
  region   = "ap-south-1.aws"
  role     = "ACCOUNTADMIN"
  oauth_access_token = "...."
}
```

Note that once this access token expires, you'll need to request a new one through an external application.

### OAuth Refresh Token

If you have an OAuth Refresh token:

```hcl
connection "snowflake" {
  plugin              = "snowflake"
  account             = "xy12345"
  user                = "steampipe"
  region              = "ap-south-1.aws"
  role                = "ACCOUNTADMIN"
  oauth_client_id     = "...."
  oauth_client_secret = "...."
  oauth_endpoint      = "...."
  oauth_redirect_url  = "...."
  oauth_refresh_token = "...."
}
```

Note: because access token have a short life; typically 1 hour, refresh token will be used to generate new access token.

### Configuration

Installing the latest snowflake plugin will create a config file (~/.steampipe/config/snowflake.spc) with a single connection named snowflake:

```hcl
connection "snowflake" {
  plugin = "snowflake"

  # Your Snowflake Account
  # account = "xy12345"

  # The user of your Snowflake Account
  # user = "steampipe"

  # The region id for Snowflake Account
  # region = "ap-south-1.aws"

  # Optional; The role that should be used by steampipe for the user.
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

  # 1.2 Key Based Authentication
  # https://docs.snowflake.com/en/user-guide/key-pair-auth.html
  # private_key_path       = "/Users/lalitbhardwaj/Turbot/prod/my_sample_codes/Go_Basics/snowflake/rsa_key.p8"
  # private_key_passphrase = "abcde"

  # OR use private key directly

  # private_key = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
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
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-snowflake
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring Snowflake Credentials

The Snowflake plugin support below authentication mechanisms

1. Using User Password
2. [Key Based Authentication](https://docs.snowflake.com/en/user-guide/key-pair-auth.html)
3. [OAuth Based Authentication](https://docs.snowflake.com/en/user-guide/oauth-custom.html)
