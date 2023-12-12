---
title: "Steampipe Table: snowflake_account_grant - Query Snowflake Account Grants using SQL"
description: "Allows users to query Snowflake Account Grants, specifically the grantee name, granted on date, and privilege details, providing insights into account-level access permissions."
---

# Table: snowflake_account_grant - Query Snowflake Account Grants using SQL

Snowflake Account Grants are resources within Snowflake that allow you to manage and monitor permissions granted at the account level. These permissions can be granted to roles, users, or other entities within the Snowflake environment. The account grant includes details about the grantee, the granted on date, and the specific privilege granted.

## Table Usage Guide

The `snowflake_account_grant` table provides insights into account-level permissions within Snowflake. As a Security Analyst, explore grant-specific details through this table, including the grantee name, granted on date, and privilege details. Utilize it to uncover information about permissions, such as who has been granted what privileges, when the privileges were granted, and the specifics of the privileges.

## Examples

### Basic info
Explore the details of your Snowflake account's access permissions to understand who has been granted what privileges, by whom, and when. This can help in maintaining security and compliance by ensuring appropriate access levels are maintained.

```sql+postgres
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

```sql+sqlite
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
Explore which privileges are associated with the account administrator role. This can be useful for understanding the level of access and permissions granted to this role within your Snowflake account.

```sql+postgres
select
  privilege,
  grant_option,
  created_on
from
  snowflake_account_grant
where
  grantee_name = 'ACCOUNTADMIN';
```

```sql+sqlite
select
  privilege,
  grant_option,
  created_on
from
  snowflake_account_grant
where
  grantee_name = 'ACCOUNTADMIN';
```