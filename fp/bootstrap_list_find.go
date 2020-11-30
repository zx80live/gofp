// bootstrap_list_find.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolList) Find(p func(bool) bool) BoolOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeBoolOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneBool
}
func (l StringList) Find(p func(string) bool) StringOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeStringOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneString
}
func (l IntList) Find(p func(int) bool) IntOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeIntOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneInt
}
func (l Int64List) Find(p func(int64) bool) Int64Option {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeInt64Option(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneInt64
}
func (l ByteList) Find(p func(byte) bool) ByteOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeByteOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneByte
}
func (l RuneList) Find(p func(rune) bool) RuneOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeRuneOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneRune
}
func (l Float32List) Find(p func(float32) bool) Float32Option {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat32Option(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat32
}
func (l Float64List) Find(p func(float64) bool) Float64Option {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat64Option(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat64
}
func (l AnyList) Find(p func(Any) bool) AnyOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeAnyOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneAny
}
func (l BoolArrayList) Find(p func([]bool) bool) BoolArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeBoolArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneBoolArray
}
func (l StringArrayList) Find(p func([]string) bool) StringArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeStringArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneStringArray
}
func (l IntArrayList) Find(p func([]int) bool) IntArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeIntArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneIntArray
}
func (l Int64ArrayList) Find(p func([]int64) bool) Int64ArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeInt64ArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneInt64Array
}
func (l ByteArrayList) Find(p func([]byte) bool) ByteArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeByteArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneByteArray
}
func (l RuneArrayList) Find(p func([]rune) bool) RuneArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeRuneArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneRuneArray
}
func (l Float32ArrayList) Find(p func([]float32) bool) Float32ArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat32ArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat32Array
}
func (l Float64ArrayList) Find(p func([]float64) bool) Float64ArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat64ArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat64Array
}
func (l AnyArrayList) Find(p func([]Any) bool) AnyArrayOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeAnyArrayOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneAnyArray
}
func (l BoolOptionList) Find(p func(BoolOption) bool) BoolOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeBoolOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneBoolOption
}
func (l StringOptionList) Find(p func(StringOption) bool) StringOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeStringOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneStringOption
}
func (l IntOptionList) Find(p func(IntOption) bool) IntOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeIntOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneIntOption
}
func (l Int64OptionList) Find(p func(Int64Option) bool) Int64OptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeInt64OptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneInt64Option
}
func (l ByteOptionList) Find(p func(ByteOption) bool) ByteOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeByteOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneByteOption
}
func (l RuneOptionList) Find(p func(RuneOption) bool) RuneOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeRuneOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneRuneOption
}
func (l Float32OptionList) Find(p func(Float32Option) bool) Float32OptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat32OptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat32Option
}
func (l Float64OptionList) Find(p func(Float64Option) bool) Float64OptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat64OptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat64Option
}
func (l AnyOptionList) Find(p func(AnyOption) bool) AnyOptionOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeAnyOptionOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneAnyOption
}
func (l BoolListList) Find(p func(BoolList) bool) BoolListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeBoolListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneBoolList
}
func (l StringListList) Find(p func(StringList) bool) StringListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeStringListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneStringList
}
func (l IntListList) Find(p func(IntList) bool) IntListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeIntListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneIntList
}
func (l Int64ListList) Find(p func(Int64List) bool) Int64ListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeInt64ListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneInt64List
}
func (l ByteListList) Find(p func(ByteList) bool) ByteListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeByteListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneByteList
}
func (l RuneListList) Find(p func(RuneList) bool) RuneListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeRuneListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneRuneList
}
func (l Float32ListList) Find(p func(Float32List) bool) Float32ListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat32ListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat32List
}
func (l Float64ListList) Find(p func(Float64List) bool) Float64ListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeFloat64ListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneFloat64List
}
func (l AnyListList) Find(p func(AnyList) bool) AnyListOption {
	xs := l
	for xs.NonEmpty() {
		if p(*xs.head) {
			return MakeAnyListOption(*xs.head)
		}
		xs = *xs.tail
	}
	return NoneAnyList
}
