package opswatProvider

import (
	"context"
	"fmt"
	opswatClient "terraform-provider-opswat/internal/connectivity"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	planmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &Workflow{}
	_ resource.ResourceWithConfigure = &Workflow{}
)

// NewWorkflow is a helper function to simplify the provider implementation.
func NewWorkflow() resource.Resource {
	return &Workflow{}
}

// Workflow is the resource implementation.
type Workflow struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *Workflow) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow"
}

// Schema defines the schema for the resource.
func (r *Workflow) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Scan workflow resource.",
		Attributes: map[string]schema.Attribute{
			"allow_cert": schema.BoolAttribute{
				Description: "Generate batch signature with certificate - Use certificate to generate batch signature flag.",
				Required:    true,
			},
			"allow_cert_cert": schema.StringAttribute{
				Description: "Certificate used for barch signing.",
				Required:    true,
			},
			"allow_cert_cert_validity": schema.Int64Attribute{
				Description: "Certificate validity (hours).",
				Required:    true,
			},
			"allow_local_files": schema.BoolAttribute{
				Description: "Process files from servers - Allow scan on server flag.",
				Required:    true,
			},
			"allow_local_files_white_list": schema.BoolAttribute{
				Description: "Process files from servers flag (false = ALLOW ALL EXCEPT, true = DENY ALL EXCEPT).",
				Required:    true,
			},
			"allow_local_files_local_paths": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Paths.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Workflow description.",
				Required:    true,
			},
			"id": schema.Int64Attribute{
				Description: "Workflow id.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"include_webhook_signature": schema.BoolAttribute{
				Description: "Webhook - Include webhook signature flag.",
				Required:    true,
			},
			"include_webhook_signature_certificate_id": schema.Int64Attribute{
				Description: "Webhook - Certificate id.",
				Required:    true,
			},
			"last_modified": schema.Int64Attribute{
				Description: "Last modified timestamp (unix epoch).",
				Computed:    true,
			},
			"mutable": schema.BoolAttribute{
				Description: "Mutable flag.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Workflow name.",
				Required:    true,
			},
			"workflow_id": schema.Int64Attribute{
				Description: "Workflow id.",
				Required:    true,
			},
			"zone_id": schema.Int64Attribute{
				Description: "Workflow network zone id.",
				Required:    true,
			},
			"scan_allowed": schema.ListAttribute{
				ElementType: types.Int64Type,
				Description: "Restrictions - Restrict access to following roles.",
				Optional:    true,
			},
			"result_allowed": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Visibility of Processing result - Visibility of scan result.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"role": schema.Int64Attribute{
							Optional: true,
						},
						"visibility": schema.Int64Attribute{
							Optional: true,
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
				Optional:    true,
			},
			"user_agents": schema.ListAttribute{
				ElementType: types.StringType,
				Description: "Restrictions - Limit to specified user agents.",
				Required:    true,
			},
		},
	}
}

func (r *Workflow) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}

// Create creates the resource and sets the initial Terraform state.
func (r *Workflow) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan workflowModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.Workflow{
		AllowCert:                 plan.AllowCert.ValueBool(),
		AllowCertCert:             plan.AllowCertCert.ValueString(),
		AllowCertCertValidity:     int(plan.AllowCertCertValidity.ValueInt64()),
		AllowLocalFiles:           plan.AllowLocalFiles.ValueBool(),
		AllowLocalFilesWhiteList:  plan.AllowLocalFilesWhiteList.ValueBool(),
		AllowLocalFilesLocalPaths: plan.AllowLocalFilesLocalPaths,
		Description:               plan.Description.ValueString(),
		IncludeWebhookSignature:   plan.IncludeWebhookSignature.ValueBool(),
		IncludeWebhookSignatureWebhookCertificateId: int(plan.IncludeWebhookSignatureCertificateID.ValueInt64()),
		Mutable:     plan.Mutable.ValueBool(),
		Name:        plan.Name.ValueString(),
		ScanAllowed: plan.ScanAllowed,
		WorkflowId:  int(plan.WorkflowID.ValueInt64()),
		ZoneId:      int(plan.ZoneID.ValueInt64()),
		UserAgents:  plan.UserAgents,
		OptionValues: opswatClient.OptionValues{
			ArchiveHandlingMaxRecursionLevel:        int(plan.OptionValues.ArchiveHandlingMaxRecursionLevel.ValueInt64()),
			ArchiveHandlingMaxSizeFiles:             int(plan.OptionValues.ArchiveHandlingMaxSizeFiles.ValueInt64()),
			ArchiveHandlingMaxNumberFiles:           int(plan.OptionValues.ArchiveHandlingMaxNumberFiles.ValueInt64()),
			ArchiveHandlingTimeout:                  int(plan.OptionValues.ArchiveHandlingTimeout.ValueInt64()),
			FiletypeAnalysisTimeout:                 int(plan.OptionValues.FiletypeAnalysisTimeout.ValueInt64()),
			ProcessInfoGlobalTimeout:                plan.OptionValues.ProcessInfoGlobalTimeout.ValueBool(),
			ProcessInfoGlobalTimeoutValue:           int(plan.OptionValues.ProcessInfoGlobalTimeoutValue.ValueInt64()),
			ProcessInfoMaxDownloadSize:              int(plan.OptionValues.ProcessInfoMaxDownloadSize.ValueInt64()),
			ProcessInfoMaxFileSize:                  int(plan.OptionValues.ProcessInfoMaxFileSize.ValueInt64()),
			ProcessInfoQuarantine:                   plan.OptionValues.ProcessInfoQuarantine.ValueBool(),
			ProcessInfoSkipHash:                     plan.OptionValues.ProcessInfoSkipHash.ValueBool(),
			ProcessInfoSkipProcessingFastSymlink:    plan.OptionValues.ProcessInfoSkipProcessingFastSymlink.ValueBool(),
			ProcessInfoWorkflowPriority:             int(plan.OptionValues.ProcessInfoWorkflowPriority.ValueInt64()),
			ScanFilescanCheckAvEngine:               plan.OptionValues.ScanFilescanCheckAvEngine.ValueBool(),
			ScanFilescanDownloadTimeout:             int(plan.OptionValues.ScanFilescanDownloadTimeout.ValueInt64()),
			ScanFilescanGlobalScanTimeout:           int(plan.OptionValues.ScanFilescanGlobalScanTimeout.ValueInt64()),
			ScanFilescanPerEngineScanTimeout:        int(plan.OptionValues.ScanFilescanPerEngineScanTimeout.ValueInt64()),
			VulFilescanTimeoutVulnerabilityScanning: int(plan.OptionValues.VulFilescanTimeoutVulnerabilityScanning.ValueInt64()),
		},
		ResultAllowed: []opswatClient.ResultAllowed{},
		Id:            int(plan.ID.ValueInt64()),
		LastModified:  int(plan.LastModified.ValueInt64()),
	}

	for _, resultsallowed := range plan.ResultAllowed {
		json.ResultAllowed = append(json.ResultAllowed, opswatClient.ResultAllowed{
			Role:       int(resultsallowed.Role.ValueInt64()),
			Visibility: int(resultsallowed.Visibility.ValueInt64()),
		})
	}

	// Add new workflow
	workflow, err := r.client.CreateWorkflow(ctx, json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating OPSWAT workflow",
			"Could not add new workflow, unexpected error: "+err.Error(),
		)
		return
	}

	// Populate computed values
	plan.ID = types.Int64Value(int64(workflow.Id))
	plan.LastModified = types.Int64Value(int64(workflow.LastModified))

	resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *Workflow) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state workflowModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed workflow config from OPSWAT
	workflow, err := r.client.GetWorkflow(ctx, int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	state = workflowModel{
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
		UserAgents:                           append(workflow.UserAgents),
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
		}}

	for _, resultsallowed := range workflow.ResultAllowed {
		// Opswat is using '#' symbol as All roles marker
		if !unicode.IsDigit(rune(resultsallowed.Role)) {
		state.ResultAllowed = append(state.ResultAllowed, ResultAllowedModel{
				Role:       types.Int64Value(int64(resultsallowed.Role)),
				Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
			})
		} else {
			state.ResultAllowed = append(state.ResultAllowed, ResultAllowedModel{
				Role:       types.Int64Value(int64(0)),
				Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
			})
		}
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *Workflow) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan workflowModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.Workflow{
		AllowCert:                 plan.AllowCert.ValueBool(),
		AllowCertCert:             plan.AllowCertCert.ValueString(),
		AllowCertCertValidity:     int(plan.AllowCertCertValidity.ValueInt64()),
		AllowLocalFiles:           plan.AllowLocalFiles.ValueBool(),
		AllowLocalFilesWhiteList:  plan.AllowLocalFilesWhiteList.ValueBool(),
		AllowLocalFilesLocalPaths: plan.AllowLocalFilesLocalPaths,
		Description:               plan.Description.ValueString(),
		//Id:                        int(plan.ID.ValueInt64()),
		IncludeWebhookSignature:                     plan.IncludeWebhookSignature.ValueBool(),
		IncludeWebhookSignatureWebhookCertificateId: int(plan.IncludeWebhookSignatureCertificateID.ValueInt64()),
		//LastModified:  int(plan.LastModified.ValueInt64()),
		//Mutable:       plan.Mutable.ValueBool(),
		Name:          plan.Name.ValueString(),
		WorkflowId:    int(plan.WorkflowID.ValueInt64()),
		ZoneId:        int(plan.ZoneID.ValueInt64()),
		ScanAllowed:   plan.ScanAllowed,
		ResultAllowed: []opswatClient.ResultAllowed{},
		UserAgents:    plan.UserAgents,
		OptionValues: opswatClient.OptionValues{
			ArchiveHandlingMaxNumberFiles:           int(plan.OptionValues.ArchiveHandlingMaxNumberFiles.ValueInt64()),
			ArchiveHandlingMaxRecursionLevel:        int(plan.OptionValues.ArchiveHandlingMaxRecursionLevel.ValueInt64()),
			ArchiveHandlingMaxSizeFiles:             int(plan.OptionValues.ArchiveHandlingMaxSizeFiles.ValueInt64()),
			ArchiveHandlingTimeout:                  int(plan.OptionValues.ArchiveHandlingTimeout.ValueInt64()),
			FiletypeAnalysisTimeout:                 int(plan.OptionValues.FiletypeAnalysisTimeout.ValueInt64()),
			ProcessInfoGlobalTimeout:                plan.OptionValues.ProcessInfoGlobalTimeout.ValueBool(),
			ProcessInfoGlobalTimeoutValue:           int(plan.OptionValues.ProcessInfoGlobalTimeoutValue.ValueInt64()),
			ProcessInfoMaxDownloadSize:              int(plan.OptionValues.ProcessInfoMaxDownloadSize.ValueInt64()),
			ProcessInfoMaxFileSize:                  int(plan.OptionValues.ProcessInfoMaxFileSize.ValueInt64()),
			ProcessInfoQuarantine:                   plan.OptionValues.ProcessInfoQuarantine.ValueBool(),
			ProcessInfoSkipHash:                     plan.OptionValues.ProcessInfoSkipHash.ValueBool(),
			ProcessInfoSkipProcessingFastSymlink:    plan.OptionValues.ProcessInfoSkipProcessingFastSymlink.ValueBool(),
			ProcessInfoWorkflowPriority:             int(plan.OptionValues.ProcessInfoWorkflowPriority.ValueInt64()),
			ScanFilescanCheckAvEngine:               plan.OptionValues.ScanFilescanCheckAvEngine.ValueBool(),
			ScanFilescanDownloadTimeout:             int(plan.OptionValues.ScanFilescanDownloadTimeout.ValueInt64()),
			ScanFilescanGlobalScanTimeout:           int(plan.OptionValues.ScanFilescanGlobalScanTimeout.ValueInt64()),
			ScanFilescanPerEngineScanTimeout:        int(plan.OptionValues.ScanFilescanPerEngineScanTimeout.ValueInt64()),
			VulFilescanTimeoutVulnerabilityScanning: int(plan.OptionValues.VulFilescanTimeoutVulnerabilityScanning.ValueInt64()),
		},
	}

	for _, resultsallowed := range plan.ResultAllowed {
		json.ResultAllowed = append(json.ResultAllowed, opswatClient.ResultAllowed{
			Role:       int(resultsallowed.Role.ValueInt64()),
			Visibility: int(resultsallowed.Visibility.ValueInt64()),
		})
	}

	// Update existing workflow based on ID
	_, err := r.client.UpdateWorkflow(ctx, int(plan.ID.ValueInt64()), json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT workflow",
			"Could not update workflow, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetWorkflow(ctx, int(plan.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	plan = workflowModel{
		AllowCert:                            types.BoolValue(result.AllowCert),
		AllowCertCert:                        types.StringValue(result.AllowCertCert),
		AllowCertCertValidity:                types.Int64Value(int64(result.AllowCertCertValidity)),
		AllowLocalFiles:                      types.BoolValue(result.AllowLocalFiles),
		AllowLocalFilesWhiteList:             types.BoolValue(result.AllowLocalFilesWhiteList),
		AllowLocalFilesLocalPaths:            append(result.AllowLocalFilesLocalPaths),
		Description:                          types.StringValue(result.Description),
		ID:                                   types.Int64Value(int64(result.Id)),
		IncludeWebhookSignature:              types.BoolValue(result.IncludeWebhookSignature),
		IncludeWebhookSignatureCertificateID: types.Int64Value(int64(result.IncludeWebhookSignatureWebhookCertificateId)),
		LastModified:                         types.Int64Value(int64(result.LastModified)),
		Mutable:                              types.BoolValue(result.Mutable),
		Name:                                 types.StringValue(result.Name),
		WorkflowID:                           types.Int64Value(int64(result.WorkflowId)),
		ZoneID:                               types.Int64Value(int64(result.ZoneId)),
		ScanAllowed:                          append(result.ScanAllowed),
		UserAgents:                           append(result.UserAgents),
		//PrefHashes:                           PrefHashesModel{DSAdvancedSettingHash: types.StringValue(workflow.PrefHashes.DSADVANCEDSETTINGHASH)},
		OptionValues: OptionValuesModel{
			ArchiveHandlingMaxNumberFiles:           types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxNumberFiles)),
			ArchiveHandlingMaxRecursionLevel:        types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxRecursionLevel)),
			ArchiveHandlingMaxSizeFiles:             types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxSizeFiles)),
			ArchiveHandlingTimeout:                  types.Int64Value(int64(result.OptionValues.ArchiveHandlingTimeout)),
			FiletypeAnalysisTimeout:                 types.Int64Value(int64(result.OptionValues.FiletypeAnalysisTimeout)),
			ProcessInfoGlobalTimeout:                types.BoolValue(result.OptionValues.ProcessInfoGlobalTimeout),
			ProcessInfoGlobalTimeoutValue:           types.Int64Value(int64(result.OptionValues.ProcessInfoGlobalTimeoutValue)),
			ProcessInfoMaxDownloadSize:              types.Int64Value(int64(result.OptionValues.ProcessInfoMaxDownloadSize)),
			ProcessInfoMaxFileSize:                  types.Int64Value(int64(result.OptionValues.ProcessInfoMaxFileSize)),
			ProcessInfoQuarantine:                   types.BoolValue(result.OptionValues.ProcessInfoQuarantine),
			ProcessInfoSkipHash:                     types.BoolValue(result.OptionValues.ProcessInfoSkipHash),
			ProcessInfoSkipProcessingFastSymlink:    types.BoolValue(result.OptionValues.ProcessInfoSkipProcessingFastSymlink),
			ProcessInfoWorkflowPriority:             types.Int64Value(int64(result.OptionValues.ProcessInfoWorkflowPriority)),
			ScanFilescanCheckAvEngine:               types.BoolValue(result.OptionValues.ScanFilescanCheckAvEngine),
			ScanFilescanDownloadTimeout:             types.Int64Value(int64(result.OptionValues.ScanFilescanDownloadTimeout)),
			ScanFilescanGlobalScanTimeout:           types.Int64Value(int64(result.OptionValues.ScanFilescanGlobalScanTimeout)),
			ScanFilescanPerEngineScanTimeout:        types.Int64Value(int64(result.OptionValues.ScanFilescanPerEngineScanTimeout)),
			VulFilescanTimeoutVulnerabilityScanning: types.Int64Value(int64(result.OptionValues.VulFilescanTimeoutVulnerabilityScanning)),
		}}

	//fmt.Println("PARSED WORKFLOWS")
	//spew.Dump(workflowState)

	for _, resultsallowed := range result.ResultAllowed {
		plan.ResultAllowed = append(plan.ResultAllowed, ResultAllowedModel{
			Role:       types.Int64Value(int64(resultsallowed.Role)),
			Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
		})
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *Workflow) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from plan
	var state workflowModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing workflow based on ID
	err := r.client.DeleteWorkflow(ctx, int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Delete OPSWAT workflow",
			"Could not update workflow, unexpected error: "+err.Error(),
		)
		return
	}
}
