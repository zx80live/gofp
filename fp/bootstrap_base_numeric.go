// bootstrap_base_numeric.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (a Int) Min(b Int) Int {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Int64) Min(b Int64) Int64 {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Byte) Min(b Byte) Byte {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Rune) Min(b Rune) Rune {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Float32) Min(b Float32) Float32 {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Float64) Min(b Float64) Float64 {
	if a <= b {
		return a
	} else {
		return b
	}
}
func (a Int) Max(b Int) Int {
	if a > b {
		return a
	} else {
		return b
	}
}
func (a Int64) Max(b Int64) Int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func (a Byte) Max(b Byte) Byte {
	if a > b {
		return a
	} else {
		return b
	}
}
func (a Rune) Max(b Rune) Rune {
	if a > b {
		return a
	} else {
		return b
	}
}
func (a Float32) Max(b Float32) Float32 {
	if a > b {
		return a
	} else {
		return b
	}
}
func (a Float64) Max(b Float64) Float64 {
	if a > b {
		return a
	} else {
		return b
	}
}
func (n Int) To(t Int) IntList {
	acc := NilIntList
	for i := n.Underlined(); i <= t.Underlined(); i++ {
		acc = acc.Cons(i)
	}
	return acc.Reverse()
}
func (n Byte) To(t Byte) ByteList {
	acc := NilByteList
	for i := n.Underlined(); i <= t.Underlined(); i++ {
		acc = acc.Cons(i)
	}
	return acc.Reverse()
}
func (n Int) Until(t Int) IntList {
	acc := NilIntList
	for i := n.Underlined(); i < t.Underlined(); i++ {
		acc = acc.Cons(i)
	}
	return acc.Reverse()
}
func (n Byte) Until(t Byte) ByteList {
	acc := NilByteList
	for i := n.Underlined(); i < t.Underlined(); i++ {
		acc = acc.Cons(i)
	}
	return acc.Reverse()
}
func (a Int) IsBetween(left, right int) bool            { return int(a) > left && int(a) < right }
func (a Int8) IsBetween(left, right int8) bool          { return int8(a) > left && int8(a) < right }
func (a Int16) IsBetween(left, right int16) bool        { return int16(a) > left && int16(a) < right }
func (a Int32) IsBetween(left, right int32) bool        { return int32(a) > left && int32(a) < right }
func (a Int64) IsBetween(left, right int64) bool        { return int64(a) > left && int64(a) < right }
func (a Uint) IsBetween(left, right uint) bool          { return uint(a) > left && uint(a) < right }
func (a Uint8) IsBetween(left, right uint8) bool        { return uint8(a) > left && uint8(a) < right }
func (a Uint16) IsBetween(left, right uint16) bool      { return uint16(a) > left && uint16(a) < right }
func (a Uint32) IsBetween(left, right uint32) bool      { return uint32(a) > left && uint32(a) < right }
func (a Uint64) IsBetween(left, right uint64) bool      { return uint64(a) > left && uint64(a) < right }
func (a Uintptr) IsBetween(left, right uintptr) bool    { return uintptr(a) > left && uintptr(a) < right }
func (a Byte) IsBetween(left, right byte) bool          { return byte(a) > left && byte(a) < right }
func (a Rune) IsBetween(left, right rune) bool          { return rune(a) > left && rune(a) < right }
func (a Float32) IsBetween(left, right float32) bool    { return float32(a) > left && float32(a) < right }
func (a Float64) IsBetween(left, right float64) bool    { return float64(a) > left && float64(a) < right }
func (a Int) IsBetweenInclusive(left, right int) bool   { return int(a) >= left && int(a) <= right }
func (a Int8) IsBetweenInclusive(left, right int8) bool { return int8(a) >= left && int8(a) <= right }
func (a Int16) IsBetweenInclusive(left, right int16) bool {
	return int16(a) >= left && int16(a) <= right
}
func (a Int32) IsBetweenInclusive(left, right int32) bool {
	return int32(a) >= left && int32(a) <= right
}
func (a Int64) IsBetweenInclusive(left, right int64) bool {
	return int64(a) >= left && int64(a) <= right
}
func (a Uint) IsBetweenInclusive(left, right uint) bool { return uint(a) >= left && uint(a) <= right }
func (a Uint8) IsBetweenInclusive(left, right uint8) bool {
	return uint8(a) >= left && uint8(a) <= right
}
func (a Uint16) IsBetweenInclusive(left, right uint16) bool {
	return uint16(a) >= left && uint16(a) <= right
}
func (a Uint32) IsBetweenInclusive(left, right uint32) bool {
	return uint32(a) >= left && uint32(a) <= right
}
func (a Uint64) IsBetweenInclusive(left, right uint64) bool {
	return uint64(a) >= left && uint64(a) <= right
}
func (a Uintptr) IsBetweenInclusive(left, right uintptr) bool {
	return uintptr(a) >= left && uintptr(a) <= right
}
func (a Byte) IsBetweenInclusive(left, right byte) bool { return byte(a) >= left && byte(a) <= right }
func (a Rune) IsBetweenInclusive(left, right rune) bool { return rune(a) >= left && rune(a) <= right }
func (a Float32) IsBetweenInclusive(left, right float32) bool {
	return float32(a) >= left && float32(a) <= right
}
func (a Float64) IsBetweenInclusive(left, right float64) bool {
	return float64(a) >= left && float64(a) <= right
}
