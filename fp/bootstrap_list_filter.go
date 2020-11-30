// bootstrap_list_filter.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolList) Filter(p BoolPredicate) BoolList {
	acc := NilBool
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringList) Filter(p StringPredicate) StringList {
	acc := NilString
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntList) Filter(p IntPredicate) IntList {
	acc := NilInt
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64List) Filter(p Int64Predicate) Int64List {
	acc := NilInt64
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteList) Filter(p BytePredicate) ByteList {
	acc := NilByte
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneList) Filter(p RunePredicate) RuneList {
	acc := NilRune
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32List) Filter(p Float32Predicate) Float32List {
	acc := NilFloat32
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64List) Filter(p Float64Predicate) Float64List {
	acc := NilFloat64
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyList) Filter(p AnyPredicate) AnyList {
	acc := NilAny
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolArrayList) Filter(p BoolArrayPredicate) BoolArrayList {
	acc := NilBoolArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringArrayList) Filter(p StringArrayPredicate) StringArrayList {
	acc := NilStringArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntArrayList) Filter(p IntArrayPredicate) IntArrayList {
	acc := NilIntArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64ArrayList) Filter(p Int64ArrayPredicate) Int64ArrayList {
	acc := NilInt64Array
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteArrayList) Filter(p ByteArrayPredicate) ByteArrayList {
	acc := NilByteArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneArrayList) Filter(p RuneArrayPredicate) RuneArrayList {
	acc := NilRuneArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32ArrayList) Filter(p Float32ArrayPredicate) Float32ArrayList {
	acc := NilFloat32Array
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64ArrayList) Filter(p Float64ArrayPredicate) Float64ArrayList {
	acc := NilFloat64Array
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyArrayList) Filter(p AnyArrayPredicate) AnyArrayList {
	acc := NilAnyArray
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolOptionList) Filter(p BoolOptionPredicate) BoolOptionList {
	acc := NilBoolOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringOptionList) Filter(p StringOptionPredicate) StringOptionList {
	acc := NilStringOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntOptionList) Filter(p IntOptionPredicate) IntOptionList {
	acc := NilIntOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64OptionList) Filter(p Int64OptionPredicate) Int64OptionList {
	acc := NilInt64Option
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteOptionList) Filter(p ByteOptionPredicate) ByteOptionList {
	acc := NilByteOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneOptionList) Filter(p RuneOptionPredicate) RuneOptionList {
	acc := NilRuneOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32OptionList) Filter(p Float32OptionPredicate) Float32OptionList {
	acc := NilFloat32Option
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64OptionList) Filter(p Float64OptionPredicate) Float64OptionList {
	acc := NilFloat64Option
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyOptionList) Filter(p AnyOptionPredicate) AnyOptionList {
	acc := NilAnyOption
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l BoolListList) Filter(p BoolListPredicate) BoolListList {
	acc := NilBoolList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l StringListList) Filter(p StringListPredicate) StringListList {
	acc := NilStringList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l IntListList) Filter(p IntListPredicate) IntListList {
	acc := NilIntList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Int64ListList) Filter(p Int64ListPredicate) Int64ListList {
	acc := NilInt64List
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l ByteListList) Filter(p ByteListPredicate) ByteListList {
	acc := NilByteList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l RuneListList) Filter(p RuneListPredicate) RuneListList {
	acc := NilRuneList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float32ListList) Filter(p Float32ListPredicate) Float32ListList {
	acc := NilFloat32List
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l Float64ListList) Filter(p Float64ListPredicate) Float64ListList {
	acc := NilFloat64List
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
func (l AnyListList) Filter(p AnyListPredicate) AnyListList {
	acc := NilAnyList
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			acc = acc.Cons(*xs.head)
		}
		xs = *xs.tail
	}
	return acc.Reverse()
}
