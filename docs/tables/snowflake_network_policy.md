# Table: snowflake_network_policy

## Examples

### Basic info

```sql
select
  name,
  comment,
  created_on,
  entries_in_allowed_ip_list,
  entries_in_blocked_ip_list
from
  snowflake_network_policy;
```

### Get blocked and allowed IP lists for a specific network policy

```sql
select
  name,
  allowed_ip_list,
  blocked_ip_list
from
  snowflake_network_policy
where
  name = 'np1' ;
```
