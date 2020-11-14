// fpredicate_logic.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

func (p1 BoolPredicate) And(p2 BoolPredicate) BoolPredicate {
	return func(e bool) bool { return p1(e) && p2(e) }
}

func (p1 RunePredicate) And(p2 RunePredicate) RunePredicate {
	return func(e rune) bool { return p1(e) && p2(e) }
}

func (p1 BytePredicate) And(p2 BytePredicate) BytePredicate {
	return func(e byte) bool { return p1(e) && p2(e) }
}

func (p1 IntPredicate) And(p2 IntPredicate) IntPredicate {
	return func(e int) bool { return p1(e) && p2(e) }
}

func (p1 Int8Predicate) And(p2 Int8Predicate) Int8Predicate {
	return func(e int8) bool { return p1(e) && p2(e) }
}

func (p1 Int16Predicate) And(p2 Int16Predicate) Int16Predicate {
	return func(e int16) bool { return p1(e) && p2(e) }
}

func (p1 Int32Predicate) And(p2 Int32Predicate) Int32Predicate {
	return func(e int32) bool { return p1(e) && p2(e) }
}

func (p1 Int64Predicate) And(p2 Int64Predicate) Int64Predicate {
	return func(e int64) bool { return p1(e) && p2(e) }
}

func (p1 UintPredicate) And(p2 UintPredicate) UintPredicate {
	return func(e uint) bool { return p1(e) && p2(e) }
}

func (p1 Uint8Predicate) And(p2 Uint8Predicate) Uint8Predicate {
	return func(e uint8) bool { return p1(e) && p2(e) }
}

func (p1 Uint16Predicate) And(p2 Uint16Predicate) Uint16Predicate {
	return func(e uint16) bool { return p1(e) && p2(e) }
}

func (p1 Uint32Predicate) And(p2 Uint32Predicate) Uint32Predicate {
	return func(e uint32) bool { return p1(e) && p2(e) }
}

func (p1 Uint64Predicate) And(p2 Uint64Predicate) Uint64Predicate {
	return func(e uint64) bool { return p1(e) && p2(e) }
}

func (p1 UintptrPredicate) And(p2 UintptrPredicate) UintptrPredicate {
	return func(e uintptr) bool { return p1(e) && p2(e) }
}

func (p1 Float32Predicate) And(p2 Float32Predicate) Float32Predicate {
	return func(e float32) bool { return p1(e) && p2(e) }
}

func (p1 Float64Predicate) And(p2 Float64Predicate) Float64Predicate {
	return func(e float64) bool { return p1(e) && p2(e) }
}

func (p1 Complex64Predicate) And(p2 Complex64Predicate) Complex64Predicate {
	return func(e complex64) bool { return p1(e) && p2(e) }
}

func (p1 Complex128Predicate) And(p2 Complex128Predicate) Complex128Predicate {
	return func(e complex128) bool { return p1(e) && p2(e) }
}

func (p1 StringPredicate) And(p2 StringPredicate) StringPredicate {
	return func(e string) bool { return p1(e) && p2(e) }
}

func (p1 AnyPredicate) And(p2 AnyPredicate) AnyPredicate {
	return func(e Any) bool { return p1(e) && p2(e) }
}

func (p1 BoolArrPredicate) And(p2 BoolArrPredicate) BoolArrPredicate {
	return func(e []bool) bool { return p1(e) && p2(e) }
}

func (p1 RuneArrPredicate) And(p2 RuneArrPredicate) RuneArrPredicate {
	return func(e []rune) bool { return p1(e) && p2(e) }
}

func (p1 ByteArrPredicate) And(p2 ByteArrPredicate) ByteArrPredicate {
	return func(e []byte) bool { return p1(e) && p2(e) }
}

func (p1 IntArrPredicate) And(p2 IntArrPredicate) IntArrPredicate {
	return func(e []int) bool { return p1(e) && p2(e) }
}

func (p1 Int8ArrPredicate) And(p2 Int8ArrPredicate) Int8ArrPredicate {
	return func(e []int8) bool { return p1(e) && p2(e) }
}

func (p1 Int16ArrPredicate) And(p2 Int16ArrPredicate) Int16ArrPredicate {
	return func(e []int16) bool { return p1(e) && p2(e) }
}

func (p1 Int32ArrPredicate) And(p2 Int32ArrPredicate) Int32ArrPredicate {
	return func(e []int32) bool { return p1(e) && p2(e) }
}

func (p1 Int64ArrPredicate) And(p2 Int64ArrPredicate) Int64ArrPredicate {
	return func(e []int64) bool { return p1(e) && p2(e) }
}

func (p1 UintArrPredicate) And(p2 UintArrPredicate) UintArrPredicate {
	return func(e []uint) bool { return p1(e) && p2(e) }
}

func (p1 Uint8ArrPredicate) And(p2 Uint8ArrPredicate) Uint8ArrPredicate {
	return func(e []uint8) bool { return p1(e) && p2(e) }
}

func (p1 Uint16ArrPredicate) And(p2 Uint16ArrPredicate) Uint16ArrPredicate {
	return func(e []uint16) bool { return p1(e) && p2(e) }
}

func (p1 Uint32ArrPredicate) And(p2 Uint32ArrPredicate) Uint32ArrPredicate {
	return func(e []uint32) bool { return p1(e) && p2(e) }
}

func (p1 Uint64ArrPredicate) And(p2 Uint64ArrPredicate) Uint64ArrPredicate {
	return func(e []uint64) bool { return p1(e) && p2(e) }
}

func (p1 UintptrArrPredicate) And(p2 UintptrArrPredicate) UintptrArrPredicate {
	return func(e []uintptr) bool { return p1(e) && p2(e) }
}

func (p1 Float32ArrPredicate) And(p2 Float32ArrPredicate) Float32ArrPredicate {
	return func(e []float32) bool { return p1(e) && p2(e) }
}

func (p1 Float64ArrPredicate) And(p2 Float64ArrPredicate) Float64ArrPredicate {
	return func(e []float64) bool { return p1(e) && p2(e) }
}

func (p1 Complex64ArrPredicate) And(p2 Complex64ArrPredicate) Complex64ArrPredicate {
	return func(e []complex64) bool { return p1(e) && p2(e) }
}

func (p1 Complex128ArrPredicate) And(p2 Complex128ArrPredicate) Complex128ArrPredicate {
	return func(e []complex128) bool { return p1(e) && p2(e) }
}

func (p1 StringArrPredicate) And(p2 StringArrPredicate) StringArrPredicate {
	return func(e []string) bool { return p1(e) && p2(e) }
}

func (p1 AnyArrPredicate) And(p2 AnyArrPredicate) AnyArrPredicate {
	return func(e []Any) bool { return p1(e) && p2(e) }
}

func (p1 BoolPredicate) Or(p2 BoolPredicate) BoolPredicate {
	return func(e bool) bool { return p1(e) || p2(e) }
}

func (p1 RunePredicate) Or(p2 RunePredicate) RunePredicate {
	return func(e rune) bool { return p1(e) || p2(e) }
}

func (p1 BytePredicate) Or(p2 BytePredicate) BytePredicate {
	return func(e byte) bool { return p1(e) || p2(e) }
}

func (p1 IntPredicate) Or(p2 IntPredicate) IntPredicate {
	return func(e int) bool { return p1(e) || p2(e) }
}

func (p1 Int8Predicate) Or(p2 Int8Predicate) Int8Predicate {
	return func(e int8) bool { return p1(e) || p2(e) }
}

func (p1 Int16Predicate) Or(p2 Int16Predicate) Int16Predicate {
	return func(e int16) bool { return p1(e) || p2(e) }
}

func (p1 Int32Predicate) Or(p2 Int32Predicate) Int32Predicate {
	return func(e int32) bool { return p1(e) || p2(e) }
}

func (p1 Int64Predicate) Or(p2 Int64Predicate) Int64Predicate {
	return func(e int64) bool { return p1(e) || p2(e) }
}

func (p1 UintPredicate) Or(p2 UintPredicate) UintPredicate {
	return func(e uint) bool { return p1(e) || p2(e) }
}

func (p1 Uint8Predicate) Or(p2 Uint8Predicate) Uint8Predicate {
	return func(e uint8) bool { return p1(e) || p2(e) }
}

func (p1 Uint16Predicate) Or(p2 Uint16Predicate) Uint16Predicate {
	return func(e uint16) bool { return p1(e) || p2(e) }
}

func (p1 Uint32Predicate) Or(p2 Uint32Predicate) Uint32Predicate {
	return func(e uint32) bool { return p1(e) || p2(e) }
}

func (p1 Uint64Predicate) Or(p2 Uint64Predicate) Uint64Predicate {
	return func(e uint64) bool { return p1(e) || p2(e) }
}

func (p1 UintptrPredicate) Or(p2 UintptrPredicate) UintptrPredicate {
	return func(e uintptr) bool { return p1(e) || p2(e) }
}

func (p1 Float32Predicate) Or(p2 Float32Predicate) Float32Predicate {
	return func(e float32) bool { return p1(e) || p2(e) }
}

func (p1 Float64Predicate) Or(p2 Float64Predicate) Float64Predicate {
	return func(e float64) bool { return p1(e) || p2(e) }
}

func (p1 Complex64Predicate) Or(p2 Complex64Predicate) Complex64Predicate {
	return func(e complex64) bool { return p1(e) || p2(e) }
}

func (p1 Complex128Predicate) Or(p2 Complex128Predicate) Complex128Predicate {
	return func(e complex128) bool { return p1(e) || p2(e) }
}

func (p1 StringPredicate) Or(p2 StringPredicate) StringPredicate {
	return func(e string) bool { return p1(e) || p2(e) }
}

func (p1 AnyPredicate) Or(p2 AnyPredicate) AnyPredicate {
	return func(e Any) bool { return p1(e) || p2(e) }
}

func (p1 BoolArrPredicate) Or(p2 BoolArrPredicate) BoolArrPredicate {
	return func(e []bool) bool { return p1(e) || p2(e) }
}

func (p1 RuneArrPredicate) Or(p2 RuneArrPredicate) RuneArrPredicate {
	return func(e []rune) bool { return p1(e) || p2(e) }
}

func (p1 ByteArrPredicate) Or(p2 ByteArrPredicate) ByteArrPredicate {
	return func(e []byte) bool { return p1(e) || p2(e) }
}

func (p1 IntArrPredicate) Or(p2 IntArrPredicate) IntArrPredicate {
	return func(e []int) bool { return p1(e) || p2(e) }
}

func (p1 Int8ArrPredicate) Or(p2 Int8ArrPredicate) Int8ArrPredicate {
	return func(e []int8) bool { return p1(e) || p2(e) }
}

func (p1 Int16ArrPredicate) Or(p2 Int16ArrPredicate) Int16ArrPredicate {
	return func(e []int16) bool { return p1(e) || p2(e) }
}

func (p1 Int32ArrPredicate) Or(p2 Int32ArrPredicate) Int32ArrPredicate {
	return func(e []int32) bool { return p1(e) || p2(e) }
}

func (p1 Int64ArrPredicate) Or(p2 Int64ArrPredicate) Int64ArrPredicate {
	return func(e []int64) bool { return p1(e) || p2(e) }
}

func (p1 UintArrPredicate) Or(p2 UintArrPredicate) UintArrPredicate {
	return func(e []uint) bool { return p1(e) || p2(e) }
}

func (p1 Uint8ArrPredicate) Or(p2 Uint8ArrPredicate) Uint8ArrPredicate {
	return func(e []uint8) bool { return p1(e) || p2(e) }
}

func (p1 Uint16ArrPredicate) Or(p2 Uint16ArrPredicate) Uint16ArrPredicate {
	return func(e []uint16) bool { return p1(e) || p2(e) }
}

func (p1 Uint32ArrPredicate) Or(p2 Uint32ArrPredicate) Uint32ArrPredicate {
	return func(e []uint32) bool { return p1(e) || p2(e) }
}

func (p1 Uint64ArrPredicate) Or(p2 Uint64ArrPredicate) Uint64ArrPredicate {
	return func(e []uint64) bool { return p1(e) || p2(e) }
}

func (p1 UintptrArrPredicate) Or(p2 UintptrArrPredicate) UintptrArrPredicate {
	return func(e []uintptr) bool { return p1(e) || p2(e) }
}

func (p1 Float32ArrPredicate) Or(p2 Float32ArrPredicate) Float32ArrPredicate {
	return func(e []float32) bool { return p1(e) || p2(e) }
}

func (p1 Float64ArrPredicate) Or(p2 Float64ArrPredicate) Float64ArrPredicate {
	return func(e []float64) bool { return p1(e) || p2(e) }
}

func (p1 Complex64ArrPredicate) Or(p2 Complex64ArrPredicate) Complex64ArrPredicate {
	return func(e []complex64) bool { return p1(e) || p2(e) }
}

func (p1 Complex128ArrPredicate) Or(p2 Complex128ArrPredicate) Complex128ArrPredicate {
	return func(e []complex128) bool { return p1(e) || p2(e) }
}

func (p1 StringArrPredicate) Or(p2 StringArrPredicate) StringArrPredicate {
	return func(e []string) bool { return p1(e) || p2(e) }
}

func (p1 AnyArrPredicate) Or(p2 AnyArrPredicate) AnyArrPredicate {
	return func(e []Any) bool { return p1(e) || p2(e) }
}

func (p1 BoolPredicate) Xor(p2 BoolPredicate) BoolPredicate {
	return func(e bool) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 RunePredicate) Xor(p2 RunePredicate) RunePredicate {
	return func(e rune) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 BytePredicate) Xor(p2 BytePredicate) BytePredicate {
	return func(e byte) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 IntPredicate) Xor(p2 IntPredicate) IntPredicate {
	return func(e int) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int8Predicate) Xor(p2 Int8Predicate) Int8Predicate {
	return func(e int8) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int16Predicate) Xor(p2 Int16Predicate) Int16Predicate {
	return func(e int16) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int32Predicate) Xor(p2 Int32Predicate) Int32Predicate {
	return func(e int32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int64Predicate) Xor(p2 Int64Predicate) Int64Predicate {
	return func(e int64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 UintPredicate) Xor(p2 UintPredicate) UintPredicate {
	return func(e uint) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint8Predicate) Xor(p2 Uint8Predicate) Uint8Predicate {
	return func(e uint8) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint16Predicate) Xor(p2 Uint16Predicate) Uint16Predicate {
	return func(e uint16) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint32Predicate) Xor(p2 Uint32Predicate) Uint32Predicate {
	return func(e uint32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint64Predicate) Xor(p2 Uint64Predicate) Uint64Predicate {
	return func(e uint64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 UintptrPredicate) Xor(p2 UintptrPredicate) UintptrPredicate {
	return func(e uintptr) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Float32Predicate) Xor(p2 Float32Predicate) Float32Predicate {
	return func(e float32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Float64Predicate) Xor(p2 Float64Predicate) Float64Predicate {
	return func(e float64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Complex64Predicate) Xor(p2 Complex64Predicate) Complex64Predicate {
	return func(e complex64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Complex128Predicate) Xor(p2 Complex128Predicate) Complex128Predicate {
	return func(e complex128) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 StringPredicate) Xor(p2 StringPredicate) StringPredicate {
	return func(e string) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 AnyPredicate) Xor(p2 AnyPredicate) AnyPredicate {
	return func(e Any) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 BoolArrPredicate) Xor(p2 BoolArrPredicate) BoolArrPredicate {
	return func(e []bool) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 RuneArrPredicate) Xor(p2 RuneArrPredicate) RuneArrPredicate {
	return func(e []rune) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 ByteArrPredicate) Xor(p2 ByteArrPredicate) ByteArrPredicate {
	return func(e []byte) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 IntArrPredicate) Xor(p2 IntArrPredicate) IntArrPredicate {
	return func(e []int) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int8ArrPredicate) Xor(p2 Int8ArrPredicate) Int8ArrPredicate {
	return func(e []int8) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int16ArrPredicate) Xor(p2 Int16ArrPredicate) Int16ArrPredicate {
	return func(e []int16) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int32ArrPredicate) Xor(p2 Int32ArrPredicate) Int32ArrPredicate {
	return func(e []int32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Int64ArrPredicate) Xor(p2 Int64ArrPredicate) Int64ArrPredicate {
	return func(e []int64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 UintArrPredicate) Xor(p2 UintArrPredicate) UintArrPredicate {
	return func(e []uint) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint8ArrPredicate) Xor(p2 Uint8ArrPredicate) Uint8ArrPredicate {
	return func(e []uint8) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint16ArrPredicate) Xor(p2 Uint16ArrPredicate) Uint16ArrPredicate {
	return func(e []uint16) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint32ArrPredicate) Xor(p2 Uint32ArrPredicate) Uint32ArrPredicate {
	return func(e []uint32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Uint64ArrPredicate) Xor(p2 Uint64ArrPredicate) Uint64ArrPredicate {
	return func(e []uint64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 UintptrArrPredicate) Xor(p2 UintptrArrPredicate) UintptrArrPredicate {
	return func(e []uintptr) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Float32ArrPredicate) Xor(p2 Float32ArrPredicate) Float32ArrPredicate {
	return func(e []float32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Float64ArrPredicate) Xor(p2 Float64ArrPredicate) Float64ArrPredicate {
	return func(e []float64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Complex64ArrPredicate) Xor(p2 Complex64ArrPredicate) Complex64ArrPredicate {
	return func(e []complex64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 Complex128ArrPredicate) Xor(p2 Complex128ArrPredicate) Complex128ArrPredicate {
	return func(e []complex128) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 StringArrPredicate) Xor(p2 StringArrPredicate) StringArrPredicate {
	return func(e []string) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p1 AnyArrPredicate) Xor(p2 AnyArrPredicate) AnyArrPredicate {
	return func(e []Any) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) }
}

func (p BoolPredicate) Neg() BoolPredicate { return func(e bool) bool { return !p(e) } }

func (p RunePredicate) Neg() RunePredicate { return func(e rune) bool { return !p(e) } }

func (p BytePredicate) Neg() BytePredicate { return func(e byte) bool { return !p(e) } }

func (p IntPredicate) Neg() IntPredicate { return func(e int) bool { return !p(e) } }

func (p Int8Predicate) Neg() Int8Predicate { return func(e int8) bool { return !p(e) } }

func (p Int16Predicate) Neg() Int16Predicate { return func(e int16) bool { return !p(e) } }

func (p Int32Predicate) Neg() Int32Predicate { return func(e int32) bool { return !p(e) } }

func (p Int64Predicate) Neg() Int64Predicate { return func(e int64) bool { return !p(e) } }

func (p UintPredicate) Neg() UintPredicate { return func(e uint) bool { return !p(e) } }

func (p Uint8Predicate) Neg() Uint8Predicate { return func(e uint8) bool { return !p(e) } }

func (p Uint16Predicate) Neg() Uint16Predicate { return func(e uint16) bool { return !p(e) } }

func (p Uint32Predicate) Neg() Uint32Predicate { return func(e uint32) bool { return !p(e) } }

func (p Uint64Predicate) Neg() Uint64Predicate { return func(e uint64) bool { return !p(e) } }

func (p UintptrPredicate) Neg() UintptrPredicate { return func(e uintptr) bool { return !p(e) } }

func (p Float32Predicate) Neg() Float32Predicate { return func(e float32) bool { return !p(e) } }

func (p Float64Predicate) Neg() Float64Predicate { return func(e float64) bool { return !p(e) } }

func (p Complex64Predicate) Neg() Complex64Predicate { return func(e complex64) bool { return !p(e) } }

func (p Complex128Predicate) Neg() Complex128Predicate {
	return func(e complex128) bool { return !p(e) }
}

func (p StringPredicate) Neg() StringPredicate { return func(e string) bool { return !p(e) } }

func (p AnyPredicate) Neg() AnyPredicate { return func(e Any) bool { return !p(e) } }

func (p BoolArrPredicate) Neg() BoolArrPredicate { return func(e []bool) bool { return !p(e) } }

func (p RuneArrPredicate) Neg() RuneArrPredicate { return func(e []rune) bool { return !p(e) } }

func (p ByteArrPredicate) Neg() ByteArrPredicate { return func(e []byte) bool { return !p(e) } }

func (p IntArrPredicate) Neg() IntArrPredicate { return func(e []int) bool { return !p(e) } }

func (p Int8ArrPredicate) Neg() Int8ArrPredicate { return func(e []int8) bool { return !p(e) } }

func (p Int16ArrPredicate) Neg() Int16ArrPredicate { return func(e []int16) bool { return !p(e) } }

func (p Int32ArrPredicate) Neg() Int32ArrPredicate { return func(e []int32) bool { return !p(e) } }

func (p Int64ArrPredicate) Neg() Int64ArrPredicate { return func(e []int64) bool { return !p(e) } }

func (p UintArrPredicate) Neg() UintArrPredicate { return func(e []uint) bool { return !p(e) } }

func (p Uint8ArrPredicate) Neg() Uint8ArrPredicate { return func(e []uint8) bool { return !p(e) } }

func (p Uint16ArrPredicate) Neg() Uint16ArrPredicate { return func(e []uint16) bool { return !p(e) } }

func (p Uint32ArrPredicate) Neg() Uint32ArrPredicate { return func(e []uint32) bool { return !p(e) } }

func (p Uint64ArrPredicate) Neg() Uint64ArrPredicate { return func(e []uint64) bool { return !p(e) } }

func (p UintptrArrPredicate) Neg() UintptrArrPredicate {
	return func(e []uintptr) bool { return !p(e) }
}

func (p Float32ArrPredicate) Neg() Float32ArrPredicate {
	return func(e []float32) bool { return !p(e) }
}

func (p Float64ArrPredicate) Neg() Float64ArrPredicate {
	return func(e []float64) bool { return !p(e) }
}

func (p Complex64ArrPredicate) Neg() Complex64ArrPredicate {
	return func(e []complex64) bool { return !p(e) }
}

func (p Complex128ArrPredicate) Neg() Complex128ArrPredicate {
	return func(e []complex128) bool { return !p(e) }
}

func (p StringArrPredicate) Neg() StringArrPredicate { return func(e []string) bool { return !p(e) } }

func (p AnyArrPredicate) Neg() AnyArrPredicate { return func(e []Any) bool { return !p(e) } }