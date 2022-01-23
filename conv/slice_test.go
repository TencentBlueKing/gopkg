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

var _ = Describe("Slice", func() {

	Describe("ToSlice", func() {

		intSlice := []int{1}
		strSlice := []string{"abc"}

		DescribeTable("ToSlice cases", func(expected int, willError bool, input interface{}) {
			data, err := conv.ToSlice(input)

			if willError {
				assert.Error(GinkgoT(), err)
			} else {
				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), expected, len(data))
			}
		},
			Entry("not a slice", 0, true, ""),
			Entry("a []int{1}", 1, false, intSlice),
			Entry("a []string{abc}", 1, false, strSlice),
		)
	})

})
