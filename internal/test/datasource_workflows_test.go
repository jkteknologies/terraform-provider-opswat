package opswatProvider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWorkflowsDataSource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configWorkflowsDataSource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestMatchResourceAttr("data.opswat_workflows.test", "workflows.0.id", regexp.MustCompile("[0-9]+")),
				),
			},
		},
	})
}

const configWorkflowsDataSource = `
data "opswat_workflows" "test" {}
output "opswat_workflows" {
  value = data.opswat_workflows.test
}
`
