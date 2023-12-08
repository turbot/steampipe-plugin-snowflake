---
title: "Steampipe Table: snowflake_session - Query Snowflake Sessions using SQL"
description: "Allows users to query Snowflake Sessions, providing detailed insights into session activities, user actions, and session-specific metadata."
---

# Table: snowflake_session - Query Snowflake Sessions using SQL

Snowflake Sessions are an integral part of the Snowflake Data Cloud platform, which allows users to manage and execute SQL queries. Each session represents a connection from a client to the Snowflake service, and it is used to execute SQL statements. The sessions are ephemeral and are automatically terminated after a period of inactivity or when the client disconnects.

## Table Usage Guide

The `snowflake_session` table provides comprehensive insights into Snowflake sessions. As a database administrator or data engineer, you can use this table to explore session-specific details, including user actions, session duration, and associated metadata. This can be particularly useful for monitoring user activity, optimizing session performance, and troubleshooting issues related to session connectivity.

## Examples

### Basic info
Explore which users have logged into your Snowflake environment and when, allowing you to monitor user activity and understand usage patterns. This information can be particularly useful for auditing and security purposes.

```sql+postgres
select
  session_id,
  user_name,
  authentication_method,
  created_on,
  client_environment ->> 'APPLICATION' as client_application
from
  snowflake_session;
```

```sql+sqlite
select
  session_id,
  user_name,
  authentication_method,
  created_on,
  json_extract(client_environment, '$.APPLICATION') as client_application
from
  snowflake_session;
```

### List distinct authentication methods used in the last year
Explore which unique authentication methods have been used by different users in the past year. This can help in understanding user behavior and enhancing security measures.

```sql+postgres
select distinct
  user_name,
  authentication_method
from
  snowflake_session
order by
  user_name;
```

```sql+sqlite
select distinct
  user_name,
  authentication_method
from
  snowflake_session
order by
  user_name;
```

### List sessions authenticated without Snowflake MFA with passsword in last 30 days
Explore which user sessions have been authenticated without the use of Snowflake multi-factor authentication (MFA) and with a password in the past 30 days. This query can help identify potential security risks and enforce stricter authentication methods.

```sql+postgres
select distinct
  user_name,
  authentication_method,
  client_environment ->> 'APPLICATION' as client_application
from
  snowflake_session
where
  split_part(authentication_method, '+', 2) = ''
  and authentication_method like 'Password%'
  and created_on > now() - interval '30 days'
order by
  user_name desc;
```

```sql+sqlite
Error: SQLite does not support split_part function.
```