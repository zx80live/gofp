package gotest

import (
  "testing"
  . "github.com/zx80live/gofp/fp"
)

func AssertEqual(a Any, b Any, t *testing.T) {
  if a != b {
    t.Errorf("'%v' is not equal '%v'", a, b)
  }
}

func toString(arr []Any) string {
  return Nil.ConsArr(arr).Reverse().MkString("[", ",", "]")
}

func AssertEqualArrays(a []Any, b []Any, t *testing.T) {
  lenA := len(a)
  lenB := len(b)

  if lenA != lenB {
    t.Errorf("arrays have different lengths: \n  %s, \n  %s", toString(a), toString(b))
    return
  }

  for i := 0; i < lenA; i++ {
    if a[i] != b[i] {
      t.Errorf("arrays are different: \n  %s, \n  %s", toString(a), toString(b))
      return
    }
  }
}

func AssertNotEqual(a Any, b Any, t *testing.T) {
  if a == b {
    t.Errorf("'%v' is equal '%v'", a, b)
  }
}

func AssertTrue(a Any, t *testing.T) {
  if a != true {
    t.Errorf("'%v' is not 'true'", a)
  }
}

func AssertFalse(a Any, t *testing.T) {
  if a == true {
    t.Errorf("'%v' is not 'false'", a)
  }
}
