# Table: snowflake_session

This Account Usage view provides information on the session, including information on the authentication method to Snowflake and the Snowflake login event. Snowflake returns one row for each session created over the last year.

**Notes**:

- This table requires a [Snowflake warehouse](https://docs.snowflake.com/en/user-guide/warehouses.html) to query. You can specify it in the `warehouse` config argument, or if not specified, the user's default warehouse will be used.
- Latency for the view may be up to 180 minutes (3 hours).

## Examples

### Basic info

```sql
select
  session_id,
  user_name,
  authentication_method,
  created_on,
  client_environment ->> 'APPLICATION' as client_application
from
  snowflake_session;
```

### List distinct authentication methods used in the last year

```sql
select distinct
  user_name,
  authentication_method
from
  snowflake_session
order by
  user_name;
```

### List sessions authenticated without Snowflake MFA with passsword in last 30 days

```sql
select distinct
  user_name,
  authentication_method,
  client_environment ->> 'APPLICATION' as client_application
from
  snowflake_session
where
  split_part(authentication_method, '+', 2) = ''
  and authentication_method like 'Password%'
  and created_on > now() - interval '30 days'
order by
  user_name desc;
```
