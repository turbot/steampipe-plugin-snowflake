---
title: "Steampipe Table: snowflake_schemata - Query Snowflake Schemata using SQL"
description: "Allows users to query Snowflake Schemata, specifically the database and schema information, providing insights into schema organization and associated metadata."
---

# Table: snowflake_schemata - Query Snowflake Schemata using SQL

Snowflake Schemata is a feature in Snowflake, a cloud-based data warehousing platform, that allows users to organize and manage data in logical groups. It provides a structure for storing, managing, and retrieving data, enabling efficient data operations. Snowflake Schemata helps users maintain control over their data, ensuring it is organized and easily accessible.

## Table Usage Guide

The `snowflake_schemata` table provides insights into the organization and structure of data within Snowflake. As a Data Engineer or Data Analyst, explore schema-specific details through this table, including the database name, schema name, and associated metadata. Utilize it to uncover information about schemas, such as their organization, the data they contain, and how they are used in data operations.

## Examples

### Basic info
Analyze the settings to understand the ownership and access details of various databases within your Snowflake account. This can help in determining which databases have managed access, are transient, and identify their respective owners, aiding in better data management and security.

```sql+postgres
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata;
```

```sql+sqlite
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata;
```

### List schemas that allow managed access
Explore which schemas in your Snowflake database are configured to allow managed access. This can help you maintain security and manageability of your data by identifying which schemas are under controlled access.

```sql+postgres
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

```sql+sqlite
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

### List transient schemas
Explore the Snowflake databases that are managed access to identify those that are transient. This can help in understanding the database structures that are temporary and managed by Snowflake, which is beneficial for maintaining data integrity and efficient resource usage.

```sql+postgres
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

```sql+sqlite
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  is_transient,
  schema_owner
from
  snowflake_schemata
where
  is_managed_access = 'YES';
```

### List schemas with a retention time greater than 15 days
Explore which schemas in your database have a retention time exceeding 15 days. This is useful for understanding and managing data storage and lifecycle within your system.

```sql+postgres
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  retention_time,
  schema_owner
from
  snowflake_schemata
where
  retention_time > 15;
```

```sql+sqlite
select
  schema_name,
  catalog_name as database_name,
  is_managed_access,
  retention_time,
  schema_owner
from
  snowflake_schemata
where
  retention_time > 15;
```