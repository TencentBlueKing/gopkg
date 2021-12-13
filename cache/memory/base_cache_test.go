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
	"errors"
	"time"

	"github.com/TencentBlueKing/gopkg/cache/memory/backend"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

func retrieveTest(k Key) (interface{}, error) {
	kStr := k.Key()
	switch kStr {
	case "a":
		return "1", nil
	case "b":
		return "2", nil
	case "error":
		return nil, errors.New("error")
	case "bool":
		return true, nil
	case "int64":
		return int64(1), nil
	case "time":
		return time.Time{}, nil
	default:
		return "", nil
	}
}

func retrieveError(k Key) (interface{}, error) {
	return nil, errors.New("test error")
}

var _ = Describe("BaseCache", func() {
	var c Cache
	var be *backend.MemoryBackend
	BeforeEach(func() {
		expiration := 5 * time.Minute
		be = backend.NewMemoryBackend("test", expiration, nil)

		c = NewBaseCache(false, retrieveTest, be)
	})

	It("Disabled", func() {
		assert.False(GinkgoT(), c.Disabled())
	})

	It("Get", func() {
		aKey := NewStringKey("a")
		x, err := c.Get(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))

		x, err = c.Get(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))
	})

	It("Disabled then get", func() {
		c = NewBaseCache(true, retrieveTest, be)

		aKey := NewStringKey("a")
		x, err := c.Get(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))

		x, err = c.Get(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))
	})

	It("Exists", func() {
		aKey := NewStringKey("a")
		// missing before first Get->Retrieve
		assert.False(GinkgoT(), c.Exists(aKey))

		c.Get(aKey)
		assert.True(GinkgoT(), c.Exists(aKey))
	})

	It("DirectGet", func() {
		aKey := NewStringKey("a")
		_, ok := c.DirectGet(aKey)
		assert.False(GinkgoT(), ok)

		c.Get(aKey)
		_, ok = c.DirectGet(aKey)
		assert.True(GinkgoT(), ok)
	})

	It("Delete", func() {
		aKey := NewStringKey("a")
		c.Get(aKey)
		assert.True(GinkgoT(), c.Exists(aKey))

		c.Delete(aKey)
		assert.False(GinkgoT(), c.Exists(aKey))
	})

	It("GetString", func() {
		aKey := NewStringKey("a")
		x, err := c.GetString(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x)

		int64Key := NewStringKey("int64")
		_, err = c.GetString(int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetString(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetBool", func() {
		aKey := NewStringKey("bool")
		x, err := c.GetBool(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), true, x)

		int64Key := NewStringKey("int64")
		_, err = c.GetBool(int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetBool(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetInt64", func() {
		aKey := NewStringKey("int64")
		x, err := c.GetInt64(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int64(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetInt64(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetInt64(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetTime", func() {
		aKey := NewStringKey("time")
		x, err := c.GetTime(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), time.Time{}, x)

		int64Key := NewStringKey("int64")
		_, err = c.GetTime(int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetTime(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetError", func() {
		aKey := NewStringKey("error")
		x, err := c.Get(aKey)
		assert.Error(GinkgoT(), err)
		assert.Nil(GinkgoT(), x)

		x, err2 := c.Get(aKey)
		assert.Error(GinkgoT(), err2)
		assert.Nil(GinkgoT(), x)

		// the error should be the same
		assert.Equal(GinkgoT(), err, err2)
	})

	It("retrieveError", func() {
		c = NewBaseCache(true, retrieveError, be)
		assert.NotNil(GinkgoT(), c)
		aKey := NewStringKey("a")
		_, err := c.Get(aKey)
		assert.Error(GinkgoT(), err)
	})

})

// TODO: mock the backend first Get fail, second Get ok

// TODO: add emptyCache here
