---
page_title: ` -opswat_userdirectory Resource - terraform-provider-opswat` -
subcategory: ` -` -
description: OPSWAT User directories
  
---

## Example Usage

```terraform
data "opswat_userdirectory" "current" {}

output "opswat_userdirectory" {
  value = data.opswat_userdirectory.current
}
```

## Schema
Read-only:
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
  - `x509_cert` - (String) Valid certificate from https://login.microsoftonline.com/xxx/federationmetadata/2007-06/federationmetadata.xml -> EntityDescriptor -> Signature -> KeyInfo -> X509Data -> X509Certificate (For SAML integration)
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
- `user_identified_by` - (String) User id/name mapping for UI (claim based) ["$${http://schemas.microsoft.com/identity/claims/displayname}"]