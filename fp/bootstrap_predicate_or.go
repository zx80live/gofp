// bootstrap_predicate_or.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (p1 BoolPredicate) Or(p2 BoolPredicate) BoolPredicate { return func(e bool) bool { return p1(e) || p2(e) } }
func (p1 StringPredicate) Or(p2 StringPredicate) StringPredicate { return func(e string) bool { return p1(e) || p2(e) } }
func (p1 IntPredicate) Or(p2 IntPredicate) IntPredicate { return func(e int) bool { return p1(e) || p2(e) } }
func (p1 Int64Predicate) Or(p2 Int64Predicate) Int64Predicate { return func(e int64) bool { return p1(e) || p2(e) } }
func (p1 BytePredicate) Or(p2 BytePredicate) BytePredicate { return func(e byte) bool { return p1(e) || p2(e) } }
func (p1 RunePredicate) Or(p2 RunePredicate) RunePredicate { return func(e rune) bool { return p1(e) || p2(e) } }
func (p1 Float32Predicate) Or(p2 Float32Predicate) Float32Predicate { return func(e float32) bool { return p1(e) || p2(e) } }
func (p1 Float64Predicate) Or(p2 Float64Predicate) Float64Predicate { return func(e float64) bool { return p1(e) || p2(e) } }
func (p1 AnyPredicate) Or(p2 AnyPredicate) AnyPredicate { return func(e Any) bool { return p1(e) || p2(e) } }
func (p1 BoolArrayPredicate) Or(p2 BoolArrayPredicate) BoolArrayPredicate { return func(e []bool) bool { return p1(e) || p2(e) } }
func (p1 StringArrayPredicate) Or(p2 StringArrayPredicate) StringArrayPredicate { return func(e []string) bool { return p1(e) || p2(e) } }
func (p1 IntArrayPredicate) Or(p2 IntArrayPredicate) IntArrayPredicate { return func(e []int) bool { return p1(e) || p2(e) } }
func (p1 Int64ArrayPredicate) Or(p2 Int64ArrayPredicate) Int64ArrayPredicate { return func(e []int64) bool { return p1(e) || p2(e) } }
func (p1 ByteArrayPredicate) Or(p2 ByteArrayPredicate) ByteArrayPredicate { return func(e []byte) bool { return p1(e) || p2(e) } }
func (p1 RuneArrayPredicate) Or(p2 RuneArrayPredicate) RuneArrayPredicate { return func(e []rune) bool { return p1(e) || p2(e) } }
func (p1 Float32ArrayPredicate) Or(p2 Float32ArrayPredicate) Float32ArrayPredicate { return func(e []float32) bool { return p1(e) || p2(e) } }
func (p1 Float64ArrayPredicate) Or(p2 Float64ArrayPredicate) Float64ArrayPredicate { return func(e []float64) bool { return p1(e) || p2(e) } }
func (p1 AnyArrayPredicate) Or(p2 AnyArrayPredicate) AnyArrayPredicate { return func(e []Any) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ArrayPredicate) Or(p2 Tuple2ArrayPredicate) Tuple2ArrayPredicate { return func(e []Tuple2) bool { return p1(e) || p2(e) } }
func (p1 BoolArrayArrayPredicate) Or(p2 BoolArrayArrayPredicate) BoolArrayArrayPredicate { return func(e [][]bool) bool { return p1(e) || p2(e) } }
func (p1 StringArrayArrayPredicate) Or(p2 StringArrayArrayPredicate) StringArrayArrayPredicate { return func(e [][]string) bool { return p1(e) || p2(e) } }
func (p1 IntArrayArrayPredicate) Or(p2 IntArrayArrayPredicate) IntArrayArrayPredicate { return func(e [][]int) bool { return p1(e) || p2(e) } }
func (p1 Int64ArrayArrayPredicate) Or(p2 Int64ArrayArrayPredicate) Int64ArrayArrayPredicate { return func(e [][]int64) bool { return p1(e) || p2(e) } }
func (p1 ByteArrayArrayPredicate) Or(p2 ByteArrayArrayPredicate) ByteArrayArrayPredicate { return func(e [][]byte) bool { return p1(e) || p2(e) } }
func (p1 RuneArrayArrayPredicate) Or(p2 RuneArrayArrayPredicate) RuneArrayArrayPredicate { return func(e [][]rune) bool { return p1(e) || p2(e) } }
func (p1 Float32ArrayArrayPredicate) Or(p2 Float32ArrayArrayPredicate) Float32ArrayArrayPredicate { return func(e [][]float32) bool { return p1(e) || p2(e) } }
func (p1 Float64ArrayArrayPredicate) Or(p2 Float64ArrayArrayPredicate) Float64ArrayArrayPredicate { return func(e [][]float64) bool { return p1(e) || p2(e) } }
func (p1 AnyArrayArrayPredicate) Or(p2 AnyArrayArrayPredicate) AnyArrayArrayPredicate { return func(e [][]Any) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ArrayArrayPredicate) Or(p2 Tuple2ArrayArrayPredicate) Tuple2ArrayArrayPredicate { return func(e [][]Tuple2) bool { return p1(e) || p2(e) } }
func (p1 BoolOptionPredicate) Or(p2 BoolOptionPredicate) BoolOptionPredicate { return func(e BoolOption) bool { return p1(e) || p2(e) } }
func (p1 StringOptionPredicate) Or(p2 StringOptionPredicate) StringOptionPredicate { return func(e StringOption) bool { return p1(e) || p2(e) } }
func (p1 IntOptionPredicate) Or(p2 IntOptionPredicate) IntOptionPredicate { return func(e IntOption) bool { return p1(e) || p2(e) } }
func (p1 Int64OptionPredicate) Or(p2 Int64OptionPredicate) Int64OptionPredicate { return func(e Int64Option) bool { return p1(e) || p2(e) } }
func (p1 ByteOptionPredicate) Or(p2 ByteOptionPredicate) ByteOptionPredicate { return func(e ByteOption) bool { return p1(e) || p2(e) } }
func (p1 RuneOptionPredicate) Or(p2 RuneOptionPredicate) RuneOptionPredicate { return func(e RuneOption) bool { return p1(e) || p2(e) } }
func (p1 Float32OptionPredicate) Or(p2 Float32OptionPredicate) Float32OptionPredicate { return func(e Float32Option) bool { return p1(e) || p2(e) } }
func (p1 Float64OptionPredicate) Or(p2 Float64OptionPredicate) Float64OptionPredicate { return func(e Float64Option) bool { return p1(e) || p2(e) } }
func (p1 AnyOptionPredicate) Or(p2 AnyOptionPredicate) AnyOptionPredicate { return func(e AnyOption) bool { return p1(e) || p2(e) } }
func (p1 Tuple2OptionPredicate) Or(p2 Tuple2OptionPredicate) Tuple2OptionPredicate { return func(e Tuple2Option) bool { return p1(e) || p2(e) } }
func (p1 BoolOptionOptionPredicate) Or(p2 BoolOptionOptionPredicate) BoolOptionOptionPredicate { return func(e BoolOptionOption) bool { return p1(e) || p2(e) } }
func (p1 StringOptionOptionPredicate) Or(p2 StringOptionOptionPredicate) StringOptionOptionPredicate { return func(e StringOptionOption) bool { return p1(e) || p2(e) } }
func (p1 IntOptionOptionPredicate) Or(p2 IntOptionOptionPredicate) IntOptionOptionPredicate { return func(e IntOptionOption) bool { return p1(e) || p2(e) } }
func (p1 Int64OptionOptionPredicate) Or(p2 Int64OptionOptionPredicate) Int64OptionOptionPredicate { return func(e Int64OptionOption) bool { return p1(e) || p2(e) } }
func (p1 ByteOptionOptionPredicate) Or(p2 ByteOptionOptionPredicate) ByteOptionOptionPredicate { return func(e ByteOptionOption) bool { return p1(e) || p2(e) } }
func (p1 RuneOptionOptionPredicate) Or(p2 RuneOptionOptionPredicate) RuneOptionOptionPredicate { return func(e RuneOptionOption) bool { return p1(e) || p2(e) } }
func (p1 Float32OptionOptionPredicate) Or(p2 Float32OptionOptionPredicate) Float32OptionOptionPredicate { return func(e Float32OptionOption) bool { return p1(e) || p2(e) } }
func (p1 Float64OptionOptionPredicate) Or(p2 Float64OptionOptionPredicate) Float64OptionOptionPredicate { return func(e Float64OptionOption) bool { return p1(e) || p2(e) } }
func (p1 AnyOptionOptionPredicate) Or(p2 AnyOptionOptionPredicate) AnyOptionOptionPredicate { return func(e AnyOptionOption) bool { return p1(e) || p2(e) } }
func (p1 Tuple2OptionOptionPredicate) Or(p2 Tuple2OptionOptionPredicate) Tuple2OptionOptionPredicate { return func(e Tuple2OptionOption) bool { return p1(e) || p2(e) } }
func (p1 BoolArrayOptionPredicate) Or(p2 BoolArrayOptionPredicate) BoolArrayOptionPredicate { return func(e BoolArrayOption) bool { return p1(e) || p2(e) } }
func (p1 StringArrayOptionPredicate) Or(p2 StringArrayOptionPredicate) StringArrayOptionPredicate { return func(e StringArrayOption) bool { return p1(e) || p2(e) } }
func (p1 IntArrayOptionPredicate) Or(p2 IntArrayOptionPredicate) IntArrayOptionPredicate { return func(e IntArrayOption) bool { return p1(e) || p2(e) } }
func (p1 Int64ArrayOptionPredicate) Or(p2 Int64ArrayOptionPredicate) Int64ArrayOptionPredicate { return func(e Int64ArrayOption) bool { return p1(e) || p2(e) } }
func (p1 ByteArrayOptionPredicate) Or(p2 ByteArrayOptionPredicate) ByteArrayOptionPredicate { return func(e ByteArrayOption) bool { return p1(e) || p2(e) } }
func (p1 RuneArrayOptionPredicate) Or(p2 RuneArrayOptionPredicate) RuneArrayOptionPredicate { return func(e RuneArrayOption) bool { return p1(e) || p2(e) } }
func (p1 Float32ArrayOptionPredicate) Or(p2 Float32ArrayOptionPredicate) Float32ArrayOptionPredicate { return func(e Float32ArrayOption) bool { return p1(e) || p2(e) } }
func (p1 Float64ArrayOptionPredicate) Or(p2 Float64ArrayOptionPredicate) Float64ArrayOptionPredicate { return func(e Float64ArrayOption) bool { return p1(e) || p2(e) } }
func (p1 AnyArrayOptionPredicate) Or(p2 AnyArrayOptionPredicate) AnyArrayOptionPredicate { return func(e AnyArrayOption) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ArrayOptionPredicate) Or(p2 Tuple2ArrayOptionPredicate) Tuple2ArrayOptionPredicate { return func(e Tuple2ArrayOption) bool { return p1(e) || p2(e) } }
func (p1 BoolListOptionPredicate) Or(p2 BoolListOptionPredicate) BoolListOptionPredicate { return func(e BoolListOption) bool { return p1(e) || p2(e) } }
func (p1 StringListOptionPredicate) Or(p2 StringListOptionPredicate) StringListOptionPredicate { return func(e StringListOption) bool { return p1(e) || p2(e) } }
func (p1 IntListOptionPredicate) Or(p2 IntListOptionPredicate) IntListOptionPredicate { return func(e IntListOption) bool { return p1(e) || p2(e) } }
func (p1 Int64ListOptionPredicate) Or(p2 Int64ListOptionPredicate) Int64ListOptionPredicate { return func(e Int64ListOption) bool { return p1(e) || p2(e) } }
func (p1 ByteListOptionPredicate) Or(p2 ByteListOptionPredicate) ByteListOptionPredicate { return func(e ByteListOption) bool { return p1(e) || p2(e) } }
func (p1 RuneListOptionPredicate) Or(p2 RuneListOptionPredicate) RuneListOptionPredicate { return func(e RuneListOption) bool { return p1(e) || p2(e) } }
func (p1 Float32ListOptionPredicate) Or(p2 Float32ListOptionPredicate) Float32ListOptionPredicate { return func(e Float32ListOption) bool { return p1(e) || p2(e) } }
func (p1 Float64ListOptionPredicate) Or(p2 Float64ListOptionPredicate) Float64ListOptionPredicate { return func(e Float64ListOption) bool { return p1(e) || p2(e) } }
func (p1 AnyListOptionPredicate) Or(p2 AnyListOptionPredicate) AnyListOptionPredicate { return func(e AnyListOption) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ListOptionPredicate) Or(p2 Tuple2ListOptionPredicate) Tuple2ListOptionPredicate { return func(e Tuple2ListOption) bool { return p1(e) || p2(e) } }
func (p1 BoolListPredicate) Or(p2 BoolListPredicate) BoolListPredicate { return func(e BoolList) bool { return p1(e) || p2(e) } }
func (p1 StringListPredicate) Or(p2 StringListPredicate) StringListPredicate { return func(e StringList) bool { return p1(e) || p2(e) } }
func (p1 IntListPredicate) Or(p2 IntListPredicate) IntListPredicate { return func(e IntList) bool { return p1(e) || p2(e) } }
func (p1 Int64ListPredicate) Or(p2 Int64ListPredicate) Int64ListPredicate { return func(e Int64List) bool { return p1(e) || p2(e) } }
func (p1 ByteListPredicate) Or(p2 ByteListPredicate) ByteListPredicate { return func(e ByteList) bool { return p1(e) || p2(e) } }
func (p1 RuneListPredicate) Or(p2 RuneListPredicate) RuneListPredicate { return func(e RuneList) bool { return p1(e) || p2(e) } }
func (p1 Float32ListPredicate) Or(p2 Float32ListPredicate) Float32ListPredicate { return func(e Float32List) bool { return p1(e) || p2(e) } }
func (p1 Float64ListPredicate) Or(p2 Float64ListPredicate) Float64ListPredicate { return func(e Float64List) bool { return p1(e) || p2(e) } }
func (p1 AnyListPredicate) Or(p2 AnyListPredicate) AnyListPredicate { return func(e AnyList) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ListPredicate) Or(p2 Tuple2ListPredicate) Tuple2ListPredicate { return func(e Tuple2List) bool { return p1(e) || p2(e) } }
func (p1 BoolArrayListPredicate) Or(p2 BoolArrayListPredicate) BoolArrayListPredicate { return func(e BoolArrayList) bool { return p1(e) || p2(e) } }
func (p1 StringArrayListPredicate) Or(p2 StringArrayListPredicate) StringArrayListPredicate { return func(e StringArrayList) bool { return p1(e) || p2(e) } }
func (p1 IntArrayListPredicate) Or(p2 IntArrayListPredicate) IntArrayListPredicate { return func(e IntArrayList) bool { return p1(e) || p2(e) } }
func (p1 Int64ArrayListPredicate) Or(p2 Int64ArrayListPredicate) Int64ArrayListPredicate { return func(e Int64ArrayList) bool { return p1(e) || p2(e) } }
func (p1 ByteArrayListPredicate) Or(p2 ByteArrayListPredicate) ByteArrayListPredicate { return func(e ByteArrayList) bool { return p1(e) || p2(e) } }
func (p1 RuneArrayListPredicate) Or(p2 RuneArrayListPredicate) RuneArrayListPredicate { return func(e RuneArrayList) bool { return p1(e) || p2(e) } }
func (p1 Float32ArrayListPredicate) Or(p2 Float32ArrayListPredicate) Float32ArrayListPredicate { return func(e Float32ArrayList) bool { return p1(e) || p2(e) } }
func (p1 Float64ArrayListPredicate) Or(p2 Float64ArrayListPredicate) Float64ArrayListPredicate { return func(e Float64ArrayList) bool { return p1(e) || p2(e) } }
func (p1 AnyArrayListPredicate) Or(p2 AnyArrayListPredicate) AnyArrayListPredicate { return func(e AnyArrayList) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ArrayListPredicate) Or(p2 Tuple2ArrayListPredicate) Tuple2ArrayListPredicate { return func(e Tuple2ArrayList) bool { return p1(e) || p2(e) } }
func (p1 BoolOptionListPredicate) Or(p2 BoolOptionListPredicate) BoolOptionListPredicate { return func(e BoolOptionList) bool { return p1(e) || p2(e) } }
func (p1 StringOptionListPredicate) Or(p2 StringOptionListPredicate) StringOptionListPredicate { return func(e StringOptionList) bool { return p1(e) || p2(e) } }
func (p1 IntOptionListPredicate) Or(p2 IntOptionListPredicate) IntOptionListPredicate { return func(e IntOptionList) bool { return p1(e) || p2(e) } }
func (p1 Int64OptionListPredicate) Or(p2 Int64OptionListPredicate) Int64OptionListPredicate { return func(e Int64OptionList) bool { return p1(e) || p2(e) } }
func (p1 ByteOptionListPredicate) Or(p2 ByteOptionListPredicate) ByteOptionListPredicate { return func(e ByteOptionList) bool { return p1(e) || p2(e) } }
func (p1 RuneOptionListPredicate) Or(p2 RuneOptionListPredicate) RuneOptionListPredicate { return func(e RuneOptionList) bool { return p1(e) || p2(e) } }
func (p1 Float32OptionListPredicate) Or(p2 Float32OptionListPredicate) Float32OptionListPredicate { return func(e Float32OptionList) bool { return p1(e) || p2(e) } }
func (p1 Float64OptionListPredicate) Or(p2 Float64OptionListPredicate) Float64OptionListPredicate { return func(e Float64OptionList) bool { return p1(e) || p2(e) } }
func (p1 AnyOptionListPredicate) Or(p2 AnyOptionListPredicate) AnyOptionListPredicate { return func(e AnyOptionList) bool { return p1(e) || p2(e) } }
func (p1 Tuple2OptionListPredicate) Or(p2 Tuple2OptionListPredicate) Tuple2OptionListPredicate { return func(e Tuple2OptionList) bool { return p1(e) || p2(e) } }
func (p1 BoolListListPredicate) Or(p2 BoolListListPredicate) BoolListListPredicate { return func(e BoolListList) bool { return p1(e) || p2(e) } }
func (p1 StringListListPredicate) Or(p2 StringListListPredicate) StringListListPredicate { return func(e StringListList) bool { return p1(e) || p2(e) } }
func (p1 IntListListPredicate) Or(p2 IntListListPredicate) IntListListPredicate { return func(e IntListList) bool { return p1(e) || p2(e) } }
func (p1 Int64ListListPredicate) Or(p2 Int64ListListPredicate) Int64ListListPredicate { return func(e Int64ListList) bool { return p1(e) || p2(e) } }
func (p1 ByteListListPredicate) Or(p2 ByteListListPredicate) ByteListListPredicate { return func(e ByteListList) bool { return p1(e) || p2(e) } }
func (p1 RuneListListPredicate) Or(p2 RuneListListPredicate) RuneListListPredicate { return func(e RuneListList) bool { return p1(e) || p2(e) } }
func (p1 Float32ListListPredicate) Or(p2 Float32ListListPredicate) Float32ListListPredicate { return func(e Float32ListList) bool { return p1(e) || p2(e) } }
func (p1 Float64ListListPredicate) Or(p2 Float64ListListPredicate) Float64ListListPredicate { return func(e Float64ListList) bool { return p1(e) || p2(e) } }
func (p1 AnyListListPredicate) Or(p2 AnyListListPredicate) AnyListListPredicate { return func(e AnyListList) bool { return p1(e) || p2(e) } }
func (p1 Tuple2ListListPredicate) Or(p2 Tuple2ListListPredicate) Tuple2ListListPredicate { return func(e Tuple2ListList) bool { return p1(e) || p2(e) } }
func (p1 Tuple2Predicate) Or(p2 Tuple2Predicate) Tuple2Predicate { return func(e Tuple2) bool { return p1(e) || p2(e) } }