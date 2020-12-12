// bootstrap_list_reverse.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolList) Reverse() BoolList {
	acc := NilBoolList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l StringList) Reverse() StringList {
	acc := NilStringList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l IntList) Reverse() IntList {
	acc := NilIntList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Int64List) Reverse() Int64List {
	acc := NilInt64List
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l ByteList) Reverse() ByteList {
	acc := NilByteList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l RuneList) Reverse() RuneList {
	acc := NilRuneList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float32List) Reverse() Float32List {
	acc := NilFloat32List
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float64List) Reverse() Float64List {
	acc := NilFloat64List
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l AnyList) Reverse() AnyList {
	acc := NilAnyList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Tuple2List) Reverse() Tuple2List {
	acc := NilTuple2List
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l BoolArrayList) Reverse() BoolArrayList {
	acc := NilBoolArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l StringArrayList) Reverse() StringArrayList {
	acc := NilStringArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l IntArrayList) Reverse() IntArrayList {
	acc := NilIntArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Int64ArrayList) Reverse() Int64ArrayList {
	acc := NilInt64ArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l ByteArrayList) Reverse() ByteArrayList {
	acc := NilByteArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l RuneArrayList) Reverse() RuneArrayList {
	acc := NilRuneArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float32ArrayList) Reverse() Float32ArrayList {
	acc := NilFloat32ArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float64ArrayList) Reverse() Float64ArrayList {
	acc := NilFloat64ArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l AnyArrayList) Reverse() AnyArrayList {
	acc := NilAnyArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Tuple2ArrayList) Reverse() Tuple2ArrayList {
	acc := NilTuple2ArrayList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l BoolOptionList) Reverse() BoolOptionList {
	acc := NilBoolOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l StringOptionList) Reverse() StringOptionList {
	acc := NilStringOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l IntOptionList) Reverse() IntOptionList {
	acc := NilIntOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Int64OptionList) Reverse() Int64OptionList {
	acc := NilInt64OptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l ByteOptionList) Reverse() ByteOptionList {
	acc := NilByteOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l RuneOptionList) Reverse() RuneOptionList {
	acc := NilRuneOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float32OptionList) Reverse() Float32OptionList {
	acc := NilFloat32OptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float64OptionList) Reverse() Float64OptionList {
	acc := NilFloat64OptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l AnyOptionList) Reverse() AnyOptionList {
	acc := NilAnyOptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Tuple2OptionList) Reverse() Tuple2OptionList {
	acc := NilTuple2OptionList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l BoolListList) Reverse() BoolListList {
	acc := NilBoolListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l StringListList) Reverse() StringListList {
	acc := NilStringListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l IntListList) Reverse() IntListList {
	acc := NilIntListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Int64ListList) Reverse() Int64ListList {
	acc := NilInt64ListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l ByteListList) Reverse() ByteListList {
	acc := NilByteListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l RuneListList) Reverse() RuneListList {
	acc := NilRuneListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float32ListList) Reverse() Float32ListList {
	acc := NilFloat32ListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Float64ListList) Reverse() Float64ListList {
	acc := NilFloat64ListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l AnyListList) Reverse() AnyListList {
	acc := NilAnyListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
func (l Tuple2ListList) Reverse() Tuple2ListList {
	acc := NilTuple2ListList
	xs := l
	for xs.NonEmpty() {
		acc = acc.Cons(*xs.head)
		xs = *xs.tail
	}
	return acc
}
