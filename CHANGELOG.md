## v0.3.0 [2022-09-26]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#12](https://github.com/turbot/steampipe-plugin-snowflake/pull/12))
- Recompiled plugin with Go version `1.19`. ([#12](https://github.com/turbot/steampipe-plugin-snowflake/pull/12))

## v0.2.0 [2022-05-19]

_What's new?_

- New tables added
  - [snowflake_resource_monitor](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_resource_monitor) ([#8](https://github.com/turbot/steampipe-plugin-snowflake/pull/8))

## v0.1.1 [2022-05-05]

_Bug fixes_

- Fixed query in inactive user example in `snowflake_user` table doc.

## v0.1.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#3](https://github.com/turbot/steampipe-plugin-snowflake/pull/3))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30). ([#4](https://github.com/turbot/steampipe-plugin-snowflake/pull/4))

## v0.0.2 [2022-04-04]

_What's new?_

- New tables added
  - [snowflake_account_parameter](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_account_parameter) ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
  - [snowflake_login_history](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_login_history) ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
  - [snowflake_schemata](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_schemata) ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
  - [snowflake_session](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_session) ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))

_Enhancements_

- Recompiled plugin with Go 1.18 ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
- Added `warehouse` configuration argument to allow users to query SNOWFLAKE.ACCOUNT_USAGE views ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
- Added `account` and `region` common columns to all tables ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))
- Added `min_cluster_count`, `max_cluster_count`, `started_clusters` and `scaling_policy` columns to `snowflake_warehouse` table ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))

_Bug fixes_

- Fixed the column type of `ext_authn_duo` column from `string` to `bool` in `snowflake_user` table ([#2](https://github.com/turbot/steampipe-plugin-snowflake/pull/2))

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
