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
resource "opswat_scan_history" "current" {
  cleanuprange = 24
}
```

## Schema

Required:

- `cleanuprange` (Int) Setting processing history clean up time (clean up records older than). Note:The clean up range
  is defined in hours.