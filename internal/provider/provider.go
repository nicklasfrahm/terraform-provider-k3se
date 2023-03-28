package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

const (
	// Namespace is the namespace of the provider.
	Namespace = "nicklasfrahm"
	// Name is the name of the provider.
	Name = "k3se"
)

var _ provider.Provider = (*k3seProvider)(nil)

type k3seProvider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &k3seProvider{}
	}
}

func (p *k3seProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *k3seProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = Name
}

func (p *k3seProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *k3seProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewClusterResource,
	}
}

func (p *k3seProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}
