// bootstrap_predicate_xor.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (p1 BoolPredicate) Xor(p2 BoolPredicate) BoolPredicate { return func(e bool) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringPredicate) Xor(p2 StringPredicate) StringPredicate { return func(e string) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntPredicate) Xor(p2 IntPredicate) IntPredicate { return func(e int) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64Predicate) Xor(p2 Int64Predicate) Int64Predicate { return func(e int64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BytePredicate) Xor(p2 BytePredicate) BytePredicate { return func(e byte) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RunePredicate) Xor(p2 RunePredicate) RunePredicate { return func(e rune) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32Predicate) Xor(p2 Float32Predicate) Float32Predicate { return func(e float32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64Predicate) Xor(p2 Float64Predicate) Float64Predicate { return func(e float64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyPredicate) Xor(p2 AnyPredicate) AnyPredicate { return func(e Any) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolArrayPredicate) Xor(p2 BoolArrayPredicate) BoolArrayPredicate { return func(e []bool) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringArrayPredicate) Xor(p2 StringArrayPredicate) StringArrayPredicate { return func(e []string) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntArrayPredicate) Xor(p2 IntArrayPredicate) IntArrayPredicate { return func(e []int) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ArrayPredicate) Xor(p2 Int64ArrayPredicate) Int64ArrayPredicate { return func(e []int64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteArrayPredicate) Xor(p2 ByteArrayPredicate) ByteArrayPredicate { return func(e []byte) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneArrayPredicate) Xor(p2 RuneArrayPredicate) RuneArrayPredicate { return func(e []rune) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ArrayPredicate) Xor(p2 Float32ArrayPredicate) Float32ArrayPredicate { return func(e []float32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ArrayPredicate) Xor(p2 Float64ArrayPredicate) Float64ArrayPredicate { return func(e []float64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyArrayPredicate) Xor(p2 AnyArrayPredicate) AnyArrayPredicate { return func(e []Any) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ArrayPredicate) Xor(p2 Tuple2ArrayPredicate) Tuple2ArrayPredicate { return func(e []Tuple2) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolArrayArrayPredicate) Xor(p2 BoolArrayArrayPredicate) BoolArrayArrayPredicate { return func(e [][]bool) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringArrayArrayPredicate) Xor(p2 StringArrayArrayPredicate) StringArrayArrayPredicate { return func(e [][]string) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntArrayArrayPredicate) Xor(p2 IntArrayArrayPredicate) IntArrayArrayPredicate { return func(e [][]int) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ArrayArrayPredicate) Xor(p2 Int64ArrayArrayPredicate) Int64ArrayArrayPredicate { return func(e [][]int64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteArrayArrayPredicate) Xor(p2 ByteArrayArrayPredicate) ByteArrayArrayPredicate { return func(e [][]byte) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneArrayArrayPredicate) Xor(p2 RuneArrayArrayPredicate) RuneArrayArrayPredicate { return func(e [][]rune) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ArrayArrayPredicate) Xor(p2 Float32ArrayArrayPredicate) Float32ArrayArrayPredicate { return func(e [][]float32) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ArrayArrayPredicate) Xor(p2 Float64ArrayArrayPredicate) Float64ArrayArrayPredicate { return func(e [][]float64) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyArrayArrayPredicate) Xor(p2 AnyArrayArrayPredicate) AnyArrayArrayPredicate { return func(e [][]Any) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ArrayArrayPredicate) Xor(p2 Tuple2ArrayArrayPredicate) Tuple2ArrayArrayPredicate { return func(e [][]Tuple2) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolOptionPredicate) Xor(p2 BoolOptionPredicate) BoolOptionPredicate { return func(e BoolOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringOptionPredicate) Xor(p2 StringOptionPredicate) StringOptionPredicate { return func(e StringOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntOptionPredicate) Xor(p2 IntOptionPredicate) IntOptionPredicate { return func(e IntOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64OptionPredicate) Xor(p2 Int64OptionPredicate) Int64OptionPredicate { return func(e Int64Option) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteOptionPredicate) Xor(p2 ByteOptionPredicate) ByteOptionPredicate { return func(e ByteOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneOptionPredicate) Xor(p2 RuneOptionPredicate) RuneOptionPredicate { return func(e RuneOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32OptionPredicate) Xor(p2 Float32OptionPredicate) Float32OptionPredicate { return func(e Float32Option) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64OptionPredicate) Xor(p2 Float64OptionPredicate) Float64OptionPredicate { return func(e Float64Option) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyOptionPredicate) Xor(p2 AnyOptionPredicate) AnyOptionPredicate { return func(e AnyOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2OptionPredicate) Xor(p2 Tuple2OptionPredicate) Tuple2OptionPredicate { return func(e Tuple2Option) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolOptionOptionPredicate) Xor(p2 BoolOptionOptionPredicate) BoolOptionOptionPredicate { return func(e BoolOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringOptionOptionPredicate) Xor(p2 StringOptionOptionPredicate) StringOptionOptionPredicate { return func(e StringOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntOptionOptionPredicate) Xor(p2 IntOptionOptionPredicate) IntOptionOptionPredicate { return func(e IntOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64OptionOptionPredicate) Xor(p2 Int64OptionOptionPredicate) Int64OptionOptionPredicate { return func(e Int64OptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteOptionOptionPredicate) Xor(p2 ByteOptionOptionPredicate) ByteOptionOptionPredicate { return func(e ByteOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneOptionOptionPredicate) Xor(p2 RuneOptionOptionPredicate) RuneOptionOptionPredicate { return func(e RuneOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32OptionOptionPredicate) Xor(p2 Float32OptionOptionPredicate) Float32OptionOptionPredicate { return func(e Float32OptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64OptionOptionPredicate) Xor(p2 Float64OptionOptionPredicate) Float64OptionOptionPredicate { return func(e Float64OptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyOptionOptionPredicate) Xor(p2 AnyOptionOptionPredicate) AnyOptionOptionPredicate { return func(e AnyOptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2OptionOptionPredicate) Xor(p2 Tuple2OptionOptionPredicate) Tuple2OptionOptionPredicate { return func(e Tuple2OptionOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolArrayOptionPredicate) Xor(p2 BoolArrayOptionPredicate) BoolArrayOptionPredicate { return func(e BoolArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringArrayOptionPredicate) Xor(p2 StringArrayOptionPredicate) StringArrayOptionPredicate { return func(e StringArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntArrayOptionPredicate) Xor(p2 IntArrayOptionPredicate) IntArrayOptionPredicate { return func(e IntArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ArrayOptionPredicate) Xor(p2 Int64ArrayOptionPredicate) Int64ArrayOptionPredicate { return func(e Int64ArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteArrayOptionPredicate) Xor(p2 ByteArrayOptionPredicate) ByteArrayOptionPredicate { return func(e ByteArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneArrayOptionPredicate) Xor(p2 RuneArrayOptionPredicate) RuneArrayOptionPredicate { return func(e RuneArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ArrayOptionPredicate) Xor(p2 Float32ArrayOptionPredicate) Float32ArrayOptionPredicate { return func(e Float32ArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ArrayOptionPredicate) Xor(p2 Float64ArrayOptionPredicate) Float64ArrayOptionPredicate { return func(e Float64ArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyArrayOptionPredicate) Xor(p2 AnyArrayOptionPredicate) AnyArrayOptionPredicate { return func(e AnyArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ArrayOptionPredicate) Xor(p2 Tuple2ArrayOptionPredicate) Tuple2ArrayOptionPredicate { return func(e Tuple2ArrayOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolListOptionPredicate) Xor(p2 BoolListOptionPredicate) BoolListOptionPredicate { return func(e BoolListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringListOptionPredicate) Xor(p2 StringListOptionPredicate) StringListOptionPredicate { return func(e StringListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntListOptionPredicate) Xor(p2 IntListOptionPredicate) IntListOptionPredicate { return func(e IntListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ListOptionPredicate) Xor(p2 Int64ListOptionPredicate) Int64ListOptionPredicate { return func(e Int64ListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteListOptionPredicate) Xor(p2 ByteListOptionPredicate) ByteListOptionPredicate { return func(e ByteListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneListOptionPredicate) Xor(p2 RuneListOptionPredicate) RuneListOptionPredicate { return func(e RuneListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ListOptionPredicate) Xor(p2 Float32ListOptionPredicate) Float32ListOptionPredicate { return func(e Float32ListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ListOptionPredicate) Xor(p2 Float64ListOptionPredicate) Float64ListOptionPredicate { return func(e Float64ListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyListOptionPredicate) Xor(p2 AnyListOptionPredicate) AnyListOptionPredicate { return func(e AnyListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ListOptionPredicate) Xor(p2 Tuple2ListOptionPredicate) Tuple2ListOptionPredicate { return func(e Tuple2ListOption) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolListPredicate) Xor(p2 BoolListPredicate) BoolListPredicate { return func(e BoolList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringListPredicate) Xor(p2 StringListPredicate) StringListPredicate { return func(e StringList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntListPredicate) Xor(p2 IntListPredicate) IntListPredicate { return func(e IntList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ListPredicate) Xor(p2 Int64ListPredicate) Int64ListPredicate { return func(e Int64List) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteListPredicate) Xor(p2 ByteListPredicate) ByteListPredicate { return func(e ByteList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneListPredicate) Xor(p2 RuneListPredicate) RuneListPredicate { return func(e RuneList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ListPredicate) Xor(p2 Float32ListPredicate) Float32ListPredicate { return func(e Float32List) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ListPredicate) Xor(p2 Float64ListPredicate) Float64ListPredicate { return func(e Float64List) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyListPredicate) Xor(p2 AnyListPredicate) AnyListPredicate { return func(e AnyList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ListPredicate) Xor(p2 Tuple2ListPredicate) Tuple2ListPredicate { return func(e Tuple2List) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolArrayListPredicate) Xor(p2 BoolArrayListPredicate) BoolArrayListPredicate { return func(e BoolArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringArrayListPredicate) Xor(p2 StringArrayListPredicate) StringArrayListPredicate { return func(e StringArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntArrayListPredicate) Xor(p2 IntArrayListPredicate) IntArrayListPredicate { return func(e IntArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ArrayListPredicate) Xor(p2 Int64ArrayListPredicate) Int64ArrayListPredicate { return func(e Int64ArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteArrayListPredicate) Xor(p2 ByteArrayListPredicate) ByteArrayListPredicate { return func(e ByteArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneArrayListPredicate) Xor(p2 RuneArrayListPredicate) RuneArrayListPredicate { return func(e RuneArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ArrayListPredicate) Xor(p2 Float32ArrayListPredicate) Float32ArrayListPredicate { return func(e Float32ArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ArrayListPredicate) Xor(p2 Float64ArrayListPredicate) Float64ArrayListPredicate { return func(e Float64ArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyArrayListPredicate) Xor(p2 AnyArrayListPredicate) AnyArrayListPredicate { return func(e AnyArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ArrayListPredicate) Xor(p2 Tuple2ArrayListPredicate) Tuple2ArrayListPredicate { return func(e Tuple2ArrayList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolOptionListPredicate) Xor(p2 BoolOptionListPredicate) BoolOptionListPredicate { return func(e BoolOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringOptionListPredicate) Xor(p2 StringOptionListPredicate) StringOptionListPredicate { return func(e StringOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntOptionListPredicate) Xor(p2 IntOptionListPredicate) IntOptionListPredicate { return func(e IntOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64OptionListPredicate) Xor(p2 Int64OptionListPredicate) Int64OptionListPredicate { return func(e Int64OptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteOptionListPredicate) Xor(p2 ByteOptionListPredicate) ByteOptionListPredicate { return func(e ByteOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneOptionListPredicate) Xor(p2 RuneOptionListPredicate) RuneOptionListPredicate { return func(e RuneOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32OptionListPredicate) Xor(p2 Float32OptionListPredicate) Float32OptionListPredicate { return func(e Float32OptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64OptionListPredicate) Xor(p2 Float64OptionListPredicate) Float64OptionListPredicate { return func(e Float64OptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyOptionListPredicate) Xor(p2 AnyOptionListPredicate) AnyOptionListPredicate { return func(e AnyOptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2OptionListPredicate) Xor(p2 Tuple2OptionListPredicate) Tuple2OptionListPredicate { return func(e Tuple2OptionList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 BoolListListPredicate) Xor(p2 BoolListListPredicate) BoolListListPredicate { return func(e BoolListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 StringListListPredicate) Xor(p2 StringListListPredicate) StringListListPredicate { return func(e StringListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 IntListListPredicate) Xor(p2 IntListListPredicate) IntListListPredicate { return func(e IntListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Int64ListListPredicate) Xor(p2 Int64ListListPredicate) Int64ListListPredicate { return func(e Int64ListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 ByteListListPredicate) Xor(p2 ByteListListPredicate) ByteListListPredicate { return func(e ByteListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 RuneListListPredicate) Xor(p2 RuneListListPredicate) RuneListListPredicate { return func(e RuneListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float32ListListPredicate) Xor(p2 Float32ListListPredicate) Float32ListListPredicate { return func(e Float32ListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Float64ListListPredicate) Xor(p2 Float64ListListPredicate) Float64ListListPredicate { return func(e Float64ListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 AnyListListPredicate) Xor(p2 AnyListListPredicate) AnyListListPredicate { return func(e AnyListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2ListListPredicate) Xor(p2 Tuple2ListListPredicate) Tuple2ListListPredicate { return func(e Tuple2ListList) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }
func (p1 Tuple2Predicate) Xor(p2 Tuple2Predicate) Tuple2Predicate { return func(e Tuple2) bool { x := p1(e); y := p2(e); return (x || y) && !(x && y) } }