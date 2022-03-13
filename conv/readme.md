# conv

`conv` 实现了一些类型转换函数


## Usage

### string - bytes

```go
import "github.com/TencentBlueKing/gopkg/conv"

conv.StringToBytes("hello world")
conv.BytesToString([]byte("hello world"))
conv.ToString(123)
```

### int64

```go
import "github.com/TencentBlueKing/gopkg/conv"

conv.ToInt64("123")
```

### slice

```go
import "github.com/TencentBlueKing/gopkg/conv"

var i interface{}
i = []int{123}
conv.ToSlice(i)
```
