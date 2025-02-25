package opswatProvider

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	opswatClient "terraform-provider-opswat/internal/connectivity"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[A-Z0-9\s]+$`),
						"Must contain only UPPERCASE alphanumeric characters",
					),
				},
			},
			"user_identified_by": schema.StringAttribute{
				Description: "User name alias via claims under profile scope",
				Required:    true,
			},
			"sp": schema.ObjectAttribute{
				Required: true,
				AttributeTypes: map[string]attr.Type{
					"login_url":            types.StringType,
					"support_logout_url":   types.BoolType,
					"support_private_key":  types.BoolType,
					"support_entity_id":    types.BoolType,
					"enable_idp_initiated": types.BoolType,
					"entity_id":            types.StringType,
				},
			},
			"idp": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"authn_request_signed": schema.BoolAttribute{
						Required: true,
					},
					"entity_id": schema.StringAttribute{
						Required: true,
					},
					"valid_until": schema.StringAttribute{
						Required: true,
					},
					"x509_cert": schema.ListAttribute{
						ElementType: types.StringType,
						Required:    true,
					},
					"login_method": schema.ObjectAttribute{
						Required: true,
						AttributeTypes: map[string]attr.Type{
							"post":     types.StringType,
							"redirect": types.StringType,
						},
					},
					"logout_method": schema.ObjectAttribute{
						Required: true,
						AttributeTypes: map[string]attr.Type{
							"redirect": types.StringType,
						},
					},
				},
			},
			"version": schema.StringAttribute{
				Description: "Version number",
				Required:    true,
			},
			"role": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"option": schema.StringAttribute{
						Required: true,
					},
					"details": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Required: true,
								},
								"values": schema.ListNestedAttribute{
									Required: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"condition": schema.StringAttribute{
												Required: true,
											},
											"role_ids": schema.ListAttribute{
												ElementType: types.StringType,
												Required:    true,
											},
											"type": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
							},
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
	json := opswatClient.UserDirectory{
		Name:             plan.Name.ValueString(),
		Type:             plan.Type.ValueString(),
		Enabled:          plan.Enabled.ValueBool(),
		UserIdentifiedBy: plan.UserIdentifiedBy.ValueString(),
		Version:          plan.Version.ValueString(),
		Sp: opswatClient.Sp{
			LoginUrl:           plan.Sp.LoginUrl.ValueString(),
			SupportLogoutUrl:   plan.Sp.SupportLogoutUrl.ValueBool(),
			SupportEntityId:    plan.Sp.SupportEntityId.ValueBool(),
			SupportPrivateKey:  plan.Sp.SupportPrivateKey.ValueBool(),
			EnableIdpInitiated: plan.Sp.EnableIdpInitiated.ValueBool(),
			EntityId:           plan.Sp.EntityId.ValueString(),
		},
		Role: opswatClient.Role{
			Option:  plan.Role.Option.ValueString(),
			Details: []opswatClient.Details{},
		},
		Idp: opswatClient.Idp{
			AuthnRequestSigned: plan.Idp.AuthnRequestSigned.ValueBool(),
			EntityId:           plan.Idp.EntityId.ValueString(),
			LoginMethod: opswatClient.LoginMethod{
				Post:     plan.Idp.LoginMethod.Post.ValueString(),
				Redirect: plan.Idp.LoginMethod.Redirect.ValueString(),
			},
			LogoutMethod: opswatClient.LogoutMethod{
				Redirect: plan.Idp.LogoutMethod.Redirect.ValueString(),
			},
			ValidUntil: plan.Idp.ValidUntil.ValueString(),
		},
	}

	json.Idp.X509Cert = append(json.Idp.X509Cert, plan.Idp.X509Cert...)

	for n, details := range plan.Role.Details {
		json.Role.Details = append(json.Role.Details, opswatClient.Details{
			Key:    details.Key.ValueString(),
			Values: []opswatClient.Values{},
		})

		json.Role.Details[n].Values = append(json.Role.Details[n].Values, opswatClient.Values{
			Condition: details.Values[0].Condition,
			RoleIds:   details.Values[0].RoleIds,
			Type:      details.Values[0].Type,
		})
	}

	// Update existing user directory
	result, err := r.client.CreateDir(ctx, json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating OPSWAT user directory",
			"Could not add new user directory, unexpected error: "+err.Error(),
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

	// Get refreshed user directory config from OPSWAT
	dir, err := r.client.GetDir(ctx, int(state.ID.ValueInt64()))

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT user directories",
			"Could not read OPSWAT user directories "+err.Error(),
		)
		return
	}

	state = dirModel{
		ID:               types.Int64Value(int64(dir.ID)),
		Type:             types.StringValue(dir.Type),
		Enabled:          types.BoolValue(dir.Enabled),
		Name:             types.StringValue(dir.Name),
		UserIdentifiedBy: types.StringValue(dir.UserIdentifiedBy),
		Sp: SPModel{
			LoginUrl:           types.StringValue(dir.Sp.LoginUrl),
			SupportLogoutUrl:   types.BoolValue(dir.Sp.SupportLogoutUrl),
			SupportPrivateKey:  types.BoolValue(dir.Sp.SupportPrivateKey),
			EnableIdpInitiated: types.BoolValue(dir.Sp.EnableIdpInitiated),
			SupportEntityId:    types.BoolValue(dir.Sp.SupportEntityId),
			EntityId:           types.StringValue(dir.Sp.EntityId),
		},
		Version: types.StringValue(dir.Version),
		Idp: IDPModel{
			AuthnRequestSigned: types.BoolValue(dir.Idp.AuthnRequestSigned),
			EntityId:           types.StringValue(dir.Idp.EntityId),
			LoginMethod: LoginMethodModel{
				Post:     types.StringValue(dir.Idp.LoginMethod.Post),
				Redirect: types.StringValue(dir.Idp.LoginMethod.Redirect),
			},
			LogoutMethod: LogoutMethodModel{
				Redirect: types.StringValue(dir.Idp.LogoutMethod.Redirect),
			},
			ValidUntil: types.StringValue(dir.Idp.ValidUntil),
		},
		Role: RoleModel{
			Details: []DetailsModel{},
			Option:  types.StringValue(dir.Role.Option),
		},
	}

	state.Idp.X509Cert = append(state.Idp.X509Cert, dir.Idp.X509Cert...)

	for n, details := range dir.Role.Details {
		state.Role.Details = append(state.Role.Details, DetailsModel{
			Key:    types.StringValue(details.Key),
			Values: []ValuesModel{},
		})

		state.Role.Details[n].Values = append(state.Role.Details[n].Values, ValuesModel{
			Condition: details.Values[0].Condition,
			RoleIds:   details.Values[0].RoleIds,
			Type:      details.Values[0].Type,
		})
	}

	// Set state
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
		Name:             strings.ToUpper(plan.Name.ValueString()),
		Type:             plan.Type.ValueString(),
		Enabled:          plan.Enabled.ValueBool(),
		UserIdentifiedBy: plan.UserIdentifiedBy.ValueString(),
		Version:          plan.Version.ValueString(),
		Sp: opswatClient.Sp{
			LoginUrl:           plan.Sp.LoginUrl.ValueString(),
			SupportLogoutUrl:   plan.Sp.SupportLogoutUrl.ValueBool(),
			SupportEntityId:    plan.Sp.SupportEntityId.ValueBool(),
			SupportPrivateKey:  plan.Sp.SupportPrivateKey.ValueBool(),
			EnableIdpInitiated: plan.Sp.EnableIdpInitiated.ValueBool(),
			EntityId:           plan.Sp.EntityId.ValueString(),
		},
		Role: opswatClient.Role{
			Option:  plan.Role.Option.ValueString(),
			Details: []opswatClient.Details{},
		},
		Idp: opswatClient.Idp{
			AuthnRequestSigned: plan.Idp.AuthnRequestSigned.ValueBool(),
			EntityId:           plan.Idp.EntityId.ValueString(),
			LoginMethod: opswatClient.LoginMethod{
				Post:     plan.Idp.LoginMethod.Post.ValueString(),
				Redirect: plan.Idp.LoginMethod.Redirect.ValueString(),
			},
			LogoutMethod: opswatClient.LogoutMethod{
				Redirect: plan.Idp.LogoutMethod.Redirect.ValueString(),
			},
			ValidUntil: plan.Idp.ValidUntil.ValueString(),
		},
	}

	json.Idp.X509Cert = append(json.Idp.X509Cert, plan.Idp.X509Cert...)

	for n, details := range plan.Role.Details {
		json.Role.Details = append(json.Role.Details, opswatClient.Details{
			Key:    details.Key.ValueString(),
			Values: []opswatClient.Values{},
		})

		json.Role.Details[n].Values = append(json.Role.Details[n].Values, opswatClient.Values{
			Condition: details.Values[0].Condition,
			RoleIds:   details.Values[0].RoleIds,
			Type:      details.Values[0].Type,
		})
	}

	// Update existing dir based on ID
	_, err := r.client.UpdateDir(ctx, int(plan.ID.ValueInt64()), json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT dir",
			"Could not update dir, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	dir, err := r.client.GetDir(ctx, int(plan.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	plan = dirModel{
		ID:               types.Int64Value(int64(dir.ID)),
		Type:             types.StringValue(dir.Type),
		Enabled:          types.BoolValue(dir.Enabled),
		Name:             types.StringValue(dir.Name),
		UserIdentifiedBy: types.StringValue(dir.UserIdentifiedBy),
		Sp: SPModel{
			LoginUrl:           types.StringValue(dir.Sp.LoginUrl),
			SupportLogoutUrl:   types.BoolValue(dir.Sp.SupportLogoutUrl),
			SupportPrivateKey:  types.BoolValue(dir.Sp.SupportPrivateKey),
			EnableIdpInitiated: types.BoolValue(dir.Sp.EnableIdpInitiated),
			SupportEntityId:    types.BoolValue(dir.Sp.SupportEntityId),
			EntityId:           types.StringValue(dir.Sp.EntityId),
		},
		Version: types.StringValue(dir.Version),
		Idp: IDPModel{
			AuthnRequestSigned: types.BoolValue(dir.Idp.AuthnRequestSigned),
			EntityId:           types.StringValue(dir.Idp.EntityId),
			LoginMethod: LoginMethodModel{
				Post:     types.StringValue(dir.Idp.LoginMethod.Post),
				Redirect: types.StringValue(dir.Idp.LoginMethod.Redirect),
			},
			LogoutMethod: LogoutMethodModel{
				Redirect: types.StringValue(dir.Idp.LogoutMethod.Redirect),
			},
			ValidUntil: types.StringValue(dir.Idp.ValidUntil),
		},
		Role: RoleModel{
			Details: []DetailsModel{},
			Option:  types.StringValue(dir.Role.Option),
		},
	}

	plan.Idp.X509Cert = append(plan.Idp.X509Cert, dir.Idp.X509Cert...)

	for n, details := range dir.Role.Details {
		plan.Role.Details = append(plan.Role.Details, DetailsModel{
			Key:    types.StringValue(details.Key),
			Values: []ValuesModel{},
		})

		plan.Role.Details[n].Values = append(plan.Role.Details[n].Values, ValuesModel{
			Condition: details.Values[0].Condition,
			RoleIds:   details.Values[0].RoleIds,
			Type:      details.Values[0].Type,
		})
	}

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
	err := r.client.DeleteDir(ctx, int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Delete OPSWAT user directory",
			"Could not delete user directory, unexpected error: "+err.Error(),
		)
		return
	}
}
