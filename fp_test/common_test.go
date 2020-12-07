package fp_test

/*
import (
	. "github.com/zx80live/gofp/fp"
	. "github.com/zx80live/gofp/gotest"
	"testing"
)

func TestEmptyPredicate(t *testing.T) {
	dataset := []Any{0, 1, -10, 'a', "Hello"}
	for _, e := range dataset {
		AssertTrue(EmptyPredicate(e), t)
	}
}

func TestEmptyFunctor(t *testing.T) {
	dataset := []Any{0, 1, -10, 'a', "Hello"}
	for _, e := range dataset {
		AssertEqual(EmptyFunctor(e), e, t)
	}
}

func TestPredicateFilter(t *testing.T) {
	var even, negatives, combined Predicate
	even = func(e Any) bool {
		return e.(int)%2 == 0
	}

	negatives = func(e Any) bool {
		return e.(int) < 0
	}

	combined = even.Filter(negatives)

	dataset := [][]Any{
		{-10, true},
		{-2, true},
		{-4, true},
		{-8, true},
		{-100, true},
		{-3, false},
		{2, false},
		{4, false},
		{10, false},
		{15, false},
	}

	for _, e := range dataset {
		n := e[0]
		expected := e[1]

		AssertEqual(combined(n), expected, t)
	}
}

func TestFunctorMap(t *testing.T) {
	var mult100, inverter, combined Functor
	mult100 = func(e Any) Any {
		return e.(int) * 100
	}
	inverter = func(e Any) Any {
		return -e.(int)
	}

	combined = mult100.Map(inverter)

	dataset := [][]Any{
		{0, 0},
		{1, -100},
		{2, -200},
		{3, -300},
		{-4, 400},
		{-5, 500},
	}

	for _, e := range dataset {
		n := e[0]
		expected := e[1]
		actual := combined(n)

		if actual != expected {
			t.Errorf("expected '%v', got '%v'", expected, actual)
		}
	}

}
*/
