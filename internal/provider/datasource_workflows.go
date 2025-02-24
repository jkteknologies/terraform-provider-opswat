package opswatProvider

import (
	"context"
	"fmt"
	opswatClient "terraform-provider-opswat/internal/connectivity"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &workflows{}
	_ datasource.DataSourceWithConfigure = &workflows{}
)

// NewGlobalWorkflows is a helper function to simplify the provider implementation.
func NewWorkflows() datasource.DataSource {
	return &workflows{}
}

// workflows is the data source implementation.
type workflows struct {
	client *opswatClient.Client
}

// Metadata returns the data source type name.
func (d *workflows) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflows"
}

// Schema defines the schema for the data source.
func (d *workflows) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global workflows datasource.",
		Attributes: map[string]schema.Attribute{
			"workflows": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_cert": schema.BoolAttribute{
							Description: "Generate batch signature with certificate - Use certificate to generate batch signature flag.",
							Computed:    true,
						},
						"allow_cert_cert": schema.StringAttribute{
							Description: "Certificate used for barch signing.",
							Computed:    true,
						},
						"allow_cert_cert_validity": schema.Int64Attribute{
							Description: "Certificate validity (hours).",
							Computed:    true,
						},
						"allow_local_files": schema.BoolAttribute{
							Description: "Process files from servers - Allow scan on server flag.",
							Computed:    true,
						},
						"allow_local_files_white_list": schema.BoolAttribute{
							Description: "Process files from servers flag (false = ALLOW ALL EXCEPT, true = DENY ALL EXCEPT).",
							Computed:    true,
						},
						"allow_local_files_local_paths": schema.ListAttribute{
							ElementType: types.StringType,
							Description: "Paths.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Workflow description.",
							Computed:    true,
						},
						"id": schema.Int64Attribute{
							Description: "Workflow id.",
							Computed:    true,
						},
						"include_webhook_signature": schema.BoolAttribute{
							Description: "Webhook - Include webhook signature flag.",
							Computed:    true,
						},
						"include_webhook_signature_certificate_id": schema.Int64Attribute{
							Description: "Webhook - Certificate id.",
							Computed:    true,
						},
						"last_modified": schema.Int64Attribute{
							Description: "Last modified timestamp (unix epoch).",
							Computed:    true,
						},
						"mutable": schema.BoolAttribute{
							Description: "Mutable flag.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Workflow name.",
							Computed:    true,
						},
						"workflow_id": schema.Int64Attribute{
							Description: "Workflow id.",
							Computed:    true,
						},
						"zone_id": schema.Int64Attribute{
							Description: "Workflow network zone id.",
							Computed:    true,
						},
						"scan_allowed": schema.ListAttribute{
							ElementType: types.Int64Type,
							Description: "Restrictions - Restrict access to following roles",
							Computed:    true,
						},
						"result_allowed": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"role": schema.Int64Attribute{
										Computed: true,
									},
									"visibility": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
						"option_values": schema.ObjectAttribute{
							AttributeTypes: map[string]attr.Type{
								"archive_handling_max_number_files":           types.Int64Type,
								"archive_handling_max_recursion_level":        types.Int64Type,
								"archive_handling_max_size_files":             types.Int64Type,
								"archive_handling_timeout":                    types.Int64Type,
								"filetype_analysis_timeout":                   types.Int64Type,
								"process_info_global_timeout":                 types.BoolType,
								"process_info_global_timeout_value":           types.Int64Type,
								"process_info_max_download_size":              types.Int64Type,
								"process_info_max_file_size":                  types.Int64Type,
								"process_info_quarantine":                     types.BoolType,
								"process_info_skip_hash":                      types.BoolType,
								"process_info_skip_processing_fast_symlink":   types.BoolType,
								"process_info_workflow_priority":              types.Int64Type,
								"scan_filescan_check_av_engine":               types.BoolType,
								"scan_filescan_download_timeout":              types.Int64Type,
								"scan_filescan_global_scan_timeout":           types.Int64Type,
								"scan_filescan_per_engine_scan_timeout":       types.Int64Type,
								"vul_filescan_timeout_vulnerability_scanning": types.Int64Type,
							},
							Description: "Options",
							Computed:    true,
						},
						"user_agents": schema.ListAttribute{
							ElementType: types.StringType,
							Description: "Restrictions - Limit to specified user agents.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (d *workflows) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*opswatClient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *opswatClient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

type workflowsModel struct {
	Workflows []workflowModel `tfsdk:"workflows"`
}

type workflowModel struct {
	AllowCert                            types.Bool           `tfsdk:"allow_cert"`
	AllowCertCert                        types.String         `tfsdk:"allow_cert_cert"`
	AllowCertCertValidity                types.Int64          `tfsdk:"allow_cert_cert_validity"`
	AllowLocalFiles                      types.Bool           `tfsdk:"allow_local_files"`
	AllowLocalFilesWhiteList             types.Bool           `tfsdk:"allow_local_files_white_list"`
	AllowLocalFilesLocalPaths            []string             `tfsdk:"allow_local_files_local_paths"`
	Description                          types.String         `tfsdk:"description"`
	IncludeWebhookSignature              types.Bool           `tfsdk:"include_webhook_signature"`
	IncludeWebhookSignatureCertificateID types.Int64          `tfsdk:"include_webhook_signature_certificate_id"`
	Mutable                              types.Bool           `tfsdk:"mutable"`
	Name                                 types.String         `tfsdk:"name"`
	WorkflowID                           types.Int64          `tfsdk:"workflow_id"`
	ZoneID                               types.Int64          `tfsdk:"zone_id"`
	ScanAllowed                          []int                `tfsdk:"scan_allowed"`
	ResultAllowed                        []ResultAllowedModel `tfsdk:"result_allowed"`
	OptionValues                         OptionValuesModel    `tfsdk:"option_values"`
	UserAgents                           []string             `tfsdk:"user_agents"`
	ID                                   types.Int64          `tfsdk:"id"`
	LastModified                         types.Int64          `tfsdk:"last_modified"`
}

// ResultAllowModel test
type ResultAllowedModel struct {
	Role       types.Int64 `tfsdk:"role"`
	Visibility types.Int64 `tfsdk:"visibility"`
}

// OptionValues
type OptionValuesModel struct {
	ArchiveHandlingMaxNumberFiles           types.Int64 `tfsdk:"archive_handling_max_number_files"`
	ArchiveHandlingMaxRecursionLevel        types.Int64 `tfsdk:"archive_handling_max_recursion_level"`
	ArchiveHandlingMaxSizeFiles             types.Int64 `tfsdk:"archive_handling_max_size_files"`
	ArchiveHandlingTimeout                  types.Int64 `tfsdk:"archive_handling_timeout"`
	FiletypeAnalysisTimeout                 types.Int64 `tfsdk:"filetype_analysis_timeout"`
	ProcessInfoGlobalTimeout                types.Bool  `tfsdk:"process_info_global_timeout"`
	ProcessInfoGlobalTimeoutValue           types.Int64 `tfsdk:"process_info_global_timeout_value"`
	ProcessInfoMaxDownloadSize              types.Int64 `tfsdk:"process_info_max_download_size"`
	ProcessInfoMaxFileSize                  types.Int64 `tfsdk:"process_info_max_file_size"`
	ProcessInfoQuarantine                   types.Bool  `tfsdk:"process_info_quarantine"`
	ProcessInfoSkipHash                     types.Bool  `tfsdk:"process_info_skip_hash"`
	ProcessInfoSkipProcessingFastSymlink    types.Bool  `tfsdk:"process_info_skip_processing_fast_symlink"`
	ProcessInfoWorkflowPriority             types.Int64 `tfsdk:"process_info_workflow_priority"`
	ScanFilescanCheckAvEngine               types.Bool  `tfsdk:"scan_filescan_check_av_engine"`
	ScanFilescanDownloadTimeout             types.Int64 `tfsdk:"scan_filescan_download_timeout"`
	ScanFilescanGlobalScanTimeout           types.Int64 `tfsdk:"scan_filescan_global_scan_timeout"`
	ScanFilescanPerEngineScanTimeout        types.Int64 `tfsdk:"scan_filescan_per_engine_scan_timeout"`
	VulFilescanTimeoutVulnerabilityScanning types.Int64 `tfsdk:"vul_filescan_timeout_vulnerability_scanning"`
}

// Read refreshes the Terraform state with the latest data.
func (d *workflows) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowsModel

	result, err := d.client.GetWorkflows(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT workflows",
			err.Error(),
		)
		return
	}

	for _, workflow := range result {
		workflowState := workflowModel{
			AllowCert:                            types.BoolValue(workflow.AllowCert),
			AllowCertCert:                        types.StringValue(workflow.AllowCertCert),
			AllowCertCertValidity:                types.Int64Value(int64(workflow.AllowCertCertValidity)),
			AllowLocalFiles:                      types.BoolValue(workflow.AllowLocalFiles),
			AllowLocalFilesWhiteList:             types.BoolValue(workflow.AllowLocalFilesWhiteList),
			AllowLocalFilesLocalPaths:            append(workflow.AllowLocalFilesLocalPaths),
			Description:                          types.StringValue(workflow.Description),
			ID:                                   types.Int64Value(int64(workflow.Id)),
			IncludeWebhookSignature:              types.BoolValue(workflow.IncludeWebhookSignature),
			IncludeWebhookSignatureCertificateID: types.Int64Value(int64(workflow.IncludeWebhookSignatureWebhookCertificateId)),
			LastModified:                         types.Int64Value(int64(workflow.LastModified)),
			Mutable:                              types.BoolValue(workflow.Mutable),
			Name:                                 types.StringValue(workflow.Name),
			WorkflowID:                           types.Int64Value(int64(workflow.WorkflowId)),
			ZoneID:                               types.Int64Value(int64(workflow.ZoneId)),
			ScanAllowed:                          append(workflow.ScanAllowed),
			OptionValues: OptionValuesModel{
				ArchiveHandlingMaxNumberFiles:           types.Int64Value(int64(workflow.OptionValues.ArchiveHandlingMaxNumberFiles)),
				ArchiveHandlingMaxRecursionLevel:        types.Int64Value(int64(workflow.OptionValues.ArchiveHandlingMaxRecursionLevel)),
				ArchiveHandlingMaxSizeFiles:             types.Int64Value(int64(workflow.OptionValues.ArchiveHandlingMaxSizeFiles)),
				ArchiveHandlingTimeout:                  types.Int64Value(int64(workflow.OptionValues.ArchiveHandlingTimeout)),
				FiletypeAnalysisTimeout:                 types.Int64Value(int64(workflow.OptionValues.FiletypeAnalysisTimeout)),
				ProcessInfoGlobalTimeout:                types.BoolValue(workflow.OptionValues.ProcessInfoGlobalTimeout),
				ProcessInfoGlobalTimeoutValue:           types.Int64Value(int64(workflow.OptionValues.ProcessInfoGlobalTimeoutValue)),
				ProcessInfoMaxDownloadSize:              types.Int64Value(int64(workflow.OptionValues.ProcessInfoMaxDownloadSize)),
				ProcessInfoMaxFileSize:                  types.Int64Value(int64(workflow.OptionValues.ProcessInfoMaxFileSize)),
				ProcessInfoQuarantine:                   types.BoolValue(workflow.OptionValues.ProcessInfoQuarantine),
				ProcessInfoSkipHash:                     types.BoolValue(workflow.OptionValues.ProcessInfoSkipHash),
				ProcessInfoSkipProcessingFastSymlink:    types.BoolValue(workflow.OptionValues.ProcessInfoSkipProcessingFastSymlink),
				ProcessInfoWorkflowPriority:             types.Int64Value(int64(workflow.OptionValues.ProcessInfoWorkflowPriority)),
				ScanFilescanCheckAvEngine:               types.BoolValue(workflow.OptionValues.ScanFilescanCheckAvEngine),
				ScanFilescanDownloadTimeout:             types.Int64Value(int64(workflow.OptionValues.ScanFilescanDownloadTimeout)),
				ScanFilescanGlobalScanTimeout:           types.Int64Value(int64(workflow.OptionValues.ScanFilescanGlobalScanTimeout)),
				ScanFilescanPerEngineScanTimeout:        types.Int64Value(int64(workflow.OptionValues.ScanFilescanPerEngineScanTimeout)),
				VulFilescanTimeoutVulnerabilityScanning: types.Int64Value(int64(workflow.OptionValues.VulFilescanTimeoutVulnerabilityScanning)),
			},
			UserAgents: append(workflow.UserAgents),
		}

		for _, resultsallowed := range workflow.ResultAllowed {
			workflowState.ResultAllowed = append(workflowState.ResultAllowed, ResultAllowedModel{
				Role:       types.Int64Value(int64(resultsallowed.Role)),
				Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
			})
		}

		state.Workflows = append(state.Workflows, workflowState)

	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
