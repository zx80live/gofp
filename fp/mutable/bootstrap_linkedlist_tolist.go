// bootstrap_linkedlist_tolist.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import . "github.com/zx80live/gofp/fp"

func (l BoolLinkedList) ToList() BoolList {
	acc := NilBoolList
	l.ForeachRight(func(e bool) { acc = acc.Cons(e) })
	return acc
}
func (l StringLinkedList) ToList() StringList {
	acc := NilStringList
	l.ForeachRight(func(e string) { acc = acc.Cons(e) })
	return acc
}
func (l IntLinkedList) ToList() IntList {
	acc := NilIntList
	l.ForeachRight(func(e int) { acc = acc.Cons(e) })
	return acc
}
func (l Int64LinkedList) ToList() Int64List {
	acc := NilInt64List
	l.ForeachRight(func(e int64) { acc = acc.Cons(e) })
	return acc
}
func (l ByteLinkedList) ToList() ByteList {
	acc := NilByteList
	l.ForeachRight(func(e byte) { acc = acc.Cons(e) })
	return acc
}
func (l RuneLinkedList) ToList() RuneList {
	acc := NilRuneList
	l.ForeachRight(func(e rune) { acc = acc.Cons(e) })
	return acc
}
func (l Float32LinkedList) ToList() Float32List {
	acc := NilFloat32List
	l.ForeachRight(func(e float32) { acc = acc.Cons(e) })
	return acc
}
func (l Float64LinkedList) ToList() Float64List {
	acc := NilFloat64List
	l.ForeachRight(func(e float64) { acc = acc.Cons(e) })
	return acc
}
func (l AnyLinkedList) ToList() AnyList {
	acc := NilAnyList
	l.ForeachRight(func(e Any) { acc = acc.Cons(e) })
	return acc
}
func (l Tuple2LinkedList) ToList() Tuple2List {
	acc := NilTuple2List
	l.ForeachRight(func(e Tuple2) { acc = acc.Cons(e) })
	return acc
}
func (l BoolArrayLinkedList) ToList() BoolArrayList {
	acc := NilBoolArrayList
	l.ForeachRight(func(e []bool) { acc = acc.Cons(e) })
	return acc
}
func (l StringArrayLinkedList) ToList() StringArrayList {
	acc := NilStringArrayList
	l.ForeachRight(func(e []string) { acc = acc.Cons(e) })
	return acc
}
func (l IntArrayLinkedList) ToList() IntArrayList {
	acc := NilIntArrayList
	l.ForeachRight(func(e []int) { acc = acc.Cons(e) })
	return acc
}
func (l Int64ArrayLinkedList) ToList() Int64ArrayList {
	acc := NilInt64ArrayList
	l.ForeachRight(func(e []int64) { acc = acc.Cons(e) })
	return acc
}
func (l ByteArrayLinkedList) ToList() ByteArrayList {
	acc := NilByteArrayList
	l.ForeachRight(func(e []byte) { acc = acc.Cons(e) })
	return acc
}
func (l RuneArrayLinkedList) ToList() RuneArrayList {
	acc := NilRuneArrayList
	l.ForeachRight(func(e []rune) { acc = acc.Cons(e) })
	return acc
}
func (l Float32ArrayLinkedList) ToList() Float32ArrayList {
	acc := NilFloat32ArrayList
	l.ForeachRight(func(e []float32) { acc = acc.Cons(e) })
	return acc
}
func (l Float64ArrayLinkedList) ToList() Float64ArrayList {
	acc := NilFloat64ArrayList
	l.ForeachRight(func(e []float64) { acc = acc.Cons(e) })
	return acc
}
func (l AnyArrayLinkedList) ToList() AnyArrayList {
	acc := NilAnyArrayList
	l.ForeachRight(func(e []Any) { acc = acc.Cons(e) })
	return acc
}
func (l Tuple2ArrayLinkedList) ToList() Tuple2ArrayList {
	acc := NilTuple2ArrayList
	l.ForeachRight(func(e []Tuple2) { acc = acc.Cons(e) })
	return acc
}
func (l BoolOptionLinkedList) ToList() BoolOptionList {
	acc := NilBoolOptionList
	l.ForeachRight(func(e BoolOption) { acc = acc.Cons(e) })
	return acc
}
func (l StringOptionLinkedList) ToList() StringOptionList {
	acc := NilStringOptionList
	l.ForeachRight(func(e StringOption) { acc = acc.Cons(e) })
	return acc
}
func (l IntOptionLinkedList) ToList() IntOptionList {
	acc := NilIntOptionList
	l.ForeachRight(func(e IntOption) { acc = acc.Cons(e) })
	return acc
}
func (l Int64OptionLinkedList) ToList() Int64OptionList {
	acc := NilInt64OptionList
	l.ForeachRight(func(e Int64Option) { acc = acc.Cons(e) })
	return acc
}
func (l ByteOptionLinkedList) ToList() ByteOptionList {
	acc := NilByteOptionList
	l.ForeachRight(func(e ByteOption) { acc = acc.Cons(e) })
	return acc
}
func (l RuneOptionLinkedList) ToList() RuneOptionList {
	acc := NilRuneOptionList
	l.ForeachRight(func(e RuneOption) { acc = acc.Cons(e) })
	return acc
}
func (l Float32OptionLinkedList) ToList() Float32OptionList {
	acc := NilFloat32OptionList
	l.ForeachRight(func(e Float32Option) { acc = acc.Cons(e) })
	return acc
}
func (l Float64OptionLinkedList) ToList() Float64OptionList {
	acc := NilFloat64OptionList
	l.ForeachRight(func(e Float64Option) { acc = acc.Cons(e) })
	return acc
}
func (l AnyOptionLinkedList) ToList() AnyOptionList {
	acc := NilAnyOptionList
	l.ForeachRight(func(e AnyOption) { acc = acc.Cons(e) })
	return acc
}
func (l Tuple2OptionLinkedList) ToList() Tuple2OptionList {
	acc := NilTuple2OptionList
	l.ForeachRight(func(e Tuple2Option) { acc = acc.Cons(e) })
	return acc
}
func (l BoolListLinkedList) ToList() BoolListList {
	acc := NilBoolListList
	l.ForeachRight(func(e BoolList) { acc = acc.Cons(e) })
	return acc
}
func (l StringListLinkedList) ToList() StringListList {
	acc := NilStringListList
	l.ForeachRight(func(e StringList) { acc = acc.Cons(e) })
	return acc
}
func (l IntListLinkedList) ToList() IntListList {
	acc := NilIntListList
	l.ForeachRight(func(e IntList) { acc = acc.Cons(e) })
	return acc
}
func (l Int64ListLinkedList) ToList() Int64ListList {
	acc := NilInt64ListList
	l.ForeachRight(func(e Int64List) { acc = acc.Cons(e) })
	return acc
}
func (l ByteListLinkedList) ToList() ByteListList {
	acc := NilByteListList
	l.ForeachRight(func(e ByteList) { acc = acc.Cons(e) })
	return acc
}
func (l RuneListLinkedList) ToList() RuneListList {
	acc := NilRuneListList
	l.ForeachRight(func(e RuneList) { acc = acc.Cons(e) })
	return acc
}
func (l Float32ListLinkedList) ToList() Float32ListList {
	acc := NilFloat32ListList
	l.ForeachRight(func(e Float32List) { acc = acc.Cons(e) })
	return acc
}
func (l Float64ListLinkedList) ToList() Float64ListList {
	acc := NilFloat64ListList
	l.ForeachRight(func(e Float64List) { acc = acc.Cons(e) })
	return acc
}
func (l AnyListLinkedList) ToList() AnyListList {
	acc := NilAnyListList
	l.ForeachRight(func(e AnyList) { acc = acc.Cons(e) })
	return acc
}
func (l Tuple2ListLinkedList) ToList() Tuple2ListList {
	acc := NilTuple2ListList
	l.ForeachRight(func(e Tuple2List) { acc = acc.Cons(e) })
	return acc
}
