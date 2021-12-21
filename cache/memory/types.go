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

package memory

import (
	"time"
)

// Key is the type for the key of a cache entry.
// a struct-like object implements the key interface, so it can be used as a key in a cache.
type Key interface {
	Key() string
}

// RetrieveFunc is the type of the retrieve function.
// it retrieves the value from database, redis, apis, etc.
type RetrieveFunc func(key Key) (interface{}, error)

// Cache is the interface for the cache.
type Cache interface {
	Get(key Key) (interface{}, error)

	GetString(key Key) (string, error)
	GetBool(key Key) (bool, error)
	GetInt(key Key) (int, error)
	GetInt8(key Key) (int8, error)
	GetInt16(key Key) (int16, error)
	GetInt32(key Key) (int32, error)
	GetInt64(key Key) (int64, error)
	GetUint(key Key) (uint, error)
	GetUint8(key Key) (uint8, error)
	GetUint16(key Key) (uint16, error)
	GetUint32(key Key) (uint32, error)
	GetUint64(key Key) (uint64, error)
	GetFloat32(key Key) (float32, error)
	GetFloat64(key Key) (float64, error)
	GetTime(key Key) (time.Time, error)

	Delete(key Key) error
	Exists(key Key) bool

	DirectGet(key Key) (interface{}, bool)

	Disabled() bool
}
