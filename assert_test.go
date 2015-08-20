package assert

import (
	"errors"
	"fmt"
	"testing"
)

type testCase struct {
	val  interface{}
	name string
}

var (
	nonNilTestCases = []testCase{
		testCase{val: 0, name: "0"},
		testCase{val: 1, name: "1"},
		testCase{val: true, name: "true"},
		testCase{val: false, name: "false"},
		testCase{val: "abc", name: `"abc"`},
		testCase{val: struct{}{}, name: "struct{}{}"},
	}
)

func TestCallerStr(t *testing.T) {
	s := callerStr(0)
	if len(s) <= 0 {
		t.Errorf("return value of callerStr is empty")
	}
}

func TestCallerStrf(t *testing.T) {
	fmtStr := "%d%d"
	val1 := 1
	val2 := 2
	res := callerStrf(0, fmtStr, val1, val2)
	if len(res) < 3 {
		t.Errorf("return value of callerStrf is not long enough")
	}
}

func TestTrue(t *testing.T) {
	True(t, true, "true wasn't reported as true")
}

func TestFalse(t *testing.T) {
	False(t, false, "false wasn't reported as false")
}

func TestNil(t *testing.T) {
	var i *int = nil
	var b *bool = nil
	var s *string = nil
	var slc *[]string = nil
	var str *struct{} = nil
	tests := []testCase{
		testCase{val: nil, name: "nil"},
		testCase{val: i, name: "*int"},
		testCase{val: b, name: "*bool"},
		testCase{val: s, name: "*string"},
		testCase{val: slc, name: "*[]string"},
		testCase{val: str, name: "*struct{}"},
	}
	for _, test := range tests {
		Nil(t, test.val, test.name)
	}
}

func TestNotNil(t *testing.T) {
	for _, test := range nonNilTestCases {
		NotNil(t, test.val, test.name)
	}
}

func TestErr(t *testing.T) {
	err11 := errors.New("err1")
	err12 := errors.New("err1")
	Err(t, err11, err12)
	err21 := fmt.Errorf("err2-%s", "a")
	err22 := fmt.Errorf("err2-%s", "a")
	Err(t, err21, err22)
}

func TestExistsErr(t *testing.T) {
	ExistsErr(t, errors.New("abc"), "error")
}

func TestNoErr(t *testing.T) {
	NoErr(t, nil)
}

func TestEqual(t *testing.T) {
	for _, test := range nonNilTestCases {
		Equal(t, test.val, test.val, test.name)
	}
}
