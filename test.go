package test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

func ok(d int, tb testing.TB, errs ...error) {
	for _, err := range errs {
		if err != nil {
			_, file, line, _ := runtime.Caller(d)
			fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
			tb.FailNow()
		}
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, errs ...error) {
	ok(2, tb, errs...)
}

func equals(d int, tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(d)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	equals(2, tb, exp, act)
}

// OkEquals return an asserter that fails if error not is nil or the value doesn't equal what was expected
func OkEquals(tb testing.TB, exp interface{}) func(act interface{}, err error) {
	return func(act interface{}, err error) {
		ok(2, tb, err)
		equals(2, tb, exp, act)
	}
}
