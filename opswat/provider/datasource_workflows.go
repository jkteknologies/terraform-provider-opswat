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
	resp.TypeName = req.ProviderTypeName + "_workflows"
}

// Schema defines the schema for the data source.
func (d *Workflows) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global workflows datasource.",
		Attributes: map[string]schema.Attribute{
			"workflows": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_cert": schema.BoolAttribute{
							Description: "Generate batch signature with certificate - Use certificate to generate batch signature flag.",
							Computed:    true,
						},
						"allow_cert_cert": schema.StringAttribute{
							Description: "Certificate used for barch signing.",
							Computed:    true,
						},
						"allow_cert_cert_validity": schema.Int64Attribute{
							Description: "Certificate validity (hours).",
							Computed:    true,
						},
						"allow_local_files": schema.BoolAttribute{
							Description: "Process files from servers - Allow scan on server flag.",
							Computed:    true,
						},
						"allow_local_files_white_list": schema.BoolAttribute{
							Description: "Process files from servers flag (false = ALLOW ALL EXCEPT, true = DENY ALL EXCEPT).",
							Computed:    true,
						},
						"allow_local_files_local_paths": schema.StringAttribute{
							Description: "Paths.",
							Computed:    true,
						},
						//"description": schema.StringAttribute{
						//	Description: "Workflow description.",
						//	Computed:    true,
						//},
						//"id": schema.Int64Attribute{
						//	Description: "Workflow id.",
						//	Computed:    true,
						//},
						//"include_webhook_signature": schema.BoolAttribute{
						//	Description: "Webhook - Include webhook signature flag.",
						//	Computed:    true,
						//},
						//"include_webhook_signature_certificate_id": schema.Int64Attribute{
						//	Description: "Webhook - Certificate id.",
						//	Computed:    true,
						//},
						//"last_modified": schema.Int64Attribute{
						//	Description: "Last modified timestamp (unix epoch).",
						//	Computed:    true,
						//},
						//"mutable": schema.BoolAttribute{
						//	Description: "mutable flag?.",
						//	Computed:    true,
						//},
						//"name": schema.StringAttribute{
						//	Description: "Workflow name.",
						//	Computed:    true,
						//},
						//"workflow_id": schema.StringAttribute{
						//	Description: "Workflow id.",
						//	Computed:    true,
						//},
						//"zone_id": schema.StringAttribute{
						//	Description: "Workflow network zone id.",
						//	Computed:    true,
						//},
						//scan_allowed
						//result_allowed
						//option_values
						//user_agents
					},
				},
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

type workflowsDataSourceModel struct {
	Workflows []workflowModel `tfsdk:"workflows"`
}

type workflowModel struct {
	AllowCert                types.Bool   `tfsdk:"allow_cert"`
	AllowCertCert            types.String `tfsdk:"allow_cert_cert"`
	AllowCertCertValidity    types.Int64  `tfsdk:"allow_cert_cert_validity"`
	AllowLocalFiles          types.Bool   `tfsdk:"allow_local_files"`
	AllowLocalFilesWhiteList types.Bool   `tfsdk:"allow_local_files_white_list"`
	AllowLocalFilesLocalPath types.String `tfsdk:"allow_local_files_local_paths"`
}

// Read refreshes the Terraform state with the latest data.
func (d *Workflows) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowsDataSourceModel

	result, err := d.client.GetWorkflows()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT workflows",
			err.Error(),
		)
		return
	}

	//fmt.Println("WORKFLOWS")
	//fmt.Printf("Workflows : %+v", result)

	//fmt.Println("RESULT")
	//fmt.Printf("Workflows : %+v", result)

	for _, workflow := range result {
		workflowState := workflowModel{
			AllowCert:                types.BoolValue(workflow.AllowCert),
			AllowCertCert:            types.StringValue(workflow.AllowCertCert),
			AllowCertCertValidity:    types.Int64Value(int64(workflow.AllowCertCertValidity)),
			AllowLocalFiles:          types.BoolValue(workflow.AllowLocalFiles),
			AllowLocalFilesWhiteList: types.BoolValue(workflow.AllowLocalFilesWhiteList),
		}

		for _, items := range workflow.AllowLocalFilesLocalPaths {
			fmt.Printf("Items : %+v", items)

			state.Workflows = append(state.Workflows, workflowModel{
				AllowLocalFilesLocalPath: types.StringValue(items),
			})
		}

		//fmt.Println("PARSED WORKFLOWS")
		//spew.Dump(workflowState)

		state.Workflows = append(state.Workflows, workflowState)

	}
	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
