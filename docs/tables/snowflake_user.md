# Table: snowflake_user

A user is an identity recognized by Snowflake and can be associated with a person or program.

**Note**: This table can only be queried by users with a role that has the `MANAGE GRANTS` global privilege. This privilege is usually granted to the `ACCOUNTADMIN` and `SECURITYADMIN` roles.

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

### List users that have passwords

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

### List users whose passwords haven't been rotated in 90 days

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
  has_password
  and password_last_set_time::timestamp < now() - interval '90 days';
```

### List users using keypair authentication

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
  (last_success_login > now() - interval '30 days')
  and last_success_login is not null;
```
