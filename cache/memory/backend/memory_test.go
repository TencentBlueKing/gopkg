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

package backend

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func newRandomDuration(seconds int) RandomExtraExpirationDurationFunc {
	return func() time.Duration {
		return time.Duration(rand.Intn(seconds*1000)) * time.Millisecond
	}
}

var _ = Describe("Memory", func() {

	It("newTTLCache", func() {
		c := newTTLCache(5*time.Second, 10*time.Second)
		assert.NotNil(GinkgoT(), c)

		c = newTTLCache(5*time.Second, 0*time.Second)
		assert.NotNil(GinkgoT(), c)
	})

	Describe("MemoryBackend", func() {
		var be *MemoryBackend
		BeforeEach(func() {
			be = NewMemoryBackend("test", 5*time.Second, newRandomDuration(5))
		})

		It("not nil", func() {
			assert.NotNil(GinkgoT(), be)
		})

		It("not exists", func() {
			_, found := be.Get("not_exists")
			assert.False(GinkgoT(), found)
		})

		It("set/get/delete", func() {
			be.Set("hello", "world", time.Duration(0))
			value, found := be.Get("hello")
			assert.True(GinkgoT(), found)
			assert.Equal(GinkgoT(), "world", value)

			_ = be.Delete("hello")
			_, found = be.Get("hello")
			assert.False(GinkgoT(), found)
		})

	})

})
