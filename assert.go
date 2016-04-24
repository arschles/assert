// Package assert provides convenience assert methods to complement
// the built in go testing library. It's intended to add onto standard
// Go tests. Example usage:
//	func TestSomething(t *testing.T) {
//		i, err := doSomething()
//		assert.NoErr(t, err)
//		assert.Equal(t, i, 123, "returned integer")
//	}
package assert

import (
	"fmt"
	"reflect"
	"runtime"
)

// Tester is a stub interface that *testing.T conforms to. It is used in all
// exported function calls in this assert library so that the library can be
// tested, or a caller can use a custom testing library. As said before,
// however, the most widely used implementation of this interface will
// be *testing.T. Example usage:
//
//	func TestSomething(t *testing.T) {
//		assert.Equal(t, "something", "something", "something")
//	}
type Tester interface {
	Fatalf(string, ...interface{})
}

// callerStr returns a string representation of the code numFrames stack
// frames above the code that called callerStr
func callerStr(numFrames int) string {
	_, file, line, _ := runtime.Caller(1 + numFrames)
	return fmt.Sprintf("%s:%d", file, line)
}

// callerStrf returns a string with fmtStr and vals in it, prefixed
// by a callerStr representation of the code numFrames above the caller of
// this function
func callerStrf(numFrames int, fmtStr string, vals ...interface{}) string {
	origStr := fmt.Sprintf(fmtStr, vals...)
	return fmt.Sprintf("%s: %s", callerStr(1+numFrames), origStr)
}

// True fails the test if b is false. on failure, it calls
// t.Fatalf(fmtStr, vals...)
func True(t Tester, b bool, fmtStr string, vals ...interface{}) {
	if !b {
		t.Fatalf(callerStrf(1, fmtStr, vals...))
	}
}

// False is the equivalent of True(t, !b, fmtStr, vals...).
func False(t Tester, b bool, fmtStr string, vals ...interface{}) {
	if b {
		t.Fatalf(callerStrf(1, fmtStr, vals...))
	}
}

// isNil checks if a specified object is nil or not, without Failing.
// copied from https://github.com/stretchr/testify/blob/master/assert/assertions.go#L310-L323
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}

// Nil uses reflect.DeepEqual(i, nil) to determine if i is nil. if it's not,
// Nil calls t.Fatalf explaining that the noun i is not nil when it should have
// been
func Nil(t Tester, i interface{}, noun string) {
	if !isNil(i) {
		t.Fatalf(callerStrf(1, "the given %s [%+v] was not nil when it should have been", noun, i))
	}
}

// NotNil uses reflect.DeepEqual(i, nil) to determine if i is nil.
// if it is, NotNil calls t.Fatalf explaining that the noun i is nil when it
// shouldn't have been.
func NotNil(t Tester, i interface{}, noun string) {
	if isNil(i) {
		t.Fatalf(callerStrf(1, "the given %s was nil when it shouldn't have been", noun))
	}
}

// Err calls t.Fatalf if expected is not equal to actual.
// it uses reflect.DeepEqual to determine if the errors are equal
func Err(t Tester, expected error, actual error) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(callerStrf(1, "expected error %s but got %s", expected, actual))
	}
}

// ExistsErr calls t.Fatalf if err == nil. The message will explain that the error
// described by noun was nil when it shouldn't have been
func ExistsErr(t Tester, err error, noun string) {
	if err == nil {
		t.Fatalf(callerStrf(1, "given error for %s was nil when it shouldn't have been", noun))
	}
}

// NoErr calls t.Fatalf if e is not nil.
func NoErr(t Tester, e error) {
	if e != nil {
		t.Fatalf(callerStrf(1, "expected no error but got %s", e))
	}
}

// Equal ensures that the actual value returned from a test was equal to an
// expected. it uses reflect.DeepEqual to do so.
// name is used to describe the values being compared. it's used in the error
// string if actual != expected.
func Equal(t Tester, actual, expected interface{}, noun string) {
	actualEqualer, actualEqualerOK := actual.(Equaler)
	expectedEqualer, expectedEqualerOK := expected.(Equaler)
	equals := false
	if actualEqualerOK && expectedEqualerOK {
		equals = actualEqualer.Equal(expectedEqualer)
	} else {
		equals = reflect.DeepEqual(actual, expected)
	}
	if !equals {
		t.Fatalf(callerStrf(1, "actual %s [%+v] != expected %s [%+v]", noun, actual, noun, expected))
	}
}
