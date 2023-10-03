package opswatProvider

import (
	"context"
	"fmt"
	"github.com/emirpasic/gods/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	opswatClient "terraform-provider-opswat/internal/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &User{}
	_ resource.ResourceWithConfigure = &User{}
)

// NewUser is a helper function to simplify the provider implementation.
func NewUser() resource.Resource {
	return &User{}
}

// User is the resource implementation.
type User struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *User) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the resource.
func (r *User) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "User resource.",
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				Description: "Apikey; This function generates valid api key for metadefander core based on following rules:\n APIKEY validation criteria\n The length of the API key must be exactly 36 characters.\n It must contain numeric and lower case a, b, c, d, e and f letter characters only\n It must contain at least 10 lower case a, b, c, d, e or f letter characters.\n It must contain at least 10 numeric characters.\n It is allowed to contain at most 3 consecutive lower case letter characters (e.g. \"abcd1a2b3c...\" is invalid because of the four consecutive letters).\n It is allowed to contain at most 3 consecutive numeric characters (e.g. \"1234a1b2c3...\" is invalid because of the four consecutive numeric characters).",
				Required:    true,
			},
			"directory_id": schema.Int64Attribute{
				Description: "User dir ID to map user to.",
				Required:    true,
			},
			"display_name": schema.StringAttribute{
				Description: "User display name.",
				Required:    true,
			},
			"email": schema.StringAttribute{
				Description: "User email.",
				Required:    true,
			},
			"id": schema.Int64Attribute{
				Description: "User id.",
				Required:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "User name.",
				Required:    true,
			},
			"roles": schema.ListAttribute{
				ElementType: types.Int64Type,
				Description: "User rights.",
				Required:    true,
			},
			"ui_settings": schema.ObjectAttribute{
				Description: "User UI settings.",
				Optional:    true,
			},
		},
	}
}

// Timeouts maps timeout schema data.
type userModel struct {
	ApiKey      types.String `tfsdk:"api_key"`
	DirectoryId types.Int64  `tfsdk:"directory_id"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	ID          types.Int64  `tfsdk:"id,omitempty"`
	Name        types.String `tfsdk:"name"`
	Roles       []int        `tfsdk:"roles"`
	UiSettings  struct{}     `tfsdk:"ui_settings"`
}

func (r *User) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *User) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan userModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	json := opswatClient.User{
		ApiKey:      plan.ApiKey.ValueString(),
		DirectoryId: int(plan.DirectoryId.ValueInt64()),
		DisplayName: plan.DisplayName.ValueString(),
		Email:       plan.Email.ValueString(),
		ID:          int(plan.ID.ValueInt64()),
		Name:        plan.Name.ValueString(),
		Roles:       plan.Roles,
		UiSettings:  plan.UiSettings,
	}

	tflog.Info(ctx, utils.ToString(json))

	// Update existing user directory
	result, err := r.client.CreateUser(json)
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
func (r *User) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state userModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed workflow config from OPSWAT
	user, err := r.client.GetUser(int(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT workflow",
			"Could not read OPSWAT workflow "+err.Error(),
		)
		return
	}

	state = userModel{
		ApiKey:      types.StringValue(user.ApiKey),
		DirectoryId: types.Int64Value(int64(user.DirectoryId)),
		DisplayName: types.StringValue(user.DisplayName),
		Email:       types.StringValue(user.Email),
		ID:          types.Int64Value(int64(user.ID)),
		Name:        types.StringValue(user.Name),
		Roles:       append(user.Roles),
		UiSettings:  user.UiSettings,
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *User) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

// Delete deletes the resource and removes the Terraform state on success.
func (r *User) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
