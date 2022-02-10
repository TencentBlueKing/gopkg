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

package stringx_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/stringx"
)

var _ = Describe("Stringx", func() {
	Describe("Truncate", func() {
		var s = "helloworld"

		DescribeTable("Truncate String cases", func(expected string, truncatedSize int) {
			assert.Equal(GinkgoT(), expected, stringx.Truncate(s, truncatedSize))
		},
			Entry("truncated size less than real size", "he", 2),
			Entry("truncated size equals to real size", s, 10),
			Entry("truncated size greater than real size", s, 20),
		)
	})

	Describe("Random String", func() {
		DescribeTable("Random Sample cases", func(sequence string, length int) {
			s := stringx.RandomSample(sequence, length)
			assert.Equal(GinkgoT(), length, len(s))
			for _, c := range s {
				assert.True(GinkgoT(), strings.ContainsRune(sequence, c))
			}
		},
			Entry("Lowercase string 0", stringx.Lowercase, 0),
			Entry("Lowercase string 10", stringx.Lowercase, 10),

			Entry("Uppercase string 0", stringx.Uppercase, 0),
			Entry("Uppercase string 10", stringx.Uppercase, 10),

			Entry("Letters string 0", stringx.Letters, 0),
			Entry("Letters string 10", stringx.Letters, 10),

			Entry("Digits string 0", stringx.Digits, 0),
			Entry("Digits string 10", stringx.Digits, 10),

			Entry("Alphanum string 0", stringx.Alphanum, 0),
			Entry("Alphanum string 10", stringx.Alphanum, 10),
		)
	})

	Describe("Random Alphanumeric String", func() {
		DescribeTable("Random Alphanumeric String cases", func(length int) {
			s := stringx.RandomAlphanum(length)
			assert.Equal(GinkgoT(), length, len(s))
			for _, c := range s {
				assert.True(GinkgoT(), strings.ContainsRune(stringx.Alphanum, c))
			}
		},
			Entry("string length 0", 0),
			Entry("string length 1", 10),
			Entry("string length 10", 10),
		)
	})
})
