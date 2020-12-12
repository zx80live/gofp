// bootstrap_queue_filter.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (q BoolQueue) Filter(p BoolPredicate) BoolQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return BoolQueue{&in, &out}
}
func (q StringQueue) Filter(p StringPredicate) StringQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return StringQueue{&in, &out}
}
func (q IntQueue) Filter(p IntPredicate) IntQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return IntQueue{&in, &out}
}
func (q Int64Queue) Filter(p Int64Predicate) Int64Queue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Int64Queue{&in, &out}
}
func (q ByteQueue) Filter(p BytePredicate) ByteQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return ByteQueue{&in, &out}
}
func (q RuneQueue) Filter(p RunePredicate) RuneQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return RuneQueue{&in, &out}
}
func (q Float32Queue) Filter(p Float32Predicate) Float32Queue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float32Queue{&in, &out}
}
func (q Float64Queue) Filter(p Float64Predicate) Float64Queue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float64Queue{&in, &out}
}
func (q AnyQueue) Filter(p AnyPredicate) AnyQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return AnyQueue{&in, &out}
}
func (q Tuple2Queue) Filter(p Tuple2Predicate) Tuple2Queue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Tuple2Queue{&in, &out}
}
func (q BoolArrayQueue) Filter(p BoolArrayPredicate) BoolArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return BoolArrayQueue{&in, &out}
}
func (q StringArrayQueue) Filter(p StringArrayPredicate) StringArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return StringArrayQueue{&in, &out}
}
func (q IntArrayQueue) Filter(p IntArrayPredicate) IntArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return IntArrayQueue{&in, &out}
}
func (q Int64ArrayQueue) Filter(p Int64ArrayPredicate) Int64ArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Int64ArrayQueue{&in, &out}
}
func (q ByteArrayQueue) Filter(p ByteArrayPredicate) ByteArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return ByteArrayQueue{&in, &out}
}
func (q RuneArrayQueue) Filter(p RuneArrayPredicate) RuneArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return RuneArrayQueue{&in, &out}
}
func (q Float32ArrayQueue) Filter(p Float32ArrayPredicate) Float32ArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float32ArrayQueue{&in, &out}
}
func (q Float64ArrayQueue) Filter(p Float64ArrayPredicate) Float64ArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float64ArrayQueue{&in, &out}
}
func (q AnyArrayQueue) Filter(p AnyArrayPredicate) AnyArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return AnyArrayQueue{&in, &out}
}
func (q Tuple2ArrayQueue) Filter(p Tuple2ArrayPredicate) Tuple2ArrayQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Tuple2ArrayQueue{&in, &out}
}
func (q BoolOptionQueue) Filter(p BoolOptionPredicate) BoolOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return BoolOptionQueue{&in, &out}
}
func (q StringOptionQueue) Filter(p StringOptionPredicate) StringOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return StringOptionQueue{&in, &out}
}
func (q IntOptionQueue) Filter(p IntOptionPredicate) IntOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return IntOptionQueue{&in, &out}
}
func (q Int64OptionQueue) Filter(p Int64OptionPredicate) Int64OptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Int64OptionQueue{&in, &out}
}
func (q ByteOptionQueue) Filter(p ByteOptionPredicate) ByteOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return ByteOptionQueue{&in, &out}
}
func (q RuneOptionQueue) Filter(p RuneOptionPredicate) RuneOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return RuneOptionQueue{&in, &out}
}
func (q Float32OptionQueue) Filter(p Float32OptionPredicate) Float32OptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float32OptionQueue{&in, &out}
}
func (q Float64OptionQueue) Filter(p Float64OptionPredicate) Float64OptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float64OptionQueue{&in, &out}
}
func (q AnyOptionQueue) Filter(p AnyOptionPredicate) AnyOptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return AnyOptionQueue{&in, &out}
}
func (q Tuple2OptionQueue) Filter(p Tuple2OptionPredicate) Tuple2OptionQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Tuple2OptionQueue{&in, &out}
}
func (q BoolListQueue) Filter(p BoolListPredicate) BoolListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return BoolListQueue{&in, &out}
}
func (q StringListQueue) Filter(p StringListPredicate) StringListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return StringListQueue{&in, &out}
}
func (q IntListQueue) Filter(p IntListPredicate) IntListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return IntListQueue{&in, &out}
}
func (q Int64ListQueue) Filter(p Int64ListPredicate) Int64ListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Int64ListQueue{&in, &out}
}
func (q ByteListQueue) Filter(p ByteListPredicate) ByteListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return ByteListQueue{&in, &out}
}
func (q RuneListQueue) Filter(p RuneListPredicate) RuneListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return RuneListQueue{&in, &out}
}
func (q Float32ListQueue) Filter(p Float32ListPredicate) Float32ListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float32ListQueue{&in, &out}
}
func (q Float64ListQueue) Filter(p Float64ListPredicate) Float64ListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Float64ListQueue{&in, &out}
}
func (q AnyListQueue) Filter(p AnyListPredicate) AnyListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return AnyListQueue{&in, &out}
}
func (q Tuple2ListQueue) Filter(p Tuple2ListPredicate) Tuple2ListQueue {
	in := (*q.in).Filter(p)
	out := (*q.out).Filter(p)
	return Tuple2ListQueue{&in, &out}
}
