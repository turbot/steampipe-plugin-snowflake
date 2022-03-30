## v0.0.2 [WIP]

_What's new?_

- New tables added
  - [snowflake_account_parameter](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_account_parameter)
  - [snowflake_login_history](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_login_history)
  - [snowflake_schemata](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_schemata)
  - [snowflake_session](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_session)

_Enhancements_

- Added `account` metadata column to all existing tables.
- Changed the column type of `ext_authn_duo` column from string to bool in `snowflake_user` table.
- Added `warehouse` configuration arguments to allow users to query SNOWFLAKE.ACCOUNT_USAGE views.

## v0.0.1 [2022-03-22]

_What's new?_

- New tables added
  - [snowflake_account_grant](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_account_grant)
  - [snowflake_database](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_database)
  - [snowflake_database_grant](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_database_grant)
  - [snowflake_network_policy](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_network_policy)
  - [snowflake_role](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_role)
  - [snowflake_role_grant](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_role_grant)
  - [snowflake_session_policy](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_session_policy)
  - [snowflake_user](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_user)
  - [snowflake_user_grant](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_user_grant)
  - [snowflake_view](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_view)
  - [snowflake_view_grant](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_view_grant)
  - [snowflake_warehouse](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_warehouse)
