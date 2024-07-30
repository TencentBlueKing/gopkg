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
	"time"

	"github.com/TencentBlueKing/gopkg/cache/memory/backend"
)

// NewCache creates a new cache object.
// - name: the cache name.
// - retrieveFunc: the function to retrieve the real data.
// - expiration: the expiration time.
// - randomExtraExpirationFunc: the function to generate a random duration, used to add extra expiration for each key.
// - options: the options for the cache . eg:
//            WithNoCache:disable cache
//            WithEmptyCache(duration): set the key EmptyCache if retrieve fail from retrieveFunc
func NewCache(
	name string,
	retrieveFunc RetrieveFunc,
	expiration time.Duration,
	randomExtraExpirationFunc backend.RandomExtraExpirationDurationFunc,
	options ...Option,
) Cache {
	be := backend.NewMemoryBackend(name, expiration, randomExtraExpirationFunc)
	return NewBaseCache(retrieveFunc, be, options...)
}

// NewMockCache create a memory cache for mock
func NewMockCache(retrieveFunc RetrieveFunc) Cache {
	be := backend.NewMemoryBackend("mockCache", 5*time.Minute, nil)

	return NewBaseCache(retrieveFunc, be)
}
