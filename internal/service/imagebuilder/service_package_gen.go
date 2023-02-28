// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []func(context.Context) (datasource.DataSourceWithConfigure, error) {
	return []func(context.Context) (datasource.DataSourceWithConfigure, error){}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []func(context.Context) (resource.ResourceWithConfigure, error) {
	return []func(context.Context) (resource.ResourceWithConfigure, error){}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{
		"aws_imagebuilder_component":                     DataSourceComponent,
		"aws_imagebuilder_components":                    DataSourceComponents,
		"aws_imagebuilder_container_recipe":              DataSourceContainerRecipe,
		"aws_imagebuilder_container_recipes":             DataSourceContainerRecipes,
		"aws_imagebuilder_distribution_configuration":    DataSourceDistributionConfiguration,
		"aws_imagebuilder_distribution_configurations":   DataSourceDistributionConfigurations,
		"aws_imagebuilder_image":                         DataSourceImage,
		"aws_imagebuilder_image_pipeline":                DataSourceImagePipeline,
		"aws_imagebuilder_image_pipelines":               DataSourceImagePipelines,
		"aws_imagebuilder_image_recipe":                  DataSourceImageRecipe,
		"aws_imagebuilder_image_recipes":                 DataSourceImageRecipes,
		"aws_imagebuilder_infrastructure_configuration":  DataSourceInfrastructureConfiguration,
		"aws_imagebuilder_infrastructure_configurations": DataSourceInfrastructureConfigurations,
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) map[string]func() *schema.Resource {
	return map[string]func() *schema.Resource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

var ServicePackage = &servicePackage{}
