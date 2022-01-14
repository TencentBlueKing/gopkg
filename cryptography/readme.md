# cryptography

`cryptography` 实现了一些加密函数

## Usage

### aes_gcm

```go
const (
	cryptoKey   = "C4QSNKR4GNPIZAH3B0RPWAIV29E7QZ66"
	aesGcmNonce = "KC9DvYrNGnPW"
)

c, err := cryptography.NewAESGcm([]byte(cryptoKey), []byte(aesGcmNonce))
if err != nil {
    return nil, fmt.Errorf("cryptos key error: %w", err)
}

plain := "hello world"

// plain
cs := c.Encrypt(plain)
ds, err := c.Decrypt(cs)

// base64
cs1 := c.EncryptToBase64(plain)
ds1, err = c.DecryptFromBase64(cs1)
```
