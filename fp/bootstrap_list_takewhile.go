// bootstrap_list_takewhile.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolList) TakeWhile(p func(bool) bool) BoolList {
	acc := NilBool
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringList) TakeWhile(p func(string) bool) StringList {
	acc := NilString
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntList) TakeWhile(p func(int) bool) IntList {
	acc := NilInt
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64List) TakeWhile(p func(int64) bool) Int64List {
	acc := NilInt64
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteList) TakeWhile(p func(byte) bool) ByteList {
	acc := NilByte
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneList) TakeWhile(p func(rune) bool) RuneList {
	acc := NilRune
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32List) TakeWhile(p func(float32) bool) Float32List {
	acc := NilFloat32
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64List) TakeWhile(p func(float64) bool) Float64List {
	acc := NilFloat64
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyList) TakeWhile(p func(Any) bool) AnyList {
	acc := NilAny
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Tuple2List) TakeWhile(p func(Tuple2) bool) Tuple2List {
	acc := NilTuple2
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolArrayList) TakeWhile(p func([]bool) bool) BoolArrayList {
	acc := NilBoolArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringArrayList) TakeWhile(p func([]string) bool) StringArrayList {
	acc := NilStringArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntArrayList) TakeWhile(p func([]int) bool) IntArrayList {
	acc := NilIntArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64ArrayList) TakeWhile(p func([]int64) bool) Int64ArrayList {
	acc := NilInt64Array
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteArrayList) TakeWhile(p func([]byte) bool) ByteArrayList {
	acc := NilByteArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneArrayList) TakeWhile(p func([]rune) bool) RuneArrayList {
	acc := NilRuneArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32ArrayList) TakeWhile(p func([]float32) bool) Float32ArrayList {
	acc := NilFloat32Array
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64ArrayList) TakeWhile(p func([]float64) bool) Float64ArrayList {
	acc := NilFloat64Array
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyArrayList) TakeWhile(p func([]Any) bool) AnyArrayList {
	acc := NilAnyArray
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Tuple2ArrayList) TakeWhile(p func([]Tuple2) bool) Tuple2ArrayList {
	acc := NilTuple2Array
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolOptionList) TakeWhile(p func(BoolOption) bool) BoolOptionList {
	acc := NilBoolOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringOptionList) TakeWhile(p func(StringOption) bool) StringOptionList {
	acc := NilStringOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntOptionList) TakeWhile(p func(IntOption) bool) IntOptionList {
	acc := NilIntOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64OptionList) TakeWhile(p func(Int64Option) bool) Int64OptionList {
	acc := NilInt64Option
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteOptionList) TakeWhile(p func(ByteOption) bool) ByteOptionList {
	acc := NilByteOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneOptionList) TakeWhile(p func(RuneOption) bool) RuneOptionList {
	acc := NilRuneOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32OptionList) TakeWhile(p func(Float32Option) bool) Float32OptionList {
	acc := NilFloat32Option
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64OptionList) TakeWhile(p func(Float64Option) bool) Float64OptionList {
	acc := NilFloat64Option
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyOptionList) TakeWhile(p func(AnyOption) bool) AnyOptionList {
	acc := NilAnyOption
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Tuple2OptionList) TakeWhile(p func(Tuple2Option) bool) Tuple2OptionList {
	acc := NilTuple2Option
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolListList) TakeWhile(p func(BoolList) bool) BoolListList {
	acc := NilBoolList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringListList) TakeWhile(p func(StringList) bool) StringListList {
	acc := NilStringList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntListList) TakeWhile(p func(IntList) bool) IntListList {
	acc := NilIntList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64ListList) TakeWhile(p func(Int64List) bool) Int64ListList {
	acc := NilInt64List
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteListList) TakeWhile(p func(ByteList) bool) ByteListList {
	acc := NilByteList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneListList) TakeWhile(p func(RuneList) bool) RuneListList {
	acc := NilRuneList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32ListList) TakeWhile(p func(Float32List) bool) Float32ListList {
	acc := NilFloat32List
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64ListList) TakeWhile(p func(Float64List) bool) Float64ListList {
	acc := NilFloat64List
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyListList) TakeWhile(p func(AnyList) bool) AnyListList {
	acc := NilAnyList
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Tuple2ListList) TakeWhile(p func(Tuple2List) bool) Tuple2ListList {
	acc := NilTuple2List
	xs := l
	for xs.NonEmpty() && p(*xs.head) {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc.Reverse()
}
