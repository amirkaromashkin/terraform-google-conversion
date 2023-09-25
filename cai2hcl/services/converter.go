package services

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/common"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/cai2hcl/services/resourcemanager"
)

var Converter = common.UberConverter{
	ConverterNamesPerAssetType: joinStringMaps([]map[string]string{
		compute.Converter.ConverterNamesPerAssetType,
		resourcemanager.Converter.ConverterNamesPerAssetType,
	}),
	AssetNameRegexpConverterPairs: append(
		compute.Converter.AssetNameRegexpConverterPairs,
		resourcemanager.Converter.AssetNameRegexpConverterPairs...,
	),
	ConverterMap: joinConverterMaps([]map[string]common.Converter{
		compute.Converter.ConverterMap,
		resourcemanager.Converter.ConverterMap,
	}),
}

func joinStringMaps(arr []map[string]string) map[string]string {
	result := make(map[string]string)

	for _, m := range arr {
		for key, value := range m {
			if _, hasKey := result[key]; hasKey {
				panic("Converters from different services are not unique")
			}

			result[key] = value
		}
	}

	return result
}

func joinConverterMaps(arr []map[string]common.Converter) map[string]common.Converter {
	result := make(map[string]common.Converter)

	for _, m := range arr {
		for key, value := range m {
			if _, hasKey := result[key]; hasKey {
				panic("Converters from different services are not unique")
			}

			result[key] = value
		}
	}

	return result
}
