---
layout: "scaleway"
page_title: "Scaleway: security_group"
sidebar_current: "docs-scaleway-resource-security-group-x"
description: |-
  Manages Scaleway security groups.
---

# scaleway_security_group

Provides security groups. This allows security groups to be created, updated and deleted.
For additional details please refer to [API documentation](https://developer.scaleway.com/#security-groups).

## Example Usage

```hcl
resource "scaleway_security_group" "test" {
  name                    = "test"
  description             = "test"
  enable_default_security = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) name of security group
* `description` - (Required) description of security group
* `enable_default_security` - (Optional) default: true. Add default security group rules

Field `name`, `description` are editable.

## Attributes Reference

The following attributes are exported:

* `id` - id of the new resource

## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import scaleway_security_group.test 5faef9cd-ea9b-4a63-9171-9e26bec03dbc
```
