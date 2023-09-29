---
subcategory: "Scan"
layout: "opswat"
page_title: "OPSWAT: opswat_queue"
sidebar_current: "docs-opswat-resource-queue"
description: |-
  OPSWAT Scan Core agent queue count.
---

## Example Usage

```terraform
resource "opswat_queue" "current" {
  max_queue_per_agent = 2500
}
```

## Schema
Required:  
- `max_queue_per_agent` (Int) Core agent scan queue count.