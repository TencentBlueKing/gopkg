/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云-gopkg available.
 * Copyright (C) 2017-2022 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package memory

import (
	"context"
	"time"

	"github.com/TencentBlueKing/gopkg/cache"
)

// RetrieveFunc is the type of the retrieve function.
// it retrieves the value from database, redis, apis, etc.
type RetrieveFunc func(ctx context.Context, key cache.Key) (interface{}, error)

// Cache is the interface for the cache.
type Cache interface {
	Get(ctx context.Context, key cache.Key) (interface{}, error)
	Set(ctx context.Context, key cache.Key, data interface{})

	GetString(ctx context.Context, key cache.Key) (string, error)
	GetBool(ctx context.Context, key cache.Key) (bool, error)
	GetInt(ctx context.Context, key cache.Key) (int, error)
	GetInt8(ctx context.Context, key cache.Key) (int8, error)
	GetInt16(ctx context.Context, key cache.Key) (int16, error)
	GetInt32(ctx context.Context, key cache.Key) (int32, error)
	GetInt64(ctx context.Context, key cache.Key) (int64, error)
	GetUint(ctx context.Context, key cache.Key) (uint, error)
	GetUint8(ctx context.Context, key cache.Key) (uint8, error)
	GetUint16(ctx context.Context, key cache.Key) (uint16, error)
	GetUint32(ctx context.Context, key cache.Key) (uint32, error)
	GetUint64(ctx context.Context, key cache.Key) (uint64, error)
	GetFloat32(ctx context.Context, key cache.Key) (float32, error)
	GetFloat64(ctx context.Context, key cache.Key) (float64, error)
	GetTime(ctx context.Context, key cache.Key) (time.Time, error)

	Delete(ctx context.Context, key cache.Key) error
	Exists(ctx context.Context, key cache.Key) bool

	DirectGet(ctx context.Context, key cache.Key) (interface{}, bool)

	Disabled() bool
}
