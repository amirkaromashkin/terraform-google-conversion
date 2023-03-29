// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeSubnetworkIAMAssetType string = "compute.googleapis.com/Subnetwork"

func resourceConverterComputeSubnetworkIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeSubnetworkIamPolicy,
	}
}

func resourceConverterComputeSubnetworkIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamBindingCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamBinding,
		MergeDelete:       MergeComputeSubnetworkIamBindingDelete,
	}
}

func resourceConverterComputeSubnetworkIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamMemberCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamMember,
		MergeDelete:       MergeComputeSubnetworkIamMemberDelete,
	}
}

func GetComputeSubnetworkIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, expandIamPolicyBindings)
}

func GetComputeSubnetworkIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, expandIamRoleBindings)
}

func GetComputeSubnetworkIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, expandIamMemberBindings)
}

func MergeComputeSubnetworkIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeSubnetworkIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeComputeSubnetworkIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeComputeSubnetworkIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeComputeSubnetworkIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newComputeSubnetworkIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ComputeSubnetworkIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeSubnetworkIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("subnetwork"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ComputeSubnetworkIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}",
		ComputeSubnetworkIAMAssetType,
	)
}