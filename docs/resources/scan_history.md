---
subcategory: "Scan"
layout: "opswat"
page_title: "OPSWAT: opswat_scan_history"
sidebar_current: "docs-opswat-resource-scan-history"
description: |-
  OPSWAT processing history clean up time
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