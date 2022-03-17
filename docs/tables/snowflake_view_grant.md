# Table: snowflake_view_grant

List all privileges that have been granted on the view.

**Note** This table requires an '=' qualifier for `view_name`, `database_name` and `schema_name` columns.

## Examples

### Basic info

```sql
select
  view_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
where
  view_name = 'ROLES'
  and database_name = 'SNOWFLAKE'
  and schema_name = 'ACCOUNT_USAGE';
```

### List view grants for `ACCOUNT_USAGE` schema in `SNOWFLAKE` database

```sql
select
  view_name,
  snowflake_view.database_name,
  snowflake_view.schema_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
  inner join
    snowflake_view
    on snowflake_view_grant.view_name = snowflake_view.name
    and snowflake_view_grant.database_name = snowflake_view.database_name
    and snowflake_view_grant.schema_name = snowflake_view.schema_name
where
  snowflake_view_grant.database_name = 'SNOWFLAKE'
  and snowflake_view_grant.schema_name = 'ACCOUNT_USAGE';
```
