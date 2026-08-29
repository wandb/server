package values

import (
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
	var intermediate map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &intermediate); err != nil {
		return nil, err
	}
	coverted := convertMapKeysToStrings(intermediate)
	return coverted, nil
}

func convertMapKeysToStrings(originalMap map[interface{}]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range originalMap {
		strKey, ok := key.(string)
		if !ok {
			continue // or handle the error as appropriate
		}
		if subMap, isSubMap := value.(map[interface{}]interface{}); isSubMap {
			value = convertMapKeysToStrings(subMap)
		}
		newMap[strKey] = value
	}
	return newMap
}
