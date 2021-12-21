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
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Key", func() {

	It("stringKey", func() {
		k := NewStringKey("hello")
		assert.NotNil(GinkgoT(), k)
		assert.Equal(GinkgoT(), "hello", k.Key())
	})

	It("intKey", func() {
		k := NewIntKey(123)
		assert.NotNil(GinkgoT(), k)
		assert.Equal(GinkgoT(), "123", k.Key())
	})

	It("int64Key", func() {
		k := NewInt64Key(int64(123))
		assert.NotNil(GinkgoT(), k)
		assert.Equal(GinkgoT(), "123", k.Key())
	})

	It("uintKey", func() {
		k := NewUintKey(uint(123))
		assert.NotNil(GinkgoT(), k)
		assert.Equal(GinkgoT(), "123", k.Key())
	})

	It("uint64Key", func() {
		k := NewUint64Key(uint64(123))
		assert.NotNil(GinkgoT(), k)
		assert.Equal(GinkgoT(), "123", k.Key())
	})

})
