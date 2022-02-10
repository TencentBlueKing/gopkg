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

package stringx

import "math/rand"

const (
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters   = Lowercase + Uppercase
	Digits    = "0123456789"
	Alphanum  = Letters + Digits
)

// Truncate string to specific length
func Truncate(s string, n int) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

// RandomSample generate a random string with string sequence and fixed length
func RandomSample(sequence string, n int) string {
	sequenceLen := len(sequence)
	b := make([]byte, n)
	for i := range b {
		b[i] = sequence[rand.Intn(sequenceLen)]
	}
	return string(b)
}

// RandomAlphanum generate a random string with alphanumeric string and fixed length
func RandomAlphanum(n int) string {
	return RandomSample(Alphanum, n)
}
