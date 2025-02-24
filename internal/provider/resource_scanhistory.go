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
	_ resource.Resource              = &scanHistory{}
	_ resource.ResourceWithConfigure = &scanHistory{}
)

// NewScanHistory is a helper function to simplify the provider implementation.
func NewScanHistory() resource.Resource {
	return &scanHistory{}
}

// scanHistory is the resource implementation.
type scanHistory struct {
	client *opswatClient.Client
}

// timeouts maps timeout schema data.
type cleanuprangeModel struct {
	Cleanuprange types.Int64 `tfsdk:"cleanuprange"`
}

// Metadata returns the resource type name.
func (r *scanHistory) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scan_history"
}

// Schema defines the schema for the resource.
func (r *scanHistory) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Processing history clean up time resource.",
		Attributes: map[string]schema.Attribute{
			"cleanuprange": schema.Int64Attribute{
				Description: "Setting processing history clean up time (clean up records older than).\n\nNote:The clean up range is defined in hours.",
				Required:    true,
			},
		},
	}
}

func (r *scanHistory) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *scanHistory) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan cleanuprangeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	cleanuprange := plan.Cleanuprange.ValueInt64()

	// Update existing item
	_, err := r.client.CreateScanHistory(ctx, int(cleanuprange))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT processing history clean up time",
			"Could not update OPSWAT processing history clean up time, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetScanHistory(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT processing history clean up time",
			"Could not read OPSWAT processing history clean up time, unexpected error: "+err.Error(),
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
func (r *scanHistory) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state cleanuprangeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed item
	result, err := r.client.GetScanHistory(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT processing history clean up time",
			"Could not read OPSWAT processing history clean up time, unexpected error: "+err.Error(),
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
func (r *scanHistory) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan cleanuprangeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	cleanuprange := plan.Cleanuprange.ValueInt64()

	// Update existing item
	_, err := r.client.UpdateScanHistory(ctx, int(cleanuprange))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT processing history clean up time",
			"Could not update OPSWAT processing history clean up time, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetScanHistory(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT processing history clean up time",
			"Could not read OPSWAT processing history clean up time, unexpected error: "+err.Error(),
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
func (r *scanHistory) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
