// bootstrap_linkedlist_iterator_next.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import . "github.com/zx80live/gofp/fp"

func (it *BoolLinkedListIterator) Next() Bool {
	next := *it.xs.value
	it.xs = it.xs.next
	return Bool(next)
}
func (it *StringLinkedListIterator) Next() String {
	next := *it.xs.value
	it.xs = it.xs.next
	return String(next)
}
func (it *IntLinkedListIterator) Next() Int {
	next := *it.xs.value
	it.xs = it.xs.next
	return Int(next)
}
func (it *Int64LinkedListIterator) Next() Int64 {
	next := *it.xs.value
	it.xs = it.xs.next
	return Int64(next)
}
func (it *ByteLinkedListIterator) Next() Byte {
	next := *it.xs.value
	it.xs = it.xs.next
	return Byte(next)
}
func (it *RuneLinkedListIterator) Next() Rune {
	next := *it.xs.value
	it.xs = it.xs.next
	return Rune(next)
}
func (it *Float32LinkedListIterator) Next() Float32 {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float32(next)
}
func (it *Float64LinkedListIterator) Next() Float64 {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float64(next)
}
func (it *AnyLinkedListIterator) Next() Any {
	next := *it.xs.value
	it.xs = it.xs.next
	return Any(next)
}
func (it *Tuple2LinkedListIterator) Next() Tuple2 {
	next := *it.xs.value
	it.xs = it.xs.next
	return Tuple2(next)
}
func (it *BoolArrayLinkedListIterator) Next() BoolArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return BoolArray(next)
}
func (it *StringArrayLinkedListIterator) Next() StringArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return StringArray(next)
}
func (it *IntArrayLinkedListIterator) Next() IntArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return IntArray(next)
}
func (it *Int64ArrayLinkedListIterator) Next() Int64Array {
	next := *it.xs.value
	it.xs = it.xs.next
	return Int64Array(next)
}
func (it *ByteArrayLinkedListIterator) Next() ByteArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return ByteArray(next)
}
func (it *RuneArrayLinkedListIterator) Next() RuneArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return RuneArray(next)
}
func (it *Float32ArrayLinkedListIterator) Next() Float32Array {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float32Array(next)
}
func (it *Float64ArrayLinkedListIterator) Next() Float64Array {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float64Array(next)
}
func (it *AnyArrayLinkedListIterator) Next() AnyArray {
	next := *it.xs.value
	it.xs = it.xs.next
	return AnyArray(next)
}
func (it *Tuple2ArrayLinkedListIterator) Next() Tuple2Array {
	next := *it.xs.value
	it.xs = it.xs.next
	return Tuple2Array(next)
}
func (it *BoolOptionLinkedListIterator) Next() BoolOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return BoolOption(next)
}
func (it *StringOptionLinkedListIterator) Next() StringOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return StringOption(next)
}
func (it *IntOptionLinkedListIterator) Next() IntOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return IntOption(next)
}
func (it *Int64OptionLinkedListIterator) Next() Int64Option {
	next := *it.xs.value
	it.xs = it.xs.next
	return Int64Option(next)
}
func (it *ByteOptionLinkedListIterator) Next() ByteOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return ByteOption(next)
}
func (it *RuneOptionLinkedListIterator) Next() RuneOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return RuneOption(next)
}
func (it *Float32OptionLinkedListIterator) Next() Float32Option {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float32Option(next)
}
func (it *Float64OptionLinkedListIterator) Next() Float64Option {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float64Option(next)
}
func (it *AnyOptionLinkedListIterator) Next() AnyOption {
	next := *it.xs.value
	it.xs = it.xs.next
	return AnyOption(next)
}
func (it *Tuple2OptionLinkedListIterator) Next() Tuple2Option {
	next := *it.xs.value
	it.xs = it.xs.next
	return Tuple2Option(next)
}
func (it *BoolListLinkedListIterator) Next() BoolList {
	next := *it.xs.value
	it.xs = it.xs.next
	return BoolList(next)
}
func (it *StringListLinkedListIterator) Next() StringList {
	next := *it.xs.value
	it.xs = it.xs.next
	return StringList(next)
}
func (it *IntListLinkedListIterator) Next() IntList {
	next := *it.xs.value
	it.xs = it.xs.next
	return IntList(next)
}
func (it *Int64ListLinkedListIterator) Next() Int64List {
	next := *it.xs.value
	it.xs = it.xs.next
	return Int64List(next)
}
func (it *ByteListLinkedListIterator) Next() ByteList {
	next := *it.xs.value
	it.xs = it.xs.next
	return ByteList(next)
}
func (it *RuneListLinkedListIterator) Next() RuneList {
	next := *it.xs.value
	it.xs = it.xs.next
	return RuneList(next)
}
func (it *Float32ListLinkedListIterator) Next() Float32List {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float32List(next)
}
func (it *Float64ListLinkedListIterator) Next() Float64List {
	next := *it.xs.value
	it.xs = it.xs.next
	return Float64List(next)
}
func (it *AnyListLinkedListIterator) Next() AnyList {
	next := *it.xs.value
	it.xs = it.xs.next
	return AnyList(next)
}
func (it *Tuple2ListLinkedListIterator) Next() Tuple2List {
	next := *it.xs.value
	it.xs = it.xs.next
	return Tuple2List(next)
}