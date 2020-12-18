// bootstrap_linkedlist_foreach.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import . "github.com/zx80live/gofp/fp"

func (n *BoolLinkedList) Foreach(f func(bool)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *StringLinkedList) Foreach(f func(string)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *IntLinkedList) Foreach(f func(int)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Int64LinkedList) Foreach(f func(int64)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *ByteLinkedList) Foreach(f func(byte)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *RuneLinkedList) Foreach(f func(rune)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float32LinkedList) Foreach(f func(float32)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float64LinkedList) Foreach(f func(float64)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *AnyLinkedList) Foreach(f func(Any)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Tuple2LinkedList) Foreach(f func(Tuple2)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *BoolArrayLinkedList) Foreach(f func([]bool)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *StringArrayLinkedList) Foreach(f func([]string)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *IntArrayLinkedList) Foreach(f func([]int)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Int64ArrayLinkedList) Foreach(f func([]int64)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *ByteArrayLinkedList) Foreach(f func([]byte)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *RuneArrayLinkedList) Foreach(f func([]rune)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float32ArrayLinkedList) Foreach(f func([]float32)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float64ArrayLinkedList) Foreach(f func([]float64)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *AnyArrayLinkedList) Foreach(f func([]Any)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Tuple2ArrayLinkedList) Foreach(f func([]Tuple2)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *BoolOptionLinkedList) Foreach(f func(BoolOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *StringOptionLinkedList) Foreach(f func(StringOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *IntOptionLinkedList) Foreach(f func(IntOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Int64OptionLinkedList) Foreach(f func(Int64Option)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *ByteOptionLinkedList) Foreach(f func(ByteOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *RuneOptionLinkedList) Foreach(f func(RuneOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float32OptionLinkedList) Foreach(f func(Float32Option)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float64OptionLinkedList) Foreach(f func(Float64Option)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *AnyOptionLinkedList) Foreach(f func(AnyOption)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Tuple2OptionLinkedList) Foreach(f func(Tuple2Option)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *BoolListLinkedList) Foreach(f func(BoolList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *StringListLinkedList) Foreach(f func(StringList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *IntListLinkedList) Foreach(f func(IntList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Int64ListLinkedList) Foreach(f func(Int64List)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *ByteListLinkedList) Foreach(f func(ByteList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *RuneListLinkedList) Foreach(f func(RuneList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float32ListLinkedList) Foreach(f func(Float32List)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Float64ListLinkedList) Foreach(f func(Float64List)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *AnyListLinkedList) Foreach(f func(AnyList)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}
func (n *Tuple2ListLinkedList) Foreach(f func(Tuple2List)) {
	h := n.head
	for {
		h = h.next
		if h == nil {
			return
		}
		f(*h.value)
	}
}