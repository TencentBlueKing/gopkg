package errorx

import (
	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

type NoIsWrapError struct {
	message string
	err     error
}

func (e NoIsWrapError) Error() string {
	return e.message
}

var _ = Describe("Types", func() {
	var e1 error
	var e2 Errorx
	var e3 Errorx
	BeforeEach(func() {
		e1 = errors.New("a")

		e2 = Errorx{
			message: "e2",
			err:     e1,
		}

		e3 = Errorx{
			message: "e3",
			err:     e2,
		}
	})

	Describe("Is", func() {
		It("target is nil", func() {
			assert.False(GinkgoT(), errors.Is(e2, nil))
		})

		It("the e.error is nil", func() {
			e := Errorx{
				message: "e",
				err: nil,
			}
			assert.False(GinkgoT(), errors.Is(e, errors.New("an error")))
		})

		It("ok", func() {
			assert.True(GinkgoT(), errors.Is(e2, e1))
			assert.True(GinkgoT(), errors.Is(e3, e1))
			assert.True(GinkgoT(), errors.Is(e3, e2))

			// false
			assert.False(GinkgoT(), errors.Is(e1, e2))
			assert.False(GinkgoT(), errors.Is(e1, e3))
			assert.False(GinkgoT(), errors.Is(e2, e3))
		})

		It("noIsWrapError", func() {
			e4 := NoIsWrapError{
				message: "no_is_wrap",
				err:     e1,
			}
			e5 := Errorx{
				message: "e5",
				err:     e4,
			}

			assert.True(GinkgoT(), errors.Is(e5, e4))
			assert.False(GinkgoT(), errors.Is(e4, e5))
		})
	})

	Describe("Unwrap", func() {
		assert.Equal(GinkgoT(), e1, e2.Unwrap())
		assert.Equal(GinkgoT(), e1, e3.Unwrap())

		assert.Equal(GinkgoT(), e1, errors.Unwrap(e2))
		assert.Equal(GinkgoT(), e1, errors.Unwrap(e3))
	})

})
