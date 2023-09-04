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
	_ datasource.DataSource              = &Workflows{}
	_ datasource.DataSourceWithConfigure = &Workflows{}
)

// NewGlobalWorkflows is a helper function to simplify the provider implementation.
func NewWorkflows() datasource.DataSource {
	return &Workflows{}
}

// Workflows is the data source implementation.
type Workflows struct {
	client *opswatClient.Client
}

// Metadata returns the data source type name.
func (d *Workflows) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_workflow"
}

// Schema defines the schema for the data source.
func (d *Workflows) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global workflows datasource.",
		Attributes: map[string]schema.Attribute{
			"allow_cert": schema.BoolAttribute{
				Description: "Generate batch signature with certificate flag.",
				Required:    false,
			},
		},
	}
}

func (d *Workflows) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
	AllowCert types.Bool `tfsdk:"allow_cert"`
}

// Read refreshes the Terraform state with the latest data.
func (d *Workflows) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state configWorkflow

	result, err := d.client.GetWorkflows()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT workflow",
			err.Error(),
		)
		return
	}

	//fmt.Println(types.Int64Value(int64(result.Timeout)))

	state.AllowCert = types.BoolValue(result.AllowCert)

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
