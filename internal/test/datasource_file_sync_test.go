package opswatProvider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccFileSyncDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configFileSyncDataSource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestMatchResourceAttr("data.opswat_file_sync.test", "timeout", regexp.MustCompile("[0-9]+")),
				),
			},
		},
	})
}

const configFileSyncDataSource = `
data "opswat_file_sync" "test" {}
output "opswat_file_sync" {
  value = data.opswat_file_sync.test
}
`
