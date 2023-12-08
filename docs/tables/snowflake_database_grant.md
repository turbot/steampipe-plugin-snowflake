---
title: "Steampipe Table: snowflake_database_grant - Query Snowflake Database Grants using SQL"
description: "Allows users to query Snowflake Database Grants, specifically providing insights into the permissions assigned to various roles and users."
---

# Table: snowflake_database_grant - Query Snowflake Database Grants using SQL

Snowflake Database Grant is a feature within Snowflake's data cloud platform that allows you to manage and assign permissions to roles and users. It is a crucial aspect of Snowflake's security model, enabling you to control who has access to your data and what they can do with it. Snowflake Database Grant helps you maintain a secure and compliant data environment by ensuring the right access levels are assigned to the right roles and users.

## Table Usage Guide

The `snowflake_database_grant` table provides insights into the database grants within Snowflake's data cloud platform. As a Database Administrator, you can explore grant-specific details through this table, including roles, users, and the specific permissions assigned to them. Utilize it to uncover information about database access levels, such as those with full permissions, the roles assigned to specific users, and the verification of user privileges.

## Examples

### Basic info
Explore which privileges have been granted to different users in a specific database. This can help in managing user access and maintaining database security.

```sql+postgres
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant where database = 'SNOWFLAKE';
```

```sql+sqlite
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant where database = 'SNOWFLAKE';
```

### List grants for all databases
Explore which privileges have been granted to various users across all databases. This can be useful for assessing security measures and understanding user access levels in a Snowflake environment.

```sql+postgres
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant
  inner join
    snowflake_database
    on snowflake_database_grant.database = snowflake_database.name;
```

```sql+sqlite
select
  database,
  privilege,
  grantee_name,
  granted_to,
  grant_option
from
  snowflake_database_grant
  inner join
    snowflake_database
    on snowflake_database_grant.database = snowflake_database.name;
```