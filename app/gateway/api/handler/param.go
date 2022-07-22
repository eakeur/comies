package handler

import (
	"encoding/json"
	"reflect"
)

func parseQueryToValue(param reflect.Type, query map[string][]string) (reflect.Value, error) {

	normalized := map[string]interface{}{}
	for key, val := range query {
		if len(val) > 1 {
			normalized[key] = val
			continue
		}
		normalized[key] = val[0]
	}

	instance := reflect.New(param)
	ptr := instance.Interface()

	data, _ := json.Marshal(normalized)
	err := json.Unmarshal(data, &ptr)
	if err != nil {
		return reflect.Value{}, err
	}

	return instance.Elem(), nil
}
