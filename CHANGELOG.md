## v0.8.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/install/steampipe.sh), as a [Postgres FDW](https://steampipe.io/install/postgres.sh), as a [SQLite extension](https://steampipe.io/install/sqlite.sh) and as a standalone [exporter](https://steampipe.io/install/export.sh).
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension.
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-snowflake/blob/main/docs/LICENSE).

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server enacapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#47](https://github.com/turbot/steampipe-plugin-snowflake/pull/47))

## v0.7.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#35](https://github.com/turbot/steampipe-plugin-snowflake/pull/35))

## v0.7.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#30](https://github.com/turbot/steampipe-plugin-snowflake/pull/30))
- Recompiled plugin with Go version `1.21`. ([#30](https://github.com/turbot/steampipe-plugin-snowflake/pull/30))

## v0.6.0 [2023-07-24]

_What's new?_

- New tables added
  - [snowflake_warehouse_metering_history](https://hub.steampipe.io/plugins/turbot/snowflake/tables/snowflake_warehouse_metering_history) ([#21](https://github.com/turbot/steampipe-plugin-snowflake/pull/21))

## v0.5.0 [2023-05-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05) which fixes increased plugin initialization time due to multiple connections causing the schema to be loaded repeatedly. ([#19](https://github.com/turbot/steampipe-plugin-snowflake/pull/19))

## v0.4.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#17](https://github.com/turbot/steampipe-plugin-snowflake/pull/17))

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
