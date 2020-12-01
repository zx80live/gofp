// bootstrap_list_equals.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (a BoolList) Equals(b BoolList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !BoolEquals(Bool(*xs1.head), Bool(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a StringList) Equals(b StringList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !StringEquals(String(*xs1.head), String(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a IntList) Equals(b IntList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !IntEquals(Int(*xs1.head), Int(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Int64List) Equals(b Int64List) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Int64Equals(Int64(*xs1.head), Int64(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a ByteList) Equals(b ByteList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !ByteEquals(Byte(*xs1.head), Byte(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a RuneList) Equals(b RuneList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !RuneEquals(Rune(*xs1.head), Rune(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float32List) Equals(b Float32List) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float32Equals(Float32(*xs1.head), Float32(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float64List) Equals(b Float64List) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float64Equals(Float64(*xs1.head), Float64(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a AnyList) Equals(b AnyList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !AnyEquals(Any(*xs1.head), Any(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a BoolArrayList) Equals(b BoolArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !BoolArrayEquals(BoolArray(*xs1.head), BoolArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a StringArrayList) Equals(b StringArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !StringArrayEquals(StringArray(*xs1.head), StringArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a IntArrayList) Equals(b IntArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !IntArrayEquals(IntArray(*xs1.head), IntArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Int64ArrayList) Equals(b Int64ArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Int64ArrayEquals(Int64Array(*xs1.head), Int64Array(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a ByteArrayList) Equals(b ByteArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !ByteArrayEquals(ByteArray(*xs1.head), ByteArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a RuneArrayList) Equals(b RuneArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !RuneArrayEquals(RuneArray(*xs1.head), RuneArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float32ArrayList) Equals(b Float32ArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float32ArrayEquals(Float32Array(*xs1.head), Float32Array(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float64ArrayList) Equals(b Float64ArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float64ArrayEquals(Float64Array(*xs1.head), Float64Array(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a AnyArrayList) Equals(b AnyArrayList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !AnyArrayEquals(AnyArray(*xs1.head), AnyArray(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a BoolOptionList) Equals(b BoolOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !BoolOptionEquals(BoolOption(*xs1.head), BoolOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a StringOptionList) Equals(b StringOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !StringOptionEquals(StringOption(*xs1.head), StringOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a IntOptionList) Equals(b IntOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !IntOptionEquals(IntOption(*xs1.head), IntOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Int64OptionList) Equals(b Int64OptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Int64OptionEquals(Int64Option(*xs1.head), Int64Option(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a ByteOptionList) Equals(b ByteOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !ByteOptionEquals(ByteOption(*xs1.head), ByteOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a RuneOptionList) Equals(b RuneOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !RuneOptionEquals(RuneOption(*xs1.head), RuneOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float32OptionList) Equals(b Float32OptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float32OptionEquals(Float32Option(*xs1.head), Float32Option(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float64OptionList) Equals(b Float64OptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float64OptionEquals(Float64Option(*xs1.head), Float64Option(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a AnyOptionList) Equals(b AnyOptionList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !AnyOptionEquals(AnyOption(*xs1.head), AnyOption(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a BoolListList) Equals(b BoolListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !BoolListEquals(BoolList(*xs1.head), BoolList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a StringListList) Equals(b StringListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !StringListEquals(StringList(*xs1.head), StringList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a IntListList) Equals(b IntListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !IntListEquals(IntList(*xs1.head), IntList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Int64ListList) Equals(b Int64ListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Int64ListEquals(Int64List(*xs1.head), Int64List(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a ByteListList) Equals(b ByteListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !ByteListEquals(ByteList(*xs1.head), ByteList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a RuneListList) Equals(b RuneListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !RuneListEquals(RuneList(*xs1.head), RuneList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float32ListList) Equals(b Float32ListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float32ListEquals(Float32List(*xs1.head), Float32List(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a Float64ListList) Equals(b Float64ListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !Float64ListEquals(Float64List(*xs1.head), Float64List(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
func (a AnyListList) Equals(b AnyListList) bool {
	len1 := a.Size()
	len2 := b.Size()
	if len1 != len2 {
		return false
	}
	xs1 := a
	xs2 := b
	for xs1.NonEmpty() {
		if !AnyListEquals(AnyList(*xs1.head), AnyList(*xs2.head)) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
