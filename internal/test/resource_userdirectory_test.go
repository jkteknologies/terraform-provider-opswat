package opswatProvider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserDirectoryResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configUserDirectoryResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_userdirectory.new", "name", "TEST03"),
				),
			},
		},
	})
}

const configUserDirectoryResource = `
resource "opswat_userdirectory" "new" {
  name = "TEST03"
  type = "Security Assertion Markup Language (SAML)"  
  version = "2.0"
  enabled = false
  idp = {
    authn_request_signed = false
    entity_id = "https://sts.windows.net/45597f60-6e37-4be7-acfb-4c9e23b261ea/"
    login_method = {
      post = "https://login.microsoftonline.com/45597f60-6e37-4be7-acfb-4c9e23b261ea/saml2"
      redirect = "https://login.microsoftonline.com/45597f60-6e37-4be7-acfb-4c9e23b261ea/saml2"
    }
    logout_method = {
      redirect = "https://login.microsoftonline.com/45597f60-6e37-4be7-acfb-4c9e23b261ea/saml2"
    }
    valid_until = ""
    x509_cert = ["MIIC9jCCAl+gAwIBAgIJAParOnPwEkKlMA0GCSqGSIb3DQEBBQUAMIGKMQswCQYDVQQGEwJMSzEQMA4GA1UECBMHV2VzdGVybjEQMA4GA1UEBxMHQ29sb21ibzEWMBQGA1UEChMNU29mdHdhcmUgVmlldzERMA8GA1UECxMIVHJhaW5pbmcxLDAqBgNVBAMTI1NvZnR3YXJlIFZpZXcgQ2VydGlmaWNhdGUgQXV0aG9yaXR5MB4XDTEwMDcxMDA2MzMxOFoXDTI0MDMxODA2MzMxOFowcjELMAkGA1UEBhMCTEsxEDAOBgNVBAgTB1dlc3Rlcm4xEDAOBgNVBAcTB0NvbG9tYm8xFjAUBgNVBAoTDVNvZnR3YXJlIFZpZXcxETAPBgNVBAsTCFRyYWluaW5nMRQwEgYDVQQDEwtKYWNrIERhbmllbDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAqAIsXru2kWzNXidrgyapDb7GdmhUwNFx1rOimDyu2RrJN9sIv0Zi2B0Kp1xSQiBPWabXbtt3wB1LzS2P19tMC+MW7BTYz0mRg4n9vSoa+mTJ3Ea6/v4W97a701BSEOlTxysVltqgO+D3gD9uNVpjiCNjXP3FlXrw44aDnXwme3sCAwEAAaN7MHkwCQYDVR0TBAIwADAdBgNVHQ4EFgQUDp+pbeXQHmYiubDctF8b+C4g6V0wHwYDVR0jBBgwFoAU1rdiaEM7sE7BtSqZhTWT9Tqn9RQwLAYJYIZIAYb4QgENBB8WHU9wZW5TU0wgR2VuZXJhdGVkIENlcnRpZmljYXRlMA0GCSqGSIb3DQEBBQUAA4GBACcLqPwC9cATSqe+Kes5r6kcgo8eN3QME+HVSQocFSaRVrZ8iOrl0NAXway2JOGdjIFCn2gU4NAkrDAzjJ1AlwrfCT/1FDL5hu4BTdY13ZpwBf5MU6LB6x2tc+Jbo4bQrskEEIfGpOcyuB/wBJtJQeONjLuY2ouX9pvaaHj2cpzS"] }
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
  user_identified_by = "$${http://schemas.microsoft.com/identity/claims/displayname}"
}
`
