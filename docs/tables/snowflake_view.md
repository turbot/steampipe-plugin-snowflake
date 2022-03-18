# Table: snowflake_view

A view is a named definition of a query.

Snowflake supports two types of views:

- `Non-materialized views`: A non-materialized view’s results are created by executing the query at the time that the view is referenced in a query.
- `Materialized views`: A materialized view’s results are stored, almost as though the results were a table. This allows faster access but requires storage space and active maintenance, both of which incur additional costs.

## Examples

### Basic info

```sql
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view;
```

### List materialized views

```sql
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_materialized;
```

### List secure views

```sql
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_secure;
```
