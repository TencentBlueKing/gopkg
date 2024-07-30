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
	"github.com/TencentBlueKing/gopkg/cache/memory/backend"
)

func retrieveTest(ctx context.Context, k cache.Key) (interface{}, error) {
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

func retrieveError(ctx context.Context, k cache.Key) (interface{}, error) {
	return nil, errors.New("test error")
}

var _ = Describe("BaseCache", func() {
	var c Cache
	var be *backend.MemoryBackend
	var ctx context.Context
	BeforeEach(func() {
		expiration := 5 * time.Minute
		be = backend.NewMemoryBackend("test", expiration, nil)

		c = NewBaseCache(retrieveTest, be, WithEmptyCache(0))
		ctx = context.Background()
	})

	It("Disabled", func() {
		assert.False(GinkgoT(), c.Disabled())
	})

	It("Get", func() {
		aKey := cache.NewStringKey("a")
		x, err := c.Get(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))

		x, err = c.Get(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))
	})

	It("Set", func() {
		setKey := cache.NewStringKey("s")
		c.Set(ctx, setKey, "1")
		x, err := c.GetString(ctx, setKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x)
	})

	It("Disabled then get", func() {
		c = NewBaseCache(retrieveTest, be, WithNoCache())

		aKey := cache.NewStringKey("a")
		x, err := c.Get(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))

		x, err = c.Get(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x.(string))
	})

	It("Exists", func() {
		aKey := cache.NewStringKey("a")
		// missing before first Get->Retrieve
		assert.False(GinkgoT(), c.Exists(ctx, aKey))

		_, _ = c.Get(ctx, aKey)
		assert.True(GinkgoT(), c.Exists(ctx, aKey))
	})

	It("DirectGet", func() {
		aKey := cache.NewStringKey("a")
		_, ok := c.DirectGet(ctx, aKey)
		assert.False(GinkgoT(), ok)

		_, _ = c.Get(ctx, aKey)
		_, ok = c.DirectGet(ctx, aKey)
		assert.True(GinkgoT(), ok)
	})

	It("Delete", func() {
		aKey := cache.NewStringKey("a")
		_, _ = c.Get(ctx, aKey)
		assert.True(GinkgoT(), c.Exists(ctx, aKey))

		_ = c.Delete(ctx, aKey)
		assert.False(GinkgoT(), c.Exists(ctx, aKey))
	})

	It("GetString", func() {
		aKey := cache.NewStringKey("a")
		x, err := c.GetString(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), "1", x)

		int64Key := cache.NewStringKey("int64")
		_, err = c.GetString(ctx, int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetString(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetBool", func() {
		aKey := cache.NewStringKey("bool")
		x, err := c.GetBool(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), true, x)

		int64Key := cache.NewStringKey("int64")
		_, err = c.GetBool(ctx, int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetBool(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetInt", func() {
		aKey := cache.NewStringKey("int")
		x, err := c.GetInt(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int(1), x)

		boolKey := cache.NewStringKey("bool")
		_, err = c.GetInt(ctx, boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetInt(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetInt8", func() {
		aKey := cache.NewStringKey("int8")
		x, err := c.GetInt8(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int8(1), x)

		boolKey := cache.NewStringKey("bool")
		_, err = c.GetInt8(ctx, boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetInt8(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetInt16", func() {
		aKey := cache.NewStringKey("int16")
		x, err := c.GetInt16(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int16(1), x)

		boolKey := cache.NewStringKey("bool")
		_, err = c.GetInt16(ctx, boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetInt16(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetInt32", func() {
		aKey := cache.NewStringKey("int32")
		x, err := c.GetInt32(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int32(1), x)

		boolKey := cache.NewStringKey("bool")
		_, err = c.GetInt32(ctx, boolKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetInt32(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetInt64", func() {
		aKey := cache.NewStringKey("int64")
		x, err := c.GetInt64(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), int64(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetInt64(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetInt64(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetUint", func() {
		aKey := cache.NewStringKey("uint")
		x, err := c.GetUint(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetUint(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetUint(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetUint8", func() {
		aKey := cache.NewStringKey("uint8")
		x, err := c.GetUint8(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint8(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetUint8(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetUint8(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint16", func() {
		aKey := cache.NewStringKey("uint16")
		x, err := c.GetUint16(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint16(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetUint16(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetUint16(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint32", func() {
		aKey := cache.NewStringKey("uint32")
		x, err := c.GetUint32(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint32(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetUint32(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetUint32(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetUint64", func() {
		aKey := cache.NewStringKey("uint64")
		x, err := c.GetUint64(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), uint64(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetUint64(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetUint64(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetFloat32", func() {
		aKey := cache.NewStringKey("float32")
		x, err := c.GetFloat32(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), float32(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetFloat32(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetFloat32(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})
	It("GetFloat64", func() {
		aKey := cache.NewStringKey("float64")
		x, err := c.GetFloat64(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), float64(1), x)

		strKey := cache.NewStringKey("a")
		_, err = c.GetFloat64(ctx, strKey)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetFloat64(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetTime", func() {
		aKey := cache.NewStringKey("time")
		x, err := c.GetTime(ctx, aKey)
		assert.NoError(GinkgoT(), err)
		assert.Equal(GinkgoT(), time.Time{}, x)

		int64Key := cache.NewStringKey("int64")
		_, err = c.GetTime(ctx, int64Key)
		assert.Error(GinkgoT(), err)

		errorKey := cache.NewStringKey("error")
		_, err = c.GetTime(ctx, errorKey)
		assert.Error(GinkgoT(), err)
	})

	It("GetError", func() {
		aKey := cache.NewStringKey("error")
		x, err := c.Get(ctx, aKey)
		assert.Error(GinkgoT(), err)
		assert.Nil(GinkgoT(), x)

		x, err2 := c.Get(ctx, aKey)
		assert.Error(GinkgoT(), err2)
		assert.Nil(GinkgoT(), x)

		// the error should be the same
		assert.Equal(GinkgoT(), err, err2)
	})

	It("retrieveError", func() {
		c = NewBaseCache(retrieveError, be, WithNoCache())
		assert.NotNil(GinkgoT(), c)
		aKey := cache.NewStringKey("a")
		_, err := c.Get(ctx, aKey)
		assert.Error(GinkgoT(), err)
	})
})

// TODO: mock the backend first Get fail, second Get ok

// TODO: add emptyCache here
