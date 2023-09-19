---
page_title: "OPSWAT Provider"  
subcategory: ""  
description: |-
---

# OPSWAT Provider
OPSWAT MetaDefenter community provider to configure environment.  

Refs:
- docs - https://docs.opswat.com/mdcore
- api - https://docs.opswat.com/mdcore/metadefender-core

## Configuration
```hcl
provider "opswat" {
  host   = "https://opswat.xxx.com"
  apikey = "xxx"
}
```

## Schema
- apikey (Required)  
  OPSWAT MetaDefender administrative API key
- host (Required)  
  OPSWAT MetaDefender base hostname (ex: https://opswat.xxx.com)