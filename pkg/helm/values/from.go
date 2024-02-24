package values

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func FromYAMLFile(path string) (Values, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromYAML(out)
}

func FromYAML(data []byte) (Values, error) {
	vals := new(Values)
	if err := yaml.Unmarshal(data, &vals); err != nil {
		return nil, err
	}
	return *vals, nil
}

func FromJSON(data []byte) (Values, error) {
	vals := new(Values)
	if err := json.Unmarshal(data, &vals); err != nil {
		return nil, err
	}
	return *vals, nil
}
