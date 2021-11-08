package set_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/gopkg/collection/set"
)

var _ = Describe("String Set", func() {
	Describe("New", func() {
		It("NewStringSet", func() {
			s := set.NewStringSet()

			assert.Len(GinkgoT(), s.Data, 0)
			assert.Equal(GinkgoT(), 0, s.Size())

			assert.False(GinkgoT(), s.Has("hello"))
		})

		It("NewStringSetWithValues", func() {
			s := set.NewStringSetWithValues([]string{"hello", "world"})

			assert.Len(GinkgoT(), s.Data, 2)
			assert.Equal(GinkgoT(), 2, s.Size())

			assert.True(GinkgoT(), s.Has("hello"))
		})

		It("NewFixedLengthStringSet", func() {
			s := set.NewFixedLengthStringSet(2)

			assert.Len(GinkgoT(), s.Data, 0)
			assert.Equal(GinkgoT(), 0, s.Size())
		})
	})

	Describe("Functions", func() {
		var s *set.StringSet

		BeforeEach(func() {
			s = set.NewStringSet()
			s.Add("hello")
		})

		It("Has", func() {
			assert.True(GinkgoT(), s.Has("hello"))
			assert.False(GinkgoT(), s.Has("world"))
		})

		It("Add", func() {
			s.Add("world")
			assert.True(GinkgoT(), s.Has("world"))
		})

		It("Append", func() {
			s.Append([]string{"abc", "def"}...)
			s.Append([]string{"def", "opq"}...)

			assert.Len(GinkgoT(), s.Data, 4)
			assert.Equal(GinkgoT(), 4, s.Size())

			assert.True(GinkgoT(), s.Has("abc"))
			assert.True(GinkgoT(), s.Has("def"))
			assert.True(GinkgoT(), s.Has("opq"))
		})

		It("Size", func() {
			assert.Equal(GinkgoT(), 1, s.Size())
		})

		It("ToSlice", func() {
			sli1 := s.ToSlice()
			assert.Len(GinkgoT(), sli1, 1)

			s.Add("world")
			sli2 := s.ToSlice()
			assert.Len(GinkgoT(), sli2, 2)
		})

		It("ToString", func() {
			s1 := s.ToString(",")
			assert.Equal(GinkgoT(), "hello", s1)

			s.Add("world")
			s2 := s.ToString(",")

			isEqual := s2 == "hello,world" || s2 == "world,hello"
			//assert.Equal(GinkgoT(), "hello,world", s2)
			assert.True(GinkgoT(), isEqual)

		})

		It("Diff", func() {
			// s = [hello, world]
			s.Add("world")

			// s1 = [world, foo]
			s1 := set.NewStringSetWithValues([]string{"world", "foo"})

			// the diff result
			s2 := s.Diff(s1)

			// the result = [hello]
			assert.Equal(GinkgoT(), 1, s2.Size())
			assert.True(GinkgoT(), s2.Has("hello"))
		})

	})

	Describe("SplitStringToSet", func() {
		It("Empty string", func() {
			s := set.SplitStringToSet("", ",")
			assert.Equal(GinkgoT(), 0, s.Size())
		})

		It("Normal string a,b,c", func() {
			s := set.SplitStringToSet("a,b,c", ",")
			assert.Equal(GinkgoT(), 3, s.Size())
			assert.True(GinkgoT(), s.Has("b"))
		})
	})

})
