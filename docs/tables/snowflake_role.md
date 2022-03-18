# Tables: snowflake_role

The role is an entity to which privileges can be granted. Roles are in turn assigned to users.

## Examples

### Basic info

```sql
select
  name,
  created_on,
  granted_roles,
  granted_to_roles
from
  snowflake_role;
```

### List idle roles

```sql
select
  name,
  created_on,
  assigned_to_users
from
  snowflake_role
where
  assigned_to_users = 0;
```

### List roles with assigned users

```sql
select
  name as role_name,
  grantee_name
from
  snowflake_role
  inner join
    snowflake_role_grant
    on snowflake_role.name = snowflake_role_grant.role
where
  assigned_to_users > 0
  and granted_to = 'USER';
```
