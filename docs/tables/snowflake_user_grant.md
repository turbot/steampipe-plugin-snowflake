# Table: snowflake_user_grant

Roles are granted to users, providing them with specific permissions.

**Notes**

- The `PUBLIC` role, which is automatically available to every user, is not listed in this table
- This table requires an '=' qualifier for the `username` column

### List all grants for a specific user

```sql
select
  username,
  role,
  granted_by,
  created_on
from
  snowflake_user_grant
where
  username = 'STEAMPIPE';
```

### List all account-level privileges for a specific user

```sql
select
  privilege,
  role,
  username,
  sug.created_on,
  sug.granted_by
from
  snowflake.snowflake_account_grant sag
  inner join
    snowflake.snowflake_user_grant sug
    on sug.role = sag.grantee_name
where
  sug.username = 'STEAMPIPE'
  and sag.granted_to = 'ROLE'
order by
  sag.grantee_name;
```
