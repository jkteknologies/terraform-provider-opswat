---
subcategory: "Access"
layout: "opswat"
page_title: "OPSWAT: opswat_user"
sidebar_current: "docs-opswat-resource-user"
description: |-
  OPSWAT User.
---

## Example Usage

```terraform
resource "opswat_user" "new" {
  # You can use /#/user/userManagement/userInventory -> Add user -> Generate Api key feature to create new apikey
  # APIKEY validation criteria:
  # The length of the API key must be exactly 36 characters.
  # must contain numeric and lower case a, b, c, d, e and f letter characters only
  # must contain at least 10 lower case a, b, c, d, e or f letter characters.
  # must contain at least 10 numeric characters.
  # allowed to contain at most 3 consecutive lower case letter characters (e.g. "abcd1a2b3c..." is invalid because of the four consecutive letters).
  # allowed to contain at most 3 consecutive numeric characters (e.g. "1234a1b2c3..." is invalid because of the four consecutive numeric characters).
  api_key = "xxx"
  directory_id = 1
  display_name = "TFUSER"
  email = "tfuser@av.swissre.cn"
  name = "TFUSER"
  roles = [1]
  # UI access password
  password = "xxx"
  # Ignore required to suppress empty password response for this api
  lifecycle {
    ignore_changes = [
      password,
    ]
  }
}
```

## Schema
Required:
- `api_key` - (String) APIKEY validation criteria:  
The length of the API key must be exactly 36 characters. 
It must contain numeric and lower case a, b, c, d, e and f letter characters only  
It must contain at least 10 lower case a, b, c, d, e or f letter characters.  
It must contain at least 10 numeric characters.  
It is allowed to contain at most 3 consecutive lower case letter characters (e.g. \"abcd1a2b3c...\" is invalid because of the four consecutive letters).  
It is allowed to contain at most 3 consecutive numeric characters (e.g. \"1234a1b2c3...\" is invalid because of the four consecutive numeric characters).
- `directory_id` - (Int) Base user template for mapping purposes.
- `display_name`- (String) User display name.
- `email` - (String) User email.
- `name` - (String) User name.
- `password` - (String) User UI access password.
- `roles` - (List of Ints) User role ids.