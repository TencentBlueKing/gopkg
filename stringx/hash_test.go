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
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/stringx"
)

var _ = Describe("Hash", func() {

	Describe("MD5Hash", func() {
		DescribeTable("MD5Hash cases", func(expected string, input string) {
			assert.Equal(GinkgoT(), expected, stringx.MD5Hash(input))
		},
			Entry("value is empty string", "d41d8cd98f00b204e9800998ecf8427e", ""),
			Entry("value is 'test'", "098f6bcd4621d373cade4e832627b4f6", "test"),
		)
	})

})
