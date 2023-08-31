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
	_ resource.Resource              = &configSession{}
	_ resource.ResourceWithConfigure = &configSession{}
)

// NewConfigSession is a helper function to simplify the provider implementation.
func NewConfigSession() resource.Resource {
	return &configSession{}
}

// configSession is the resource implementation.
type configSession struct {
	client *opswatClient.Client
}

// Metadata returns the resource type name.
func (r *configSession) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_session"
}

// Schema defines the schema for the resource.
func (r *configSession) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global file sync can timeout resource.",
		Attributes: map[string]schema.Attribute{
			"absolutesessiontimeout": schema.Int64Attribute{
				Description: "The interval (in milliseconds) for overall session length timeout (regardless of activity). minimal 300000. 0 - for infinity sessions.",
				Required:    true,
			},
			"allowcrossipsessions": schema.BoolAttribute{
				Description: "Allow requests from the same user to come from different IPs.",
				Required:    true,
			},
			"allowduplicatesession": schema.BoolAttribute{
				Description: "Allow same user to have multiple active sessions.",
				Required:    true,
			},
			"sessiontimeout": schema.Int64Attribute{
				Description: "The interval (in milliseconds) for the user's session timeout, based on last activity. Timer starts after the last activity for the apikey. minimal - 60000. 0 - for infinity sessions.",
				Required:    true,
			},
		},
	}
}

// timeouts maps timeout schema data.
type configSessionModel struct {
	AbsoluteSessionTimeout types.Int64 `tfsdk:"absolutesessiontimeout"`
	AllowCrossIpSessions   types.Bool  `tfsdk:"allowcrossipsessions"`
	AllowDuplicateSession  types.Bool  `tfsdk:"allowduplicatesession"`
	SessionTimeout         types.Int64 `tfsdk:"sessiontimeout"`
}

func (r *configSession) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *configSession) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan configSessionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.ConfigSession{
		AbsoluteSessionTimeout: int(plan.AbsoluteSessionTimeout.ValueInt64()),
		AllowCrossIpSessions:   bool(plan.AllowCrossIpSessions.ValueBool()),
		AllowDuplicateSession:  bool(plan.AllowDuplicateSession.ValueBool()),
		SessionTimeout:         int(plan.SessionTimeout.ValueInt64()),
	}

	// Update existing order
	_, err := r.client.CreateConfigSession(json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT session config",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not populated.
	result, err := r.client.GetConfigSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	plan.AbsoluteSessionTimeout = types.Int64Value(int64(result.AbsoluteSessionTimeout))
	plan.AllowCrossIpSessions = types.BoolValue(result.AllowCrossIpSessions)
	plan.AllowDuplicateSession = types.BoolValue(result.AllowDuplicateSession)
	plan.SessionTimeout = types.Int64Value(int64(result.SessionTimeout))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *configSession) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state configSessionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed order value from OPSWAT
	result, err := r.client.GetConfigSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	state.AbsoluteSessionTimeout = types.Int64Value(int64(result.AbsoluteSessionTimeout))
	state.AllowCrossIpSessions = types.BoolValue(result.AllowCrossIpSessions)
	state.AllowDuplicateSession = types.BoolValue(result.AllowDuplicateSession)
	state.SessionTimeout = types.Int64Value(int64(result.SessionTimeout))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *configSession) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan configSessionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	json := opswatClient.ConfigSession{
		AbsoluteSessionTimeout: int(plan.AbsoluteSessionTimeout.ValueInt64()),
		AllowCrossIpSessions:   bool(plan.AllowCrossIpSessions.ValueBool()),
		AllowDuplicateSession:  bool(plan.AllowDuplicateSession.ValueBool()),
		SessionTimeout:         int(plan.SessionTimeout.ValueInt64()),
	}

	// Update existing order
	_, err := r.client.UpdateConfigSession(json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating OPSWAT session config",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}

	// Fetch updated items from GetOrder as UpdateOrder items are not populated.
	result, err := r.client.GetConfigSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT session config",
			"Could not read OPSWAT session config "+err.Error(),
		)
		return
	}

	plan.AbsoluteSessionTimeout = types.Int64Value(int64(result.AbsoluteSessionTimeout))
	plan.AllowCrossIpSessions = types.BoolValue(result.AllowCrossIpSessions)
	plan.AllowDuplicateSession = types.BoolValue(result.AllowDuplicateSession)
	plan.SessionTimeout = types.Int64Value(int64(result.SessionTimeout))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *configSession) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
