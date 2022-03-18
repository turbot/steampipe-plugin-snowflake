# Table: snowflake_account_grant

List all account-level, i.e., global, privileges that have been granted to roles.

## Examples

### Basic info

```sql
select
  name,
  privilege,
  grantee_name,
  granted_to,
  grant_option,
  created_on
from
  snowflake_account_grant;
```

### List privileges with the ACCOUNTADMIN role

```sql
select
  privilege,
  grant_option,
  created_on
from
  snowflake_account_grant
where
  grantee_name = 'ACCOUNTADMIN';
```
