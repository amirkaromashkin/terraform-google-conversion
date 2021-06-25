// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

func GetIapWebTypeAppEngineIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeAppEngineIamAsset(d, config, expandIamPolicyBindings)
}

func GetIapWebTypeAppEngineIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeAppEngineIamAsset(d, config, expandIamRoleBindings)
}

func GetIapWebTypeAppEngineIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapWebTypeAppEngineIamAsset(d, config, expandIamMemberBindings)
}

func MergeIapWebTypeAppEngineIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapWebTypeAppEngineIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeIapWebTypeAppEngineIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeIapWebTypeAppEngineIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeIapWebTypeAppEngineIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newIapWebTypeAppEngineIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//iap.googleapis.com/{{webtypeappengine}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "iap.googleapis.com/WebTypeAppEngine",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapWebTypeAppEngineIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{webtypeappengine}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		IapWebTypeAppEngineIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/{{webtypeappengine}}",
		"iap.googleapis.com/WebTypeAppEngine",
	)
}