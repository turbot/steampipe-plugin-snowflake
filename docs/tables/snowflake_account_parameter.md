# Table: snowflake_account_parameter

Snowflake provides parameters that let you control the behavior of your account, individual user sessions, and objects. All the parameters have default values, which can be set and then overridden at different levels depending on the parameter type (Account, Session, or Object).

Account parameters can be set only at the account level by users with the appropriate administrator role.

## Examples

### Basic info

```sql
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter;
```

### Check whether account allows MFA caching

```sql
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'ALLOW_CLIENT_MFA_CACHING';
```

### Get number of days Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on the object

```sql
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'DATA_RETENTION_TIME_IN_DAYS';
```
