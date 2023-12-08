---
title: "Steampipe Table: snowflake_warehouse_metering_history - Query Snowflake Warehouse Metering History using SQL"
description: "Allows users to query Snowflake Warehouse Metering History, specifically the consumption of Snowflake credits by virtual warehouses over time, providing insights into resource utilization and potential cost optimization."
---

# Table: snowflake_warehouse_metering_history - Query Snowflake Warehouse Metering History using SQL

Snowflake Warehouse Metering History is a feature within Snowflake that tracks the consumption of Snowflake credits by virtual warehouses over time. It provides a detailed breakdown of resource utilization, enabling users to monitor and manage their Snowflake credit usage. This can aid in identifying patterns, optimizing costs, and managing resources more efficiently.

## Table Usage Guide

The `snowflake_warehouse_metering_history` table provides insights into the consumption of Snowflake credits by virtual warehouses over time. As a data analyst or a cloud cost manager, explore detailed breakdowns of resource utilization through this table, including the number of Snowflake credits consumed, the time period of consumption, and the specific warehouses involved. Utilize it to uncover information about credit consumption patterns, cost optimization opportunities, and efficient resource management.

## Examples

### Basic info
Discover the segments that have utilized resources in your Snowflake warehouse. This query is beneficial as it allows you to analyze the consumption of credits, providing insights into resource usage and aiding in efficient resource management.

```sql+postgres
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history;
```

```sql+sqlite
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history;
```

### List the metering history for a particular warehouse
Gain insights into the usage history of a specific warehouse by analyzing its consumption of credits over time. This aids in cost management and optimization by tracking resource usage.

```sql+postgres
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  warehouse_name = 'COMPUTE_WH';
```

```sql+sqlite
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  warehouse_name = 'COMPUTE_WH';
```

### List the metering history for the inactive warehouses
Explore the metering history of warehouses that are currently inactive. This can be useful to analyze past resource usage and expenditure for warehouses that are no longer in use.

```sql+postgres
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history as h,
  snowflake_warehouse as w
where
  h.warehouse_name = w.name
  and state = 'SUSPENDED';
```

```sql+sqlite
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history as h,
  snowflake_warehouse as w
where
  h.warehouse_name = w.name
  and state = 'SUSPENDED';
```

### List the metering history for the last 10 days
Explore the credit usage of your warehouse in the last 10 days. This helps in understanding the resource consumption for better planning and management.

```sql+postgres
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  start_time >= now() - interval '10' day;
```

```sql+sqlite
select
  warehouse_name,
  warehouse_id,
  start_time,
  end_time,
  credits_used,
  credits_used_compute,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  start_time >= datetime('now', '-10 days');
```

### List the top 5 warehouses with the highest credits used for cloud services in a particular account
Explore the top five warehouses with the highest usage of credits for cloud services within a specific account. This can be beneficial in identifying potential areas of cost savings and optimizing resource allocation.

```sql+postgres
select
  warehouse_id,
  warehouse_name,
  account,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  account = 'desired_account'
order by
  credits_used_cloud_services desc
limit 5;
```

```sql+sqlite
select
  warehouse_id,
  warehouse_name,
  account,
  credits_used_cloud_services
from
  snowflake_warehouse_metering_history
where
  account = 'desired_account'
order by
  credits_used_cloud_services desc
limit 5;
```

### Calculate the average credits used per hour for each warehouse
Analyze the usage of each warehouse by calculating the average credits consumed per hour. This can help in cost optimization and efficient resource allocation.

```sql+postgres
select
  warehouse_id,
  warehouse_name,
  AVG(credits_used) as avg_credits_per_hour
from
  snowflake_warehouse_metering_history
group by
  warehouse_id,
  warehouse_name;
```

```sql+sqlite
select
  warehouse_id,
  warehouse_name,
  AVG(credits_used) as avg_credits_per_hour
from
  snowflake_warehouse_metering_history
group by
  warehouse_id,
  warehouse_name;
```

### Calculate the percentage of cloud services credits used compared to total credits for each warehouse
Determine the proportion of cloud services credits utilized in relation to the total credits for each warehouse. This is useful for understanding the extent of cloud services usage and managing resource allocation effectively.

```sql+postgres
select
  warehouse_id,
  warehouse_name,
  (credits_used_cloud_services / credits_used) * 100 as cloud_services_percentage
from
  snowflake_warehouse_metering_history
where
  credits_used > 0;
```

```sql+sqlite
select
  warehouse_id,
  warehouse_name,
  (credits_used_cloud_services / credits_used) * 100 as cloud_services_percentage
from
  snowflake_warehouse_metering_history
where
  credits_used > 0;
```