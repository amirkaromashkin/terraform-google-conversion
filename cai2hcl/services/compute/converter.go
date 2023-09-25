package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
)

var Converter = common.UberConverter{
	ConverterNamesPerAssetType: map[string]string{
		ComputeInstanceAssetType:    "google_compute_instance",
		ComputeHealthCheckAssetType: "google_compute_health_check",
	},
	AssetNameRegexpConverterPairs: []common.RegexpNamePair{
		{AssetNameRegexp: ComputeForwardingRuleAssetNameRegex, ConverterName: "google_compute_forwarding_rule"},
		{AssetNameRegexp: ComputeGlobalForwardingRuleAssetNameRegex, ConverterName: "google_compute_global_forwarding_rule"},
		{AssetNameRegexp: ComputeRegionBackendServiceAssetNameRegex, ConverterName: "google_compute_region_backend_service"},
		{AssetNameRegexp: ComputeBackendServiceAssetNameRegex, ConverterName: "google_compute_backend_service"},
	},
	ConverterMap: common.CreateConverterMap(map[string]common.ConverterFactory{
		"google_compute_instance":               NewComputeInstanceConverter,
		"google_compute_forwarding_rule":        NewComputeForwardingRuleConverter,
		"google_compute_global_forwarding_rule": NewComputeGlobalForwardingRuleConverter,
		"google_compute_region_backend_service": NewComputeRegionBackendServiceConverter,
		"google_compute_backend_service":        NewComputeBackendServiceConverter,
		"google_compute_health_check":           NewComputeHealthCheckConverter,
	}),
}
