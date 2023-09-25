package common

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
)

// Converter which aggregates all service-specific converters in the same interface.
type UberConverter struct {
	// Mapping between asset type (i.e. compute.googleapis.com/Instance) to converter name.
	ConverterNamesPerAssetType map[string]string

	// Collection of asset name formats (i.e. projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/forwardingRules)
	// together with corresponding converter name.
	AssetNameRegexpConverterPairs []RegexpNamePair

	// Mapping between converter name and converter constructor.
	ConverterMap map[string]Converter
}

// Convert assets of any of known types to the list of HCL blocks.
func (c UberConverter) Convert(assets []*caiasset.Asset) ([]*HCLResourceBlock, error) {
	// Group resources from the same tf resource type for convert.
	// tf -> cai has 1:N mappings occasionally
	groups := make(map[string][]*caiasset.Asset)
	for _, asset := range assets {
		name, ok := c.ConverterNamesPerAssetType[asset.Type]

		if !ok {
			name, ok = FindMatchingConverterName(asset.Name, c.AssetNameRegexpConverterPairs)
		}

		if ok {
			groups[name] = append(groups[name], asset)
		}
	}

	allBlocks := []*HCLResourceBlock{}
	for name, assets := range groups {
		converter, ok := c.ConverterMap[name]
		if !ok {
			continue
		}
		newBlocks, err := converter.Convert(assets)
		if err != nil {
			return nil, err
		}

		allBlocks = append(allBlocks, newBlocks...)
	}

	return allBlocks, nil
}
