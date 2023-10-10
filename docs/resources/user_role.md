---
subcategory: "Access"
layout: "opswat"
page_title: "OPSWAT: opswat_userrole"
sidebar_current: "docs-opswat-resource-userrole"
description: |-
  OPSWAT User role.
---

## Example Usage

```terraform
resource "opswat_user_role" "new" {
  name         = "TFROLE"
  display_name = "TFROLE"
  rights = {
    fetch    = ["selfonly"]
    download = ["selfonly"]
  }
}
```

## Schema
Required:
- `name` - (String) User role name.
- `display_name` - (String) User role display name.
- `rights` - (Map of Lists) Access rights for specific features. 'fetch - Processing result fetching' and 'download - Download processed file' are only supported features as for 0.1.8 provider version. Access right options: ["selfonly", "anyone"]