# Table: snowflake_resource_monitor

A resource monitor can be used to monitor credit usage by user-managed virtual warehouses and virtual warehouses used by cloud services.

## Examples

### Basic info

```sql
select
  name as warehouse,
  credit_quota,
  used_credits,
  remaining_credits
from
  snowflake_resource_monitor;
```

### List warehouses and % credit left

```sql
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

### List warehouses which have used all their credits

```sql
select
  account,
  name as warehouse
from
  snowflake_resource_monitor
where
  remaining_credits < 1;
```
