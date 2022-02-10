# cryptography

`cryptography` 实现了一些加密函数

## Usage

### aes_gcm

```go
import "github.com/TencentBlueKing/gopkg/cryptography"

const (
	cryptoKey   = "C4QSNKR4GNPIZAH3B0RPWAIV29E7QZ66"
	aesGcmNonce = "KC9DvYrNGnPW"
)

c, err := cryptography.NewAESGcm([]byte(cryptoKey), []byte(aesGcmNonce))
if err != nil {
    return nil, fmt.Errorf("cryptos key error: %w", err)
}

plain := []byte("hello world")

// bytes
cs := c.Encrypt(plain)
ds, err := c.Decrypt(cs)

// string
cs1 := c.EncryptToString(plain)
ds1, err := c.DecryptString(cs1)
```
