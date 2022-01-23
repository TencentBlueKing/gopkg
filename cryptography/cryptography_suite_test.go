package cryptography_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCryptography(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cryptography Suite")
}
