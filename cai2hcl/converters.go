package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/services/resourcemanager"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tpg_provider "github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
)

var provider *schema.Provider
var Converters []common.Converter

func GetConverters() []common.Converter {
	provider = tpg_provider.Provider()

	return []common.Converter{
		compute.NewComputeInstanceConverter(provider),
		compute.NewComputeForwardingRuleConverter(provider),
		resourcemanager.NewProjectConverter(provider),
	}
}
