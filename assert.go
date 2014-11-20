//package assert provides convenience assert methods to complement
//the built in go testing library. It's intended to add onto standard
//Go tests. Example usage:
//  func TestSomething(t *testing.T) {
//    i, err := doSomething()
//    assert.NoErr(err)
//    assert.Equal(i, 123, "returned integer")
//  }
package assert

import (
	"fmt"
	"runtime"
	"testing"
	"reflect"
)

//callerStr returns a string representation of the code numFrames stack
//frames above the code that called callerStr
func callerStr(numFrames int) string {
	_, file, line, _ := runtime.Caller(1 + numFrames)
	return fmt.Sprintf("%s:%d", file, line)
}

//callerStrErrorf calls t.Errorf with fmtStr and vals in it, prefixed
//by a callerStr representation of the code numFrames above the caller of
//this function
func callerStrf(numFrames int, fmtStr string, vals ...interface{}) string {
	origStr := fmt.Sprintf(fmtStr, vals...)
	return fmt.Sprintf("%s: %s", callerStr(1+numFrames), origStr)
}

//True calls t.Errorf if the provided bool is false, does nothing
//otherwise
func True(t *testing.T, b bool, fmtStr string, vals ...interface{}) {
	if !b {
		t.Errorf(callerStrf(1, fmtStr, vals...))
	}
}

//False is the equivalent of True(t, !b, fmtStr, vals...)
func False(t *testing.T, b bool, fmtStr string, vals ...interface{}) {
	if b {
		t.Errorf(callerStrf(1, fmtStr, vals...))
	}
}

//Nil calls t.Errorf if i is not nil
func Nil(t *testing.T, i interface{}, fmtStr string, vals ...interface{}) {
	if i != nil {
		t.Errorf(callerStrf(1, fmtStr, vals...))
	}
}

//NoErr calls t.Errorf if e is not nil
func NoErr(t *testing.T, e error) {
	if e != nil {
		t.Errorf(callerStrf(1, "expected no error but got %s", e))
	}
}

//Err calls t.Errorf if expected is not equal to actual
func Err(t *testing.T, expected error, actual error) {
	if expected != actual {
		t.Errorf(callerStrf(1, "expected error %s but got %s", expected, actual))
	}
}

//Equal ensures that the actual value returned from a test was equal to an
//expected. it uses reflect.DeepEqual to do so.
//name is the name used to describe the values being compared.
//it's used in the error string if actual != expected
func Equal(t *testing.T, actual, expected interface{}, noun string) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(callerStrf(1, "actual %s [%+v] != expected %s [%+v]", noun, actual, noun, expected))
	}
}
