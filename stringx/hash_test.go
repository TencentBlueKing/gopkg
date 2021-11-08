package stringx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/stringx"
)

var _ = Describe("Hash", func() {

	Describe("MD5Hash", func() {
		DescribeTable("MD5Hash cases", func(expected string, input string) {
			assert.Equal(GinkgoT(), expected, stringx.MD5Hash(input))
		},
			Entry("value is empty string", "d41d8cd98f00b204e9800998ecf8427e", ""),
			Entry("value is 'test'", "098f6bcd4621d373cade4e832627b4f6", "test"),
		)
	})


})
