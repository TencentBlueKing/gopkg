package errorx

/*
Package `errorx` implements a custom error, wrap/wrapf with detail formatted message

The usage:

1. raw wrap

    import "github.com/TencentBlueKing/gopkg/errorx"

    cnt, err := l.relationManager.GetMemberCount(_type, id)
    if err != nil {
        return errorx.Wrapf(err, "ServiceLayer", "GetMemberCount",
             "relationManager.GetMemberCount _type=`%s`, id=`%s` fail", _type, id)
    }

2. in func with multiple returns

    import "github.com/TencentBlueKing/gopkg/errorx"

    // create a func with layer name and function name
    errorWrapf := errorx.NewLayerFunctionErrorWrapf("ServiceLayer", "BulkDeleteSubjectMember")

    if err != nil {
        return errorWrapf(err, "relationManager.UpdateExpiredAt relations=`%+v` fail", relations)
    }

    // ...

    if err != nil {
        return errorWrapf(err, "relationManager.DoSomething relations=`%+v` fail", relations)
    }
*/

import (
	"errors"
	"fmt"
)

// make the message for error wrap
func makeMessage(err error, layer, function, msg string) string {
	var message string
	var e Errorx
	if errors.As(err, &e) {
		message = fmt.Sprintf("[%s:%s] %s => %s", layer, function, msg, err.Error())
	} else {
		message = fmt.Sprintf("[%s:%s] %s => [Raw:Error] %v", layer, function, msg, err.Error())
	}

	return message
}

// Wrap the error with message
func Wrap(err error, layer string, function string, message string) error {
	if err == nil {
		return nil
	}

	return Errorx{
		message: makeMessage(err, layer, function, message),
		err:     err,
	}
}

// Wrapf the error with formatted message, shortcut for
func Wrapf(err error, layer string, function string, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	return Errorx{
		message: makeMessage(err, layer, function, msg),
		err:     err,
	}
}

// WrapFuncWithLayerFunction define the func of wrapError for partial specific layer name and function name
type WrapFuncWithLayerFunction func(err error, message string) error

// WrapfFuncWithLayerFunction define the func of wrapfError for partial specific layer name and function name
type WrapfFuncWithLayerFunction func(err error, format string, args ...interface{}) error

// NewLayerFunctionErrorWrap create the wrapError func with specific layer and func
func NewLayerFunctionErrorWrap(layer string, function string) WrapFuncWithLayerFunction {
	return func(err error, message string) error {
		return Wrap(err, layer, function, message)
	}
}

// NewLayerFunctionErrorWrapf create the wrapfError func with specific layer and func
func NewLayerFunctionErrorWrapf(layer string, function string) WrapfFuncWithLayerFunction {
	return func(err error, format string, args ...interface{}) error {
		return Wrapf(err, layer, function, format, args...)
	}
}
