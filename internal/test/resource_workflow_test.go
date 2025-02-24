package opswatProvider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
)

func TestAccWorkflowResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: providerConfig + configWorkflowResource,
				Check: resource.ComposeTestCheckFunc(
					// Verify value returned
					resource.TestCheckResourceAttr("opswat_workflow.new", "name", "TEST"),
				),
			},
		},
	})
}

const configWorkflowResource = `
resource "opswat_workflow" "new" {
  allow_cert               = false
  allow_cert_cert          = "None"
  allow_cert_cert_validity = 1
  allow_local_files        = false
  allow_local_files_local_paths = []
  allow_local_files_white_list             = true
  description                              = "TEST workflow"
  include_webhook_signature                = false
  include_webhook_signature_certificate_id = 0
  mutable                                  = true
  name                                     = "TEST"
  option_values = {
    archive_handling_max_number_files           = 1000000
    archive_handling_max_recursion_level        = 10
    archive_handling_max_size_files             = 100000
    archive_handling_timeout                    = 25
    filetype_analysis_timeout                   = 25
    process_info_global_timeout                 = false
    process_info_global_timeout_value           = 600
    process_info_max_download_size              = 20000
    process_info_max_file_size                  = 20000
    process_info_quarantine                     = true
    process_info_skip_hash                      = true
    process_info_skip_processing_fast_symlink   = true
    process_info_workflow_priority              = 5
    scan_filescan_check_av_engine               = true
    scan_filescan_download_timeout              = 25
    scan_filescan_global_scan_timeout           = 25
    scan_filescan_per_engine_scan_timeout       = 25
    vul_filescan_timeout_vulnerability_scanning = 5
  }
  result_allowed = [
    {
      role = "1"
      visibility ="3"
    },
    {
      role = "#"
      visibility ="3"
    }
  ]
  scan_allowed = [4]
  workflow_id = 1
  zone_id     = 1
  user_agents = []
}
`
