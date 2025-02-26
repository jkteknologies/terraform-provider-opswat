package opswatProvider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserDirectoryDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configDirectoryDataSource,
				Check: resource.ComposeTestCheckFunc(
					// Verify dir returned
					resource.TestCheckResourceAttr("data.opswat_userdirectories.test", "dirs.0.id", "1"),
				),
			},
		},
	})
}

const configDirectoryDataSource = `
data "opswat_userdirectories" "test" {}
output "opswat_userdirectories" {
  value = data.opswat_userdirectories.test
}
`
