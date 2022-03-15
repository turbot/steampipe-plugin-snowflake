---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/snowflake.svg"
brand_color: "#2596BE"
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
  last_success_login
from
  snowflake_user
where
  (last_success_login > now() - interval '30 days')
  and last_success_login is not null;
```

```
+-----------+------------------+----------+--------------+--------------+--------------------+
| name      | email            | disabled | default_role | has_password | has_rsa_public_key |
+-----------+------------------+----------+--------------+--------------+--------------------+
| ROHIT     | rohit@xyz.com    | false    | ACCOUNTADMIN | true         | false              |
| SUMIT     | sumit@xyz.com    | false    | PUBLIC       | true         | true               |
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

In all authentication methods account, user, and region is required.

#### Using username and password

- [Create and manage user with Web Interface](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-the-web-interface)
- [Create and manage user using SQL](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-sql)

```hcl
connection "snowflake" {
  plugin   = "snowflake"
  account  = "xy12345"
  region   = "ap-south-1.aws"
  role     = "ACCOUNTADMIN"
  user     = "steampipe"
  password = "~dummy@pass"
}
```

#### Keypair Authentication

You can [generate the public and private keys](https://docs.snowflake.com/en/user-guide/key-pair-auth.html) here.

```hcl
connection "snowflake" {
  plugin                   = "snowflake"
  account                  = "xy12345"
  user                     = "steampipe"
  region                   = "ap-south-1.aws"
  role                     = "ACCOUNTADMIN"
  private_key_path         = "/path/to/rsa_key.p8"
  # private_key            = "-----BEGIN ENCRYPTED PRIVATE KEY-----\nMIIFHzBJ....au/BUg==\n-----END ENCRYPTED PRIVATE KEY-----"
  # private_key_passphrase = "abcde" # If key supports passphrase authentication
}

```

#### OAuth Access Token

If you have an OAuth access token:

```hcl
connection "snowflake" {
  plugin             = "snowflake"
  account            = "xy12345"
  user               = "steampipe"
  region             = "ap-south-1.aws"
  role               = "ACCOUNTADMIN"
  oauth_access_token = "eyJraWQiOiJLWjN....jwqt1uCG8Z94ZYZp_LK3YhQbWLkWA"
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
  oauth_client_id     = "0oa44dah4cudhAkPU5d7"
  oauth_client_secret = "wkQYoty7kCRrBzmkqBbubxK-egaJDJ5gT1BH-4b-"
  oauth_endpoint      = "https://xyz.abc.com/oauth2/auFGTkTZs5d7/v1/token"
  oauth_redirect_url  = "https://xy1234.ap-south-1.aws.snowflakecomputing.com/"
  oauth_refresh_token = "0oa44dah4cudhAkPU5d70oa44dah4cudhAkPU5d7"
}
```

Note: because access token have a short life; typically 1 hour, refresh token will be used to generate new access token.

### Configuration

Installing the latest snowflake plugin will create a config file (~/.steampipe/config/snowflake.spc) with a single connection named snowflake:

```hcl
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
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-snowflake
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring Snowflake Credentials

The Snowflake plugin support below authentication mechanisms

1. [Using User Password](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-the-web-interface)
2. [Key Based Authentication](https://docs.snowflake.com/en/user-guide/key-pair-auth.html)
3. [OAuth Access Token Authentication](https://docs.snowflake.com/en/user-guide/oauth-custom.html).

   - [Configure Okta for External OAuth](https://docs.snowflake.com/en/user-guide/oauth-okta.html#label-ext-oauth-integration-okta)

4. [OAuth Refresh Token Authentication](https://docs.snowflake.com/en/user-guide/oauth-custom.html)

   - [Refresh access tokens in Okta](https://developer.okta.com/docs/guides/refresh-tokens/main/)
