// bootstrap_linkedlist_foreachright.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import . "github.com/zx80live/gofp/fp"

func (n *BoolLinkedList) ForeachRight(f func(bool)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *StringLinkedList) ForeachRight(f func(string)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *IntLinkedList) ForeachRight(f func(int)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Int64LinkedList) ForeachRight(f func(int64)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *ByteLinkedList) ForeachRight(f func(byte)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *RuneLinkedList) ForeachRight(f func(rune)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float32LinkedList) ForeachRight(f func(float32)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float64LinkedList) ForeachRight(f func(float64)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *AnyLinkedList) ForeachRight(f func(Any)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Tuple2LinkedList) ForeachRight(f func(Tuple2)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *BoolArrayLinkedList) ForeachRight(f func([]bool)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *StringArrayLinkedList) ForeachRight(f func([]string)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *IntArrayLinkedList) ForeachRight(f func([]int)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Int64ArrayLinkedList) ForeachRight(f func([]int64)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *ByteArrayLinkedList) ForeachRight(f func([]byte)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *RuneArrayLinkedList) ForeachRight(f func([]rune)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float32ArrayLinkedList) ForeachRight(f func([]float32)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float64ArrayLinkedList) ForeachRight(f func([]float64)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *AnyArrayLinkedList) ForeachRight(f func([]Any)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Tuple2ArrayLinkedList) ForeachRight(f func([]Tuple2)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *BoolOptionLinkedList) ForeachRight(f func(BoolOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *StringOptionLinkedList) ForeachRight(f func(StringOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *IntOptionLinkedList) ForeachRight(f func(IntOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Int64OptionLinkedList) ForeachRight(f func(Int64Option)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *ByteOptionLinkedList) ForeachRight(f func(ByteOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *RuneOptionLinkedList) ForeachRight(f func(RuneOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float32OptionLinkedList) ForeachRight(f func(Float32Option)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float64OptionLinkedList) ForeachRight(f func(Float64Option)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *AnyOptionLinkedList) ForeachRight(f func(AnyOption)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Tuple2OptionLinkedList) ForeachRight(f func(Tuple2Option)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *BoolLinkedListLinkedList) ForeachRight(f func(BoolLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *StringLinkedListLinkedList) ForeachRight(f func(StringLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *IntLinkedListLinkedList) ForeachRight(f func(IntLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Int64LinkedListLinkedList) ForeachRight(f func(Int64LinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *ByteLinkedListLinkedList) ForeachRight(f func(ByteLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *RuneLinkedListLinkedList) ForeachRight(f func(RuneLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float32LinkedListLinkedList) ForeachRight(f func(Float32LinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Float64LinkedListLinkedList) ForeachRight(f func(Float64LinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *AnyLinkedListLinkedList) ForeachRight(f func(AnyLinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
func (n *Tuple2LinkedListLinkedList) ForeachRight(f func(Tuple2LinkedList)) {
	e := n.end
	for e.value != nil {
		f(*e.value)
		e = e.prev
	}
}
