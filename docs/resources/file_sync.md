---
subcategory: "Connection"
layout: "opswat"
page_title: "OPSWAT: opswat_file_sync"
sidebar_current: "docs-opswat-resource-file-sync"
description: |-
  OPSWAT Global sync connection timeout
---

## Example Usage

```terraform
resource "opswat_file_sync" "current" {
  timeout = 10
}
```

## Schema
Required:  
- `timeout` (Int) OPSWAT Global sync connection timeout in minutes