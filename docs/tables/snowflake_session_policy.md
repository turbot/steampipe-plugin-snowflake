# Table: snowflake_session_policy

A session policy defines the idle session timeout period in minutes. The idle session timeout refers to a period of inactivity with either the web interface or a programmatic client (e.g. SnowSQL, JDBC driver). When the idle session timeout period expires, users must authenticate to Snowflake again.

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
