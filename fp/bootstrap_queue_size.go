// bootstrap_queue_size.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (q BoolQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q StringQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q IntQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Int64Queue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q ByteQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q RuneQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float32Queue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float64Queue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q AnyQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Tuple2Queue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q BoolArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q StringArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q IntArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Int64ArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q ByteArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q RuneArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float32ArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float64ArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q AnyArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Tuple2ArrayQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q BoolOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q StringOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q IntOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Int64OptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q ByteOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q RuneOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float32OptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float64OptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q AnyOptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Tuple2OptionQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q BoolListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q StringListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q IntListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Int64ListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q ByteListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q RuneListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float32ListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Float64ListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q AnyListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}
func (q Tuple2ListQueue) Size() int {
	count := 0
	xs := q
	for xs.NonEmpty() {
		count++
		_, xs = xs.Dequeue()
	}
	return count
}