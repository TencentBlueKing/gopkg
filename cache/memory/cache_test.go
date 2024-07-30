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
	"errors"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/cache"
)

var cacheError = errors.New("cache error")

func retrieveOK(ctx context.Context, k cache.Key) (interface{}, error) {
	return "", nil
}

func retrieveErr(ctx context.Context, k cache.Key) (interface{}, error) {
	return nil, cacheError
}

var _ = Describe("Cache", func() {

	var ctx context.Context

	It("New", func() {
		expiration := 5 * time.Minute

		c := NewCache("test", retrieveOK, expiration, nil)
		assert.NotNil(GinkgoT(), c)
	})

	It("NewMock", func() {
		c := NewMockCache(retrieveOK)
		assert.NotNil(GinkgoT(), c)
	})

	It("Cache Disable", func() {
		expiration := 5 * time.Minute
		c := NewCache("test", retrieveOK, expiration, nil, WithNoCache())
		assert.True(GinkgoT(), c.Disabled())
	})

	It("Cache WithEmptyCache", func() {
		aKey := cache.NewStringKey("test")
		expiration := 5 * time.Minute
		c := NewCache("test", retrieveErr, expiration, nil, WithEmptyCache(0))
		_, err := c.Get(ctx, aKey)
		assert.ErrorIs(GinkgoT(), err, cacheError)

	})

})
