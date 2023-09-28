package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccUserDirectoryDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + testAccExampleDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify dir returned
					resource.TestCheckResourceAttr("data.opswat_userdirectory.test", "userdirectory.0.id", "1"),
				),
			},
		},
	})
}

const testAccExampleDataSourceConfig = `
data "opswat_userdirectory" "test" {}
`
