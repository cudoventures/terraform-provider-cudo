---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cudo_ssh_keys Data Source - cudo-terraform-provider-pf"
subcategory: ""
description: |-
  SshKeys data source
---

# cudo_ssh_keys (Data Source)

SshKeys data source



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) Placeholder identifier attribute.
- `images` (Attributes List) List of SSH keys (see [below for nested schema](#nestedatt--images))

<a id="nestedatt--images"></a>
### Nested Schema for `images`

Read-Only:

- `comment` (String) SSH key comment
- `fingerprint` (String) SSH key finger print
- `id` (String) SSH key identifier
- `public_key` (String) SSH key public key
- `type` (String) SSH key type

