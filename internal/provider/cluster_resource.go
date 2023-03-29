package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = (*clusterResource)(nil)

// NewClusterResource is a helper function to simplify the provider implementation.
func NewClusterResource() resource.Resource {
	return &clusterResource{}
}

// clusterResource is a resource implementation that manages a cluster.
type clusterResource struct {
}

// Metadata returns the the resource type name.
func (r *clusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Schema defines the schema for the resource.
func (r *clusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	serverSchema := schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"tls_san": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"write_kubeconfig_mode": schema.StringAttribute{
				Optional: true,
			},
			"node_label": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"node_ip": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"node_external_ip": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"flannel_iface": schema.StringAttribute{
				Optional: true,
			},
		},
	}

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"version": schema.StringAttribute{
				Required: true,
			},
			"cluster": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"server": serverSchema,
				},
			},
			"nodes": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"role": schema.StringAttribute{
							Required: true,
						},
						"ssh": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Required: true,
								},
								"user": schema.StringAttribute{
									Optional: true,
								},
								"key_file": schema.StringAttribute{
									Optional: true,
								},
							},
						},
						"server": serverSchema,
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *clusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (e *clusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (e *clusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (e *clusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
