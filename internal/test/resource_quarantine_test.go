package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccQuarantineResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configQuarantineResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_quarantine.new", "cleanup_range", "168"),
				),
			},
		},
	})
}

const configQuarantineResource = `
resource "opswat_quarantine" "new" {
  cleanup_range = 168
}
`
