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
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
)

/*
Package `metric` implements metric collector for prometheus.

The usage:

collector := NewRedisCollector(serviceName, clients)
prometheus.MustRegister(collector)
*/

// RedisCollector ...
type RedisCollector struct {
	name string

	clients map[string]*redis.Client
}

// NewRedisCollector ...
func NewRedisCollector(serviceName string, clients map[string]*redis.Client) *RedisCollector {
	return &RedisCollector{
		name:    serviceName + "_redis_available",
		clients: clients,
	}
}

// Describe ...
func (c *RedisCollector) Describe(ch chan<- *prometheus.Desc) {
	desc := prometheus.NewDesc(c.name, "redis available", []string{"id"}, nil)
	ch <- desc
}

// Collect ...
func (c *RedisCollector) Collect(ch chan<- prometheus.Metric) {
	for id, cli := range c.clients {
		_, err := cli.Ping(context.TODO()).Result()
		var value float64
		if err != nil {
			value = 0
		} else {
			value = 1
		}

		ch <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(c.name, "redis available", []string{"id"}, nil),
			prometheus.GaugeValue,
			value,
			id,
		)

	}
}
