// bootstrap_list_iterator_next.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (it *BoolListIterator) Next() bool {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *StringListIterator) Next() string {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *IntListIterator) Next() int {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Int64ListIterator) Next() int64 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *ByteListIterator) Next() byte {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *RuneListIterator) Next() rune {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float32ListIterator) Next() float32 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float64ListIterator) Next() float64 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *AnyListIterator) Next() Any {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Tuple2ListIterator) Next() Tuple2 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *BoolArrayListIterator) Next() []bool {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *StringArrayListIterator) Next() []string {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *IntArrayListIterator) Next() []int {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Int64ArrayListIterator) Next() []int64 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *ByteArrayListIterator) Next() []byte {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *RuneArrayListIterator) Next() []rune {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float32ArrayListIterator) Next() []float32 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float64ArrayListIterator) Next() []float64 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *AnyArrayListIterator) Next() []Any {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Tuple2ArrayListIterator) Next() []Tuple2 {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *BoolOptionListIterator) Next() BoolOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *StringOptionListIterator) Next() StringOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *IntOptionListIterator) Next() IntOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Int64OptionListIterator) Next() Int64Option {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *ByteOptionListIterator) Next() ByteOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *RuneOptionListIterator) Next() RuneOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float32OptionListIterator) Next() Float32Option {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float64OptionListIterator) Next() Float64Option {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *AnyOptionListIterator) Next() AnyOption {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Tuple2OptionListIterator) Next() Tuple2Option {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *BoolListListIterator) Next() BoolList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *StringListListIterator) Next() StringList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *IntListListIterator) Next() IntList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Int64ListListIterator) Next() Int64List {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *ByteListListIterator) Next() ByteList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *RuneListListIterator) Next() RuneList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float32ListListIterator) Next() Float32List {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Float64ListListIterator) Next() Float64List {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *AnyListListIterator) Next() AnyList {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
func (it *Tuple2ListListIterator) Next() Tuple2List {
	next := *it.xs.head
	it.xs = it.xs.tail
	return next
}
