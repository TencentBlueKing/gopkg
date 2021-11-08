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

package stringx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
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

		Describe("Random", func() {
			DescribeTable("Random String cases", func(length int) {
				assert.Equal(GinkgoT(), length, len(stringx.Random(length)))
			},
				Entry("string length 0", 0),
				Entry("string length 1", 10),
				Entry("string length 10", 10),
			)
		})
	})


})
