# mapx

> mapx 是一个 Golang 的 map 工具包，包含一些常见的 map 使用快捷方法。

## Usage

你可以通过以下方式使用 mapx 包：

```go
import "github.com/TencentBlueKing/gopkg/mapx"
```

## Spec

### ExistsKey

判断 key 是否存在于 map 中。

```go
m := map[string]interface{}{"k1": "v1"}

// true
mapx.ExistsKey(m, "k1")

// false
mapx.ExistsKey(m, "k2")
```

### Differ

对两个 map 进行对比，输出差异项，支持比较子项。

**备注**
- 子项比较只支持 `[]interface{}`, `map[string]interface{}`
- 若某个 key 包含 `.` 则输出结果中会包含小括号，如 `[]string{"k1", "k2.2", "k3"} -> "k1.(k2.2).k3"`
- 更多样例参考 [differ_test.go](./differ_test.go)

```go
o := map[string]interface{}{
    "k1": "v1", 
    "k2": "v2",
    "k3": map[string]interface{}{
        "k4": "v4",
    },
}

n := map[string]interface{}{
    "k1": "v1.1", 
    "k3": map[string]interface{}{
        "k4": "v4.1",
    }, 
    "k5": "v5",
}

/* 
[
    {"Add",     "k5",     <nil>,  "v5"}
    {"Change",  "k1",     "v1",   "v1.1"}
    {"Change",  "k3.k4",  "v4",   "v4.1"}
    {"Remove",  "k2",     "v2",   <nil>}
]
*/
diffRets := mapx.NewDiffer(o, n).Do()

for _, r := range diffRets {
    /*
        Add k5: v5
        Change k1: v1 -> v1.1
        Change k3.k4: v4 -> v4.1
        Remove k2: v2
     */
    s := r.Repr()
}
```

### GetItems

根据指定路径从嵌套的 `map[string]interface{}` 获取 value 的方法。

```go
m := map[string]interface{}{
    "a1": map[string]interface{}{
        "b1": map[string]interface{}{
            "c1": map[string]interface{}{
                "d1": "v1", 
                "d2": "v2", 
                "d.3": 3,
            },
        },
    },
}

// d1val: v1
d1Val, _ := mapx.GetItems(m, "a1.b1.c1.d1")

// 路径中某 key 存在 `.`，可使用 []string 作为参数
// dDot3Val: 3
dDot3Val, _ := mapx.GetItems(m, []string{"a1", "b1", "c1", "d.3"})

// key 不存在或某中间值类型非 map[string]interface{}，返回错误
// err: key c2 not exist
_, err := mapx.GetItems(m, "a1.b1.c2")
```

### Get

`GetItems` 的快捷方法，支持设置默认值，当 `GetItems` 返回的 `err != nil` 时，返回默认值。

```go
m := ...

// d1val: v1
d1Val := mapx.Get(m, "a1.b1.c1.d1", "default")

// c2Val: default
c2Val := mapx.GetItems(m, "a1.b1.c2", "default")
```

### GetBool

`Get` 的快捷方法，默认返回值为 `false`

### GetInt64

`Get` 的快捷方法，默认返回值为 `int64(0)`

### GetStr

`Get` 的快捷方法，默认返回值为 `""`

### GetList

`Get` 的快捷方法，默认返回值为 `[]interface{}{}`

### GetMap

`Get` 的快捷方法，默认返回值为 `map[string]interface{}{}`

### SetItems

根据指定路径向嵌套的 `map[string]interface{}` 设置 value 的方法。

```go
m := map[string]interface{}{
    "a1": map[string]interface{}{
        "b1": map[string]interface{}{
            "c1": []interface{}{
                "d1", "d2", "d3",
            },
        },
    },
}

/*
   m = map[string]interface{}{
        "a1": map[string]interface{}{
            "b1": map[string]interface{}{
                "c1": []interface{}{
                    "d1", "d2", "d3",
                },
                "c2": "d4",
            },
        },
    }
 */
_ = mapx.SetItems(m, "a1.b1.c2", "d4")

// 当某中间值不存在或其值非 `map[string]interface{}`，返回错误
// err: key c1 not exists or obj[key] not map[string]interface{} type
err := mapx.SetItems(m, "a1.b1.c1.d1", "d5")
```

