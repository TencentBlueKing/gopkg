package stringx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/stringx"
)

var _ = Describe("Stringx", func() {
	Describe("Truncate", func() {
		var s = "helloworld"

		DescribeTable("Truncate String cases", func(expected string, truncatedSize int) {
			assert.Equal(GinkgoT(), expected, stringx.Truncate(s, truncatedSize))
		},
			Entry("truncated size less than real size", "he", 2),
			Entry("truncated size equals to real size", s, 10),
			Entry("truncated size greater than real size", s, 20),
		)

		Describe("Random", func() {
			DescribeTable("Random String cases", func(length int) {
				assert.Equal(GinkgoT(), length, len(stringx.Random(length)))
			},
				Entry("string length 0", 0),
				Entry("string length 1", 10),
				Entry("string length 10", 10),
			)
		})
	})


})
