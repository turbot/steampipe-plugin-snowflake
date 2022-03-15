# Table: snowflake_user_grant

List all privileges and roles granted to a User.

**Note** This table requires an '=' qualifier for `username` columns

### List all grants for a specific user

```sql
select
  username,
  role,
  granted_by,
  created_on
from
  snowflake_user_grant
where
  username = 'STEAMPIPE';
```
