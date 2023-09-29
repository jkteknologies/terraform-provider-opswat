package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccSessionResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configSessionResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_session.new", "absolute_session_timeout", "0"),
					resource.TestCheckResourceAttr("opswat_session.new", "allow_crossip_sessions", "true"),
					resource.TestCheckResourceAttr("opswat_session.new", "allow_duplicate_session", "true"),
					resource.TestCheckResourceAttr("opswat_session.new", "session_timeout", "0"),
				),
			},
		},
	})
}

const configSessionResource = `
resource "opswat_session" "new" {
  absolute_session_timeout = 0
  allow_crossip_sessions   = true
  allow_duplicate_session  = true
  session_timeout          = 0
}
`
