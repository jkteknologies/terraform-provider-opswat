package opswatProvider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/internal/connectivity"
)

import (
	"context"
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
	resp.TypeName = req.ProviderTypeName + "_file_sync"
}

// Schema defines the schema for the data source.
func (d *globalSyncDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global file sync can timeout datasource.",
		Attributes: map[string]schema.Attribute{
			"timeout": schema.Int64Attribute{
				Description: "Global file sync can timeout.",
				Computed:    true,
			},
		},
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
			fmt.Sprintf("Expected *opswatClient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

// timeouts maps timeout schema data.
type timeoutModel struct {
	Timeout types.Int64 `tfsdk:"timeout"`
}

// Read refreshes the Terraform state with the latest data.
func (d *globalSyncDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state timeoutModel

	result, err := d.client.GetGlobalSync(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT Global sync timeout",
			err.Error(),
		)
		return
	}

	//fmt.Println(types.Int64Value(int64(result.Timeout))) test

	state.Timeout = types.Int64Value(int64(result.Timeout))

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
