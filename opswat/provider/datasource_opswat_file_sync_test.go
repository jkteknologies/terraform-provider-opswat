package opswatProvider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNewGlobalSyncDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + `data "opswat_file_sync" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify timeout returned
					resource.TestCheckResourceAttr("data.opswat_global_sync.test", "timeout", "10"),
				),
			},
		},
	})
}
