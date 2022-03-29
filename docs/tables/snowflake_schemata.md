# Table: snowflake_schemata

This Information Schema view displays a row for each schema in the specified (or current) database, including the INFORMATION_SCHEMA schema itself.

**Note**: This table requires a `Snowflake warehouse` to query. You can set it by `warehouse` config argument in Steampipe connection config.

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
