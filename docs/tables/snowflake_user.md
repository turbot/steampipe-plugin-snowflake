# Table: snowflake_user

A user identity recognized by Snowflake, whether associated with a person or program.

## Examples

### Basic info

```sql
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

### Users that have passwords in Snowflake

```sql
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

### Users with passwords haven't changed password for last 90 days

```sql
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
  has_password and
  password_last_set_time::timestamp < now() - interval '90 days';
```

### Users using keypair authentication

```sql
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

### List inactive users

```sql
select
  name,
  email,
  disabled,
  default_role,
  has_password,
  has_rsa_public_key
from
  snowflake_user
where
  (last_success_login > now() - interval '30 days') and
  last_success_login is not null
```