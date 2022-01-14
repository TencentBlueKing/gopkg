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

package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"

	"github.com/TencentBlueKing/gopkg/conv"
)

// reference: https://golang.org/src/crypto/cipher/example_test.go

const (
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256)

	ValidAES128KeySize int = 16
	ValidAES256KeySize int = 32

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.

	NonceByteSize int = 12
)

type AESGcm struct {
	key   []byte
	nonce []byte
	// authenticated encryption with associated data (AEAD)
	aead cipher.AEAD
}

func NewAESGcm(key []byte, nonce []byte) (aesGcm *AESGcm, err error) {
	// check key and nonce length
	if len(key) != ValidAES128KeySize && len(key) != ValidAES256KeySize {
		return nil, errors.New("invalid key, should be 16 or 32 bytes")
	}

	if len(nonce) != NonceByteSize {
		return nil, errors.New("invalid nonce, should be 12 bytes")
	}

	// create AEAD
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	return &AESGcm{
		key:   key,
		nonce: nonce,
		aead:  aead,
	}, nil
}

func (a *AESGcm) Encrypt(plaintext []byte) []byte {
	encryptedText := a.aead.Seal(plaintext[:0], a.nonce, plaintext, nil)
	return encryptedText
}

func (a *AESGcm) Decrypt(encryptedText []byte) ([]byte, error) {
	plaintext, err := a.aead.Open(nil, a.nonce, encryptedText, nil)
	return plaintext, err
}

func (a *AESGcm) EncryptToBase64(plaintext string) string {
	plaintextBytes := conv.StringToBytes(plaintext)
	encryptedText := a.Encrypt(plaintextBytes)
	return base64.StdEncoding.EncodeToString(encryptedText)
}

func (a *AESGcm) DecryptFromBase64(encryptedTextB64 string) (plaintext string, err error) {
	var encryptedText []byte
	encryptedText, err = base64.StdEncoding.DecodeString(encryptedTextB64)
	if err != nil {
		return
	}

	var plaintextBytes []byte
	plaintextBytes, err = a.Decrypt(encryptedText)
	if err != nil {
		return
	}

	return conv.BytesToString(plaintextBytes), err
}
