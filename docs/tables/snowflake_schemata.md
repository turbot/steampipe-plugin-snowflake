# Table: snowflake_schemata

This Information Schema view displays a row for each schema in the specified (or current) database, including the INFORMATION_SCHEMA schema itself.

**Note**:

- This table requires a `Snowflake warehouse` to query. You can set it by `warehouse` config argument in Steampipe connection config.
- The view only displays objects for which the current role for the session has been granted access privileges.
- Latency for the view may be up to 120 minutes (2 hours).

## Examples

### Basic info

```sql
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata;
```

### List schemas that allow managed access

```sql
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

### List schemas that are transient

```sql
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

### List schemas with a retention time greater than 15 days

```sql
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  retention_time,
  schema_owner
from
  snowflake_schemata
where
  retention_time > 10;
```
