package main

import (
	"fmt"
)

type LazyInt struct {
	value     func() int
	evaluated *int
}

func (n LazyInt) Value() int {
	if n.evaluated != nil {
		//fmt.Println("evaluated", *n.evaluated)
		return *n.evaluated
	} else {
		//fmt.Println("calculated", n.value())
		return n.value()
	}
}

func (n LazyInt) Evaluate() LazyInt {
	evaluated := n.value()
	return LazyInt{n.value, &evaluated}
}

type LazyState = func() State

type State struct {
	head *LazyInt
	tail *LazyList
}

var EmptyState State = State{nil, nil}
var NilLazyList LazyList = LazyList{nil}

type LazyList struct {
	state *LazyState
}

func (l LazyList) IsEmpty() bool { return l == NilLazyList }

func (l LazyList) NonEmpty() bool { return !l.IsEmpty() }

func (l LazyList) Cons(i LazyInt) LazyList {
	xs := LazyList{l.state}
	state := func() State { return State{&i, &xs} }

	return LazyList{&state}
}

func (l LazyList) Map(f func(e int) int) LazyList {
	newState := func() State {
		state := (*l.state)()
		h := *state.head
		mappedValue := func() int {
			return f(h.value())
		}
		mappedH := LazyInt{mappedValue, nil}
		t := state.tail.Map(f)

		return State{&mappedH, &t}
	}
	return LazyList{&newState}
}

func (l LazyList) Filter(p func(e int) bool) LazyList {
	restRef := l

	newState := func() State {

		var elem LazyInt
		var found = false
		var rest = restRef

		for !found && rest.NonEmpty() {
			restState := (*rest.state)()
			elem = (*restState.head).Evaluate()
			found = p(*elem.evaluated)
			rest = *restState.tail
			restRef = rest
		}

		if found {
			t := rest.Filter(p)
			return State{&elem, &t}
		} else {
			return EmptyState
		}
	}

	return LazyList{&newState}
}

func (l LazyList) Take(n int) LazyList {
	if n <= 0 {
		return NilLazyList
	}

	newState := func() State {
		state := (*l.state)()
		t := state.tail.Take(n - 1)
		return State{state.head, &t}
	}
	return LazyList{&newState}
}

// Creates infinite lazy list of numbers
func InfiniteIntList(from int) LazyList {
	el := LazyInt{func() int { return from }, nil}
	state := func() State {
		rest := InfiniteIntList(from + 1)
		return State{&el, &rest}
	}
	return LazyList{&state}
}

// Creates finite lazy list of numbers
func FiniteIntList(from, to int) LazyList {
	if from == to {
		return NilLazyList
	}

	el := LazyInt{func() int {
		return from
	}, nil}

	state := func() State {
		rest := FiniteIntList(from+1, to)
		return State{&el, &rest}
	}
	return LazyList{&state}
}

func main() {
	fmt.Println("Hello Lazy")

	xs := InfiniteIntList(1).
		//Map(func(e int) int { return e + 1 }).
		//Map(func(e int) int { return e - 1 }).
		Filter(func(e int) bool { return e%2 != 0 }).
		Filter(func(e int) bool { return e > 10 }).Take(5)

	for i := 0; xs.NonEmpty(); i++ {
		state := (*xs.state)()
		fmt.Println((*state.head).Value())
		xs = *state.tail

		if i == 100 {
			break
		}
	}
}
