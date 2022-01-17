package cryptography

import (
	"encoding/base64"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

const (
	AESTestKey string = "AES256Key-32Characters1234567890"
	Nonce      string = "WA7ChYcNFnCS"
)

var _ = Describe("Cryptography", func() {

	Describe("NewAESGcm", func() {
		It("ok", func() {
			a, err := NewAESGcm([]byte(AESTestKey), []byte(Nonce))
			assert.NoError(GinkgoT(), err)
			assert.NotNil(GinkgoT(), a)
		})

		It("invalid key", func() {
			_, err := NewAESGcm([]byte("abc"), []byte(Nonce))
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "invalid key, should be 16 or 32 bytes")
		})

		It("invalid nonce", func() {
			_, err := NewAESGcm([]byte(AESTestKey), []byte("abc"))
			assert.Error(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), "invalid nonce, should be 12 bytes")
		})
	})

	Describe("Encrypt and Decrypt", func() {
		var aesGcm *AESGcm
		var plaintext []byte
		var encryptedText []byte
		var encryptedTextBase64 string
		BeforeEach(func() {
			aesGcm, _ = NewAESGcm([]byte(AESTestKey), []byte(Nonce))
			plaintext = []byte("exampleplaintext")
			encryptedText = []byte{148, 205, 172, 75, 6, 179, 220, 244, 255, 30, 115, 122, 55, 205, 243, 240, 125, 149, 164, 203, 228, 253, 252, 76, 222, 14, 124, 180, 56, 36, 142, 80}
			encryptedTextBase64 = base64.StdEncoding.EncodeToString(encryptedText)
		})

		Describe("Encrypt", func() {
			It("ok", func() {
				et := aesGcm.Encrypt(plaintext)
				assert.Equal(GinkgoT(), encryptedText, et)
			})
		})

		Describe("Decrypt", func() {
			It("ok", func() {
				dt, err := aesGcm.Decrypt(encryptedText)
				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), plaintext, dt)
			})
		})

		Describe("EncryptToBase64", func() {
			It("ok", func() {
				et := aesGcm.EncryptToBase64(plaintext)
				assert.Equal(GinkgoT(), encryptedTextBase64, et)
			})
		})

		Describe("DecryptFromBase64", func() {
			It("ok", func() {
				dt, err := aesGcm.DecryptFromBase64([]byte(encryptedTextBase64))
				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), string(plaintext), dt)
			})

		})
	})

})

func setup() []byte {
	nonce := []byte(strconv.Itoa(int(time.Now().UTC().Unix())))[:NonceByteSize]

	return nonce
}

func benchmarkAESGCMEncrypt(b *testing.B) {
	text := "http://www.test.com?foo=bar&hello=world"
	nonce := setup()
	aesgcm, _ := NewAESGcm([]byte(AESTestKey), nonce)

	input := []byte(text)
	for i := 0; i < b.N; i++ {
		aesgcm.Encrypt(input)
	}
}

func benchmarkAESGCMDecrypt(b *testing.B) {
	text := "http://www.test.com?foo=bar&hello=world"
	nonce := setup()
	aesgcm, _ := NewAESGcm([]byte(AESTestKey), nonce)

	input := []byte(text)
	encryptedText := aesgcm.Encrypt(input)
	for i := 0; i < b.N; i++ {
		_, _ = aesgcm.Decrypt(encryptedText)
	}
}

func benchmarkAESGCMEncryptToBase64(b *testing.B) {
	text := []byte("http://www.test.com?foo=bar&hello=world")
	nonce := setup()
	aesgcm, _ := NewAESGcm([]byte(AESTestKey), nonce)

	// input := []byte(text)
	for i := 0; i < b.N; i++ {
		aesgcm.EncryptToBase64(text)
	}
}

func benchmarkAESGCMDecryptFromBase64(b *testing.B) {
	text := []byte("http://www.test.com?foo=bar&hello=world")
	nonce := setup()
	aesgcm, _ := NewAESGcm([]byte(AESTestKey), nonce)

	// input := []byte(text)
	encryptedText := []byte(aesgcm.EncryptToBase64(text))
	for i := 0; i < b.N; i++ {
		_, _ = aesgcm.DecryptFromBase64(encryptedText)
	}
}

func BenchmarkAESGCMEncryptDecrypt(b *testing.B) {
	b.Run("cipher", func(b *testing.B) {
		b.Run("Encrypt", benchmarkAESGCMEncrypt)
		b.Run("Decrypt", benchmarkAESGCMDecrypt)
		b.Run("EncryptToBase64", benchmarkAESGCMEncryptToBase64)
		b.Run("DecryptFromBase64", benchmarkAESGCMDecryptFromBase64)
	})
}
