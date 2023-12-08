---
title: "Steampipe Table: snowflake_login_history - Query Snowflake Login Histories using SQL"
description: "Allows users to query Snowflake Login Histories, specifically providing insights into user login events and activity patterns."
---

# Table: snowflake_login_history - Query Snowflake Login Histories using SQL

Snowflake Login History is a feature within Snowflake that allows you to monitor and track user login events across your Snowflake environment. It provides a comprehensive log of user login activity, including details such as user name, login time, IP address, and more. Snowflake Login History helps you stay informed about the login activities in your Snowflake environment and take appropriate actions when anomalies are detected.

## Table Usage Guide

The `snowflake_login_history` table provides insights into user login events within Snowflake. As a Security Analyst, explore user-specific login details through this table, including user name, login time, IP address, and more. Utilize it to uncover information about user login activities, such as login frequency, login timings, and the source IP addresses of the logins.

## Examples

### Basic info
Analyze login history to understand the success rate of user authentications and pinpoint specific instances where the first authentication factor was used. This could be beneficial in assessing the security measures and identifying potential vulnerabilities.

```sql+postgres
select
  user_name,
  first_authentication_factor,
  is_success,
  event_timestamp
from
  snowflake_login_history;
```

```sql+sqlite
select
  user_name,
  first_authentication_factor,
  is_success,
  event_timestamp
from
  snowflake_login_history;
```

### List all authentication methods used in the last 30 days
Explore the variety of successful authentication methods utilized by users in the past month. This can provide insights into user behaviour and security practices, aiding in the enhancement of system security protocols.

```sql+postgres
select distinct
  user_name,
  first_authentication_factor
from
  snowflake_login_history
where
  is_success = 'YES'
  and event_timestamp > now() - interval '30 days'
order by
  user_name;
```

```sql+sqlite
select distinct
  user_name,
  first_authentication_factor
from
  snowflake_login_history
where
  is_success = 'YES'
  and event_timestamp > datetime('now', '-30 days')
order by
  user_name;
```