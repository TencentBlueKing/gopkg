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
	case "int":
		return 1, nil
	case "int8":
		return int8(1), nil
	case "int16":
		return int16(1), nil
	case "int32":
		return int32(1), nil
	case "int64":
		return int64(1), nil
	case "uint":
		return uint(1), nil
	case "uint8":
		return uint8(1), nil
	case "uint16":
		return uint16(1), nil
	case "uint32":
		return uint32(1), nil
	case "uint64":
		return uint64(1), nil
	case "float32":
		return float32(1), nil
	case "float64":
		return float64(1), nil
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

		_, _ = c.Get(aKey)
		assert.True(GinkgoT(), c.Exists(aKey))
	})

	It("DirectGet", func() {
		aKey := NewStringKey("a")
		_, ok := c.DirectGet(aKey)
		assert.False(GinkgoT(), ok)

		_, _ = c.Get(aKey)
		_, ok = c.DirectGet(aKey)
		assert.True(GinkgoT(), ok)
	})

	It("Delete", func() {
		aKey := NewStringKey("a")
		_, _ = c.Get(aKey)
		assert.True(GinkgoT(), c.Exists(aKey))

		_ = c.Delete(aKey)
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

	It("GetInt", func() {
		aKey := NewStringKey("int")
		x, err := c.GetInt(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int(1), x)

		boolKey := NewStringKey("bool")
		_, err = c.GetInt(boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetInt(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetInt8", func() {
		aKey := NewStringKey("int8")
		x, err := c.GetInt8(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int8(1), x)

		boolKey := NewStringKey("bool")
		_, err = c.GetInt8(boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetInt8(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetInt16", func() {
		aKey := NewStringKey("int16")
		x, err := c.GetInt16(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int16(1), x)

		boolKey := NewStringKey("bool")
		_, err = c.GetInt16(boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetInt16(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetInt32", func() {
		aKey := NewStringKey("int32")
		x, err := c.GetInt32(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int32(1), x)

		boolKey := NewStringKey("bool")
		_, err = c.GetInt32(boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetInt32(errorKey)
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

	It("GetUint", func() {
		aKey := NewStringKey("uint")
		x, err := c.GetUint(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetUint(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetUint(errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetUint8", func() {
		aKey := NewStringKey("uint8")
		x, err := c.GetUint8(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint8(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetUint8(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetUint8(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint16", func() {
		aKey := NewStringKey("uint16")
		x, err := c.GetUint16(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint16(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetUint16(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetUint16(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint32", func() {
		aKey := NewStringKey("uint32")
		x, err := c.GetUint32(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint32(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetUint32(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetUint32(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint64", func() {
		aKey := NewStringKey("uint64")
		x, err := c.GetUint64(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint64(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetUint64(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetUint64(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetFloat32", func() {
		aKey := NewStringKey("float32")
		x, err := c.GetFloat32(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), float32(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetFloat32(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetFloat32(errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetFloat64", func() {
		aKey := NewStringKey("float64")
		x, err := c.GetFloat64(aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), float64(1), x)

		strKey := NewStringKey("a")
		_, err = c.GetFloat64(strKey)
		assert.Error(GinkgoT(), err)

		errorKey := NewStringKey("error")
		_, err = c.GetFloat64(errorKey)
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
