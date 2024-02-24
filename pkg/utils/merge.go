package utils

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrTypeNotMatch = errors.New("type not match")

func convertSlice(i interface{}) []interface{} {
	ret := []interface{}{}

	switch v := i.(type) {
	case []interface{}:
		return v
	case []string:
		for _, v := range v {
			ret = append(ret, v)
		}
		return ret
	case []int:
		for _, v := range v {
			ret = append(ret, v)
		}
		return ret
	case []float64:
		for _, v := range v {
			ret = append(ret, v)
		}
		return ret
	case []float32:
		for _, v := range v {
			ret = append(ret, v)
		}
		return ret
	case []byte:
		return append(ret, v)
	}
	return nil
}

func Merge(src, dst interface{}) (interface{}, error) {
	srcType := reflect.TypeOf(src)
	dstType := reflect.TypeOf(dst)
	if srcType.Kind() != dstType.Kind() {
		return nil, ErrTypeNotMatch
	}

	switch srcType.Kind() {
	case reflect.Map:
		srcMap := src.(map[string]interface{})
		for k, dstVal := range dst.(map[string]interface{}) {
			srcVal, ok := srcMap[k]
			if !ok {
				srcMap[k] = dstVal
			} else {
				mergedVal, err := Merge(srcVal, dstVal)
				if err != nil {
					return nil, err
				}
				srcMap[k] = mergedVal
			}
		}
		return src, nil
	case reflect.Slice:
		srcSlice := convertSlice(src)
		dstSlice := convertSlice(dst)
		return append(srcSlice, dstSlice...), nil
	default:
		return src, nil
	}
}

func MergeMapString(a, b map[string]interface{}) (map[string]interface{}, error) {
	mr, err := Merge(a, b)
	if err != nil {
		return nil, err
	}

	if m, ok := mr.(map[string]interface{}); ok {
		return m, nil
	}

	return nil, fmt.Errorf("failed to merge map[string]interface{}: %v", mr)
}
