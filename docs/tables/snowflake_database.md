---
title: "Steampipe Table: snowflake_database - Query Snowflake Databases using SQL"
description: "Allows users to query Snowflake Databases, providing detailed information about each database within the Snowflake data warehousing platform."
---

# Table: snowflake_database - Query Snowflake Databases using SQL

Snowflake is a cloud-based data warehousing platform that provides comprehensive solutions for data storage, processing, and analysis. It supports a wide range of data types, including structured and semi-structured data, and allows for seamless scaling of resources to meet the demands of any size of data workload. Snowflake databases are a key resource within this platform, housing the data that is processed and analyzed.

## Table Usage Guide

The `snowflake_database` table provides insights into databases within the Snowflake data warehousing platform. As a data engineer or analyst, explore database-specific details through this table, including ownership, creation time, and associated metadata. Utilize it to uncover information about databases, such as their size, the number of tables they contain, and their overall usage statistics.

## Examples

### Basic info
Explore the creation dates, current status, origin, owner, and retention time of your databases in Snowflake to understand their management and maintenance history. This is useful to assess the operational aspects of your databases and for future planning.

```sql+postgres
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database;
```

```sql+sqlite
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database;
```

### List databases with retention time greater than 1 day
Explore which databases have a retention time greater than a day. This is useful for understanding the longevity of your data and managing storage resources effectively.

```sql+postgres
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database
where
  retention_time > 1;
```

```sql+sqlite
select
  name,
  created_on,
  is_current,
  origin,
  owner,
  retention_time
from
  snowflake_database
where
  retention_time > 1;
```