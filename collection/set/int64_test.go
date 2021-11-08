package set_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/collection/set"
)

var _ = Describe("Int64", func() {
	var s *set.Int64Set

	BeforeEach(func() {
		s = set.NewInt64Set()
	})

	It("NewInt64Set", func() {
		//s := util.NewInt64Set()
		assert.Len(GinkgoT(), s.Data, 0)
		assert.Equal(GinkgoT(), 0, s.Size())
	})

	It("NewInt64SetWithValues", func() {
		s1 := set.NewInt64SetWithValues([]int64{123, 456})

		assert.Len(GinkgoT(), s1.Data, 2)
		assert.Equal(GinkgoT(), 2, s1.Size())

		assert.True(GinkgoT(), s1.Has(123))
	})

	It("NewFixedLengthInt64Set", func() {
		s1 := set.NewFixedLengthInt64Set(2)

		assert.Len(GinkgoT(), s1.Data, 0)
		assert.Equal(GinkgoT(), 0, s1.Size())
	})

	It("Add one, check size", func() {
		s.Add(123)

		assert.Len(GinkgoT(), s.Data, 1)
		assert.Equal(GinkgoT(), 1, s.Size())
	})

	It("Append", func() {
		s.Append([]int64{123, 456}...)
		s.Append([]int64{456, 789}...)

		assert.Len(GinkgoT(), s.Data, 3)
		assert.Equal(GinkgoT(), 3, s.Size())

		assert.True(GinkgoT(), s.Has(int64(123)))
		assert.True(GinkgoT(), s.Has(int64(456)))
		assert.True(GinkgoT(), s.Has(int64(789)))
	})

	It("Has 123", func() {
		assert.False(GinkgoT(), s.Has(123))
		s.Add(123)
		assert.True(GinkgoT(), s.Has(123))
	})

	It("ToSlice", func() {
		s.Add(123)
		sli1 := s.ToSlice()
		assert.Len(GinkgoT(), sli1, 1)

		s.Add(456)

		sli2 := s.ToSlice()
		assert.Len(GinkgoT(), sli2, 2)
	})
})
