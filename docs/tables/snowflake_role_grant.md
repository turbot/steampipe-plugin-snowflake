# Table: snowflake_role_grant

List all privileges and roles granted to a role.

**Note**: This table requires an '=' qualifier for the `role` column.

## Examples

### List users granted the ACCOUNTADMIN role

```sql
select
  role,
  granted_to,
  grantee_name,
  granted_by,
  created_on
from
  snowflake_role_grant
where
  role = 'ACCOUNTADMIN' and
  granted_to = 'USER';
```

### List roles granted the SYSADMIN role

```sql
select
  role,
  granted_to,
  grantee_name,
  granted_by,
  created_on
from
  snowflake_role_grant
where
  role = 'SYSADMIN' and
  granted_to = 'ROLE';
```
