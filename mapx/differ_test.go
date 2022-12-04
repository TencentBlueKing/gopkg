/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云 - gopkg available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package mapx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var oldMap = map[string]interface{}{
	"a1": map[string]interface{}{
		"b1": map[string]interface{}{
			"c1": map[string]interface{}{
				"d1": "v1",
				"d2": "v2",
				"d3": 3,
				"d4": []interface{}{4, 5},
				"d5": nil,
				"d6": []interface{}{
					6.1, 6.2, 6.3, 6.4, 6.5,
				},
			},
		},
	},
	"a2": []interface{}{
		map[string]interface{}{
			"b2": map[string]interface{}{
				"c2": []interface{}{
					"d1",
					map[string]interface{}{
						"e1": "v1",
						"e2": "v2",
					},
					map[string]string{
						"e3": "v3",
						"e4": "v4",
					},
					2,
				},
			},
			"b3": []interface{}{
				"c3", "c4", 5,
			},
		},
	},
}

var newMap = map[string]interface{}{
	"a1": map[string]interface{}{
		"b1": map[string]interface{}{
			"c1": map[string]interface{}{
				"d1": "v1",
				// change a1.b1.c1.d2 v2-v1
				"d2": "v1",
				// remove a1.b1.c1.d3 ...
				// add a1.b1.c1.d7 ...
				"d7": 3,
				// remove a1.b1.c1.d4[1] ...
				"d4": []interface{}{4},
				// change a1.b1.c1.d5 nil->"nil"
				"d5": "nil",
				// change a1.b1.c1.d6[2] 6.3->6.4
				// change a1.b1.c1.d6[3] 6.4->6.5
				// change a1.b1.c1.d6[4] 6.5->6.3
				"d6": []interface{}{
					6.1, 6.2, 6.4, 6.5, 6.3,
				},
			},
		},
	},
	"a2": []interface{}{
		map[string]interface{}{
			"b2": map[string]interface{}{
				"c2": []interface{}{
					// change a2[0].b2.c2[0] d1->d2
					"d2",
					map[string]interface{}{
						// change a2[0].b2.c2[1].e1 v1->v2
						"e1": "v2",
						// remove a2[0].b2.c2[1].e2 ...
						// add a2[0].b2.c2[1].e3 ...
						"e3": "v2",
						// add a2[0].b2.c2[1].e4 ...
						"e4": "v4",
						// add a2[0].b2.c2[1].(e5.f1) ...
						"e5.f1": "v5",
					},
					// change a2[0].b2.c2[2] ...
					map[string]string{
						"e3": "v4", // 只是 v3->v4, 但是 map[string]string 不会展开
						"e4": "v4",
					},
					// change a2[0].b2.c2[3] 2->1
					1,
					// add a2[0].b2.c2[4] 2
					2,
				},
			},
			// change a2[0].b3[0] "c3"->"c4"
			// change a2[0].b3[2] 5->6
			// add a2[0].b3[3] 7
			"b3": []interface{}{
				"c4", "c4", 6, 7,
			},
		},
	},
	// add a3 ...
	"a3": map[string]interface{}{
		"b4": "v1",
	},
}

var exceptedDiffRets = DiffRetList{
	{ActionAdd, "a1.b1.c1.d7", nil, 3},
	{ActionAdd, "a2[0].b2.c2[1].(e5.f1)", nil, "v5"},
	{ActionAdd, "a2[0].b2.c2[1].e3", nil, "v2"},
	{ActionAdd, "a2[0].b2.c2[1].e4", nil, "v4"},
	{ActionAdd, "a2[0].b2.c2[4]", nil, 2},
	{ActionAdd, "a2[0].b3[3]", nil, 7},
	{ActionAdd, "a3", nil, map[string]interface{}{"b4": "v1"}},
	{ActionChange, "a1.b1.c1.d2", "v2", "v1"},
	{ActionChange, "a1.b1.c1.d5", nil, "nil"},
	{ActionChange, "a1.b1.c1.d6[2]", 6.3, 6.4},
	{ActionChange, "a1.b1.c1.d6[3]", 6.4, 6.5},
	{ActionChange, "a1.b1.c1.d6[4]", 6.5, 6.3},
	{ActionChange, "a2[0].b2.c2[0]", "d1", "d2"},
	{ActionChange, "a2[0].b2.c2[1].e1", "v1", "v2"},
	{
		ActionChange,
		"a2[0].b2.c2[2]",
		map[string]string{"e3": "v3", "e4": "v4"},
		map[string]string{"e3": "v4", "e4": "v4"},
	},
	{ActionChange, "a2[0].b2.c2[3]", 2, 1},
	{ActionChange, "a2[0].b3[0]", "c3", "c4"},
	{ActionChange, "a2[0].b3[2]", 5, 6},
	{ActionRemove, "a1.b1.c1.d3", 3, nil},
	{ActionRemove, "a1.b1.c1.d4[1]", 5, nil},
	{ActionRemove, "a2[0].b2.c2[1].e2", "v2", nil},
}

func TestDiffer(t *testing.T) {
	diffRets := NewDiffer(oldMap, newMap).Do()
	assert.Equal(t, exceptedDiffRets, diffRets)
}

func TestDiffRetRepr(t *testing.T) {
	addDiffRet := DiffRet{ActionAdd, "a1.b1.c1.d7", nil, 3}
	assert.Equal(t, "Add a1.b1.c1.d7: 3", addDiffRet.Repr())

	changeDiffRet := DiffRet{ActionChange, "a1.b1.c1.d5", nil, "nil"}
	assert.Equal(t, "Change a1.b1.c1.d5: <nil> -> nil", changeDiffRet.Repr())

	changeDiffRet = DiffRet{ActionChange, "a1.b1.c1.d2", "v2", "v1"}
	assert.Equal(t, "Change a1.b1.c1.d2: v2 -> v1", changeDiffRet.Repr())

	removeDiffRet := DiffRet{ActionRemove, "a1.b1.c1.d4[1]", 5, nil}
	assert.Equal(t, "Remove a1.b1.c1.d4[1]: 5", removeDiffRet.Repr())
}
