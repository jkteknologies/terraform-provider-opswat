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
	_ resource.Resource              = &Quarantine{}
	_ resource.ResourceWithConfigure = &Quarantine{}
)

// NewQuarantine is a helper function to simplify the provider implementation.
func NewQuarantine() resource.Resource {
	return &Quarantine{}
}

// Quarantine is the resource implementation.
type Quarantine struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *Quarantine) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_quarantine"
}

// Schema defines the schema for the resource.
func (r *Quarantine) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global file sync can timeout resource.",
		Attributes: map[string]schema.Attribute{
			"cleanup_range": schema.Int64Attribute{
				Description: "Setting quarantine clean up time (clean up records older than). Note:The clean up range is defined in hours.",
				Required:    true,
			},
		},
	}
}

// QuarantineModel schema
type QuarantineModel struct {
	Cleanuprange types.Int64 `tfsdk:"cleanup_range"`
}

func (r *Quarantine) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *Quarantine) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan QuarantineModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.Quarantine{
		Cleanuprange: int(plan.Cleanuprange.ValueInt64()),
	}

	// Update existing item
	_, err := r.client.CreateQuarantine(ctx, json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT quarantine config",
			"Could not add qurantine config, quarantine error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetQuarantine(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT quarantine config",
			"Could not read OPSWAT quarantine config "+err.Error(),
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
func (r *Quarantine) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state QuarantineModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed item
	result, err := r.client.GetQuarantine(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT quarantine config",
			"Could not read OPSWAT quarantine config "+err.Error(),
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
func (r *Quarantine) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan QuarantineModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.Quarantine{
		Cleanuprange: int(plan.Cleanuprange.ValueInt64()),
	}

	// Update existing item
	_, err := r.client.UpdateQuarantine(ctx, json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT quarantine config",
			"Could not update quarantine config, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetQuarantine(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT quarantine config",
			"Could not read OPSWAT quarantine config "+err.Error(),
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
func (r *Quarantine) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
