/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云-gopkg available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package cache

import (
	"time"
)

type Key interface {
	Key() string
}

type RetrieveFunc func(key Key) (interface{}, error)

type Cache interface {
	Get(key Key) (interface{}, error)

	GetString(key Key) (string, error)
	GetBool(key Key) (bool, error)
	GetInt64(key Key) (int64, error)
	GetTime(key Key) (time.Time, error)

	Delete(key Key) error
	Exists(key Key) bool

	DirectGet(key Key) (interface{}, bool)

	Disabled() bool
}
