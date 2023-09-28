---
page_title: "opswat_global_sync Resource - terraform-provider-opswat"
subcategory: ""
description: OPSWAT Global sync connection timeout
  
---

## Example Usage

```terraform
resource "opswat_file_sync" "new" {
  timeout = 10
}
```


## Argument Reference
- `timeout` (Required) OPSWAT Global sync connection timeout in minutes