// flist_foreach.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

func (l ListBool) Foreach(f func(bool)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListRune) Foreach(f func(rune)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListByte) Foreach(f func(byte)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt) Foreach(f func(int)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt8) Foreach(f func(int8)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt16) Foreach(f func(int16)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt32) Foreach(f func(int32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt64) Foreach(f func(int64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint) Foreach(f func(uint)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint8) Foreach(f func(uint8)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint16) Foreach(f func(uint16)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint32) Foreach(f func(uint32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint64) Foreach(f func(uint64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUintptr) Foreach(f func(uintptr)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListFloat32) Foreach(f func(float32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListFloat64) Foreach(f func(float64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListComplex64) Foreach(f func(complex64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListComplex128) Foreach(f func(complex128)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListString) Foreach(f func(string)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInterface) Foreach(f func(interface{})) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListBoolArr) Foreach(f func([]bool)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListRuneArr) Foreach(f func([]rune)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListByteArr) Foreach(f func([]byte)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListIntArr) Foreach(f func([]int)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt8Arr) Foreach(f func([]int8)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt16Arr) Foreach(f func([]int16)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt32Arr) Foreach(f func([]int32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInt64Arr) Foreach(f func([]int64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUintArr) Foreach(f func([]uint)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint8Arr) Foreach(f func([]uint8)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint16Arr) Foreach(f func([]uint16)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint32Arr) Foreach(f func([]uint32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUint64Arr) Foreach(f func([]uint64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListUintptrArr) Foreach(f func([]uintptr)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListFloat32Arr) Foreach(f func([]float32)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListFloat64Arr) Foreach(f func([]float64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListComplex64Arr) Foreach(f func([]complex64)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListComplex128Arr) Foreach(f func([]complex128)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListStringArr) Foreach(f func([]string)) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
func (l ListInterfaceArr) Foreach(f func([]interface{})) {
	if l.IsNotEmpty() {
		f(*l.head)
		l.tail.Foreach(f)
	}
}
