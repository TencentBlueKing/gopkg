package errorx

import (
	"errors"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Errorx", func() {

	Describe("makeMessage", func() {

		It("raw error", func() {
			err := errors.New("it's a raw error")

			msg := makeMessage(err, "Test", "makeMessage", "wrapped-message")
			assert.Equal(GinkgoT(), "[Test:makeMessage] wrapped-message => [Raw:Error] it's a raw error", msg)
		})

		It("errorx.Error", func() {
			err := errors.New("it's a raw error")
			errx := Errorx{
				message: "wrap errorx message",
				err:     err,
			}

			msg := makeMessage(errx, "Test", "makeMessage", "wrapped-message")
			assert.Equal(GinkgoT(), "[Test:makeMessage] wrapped-message => wrap errorx message", msg)
		})

		It("Wrap nil", func() {
			var err error
			err = nil
			assert.Nil(GinkgoT(), Wrap(err, "Test", "Wrap", "wrapped-message"))
		})

		It("Wrap", func() {
			err := errors.New("it's a raw error")

			e := Wrap(err, "Test", "Wrap", "wrapped-message")
			assert.Equal(GinkgoT(), "[Test:Wrap] wrapped-message => [Raw:Error] it's a raw error", e.Error())
		})

		It("Wrap-Wrap", func() {
			err := errors.New("it's a raw error")

			e := Wrap(err, "Test", "Wrap", "wrapped-message")

			e2 := Wrap(e, "Test2", "Wrap2", "wrapped-message-2")
			assert.Equal(GinkgoT(), "[Test2:Wrap2] wrapped-message-2 => [Test:Wrap] wrapped-message => [Raw:Error] it's a raw error", e2.Error())
		})

		It("Wrapf nil", func() {
			var err error
			err = nil
			assert.Nil(GinkgoT(), Wrapf(err, "Test", "Wrap", "wrapped-message %d", 100))
		})

		It("Wrapf", func() {
			err := errors.New("it's a raw error")

			e := Wrapf(err, "Test", "Wrap", "wrapped-message %d", 100)
			assert.Equal(GinkgoT(), "[Test:Wrap] wrapped-message 100 => [Raw:Error] it's a raw error", e.Error())
		})

		It("NewLayerFunctionErrorWrap", func() {
			err := errors.New("it's a raw error")

			f := NewLayerFunctionErrorWrap("Test", "call")
			e := f(err, "hello")
			assert.Equal(GinkgoT(), "[Test:call] hello => [Raw:Error] it's a raw error", e.Error())
		})

		It("NewLayerFunctionErrorWrapf", func() {
			err := errors.New("it's a raw error")

			f := NewLayerFunctionErrorWrapf("Test", "call")
			e := f(err, "hello %s", "world")
			assert.Equal(GinkgoT(), "[Test:call] hello world => [Raw:Error] it's a raw error", e.Error())
		})

	})

})
