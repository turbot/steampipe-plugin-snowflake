---
title: "Steampipe Table: snowflake_resource_monitor - Query Snowflake Resource Monitors using SQL"
description: "Allows users to query Resource Monitors in Snowflake, specifically providing insights into the usage and limits of resources."
---

# Table: snowflake_resource_monitor - Query Snowflake Resource Monitors using SQL

A Snowflake Resource Monitor is a tool within Snowflake that allows you to track and control the usage of resources within your Snowflake account. It provides a way to set up and manage alerts for various resources, including virtual warehouses, databases, and more. Snowflake Resource Monitor helps you stay informed about the health and performance of your Snowflake resources and take appropriate actions when predefined conditions are met.

## Table Usage Guide

The `snowflake_resource_monitor` table provides insights into Resource Monitors within Snowflake. As a Database Administrator, explore monitor-specific details through this table, including usage, limits, and associated metadata. Utilize it to uncover information about resources, such as those nearing their limits, the usage patterns, and the verification of resource usage.

## Examples

### Basic info
Analyze your Snowflake resource monitor to understand your warehouse's credit usage. This can help you manage your resources more efficiently by showing you how much credit quota is available, how much has been used, and what remains.

```sql+postgres
select
  name as warehouse,
  credit_quota,
  used_credits,
  remaining_credits
from
  snowflake_resource_monitor;
```

```sql+sqlite
select
  name as warehouse,
  credit_quota,
  used_credits,
  remaining_credits
from
  snowflake_resource_monitor;
```

### List warehouses and % credit left
Determine the areas in which your warehouse's credit usage exceeds 75% of the total quota. This query helps in monitoring resource consumption, alerting you to potential overruns before they occur.

```sql+postgres
select account,
    name as warehouse,
    credit_quota,
    used_credits,
    remaining_credits,
    round((used_credits/credit_quota*100)::numeric, 1) as percent_used,
    case
      when used_credits/credit_quota*100 > 90 then 'alert'
      when used_credits/credit_quota*100 > 75 then 'warning'
      else 'ok'
    end as type
from
  snowflake_resource_monitor
where
  used_credits/credit_quota*100 > 75
order by
  used_credits/credit_quota desc;
```

```sql+sqlite
select account,
    name as warehouse,
    credit_quota,
    used_credits,
    remaining_credits,
    round((used_credits/credit_quota*100), 1) as percent_used,
    case
      when used_credits/credit_quota*100 > 90 then 'alert'
      when used_credits/credit_quota*100 > 75 then 'warning'
      else 'ok'
    end as type
from
  snowflake_resource_monitor
where
  used_credits/credit_quota*100 > 75
order by
  used_credits/credit_quota desc;
```

### List warehouses which have used all their credits
The query is used to identify Snowflake warehouses that have exhausted their credit allocation. This is beneficial in managing resources effectively and avoiding potential disruptions in data processing tasks.

```sql+postgres
select
  account,
  name as warehouse
from
  snowflake_resource_monitor
where
  remaining_credits < 1;
```

```sql+sqlite
select
  account,
  name as warehouse
from
  snowflake_resource_monitor
where
  remaining_credits < 1;
```