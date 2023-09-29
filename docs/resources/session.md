---
subcategory: "Scan"
layout: "opswat"
page_title: "OPSWAT: opswat_session"
sidebar_current: "docs-opswat-resource-session"
description: |-
  OPSWAT Global session timeouts resource.
---

## Example Usage

```terraform
resource "opswat_session" "current" {
  absolute_session_timeout = 0
  allow_crossip_sessions = true
  allow_duplicate_session = true
  session_timeout = 0
}
```

## Schema
Required:  
- `absolute_session_timeout` (Int) The interval (in milliseconds) for overall session length timeout (regardless of activity). minimal 300000. 0 - for infinity sessions..
- `allow_crossip_sessions` (Bool) Allow requests from the same user to come from different IPs flag.
- `allow_duplicate_session` (Bool) Allow same user to have multiple active sessions.
- `session_timeout` (Int) The interval (in milliseconds) for the user's session timeout, based on last activity. Timer starts after the last activity for the apikey. minimal - 60000. 0 - for infinity sessions.