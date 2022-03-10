# Table: snowflake_session_policy

A session policy defines the idle session timeout period in minutes and provides the option to override the default idle timeout value of 4 hours.

The session policy can be set for an account or user with configurable idle timeout periods to address compliance requirements. If a user is associated with both an account and user-level session policy, the user-level session policy takes precedence.

## Examples

### Basic info

```sql
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy;
```

### List policies with idle timeout more that an hour

```sql
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy
where
  session_idle_timeout_mins > 60 or
  session_ui_idle_timeout_mins > 60;
```
