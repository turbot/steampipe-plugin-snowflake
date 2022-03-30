# Table: snowflake_session

This Account Usage view provides information on the session, including information on the authentication method to Snowflake and the Snowflake login event. Snowflake returns one row for each session created over the last year.

**Note**:

- This table requires a `Snowflake warehouse` to query. You can set it by `warehouse` config argument in Steampipe connection config.
- Latency for the view may be up to 180 minutes (3 hours).

## Examples

### Basic info

```sql
select
  session_id,
  user_name,
  authentication_method,
  created_on
```
