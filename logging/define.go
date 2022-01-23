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

package logging

import (
	"sync"

	"logur.dev/logur"
)

// Looger is a structured logger interface.
type Logger = logur.LoggerFacade

// DefaultLoggerName is the name of the default logger.
const DefaultLoggerName = ""

var loggerAlias, loggers *sync.Map

func tryGetLogger(name string) (Logger, bool) {
	v, ok := loggers.Load(name)
	if !ok {
		return nil, false
	}

	logger, ok := v.(Logger)
	if !ok {
		return nil, false
	}

	return logger, true
}

// GetLogger returns a named logger. If the logger does not exist, it will return the default logger.
func GetLogger(name string) Logger {
	logger, ok := tryGetLogger(name)
	if ok {
		return logger
	}

	// search for alias when name is not found
	realName, ok := loggerAlias.Load(name)
	if !ok {
		logger, _ = tryGetLogger(DefaultLoggerName)
		return logger
	}

	logger, ok = tryGetLogger(realName.(string))
	if ok {
		return logger
	}

	logger, _ = tryGetLogger(DefaultLoggerName)
	return logger
}

// SetLogger sets a named logger.
func SetLogger(name string, logger Logger) {
	loggers.Store(name, logger)
}

// SetAlias is used to define a logger alias.
func SetAlias(name string, aliases ...string) {
	for _, alias := range aliases {
		loggerAlias.Store(alias, name)
	}
}

func init() {
	loggerAlias = &sync.Map{}
	loggers = &sync.Map{}

	// make sure the default logger is created
	SetLogger(DefaultLoggerName, logur.NoopLogger{})
}
