# Table: snowflake_warehouse

A warehouse is a cluster of compute resources in Snowflake. Warehouses provide the required resources, such as CPU, memory, and temporary storage, to perform queries.

## Examples

### Basic info

```sql
select
  name,
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
    round(used_credits/credit_quota*100, 1) as "% Used",
    case
        when used_credits/credit_quota*100 > 90 then 'alert'
        when used_credits/credit_quota*100 > 75 then 'warning'
        else 'ok'
    end as type
from snowflake_resource_monitor
where used_credits/credit_quota*100 > 75
order by used_credits/credit_quota desc
```

### List warehouses which have used all their credits

```sql
select account as "Account",
    name as "Warehouse"
from snowflake_resource_monitor
where remaining_credits < 1
```
