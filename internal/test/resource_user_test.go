package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccUserResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configUserResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_user.new", "name", "TEST"),
				),
			},
		},
	})
}

const configUserResource = `
resource "opswat_user" "new" {
  api_key = "e4b6e5d6d0e40df0da5e6b2aa3bc98b29d25"
  directory_id = 1
  display_name = "TEST"
  email = "null@opswat.xxx.com"
  name = "TEST"
  roles = [
    "5" // project related role
  ]
  password = "abc"
  lifecycle {
    ignore_changes = [
      password,
    ]
  }
}
`
