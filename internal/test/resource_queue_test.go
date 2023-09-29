package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccQueueResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configQueueResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_queue.new", "max_queue_per_agent", "2500"),
				),
			},
		},
	})
}

const configQueueResource = `
resource "opswat_queue" "new" {
  max_queue_per_agent = 2500
}
`
