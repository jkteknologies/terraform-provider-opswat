package opswatProvider

import (
	"context"
	"fmt"
	opswatClient "terraform-provider-opswat/internal/connectivity"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &UserRole{}
	_ resource.ResourceWithConfigure = &UserRole{}
)

// NewUser is a helper function to simplify the provider implementation.
func NewUserRole() resource.Resource {
	return &UserRole{}
}

// User is the resource implementation.
type UserRole struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *UserRole) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_role"
}

// Schema defines the schema for the resource.
func (r *UserRole) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "User role resource.",
		Attributes: map[string]schema.Attribute{
			"display_name": schema.StringAttribute{
				Description: "User role display name.",
				Optional:    true,
			},
			"id": schema.Int64Attribute{
				Description: "User id.",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "User role name.",
				Optional:    true,
			},
			"rights": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					//"scanlog": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"statistics": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"quarantine": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"updatelog": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"configlog": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"rule": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"workflow": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"zone": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"agents": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"engines": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"external": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"skip": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"cert": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"webhookauth": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"retention": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"users": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"license": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"update": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"scan": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					//"healthcheck": schema.ListAttribute{
					//	ElementType: types.StringType,
					//	Optional:    true,
					//},
					"download": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"fetch": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
				},
			},
		},
	}
}

// Timeouts maps timeout schema data.
type userRoleModel struct {
	DisplayName types.String    `tfsdk:"display_name"`
	ID          types.Int64     `tfsdk:"id"`
	Name        types.String    `tfsdk:"name"`
	UserRights  UserRightsModel `tfsdk:"rights"`
}

type UserRightsModel struct {
	//Scanlog     []string `tfsdk:"scanlog"`
	//Statistics  []string `tfsdk:"statistics"`
	//Quarantine  []string `tfsdk:"quarantine"`
	//Updatelog   []string `tfsdk:"updatelog"`
	//Configlog   []string `tfsdk:"configlog"`
	//Rule        []string `tfsdk:"rule"`
	//Workflow    []string `tfsdk:"workflow"`
	//Zone        []string `tfsdk:"zone"`
	//Agents      []string `tfsdk:"agents"`
	//Engines     []string `tfsdk:"engines"`
	//External    []string `tfsdk:"external"`
	//Skip        []string `tfsdk:"skip"`
	//Cert        []string `tfsdk:"cert"`
	//WebhookAuth []string `tfsdk:"webhook_auth"`
	//Retention   []string `tfsdk:"retention"`
	//Users       []string `tfsdk:"users"`
	//License     []string `tfsdk:"license"`
	//Update      []string `tfsdk:"update"`
	//Scan        []string `tfsdk:"scan"`
	//Healthcheck []string `tfsdk:"healthcheck"`
	Fetch    []string `tfsdk:"fetch"`
	Download []string `tfsdk:"download"`
}

func (r *UserRole) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *UserRole) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan userRoleModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.UserRole{
		Name:        plan.Name.ValueString(),
		DisplayName: plan.DisplayName.ValueString(),
		ID:          int(plan.ID.ValueInt64()),
		UserRights: opswatClient.UserRights{
			//Scanlog:     plan.UserRights.Scanlog,
			//Statistics:  plan.UserRights.Statistics,
			//Quarantine:  plan.UserRights.Quarantine,
			//Updatelog:   plan.UserRights.Updatelog,
			//Configlog:   plan.UserRights.Configlog,
			//Rule:        plan.UserRights.Rule,
			//Workflow:    plan.UserRights.Workflow,
			//Zone:        plan.UserRights.Zone,
			//Agents:      plan.UserRights.Agents,
			//Engines:     plan.UserRights.Engines,
			//External:    plan.UserRights.External,
			//Skip:        plan.UserRights.Skip,
			//Cert:        plan.UserRights.Cert,
			//WebhookAuth: plan.UserRights.WebhookAuth,
			//Retention:   plan.UserRights.Retention,
			//Users:       plan.UserRights.Users,
			//License:     plan.UserRights.License,
			//Update:      plan.UserRights.Update,
			//Scan:        plan.UserRights.Scan,
			//Healthcheck: plan.UserRights.Healthcheck,
			Fetch:    plan.UserRights.Fetch,
			Download: plan.UserRights.Download,
		},
	}

	// Update existing user
	result, err := r.client.CreateUserRole(ctx, json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating OPSWAT user role",
			"Could not add new user role, unexpected error: "+err.Error(),
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
func (r *UserRole) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state userRoleModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed workflow config from OPSWAT
	userRole, err := r.client.GetUserRole(ctx, int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT user role",
			"Could not read OPSWAT user role "+err.Error(),
		)
		return
	}

	state = userRoleModel{
		DisplayName: types.StringValue(userRole.DisplayName),
		ID:          types.Int64Value(int64(userRole.ID)),
		Name:        types.StringValue(userRole.Name),
		UserRights: UserRightsModel{
			//Scanlog:     userRole.UserRights.Scanlog,
			//Statistics:  userRole.UserRights.Statistics,
			//Quarantine:  userRole.UserRights.Quarantine,
			//Updatelog:   userRole.UserRights.Updatelog,
			//Configlog:   userRole.UserRights.Configlog,
			//Rule:        userRole.UserRights.Rule,
			//Workflow:    userRole.UserRights.Workflow,
			//Zone:        userRole.UserRights.Zone,
			//Agents:      userRole.UserRights.Agents,
			//Engines:     userRole.UserRights.Engines,
			//External:    userRole.UserRights.External,
			//Skip:        userRole.UserRights.Skip,
			//Cert:        userRole.UserRights.Cert,
			//WebhookAuth: userRole.UserRights.WebhookAuth,
			//Retention:   userRole.UserRights.Retention,
			//Users:       userRole.UserRights.Users,
			//License:     userRole.UserRights.License,
			//Update:      userRole.UserRights.Update,
			//Scan:        userRole.UserRights.Scan,
			//Healthcheck: userRole.UserRights.Healthcheck,
			Fetch:    userRole.UserRights.Fetch,
			Download: userRole.UserRights.Download,
		},
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update userRole resource and sets the updated Terraform state on success.
func (r *UserRole) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan userRoleModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state userRoleModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	displayname := ""
	name := ""

	if !plan.DisplayName.Equal(state.DisplayName) {
		displayname = plan.DisplayName.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		name = plan.Name.ValueString()
	}

	// Generate API request body from plan
	json := opswatClient.UserRole{
		DisplayName: displayname,
		Name:        name,
		UserRights: opswatClient.UserRights{
			Download: state.UserRights.Download,
			Fetch:    state.UserRights.Fetch,
		},
	}

	// Update existing user
	_, err := r.client.UpdateUserRole(ctx, int(plan.ID.ValueInt64()), json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT user",
			"Could not update existing user, unexpected error: "+err.Error(),
		)
		return
	}

	// Get refreshed workflow config from OPSWAT
	userRole, err := r.client.GetUserRole(ctx, int(plan.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	plan = userRoleModel{
		DisplayName: types.StringValue(userRole.DisplayName),
		ID:          types.Int64Value(int64(userRole.ID)),
		Name:        types.StringValue(userRole.Name),
		UserRights:  UserRightsModel(userRole.UserRights),
	}

	// Set refreshed state
	resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *UserRole) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from plan
	var state userRoleModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update existing dir based on ID
	err := r.client.DeleteUserRole(ctx, int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Delete OPSWAT user role",
			"Could not delete user role unexpected error: "+err.Error(),
		)
		return
	}
}
