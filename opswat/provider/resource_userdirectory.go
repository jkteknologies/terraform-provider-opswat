package opswatProvider

import (
	"context"
	"fmt"
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
				Optional:    true,
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
			/*"role": schema.MapNestedAttribute{
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
			},*/
			"version": schema.StringAttribute{
				Description: "Version number",
				Required:    true,
			},
			/*"idp": schema.MapNestedAttribute{
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
			},-*/
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
			"Error Creating OPSWAT user directory",
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
	// Get current state
	var state dirModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from OPSWAT
	dir, err := r.client.GetDir(int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT user directory",
			"Could not read OPSWAT user directory "+err.Error(),
		)
		return
	}

	state = dirModel{
		ID:               types.Int64Value(int64(dir.ID)),
		Type:             types.StringValue(dir.Type),
		Enabled:          types.BoolValue(dir.Enabled),
		Name:             types.StringValue(dir.Name),
		UserIdentifiedBy: types.StringValue(dir.UserIdentifiedBy),
		Sp:               SPModel{},
		//Role:             RoleModel{},
		Version: types.StringValue(dir.Version),
		//Idp:              IDPModel{},
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *Dir) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan dirModel
	diags := req.Plan.Get(ctx, &plan)

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.UserDirectory{
		ID:               int(plan.ID.ValueInt64()),
		Type:             plan.Type.ValueString(),
		Enabled:          plan.Enabled.ValueBool(),
		Name:             plan.Name.ValueString(),
		UserIdentifiedBy: plan.UserIdentifiedBy.ValueString(),
		Sp:               opswatClient.SpModel{},
		Role:             opswatClient.RoleModel{},
		Version:          plan.Version.ValueString(),
		Idp:              opswatClient.IdpModel{},
	}

	// Update existing dir based on ID
	_, err := r.client.UpdateDir(int(plan.ID.ValueInt64()), json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT dir",
			"Could not update dir, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetDir(int(plan.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	plan = dirModel{
		ID:               types.Int64Value(int64(result.ID)),
		Type:             types.StringValue(result.Type),
		Enabled:          types.BoolValue(result.Enabled),
		Name:             types.StringValue(result.Name),
		UserIdentifiedBy: types.StringValue(result.UserIdentifiedBy),
		Sp: SPModel{
			LoginUrl:           types.StringValue(result.Sp.LoginUrl),
			SupportLogoutUrl:   types.BoolValue(result.Sp.SupportLogoutUrl),
			SupportPrivateKey:  types.BoolValue(result.Sp.SupportPrivateKey),
			EnableIdpInitiated: types.BoolValue(result.Sp.EnableIdpInitiated),
			SupportEntityId:    types.BoolValue(result.Sp.SupportEntityId),
			EntityId:           types.StringValue(result.Sp.EntityId),
		},
		//Role:    RoleModel{},
		Version: types.StringValue(result.Version),
		/*Idp: IDPModel{
			AuthnRequestSigned: types.BoolValue(result.Idp.AuthnRequestSigned),
			EntityId:           types.StringValue(result.Idp.EntityId),
			LoginMethod: LoginMethodModel{
				Post:     types.StringValue(result.Idp.LoginMethod.Post),
				Redirect: types.StringValue(result.Idp.LoginMethod.Redirect),
			},
			LogoutMethod: LogoutMethodModel{
				Redirect: types.StringValue(result.Idp.LogoutMethod.Redirect),
			},
			ValidUntil: types.StringValue(result.Idp.ValidUntil),
			X509Cert:   types.StringValue(result.Idp.X509Cert),
		},*/
	}

	/*for _, details := range result.Role.Details {
		plan.Role.Details = append(plan.Role.Details, DetailsModel{
			Key:    types.StringValue(details.Key),
			Values: []ValuesModel{},
		})
	}*/

	//for _, detailsValues := range result.Role.Details {
	//	plan.Role.Details.Values = append(plan.Role.Details, DetailsModel{})
	//}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *Dir) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from plan
	var state dirModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing dir based on ID
	err := r.client.DeleteDir(int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Delete OPSWAT user directory",
			"Could not update user directory, unexpected error: "+err.Error(),
		)
		return
	}
}
