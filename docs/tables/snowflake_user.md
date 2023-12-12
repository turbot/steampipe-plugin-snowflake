---
title: "Steampipe Table: snowflake_user - Query OCI Snowflake Users using SQL"
description: "Allows users to query Snowflake Users, presenting detailed information about each user's properties, roles, and status."
---

# Table: snowflake_user - Query OCI Snowflake Users using SQL

Snowflake is a cloud-based data warehousing platform that enables data storage, processing, and analytic solutions. It is designed to support and manage all aspects of data, analytics, and application integration. Snowflake is built on top of the Amazon Web Services (AWS) cloud infrastructure and is a scalable and elastic solution that can handle high volumes of data and concurrent workloads.

## Table Usage Guide

The `snowflake_user` table provides insights into users within the OCI Snowflake service. As a data analyst or database administrator, explore user-specific details through this table, including user properties, roles, and status. Utilize it to manage user access, track user activity, and ensure compliance with your organization's data usage policies.

## Examples

### Basic info
Explore user profiles on your Snowflake platform to understand their access level and recent activity. This aids in maintaining security by identifying unusual behavior or inactive accounts.

```sql+postgres
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse,
  has_password,
  has_rsa_public_key,
  last_success_login
from
  snowflake_user;
```

```sql+sqlite
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse,
  has_password,
  has_rsa_public_key,
  last_success_login
from
  snowflake_user;
```

### List users that have passwords
Discover the segments that have passwords in the Snowflake user base to assess the elements within the user configuration. This can help in identifying instances where users may have potential security risks.

```sql+postgres
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse
from
  snowflake_user
where
  has_password;
```

```sql+sqlite
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse
from
  snowflake_user
where
  has_password = 1;
```

### List users whose passwords haven't been rotated in 90 days
Assess the elements within your user base to identify those who haven't updated their passwords in the past 90 days. This can be useful for enforcing security standards and ensuring regular password rotation.

```sql+postgres
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse,
  has_password,
  password_last_set_time::timestamp
from
  snowflake_user
where
  has_password
  and password_last_set_time::timestamp < now() - interval '90 days';
```

```sql+sqlite
select
  name,
  login_name,
  disabled,
  default_role,
  default_warehouse,
  has_password,
  datetime(password_last_set_time)
from
  snowflake_user
where
  has_password
  and datetime(password_last_set_time) < datetime('now','-90 days');
```

### List users using keypair authentication
Discover the segments that use keypair authentication to gain insights into the security measures in place for user access. This can be useful in assessing the strength and variety of authentication methods employed within your system.

```sql+postgres
select
  name,
  login_name,
  disabled,
  rsa_public_key,
  rsa_public_key_fp,
  rsa_public_key_2,
  rsa_public_key_2_fp
from
  snowflake_user
where
  has_rsa_public_key;
```

```sql+sqlite
select
  name,
  login_name,
  disabled,
  rsa_public_key,
  rsa_public_key_fp,
  rsa_public_key_2,
  rsa_public_key_2_fp
from
  snowflake_user
where
  has_rsa_public_key = 1;
```

### List users that have not logged in for 30 days
Identify users who haven't engaged with your platform in the last month. This can help in tailoring re-engagement strategies and understanding user activity patterns.

```sql+postgres
select
  name,
  email,
  disabled,
  last_success_login
from
  snowflake_user
where
  last_success_login is null
  or (last_success_login < now() - interval '30 days');
```

```sql+sqlite
select
  name,
  email,
  disabled,
  last_success_login
from
  snowflake_user
where
  last_success_login is null
  or (last_success_login < datetime('now', '-30 days'));
```