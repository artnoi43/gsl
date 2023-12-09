package soyutils

import (
	"encoding/json"
	"os"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/soyart/gsl"
)

const testFile = "./test_assets/marshal_test"

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("marshal yaml to file", func(t *testing.T) {
		testMarshalAndWriteToFile(t, testFile+".yaml", yaml.Marshal)
	})

	t.Run("unmarshal yaml from file", func(t *testing.T) {
		testReadAndUnmarshalFile(t, testFile+".yaml", yaml.Unmarshal)
	})

	t.Run("marshal json to file", func(t *testing.T) {
		testMarshalAndWriteToFile(t, testFile+".json", json.Marshal)
	})

	t.Run("unmarshal json from file", func(t *testing.T) {
		testReadAndUnmarshalFile(t, testFile+".json", json.Unmarshal)
	})
}

func init() {
	soy, _ := getSoy()

	bytesJsonSoy, err := json.Marshal(soy)
	if err != nil {
		panic("failed to marshal json test file from soy")
	}

	bytesYamlSoy, err := yaml.Marshal(soy)
	if err != nil {
		panic("failed to marshal yaml test file from soy")
	}

	err = os.WriteFile(testFile+".json", bytesJsonSoy, os.ModePerm)
	if err != nil {
		panic("failed to write test file json")
	}

	err = os.WriteFile(testFile+".json", bytesYamlSoy, os.ModePerm)
	if err != nil {
		panic("failed to write test file yaml")
	}
}

type soy struct {
	S string `json:"S" yaml:"S"`
	I int    `json:"I" yaml:"I"`
}

func getSoy() (soy, map[string]interface{}) {
	return soy{
			S: "soytest",
			I: 70,
		},
		map[string]interface{}{
			"S": "soytest",
			"I": 70,
		}
}

func testMarshalAndWriteToFile(t *testing.T, filename string, f marshalFunc) {
	v, _ := getSoy()
	err := MarshalAndWriteToFile(v, f, filename)
	if err != nil {
		t.Error(err.Error())
	}
}

func testReadAndUnmarshalFile(t *testing.T, filename string, f unmarshalFunc) {
	// Test unmarshaling into type soy
	actualSoy, err := ReadAndUnmarshalFile[soy](filename, f)
	if err != nil {
		t.Error(err.Error())
	}

	expectedSoy, expectedMap := getSoy()
	ok, err := gsl.CompareInterfaceValues[soy](actualSoy, expectedSoy)
	if err != nil {
		t.Error(err.Error())
	}
	if !ok {
		t.Error("soy values not matched")
	}

	// Test unmarshaling into map[string]interface{}
	m, err := ReadAndUnmarshalFile[any](filename, f)
	if err != nil {
		t.Error(err.Error())
	}

	actual := m.(map[string]interface{})
	for k, expectedValue := range expectedMap {
		actualValue, ok := actual[k]
		if !ok {
			t.Errorf("actual value not found in key %s", k)
			continue
		}

		switch k {
		case "S":
			ok, err := gsl.CompareInterfaceValues[string](actualValue, expectedValue)
			if err != nil {
				t.Error(err.Error())
			}

			if !ok {
				t.Fatalf("values not matched for key %s\n", k)
			}

		case "I":
			ok, err := gsl.CompareInterfaceValues[int](actualValue, expectedValue)
			if err != nil {
				t.Error(err.Error())
			}

			if !ok {
				t.Fatalf("values not matched for key %s\n", k)
			}
		}
	}
}
