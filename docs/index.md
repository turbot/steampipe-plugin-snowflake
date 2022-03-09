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

- TODO

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

  # Authentication to snowflake can be done in three ways
  # 1. Using Password
  # 2. Key Based Authentication
  # 3. OAuth Based Authentication

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

  # 1.3 OAuth Based Authentication
  # TODO
}
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-snowflake
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring Snowflake Credentials

The Snowflake plugin support below authentication mechanisms

1. `Using User Password`
2. [Key Based Authentication](https://docs.snowflake.com/en/user-guide/key-pair-auth.html)
3. [OAuth Based Authentication](https://docs.snowflake.com/en/user-guide/oauth-custom.html)
