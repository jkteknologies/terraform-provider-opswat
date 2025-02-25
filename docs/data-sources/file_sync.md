---
subcategory: "Qurantine"
layout: "opswat"
page_title: "OPSWAT: opswat_file_sync"
sidebar_current: "docs-opswat-datasource-file-sync"
description: |-
  OPSWAT Global sync connection timeout.
---

## Example Usage

```terraform
data "opswat_file_sync" "current" {}

output "opswat_file_sync" {
  value = data.opswat_file_sync.current
}
```

## Schema
Read-only:
- `timeout` (Int) OPSWAT Global sync connection timeout in minutes