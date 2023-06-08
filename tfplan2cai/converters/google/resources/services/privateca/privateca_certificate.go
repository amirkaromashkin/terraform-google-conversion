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

package privateca

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const PrivatecaCertificateAssetType string = "privateca.googleapis.com/Certificate"

func ResourceConverterPrivatecaCertificate() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: PrivatecaCertificateAssetType,
		Convert:   GetPrivatecaCertificateCaiObject,
	}
}

func GetPrivatecaCertificateCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificates/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetPrivatecaCertificateApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: PrivatecaCertificateAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/privateca/v1/rest",
				DiscoveryName:        "Certificate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetPrivatecaCertificateApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	lifetimeProp, err := expandPrivatecaCertificateLifetime(d.Get("lifetime"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("lifetime"); !tpgresource.IsEmptyValue(reflect.ValueOf(lifetimeProp)) && (ok || !reflect.DeepEqual(v, lifetimeProp)) {
		obj["lifetime"] = lifetimeProp
	}
	certificateTemplateProp, err := expandPrivatecaCertificateCertificateTemplate(d.Get("certificate_template"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("certificate_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(certificateTemplateProp)) && (ok || !reflect.DeepEqual(v, certificateTemplateProp)) {
		obj["certificateTemplate"] = certificateTemplateProp
	}
	labelsProp, err := expandPrivatecaCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	pemCsrProp, err := expandPrivatecaCertificatePemCsr(d.Get("pem_csr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("pem_csr"); !tpgresource.IsEmptyValue(reflect.ValueOf(pemCsrProp)) && (ok || !reflect.DeepEqual(v, pemCsrProp)) {
		obj["pemCsr"] = pemCsrProp
	}
	configProp, err := expandPrivatecaCertificateConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}

	return obj, nil
}

func expandPrivatecaCertificateLifetime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateCertificateTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandPrivatecaCertificatePemCsr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedX509Config, err := expandPrivatecaCertificateConfigX509Config(original["x509_config"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["x509Config"] = transformedX509Config
	}

	transformedSubjectConfig, err := expandPrivatecaCertificateConfigSubjectConfig(original["subject_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectConfig"] = transformedSubjectConfig
	}

	transformedPublicKey, err := expandPrivatecaCertificateConfigPublicKey(original["public_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPublicKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["publicKey"] = transformedPublicKey
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigX509Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return v, nil
	}
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	if len(original) == 0 {
		return nil, nil
	}
	transformed := make(map[string]interface{})

	caOptions, err := expandPrivatecaCertificateConfigX509ConfigCaOptions(original["ca_options"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["caOptions"] = caOptions

	keyUsage, err := expandPrivatecaCertificateConfigX509ConfigKeyUsage(original["key_usage"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["keyUsage"] = keyUsage

	policyIds, err := expandPrivatecaCertificateConfigX509ConfigPolicyIds(original["policy_ids"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["policyIds"] = policyIds

	aiaOcspServers, err := expandPrivatecaCertificateConfigX509ConfigAiaOcspServers(original["aia_ocsp_servers"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["aiaOcspServers"] = aiaOcspServers

	addExts, err := expandPrivatecaCertificateConfigX509ConfigAdditionalExtensions(original["additional_extensions"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["additionalExtensions"] = addExts

	nameConstraints, err := expandPrivatecaCertificateConfigX509ConfigNameConstraints(original["name_constraints"], d, config)
	if err != nil {
		return nil, err
	}
	transformed["nameConstraints"] = nameConstraints
	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSubject, err := expandPrivatecaCertificateConfigSubjectConfigSubject(original["subject"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubject); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subject"] = transformedSubject
	}

	transformedSubjectAltName, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltName(original["subject_alt_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSubjectAltName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["subjectAltName"] = transformedSubjectAltName
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCountryCode, err := expandPrivatecaCertificateConfigSubjectConfigSubjectCountryCode(original["country_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCountryCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["countryCode"] = transformedCountryCode
	}

	transformedOrganization, err := expandPrivatecaCertificateConfigSubjectConfigSubjectOrganization(original["organization"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganization); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organization"] = transformedOrganization
	}

	transformedOrganizationalUnit, err := expandPrivatecaCertificateConfigSubjectConfigSubjectOrganizationalUnit(original["organizational_unit"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOrganizationalUnit); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["organizationalUnit"] = transformedOrganizationalUnit
	}

	transformedLocality, err := expandPrivatecaCertificateConfigSubjectConfigSubjectLocality(original["locality"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocality); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locality"] = transformedLocality
	}

	transformedProvince, err := expandPrivatecaCertificateConfigSubjectConfigSubjectProvince(original["province"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvince); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["province"] = transformedProvince
	}

	transformedStreetAddress, err := expandPrivatecaCertificateConfigSubjectConfigSubjectStreetAddress(original["street_address"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStreetAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["streetAddress"] = transformedStreetAddress
	}

	transformedPostalCode, err := expandPrivatecaCertificateConfigSubjectConfigSubjectPostalCode(original["postal_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostalCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postalCode"] = transformedPostalCode
	}

	transformedCommonName, err := expandPrivatecaCertificateConfigSubjectConfigSubjectCommonName(original["common_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCommonName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["commonName"] = transformedCommonName
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectCountryCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectOrganization(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectOrganizationalUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectLocality(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectProvince(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectStreetAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectPostalCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectCommonName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDnsNames, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameDnsNames(original["dns_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dnsNames"] = transformedDnsNames
	}

	transformedUris, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedEmailAddresses, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameEmailAddresses(original["email_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEmailAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["emailAddresses"] = transformedEmailAddresses
	}

	transformedIpAddresses, err := expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameIpAddresses(original["ip_addresses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ipAddresses"] = transformedIpAddresses
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameDnsNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameEmailAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigSubjectConfigSubjectAltNameIpAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigPublicKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKey, err := expandPrivatecaCertificateConfigPublicKeyKey(original["key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["key"] = transformedKey
	}

	transformedFormat, err := expandPrivatecaCertificateConfigPublicKeyFormat(original["format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["format"] = transformedFormat
	}

	return transformed, nil
}

func expandPrivatecaCertificateConfigPublicKeyKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandPrivatecaCertificateConfigPublicKeyFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}