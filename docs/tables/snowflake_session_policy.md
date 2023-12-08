---
title: "Steampipe Table: snowflake_session_policy - Query Snowflake Session Policies using SQL"
description: "Allows users to query Snowflake Session Policies, specifically providing details about each session policy created in the Snowflake account."
---

# Table: snowflake_session_policy - Query Snowflake Session Policies using SQL

A Snowflake Session Policy is a feature in Snowflake that allows users to define and enforce specific settings or behaviors for a user session. These policies can include setting the time zone, date and time formats, and other session parameters. Snowflake Session Policies help in maintaining consistency and control over user sessions.

## Table Usage Guide

The `snowflake_session_policy` table provides insights into session policies within Snowflake. As a database administrator, explore specific details about each session policy, including its name, comment, and the values of the parameters it sets. Utilize it to uncover information about session policies, such as their current state, and the specific settings or behaviors they enforce.

## Examples

### Basic info
Identify the policies regarding session timeouts within your Snowflake environment. This can help manage idle sessions and optimize resource usage.

```sql+postgres
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy;
```

```sql+sqlite
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy;
```

### List policies with idle timeout more than an hour
Explore the policies that have an idle timeout exceeding one hour. This query can be useful in identifying potential areas for improved resource allocation and efficiency.

```sql+postgres
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy
where
  session_idle_timeout_mins > 60 or
  session_ui_idle_timeout_mins > 60;
```

```sql+sqlite
select
  name,
  database_name,
  schema_name,
  session_idle_timeout_mins,
  session_ui_idle_timeout_mins
from
  snowflake_session_policy
where
  session_idle_timeout_mins > 60 or
  session_ui_idle_timeout_mins > 60;
```