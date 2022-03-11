# Table: snowflake_database_grant

Lists privileges that have been granted on the database.

**Note** This table requires an '=' qualifier for `database` columns.

## Examples

### Basic info

```sql
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant where database = 'SNOWFLAKE'
```

### List grants for all databases in Snowflake

```sql
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant
  inner join
    snowflake_database
    on snowflake_database_grant.database = snowflake_database.name
```
