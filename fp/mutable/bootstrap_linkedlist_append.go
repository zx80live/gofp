// bootstrap_linkedlist_append.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import . "github.com/zx80live/gofp/fp"

func (l *BoolLinkedList) Append(e bool) *BoolLinkedList {
	nn := BoolLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *StringLinkedList) Append(e string) *StringLinkedList {
	nn := StringLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *IntLinkedList) Append(e int) *IntLinkedList {
	nn := IntLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Int64LinkedList) Append(e int64) *Int64LinkedList {
	nn := Int64LinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *ByteLinkedList) Append(e byte) *ByteLinkedList {
	nn := ByteLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *RuneLinkedList) Append(e rune) *RuneLinkedList {
	nn := RuneLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float32LinkedList) Append(e float32) *Float32LinkedList {
	nn := Float32LinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float64LinkedList) Append(e float64) *Float64LinkedList {
	nn := Float64LinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *AnyLinkedList) Append(e Any) *AnyLinkedList {
	nn := AnyLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Tuple2LinkedList) Append(e Tuple2) *Tuple2LinkedList {
	nn := Tuple2LinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *BoolArrayLinkedList) Append(e []bool) *BoolArrayLinkedList {
	nn := BoolArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *StringArrayLinkedList) Append(e []string) *StringArrayLinkedList {
	nn := StringArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *IntArrayLinkedList) Append(e []int) *IntArrayLinkedList {
	nn := IntArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Int64ArrayLinkedList) Append(e []int64) *Int64ArrayLinkedList {
	nn := Int64ArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *ByteArrayLinkedList) Append(e []byte) *ByteArrayLinkedList {
	nn := ByteArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *RuneArrayLinkedList) Append(e []rune) *RuneArrayLinkedList {
	nn := RuneArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float32ArrayLinkedList) Append(e []float32) *Float32ArrayLinkedList {
	nn := Float32ArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float64ArrayLinkedList) Append(e []float64) *Float64ArrayLinkedList {
	nn := Float64ArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *AnyArrayLinkedList) Append(e []Any) *AnyArrayLinkedList {
	nn := AnyArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Tuple2ArrayLinkedList) Append(e []Tuple2) *Tuple2ArrayLinkedList {
	nn := Tuple2ArrayLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *BoolOptionLinkedList) Append(e BoolOption) *BoolOptionLinkedList {
	nn := BoolOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *StringOptionLinkedList) Append(e StringOption) *StringOptionLinkedList {
	nn := StringOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *IntOptionLinkedList) Append(e IntOption) *IntOptionLinkedList {
	nn := IntOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Int64OptionLinkedList) Append(e Int64Option) *Int64OptionLinkedList {
	nn := Int64OptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *ByteOptionLinkedList) Append(e ByteOption) *ByteOptionLinkedList {
	nn := ByteOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *RuneOptionLinkedList) Append(e RuneOption) *RuneOptionLinkedList {
	nn := RuneOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float32OptionLinkedList) Append(e Float32Option) *Float32OptionLinkedList {
	nn := Float32OptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float64OptionLinkedList) Append(e Float64Option) *Float64OptionLinkedList {
	nn := Float64OptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *AnyOptionLinkedList) Append(e AnyOption) *AnyOptionLinkedList {
	nn := AnyOptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Tuple2OptionLinkedList) Append(e Tuple2Option) *Tuple2OptionLinkedList {
	nn := Tuple2OptionLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *BoolListLinkedList) Append(e BoolList) *BoolListLinkedList {
	nn := BoolListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *StringListLinkedList) Append(e StringList) *StringListLinkedList {
	nn := StringListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *IntListLinkedList) Append(e IntList) *IntListLinkedList {
	nn := IntListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Int64ListLinkedList) Append(e Int64List) *Int64ListLinkedList {
	nn := Int64ListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *ByteListLinkedList) Append(e ByteList) *ByteListLinkedList {
	nn := ByteListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *RuneListLinkedList) Append(e RuneList) *RuneListLinkedList {
	nn := RuneListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float32ListLinkedList) Append(e Float32List) *Float32ListLinkedList {
	nn := Float32ListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Float64ListLinkedList) Append(e Float64List) *Float64ListLinkedList {
	nn := Float64ListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *AnyListLinkedList) Append(e AnyList) *AnyListLinkedList {
	nn := AnyListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
func (l *Tuple2ListLinkedList) Append(e Tuple2List) *Tuple2ListLinkedList {
	nn := Tuple2ListLinkedList{&e, l.head, nil, l, nil}
	nn.end = &nn
	l.next = &nn
	return &nn
}
