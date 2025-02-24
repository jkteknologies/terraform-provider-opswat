package opswatProvider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/internal/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &globalSync{}
	_ resource.ResourceWithConfigure = &globalSync{}
)

// NewGlobalSync is a helper function to simplify the provider implementation.
func NewGlobalSync() resource.Resource {
	return &globalSync{}
}

// globalSync is the resource implementation.
type globalSync struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *globalSync) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file_sync"
}

// Schema defines the schema for the resource.
func (r *globalSync) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global file sync timeout resource.",
		Attributes: map[string]schema.Attribute{
			"timeout": schema.Int64Attribute{
				Description: "Global file sync can timeout.",
				Required:    true,
			},
		},
	}
}

func (r *globalSync) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *globalSync) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan timeoutModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	timeout := plan.Timeout.ValueInt64()

	// Update existing item
	_, err := r.client.CreateGlobalSync(ctx, int(timeout))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT Global sync value",
			"Could not update global sync, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetGlobalSync(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Global sync timeout",
			"Could not read OPSWAT Global sync timeout "+err.Error(),
		)
		return
	}

	plan.Timeout = types.Int64Value(int64(result.Timeout))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *globalSync) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state timeoutModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed item
	result, err := r.client.GetGlobalSync(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Global sync timeout",
			"Could not read OPSWAT Global sync timeout "+err.Error(),
		)
		return
	}

	state.Timeout = types.Int64Value(int64(result.Timeout))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *globalSync) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan timeoutModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	timeout := plan.Timeout.ValueInt64()

	// Update existing item
	_, err := r.client.UpdateGlobalSync(ctx, int(timeout))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT Global sync timeout",
			"Could not update lobal sync timeout value, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetGlobalSync(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Global sync timeout",
			"Could not read OPSWAT Global sync timeout "+err.Error(),
		)
		return
	}

	plan.Timeout = types.Int64Value(int64(result.Timeout))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *globalSync) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
