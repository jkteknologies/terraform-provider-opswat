package opswatProvider

import (
	"context"
	"fmt"
	opswatClient "terraform-provider-opswat/internal/connectivity"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &userDirectory{}
	_ datasource.DataSourceWithConfigure = &userDirectory{}
)

// NewGlobalUserDirectory is a helper function to simplify the provider implementation.
func NewUserDirectory() datasource.DataSource {
	return &userDirectory{}
}

// userDirectory is the data source implementation.
type userDirectory struct {
	client *opswatClient.Client
}

// userDirs model maps the data source schema data.
type dirModels struct {
	Dirs []dirModel `tfsdk:"dirs"`
}

type dirModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Type             types.String `tfsdk:"type"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	Name             types.String `tfsdk:"name"`
	UserIdentifiedBy types.String `tfsdk:"user_identified_by"`
	Sp               SPModel      `tfsdk:"sp"`
	Role             RoleModel    `tfsdk:"role"`
	Version          types.String `tfsdk:"version"`
	Idp              IDPModel     `tfsdk:"idp"`
}

type SPModel struct {
	LoginUrl           types.String `tfsdk:"login_url"`
	SupportLogoutUrl   types.Bool   `tfsdk:"support_logout_url"`
	SupportPrivateKey  types.Bool   `tfsdk:"support_private_key"`
	SupportEntityId    types.Bool   `tfsdk:"support_entity_id"`
	EnableIdpInitiated types.Bool   `tfsdk:"enable_idp_initiated"`
	EntityId           types.String `tfsdk:"entity_id"`
}

type RoleModel struct {
	Option  types.String   `tfsdk:"option"`
	Details []DetailsModel `tfsdk:"details"`
}

type DetailsModel struct {
	Key    types.String  `tfsdk:"key"`
	Values []ValuesModel `tfsdk:"values"`
}

type ValuesModel struct {
	Condition string   `tfsdk:"condition"`
	RoleIds   []string `tfsdk:"role_ids"`
	Type      string   `tfsdk:"type"`
}

type IDPModel struct {
	AuthnRequestSigned types.Bool        `tfsdk:"authn_request_signed"`
	EntityId           types.String      `tfsdk:"entity_id"`
	LoginMethod        LoginMethodModel  `tfsdk:"login_method"`
	LogoutMethod       LogoutMethodModel `tfsdk:"logout_method"`
	ValidUntil         types.String      `tfsdk:"valid_until"`
	X509Cert           []string          `tfsdk:"x509_cert"`
}

type LoginMethodModel struct {
	Post     types.String `tfsdk:"post"`
	Redirect types.String `tfsdk:"redirect"`
}

type LogoutMethodModel struct {
	Redirect types.String `tfsdk:"redirect"`
}

// Metadata returns the data source type name.
func (d *userDirectory) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_userdirectories"
}

// Schema defines the schema for the data source.
func (d *userDirectory) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"dirs": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Userdirectory id.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Local, AD, LDAP, OIDC or SAML",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Enabled flag",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Directory name",
							Computed:    true,
						},
						"user_identified_by": schema.StringAttribute{
							Description: "User name alias via claims under profile scope",
							Computed:    true,
						},
						"sp": schema.ObjectAttribute{
							Computed: true,
							AttributeTypes: map[string]attr.Type{
								"login_url":            types.StringType,
								"support_logout_url":   types.BoolType,
								"support_private_key":  types.BoolType,
								"support_entity_id":    types.BoolType,
								"enable_idp_initiated": types.BoolType,
								"entity_id":            types.StringType,
							},
						},
						"idp": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"authn_request_signed": schema.BoolAttribute{
									Computed: true,
								},
								"entity_id": schema.StringAttribute{
									Computed: true,
								},
								"valid_until": schema.StringAttribute{
									Computed: true,
								},
								"x509_cert": schema.ListAttribute{
									ElementType: types.StringType,
									Computed:    true,
								},
								"login_method": schema.ObjectAttribute{
									Computed: true,
									AttributeTypes: map[string]attr.Type{
										"post":     types.StringType,
										"redirect": types.StringType,
									},
								},
								"logout_method": schema.ObjectAttribute{
									Computed: true,
									AttributeTypes: map[string]attr.Type{
										"redirect": types.StringType,
									},
								},
							},
						},
						"version": schema.StringAttribute{
							Description: "Version number",
							Computed:    true,
						},
						"role": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"option": schema.StringAttribute{
									Computed: true,
								},
								"details": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Computed: true,
											},
											"values": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"condition": schema.StringAttribute{
															Computed: true,
														},
														"role_ids": schema.ListAttribute{
															ElementType: types.StringType,
															Computed:    true,
														},
														"type": schema.StringAttribute{
															Computed: true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *userDirectory) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Read refreshes the Terraform state with the latest data.
func (d *userDirectory) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Get current state
	var state dirModels

	// Get refreshed session value from OPSWAT
	userDirs, err := d.client.GetDirs(ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading OPSWAT user directories",
			"Could not read OPSWAT user directories "+err.Error(),
		)
		return
	}

	for _, result := range userDirs {
		dirState := dirModel{
			ID:               types.Int64Value(int64(result.ID)),
			Type:             types.StringValue(result.Type),
			Enabled:          types.BoolValue(result.Enabled),
			Name:             types.StringValue(result.Name),
			UserIdentifiedBy: types.StringValue(result.UserIdentifiedBy),
			Sp: SPModel{
				LoginUrl:           types.StringValue(result.Sp.LoginUrl),
				SupportLogoutUrl:   types.BoolValue(result.Sp.SupportLogoutUrl),
				SupportPrivateKey:  types.BoolValue(result.Sp.SupportPrivateKey),
				EnableIdpInitiated: types.BoolValue(result.Sp.EnableIdpInitiated),
				SupportEntityId:    types.BoolValue(result.Sp.SupportEntityId),
				EntityId:           types.StringValue(result.Sp.EntityId),
			},
			Version: types.StringValue(result.Version),
			Idp: IDPModel{
				AuthnRequestSigned: types.BoolValue(result.Idp.AuthnRequestSigned),
				EntityId:           types.StringValue(result.Idp.EntityId),
				LoginMethod: LoginMethodModel{
					Post:     types.StringValue(result.Idp.LoginMethod.Post),
					Redirect: types.StringValue(result.Idp.LoginMethod.Redirect),
				},
				LogoutMethod: LogoutMethodModel{
					Redirect: types.StringValue(result.Idp.LogoutMethod.Redirect),
				},
				ValidUntil: types.StringValue(result.Idp.ValidUntil),
			},
			Role: RoleModel{
				Details: []DetailsModel{},
				Option:  types.StringValue(result.Role.Option),
			},
		}

		dirState.Idp.X509Cert = append(dirState.Idp.X509Cert, result.Idp.X509Cert...)

		for n, details := range result.Role.Details {
			dirState.Role.Details = append(dirState.Role.Details, DetailsModel{
				Key:    types.StringValue(details.Key),
				Values: []ValuesModel{},
			})

			dirState.Role.Details[n].Values = append(dirState.Role.Details[n].Values, ValuesModel{
				Condition: details.Values[0].Condition,
				RoleIds:   details.Values[0].RoleIds,
				Type:      details.Values[0].Type,
			})
		}

		state.Dirs = append(state.Dirs, dirState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
