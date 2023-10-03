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

package databasemigrationservice

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DatabaseMigrationServicePrivateConnectionAssetType string = "datamigration.googleapis.com/PrivateConnection"

func ResourceConverterDatabaseMigrationServicePrivateConnection() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DatabaseMigrationServicePrivateConnectionAssetType,
		Convert:   GetDatabaseMigrationServicePrivateConnectionCaiObject,
	}
}

func GetDatabaseMigrationServicePrivateConnectionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//datamigration.googleapis.com/projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDatabaseMigrationServicePrivateConnectionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DatabaseMigrationServicePrivateConnectionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datamigration/v1/rest",
				DiscoveryName:        "PrivateConnection",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDatabaseMigrationServicePrivateConnectionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDatabaseMigrationServicePrivateConnectionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	vpcPeeringConfigProp, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(d.Get("vpc_peering_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc_peering_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcPeeringConfigProp)) && (ok || !reflect.DeepEqual(v, vpcPeeringConfigProp)) {
		obj["vpcPeeringConfig"] = vpcPeeringConfigProp
	}
	labelsProp, err := expandDatabaseMigrationServicePrivateConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDatabaseMigrationServicePrivateConnectionDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVpcName, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(original["vpc_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVpcName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vpcName"] = transformedVpcName
	}

	transformedSubnet, err := expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(original["subnet"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subnet"] = transformedSubnet
	}

	return transformed, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigVpcName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionVpcPeeringConfigSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatabaseMigrationServicePrivateConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}