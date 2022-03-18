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

For example, to list inactive users:

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

The Snowflake plugin supports multiple ways to authenticate:

- Password
- Key pair authentication
- OAuth access token
- OAuth refresh token

For all authentication methods, `account`, `user`, and `region` are required.

#### Password

You can manage your password through the [Web Interface](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-the-web-interface) or using [SQL](https://docs.snowflake.com/en/user-guide/admin-user-management.html#using-sql).

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

#### Key Pair Authentication

To generate your key pair, please see [Key Pair Authentication](https://docs.snowflake.com/en/user-guide/key-pair-auth.html).

```hcl
connection "snowflake" {
  plugin                   = "snowflake"
  account                  = "xy12345"
  user                     = "steampipe"
  region                   = "ap-south-1.aws"
  role                     = "ACCOUNTADMIN"
  private_key_path         = "/path/to/rsa_key.p8"
}

```

#### OAuth Access Token

To create your OAuth access token, please see [Configure Snowflake OAuth for Custom Clients](https://docs.snowflake.com/en/user-guide/oauth-custom.html).

If using Okta, please see [Configure Okta for External OAuth](https://docs.snowflake.com/en/user-guide/oauth-okta.html#label-ext-oauth-integration-okta).

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

Note: Once the access token in `oauth_access_token` expires, you'll need to request a new one through an external application and update your connnection config.

#### OAuth Refresh Token

Because OAuth access tokens typically have a short life, e.g., 10 minutes, refresh tokens may be a better authentication method as they will automatically obtain new access tokens once expired.

To request your OAuth refresh token, please see [Configure Snowflake OAuth for Custom Clients](https://docs.snowflake.com/en/user-guide/oauth-custom.html)

If using Okta, please see [Configure Okta for External OAuth](https://docs.snowflake.com/en/user-guide/oauth-okta.html#label-ext-oauth-integration-okta).

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

### Configuration

Installing the latest snowflake plugin will create a config file (~/.steampipe/config/snowflake.spc) with a single connection named snowflake:

```hcl
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
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-snowflake
- Community: [Slack Channel](https://steampipe.io/community/join)
