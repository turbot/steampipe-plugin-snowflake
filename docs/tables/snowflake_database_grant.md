# Table: snowflake_database_grant

Lists all privileges that have been granted on a database.

**Note** This tables requires an '=' qualifier for `database` columns

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

### List grants for all the database in Snoflake

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