# Table: snowflake_database

All data in Snowflake is maintained in databases. Each database consists of one or more schemas, which are logical groupings of database objects, such as tables and views.

## Examples

### Basic info

```sql
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database;
```

### List databases with retention time greater than 1 day

```sql
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database
where
  retention_time > 1;
```
