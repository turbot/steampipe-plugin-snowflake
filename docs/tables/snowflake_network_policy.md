# Table: snowflake_network_policy

Network policies enable restricting access to your account based on user IP address.

**Note**: Only the network policy owner (i.e. role with the `OWNERSHIP` privilege on the network policy) or higher can query this table.

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
