---
title: "Steampipe Table: snowflake_role_grant - Query Snowflake Role Grants using SQL"
description: "Allows users to query Snowflake Role Grants, specifically providing details about the roles and the privileges granted to them."
---

# Table: snowflake_role_grant - Query Snowflake Role Grants using SQL

Snowflake Role Grant is a feature within Snowflake that allows you to manage and control access to database objects. It provides a way to assign privileges to roles, which can then be granted to users or other roles. Snowflake Role Grant helps you ensure appropriate access levels and permissions across your Snowflake resources.

## Table Usage Guide

The `snowflake_role_grant` table provides insights into Role Grants within Snowflake. As a Database Administrator, you can use this table to explore details about roles and the privileges granted to them. This information can be used to manage access control, enforce security policies, and audit role-based permissions in your Snowflake environment.

## Examples

### List users granted the ACCOUNTADMIN role
Determine the users who have been granted the highest level of access within your system. This is useful for auditing security and managing permissions.

```sql+postgres
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

```sql+sqlite
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
Identify instances where the SYSADMIN role has been granted. This query is useful for understanding who has been given this high-level access and by whom, helping maintain security and manage permissions effectively.

```sql+postgres
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

```sql+sqlite
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