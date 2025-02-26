package opswatProvider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserRoleResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configUserRoleResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_user_role.new", "name", "TEST"),
				),
			},
		},
	})
}

const configUserRoleResource = `
resource "opswat_user_role" "new" {
  name         = "TEST"
  display_name = "TEST"
  rights = {
    #    scanlog     = []
    #    statistics  = []
    #    quarantine  = []
    #    updatelog   = []
    #    configlog   = []
    #    rule        = []
    #    workflow    = []
    #    zone        = []
    #    agents      = []
    #    engines     = []
    #    external    = []
    #    skip        = []
    #    cert        = []
    #    webhookauth = []
    #    retention   = []
    #    users       = []
    #    license     = []
    #    update      = []
    #    scan        = []
    #    healthcheck = []
    fetch    = ["selfonly"]
    download = ["selfonly"]
  }
}
`
