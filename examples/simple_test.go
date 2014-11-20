package examples

import (
  "github.com/arschles/assert"
  "testing"
)

func TestNil(t *testing.T) {
  assert.Nil(t, nil, "nil was not nil")
}

func TestTrue(t *testing.T) {
  assert.True(t, true, "true was not true")
}

func TestEqual(t *testing.T) {
  s1 := struct{
    a string
    b int
  } {"testString", 1}
  s2 := struct{
    a string
    b int
  }{"testString", 1}
  assert.Equal(t, s1, s2, "struct")
}
