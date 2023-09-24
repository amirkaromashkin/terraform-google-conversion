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

package gkebackup

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const GKEBackupRestorePlanIAMAssetType string = "gkebackup.googleapis.com/RestorePlan"

func ResourceConverterGKEBackupRestorePlanIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupRestorePlanIAMAssetType,
		Convert:           GetGKEBackupRestorePlanIamPolicyCaiObject,
		MergeCreateUpdate: MergeGKEBackupRestorePlanIamPolicy,
	}
}

func ResourceConverterGKEBackupRestorePlanIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupRestorePlanIAMAssetType,
		Convert:           GetGKEBackupRestorePlanIamBindingCaiObject,
		FetchFullResource: FetchGKEBackupRestorePlanIamPolicy,
		MergeCreateUpdate: MergeGKEBackupRestorePlanIamBinding,
		MergeDelete:       MergeGKEBackupRestorePlanIamBindingDelete,
	}
}

func ResourceConverterGKEBackupRestorePlanIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupRestorePlanIAMAssetType,
		Convert:           GetGKEBackupRestorePlanIamMemberCaiObject,
		FetchFullResource: FetchGKEBackupRestorePlanIamPolicy,
		MergeCreateUpdate: MergeGKEBackupRestorePlanIamMember,
		MergeDelete:       MergeGKEBackupRestorePlanIamMemberDelete,
	}
}

func GetGKEBackupRestorePlanIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupRestorePlanIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetGKEBackupRestorePlanIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupRestorePlanIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetGKEBackupRestorePlanIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupRestorePlanIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeGKEBackupRestorePlanIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeGKEBackupRestorePlanIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeGKEBackupRestorePlanIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeGKEBackupRestorePlanIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeGKEBackupRestorePlanIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newGKEBackupRestorePlanIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/restorePlans/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: GKEBackupRestorePlanIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchGKEBackupRestorePlanIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		GKEBackupRestorePlanIamUpdaterProducer,
		d,
		config,
		"//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/restorePlans/{{name}}",
		GKEBackupRestorePlanIAMAssetType,
	)
}