---
title: "Steampipe Table: snowflake_view_grant - Query Snowflake View Grants using SQL"
description: "Allows users to query Snowflake View Grants, providing insights into the permissions granted to specific views within a Snowflake database."
---

# Table: snowflake_view_grant - Query Snowflake View Grants using SQL

Snowflake View Grants represent permissions that are granted to specific views within a Snowflake database. They are a crucial aspect of managing security and access control in Snowflake, ensuring that only authorized users or roles can access and manipulate specific views. They can be used to grant or revoke privileges such as SELECT, INSERT, UPDATE, DELETE, TRUNCATE, REFERENCES, or TRANSFER OWNERSHIP on a specific view to a specific role.

## Table Usage Guide

The `snowflake_view_grant` table provides insights into the permissions granted to specific views within a Snowflake database. As a database administrator or security officer, explore the details of these grants through this table, including the granted role, privilege type, and associated metadata. Utilize it to manage and monitor access control, ensuring that only authorized users or roles can access and manipulate specific views.

## Examples

### Basic info
Explore which privileges have been granted to specific users within the Snowflake database. This could be particularly useful for administrators looking to review access controls for security purposes.

```sql+postgres
select
  view_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
where
  view_name = 'ROLES'
  and database_name = 'SNOWFLAKE'
  and schema_name = 'ACCOUNT_USAGE';
```

```sql+sqlite
select
  view_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
where
  view_name = 'ROLES'
  and database_name = 'SNOWFLAKE'
  and schema_name = 'ACCOUNT_USAGE';
```

### List view grants for `ACCOUNT_USAGE` schema in `SNOWFLAKE` database
Discover the segments that have been granted access to view certain data in a specific Snowflake database schema. This query is useful for auditing and managing data access permissions in your organization.

```sql+postgres
select
  view_name,
  snowflake_view.database_name,
  snowflake_view.schema_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
  inner join
    snowflake_view
    on snowflake_view_grant.view_name = snowflake_view.name
    and snowflake_view_grant.database_name = snowflake_view.database_name
    and snowflake_view_grant.schema_name = snowflake_view.schema_name
where
  snowflake_view_grant.database_name = 'SNOWFLAKE'
  and snowflake_view_grant.schema_name = 'ACCOUNT_USAGE';
```

```sql+sqlite
select
  view_name,
  snowflake_view.database_name,
  snowflake_view.schema_name,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_view_grant
  inner join
    snowflake_view
    on snowflake_view_grant.view_name = snowflake_view.name
    and snowflake_view_grant.database_name = snowflake_view.database_name
    and snowflake_view_grant.schema_name = snowflake_view.schema_name
where
  snowflake_view_grant.database_name = 'SNOWFLAKE'
  and snowflake_view_grant.schema_name = 'ACCOUNT_USAGE';
```