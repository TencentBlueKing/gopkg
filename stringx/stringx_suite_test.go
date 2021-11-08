package stringx_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringx Suite")
}
