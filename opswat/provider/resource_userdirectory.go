package opswatProvider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/opswat/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &Dir{}
	_ resource.ResourceWithConfigure = &Dir{}
)

// NewDir is a helper function to simplify the provider implementation.
func NewDir() resource.Resource {
	return &Dir{}
}

// Dir is the resource implementation.
type Dir struct {
	client *opswatClient.Client
}

type dirModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Type             types.String `tfsdk:"type"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	Name             types.String `tfsdk:"name"`
	UserIdentifiedBy types.String `tfsdk:"user_identified_by"`
	Sp               struct {
		LoginUrl           types.String `tfsdk:"login_url"`
		SupportLogoutUrl   types.Bool   `tfsdk:"support_logout_url"`
		SupportPrivateKey  types.Bool   `tfsdk:"support_private_key"`
		SupportEntityId    types.Bool   `tfsdk:"support_entity_id"`
		EnableIdpInitiated types.Bool   `tfsdk:"enable_idp_initiated"`
		EntityId           types.String `tfsdk:"entity_id"`
	} `tfsdk:"sp"`
	Role struct {
		Option  types.String `tfsdk:"option"`
		Details struct {
			Default types.Int64 `tfsdk:"default"`
		} `tfsdk:"details"`
	} `json:"role"`
	Version types.String `tfsdk:"version"`
	Idp     struct {
		AuthnRequestSigned types.Bool   `tfsdk:"authn_request_signed"`
		EntityId           types.String `tfsdk:"entity_id"`
		LoginMethod        struct {
			Post     types.String `tfsdk:"post"`
			Redirect types.String `tfsdk:"redirect"`
		} `tfsdk:"login_method"`
		LogoutMethod struct {
			Redirect types.String `tfsdk:"redirect"`
		} `json:"logout_method"`
		ValidUntil types.String `tfsdk:"valid_until"`
		X509Cert   types.String `tfsdk:"x509_cert"`
	} `tfsdk:"idp"`
}

// Metadata returns the resource type name.
func (r *Dir) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_userdirectory"
}

// Schema defines the schema for the resource.
func (r *Dir) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Scan agent dir resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Userdirectory id.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				Description: "Local, AD, LDAP, OIDC or SAML",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Enabled flag",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Directory name",
				Required:    true,
			},
			"useridentifiedby": schema.StringAttribute{
				Description: "User name alias via claims under profile scope",
				Required:    true,
			},
			"sp": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"loginurl": schema.StringAttribute{
							Required: true,
						},
						"supportlogouturl": schema.BoolAttribute{
							Required: true,
						},
						"supportprivatekey": schema.BoolAttribute{
							Required: true,
						},
						"supportentityid": schema.BoolAttribute{
							Required: true,
						},
						"enableidpinitiated": schema.BoolAttribute{
							Required: true,
						},
						"entityid": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"role": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"option": schema.StringAttribute{
							Required: true,
						},
						"details": schema.ObjectAttribute{
							Required: true,
							AttributeTypes: map[string]attr.Type{
								"default": types.Int64Type,
							},
						},
					},
				},
			},
			"version": schema.StringAttribute{
				Description: "Version number",
				Required:    true,
			},
			"idp": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"authnrequestsigned": schema.BoolAttribute{
							Required: true,
						},
						"entityid": schema.StringAttribute{
							Required: true,
						},
						"loginmethod": schema.ObjectAttribute{
							Required: true,
							AttributeTypes: map[string]attr.Type{
								"post":     types.StringType,
								"redirect": types.StringType,
							},
						},
						"validuntil": schema.StringAttribute{
							Required: true,
						},
						"x509cert": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *Dir) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *Dir) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan dirModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.UserDirectory{}

	// Update existing order
	result, err := r.client.CreateDir(json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating OPSWAT dir",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Populate computed values
	plan.ID = types.Int64Value(int64(result.ID))

	resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *Dir) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//// Get current state
	//var state dirModel
	//diags := req.State.Get(ctx, &state)
	//resp.Diagnostics.Append(diags...)
	//if resp.Diagnostics.HasError() {
	//	return
	//}
	//
	//// Get refreshed order value from OPSWAT
	//dir, err := r.client.GetDir(int(state.ID.ValueInt64()))
	//if err != nil {
	//	resp.Diagnostics.AddError(
	//		"Error Reading OPSWAT dir",
	//		"Could not read OPSWAT Scan agent dir count "+err.Error(),
	//	)
	//	return
	//}
	//
	//state = dirModel{
	//
	//	}}
	//
	//// Set refreshed state
	////diags = resp.State.Set(ctx, &state)
	////resp.Diagnostics.Append(diags...)
	////if resp.Diagnostics.HasError() {
	////	return
	////}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *Dir) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//	// Retrieve values from plan
	//	var plan dirModel
	//	diags := req.Plan.Get(ctx, &plan)
	//
	//	resp.Diagnostics.Append(diags...)
	//	if resp.Diagnostics.HasError() {
	//		return
	//	}
	//
	//	// Generate API request body from plan
	//	json := opswatClient.Dir{
	//		AllowCert:                 plan.AllowCert.ValueBool(),
	//		AllowCertCert:             plan.AllowCertCert.ValueString(),
	//		AllowCertCertValidity:     int(plan.AllowCertCertValidity.ValueInt64()),
	//		AllowLocalFiles:           plan.AllowLocalFiles.ValueBool(),
	//		AllowLocalFilesWhiteList:  plan.AllowLocalFilesWhiteList.ValueBool(),
	//		AllowLocalFilesLocalPaths: plan.AllowLocalFilesLocalPaths,
	//		Description:               plan.Description.ValueString(),
	//		//Id:                        int(plan.ID.ValueInt64()),
	//		IncludeWebhookSignature:                     plan.IncludeWebhookSignature.ValueBool(),
	//		IncludeWebhookSignatureWebhookCertificateId: int(plan.IncludeWebhookSignatureCertificateID.ValueInt64()),
	//		//LastModified:  int(plan.LastModified.ValueInt64()),
	//		//Mutable:       plan.Mutable.ValueBool(),
	//		Name:          plan.Name.ValueString(),
	//		DirId:         int(plan.DirID.ValueInt64()),
	//		ZoneId:        int(plan.ZoneID.ValueInt64()),
	//		ScanAllowed:   plan.ScanAllowed,
	//		ResultAllowed: []opswatClient.ResultAllowed{},
	//		UserAgents:    plan.UserAgents,
	//		OptionValues: opswatClient.OptionValues{
	//			ArchiveHandlingMaxNumberFiles:           int(plan.OptionValues.ArchiveHandlingMaxNumberFiles.ValueInt64()),
	//			ArchiveHandlingMaxRecursionLevel:        int(plan.OptionValues.ArchiveHandlingMaxRecursionLevel.ValueInt64()),
	//			ArchiveHandlingMaxSizeFiles:             int(plan.OptionValues.ArchiveHandlingMaxSizeFiles.ValueInt64()),
	//			ArchiveHandlingTimeout:                  int(plan.OptionValues.ArchiveHandlingTimeout.ValueInt64()),
	//			FiletypeAnalysisTimeout:                 int(plan.OptionValues.FiletypeAnalysisTimeout.ValueInt64()),
	//			ProcessInfoGlobalTimeout:                plan.OptionValues.ProcessInfoGlobalTimeout.ValueBool(),
	//			ProcessInfoGlobalTimeoutValue:           int(plan.OptionValues.ProcessInfoGlobalTimeoutValue.ValueInt64()),
	//			ProcessInfoMaxDownloadSize:              int(plan.OptionValues.ProcessInfoMaxDownloadSize.ValueInt64()),
	//			ProcessInfoMaxFileSize:                  int(plan.OptionValues.ProcessInfoMaxFileSize.ValueInt64()),
	//			ProcessInfoQuarantine:                   plan.OptionValues.ProcessInfoQuarantine.ValueBool(),
	//			ProcessInfoSkipHash:                     plan.OptionValues.ProcessInfoSkipHash.ValueBool(),
	//			ProcessInfoSkipProcessingFastSymlink:    plan.OptionValues.ProcessInfoSkipProcessingFastSymlink.ValueBool(),
	//			ProcessInfoDirPriority:                  int(plan.OptionValues.ProcessInfoDirPriority.ValueInt64()),
	//			ScanFilescanCheckAvEngine:               plan.OptionValues.ScanFilescanCheckAvEngine.ValueBool(),
	//			ScanFilescanDownloadTimeout:             int(plan.OptionValues.ScanFilescanDownloadTimeout.ValueInt64()),
	//			ScanFilescanGlobalScanTimeout:           int(plan.OptionValues.ScanFilescanGlobalScanTimeout.ValueInt64()),
	//			ScanFilescanPerEngineScanTimeout:        int(plan.OptionValues.ScanFilescanPerEngineScanTimeout.ValueInt64()),
	//			VulFilescanTimeoutVulnerabilityScanning: int(plan.OptionValues.VulFilescanTimeoutVulnerabilityScanning.ValueInt64()),
	//		},
	//	}
	//
	//	for _, resultsallowed := range plan.ResultAllowed {
	//		json.ResultAllowed = append(json.ResultAllowed, opswatClient.ResultAllowed{
	//			Role:       int(resultsallowed.Role.ValueInt64()),
	//			Visibility: int(resultsallowed.Visibility.ValueInt64()),
	//		})
	//	}
	//
	//	// Update existing dir based on ID
	//	_, err := r.client.UpdateDir(int(plan.ID.ValueInt64()), json)
	//	if err != nil {
	//		resp.Diagnostics.AddError(
	//			"Error Updating OPSWAT dir",
	//			"Could not update dir, unexpected error: "+err.Error(),
	//		)
	//		return
	//	}
	//
	//	// Fetch updated items
	//	result, err := r.client.GetDir(int(plan.ID.ValueInt64()))
	//	if err != nil {
	//		resp.Diagnostics.AddError(
	//			"Error Reading OPSWAT dir",
	//			"Could not read OPSWAT dir "+err.Error(),
	//		)
	//		return
	//	}
	//
	//	plan = dirModel{
	//		AllowCert:                            types.BoolValue(result.AllowCert),
	//		AllowCertCert:                        types.StringValue(result.AllowCertCert),
	//		AllowCertCertValidity:                types.Int64Value(int64(result.AllowCertCertValidity)),
	//		AllowLocalFiles:                      types.BoolValue(result.AllowLocalFiles),
	//		AllowLocalFilesWhiteList:             types.BoolValue(result.AllowLocalFilesWhiteList),
	//		AllowLocalFilesLocalPaths:            append(result.AllowLocalFilesLocalPaths),
	//		Description:                          types.StringValue(result.Description),
	//		ID:                                   types.Int64Value(int64(result.Id)),
	//		IncludeWebhookSignature:              types.BoolValue(result.IncludeWebhookSignature),
	//		IncludeWebhookSignatureCertificateID: types.Int64Value(int64(result.IncludeWebhookSignatureWebhookCertificateId)),
	//		LastModified:                         types.Int64Value(int64(result.LastModified)),
	//		Mutable:                              types.BoolValue(result.Mutable),
	//		Name:                                 types.StringValue(result.Name),
	//		DirID:                                types.Int64Value(int64(result.DirId)),
	//		ZoneID:                               types.Int64Value(int64(result.ZoneId)),
	//		ScanAllowed:                          append(result.ScanAllowed),
	//		UserAgents:                           append(result.UserAgents),
	//		//PrefHashes:                           PrefHashesModel{DSAdvancedSettingHash: types.StringValue(dir.PrefHashes.DSADVANCEDSETTINGHASH)},
	//		OptionValues: OptionValuesModel{
	//			ArchiveHandlingMaxNumberFiles:           types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxNumberFiles)),
	//			ArchiveHandlingMaxRecursionLevel:        types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxRecursionLevel)),
	//			ArchiveHandlingMaxSizeFiles:             types.Int64Value(int64(result.OptionValues.ArchiveHandlingMaxSizeFiles)),
	//			ArchiveHandlingTimeout:                  types.Int64Value(int64(result.OptionValues.ArchiveHandlingTimeout)),
	//			FiletypeAnalysisTimeout:                 types.Int64Value(int64(result.OptionValues.FiletypeAnalysisTimeout)),
	//			ProcessInfoGlobalTimeout:                types.BoolValue(result.OptionValues.ProcessInfoGlobalTimeout),
	//			ProcessInfoGlobalTimeoutValue:           types.Int64Value(int64(result.OptionValues.ProcessInfoGlobalTimeoutValue)),
	//			ProcessInfoMaxDownloadSize:              types.Int64Value(int64(result.OptionValues.ProcessInfoMaxDownloadSize)),
	//			ProcessInfoMaxFileSize:                  types.Int64Value(int64(result.OptionValues.ProcessInfoMaxFileSize)),
	//			ProcessInfoQuarantine:                   types.BoolValue(result.OptionValues.ProcessInfoQuarantine),
	//			ProcessInfoSkipHash:                     types.BoolValue(result.OptionValues.ProcessInfoSkipHash),
	//			ProcessInfoSkipProcessingFastSymlink:    types.BoolValue(result.OptionValues.ProcessInfoSkipProcessingFastSymlink),
	//			ProcessInfoDirPriority:                  types.Int64Value(int64(result.OptionValues.ProcessInfoDirPriority)),
	//			ScanFilescanCheckAvEngine:               types.BoolValue(result.OptionValues.ScanFilescanCheckAvEngine),
	//			ScanFilescanDownloadTimeout:             types.Int64Value(int64(result.OptionValues.ScanFilescanDownloadTimeout)),
	//			ScanFilescanGlobalScanTimeout:           types.Int64Value(int64(result.OptionValues.ScanFilescanGlobalScanTimeout)),
	//			ScanFilescanPerEngineScanTimeout:        types.Int64Value(int64(result.OptionValues.ScanFilescanPerEngineScanTimeout)),
	//			VulFilescanTimeoutVulnerabilityScanning: types.Int64Value(int64(result.OptionValues.VulFilescanTimeoutVulnerabilityScanning)),
	//		}}
	//
	//	//fmt.Println("PARSED WORKFLOWS")
	//	//spew.Dump(dirState)
	//
	//	for _, resultsallowed := range result.ResultAllowed {
	//		plan.ResultAllowed = append(plan.ResultAllowed, ResultAllowedModel{
	//			Role:       types.Int64Value(int64(resultsallowed.Role)),
	//			Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
	//		})
	//	}
	//
	//	diags = resp.State.Set(ctx, plan)
	//	resp.Diagnostics.Append(diags...)
	//	if resp.Diagnostics.HasError() {
	//		return
	//	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *Dir) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//// Retrieve values from plan
	//var state dirModel
	//diags := req.State.Get(ctx, &state)
	//resp.Diagnostics.Append(diags...)
	//if resp.Diagnostics.HasError() {
	//	return
	//}
	//
	//// Update existing dir based on ID
	//err := r.client.DeleteDir(int(state.ID.ValueInt64()))
	//if err != nil {
	//	resp.Diagnostics.AddError(
	//		"Error Delete OPSWAT dir",
	//		"Could not update dir, unexpected error: "+err.Error(),
	//	)
	//	return
	//}
}
