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

package firebase

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseAppleAppAssetType string = "firebase.googleapis.com/AppleApp"

func ResourceConverterFirebaseAppleApp() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseAppleAppAssetType,
		Convert:   GetFirebaseAppleAppCaiObject,
	}
}

func GetFirebaseAppleAppCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebase.googleapis.com/projects/{{project}}/iosApps/{{app_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseAppleAppApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseAppleAppAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebase/v1beta1/rest",
				DiscoveryName:        "AppleApp",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseAppleAppApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandFirebaseAppleAppDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bundleIdProp, err := expandFirebaseAppleAppBundleId(d.Get("bundle_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bundle_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(bundleIdProp)) && (ok || !reflect.DeepEqual(v, bundleIdProp)) {
		obj["bundleId"] = bundleIdProp
	}
	appStoreIdProp, err := expandFirebaseAppleAppAppStoreId(d.Get("app_store_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("app_store_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(appStoreIdProp)) && (ok || !reflect.DeepEqual(v, appStoreIdProp)) {
		obj["appStoreId"] = appStoreIdProp
	}
	teamIdProp, err := expandFirebaseAppleAppTeamId(d.Get("team_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("team_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(teamIdProp)) && (ok || !reflect.DeepEqual(v, teamIdProp)) {
		obj["teamId"] = teamIdProp
	}

	return obj, nil
}

func expandFirebaseAppleAppDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppBundleId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppAppStoreId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseAppleAppTeamId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
