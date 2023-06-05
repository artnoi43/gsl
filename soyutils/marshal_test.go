package soyutils

import (
	"encoding/json"
	"testing"

	"github.com/soyart/gsl"
	"gopkg.in/yaml.v3"
)

func TestMarshalUnmarshal(t *testing.T) {
	t.Run("marshal yaml to file", func(t *testing.T) {
		testMarshalAndWriteToFile(t, yaml.Marshal)
	})
	t.Run("unmarshal yaml from file", func(t *testing.T) {
		testReadAndUnmarshalFile(t, yaml.Unmarshal)
	})
	t.Run("marshal json to file", func(t *testing.T) {
		testMarshalAndWriteToFile(t, json.Marshal)
	})
	t.Run("unmarshal json from file", func(t *testing.T) {
		testReadAndUnmarshalFile(t, json.Unmarshal)
	})
}

type soy struct {
	S string `json:"S" yaml:"S"`
	I int    `json:"I" yaml:"I"`
}

const testFile = "./assets/marshal_test"

func getSoy() (soy, map[string]interface{}) {
	return soy{S: "soytest", I: 69}, map[string]interface{}{
		"S": "soytest",
		"I": 69,
	}
}

func testMarshalAndWriteToFile(t *testing.T, f marshalFunc) {
	v, _ := getSoy()
	err := MarshalAndWriteToFile(v, f, testFile)
	if err != nil {
		t.Error(err.Error())
	}
}

func testReadAndUnmarshalFile(t *testing.T, f unmarshalFunc) {
	// Test unmarshaling into type soy
	actualSoy, err := ReadAndUnmarshalFile[soy](testFile, f)
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
	m, err := ReadAndUnmarshalFile[any](testFile, f)
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
				t.Errorf("values not matched for key %s\n", k)
			}
		case "I":
			ok, err := gsl.CompareInterfaceValues[int](actualValue, expectedValue)
			if err != nil {
				t.Error(err.Error())
			}
			if !ok {
				t.Errorf("values not matched for key %s\n", k)
			}
		}
	}
}
