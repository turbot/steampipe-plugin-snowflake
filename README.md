![image](https://hub.steampipe.io/images/plugins/turbot/snowflake-social-graphic.png)

# Snowflake Plugin for Steampipe

Use SQL to query roles, databases, and more from Snowflake.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/snowflake)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/snowflake/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-snowflake/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install snowflake
```

Run a query:

```sql
select
  name,
  email,
  disabled,
  last_success_login
from
  snowflake_user
where
  (last_success_login > now() - interval '30 days')
  and last_success_login is not null;
```
## Examples

```sql
select account as "Account",
    name as "Warehouse",
    credit_quota as "Credit Quota",
    used_credits as "Used Credits",
    remaining_credits as "Remaining Credits",
    to_char(used_credits/credit_quota*100,'9990D9') as "% Used",
    case
        when used_credits/credit_quota*100 > 90 then 'alert'
        when used_credits/credit_quota*100 > 75 then 'warning'
        else 'ok'
    end as type
from snowflake_resource_monitor
where used_credits/credit_quota*100 > 75
order by used_credits/credit_quota desc
```

```sql
select account as "Account",
    name as "Warehouse"
from snowflake_resource_monitor
where remaining_credits < 1
```
  
## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-snowflake.git
cd steampipe-plugin-snowflake
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/snowflake.spc
```

Try it!

```
steampipe query
> .inspect snowflake
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-snowflake/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Snowflake Plugin](https://github.com/turbot/steampipe-plugin-snowflake/labels/help%20wanted)
