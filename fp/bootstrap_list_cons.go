// bootstrap_list_cons.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func MkBoolList(elements ...bool) BoolList {
	xs := NilBool
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkStringList(elements ...string) StringList {
	xs := NilString
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkIntList(elements ...int) IntList {
	xs := NilInt
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkInt64List(elements ...int64) Int64List {
	xs := NilInt64
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkByteList(elements ...byte) ByteList {
	xs := NilByte
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkRuneList(elements ...rune) RuneList {
	xs := NilRune
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat32List(elements ...float32) Float32List {
	xs := NilFloat32
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat64List(elements ...float64) Float64List {
	xs := NilFloat64
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkAnyList(elements ...Any) AnyList {
	xs := NilAny
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkTuple2List(elements ...Tuple2) Tuple2List {
	xs := NilTuple2
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkBoolArrayList(elements ...[]bool) BoolArrayList {
	xs := NilBoolArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkStringArrayList(elements ...[]string) StringArrayList {
	xs := NilStringArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkIntArrayList(elements ...[]int) IntArrayList {
	xs := NilIntArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkInt64ArrayList(elements ...[]int64) Int64ArrayList {
	xs := NilInt64Array
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkByteArrayList(elements ...[]byte) ByteArrayList {
	xs := NilByteArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkRuneArrayList(elements ...[]rune) RuneArrayList {
	xs := NilRuneArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat32ArrayList(elements ...[]float32) Float32ArrayList {
	xs := NilFloat32Array
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat64ArrayList(elements ...[]float64) Float64ArrayList {
	xs := NilFloat64Array
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkAnyArrayList(elements ...[]Any) AnyArrayList {
	xs := NilAnyArray
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkTuple2ArrayList(elements ...[]Tuple2) Tuple2ArrayList {
	xs := NilTuple2Array
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkBoolOptionList(elements ...BoolOption) BoolOptionList {
	xs := NilBoolOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkStringOptionList(elements ...StringOption) StringOptionList {
	xs := NilStringOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkIntOptionList(elements ...IntOption) IntOptionList {
	xs := NilIntOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkInt64OptionList(elements ...Int64Option) Int64OptionList {
	xs := NilInt64Option
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkByteOptionList(elements ...ByteOption) ByteOptionList {
	xs := NilByteOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkRuneOptionList(elements ...RuneOption) RuneOptionList {
	xs := NilRuneOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat32OptionList(elements ...Float32Option) Float32OptionList {
	xs := NilFloat32Option
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat64OptionList(elements ...Float64Option) Float64OptionList {
	xs := NilFloat64Option
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkAnyOptionList(elements ...AnyOption) AnyOptionList {
	xs := NilAnyOption
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkTuple2OptionList(elements ...Tuple2Option) Tuple2OptionList {
	xs := NilTuple2Option
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkBoolListList(elements ...BoolList) BoolListList {
	xs := NilBoolList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkStringListList(elements ...StringList) StringListList {
	xs := NilStringList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkIntListList(elements ...IntList) IntListList {
	xs := NilIntList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkInt64ListList(elements ...Int64List) Int64ListList {
	xs := NilInt64List
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkByteListList(elements ...ByteList) ByteListList {
	xs := NilByteList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkRuneListList(elements ...RuneList) RuneListList {
	xs := NilRuneList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat32ListList(elements ...Float32List) Float32ListList {
	xs := NilFloat32List
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkFloat64ListList(elements ...Float64List) Float64ListList {
	xs := NilFloat64List
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkAnyListList(elements ...AnyList) AnyListList {
	xs := NilAnyList
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
func MkTuple2ListList(elements ...Tuple2List) Tuple2ListList {
	xs := NilTuple2List
	for _, e := range elements {
		xs = xs.Cons(e)
	}
	return xs.Reverse()
}
