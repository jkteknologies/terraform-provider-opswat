---
subcategory: "Qurantine"
layout: "opswat"
page_title: "OPSWAT: opswat_quarantine"
sidebar_current: "docs-opswat-resource-quarantine"
description: |-
  OPSWAT Quarantine clean up time.
---

## Example Usage

```terraform
resource "opswat_quarantine" "current" {
  cleanup_range = 168
}
```

## Schema
Required:  
- `timeout` (Int) Quarantine clean up time (clean up records older than). The clean up range is defined in hours.