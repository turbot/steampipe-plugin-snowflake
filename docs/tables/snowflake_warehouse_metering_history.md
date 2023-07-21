# Table: snowflake_warehouse_metering_history

The `warehouse_metering_history` table in the `ACCOUNT_USAGE` schema in Snowflake is a system-generated table that stores historical information about the usage and consumption of virtual warehouses in your Snowflake account. It provides detailed metrics related to the performance and resource utilization of the virtual warehouses over time.

## Examples

### Basic info

```sql
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

```sql
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

```sql
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

```sql
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

### List the top 5 warehouses with the highest credits used for cloud services in a particular account

```sql
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

```sql
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

```sql
select
  warehouse_id,
  warehouse_name,
  (credits_used_cloud_services / credits_used) * 100 as cloud_services_percentage
from
  snowflake_warehouse_metering_history
where
  credits_used > 0;
```