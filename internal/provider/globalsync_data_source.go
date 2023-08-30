package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &globalSyncDataSource{}
	_ datasource.DataSourceWithConfigure = &globalSyncDataSource{}
)

// NewGlobalSyncDataSource is a helper function to simplify the provider implementation.
func NewGlobalSyncDataSource() datasource.DataSource {
	return &globalSyncDataSource{}
}

// globalSyncDataSource is the data source implementation.
type globalSyncDataSource struct {
	client *opswatClient.Client
}

// Metadata returns the data source type name.
func (d *globalSyncDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_globalSync"
}

// Schema defines the schema for the data source.
func (d *globalSyncDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"timeout": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

// timeouts maps the data source schema data.
type timeoutsDataSourceModel struct {
	Timeout timeoutsModel `tfsdk:"timeout"`
}

// timeouts maps coffees schema data.
type timeoutsModel struct {
	Timeout types.Int64 `tfsdk:"timeout"`
}

// Read refreshes the Terraform state with the latest data.
func (d *globalSyncDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state timeoutsDataSourceModel

	result, err := d.client.GetGlobalSync()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT Global sync timeout",
			err.Error(),
		)
		return
	}

	// Map response body to model
	timeoutState := timeoutsModel{
		Timeout: types.Int64Value(int64(result.Timeout)),
	}

	fmt.Println("result output: " + fmt.Sprintf("%d", result))

	var test = types.Int64Value(int64(result.Timeout))
	fmt.Println("test output: " + fmt.Sprintf("%d", test))

	state.Timeout = timeoutState

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *globalSyncDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*opswatClient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
