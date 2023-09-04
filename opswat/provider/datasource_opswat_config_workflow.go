package opswatProvider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/opswat/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &globalConfigWorkflow{}
	_ datasource.DataSourceWithConfigure = &globalConfigWorkflow{}
)

// NewGlobalConfigWorkflow is a helper function to simplify the provider implementation.
func NewGlobalConfigWorkflow() datasource.DataSource {
	return &globalConfigWorkflow{}
}

// globalConfigWorkflow is the data source implementation.
type globalConfigWorkflow struct {
	client *opswatClient.Client
}

// Metadata returns the data source type name.
func (d *globalConfigWorkflow) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_workflow"
}

// Schema defines the schema for the data source.
func (d *globalConfigWorkflow) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global workflows datasource.",
		Attributes: map[string]schema.Attribute{
			"AllowCert": schema.BoolAttribute{
				Description: "Global file sync can timeout.",
				Required:    true,
			},
		},
	}
}

func (d *globalConfigWorkflow) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

// timeouts maps timeout schema data.
type configWorkflow struct {
	Timeout types.Int64 `tfsdk:"timeout"`
}

// Read refreshes the Terraform state with the latest data.
func (d *globalConfigWorkflow) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state timeoutModel

	result, err := d.client.GetGlobalSync()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT Global sync timeout",
			err.Error(),
		)
		return
	}

	//fmt.Println(types.Int64Value(int64(result.Timeout)))

	state.Timeout = types.Int64Value(int64(result.Timeout))

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
