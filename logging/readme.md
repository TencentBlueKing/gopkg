# 统一日志界面
为了方便项目适配其日志实现，请使用统一日志界面来记录日志。
该日志界面主要基于 logur 统一日志接口，项目可以在实际使用时，可注入具体的日志实现。

## 基本用法
类似 python logging 模块，需要使用名字获取对应的日志对象，请使用模块导入地址来设置和获取，如：
```go
// 设置模块日志
logger := logging.GetLogger("github.com/TencentBlueKing/gopkg")

// 如果一个模块有多个日志对象，用冒号隔开
loggerA := logging.GetLogger("github.com/TencentBlueKing/gopkg:a")
loggerB := logging.GetLogger("github.com/TencentBlueKing/gopkg:b")
```

## 设置日志实现
可以使用 log/zap/logrus 来作为日志实现，默认不选择任何实现，按需开启。你也可以实现 `logging.Logger` 接口，自定义日志实现，如：
```go
// 实现 `logging.Logger` 接口，作为模块日志实现
logging.SetLogger("github.com/TencentBlueKing/gopkg", myLogger)
```

## 设置别名
为了方便配置，可以使用别名，多个日志会使用同一个实现，如：
```go
// 设置一个日志
logging.SetLogger("github.com/TencentBlueKing/gopkg", myLogger)

// 设置别名
logging.SetLogger(
    "github.com/TencentBlueKing/gopkg",  // 具体的实现
    "github.com/TencentBlueKing/gopkg:a",  // 别名1
    "github.com/TencentBlueKing/gopkg:b",  // 别名2
)
```

## 设置默认实现
如果你想使用默认的日志实现，请使用以下方法：
```go
// 设置默认日志实现，当未匹配到名称和别名时，会尝试使用默认实现
SetLogger(DefaultLoggerName, myLogger)
```