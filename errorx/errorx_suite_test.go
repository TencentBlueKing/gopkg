package errorx_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestErrorx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errorx Suite")
}
