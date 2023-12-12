---
title: "Steampipe Table: snowflake_view - Query Snowflake Views using SQL"
description: "Allows users to query views in Snowflake, specifically providing detailed information about the views, their structure, and other related metadata."
---

# Table: snowflake_view - Query Snowflake Views using SQL

Snowflake is a cloud-based data warehousing platform that provides a comprehensive solution for data storage, processing, and analytics. It offers a unique architecture that separates storage and compute resources, allowing each to scale independently. One of the key features of Snowflake is its support for views, which are virtual tables based on the result-set of an SQL statement.

## Table Usage Guide

The `snowflake_view` table provides insights into views within Snowflake. As a data engineer, explore view-specific details through this table, including the view's definition, database, schema, and more. Utilize it to uncover information about views, such as their structure, the SQL statement used to create them, and other related metadata.

## Examples

### Basic info
Analyze the settings to understand the security status, creation date, and whether the views are materialized in your Snowflake database. This information can be useful for database management, particularly in maintaining security and optimizing performance.

```sql+postgres
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view;
```

```sql+sqlite
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view;
```

### List materialized views
Uncover the details of all materialized views within a Snowflake database, including their names, security status, and creation dates. This is useful for understanding the structure of your data and ensuring it is appropriately secured.

```sql+postgres
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_materialized;
```

```sql+sqlite
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_materialized;
```

### List secure views
Discover the segments that have secure views in your Snowflake database. This can help enhance security by identifying views that are configured to restrict data access.

```sql+postgres
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_secure;
```

```sql+sqlite
select
  name,
  database_name,
  schema_name,
  is_materialized,
  is_secure,
  created_on
from
  snowflake_view
where
  is_secure = 1;
```