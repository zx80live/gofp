// fequal.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

func EqualBool(a, b bool) bool             { return a == b }
func EqualRune(a, b rune) bool             { return a == b }
func EqualByte(a, b byte) bool             { return a == b }
func EqualInt(a, b int) bool               { return a == b }
func EqualInt8(a, b int8) bool             { return a == b }
func EqualInt16(a, b int16) bool           { return a == b }
func EqualInt32(a, b int32) bool           { return a == b }
func EqualInt64(a, b int64) bool           { return a == b }
func EqualUint(a, b uint) bool             { return a == b }
func EqualUint8(a, b uint8) bool           { return a == b }
func EqualUint16(a, b uint16) bool         { return a == b }
func EqualUint32(a, b uint32) bool         { return a == b }
func EqualUint64(a, b uint64) bool         { return a == b }
func EqualUintptr(a, b uintptr) bool       { return a == b }
func EqualFloat32(a, b float32) bool       { return a == b }
func EqualFloat64(a, b float64) bool       { return a == b }
func EqualComplex64(a, b complex64) bool   { return a == b }
func EqualComplex128(a, b complex128) bool { return a == b }
func EqualString(a, b string) bool         { return a == b }
func EqualAny(a, b Any) bool               { return a == b }
func EqualBoolArr(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualRuneArr(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualByteArr(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualIntArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualInt8Arr(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualInt16Arr(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualInt32Arr(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualInt64Arr(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUintArr(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUint8Arr(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUint16Arr(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUint32Arr(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUint64Arr(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualUintptrArr(a, b []uintptr) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualFloat32Arr(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualFloat64Arr(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualComplex64Arr(a, b []complex64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualComplex128Arr(a, b []complex128) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualStringArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualAnyArr(a, b []Any) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func EqualOptionBool(a, b OptionBool) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualBool(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionRune(a, b OptionRune) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualRune(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionByte(a, b OptionByte) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualByte(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt(a, b OptionInt) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt8(a, b OptionInt8) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt8(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt16(a, b OptionInt16) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt16(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt32(a, b OptionInt32) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt32(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt64(a, b OptionInt64) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt64(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint(a, b OptionUint) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint8(a, b OptionUint8) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint8(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint16(a, b OptionUint16) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint16(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint32(a, b OptionUint32) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint32(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint64(a, b OptionUint64) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint64(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUintptr(a, b OptionUintptr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUintptr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionFloat32(a, b OptionFloat32) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualFloat32(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionFloat64(a, b OptionFloat64) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualFloat64(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionComplex64(a, b OptionComplex64) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualComplex64(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionComplex128(a, b OptionComplex128) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualComplex128(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionString(a, b OptionString) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualString(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionAny(a, b OptionAny) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualAny(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionBoolArr(a, b OptionBoolArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualBoolArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionRuneArr(a, b OptionRuneArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualRuneArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionByteArr(a, b OptionByteArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualByteArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionIntArr(a, b OptionIntArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualIntArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt8Arr(a, b OptionInt8Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt8Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt16Arr(a, b OptionInt16Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt16Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt32Arr(a, b OptionInt32Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt32Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionInt64Arr(a, b OptionInt64Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualInt64Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUintArr(a, b OptionUintArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUintArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint8Arr(a, b OptionUint8Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint8Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint16Arr(a, b OptionUint16Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint16Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint32Arr(a, b OptionUint32Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint32Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUint64Arr(a, b OptionUint64Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUint64Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionUintptrArr(a, b OptionUintptrArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualUintptrArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionFloat32Arr(a, b OptionFloat32Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualFloat32Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionFloat64Arr(a, b OptionFloat64Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualFloat64Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionComplex64Arr(a, b OptionComplex64Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualComplex64Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionComplex128Arr(a, b OptionComplex128Arr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualComplex128Arr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionStringArr(a, b OptionStringArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualStringArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualOptionAnyArr(a, b OptionAnyArr) bool {
	if a.IsDefined() {
		if b.IsDefined() {
			return EqualAnyArr(*a.value, *b.value)
		} else {
			return false
		}
	} else if b.IsDefined() {
		return false
	} else {
		return true
	}
}

func EqualBoolList(a, b BoolList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualBool(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualRuneList(a, b RuneList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualRune(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualByteList(a, b ByteList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualByte(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualIntList(a, b IntList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt8List(a, b Int8List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt8(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt16List(a, b Int16List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt16(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt32List(a, b Int32List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt32(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt64List(a, b Int64List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt64(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUintList(a, b UintList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint8List(a, b Uint8List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint8(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint16List(a, b Uint16List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint16(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint32List(a, b Uint32List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint32(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint64List(a, b Uint64List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint64(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUintptrList(a, b UintptrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUintptr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualFloat32List(a, b Float32List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualFloat32(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualFloat64List(a, b Float64List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualFloat64(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualComplex64List(a, b Complex64List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualComplex64(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualComplex128List(a, b Complex128List) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualComplex128(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualStringList(a, b StringList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualString(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualAnyList(a, b AnyList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualAny(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualBoolArrList(a, b BoolArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualBoolArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualRuneArrList(a, b RuneArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualRuneArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualByteArrList(a, b ByteArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualByteArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualIntArrList(a, b IntArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualIntArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt8ArrList(a, b Int8ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt8Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt16ArrList(a, b Int16ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt16Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt32ArrList(a, b Int32ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt32Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualInt64ArrList(a, b Int64ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualInt64Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUintArrList(a, b UintArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUintArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint8ArrList(a, b Uint8ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint8Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint16ArrList(a, b Uint16ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint16Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint32ArrList(a, b Uint32ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint32Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUint64ArrList(a, b Uint64ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUint64Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualUintptrArrList(a, b UintptrArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualUintptrArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualFloat32ArrList(a, b Float32ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualFloat32Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualFloat64ArrList(a, b Float64ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualFloat64Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualComplex64ArrList(a, b Complex64ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualComplex64Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualComplex128ArrList(a, b Complex128ArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualComplex128Arr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualStringArrList(a, b StringArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualStringArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}

func EqualAnyArrList(a, b AnyArrList) bool {
	if a.Size() != b.Size() {
		return false
	}

	xs1 := a
	xs2 := b
	for xs1.IsNotEmpty() {
		if !EqualAnyArr(*xs1.head, *xs2.head) {
			return false
		}
		xs1 = *xs1.tail
		xs2 = *xs2.tail
	}
	return true
}
