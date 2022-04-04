# Table: snowflake_login_history

This Account Usage view can be used to query login attempts by Snowflake users within the last 365 days (1 year).

**Notes:**

- This table requires a [Snowflake warehouse](https://docs.snowflake.com/en/user-guide/warehouses.html) to query. You can specify it in the `warehouse` config argument, or if not specified, the user's default warehouse will be used.
- Latency for the view may be up to 120 minutes (2 hours).

## Examples

### Basic info

```sql
select
  user_name,
  first_authentication_factor,
  is_success,
  event_timestamp
from
  snowflake_login_history;
```

### List all authentication methods used in the last 30 days

```sql
select distinct
  user_name,
  first_authentication_factor
from
  snowflake_login_history
where
  is_success = 'YES'
  and event_timestamp > now() - interval '30 days'
order by
  user_name;
```
