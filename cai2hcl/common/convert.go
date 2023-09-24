package common

import (
	"fmt"
	"regexp"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	"github.com/hashicorp/hcl/hcl/printer"
	"github.com/hashicorp/hcl/v2/hclwrite"
	tpg_provider "github.com/hashicorp/terraform-provider-google-beta/google-beta/provider"
	"github.com/zclconf/go-cty/cty"
)

// Initializes map of converters.
func CreateConverterMap(converterFactories map[string]ConverterFactory) map[string]Converter {
	provider := tpg_provider.Provider()

	result := make(map[string]Converter, len(converterFactories))
	for name, factory := range converterFactories {
		result[name] = factory(name, provider.ResourcesMap[name].Schema)
	}

	return result
}

// Implementation for the entry point conversion from Assets to HCL.
func Convert(assets []*caiasset.Asset, converterNamesPerAssetType map[string]string, assetNameRegexpConverterPairs []RegexpNamePair, converterMap map[string]Converter) ([]byte, error) {
	// Group resources from the same tf resource type for convert.
	// tf -> cai has 1:N mappings occasionally
	groups := make(map[string][]*caiasset.Asset)
	for _, asset := range assets {
		name, ok := converterNamesPerAssetType[asset.Type]

		if !ok {
			name, ok = findMatchingConverterName(asset.Name, assetNameRegexpConverterPairs)
		}

		if ok {
			groups[name] = append(groups[name], asset)
		}
	}

	f := hclwrite.NewFile()
	rootBody := f.Body()
	for name, v := range groups {
		converter, ok := converterMap[name]
		if !ok {
			continue
		}
		items, err := converter.Convert(v)
		if err != nil {
			return nil, err
		}

		for _, resourceBlock := range items {
			hclBlock := rootBody.AppendNewBlock("resource", resourceBlock.Labels)
			if err := hclWriteBlock(resourceBlock.Value, hclBlock.Body()); err != nil {
				return nil, err
			}
		}
		if err != nil {
			return nil, err
		}
	}

	return printer.Format(f.Bytes())
}

func hclWriteBlock(val cty.Value, body *hclwrite.Body) error {
	if val.IsNull() {
		return nil
	}
	if !val.Type().IsObjectType() {
		return fmt.Errorf("expect object type only, but type = %s", val.Type().FriendlyName())
	}
	it := val.ElementIterator()
	for it.Next() {
		objKey, objVal := it.Element()
		if objVal.IsNull() {
			continue
		}
		objValType := objVal.Type()
		switch {
		case objValType.IsObjectType():
			newBlock := body.AppendNewBlock(objKey.AsString(), nil)
			if err := hclWriteBlock(objVal, newBlock.Body()); err != nil {
				return err
			}
		case objValType.IsCollectionType():
			if objVal.LengthInt() == 0 {
				continue
			}
			// Presumes map should not contain object type.
			if !objValType.IsMapType() && objValType.ElementType().IsObjectType() {
				listIterator := objVal.ElementIterator()
				for listIterator.Next() {
					_, listVal := listIterator.Element()
					subBlock := body.AppendNewBlock(objKey.AsString(), nil)
					if err := hclWriteBlock(listVal, subBlock.Body()); err != nil {
						return err
					}
				}
				continue
			}
			fallthrough
		default:
			if objValType.FriendlyName() == "string" && objVal.AsString() == "" {
				continue
			}
			body.SetAttributeValue(objKey.AsString(), objVal)
		}
	}
	return nil
}

func findMatchingConverterName(str string, assetNameRegexpConverterPairs []RegexpNamePair) (string, bool) {
	for _, pair := range assetNameRegexpConverterPairs {
		match, _ := regexp.MatchString(pair.AssetNameRegexp, str)
		if match {
			return pair.ConverterName, true
		}
	}

	return "", false
}