package values

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/wandb/server/pkg/utils"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/strvals"
)

// Values holds an arbitrary tree-like data structure, such as parsed JSON or
// YAML. It can be used to represent Helm-like values.
//
// You can use convenient accessors to work with this data structure.
//
// Note: Some use-cases, including Helm SDK, make assumptions about the data
// type of the embedded lists and objects. To avoid any error or an unexpected
// behavior use `[]interface{}` for lists and `map[string]interface{}` for
// nested objects.
type Values map[string]interface{}

// AsMap returns the underlying map. This maybe be useful for external libraries.
func (v Values) AsMap() map[string]interface{} {
	return v
}

// GetValue retrieves the value that is addressed with a dot-separated key, e.g.
// `x.y.z`. The key format does not support array indexing.
//
// It will return an error, if the key does not exist or can not be traversed,
// for example nested keys of a leaf node of the tree.
func (v Values) GetValue(key string) (interface{}, error) {
	cursor := v
	keyElements := []string{}

	if key != "" {
		keyElements = strings.Split(key, ".")
	}

	var target interface{} = cursor
	for idx, elm := range keyElements {
		target = cursor[elm]

		if target == nil {
			if idx < len(keyElements)-1 {
				return nil, errors.Errorf("missing element at %s for %s key", elm, key)
			}
		}

		if targetAsMap, ok := target.(map[string]interface{}); ok {
			cursor = targetAsMap
		} else {
			if idx < len(keyElements)-1 {
				return nil, errors.Errorf("leaf element at %s for %s key", elm, key)
			}
		}
	}

	return target, nil
}

// GetString retrieves the value of type `string` that is addressed with a
// dot-separated key, e.g. `x.y.z`. The key format does not support array
// indexing.
//
// If the key does not exist it will return the optional `defaultValue` or an
// empty string. When the key exists but its type is not `string` it tries to
// format it as `string`. If it fails it will return either the `defaultValue`
// or an empty string.
//
// Use `HasKey` to check if the `key` exist.
func (v Values) GetString(key string, defaultValue ...string) string {
	defaultValueToUse := ""
	if len(defaultValue) > 0 {
		defaultValueToUse = defaultValue[0]
	}

	if val, err := v.GetValue(key); err == nil {
		if strVal, isStr := val.(string); isStr {
			return strVal
		}

		if val != nil {
			/* You are at the mercy of fmt.Sprint */
			return fmt.Sprint(val)
		}
	}

	return defaultValueToUse
}

// GetBool retrieves the value of type `bool` that is addressed with a
// dot-separated key, e.g. `x.y.z`. The key format does not support array
// indexing.
//
// If the key does not exist or its type is not `bool` it will return the
// optional `defaultValue` or `false`.
//
// Use `HasKey` to check if the `key` exist.
func (v Values) GetBool(key string, defaultValue ...bool) bool {
	defaultValueToUse := false
	if len(defaultValue) > 0 {
		defaultValueToUse = defaultValue[0]
	}

	if val, err := v.GetValue(key); err == nil {
		if boolVal, isBool := val.(bool); isBool {
			return boolVal
		}
	}

	return defaultValueToUse
}

// HasKey checks if a specific key exists. The `key` is in dot-separated format
// and it does not support array indexing. When it returns `true` the getter
// methods, e.g. `GetValue`, can successfully retrieve the associated value of
// the key.
func (v Values) HasKey(key string) bool {
	if _, err := v.GetValue(key); err == nil {
		return true
	}

	return false
}

// SetValue set the value that is addressed with a dot-separated key, e.g.
// `x.y.z`. The key format does not support array indexing.
//
// It does not do any type checking and will override the targeted element with
// the provided value. It will create any intermediate nested object that
// doesn't exist. But it will return an error when the key can not be traversed,
// for example nested keys of a leaf node of the tree.
func (v Values) SetValue(key string, value interface{}) error {
	if key == "" {
		return errors.Errorf("can not set the root element")
	}

	cursor := v
	keyElements := strings.Split(key, ".")

	var target interface{} = cursor
	for idx, elm := range keyElements {
		target = cursor[elm]

		if target == nil {
			if idx < len(keyElements)-1 {
				target = map[string]interface{}{}
				cursor[elm] = target
			}
		}

		if targetAsMap, ok := target.(map[string]interface{}); ok {
			cursor = targetAsMap
		} else {
			if idx < len(keyElements)-1 {
				return errors.Errorf("leaf element at %s for %s key", elm, key)
			}
		}
	}

	cursor[keyElements[len(keyElements)-1]] = value

	return nil
}

// Merge uses deep copy to merge the content of another data structure. It
// overrides the existing keys with their associated new values and copies the
// missing keys.
func (v Values) Merge(newValues Values) (Values, error) {
	return utils.MergeMapString(v.AsMap(), newValues.AsMap())
}

// Coalesce uses Helm-style coalesce to merge the content of another data
// structure.
//
// Coalesce can traverse nested values. As opposed to merge, it only inserts the
// missing key-value pairs and does not override the existing keys.
//
// It uses Helm SDK coalesce function and does not return an error when fails
// to merge specific entries. Therefore it may lead to an incorrect output.
func (v Values) Coalesce(newValues Values) {
	chartutil.CoalesceTables(v, newValues)
}

// AddHelmValue sets the specified value using Helm style key and value format.
// This is similar to `--set key=value` command argument of Helm. It does not
// support multiple keys and values. It returns an error if it can not parse
// the key or assign the value.
func (v Values) AddHelmValue(key, value string) error {
	if err := strvals.ParseInto(fmt.Sprintf("%s=%s", key, value), v); err != nil {
		return errors.Wrapf(err, "failed to set value: %s=%s", key, value)
	}

	return nil
}
