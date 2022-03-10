# Table: snowflake_warehouse

A warehouse is a cluster of compute resources in Snowflake. Warehouse provides the required resources, such as CPU, memory, and temporary storage, to perform queries.

## Examples

### Basic info

```sql
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse;
```

### List active warehouses

```sql
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse
where
  state = 'STARTED'
```

### Get count of warehouses grouped by size

```sql
select
  count(*),
  size
from
  snowflake_warehouse
group by
  size;
```

### List warehouses with auto_resume turned off

```sql
select
  name,
  type,
  size,
  auto_resume
from
  snowflake_warehouse
where
  not auto_resume;
```
