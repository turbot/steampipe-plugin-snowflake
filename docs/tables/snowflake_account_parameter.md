---
title: "Steampipe Table: snowflake_account_parameter - Query Snowflake Account Parameters using SQL"
description: "Allows users to query Snowflake Account Parameters, specifically retrieving the current settings and defaults for various parameters in Snowflake accounts."
---

# Table: snowflake_account_parameter - Query Snowflake Account Parameters using SQL

Snowflake Account Parameters are a collection of settings and defaults that govern the behavior of Snowflake accounts. These parameters include settings related to data storage, query processing, security, and other operational aspects of Snowflake accounts. They provide a way to customize and tune the behavior of Snowflake accounts to meet specific requirements or preferences.

## Table Usage Guide

The `snowflake_account_parameter` table provides insights into the settings and defaults of Snowflake Account Parameters. As a Database Administrator or Data Engineer, explore parameter-specific details through this table, including names, values, and descriptions. Utilize it to uncover information about parameters, such as their current settings, default values, and the impact of these parameters on the operation of Snowflake accounts.

## Examples

### Basic info
Explore which account parameters are set in your Snowflake account to understand and manage your account configurations better. This can be particularly useful when auditing your account settings or troubleshooting issues related to account parameters.

```sql+postgres
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter;
```

```sql+sqlite
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter;
```

### Check whether account allows MFA caching
Assess the elements within your account to understand if multi-factor authentication (MFA) caching is permitted. This is useful for enhancing security measures by managing how user authentication data is stored.

```sql+postgres
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'ALLOW_CLIENT_MFA_CACHING';
```

```sql+sqlite
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'ALLOW_CLIENT_MFA_CACHING';
```

### Get number of days Snowflake retains historical data for performing Time Travel actions (SELECT, CLONE, UNDROP) on the object
Analyze the settings to understand the duration for which Snowflake preserves historical data, which can be crucial for executing actions such as Time Travel on objects. This can be beneficial for data recovery and auditing purposes.

```sql+postgres
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'DATA_RETENTION_TIME_IN_DAYS';
```

```sql+sqlite
select
  key,
  value,
  level,
  description
from
  snowflake_account_parameter
where
  key = 'DATA_RETENTION_TIME_IN_DAYS';
```