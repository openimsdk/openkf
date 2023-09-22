// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"encoding/json"
	"errors"
	"reflect"
)

func StructToMapWithJSONTag(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, errors.New("data parameter must be a struct or a pointer to struct")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func StructToMapString(data interface{}) (map[string]string, error) {
	result := make(map[string]string)

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, errors.New("data parameter must be a struct or a pointer to struct")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FlattenMap(input map[string]interface{}, infix string) map[string]interface{} {
	flattened := make(map[string]interface{})
	for k, v := range input {
		switch v := v.(type) {
		case map[string]interface{}:
			dfsMap(flattened, v, k, infix)
		default:
			flattened[k] = v
		}
	}
	return flattened
}

func dfsMap(flattened, data map[string]interface{}, path string, infix string) {
	for k, v := range data {
		newPath := path + infix + k
		switch v := v.(type) {
		case map[string]interface{}:
			dfsMap(flattened, v, newPath, infix)
		default:
			flattened[newPath] = v
		}
	}
}
