---
title: "Steampipe Table: snowflake_network_policy - Query OCI Snowflake Network Policies using SQL"
description: "Allows users to query Network Policies in Snowflake, providing insights into network access control configurations and potential vulnerabilities."
---

# Table: snowflake_network_policy - Query OCI Snowflake Network Policies using SQL

A Network Policy in Snowflake is a set of rules that govern the network access control for virtual warehouses. It allows administrators to define IP whitelisting rules to restrict access to the Snowflake account only from allowed IP addresses. Network Policies can be associated with individual users or the entire account.

## Table Usage Guide

The `snowflake_network_policy` table provides insights into Network Policies within OCI Snowflake. As a Network Administrator, explore policy-specific details through this table, including allowed IP addresses, blocked IP addresses, and associated metadata. Use it to uncover information about policies, such as those with unrestricted access, the IP address restrictions in place, and the verification of network control configurations.

## Examples

### Basic info
Explore which network policies have been implemented within Snowflake, focusing on when they were created and the number of entries in both the allowed and blocked IP lists. This can help identify potential security gaps and understand the overall network security posture.

```sql+postgres
select
  name,
  comment,
  created_on,
  entries_in_allowed_ip_list,
  entries_in_blocked_ip_list
from
  snowflake_network_policy;
```

```sql+sqlite
select
  name,
  comment,
  created_on,
  entries_in_allowed_ip_list,
  entries_in_blocked_ip_list
from
  snowflake_network_policy;
```

### Get blocked and allowed IP lists for a specific network policy
Analyze the settings to understand the blocked and allowed IP addresses associated with a specific network policy. This can help in assessing the security measures and identifying any potential vulnerabilities in the network access.

```sql+postgres
select
  name,
  allowed_ip_list,
  blocked_ip_list
from
  snowflake_network_policy
where
  name = 'np1';
```

```sql+sqlite
select
  name,
  allowed_ip_list,
  blocked_ip_list
from
  snowflake_network_policy
where
  name = 'np1';
```