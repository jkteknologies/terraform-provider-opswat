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
	_ resource.Resource              = &Queue{}
	_ resource.ResourceWithConfigure = &Queue{}
)

// NewQueue is a helper function to simplify the provider implementation.
func NewQueue() resource.Resource {
	return &Queue{}
}

// Queue is the resource implementation.
type Queue struct {
	client *opswatClient.Client
}

// Queue
type queueModel struct {
	MaxQueuePerAgent types.Int64 `tfsdk:"max_queue_per_agent"`
}

// Metadata returns the resource type name.
func (r *Queue) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue"
}

// Schema defines the schema for the resource.
func (r *Queue) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Scan agent queue resource.",
		Attributes: map[string]schema.Attribute{
			"max_queue_per_agent": schema.Int64Attribute{
				Description: "Scan agent queue count.",
				Required:    true,
			},
		},
	}
}

func (r *Queue) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *Queue) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan queueModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	queue := plan.MaxQueuePerAgent.ValueInt64()

	// Add new scan queue
	_, err := r.client.CreateQueue(ctx, int(queue))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT Scan agent queue count",
			"Could not update Scan agent queue value, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetQueue(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Scan agent queue count",
			"Could not read OPSWAT Scan agent queue count "+err.Error(),
		)
		return
	}

	plan.MaxQueuePerAgent = types.Int64Value(int64(result.MaxQueuePerAgent))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *Queue) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state queueModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed item
	result, err := r.client.GetQueue(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Scan agent queue count",
			"Could not read OPSWAT Scan agent queue count "+err.Error(),
		)
		return
	}

	state.MaxQueuePerAgent = types.Int64Value(int64(result.MaxQueuePerAgent))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *Queue) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan queueModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	queue := plan.MaxQueuePerAgent.ValueInt64()

	// Update existing item
	_, err := r.client.UpdateQueue(ctx, int(queue))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT Scan agent queue count",
			"Could not update Scan agent queue value, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items
	result, err := r.client.GetQueue(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT Scan agent queue count",
			"Could not read OPSWAT Scan agent queue count "+err.Error(),
		)
		return
	}

	plan.MaxQueuePerAgent = types.Int64Value(int64(result.MaxQueuePerAgent))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *Queue) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
