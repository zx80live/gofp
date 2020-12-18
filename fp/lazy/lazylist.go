package main

import (
	"fmt"
)

type LazyInt struct {
	eval   func() int
	cached *int
}

func (n LazyInt) Value() int {
	if n.cached != nil {
		fmt.Println(" int.cached", *n.cached)
		return *n.cached
	} else {
		fmt.Println("*int.eval", n.eval())
		return n.eval()
	}
}

func (n LazyInt) Eval() LazyInt {
	if n.cached != nil {
		fmt.Println(" int.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		fmt.Println("*int.eval", cached)
		return LazyInt{n.eval, &cached}
	}
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
			return f(h.eval())
		}
		mappedH := LazyInt{mappedValue, nil}
		t := state.tail.Map(f)

		return State{&mappedH, &t}
	}
	return LazyList{&newState}
}

func (l LazyList) Filter(p func(e int) bool) LazyList {
	newState := func() State {
		var xs = l
		var h LazyInt
		var found = false

		for !found && xs.NonEmpty() {
			s := (*xs.state)()
			h = (*s.head).Eval()
			found = p(*h.cached)
			xs = *s.tail
		}

		if found {
			t := xs.Filter(p)
			return State{&h, &t}
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
		Filter(func(e int) bool { return e%2 == 0 }).
		Filter(func(e int) bool { return e > -100 }).
		Filter(func(e int) bool { return e > 0 }).
		//Map(func(e int) int { return e + 1 }).
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
