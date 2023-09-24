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

package dialogflowcx

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DialogflowCXSecuritySettingsAssetType string = "{{location}}-dialogflow.googleapis.com/SecuritySettings"

func ResourceConverterDialogflowCXSecuritySettings() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DialogflowCXSecuritySettingsAssetType,
		Convert:   GetDialogflowCXSecuritySettingsCaiObject,
	}
}

func GetDialogflowCXSecuritySettingsCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDialogflowCXSecuritySettingsApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DialogflowCXSecuritySettingsAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "SecuritySettings",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDialogflowCXSecuritySettingsApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXSecuritySettingsDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	redactionStrategyProp, err := expandDialogflowCXSecuritySettingsRedactionStrategy(d.Get("redaction_strategy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redaction_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(redactionStrategyProp)) && (ok || !reflect.DeepEqual(v, redactionStrategyProp)) {
		obj["redactionStrategy"] = redactionStrategyProp
	}
	redactionScopeProp, err := expandDialogflowCXSecuritySettingsRedactionScope(d.Get("redaction_scope"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redaction_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(redactionScopeProp)) && (ok || !reflect.DeepEqual(v, redactionScopeProp)) {
		obj["redactionScope"] = redactionScopeProp
	}
	inspectTemplateProp, err := expandDialogflowCXSecuritySettingsInspectTemplate(d.Get("inspect_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("inspect_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(inspectTemplateProp)) && (ok || !reflect.DeepEqual(v, inspectTemplateProp)) {
		obj["inspectTemplate"] = inspectTemplateProp
	}
	deidentifyTemplateProp, err := expandDialogflowCXSecuritySettingsDeidentifyTemplate(d.Get("deidentify_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deidentify_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(deidentifyTemplateProp)) && (ok || !reflect.DeepEqual(v, deidentifyTemplateProp)) {
		obj["deidentifyTemplate"] = deidentifyTemplateProp
	}
	purgeDataTypesProp, err := expandDialogflowCXSecuritySettingsPurgeDataTypes(d.Get("purge_data_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purge_data_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(purgeDataTypesProp)) && (ok || !reflect.DeepEqual(v, purgeDataTypesProp)) {
		obj["purgeDataTypes"] = purgeDataTypesProp
	}
	audioExportSettingsProp, err := expandDialogflowCXSecuritySettingsAudioExportSettings(d.Get("audio_export_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("audio_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(audioExportSettingsProp)) && (ok || !reflect.DeepEqual(v, audioExportSettingsProp)) {
		obj["audioExportSettings"] = audioExportSettingsProp
	}
	insightsExportSettingsProp, err := expandDialogflowCXSecuritySettingsInsightsExportSettings(d.Get("insights_export_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("insights_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(insightsExportSettingsProp)) && (ok || !reflect.DeepEqual(v, insightsExportSettingsProp)) {
		obj["insightsExportSettings"] = insightsExportSettingsProp
	}
	retentionWindowDaysProp, err := expandDialogflowCXSecuritySettingsRetentionWindowDays(d.Get("retention_window_days"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionWindowDaysProp)) && (ok || !reflect.DeepEqual(v, retentionWindowDaysProp)) {
		obj["retentionWindowDays"] = retentionWindowDaysProp
	}
	retentionStrategyProp, err := expandDialogflowCXSecuritySettingsRetentionStrategy(d.Get("retention_strategy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionStrategyProp)) && (ok || !reflect.DeepEqual(v, retentionStrategyProp)) {
		obj["retentionStrategy"] = retentionStrategyProp
	}

	return obj, nil
}

func expandDialogflowCXSecuritySettingsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRedactionStrategy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRedactionScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsInspectTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsDeidentifyTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsPurgeDataTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcsBucket, err := expandDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(original["gcs_bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcsBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcsBucket"] = transformedGcsBucket
	}

	transformedAudioExportPattern, err := expandDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(original["audio_export_pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudioExportPattern); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audioExportPattern"] = transformedAudioExportPattern
	}

	transformedEnableAudioRedaction, err := expandDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(original["enable_audio_redaction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableAudioRedaction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableAudioRedaction"] = transformedEnableAudioRedaction
	}

	transformedAudioFormat, err := expandDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(original["audio_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudioFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audioFormat"] = transformedAudioFormat
	}

	return transformed, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsInsightsExportSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableInsightsExport, err := expandDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(original["enable_insights_export"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableInsightsExport); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableInsightsExport"] = transformedEnableInsightsExport
	}

	return transformed, nil
}

func expandDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRetentionWindowDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRetentionStrategy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}