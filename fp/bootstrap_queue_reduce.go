// bootstrap_queue_reduce.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (q BoolQueue) Reduce(f func(bool, bool) bool) bool {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q StringQueue) Reduce(f func(string, string) string) string {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q IntQueue) Reduce(f func(int, int) int) int {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Int64Queue) Reduce(f func(int64, int64) int64) int64 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q ByteQueue) Reduce(f func(byte, byte) byte) byte {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q RuneQueue) Reduce(f func(rune, rune) rune) rune {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float32Queue) Reduce(f func(float32, float32) float32) float32 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float64Queue) Reduce(f func(float64, float64) float64) float64 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q AnyQueue) Reduce(f func(Any, Any) Any) Any {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Tuple2Queue) Reduce(f func(Tuple2, Tuple2) Tuple2) Tuple2 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q BoolArrayQueue) Reduce(f func([]bool, []bool) []bool) []bool {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q StringArrayQueue) Reduce(f func([]string, []string) []string) []string {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q IntArrayQueue) Reduce(f func([]int, []int) []int) []int {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Int64ArrayQueue) Reduce(f func([]int64, []int64) []int64) []int64 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q ByteArrayQueue) Reduce(f func([]byte, []byte) []byte) []byte {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q RuneArrayQueue) Reduce(f func([]rune, []rune) []rune) []rune {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float32ArrayQueue) Reduce(f func([]float32, []float32) []float32) []float32 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float64ArrayQueue) Reduce(f func([]float64, []float64) []float64) []float64 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q AnyArrayQueue) Reduce(f func([]Any, []Any) []Any) []Any {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Tuple2ArrayQueue) Reduce(f func([]Tuple2, []Tuple2) []Tuple2) []Tuple2 {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q BoolOptionQueue) Reduce(f func(BoolOption, BoolOption) BoolOption) BoolOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q StringOptionQueue) Reduce(f func(StringOption, StringOption) StringOption) StringOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q IntOptionQueue) Reduce(f func(IntOption, IntOption) IntOption) IntOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Int64OptionQueue) Reduce(f func(Int64Option, Int64Option) Int64Option) Int64Option {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q ByteOptionQueue) Reduce(f func(ByteOption, ByteOption) ByteOption) ByteOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q RuneOptionQueue) Reduce(f func(RuneOption, RuneOption) RuneOption) RuneOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float32OptionQueue) Reduce(f func(Float32Option, Float32Option) Float32Option) Float32Option {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float64OptionQueue) Reduce(f func(Float64Option, Float64Option) Float64Option) Float64Option {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q AnyOptionQueue) Reduce(f func(AnyOption, AnyOption) AnyOption) AnyOption {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Tuple2OptionQueue) Reduce(f func(Tuple2Option, Tuple2Option) Tuple2Option) Tuple2Option {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q BoolListQueue) Reduce(f func(BoolList, BoolList) BoolList) BoolList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q StringListQueue) Reduce(f func(StringList, StringList) StringList) StringList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q IntListQueue) Reduce(f func(IntList, IntList) IntList) IntList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Int64ListQueue) Reduce(f func(Int64List, Int64List) Int64List) Int64List {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q ByteListQueue) Reduce(f func(ByteList, ByteList) ByteList) ByteList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q RuneListQueue) Reduce(f func(RuneList, RuneList) RuneList) RuneList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float32ListQueue) Reduce(f func(Float32List, Float32List) Float32List) Float32List {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Float64ListQueue) Reduce(f func(Float64List, Float64List) Float64List) Float64List {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q AnyListQueue) Reduce(f func(AnyList, AnyList) AnyList) AnyList {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
func (q Tuple2ListQueue) Reduce(f func(Tuple2List, Tuple2List) Tuple2List) Tuple2List {
	if q.IsEmpty() {
		panic("Can't reduce empty queue")
	} else {
		h, t := q.Dequeue()
		if t.IsEmpty() {
			return h
		}
		return f(h, t.Reduce(f))
	}
}
