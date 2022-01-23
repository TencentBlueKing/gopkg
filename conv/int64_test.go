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

package conv_test

import (
	"github.com/TencentBlueKing/gopkg/conv"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Int64", func() {
	DescribeTable("ToInt64 cases", func(value interface{}, expected int64, willError bool) {
		v, err := conv.ToInt64(value)
		if willError {
			assert.Error(GinkgoT(), err)
		} else {
			assert.NoError(GinkgoT(), err)
			assert.Equal(GinkgoT(), expected, v)
		}
	},
		Entry("value is positive int", 123, int64(123), false),
		Entry("value is negative int", -123, int64(-123), false),
		Entry("value is positive int64", int64(123), int64(123), false),
		Entry("value is negative int64", int64(-123), int64(-123), false),
		Entry("value is string numberic", "123", int64(123), false),
		Entry("value is string numberic", "-123", int64(-123), false),
		Entry("value is float", float64(123.45), int64(123), false),
		Entry("value is nil", nil, int64(0), false),
		Entry("value is int32", int32(123), int64(0), true),
		Entry("value is not numberic", "abc", int64(0), true),
	)

})
