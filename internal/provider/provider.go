package opswatProvider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	opswatClient "terraform-provider-opswat/internal/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &opswatProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &opswatProvider{
			version: version,
		}
	}
}

// opswatProvider is the provider implementation.
type opswatProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *opswatProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "opswat"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *opswatProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Interact with OPSWAT.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for OPSWAT API. May also be provided via OPSWAT_HOST environment variable.",
				Required:    true,
			},
			"apikey": schema.StringAttribute{
				Description: "API KEY for OPSWAT API. May also be provided via OPSWAT_APIKEY environment variable.",
				Required:    true,
			},
		},
	}
}

// opswatProviderModel maps provider schema data to a Go type.
type opswatProviderModel struct {
	Host   types.String `tfsdk:"host"`
	Apikey types.String `tfsdk:"apikey"`
}

func (p *opswatProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring OPSWAT client")

	// Retrieve provider data from configuration
	var config opswatProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown OPSWAT API Host",
			"The provider cannot create the OPSWAT API client as there is an unknown configuration value for the OPSWAT API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPSWAT_HOST environment variable.",
		)
	}

	if config.Apikey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("apikey"),
			"Unknown OPSWAT API key",
			"The provider cannot create the OPSWAT API client as there is an unknown configuration value for the OPSWAT API apikey. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the OPSWAT_APIKEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("OPSWAT_HOST")
	apikey := os.Getenv("OPSWAT_APIKEY")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Apikey.IsNull() {
		apikey = config.Apikey.ValueString()
	}

	ctx = tflog.SetField(ctx, "OPSWAT_HOST", host)
	ctx = tflog.SetField(ctx, "OPSWAT_APIKEY", apikey)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "OPSWAT_APIKEY")

	tflog.Debug(ctx, "Creating OPSWAT client")

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing OPSWAT API hostname",
			"The provider cannot create the OPSWAT API client as there is a missing or empty value for the OPSWAT API host. "+
				"Set the host value in the configuration or use the OPSWAT_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apikey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("apikey"),
			"Missing OPSWAT API key",
			"The provider cannot create the OPSWAT API client as there is a missing or empty value for the OPSWAT API apikey. "+
				"Set the apikey value in the configuration or use the OPSWAT_APIKEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new opswat client using the configuration values

	// init http request
	client, err := opswatClient.NewClient(&host, &apikey)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OPSWAT API Client",
			"An unexpected error occurred when creating the OPSWAT API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"OPSWAT Client Error: "+err.Error(),
		)
		return
	}

	// Make the opswat client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured OPSWAT client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *opswatProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGlobalSyncDataSource,
		NewWorkflows,
		NewUserDirectory,
	}
}

// Resources defines the resources implemented in the provider.
func (p *opswatProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGlobalSync,
		NewSession,
		NewQuarantine,
		NewQueue,
		NewWorkflow,
		NewDir,
		NewUser,
		NewUserRole,
		NewScanHistory,
	}
}
