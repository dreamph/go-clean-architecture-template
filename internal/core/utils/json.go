package utils

import (
	"backend/internal/core/json"

	"github.com/pkg/errors"
)

func ToJson(val interface{}) string {
	jsonBytes := ToJsonBytes(val)
	if jsonBytes == nil {
		return ""
	}
	return string(jsonBytes)
}

func ToJsonBytes(val interface{}) []byte {
	if val == nil {
		return nil
	}
	result, err := json.Marshal(val)
	if err != nil {
		return nil
	}
	return result
}

func JsonToObj(jsonBytes []byte, objPointer interface{}) error {
	err := json.Unmarshal(jsonBytes, objPointer)
	if err != nil {
		return err
	}
	return nil
}

func ObjToJson(val interface{}) (string, error) {
	dataBytes, err := ObjToJsonBytes(val)
	if err != nil {
		return "", err
	}
	return string(dataBytes), nil
}

func ObjToJsonBytes(val interface{}) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	result, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ObjToMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func MapToObj(m map[string]interface{}, objPointer interface{}) error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonBytes, objPointer)
	if err != nil {
		return err
	}
	return nil
}

func ChangeValueInJson(jsonBytes []byte, onValue func(key string, val interface{}) (interface{}, error)) ([]byte, error) {
	var obj map[string]interface{}
	err := JsonToObj(jsonBytes, &obj)
	if err != nil {
		return nil, errors.Cause(err)
	}

	newData, err := CopyMap(obj, onValue)
	if err != nil {
		return nil, errors.Cause(err)
	}

	newBytes, err := json.Marshal(newData)
	if err != nil {
		return nil, errors.Cause(err)
	}
	return newBytes, nil
}

func CopyMap(m map[string]interface{}, onValue func(key string, val interface{}) (interface{}, error)) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range m {
		mapValue, isMap := v.(map[string]interface{})
		if isMap {
			val, err := onValue(k, mapValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			val, err = CopyMap(val.(map[string]interface{}), onValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			result[k] = val
			continue
		}

		sliceValue, isSlice := v.([]interface{})
		if isSlice {
			val, err := onValue(k, sliceValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			val, err = CopySlice(val.([]interface{}), onValue)
			if err != nil {
				return nil, errors.Cause(err)
			}

			result[k] = val
			continue
		}

		val, err := onValue(k, v)
		if err != nil {
			return nil, errors.Cause(err)
		}
		result[k] = val
	}

	return result, nil
}

func CopySlice(s []interface{}, onValue func(key string, val interface{}) (interface{}, error)) ([]interface{}, error) {
	var result []interface{}

	for _, v := range s {
		// Handle maps
		mapValue, isMap := v.(map[string]interface{})
		if isMap {
			val, err := onValue("", mapValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			val, err = CopyMap(val.(map[string]interface{}), onValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			result = append(result, val)
			continue
		}

		// Handle slices
		sliceValue, isSlice := v.([]interface{})
		if isSlice {
			val, err := onValue("", sliceValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			val, err = CopySlice(val.([]interface{}), onValue)
			if err != nil {
				return nil, errors.Cause(err)
			}
			result = append(result, val)
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
