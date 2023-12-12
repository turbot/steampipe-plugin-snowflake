---
title: "Steampipe Table: snowflake_warehouse - Query Snowflake Warehouses using SQL"
description: "Allows users to query Snowflake Warehouses, providing insights into warehouse configurations and usage statistics."
---

# Table: snowflake_warehouse - Query Snowflake Warehouses using SQL

A Snowflake Warehouse is a virtual warehouse in Snowflake, a cloud-based data warehousing platform. A warehouse is the computational resource that executes all data processing tasks, including loading data, executing transformations, and running queries. Each virtual warehouse is an independent compute resource that does not share compute resources with other virtual warehouses.

## Table Usage Guide

The `snowflake_warehouse` table provides insights into Snowflake Warehouses. As a data engineer or analyst, you can explore details about each warehouse, including its size, state, and usage statistics. Use this table to understand the performance of your warehouses and to identify potential areas for optimization.

## Examples

### Basic info
Analyze the settings to understand the status and characteristics of your Snowflake warehouses. This can be useful for capacity planning and resource allocation.

```sql+postgres
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse;
```

```sql+sqlite
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse;
```

### List active warehouses
Explore which warehouses are currently active in your Snowflake environment. This can help in managing resources efficiently and ensuring optimal performance.

```sql+postgres
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse
where
  state = 'STARTED';
```

```sql+sqlite
select
  name,
  size,
  type,
  state
from
  snowflake_warehouse
where
  state = 'STARTED';
```

### Get a count of warehouses grouped by size
Determine the distribution of warehouse sizes within your infrastructure to better manage resources and planning. This query is useful for understanding the scale of your operations.

```sql+postgres
select
  count(*),
  size
from
  snowflake_warehouse
group by
  size;
```

```sql+sqlite
select
  count(*),
  size
from
  snowflake_warehouse
group by
  size;
```

### List warehouses with auto-resume disabled
Determine the areas in which warehouses have the auto-resume feature disabled to assess potential inefficiencies in warehouse management.

```sql+postgres
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

```sql+sqlite
select
  name,
  type,
  size,
  auto_resume
from
  snowflake_warehouse
where
  auto_resume = 0;
```