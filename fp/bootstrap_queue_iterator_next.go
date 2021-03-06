// bootstrap_queue_iterator_next.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (it *BoolQueueIterator) Next() Bool {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Bool(next)
}
func (it *StringQueueIterator) Next() String {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return String(next)
}
func (it *IntQueueIterator) Next() Int {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Int(next)
}
func (it *Int64QueueIterator) Next() Int64 {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Int64(next)
}
func (it *ByteQueueIterator) Next() Byte {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Byte(next)
}
func (it *RuneQueueIterator) Next() Rune {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Rune(next)
}
func (it *Float32QueueIterator) Next() Float32 {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float32(next)
}
func (it *Float64QueueIterator) Next() Float64 {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float64(next)
}
func (it *AnyQueueIterator) Next() Any {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Any(next)
}
func (it *Tuple2QueueIterator) Next() Tuple2 {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Tuple2(next)
}
func (it *BoolArrayQueueIterator) Next() BoolArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return BoolArray(next)
}
func (it *StringArrayQueueIterator) Next() StringArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return StringArray(next)
}
func (it *IntArrayQueueIterator) Next() IntArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return IntArray(next)
}
func (it *Int64ArrayQueueIterator) Next() Int64Array {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Int64Array(next)
}
func (it *ByteArrayQueueIterator) Next() ByteArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return ByteArray(next)
}
func (it *RuneArrayQueueIterator) Next() RuneArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return RuneArray(next)
}
func (it *Float32ArrayQueueIterator) Next() Float32Array {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float32Array(next)
}
func (it *Float64ArrayQueueIterator) Next() Float64Array {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float64Array(next)
}
func (it *AnyArrayQueueIterator) Next() AnyArray {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return AnyArray(next)
}
func (it *Tuple2ArrayQueueIterator) Next() Tuple2Array {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Tuple2Array(next)
}
func (it *BoolOptionQueueIterator) Next() BoolOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return BoolOption(next)
}
func (it *StringOptionQueueIterator) Next() StringOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return StringOption(next)
}
func (it *IntOptionQueueIterator) Next() IntOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return IntOption(next)
}
func (it *Int64OptionQueueIterator) Next() Int64Option {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Int64Option(next)
}
func (it *ByteOptionQueueIterator) Next() ByteOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return ByteOption(next)
}
func (it *RuneOptionQueueIterator) Next() RuneOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return RuneOption(next)
}
func (it *Float32OptionQueueIterator) Next() Float32Option {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float32Option(next)
}
func (it *Float64OptionQueueIterator) Next() Float64Option {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float64Option(next)
}
func (it *AnyOptionQueueIterator) Next() AnyOption {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return AnyOption(next)
}
func (it *Tuple2OptionQueueIterator) Next() Tuple2Option {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Tuple2Option(next)
}
func (it *BoolListQueueIterator) Next() BoolList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return BoolList(next)
}
func (it *StringListQueueIterator) Next() StringList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return StringList(next)
}
func (it *IntListQueueIterator) Next() IntList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return IntList(next)
}
func (it *Int64ListQueueIterator) Next() Int64List {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Int64List(next)
}
func (it *ByteListQueueIterator) Next() ByteList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return ByteList(next)
}
func (it *RuneListQueueIterator) Next() RuneList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return RuneList(next)
}
func (it *Float32ListQueueIterator) Next() Float32List {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float32List(next)
}
func (it *Float64ListQueueIterator) Next() Float64List {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Float64List(next)
}
func (it *AnyListQueueIterator) Next() AnyList {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return AnyList(next)
}
func (it *Tuple2ListQueueIterator) Next() Tuple2List {
	next, t := it.xs.Dequeue()
	it.xs = &t
	return Tuple2List(next)
}
