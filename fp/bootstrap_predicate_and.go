// bootstrap_predicate_and.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (p1 BoolPredicate) And(p2 BoolPredicate) BoolPredicate {
	return func(e bool) bool { return p1(e) && p2(e) }
}
func (p1 StringPredicate) And(p2 StringPredicate) StringPredicate {
	return func(e string) bool { return p1(e) && p2(e) }
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
func (p1 BytePredicate) And(p2 BytePredicate) BytePredicate {
	return func(e byte) bool { return p1(e) && p2(e) }
}
func (p1 RunePredicate) And(p2 RunePredicate) RunePredicate {
	return func(e rune) bool { return p1(e) && p2(e) }
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
func (p1 AnyPredicate) And(p2 AnyPredicate) AnyPredicate {
	return func(e Any) bool { return p1(e) && p2(e) }
}
func (p1 BoolArrayPredicate) And(p2 BoolArrayPredicate) BoolArrayPredicate {
	return func(e []bool) bool { return p1(e) && p2(e) }
}
func (p1 StringArrayPredicate) And(p2 StringArrayPredicate) StringArrayPredicate {
	return func(e []string) bool { return p1(e) && p2(e) }
}
func (p1 IntArrayPredicate) And(p2 IntArrayPredicate) IntArrayPredicate {
	return func(e []int) bool { return p1(e) && p2(e) }
}
func (p1 Int64ArrayPredicate) And(p2 Int64ArrayPredicate) Int64ArrayPredicate {
	return func(e []int64) bool { return p1(e) && p2(e) }
}
func (p1 ByteArrayPredicate) And(p2 ByteArrayPredicate) ByteArrayPredicate {
	return func(e []byte) bool { return p1(e) && p2(e) }
}
func (p1 RuneArrayPredicate) And(p2 RuneArrayPredicate) RuneArrayPredicate {
	return func(e []rune) bool { return p1(e) && p2(e) }
}
func (p1 Float32ArrayPredicate) And(p2 Float32ArrayPredicate) Float32ArrayPredicate {
	return func(e []float32) bool { return p1(e) && p2(e) }
}
func (p1 Float64ArrayPredicate) And(p2 Float64ArrayPredicate) Float64ArrayPredicate {
	return func(e []float64) bool { return p1(e) && p2(e) }
}
func (p1 AnyArrayPredicate) And(p2 AnyArrayPredicate) AnyArrayPredicate {
	return func(e []Any) bool { return p1(e) && p2(e) }
}
func (p1 BoolArrayArrayPredicate) And(p2 BoolArrayArrayPredicate) BoolArrayArrayPredicate {
	return func(e [][]bool) bool { return p1(e) && p2(e) }
}
func (p1 StringArrayArrayPredicate) And(p2 StringArrayArrayPredicate) StringArrayArrayPredicate {
	return func(e [][]string) bool { return p1(e) && p2(e) }
}
func (p1 IntArrayArrayPredicate) And(p2 IntArrayArrayPredicate) IntArrayArrayPredicate {
	return func(e [][]int) bool { return p1(e) && p2(e) }
}
func (p1 Int64ArrayArrayPredicate) And(p2 Int64ArrayArrayPredicate) Int64ArrayArrayPredicate {
	return func(e [][]int64) bool { return p1(e) && p2(e) }
}
func (p1 ByteArrayArrayPredicate) And(p2 ByteArrayArrayPredicate) ByteArrayArrayPredicate {
	return func(e [][]byte) bool { return p1(e) && p2(e) }
}
func (p1 RuneArrayArrayPredicate) And(p2 RuneArrayArrayPredicate) RuneArrayArrayPredicate {
	return func(e [][]rune) bool { return p1(e) && p2(e) }
}
func (p1 Float32ArrayArrayPredicate) And(p2 Float32ArrayArrayPredicate) Float32ArrayArrayPredicate {
	return func(e [][]float32) bool { return p1(e) && p2(e) }
}
func (p1 Float64ArrayArrayPredicate) And(p2 Float64ArrayArrayPredicate) Float64ArrayArrayPredicate {
	return func(e [][]float64) bool { return p1(e) && p2(e) }
}
func (p1 AnyArrayArrayPredicate) And(p2 AnyArrayArrayPredicate) AnyArrayArrayPredicate {
	return func(e [][]Any) bool { return p1(e) && p2(e) }
}
func (p1 BoolOptionPredicate) And(p2 BoolOptionPredicate) BoolOptionPredicate {
	return func(e BoolOption) bool { return p1(e) && p2(e) }
}
func (p1 StringOptionPredicate) And(p2 StringOptionPredicate) StringOptionPredicate {
	return func(e StringOption) bool { return p1(e) && p2(e) }
}
func (p1 IntOptionPredicate) And(p2 IntOptionPredicate) IntOptionPredicate {
	return func(e IntOption) bool { return p1(e) && p2(e) }
}
func (p1 Int64OptionPredicate) And(p2 Int64OptionPredicate) Int64OptionPredicate {
	return func(e Int64Option) bool { return p1(e) && p2(e) }
}
func (p1 ByteOptionPredicate) And(p2 ByteOptionPredicate) ByteOptionPredicate {
	return func(e ByteOption) bool { return p1(e) && p2(e) }
}
func (p1 RuneOptionPredicate) And(p2 RuneOptionPredicate) RuneOptionPredicate {
	return func(e RuneOption) bool { return p1(e) && p2(e) }
}
func (p1 Float32OptionPredicate) And(p2 Float32OptionPredicate) Float32OptionPredicate {
	return func(e Float32Option) bool { return p1(e) && p2(e) }
}
func (p1 Float64OptionPredicate) And(p2 Float64OptionPredicate) Float64OptionPredicate {
	return func(e Float64Option) bool { return p1(e) && p2(e) }
}
func (p1 AnyOptionPredicate) And(p2 AnyOptionPredicate) AnyOptionPredicate {
	return func(e AnyOption) bool { return p1(e) && p2(e) }
}
func (p1 BoolOptionOptionPredicate) And(p2 BoolOptionOptionPredicate) BoolOptionOptionPredicate {
	return func(e BoolOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 StringOptionOptionPredicate) And(p2 StringOptionOptionPredicate) StringOptionOptionPredicate {
	return func(e StringOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 IntOptionOptionPredicate) And(p2 IntOptionOptionPredicate) IntOptionOptionPredicate {
	return func(e IntOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 Int64OptionOptionPredicate) And(p2 Int64OptionOptionPredicate) Int64OptionOptionPredicate {
	return func(e Int64OptionOption) bool { return p1(e) && p2(e) }
}
func (p1 ByteOptionOptionPredicate) And(p2 ByteOptionOptionPredicate) ByteOptionOptionPredicate {
	return func(e ByteOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 RuneOptionOptionPredicate) And(p2 RuneOptionOptionPredicate) RuneOptionOptionPredicate {
	return func(e RuneOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 Float32OptionOptionPredicate) And(p2 Float32OptionOptionPredicate) Float32OptionOptionPredicate {
	return func(e Float32OptionOption) bool { return p1(e) && p2(e) }
}
func (p1 Float64OptionOptionPredicate) And(p2 Float64OptionOptionPredicate) Float64OptionOptionPredicate {
	return func(e Float64OptionOption) bool { return p1(e) && p2(e) }
}
func (p1 AnyOptionOptionPredicate) And(p2 AnyOptionOptionPredicate) AnyOptionOptionPredicate {
	return func(e AnyOptionOption) bool { return p1(e) && p2(e) }
}
func (p1 BoolArrayOptionPredicate) And(p2 BoolArrayOptionPredicate) BoolArrayOptionPredicate {
	return func(e BoolArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 StringArrayOptionPredicate) And(p2 StringArrayOptionPredicate) StringArrayOptionPredicate {
	return func(e StringArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 IntArrayOptionPredicate) And(p2 IntArrayOptionPredicate) IntArrayOptionPredicate {
	return func(e IntArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 Int64ArrayOptionPredicate) And(p2 Int64ArrayOptionPredicate) Int64ArrayOptionPredicate {
	return func(e Int64ArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 ByteArrayOptionPredicate) And(p2 ByteArrayOptionPredicate) ByteArrayOptionPredicate {
	return func(e ByteArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 RuneArrayOptionPredicate) And(p2 RuneArrayOptionPredicate) RuneArrayOptionPredicate {
	return func(e RuneArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 Float32ArrayOptionPredicate) And(p2 Float32ArrayOptionPredicate) Float32ArrayOptionPredicate {
	return func(e Float32ArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 Float64ArrayOptionPredicate) And(p2 Float64ArrayOptionPredicate) Float64ArrayOptionPredicate {
	return func(e Float64ArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 AnyArrayOptionPredicate) And(p2 AnyArrayOptionPredicate) AnyArrayOptionPredicate {
	return func(e AnyArrayOption) bool { return p1(e) && p2(e) }
}
func (p1 BoolListOptionPredicate) And(p2 BoolListOptionPredicate) BoolListOptionPredicate {
	return func(e BoolListOption) bool { return p1(e) && p2(e) }
}
func (p1 StringListOptionPredicate) And(p2 StringListOptionPredicate) StringListOptionPredicate {
	return func(e StringListOption) bool { return p1(e) && p2(e) }
}
func (p1 IntListOptionPredicate) And(p2 IntListOptionPredicate) IntListOptionPredicate {
	return func(e IntListOption) bool { return p1(e) && p2(e) }
}
func (p1 Int64ListOptionPredicate) And(p2 Int64ListOptionPredicate) Int64ListOptionPredicate {
	return func(e Int64ListOption) bool { return p1(e) && p2(e) }
}
func (p1 ByteListOptionPredicate) And(p2 ByteListOptionPredicate) ByteListOptionPredicate {
	return func(e ByteListOption) bool { return p1(e) && p2(e) }
}
func (p1 RuneListOptionPredicate) And(p2 RuneListOptionPredicate) RuneListOptionPredicate {
	return func(e RuneListOption) bool { return p1(e) && p2(e) }
}
func (p1 Float32ListOptionPredicate) And(p2 Float32ListOptionPredicate) Float32ListOptionPredicate {
	return func(e Float32ListOption) bool { return p1(e) && p2(e) }
}
func (p1 Float64ListOptionPredicate) And(p2 Float64ListOptionPredicate) Float64ListOptionPredicate {
	return func(e Float64ListOption) bool { return p1(e) && p2(e) }
}
func (p1 AnyListOptionPredicate) And(p2 AnyListOptionPredicate) AnyListOptionPredicate {
	return func(e AnyListOption) bool { return p1(e) && p2(e) }
}
func (p1 BoolListPredicate) And(p2 BoolListPredicate) BoolListPredicate {
	return func(e BoolList) bool { return p1(e) && p2(e) }
}
func (p1 StringListPredicate) And(p2 StringListPredicate) StringListPredicate {
	return func(e StringList) bool { return p1(e) && p2(e) }
}
func (p1 IntListPredicate) And(p2 IntListPredicate) IntListPredicate {
	return func(e IntList) bool { return p1(e) && p2(e) }
}
func (p1 Int64ListPredicate) And(p2 Int64ListPredicate) Int64ListPredicate {
	return func(e Int64List) bool { return p1(e) && p2(e) }
}
func (p1 ByteListPredicate) And(p2 ByteListPredicate) ByteListPredicate {
	return func(e ByteList) bool { return p1(e) && p2(e) }
}
func (p1 RuneListPredicate) And(p2 RuneListPredicate) RuneListPredicate {
	return func(e RuneList) bool { return p1(e) && p2(e) }
}
func (p1 Float32ListPredicate) And(p2 Float32ListPredicate) Float32ListPredicate {
	return func(e Float32List) bool { return p1(e) && p2(e) }
}
func (p1 Float64ListPredicate) And(p2 Float64ListPredicate) Float64ListPredicate {
	return func(e Float64List) bool { return p1(e) && p2(e) }
}
func (p1 AnyListPredicate) And(p2 AnyListPredicate) AnyListPredicate {
	return func(e AnyList) bool { return p1(e) && p2(e) }
}
func (p1 BoolArrayListPredicate) And(p2 BoolArrayListPredicate) BoolArrayListPredicate {
	return func(e BoolArrayList) bool { return p1(e) && p2(e) }
}
func (p1 StringArrayListPredicate) And(p2 StringArrayListPredicate) StringArrayListPredicate {
	return func(e StringArrayList) bool { return p1(e) && p2(e) }
}
func (p1 IntArrayListPredicate) And(p2 IntArrayListPredicate) IntArrayListPredicate {
	return func(e IntArrayList) bool { return p1(e) && p2(e) }
}
func (p1 Int64ArrayListPredicate) And(p2 Int64ArrayListPredicate) Int64ArrayListPredicate {
	return func(e Int64ArrayList) bool { return p1(e) && p2(e) }
}
func (p1 ByteArrayListPredicate) And(p2 ByteArrayListPredicate) ByteArrayListPredicate {
	return func(e ByteArrayList) bool { return p1(e) && p2(e) }
}
func (p1 RuneArrayListPredicate) And(p2 RuneArrayListPredicate) RuneArrayListPredicate {
	return func(e RuneArrayList) bool { return p1(e) && p2(e) }
}
func (p1 Float32ArrayListPredicate) And(p2 Float32ArrayListPredicate) Float32ArrayListPredicate {
	return func(e Float32ArrayList) bool { return p1(e) && p2(e) }
}
func (p1 Float64ArrayListPredicate) And(p2 Float64ArrayListPredicate) Float64ArrayListPredicate {
	return func(e Float64ArrayList) bool { return p1(e) && p2(e) }
}
func (p1 AnyArrayListPredicate) And(p2 AnyArrayListPredicate) AnyArrayListPredicate {
	return func(e AnyArrayList) bool { return p1(e) && p2(e) }
}
func (p1 BoolOptionListPredicate) And(p2 BoolOptionListPredicate) BoolOptionListPredicate {
	return func(e BoolOptionList) bool { return p1(e) && p2(e) }
}
func (p1 StringOptionListPredicate) And(p2 StringOptionListPredicate) StringOptionListPredicate {
	return func(e StringOptionList) bool { return p1(e) && p2(e) }
}
func (p1 IntOptionListPredicate) And(p2 IntOptionListPredicate) IntOptionListPredicate {
	return func(e IntOptionList) bool { return p1(e) && p2(e) }
}
func (p1 Int64OptionListPredicate) And(p2 Int64OptionListPredicate) Int64OptionListPredicate {
	return func(e Int64OptionList) bool { return p1(e) && p2(e) }
}
func (p1 ByteOptionListPredicate) And(p2 ByteOptionListPredicate) ByteOptionListPredicate {
	return func(e ByteOptionList) bool { return p1(e) && p2(e) }
}
func (p1 RuneOptionListPredicate) And(p2 RuneOptionListPredicate) RuneOptionListPredicate {
	return func(e RuneOptionList) bool { return p1(e) && p2(e) }
}
func (p1 Float32OptionListPredicate) And(p2 Float32OptionListPredicate) Float32OptionListPredicate {
	return func(e Float32OptionList) bool { return p1(e) && p2(e) }
}
func (p1 Float64OptionListPredicate) And(p2 Float64OptionListPredicate) Float64OptionListPredicate {
	return func(e Float64OptionList) bool { return p1(e) && p2(e) }
}
func (p1 AnyOptionListPredicate) And(p2 AnyOptionListPredicate) AnyOptionListPredicate {
	return func(e AnyOptionList) bool { return p1(e) && p2(e) }
}
func (p1 BoolListListPredicate) And(p2 BoolListListPredicate) BoolListListPredicate {
	return func(e BoolListList) bool { return p1(e) && p2(e) }
}
func (p1 StringListListPredicate) And(p2 StringListListPredicate) StringListListPredicate {
	return func(e StringListList) bool { return p1(e) && p2(e) }
}
func (p1 IntListListPredicate) And(p2 IntListListPredicate) IntListListPredicate {
	return func(e IntListList) bool { return p1(e) && p2(e) }
}
func (p1 Int64ListListPredicate) And(p2 Int64ListListPredicate) Int64ListListPredicate {
	return func(e Int64ListList) bool { return p1(e) && p2(e) }
}
func (p1 ByteListListPredicate) And(p2 ByteListListPredicate) ByteListListPredicate {
	return func(e ByteListList) bool { return p1(e) && p2(e) }
}
func (p1 RuneListListPredicate) And(p2 RuneListListPredicate) RuneListListPredicate {
	return func(e RuneListList) bool { return p1(e) && p2(e) }
}
func (p1 Float32ListListPredicate) And(p2 Float32ListListPredicate) Float32ListListPredicate {
	return func(e Float32ListList) bool { return p1(e) && p2(e) }
}
func (p1 Float64ListListPredicate) And(p2 Float64ListListPredicate) Float64ListListPredicate {
	return func(e Float64ListList) bool { return p1(e) && p2(e) }
}
func (p1 AnyListListPredicate) And(p2 AnyListListPredicate) AnyListListPredicate {
	return func(e AnyListList) bool { return p1(e) && p2(e) }
}
