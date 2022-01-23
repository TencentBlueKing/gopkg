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

package log

import (
	"log"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"logur.dev/logur"
)

var _ = Describe("Logger", func() {
	var (
		output strings.Builder
		logger *log.Logger
	)

	BeforeEach(func() {
		output.Reset()
		logger = log.New(&output, "", 0)
	})

	Context("Log", func() {
		It("should log a message", func() {
			l := New(logger, logur.Info)
			l.log(logur.Info, "TEST", "log")
			Expect(output.String()).To(Equal("TEST log\n"))
		})

		It("should not log a message", func() {
			l := New(logger, logur.Error)
			l.log(logur.Info, "TEST", "log")
			Expect(output.String()).To(Equal(""))
		})

		It("should log a message with fields", func() {
			l := New(logger, logur.Info)
			l.log(logur.Info, "TEST", "log", map[string]interface{}{
				"value": 123.456,
			})
			Expect(output.String()).To(Equal("TEST log value=123.456\n"))
		})
	})

	Context("Log by level", func() {
		var adaptedLogger *Logger

		BeforeEach(func() {
			adaptedLogger = New(logger, logur.Trace)
		})

		It("Trace", func() {
			adaptedLogger.Trace("testing")
			Expect(output.String()).To(Equal("TRACE testing\n"))
		})

		It("Debug", func() {
			adaptedLogger.Debug("testing")
			Expect(output.String()).To(Equal("DEBUG testing\n"))
		})

		It("Info", func() {
			adaptedLogger.Info("testing")
			Expect(output.String()).To(Equal("INFO testing\n"))
		})

		It("Warn", func() {
			adaptedLogger.Warn("testing")
			Expect(output.String()).To(Equal("WARN testing\n"))
		})

		It("Error", func() {
			adaptedLogger.Error("testing")
			Expect(output.String()).To(Equal("ERROR testing\n"))
		})
	})
})
