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

package conv_test

import (
	"errors"
	"math"
	"testing"

	"github.com/TencentBlueKing/gopkg/conv"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("String", func() {
	It("string to bytes", func() {
		b := conv.StringToBytes("abc")
		assert.Equal(GinkgoT(), []byte("abc"), b)
	})

	It("bytes to string", func() {
		s := conv.BytesToString([]byte("abc"))
		assert.Equal(GinkgoT(), "abc", s)
	})

	Describe("ToString", func() {
		DescribeTable("ToString cases", func(expected string, value interface{}) {
			v := conv.ToString(value)
			assert.Equal(GinkgoT(), expected, v)
		},
			Entry("", "", nil),
			Entry("", "", ""),
			Entry("", "foo", "foo"),
			Entry("", "true", true),
			Entry("", "42", 42),
			Entry("", "3.14", 3.14),
			Entry("", "-127", -127),
			Entry("", "255", 0xFF),
			Entry("", "", []byte{}),
			Entry("", "abc", errors.New("abc")),
			Entry("", "42", int(42)),
			Entry("", "42", int8(42)),
			Entry("", "42", int16(42)),
			Entry("", "42", int32(42)),
			Entry("", "42", int64(42)),
			Entry("", "42", uint(42)),
			Entry("", "42", uint8(42)),
			Entry("", "42", uint16(42)),
			Entry("", "42", uint32(42)),
			Entry("", "42", uint64(42)),
			Entry("", "3.141592653589793", math.Pi),
			Entry("", "NaN", math.NaN()),
			Entry("", "+Inf", math.Inf(1)),
			Entry("", "-Inf", math.Inf(-1)),
		)
	})
})

func BenchmarkBytesToString(b *testing.B) {
	bs := []byte("hello world")
	for i := 0; i < b.N; i++ {
		conv.BytesToString(bs)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	s := "hello world"
	for i := 0; i < b.N; i++ {
		conv.StringToBytes(s)
	}
}
