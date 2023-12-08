---
title: "Steampipe Table: snowflake_role - Query Snowflake Roles using SQL"
description: "Allows users to query Snowflake Roles, providing insights into the access and permissions assigned to each role in the Snowflake data warehousing service."
---

# Table: snowflake_role - Query Snowflake Roles using SQL

Snowflake Roles are a key component of the Snowflake data warehousing service's access control architecture. They dictate the level of access that a user has to Snowflake objects, such as databases, schemas, and warehouses. Roles can be assigned to users, other roles, and integration objects to facilitate granular, role-based access control.

## Table Usage Guide

The `snowflake_role` table provides insights into the roles within Snowflake's access control architecture. If you are a security analyst or administrator, you can use this table to explore role-specific details, including the permissions assigned to each role and the users and roles to which each role is assigned. This table can be particularly useful for auditing access controls, identifying overly permissive roles, and ensuring compliance with your organization's access policies.

## Examples

### Basic info
Explore the roles within your Snowflake environment, including when they were created and their associated permissions. This can help enhance security by ensuring only necessary permissions are granted.

```sql+postgres
select
  name,
  created_on,
  granted_roles,
  granted_to_roles
from
  snowflake_role;
```

```sql+sqlite
select
  name,
  created_on,
  granted_roles,
  granted_to_roles
from
  snowflake_role;
```

### List idle roles
Discover the roles that are currently idle, meaning they are not assigned to any users. This can be useful for identifying potential areas of resource optimization or unnecessary access permissions.

```sql+postgres
select
  name,
  created_on,
  assigned_to_users
from
  snowflake_role
where
  assigned_to_users = 0;
```

```sql+sqlite
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
Explore which roles have been assigned to users in Snowflake, allowing you to manage user access and permissions effectively. This is useful in maintaining security and ensuring only authorized users have certain privileges.

```sql+postgres
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

```sql+sqlite
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