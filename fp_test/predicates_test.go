package fp_test

import (
	. "github.com/zx80live/gofp/fp"
	. "github.com/zx80live/gofp/gotest"
	"regexp"
	"testing"
)

func TestEq(t *testing.T) {
	l := MakeList(1, 2, 3, 4, 5)

	AssertEqualArrays(l.Filter(Eq(4)).ToArray(), []Any{4}, t)
	AssertEqualArrays(l.Filter(Eq(10)).ToArray(), []Any{}, t)
	AssertEqual(l.Find(Eq(4)), 4, t)
	AssertEqual(l.Find(Eq(10)), nil, t)
}

func TestEven(t *testing.T) {
	l := MakeList(1, 2, 3, 4, 5)
	AssertEqualArrays(l.Filter(Even).ToArray(), []Any{2, 4}, t)
}

func TestOdd(t *testing.T) {
	l := MakeList(1, 2, 3, 4, 5)
	AssertEqualArrays(l.Filter(Odd).ToArray(), []Any{1, 3, 5}, t)
}

func TestPositive(t *testing.T) {
	l := MakeList(1, -2, 3, 4, -5)
	AssertEqualArrays(l.Filter(Positive).ToArray(), []Any{1, 3, 4}, t)
}

func TestNegative(t *testing.T) {
	l := MakeList(1, -2, 3, 4, -5)
	AssertEqualArrays(l.Filter(Negative).ToArray(), []Any{-2, -5}, t)
}

func TestRegexpString(t *testing.T) {
	l := MakeList("f-0", "f-1", "aa", "bb", "cc", "ff-2", "dd")
	res := l.Filter(RegexpString("^\\w{1}\\-\\d{1}"))
	AssertEqualArrays(res.ToArray(), []Any{"f-0", "f-1"}, t)
}

func TestRegexp(t *testing.T) {
	l := MakeList("f-0", "f-1", "aa", "bb", "cc", "ff-2", "dd")
	res := l.Filter(Regexp(regexp.MustCompile("^\\w{1}\\-\\d{1}")))
	AssertEqualArrays(res.ToArray(), []Any{"f-0", "f-1"}, t)
}
