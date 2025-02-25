---
subcategory: "SSO"
layout: "opswat"
page_title: "OPSWAT: opswat_userdirectory"
sidebar_current: "docs-opswat-resource-userdirectory"
description: |-
  OPSWAT User directory.
---

-> NOTE: Only SAML integration supported.

## Example Usage

```terraform
resource "opswat_userdirectory" "new" {
  name = "TEST"
  type = "Security Assertion Markup Language (SAML)"
  version = "2.0"
  enabled = false
  idp = {
    authn_request_signed = false
    entity_id = "https://sts.windows.net/xxx/"
    login_method = {
      post = "https://login.microsoftonline.com/xxx/saml2"
      redirect = "https://login.microsoftonline.com/xxx/saml2"
    }
    logout_method = {
      redirect = "https://login.microsoftonline.com/xxx/saml2"
    }
    valid_until = ""
    x509_cert = ["xxx"]
  }
  role = {
    details = [{
      key = "http://schemas.microsoft.com/ws/2008/06/identity/claims/role"
      values = [{
        condition = "admin$"
        role_ids = ["1"]
        type = "regex"
      }]
    }]
    option = "role mapping"
  }
  sp = {
    enable_idp_initiated = false
    entity_id = "xxx"
    login_url = "https://opswat.xxx.com/ssologin/saml/XXXX"
    support_entity_id = true
    support_logout_url = false
    support_private_key = false
  }
  user_identified_by = "${http://schemas.microsoft.com/identity/claims/displayname}"
}
```

## Schema
Required:  
- `name` - (String) SSO configuration name (uppercase)
- `type` - (String) SSO configuration type ["OpenID Connect (OIDC)", "Security Assertion Markup Language (SAML)"]
- `version`- (String) Configuration version "2.0" for SAML "1.0" for OIDC
- `enabled` - (Bool) Enable flag for SSO current configuration
- `idp` - (Nested object) IDP configuration
    - `authn_request_signed` - (Bool) Enable Assertion Decryption flag (private key must be provided)
    - `entity_id` - (String) IDP Directory (tenant) ID full url
    - `login_method`
        - `post` - (String) SAML-P sign-on endpoint (For SAML integration)
        - `redirect` - (String) SAML-P sign-on endpoint (For SAML integration)
    - `logout_method`
        - `redirect` - (String) SAML-P sign-on endpoint (For SAML integration)
    - `valid_until` - (String) ""
    - `x509_cert` - (List of strings) Valid certificates from https://login.microsoftonline.com/xxx/federationmetadata/2007-06/federationmetadata.xml -> EntityDescriptor -> Signature -> KeyInfo -> X509Data -> X509Certificate (For SAML integration)
- `role` - (Nested object) Roles mapping
    - `details` - (List of objects)
        - `key` - (String) Role claim (For example: "http://schemas.microsoft.com/ws/2008/06/identity/claims/role" for Azure AD SAML integration)
        - `values` -
            - `condition` - (String) Role mapping regexp (Example: admin$)
            - `role_ids` - (List of ints) Role ID to map IDP role to
            - `type` - (String) "regex"
    - `option` - (String) "role mapping"
- `sp`
    - `enable_idp_initiated` - (Bool) Sign in MetaDefender Core via Identity Provider site flag
    - `entity_id` - (String) IDP Application client ID
    - `login_url` - (String) IDP redirect url [hostname + /ssologin/saml/XXXX]
    - `support_entity_id` - (Bool) Custom entity (client id) usage [should be true in case of SAML integration]
    - `support_logout_url` - (Bool) Custom logout url support [might be false]
    - `support_private_key` - (Bool) Private key flag [might be false]
- `user_identified_by` - (String) User id/name mapping for UI (claim based) ["${http://schemas.microsoft.com/identity/claims/displayname}"]