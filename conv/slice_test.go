package conv_test

import (
	"github.com/TencentBlueKing/gopkg/conv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Slice", func() {

	Describe("ToSlice", func() {

		intSlice := []int{1}
		strSlice := []string{"abc"}

		DescribeTable("ToSlice cases", func(expected int, willError bool, input interface{}) {
			data, err := conv.ToSlice(input)

			if willError {
				assert.Error(GinkgoT(), err)
			} else {
				assert.NoError(GinkgoT(), err)
				assert.Equal(GinkgoT(), expected, len(data))
			}
		},
			Entry("not a slice", 0, true, ""),
			Entry("a []int{1}", 1, false, intSlice),
			Entry("a []string{abc}", 1, false, strSlice),
		)
	})

})
