# Table: snowflake_role_grant

Lists all privileges and roles granted to a role.

**Note** This table requires an '=' qualifier for `role` columns

## Examples

### List users granted ACCOUNTADMIN role

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

### List roles granted SYSADMIN role

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
