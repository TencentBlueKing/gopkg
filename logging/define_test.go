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

package logging

import (
	"sync"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"logur.dev/logur"
)

var _ = Describe("Define", func() {
	var (
		globalLoggers, globalLoggerAlias *sync.Map
		mockLogger                       logur.NoopLogger
	)

	BeforeEach(func() {
		globalLoggers = loggers
		loggers = &sync.Map{}
		globalLoggers.Range(func(key, value interface{}) bool {
			loggers.Store(key, value)
			return true
		})

		globalLoggerAlias = loggerAlias
		loggerAlias = &sync.Map{}
		globalLoggerAlias.Range(func(key, value interface{}) bool {
			loggerAlias.Store(key, value)
			return true
		})
	})

	AfterEach(func() {
		loggers = globalLoggers
		loggerAlias = globalLoggerAlias
	})

	It("Define a logger", func() {
		SetLogger("test", mockLogger)

		assert.Equal(GinkgoT(), GetLogger("test"), mockLogger)
	})

	It("Define an alias", func() {
		SetLogger("test", mockLogger)
		SetAlias("test", "alias")

		assert.Equal(GinkgoT(), GetLogger("test"), GetLogger("alias"))
	})

	It("Get a non-existing logger", func() {
		assert.Equal(GinkgoT(), GetLogger(DefaultLoggerName), GetLogger("non-existing"))
	})

	It("Define aliases", func() {
		SetLogger("test", mockLogger)
		SetAlias("test", "aliasA", "aliasB")

		assert.Equal(GinkgoT(), GetLogger("test"), GetLogger("aliasA"))
		assert.Equal(GinkgoT(), GetLogger("test"), GetLogger("aliasB"))
	})
})
