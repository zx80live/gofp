// flist_make.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

func MakeBoolList(elements ...bool) BoolList {
	l := NilBool
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeRuneList(elements ...rune) RuneList {
	l := NilRune
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeByteList(elements ...byte) ByteList {
	l := NilByte
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeIntList(elements ...int) IntList {
	l := NilInt
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt8List(elements ...int8) Int8List {
	l := NilInt8
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt16List(elements ...int16) Int16List {
	l := NilInt16
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt32List(elements ...int32) Int32List {
	l := NilInt32
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt64List(elements ...int64) Int64List {
	l := NilInt64
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUintList(elements ...uint) UintList {
	l := NilUint
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint8List(elements ...uint8) Uint8List {
	l := NilUint8
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint16List(elements ...uint16) Uint16List {
	l := NilUint16
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint32List(elements ...uint32) Uint32List {
	l := NilUint32
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint64List(elements ...uint64) Uint64List {
	l := NilUint64
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUintptrList(elements ...uintptr) UintptrList {
	l := NilUintptr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeFloat32List(elements ...float32) Float32List {
	l := NilFloat32
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeFloat64List(elements ...float64) Float64List {
	l := NilFloat64
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeComplex64List(elements ...complex64) Complex64List {
	l := NilComplex64
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeComplex128List(elements ...complex128) Complex128List {
	l := NilComplex128
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeStringList(elements ...string) StringList {
	l := NilString
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeAnyList(elements ...Any) AnyList {
	l := NilAny
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeBoolArrList(elements ...[]bool) BoolArrList {
	l := NilBoolArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeRuneArrList(elements ...[]rune) RuneArrList {
	l := NilRuneArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeByteArrList(elements ...[]byte) ByteArrList {
	l := NilByteArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeIntArrList(elements ...[]int) IntArrList {
	l := NilIntArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt8ArrList(elements ...[]int8) Int8ArrList {
	l := NilInt8Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt16ArrList(elements ...[]int16) Int16ArrList {
	l := NilInt16Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt32ArrList(elements ...[]int32) Int32ArrList {
	l := NilInt32Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeInt64ArrList(elements ...[]int64) Int64ArrList {
	l := NilInt64Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUintArrList(elements ...[]uint) UintArrList {
	l := NilUintArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint8ArrList(elements ...[]uint8) Uint8ArrList {
	l := NilUint8Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint16ArrList(elements ...[]uint16) Uint16ArrList {
	l := NilUint16Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint32ArrList(elements ...[]uint32) Uint32ArrList {
	l := NilUint32Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUint64ArrList(elements ...[]uint64) Uint64ArrList {
	l := NilUint64Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeUintptrArrList(elements ...[]uintptr) UintptrArrList {
	l := NilUintptrArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeFloat32ArrList(elements ...[]float32) Float32ArrList {
	l := NilFloat32Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeFloat64ArrList(elements ...[]float64) Float64ArrList {
	l := NilFloat64Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeComplex64ArrList(elements ...[]complex64) Complex64ArrList {
	l := NilComplex64Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeComplex128ArrList(elements ...[]complex128) Complex128ArrList {
	l := NilComplex128Arr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeStringArrList(elements ...[]string) StringArrList {
	l := NilStringArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
func MakeAnyArrList(elements ...[]Any) AnyArrList {
	l := NilAnyArr
	for i := len(elements) - 1; i >= 0; i-- {
		l = l.Cons(elements[i])
	}
	return l
}
