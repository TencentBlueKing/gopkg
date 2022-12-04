/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云 - gopkg available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package mapx

import (
	"errors"
	"fmt"
	"strings"
)

// GetItems 获取嵌套定义的 Map 值
// paths 参数支持 []string 类型，如 []string{"metadata", "namespace"}
// 或 string 类型（以 '.' 为分隔符），如 "spec.template.spec.containers"
func GetItems(obj map[string]interface{}, paths interface{}) (interface{}, error) {
	switch t := paths.(type) {
	case string:
		return getItems(obj, strings.Split(paths.(string), "."))
	case []string:
		return getItems(obj, paths.([]string))
	default:
		return nil, fmt.Errorf("paths's type must one of (string, []string), get %v", t)
	}
}

func getItems(obj map[string]interface{}, paths []string) (interface{}, error) {
	if len(paths) == 0 {
		return nil, errors.New("paths is empty list")
	}
	ret, exists := obj[paths[0]]
	if !exists {
		return nil, fmt.Errorf("key %s not exist", paths[0])
	}
	if len(paths) == 1 {
		return ret, nil
	} else if subMap, ok := obj[paths[0]].(map[string]interface{}); ok {
		return getItems(subMap, paths[1:])
	}
	return nil, fmt.Errorf("key %s, val not map[string]interface{} type", paths[0])
}

// Get 若指定值不存在，则返回默认值
func Get(obj map[string]interface{}, paths interface{}, defVal interface{}) interface{} {
	ret, err := GetItems(obj, paths)
	if err != nil {
		return defVal
	}
	return ret
}

// GetBool 获取 Bool 类型快捷方法，默认值为 false
func GetBool(obj map[string]interface{}, paths interface{}) bool {
	return Get(obj, paths, false).(bool)
}

// GetInt64 获取 int64 类型快捷方法，默认值为 int64(0)
func GetInt64(obj map[string]interface{}, paths interface{}) int64 {
	return Get(obj, paths, int64(0)).(int64)
}

// GetStr 获取 string 类型快捷方法，默认值为 ""
func GetStr(obj map[string]interface{}, paths interface{}) string {
	return Get(obj, paths, "").(string)
}

// GetList 获取 []interface{} 类型快捷方法，默认值为 []interface{}{}
func GetList(obj map[string]interface{}, paths interface{}) []interface{} {
	return Get(obj, paths, []interface{}{}).([]interface{})
}

// GetMap 获取 map[string]interface{} 类型快捷方法，默认值为 map[string]interface{}
func GetMap(obj map[string]interface{}, paths interface{}) map[string]interface{} {
	return Get(obj, paths, map[string]interface{}{}).(map[string]interface{})
}
