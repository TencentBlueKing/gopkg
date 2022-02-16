# stringx

`stringx`包含一些字符串常用函数

## Usage


```go
import "github.com/TencentBlueKing/gopkg/stringx"

text := "hello world"

stringx.MD5Hash(text)
stringx.Truncate(text, 10)
stringx.RandomSample(stringx.LowercaseLetters+stringx.Digits, 10)
stringx.RandomAlphanum(10)
```

