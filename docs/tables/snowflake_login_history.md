# Table: snowflake_login_history

This Account Usage view can be used to query login attempts by Snowflake users within the last 365 days (1 year).

**Note:** This table requires a `Snowflake warehouse` to query. You can set it by `warehouse` config argument in Steampipe connection config.

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

### List all authentication methods used for login in 30 days

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
