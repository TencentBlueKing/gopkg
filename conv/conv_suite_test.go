package conv_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Conv Suite")
}
