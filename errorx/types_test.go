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

package errorx

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

type NoIsWrapError struct {
	message string
	err     error
}

func (e NoIsWrapError) Error() string {
	return e.message
}

var _ = Describe("Types", func() {
	var e1 error
	var e2 Errorx
	var e3 Errorx
	BeforeEach(func() {
		e1 = errors.New("a")

		e2 = Errorx{
			message: "e2",
			err:     e1,
		}

		e3 = Errorx{
			message: "e3",
			err:     e2,
		}
	})

	Describe("Is", func() {
		It("target is nil", func() {
			assert.False(GinkgoT(), errors.Is(e2, nil))
		})

		It("the e.error is nil", func() {
			e := Errorx{
				message: "e",
				err:     nil,
			}
			assert.False(GinkgoT(), errors.Is(e, errors.New("an error")))
		})

		It("ok", func() {
			assert.True(GinkgoT(), errors.Is(e2, e1))
			assert.True(GinkgoT(), errors.Is(e3, e1))
			assert.True(GinkgoT(), errors.Is(e3, e2))

			// false
			assert.False(GinkgoT(), errors.Is(e1, e2))
			assert.False(GinkgoT(), errors.Is(e1, e3))
			assert.False(GinkgoT(), errors.Is(e2, e3))
		})

		It("noIsWrapError", func() {
			e4 := NoIsWrapError{
				message: "no_is_wrap",
				err:     e1,
			}
			e5 := Errorx{
				message: "e5",
				err:     e4,
			}

			assert.True(GinkgoT(), errors.Is(e5, e4))
			assert.False(GinkgoT(), errors.Is(e4, e5))
		})
	})

	Describe("Unwrap", func() {
		assert.Equal(GinkgoT(), e1, e2.Unwrap())
		assert.Equal(GinkgoT(), e1, e3.Unwrap())

		assert.Equal(GinkgoT(), e1, errors.Unwrap(e2))
		assert.Equal(GinkgoT(), e1, errors.Unwrap(e3))
	})

})
