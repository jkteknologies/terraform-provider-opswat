package opswatProvider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/opswat/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &configQuarantine{}
	_ resource.ResourceWithConfigure = &configQuarantine{}
)

// NewConfigQuarantine is a helper function to simplify the provider implementation.
func NewConfigQuarantine() resource.Resource {
	return &configQuarantine{}
}

// configQuarantine is the resource implementation.
type configQuarantine struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *configQuarantine) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_quarantine"
}

// Schema defines the schema for the resource.
func (r *configQuarantine) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global file sync can timeout resource.",
		Attributes: map[string]schema.Attribute{
			"cleanuprange": schema.Int64Attribute{
				Description: "Setting quarantine clean up time (clean up records older than). Note:The clean up range is defined in hours.",
				Required:    true,
			},
		},
	}
}

// timeouts maps timeout schema data.
type configQuarantineModel struct {
	Cleanuprange types.Int64 `tfsdk:"cleanuprange"`
}

func (r *configQuarantine) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *configQuarantine) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan configQuarantineModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.ConfigQuarantine{
		Cleanuprange: int(plan.Cleanuprange.ValueInt64()),
	}

	// Update existing order
	_, err := r.client.CreateConfigQuarantine(json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT session config",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not populated.
	result, err := r.client.GetConfigQuarantine()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	plan.Cleanuprange = types.Int64Value(int64(result.Cleanuprange))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *configQuarantine) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state configQuarantineModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from OPSWAT
	result, err := r.client.GetConfigQuarantine()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	state.Cleanuprange = types.Int64Value(int64(result.Cleanuprange))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *configQuarantine) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan configQuarantineModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.ConfigQuarantine{
		Cleanuprange: int(plan.Cleanuprange.ValueInt64()),
	}

	// Update existing order
	_, err := r.client.UpdateConfigQuarantine(json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT session config",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not populated.
	result, err := r.client.GetConfigQuarantine()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	plan.Cleanuprange = types.Int64Value(int64(result.Cleanuprange))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *configQuarantine) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
