package opswatProvider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFileSyncResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configFileSyncResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_file_sync.new", "timeout", "10"),
				),
			},
		},
	})
}

const configFileSyncResource = `
resource "opswat_file_sync" "new" {
  timeout = 10
}
`
