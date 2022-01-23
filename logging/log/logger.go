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
package log

import (
	"context"
	"fmt"
	"log"
	"strings"

	"logur.dev/logur"
)

// Logger is a Logur adapter for TEMPLATE.
type Logger struct {
	level  logur.Level
	logger *log.Logger
}

// New returns a new Logur logger.
// If logger is nil, a default instance is created.
func New(logger *log.Logger, level logur.Level) *Logger {
	return &Logger{
		level:  level,
		logger: logger,
	}
}

func (l *Logger) log(level logur.Level, prefix string, msg string, fields ...map[string]interface{}) {
	if level < l.level {
		return
	}

	fieldCount := 2
	if len(fields) != 0 {
		fieldCount += len(fields[0])
	}

	parts := make([]string, 0, fieldCount)
	parts = append(parts, prefix, msg)

	if len(fields) != 0 {
		for key, value := range fields[0] {
			parts = append(parts, fmt.Sprintf("%v=%v", key, value))
		}
	}

	l.logger.Print(strings.Join(parts, " "))
}

// Trace implements the Logur Logger interface.
func (l *Logger) Trace(msg string, fields ...map[string]interface{}) {
	l.log(logur.Trace, "TRACE", msg, fields...)
}

// Debug implements the Logur Logger interface.
func (l *Logger) Debug(msg string, fields ...map[string]interface{}) {
	l.log(logur.Debug, "DEBUG", msg, fields...)
}

// Info implements the Logur Logger interface.
func (l *Logger) Info(msg string, fields ...map[string]interface{}) {
	l.log(logur.Info, "INFO", msg, fields...)
}

// Warn implements the Logur Logger interface.
func (l *Logger) Warn(msg string, fields ...map[string]interface{}) {
	l.log(logur.Warn, "WARN", msg, fields...)
}

// Error implements the Logur Logger interface.
func (l *Logger) Error(msg string, fields ...map[string]interface{}) {
	l.log(logur.Error, "ERROR", msg, fields...)
}

func (l *Logger) TraceContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Trace(msg, fields...)
}

func (l *Logger) DebugContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(_ context.Context, msg string, fields ...map[string]interface{}) {
	l.Error(msg, fields...)
}
