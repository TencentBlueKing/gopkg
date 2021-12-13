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

package cache

import "fmt"

type StringKey struct {
	key string
}

func NewStringKey(key string) StringKey {
	return StringKey{
		key: key,
	}
}

func (s StringKey) Key() string {
	return s.key
}

type IntKey struct {
	key int
}

func NewIntKey(key int) IntKey {
	return IntKey{
		key: key,
	}
}

func (k IntKey) Key() string {
	return fmt.Sprintf("%d", k.key)
}

type Int64Key struct {
	key int64
}

func NewInt64Key(key int64) Int64Key {
	return Int64Key{
		key: key,
	}
}

func (k Int64Key) Key() string {
	return fmt.Sprintf("%d", k.key)
}

type UintKey struct {
	key uint
}

func NewUintKey(key uint) UintKey {
	return UintKey{
		key: key,
	}
}

func (k UintKey) Key() string {
	return fmt.Sprintf("%d", k.key)
}

type Uint64Key struct {
	key uint64
}

func NewUint64Key(key uint64) Uint64Key {
	return Uint64Key{
		key: key,
	}
}

func (k Uint64Key) Key() string {
	return fmt.Sprintf("%d", k.key)
}
