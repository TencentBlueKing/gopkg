package conv_test

import (
	"github.com/TencentBlueKing/gopkg/conv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Int64", func() {
	DescribeTable("ToInt64 cases", func(value interface{}, expected int64, willError bool) {
		v, err := conv.ToInt64(value)
		if willError {
			assert.Error(GinkgoT(), err)
		} else {
			assert.NoError(GinkgoT(), err)
			assert.Equal(GinkgoT(), expected, v)
		}
	},
		Entry("value is positive int", 123, int64(123), false),
		Entry("value is negative int", -123, int64(-123), false),
		Entry("value is positive int64", int64(123), int64(123), false),
		Entry("value is negative int64", int64(-123), int64(-123), false),
		Entry("value is string numberic", "123", int64(123), false),
		Entry("value is string numberic", "-123", int64(-123), false),
		Entry("value is float", float64(123.45), int64(123), false),
		Entry("value is nil", nil, int64(0), false),
		Entry("value is int32", int32(123), int64(0), true),
		Entry("value is not numberic", "abc", int64(0), true),
	)

})
