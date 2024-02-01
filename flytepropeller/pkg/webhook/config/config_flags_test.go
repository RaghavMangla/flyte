// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

var dereferencableKindsConfig = map[reflect.Kind]struct{}{
	reflect.Array: {}, reflect.Chan: {}, reflect.Map: {}, reflect.Ptr: {}, reflect.Slice: {},
}

// Checks if t is a kind that can be dereferenced to get its underlying type.
func canGetElementConfig(t reflect.Kind) bool {
	_, exists := dereferencableKindsConfig[t]
	return exists
}

// This decoder hook tests types for json unmarshaling capability. If implemented, it uses json unmarshal to build the
// object. Otherwise, it'll just pass on the original data.
func jsonUnmarshalerHookConfig(_, to reflect.Type, data interface{}) (interface{}, error) {
	unmarshalerType := reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()
	if to.Implements(unmarshalerType) || reflect.PtrTo(to).Implements(unmarshalerType) ||
		(canGetElementConfig(to.Kind()) && to.Elem().Implements(unmarshalerType)) {

		raw, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Failed to marshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		res := reflect.New(to).Interface()
		err = json.Unmarshal(raw, &res)
		if err != nil {
			fmt.Printf("Failed to umarshal Data: %v. Error: %v. Skipping jsonUnmarshalHook", data, err)
			return data, nil
		}

		return res, nil
	}

	return data, nil
}

func decode_Config(input, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			jsonUnmarshalerHookConfig,
		),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func join_Config(arr interface{}, sep string) string {
	listValue := reflect.ValueOf(arr)
	strs := make([]string, 0, listValue.Len())
	for i := 0; i < listValue.Len(); i++ {
		strs = append(strs, fmt.Sprintf("%v", listValue.Index(i)))
	}

	return strings.Join(strs, sep)
}

func testDecodeJson_Config(t *testing.T, val, result interface{}) {
	assert.NoError(t, decode_Config(val, result))
}

func testDecodeRaw_Config(t *testing.T, vStringSlice, result interface{}) {
	assert.NoError(t, decode_Config(vStringSlice, result))
}

func TestConfig_GetPFlagSet(t *testing.T) {
	val := Config{}
	cmdFlags := val.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())
}

func TestConfig_SetFlags(t *testing.T) {
	actual := Config{}
	cmdFlags := actual.GetPFlagSet("")
	assert.True(t, cmdFlags.HasFlags())

	t.Run("Test_metrics-prefix", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("metrics-prefix", testValue)
			if vString, err := cmdFlags.GetString("metrics-prefix"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.MetricsPrefix)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_certDir", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("certDir", testValue)
			if vString, err := cmdFlags.GetString("certDir"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.CertDir)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_localCert", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("localCert", testValue)
			if vBool, err := cmdFlags.GetBool("localCert"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.LocalCert)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_listenPort", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("listenPort", testValue)
			if vInt, err := cmdFlags.GetInt("listenPort"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt), &actual.ListenPort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_serviceName", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("serviceName", testValue)
			if vString, err := cmdFlags.GetString("serviceName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.ServiceName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_servicePort", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("servicePort", testValue)
			if vInt32, err := cmdFlags.GetInt32("servicePort"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vInt32), &actual.ServicePort)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_secretName", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("secretName", testValue)
			if vString, err := cmdFlags.GetString("secretName"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.SecretName)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_awsSecretManager.sidecarImage", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("awsSecretManager.sidecarImage", testValue)
			if vString, err := cmdFlags.GetString("awsSecretManager.sidecarImage"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.AWSSecretManagerConfig.SidecarImage)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_gcpSecretManager.sidecarImage", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("gcpSecretManager.sidecarImage", testValue)
			if vString, err := cmdFlags.GetString("gcpSecretManager.sidecarImage"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.GCPSecretManagerConfig.SidecarImage)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_vaultSecretManager.role", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("vaultSecretManager.role", testValue)
			if vString, err := cmdFlags.GetString("vaultSecretManager.role"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.VaultSecretManagerConfig.Role)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_embeddedSecretManagerConfig.enabled", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("embeddedSecretManagerConfig.enabled", testValue)
			if vBool, err := cmdFlags.GetBool("embeddedSecretManagerConfig.enabled"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vBool), &actual.EmbeddedSecretManagerConfig.Enabled)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_embeddedSecretManagerConfig.type", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("embeddedSecretManagerConfig.type", testValue)
			if vString, err := cmdFlags.GetString("embeddedSecretManagerConfig.type"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.EmbeddedSecretManagerConfig.Type)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_embeddedSecretManagerConfig.awsConfig.region", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("embeddedSecretManagerConfig.awsConfig.region", testValue)
			if vString, err := cmdFlags.GetString("embeddedSecretManagerConfig.awsConfig.region"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.EmbeddedSecretManagerConfig.AWSConfig.Region)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
	t.Run("Test_embeddedSecretManagerConfig.gcpConfig.project", func(t *testing.T) {

		t.Run("Override", func(t *testing.T) {
			testValue := "1"

			cmdFlags.Set("embeddedSecretManagerConfig.gcpConfig.project", testValue)
			if vString, err := cmdFlags.GetString("embeddedSecretManagerConfig.gcpConfig.project"); err == nil {
				testDecodeJson_Config(t, fmt.Sprintf("%v", vString), &actual.EmbeddedSecretManagerConfig.GCPConfig.Project)

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	})
}
