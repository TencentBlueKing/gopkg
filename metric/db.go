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

package metric

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

/*
Package `metric` implements metric collector for prometheus.

The usage:

collector := NewDatabaseCollector(serviceName, connections)
prometheus.MustRegister(collector)
*/

// DatabaseCollector ...
type DatabaseCollector struct {
	name string

	connections map[string]*sql.DB
}

// NewDatabaseCollector ...
func NewDatabaseCollector(serviceName string, connections map[string]*sql.DB) *DatabaseCollector {
	return &DatabaseCollector{
		name:        serviceName + "_database_available",
		connections: connections,
	}
}

// Describe ...
func (c *DatabaseCollector) Describe(ch chan<- *prometheus.Desc) {
	desc := prometheus.NewDesc(c.name, "database available", []string{"id"}, nil)
	ch <- desc
}

// Collect ...
func (c *DatabaseCollector) Collect(ch chan<- prometheus.Metric) {
	for id, conn := range c.connections {
		err := conn.Ping()
		var value float64
		if err != nil {
			value = 0
		} else {
			value = 1
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(c.name, "database available", []string{"id"}, nil),
			prometheus.GaugeValue,
			value,
			id,
		)
	}
}
