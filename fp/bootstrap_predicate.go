// bootstrap_predicate.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


type BoolPredicate func(t bool) bool
type StringPredicate func(t string) bool
type IntPredicate func(t int) bool
type Int8Predicate func(t int8) bool
type Int16Predicate func(t int16) bool
type Int32Predicate func(t int32) bool
type Int64Predicate func(t int64) bool
type UintPredicate func(t uint) bool
type Uint8Predicate func(t uint8) bool
type Uint16Predicate func(t uint16) bool
type Uint32Predicate func(t uint32) bool
type Uint64Predicate func(t uint64) bool
type UintptrPredicate func(t uintptr) bool
type BytePredicate func(t byte) bool
type RunePredicate func(t rune) bool
type Float32Predicate func(t float32) bool
type Float64Predicate func(t float64) bool
type Complex64Predicate func(t complex64) bool
type Complex128Predicate func(t complex128) bool
type AnyPredicate func(t Any) bool
type BoolArrPredicate func(t []bool) bool
type StringArrPredicate func(t []string) bool
type IntArrPredicate func(t []int) bool
type Int8ArrPredicate func(t []int8) bool
type Int16ArrPredicate func(t []int16) bool
type Int32ArrPredicate func(t []int32) bool
type Int64ArrPredicate func(t []int64) bool
type UintArrPredicate func(t []uint) bool
type Uint8ArrPredicate func(t []uint8) bool
type Uint16ArrPredicate func(t []uint16) bool
type Uint32ArrPredicate func(t []uint32) bool
type Uint64ArrPredicate func(t []uint64) bool
type UintptrArrPredicate func(t []uintptr) bool
type ByteArrPredicate func(t []byte) bool
type RuneArrPredicate func(t []rune) bool
type Float32ArrPredicate func(t []float32) bool
type Float64ArrPredicate func(t []float64) bool
type Complex64ArrPredicate func(t []complex64) bool
type Complex128ArrPredicate func(t []complex128) bool
type AnyArrPredicate func(t []Any) bool
type BoolArrArrPredicate func(t [][]bool) bool
type StringArrArrPredicate func(t [][]string) bool
type IntArrArrPredicate func(t [][]int) bool
type Int8ArrArrPredicate func(t [][]int8) bool
type Int16ArrArrPredicate func(t [][]int16) bool
type Int32ArrArrPredicate func(t [][]int32) bool
type Int64ArrArrPredicate func(t [][]int64) bool
type UintArrArrPredicate func(t [][]uint) bool
type Uint8ArrArrPredicate func(t [][]uint8) bool
type Uint16ArrArrPredicate func(t [][]uint16) bool
type Uint32ArrArrPredicate func(t [][]uint32) bool
type Uint64ArrArrPredicate func(t [][]uint64) bool
type UintptrArrArrPredicate func(t [][]uintptr) bool
type ByteArrArrPredicate func(t [][]byte) bool
type RuneArrArrPredicate func(t [][]rune) bool
type Float32ArrArrPredicate func(t [][]float32) bool
type Float64ArrArrPredicate func(t [][]float64) bool
type Complex64ArrArrPredicate func(t [][]complex64) bool
type Complex128ArrArrPredicate func(t [][]complex128) bool
type AnyArrArrPredicate func(t [][]Any) bool
type BoolOptionArrPredicate func(t []BoolOption) bool
type StringOptionArrPredicate func(t []StringOption) bool
type IntOptionArrPredicate func(t []IntOption) bool
type Int8OptionArrPredicate func(t []Int8Option) bool
type Int16OptionArrPredicate func(t []Int16Option) bool
type Int32OptionArrPredicate func(t []Int32Option) bool
type Int64OptionArrPredicate func(t []Int64Option) bool
type UintOptionArrPredicate func(t []UintOption) bool
type Uint8OptionArrPredicate func(t []Uint8Option) bool
type Uint16OptionArrPredicate func(t []Uint16Option) bool
type Uint32OptionArrPredicate func(t []Uint32Option) bool
type Uint64OptionArrPredicate func(t []Uint64Option) bool
type UintptrOptionArrPredicate func(t []UintptrOption) bool
type ByteOptionArrPredicate func(t []ByteOption) bool
type RuneOptionArrPredicate func(t []RuneOption) bool
type Float32OptionArrPredicate func(t []Float32Option) bool
type Float64OptionArrPredicate func(t []Float64Option) bool
type Complex64OptionArrPredicate func(t []Complex64Option) bool
type Complex128OptionArrPredicate func(t []Complex128Option) bool
type AnyOptionArrPredicate func(t []AnyOption) bool
type BoolOptionPredicate func(t BoolOption) bool
type StringOptionPredicate func(t StringOption) bool
type IntOptionPredicate func(t IntOption) bool
type Int8OptionPredicate func(t Int8Option) bool
type Int16OptionPredicate func(t Int16Option) bool
type Int32OptionPredicate func(t Int32Option) bool
type Int64OptionPredicate func(t Int64Option) bool
type UintOptionPredicate func(t UintOption) bool
type Uint8OptionPredicate func(t Uint8Option) bool
type Uint16OptionPredicate func(t Uint16Option) bool
type Uint32OptionPredicate func(t Uint32Option) bool
type Uint64OptionPredicate func(t Uint64Option) bool
type UintptrOptionPredicate func(t UintptrOption) bool
type ByteOptionPredicate func(t ByteOption) bool
type RuneOptionPredicate func(t RuneOption) bool
type Float32OptionPredicate func(t Float32Option) bool
type Float64OptionPredicate func(t Float64Option) bool
type Complex64OptionPredicate func(t Complex64Option) bool
type Complex128OptionPredicate func(t Complex128Option) bool
type AnyOptionPredicate func(t AnyOption) bool
type BoolArrOptionPredicate func(t BoolArrOption) bool
type StringArrOptionPredicate func(t StringArrOption) bool
type IntArrOptionPredicate func(t IntArrOption) bool
type Int8ArrOptionPredicate func(t Int8ArrOption) bool
type Int16ArrOptionPredicate func(t Int16ArrOption) bool
type Int32ArrOptionPredicate func(t Int32ArrOption) bool
type Int64ArrOptionPredicate func(t Int64ArrOption) bool
type UintArrOptionPredicate func(t UintArrOption) bool
type Uint8ArrOptionPredicate func(t Uint8ArrOption) bool
type Uint16ArrOptionPredicate func(t Uint16ArrOption) bool
type Uint32ArrOptionPredicate func(t Uint32ArrOption) bool
type Uint64ArrOptionPredicate func(t Uint64ArrOption) bool
type UintptrArrOptionPredicate func(t UintptrArrOption) bool
type ByteArrOptionPredicate func(t ByteArrOption) bool
type RuneArrOptionPredicate func(t RuneArrOption) bool
type Float32ArrOptionPredicate func(t Float32ArrOption) bool
type Float64ArrOptionPredicate func(t Float64ArrOption) bool
type Complex64ArrOptionPredicate func(t Complex64ArrOption) bool
type Complex128ArrOptionPredicate func(t Complex128ArrOption) bool
type AnyArrOptionPredicate func(t AnyArrOption) bool
type BoolArrArrOptionPredicate func(t BoolArrArrOption) bool
type StringArrArrOptionPredicate func(t StringArrArrOption) bool
type IntArrArrOptionPredicate func(t IntArrArrOption) bool
type Int8ArrArrOptionPredicate func(t Int8ArrArrOption) bool
type Int16ArrArrOptionPredicate func(t Int16ArrArrOption) bool
type Int32ArrArrOptionPredicate func(t Int32ArrArrOption) bool
type Int64ArrArrOptionPredicate func(t Int64ArrArrOption) bool
type UintArrArrOptionPredicate func(t UintArrArrOption) bool
type Uint8ArrArrOptionPredicate func(t Uint8ArrArrOption) bool
type Uint16ArrArrOptionPredicate func(t Uint16ArrArrOption) bool
type Uint32ArrArrOptionPredicate func(t Uint32ArrArrOption) bool
type Uint64ArrArrOptionPredicate func(t Uint64ArrArrOption) bool
type UintptrArrArrOptionPredicate func(t UintptrArrArrOption) bool
type ByteArrArrOptionPredicate func(t ByteArrArrOption) bool
type RuneArrArrOptionPredicate func(t RuneArrArrOption) bool
type Float32ArrArrOptionPredicate func(t Float32ArrArrOption) bool
type Float64ArrArrOptionPredicate func(t Float64ArrArrOption) bool
type Complex64ArrArrOptionPredicate func(t Complex64ArrArrOption) bool
type Complex128ArrArrOptionPredicate func(t Complex128ArrArrOption) bool
type AnyArrArrOptionPredicate func(t AnyArrArrOption) bool
type BoolOptionArrOptionPredicate func(t BoolOptionArrOption) bool
type StringOptionArrOptionPredicate func(t StringOptionArrOption) bool
type IntOptionArrOptionPredicate func(t IntOptionArrOption) bool
type Int8OptionArrOptionPredicate func(t Int8OptionArrOption) bool
type Int16OptionArrOptionPredicate func(t Int16OptionArrOption) bool
type Int32OptionArrOptionPredicate func(t Int32OptionArrOption) bool
type Int64OptionArrOptionPredicate func(t Int64OptionArrOption) bool
type UintOptionArrOptionPredicate func(t UintOptionArrOption) bool
type Uint8OptionArrOptionPredicate func(t Uint8OptionArrOption) bool
type Uint16OptionArrOptionPredicate func(t Uint16OptionArrOption) bool
type Uint32OptionArrOptionPredicate func(t Uint32OptionArrOption) bool
type Uint64OptionArrOptionPredicate func(t Uint64OptionArrOption) bool
type UintptrOptionArrOptionPredicate func(t UintptrOptionArrOption) bool
type ByteOptionArrOptionPredicate func(t ByteOptionArrOption) bool
type RuneOptionArrOptionPredicate func(t RuneOptionArrOption) bool
type Float32OptionArrOptionPredicate func(t Float32OptionArrOption) bool
type Float64OptionArrOptionPredicate func(t Float64OptionArrOption) bool
type Complex64OptionArrOptionPredicate func(t Complex64OptionArrOption) bool
type Complex128OptionArrOptionPredicate func(t Complex128OptionArrOption) bool
type AnyOptionArrOptionPredicate func(t AnyOptionArrOption) bool
type BoolListOptionPredicate func(t BoolListOption) bool
type StringListOptionPredicate func(t StringListOption) bool
type IntListOptionPredicate func(t IntListOption) bool
type Int8ListOptionPredicate func(t Int8ListOption) bool
type Int16ListOptionPredicate func(t Int16ListOption) bool
type Int32ListOptionPredicate func(t Int32ListOption) bool
type Int64ListOptionPredicate func(t Int64ListOption) bool
type UintListOptionPredicate func(t UintListOption) bool
type Uint8ListOptionPredicate func(t Uint8ListOption) bool
type Uint16ListOptionPredicate func(t Uint16ListOption) bool
type Uint32ListOptionPredicate func(t Uint32ListOption) bool
type Uint64ListOptionPredicate func(t Uint64ListOption) bool
type UintptrListOptionPredicate func(t UintptrListOption) bool
type ByteListOptionPredicate func(t ByteListOption) bool
type RuneListOptionPredicate func(t RuneListOption) bool
type Float32ListOptionPredicate func(t Float32ListOption) bool
type Float64ListOptionPredicate func(t Float64ListOption) bool
type Complex64ListOptionPredicate func(t Complex64ListOption) bool
type Complex128ListOptionPredicate func(t Complex128ListOption) bool
type AnyListOptionPredicate func(t AnyListOption) bool
type BoolOptionOptionPredicate func(t BoolOptionOption) bool
type StringOptionOptionPredicate func(t StringOptionOption) bool
type IntOptionOptionPredicate func(t IntOptionOption) bool
type Int8OptionOptionPredicate func(t Int8OptionOption) bool
type Int16OptionOptionPredicate func(t Int16OptionOption) bool
type Int32OptionOptionPredicate func(t Int32OptionOption) bool
type Int64OptionOptionPredicate func(t Int64OptionOption) bool
type UintOptionOptionPredicate func(t UintOptionOption) bool
type Uint8OptionOptionPredicate func(t Uint8OptionOption) bool
type Uint16OptionOptionPredicate func(t Uint16OptionOption) bool
type Uint32OptionOptionPredicate func(t Uint32OptionOption) bool
type Uint64OptionOptionPredicate func(t Uint64OptionOption) bool
type UintptrOptionOptionPredicate func(t UintptrOptionOption) bool
type ByteOptionOptionPredicate func(t ByteOptionOption) bool
type RuneOptionOptionPredicate func(t RuneOptionOption) bool
type Float32OptionOptionPredicate func(t Float32OptionOption) bool
type Float64OptionOptionPredicate func(t Float64OptionOption) bool
type Complex64OptionOptionPredicate func(t Complex64OptionOption) bool
type Complex128OptionOptionPredicate func(t Complex128OptionOption) bool
type AnyOptionOptionPredicate func(t AnyOptionOption) bool
type BoolArrOptionOptionPredicate func(t BoolArrOptionOption) bool
type StringArrOptionOptionPredicate func(t StringArrOptionOption) bool
type IntArrOptionOptionPredicate func(t IntArrOptionOption) bool
type Int8ArrOptionOptionPredicate func(t Int8ArrOptionOption) bool
type Int16ArrOptionOptionPredicate func(t Int16ArrOptionOption) bool
type Int32ArrOptionOptionPredicate func(t Int32ArrOptionOption) bool
type Int64ArrOptionOptionPredicate func(t Int64ArrOptionOption) bool
type UintArrOptionOptionPredicate func(t UintArrOptionOption) bool
type Uint8ArrOptionOptionPredicate func(t Uint8ArrOptionOption) bool
type Uint16ArrOptionOptionPredicate func(t Uint16ArrOptionOption) bool
type Uint32ArrOptionOptionPredicate func(t Uint32ArrOptionOption) bool
type Uint64ArrOptionOptionPredicate func(t Uint64ArrOptionOption) bool
type UintptrArrOptionOptionPredicate func(t UintptrArrOptionOption) bool
type ByteArrOptionOptionPredicate func(t ByteArrOptionOption) bool
type RuneArrOptionOptionPredicate func(t RuneArrOptionOption) bool
type Float32ArrOptionOptionPredicate func(t Float32ArrOptionOption) bool
type Float64ArrOptionOptionPredicate func(t Float64ArrOptionOption) bool
type Complex64ArrOptionOptionPredicate func(t Complex64ArrOptionOption) bool
type Complex128ArrOptionOptionPredicate func(t Complex128ArrOptionOption) bool
type AnyArrOptionOptionPredicate func(t AnyArrOptionOption) bool
type BoolArrArrOptionOptionPredicate func(t BoolArrArrOptionOption) bool
type StringArrArrOptionOptionPredicate func(t StringArrArrOptionOption) bool
type IntArrArrOptionOptionPredicate func(t IntArrArrOptionOption) bool
type Int8ArrArrOptionOptionPredicate func(t Int8ArrArrOptionOption) bool
type Int16ArrArrOptionOptionPredicate func(t Int16ArrArrOptionOption) bool
type Int32ArrArrOptionOptionPredicate func(t Int32ArrArrOptionOption) bool
type Int64ArrArrOptionOptionPredicate func(t Int64ArrArrOptionOption) bool
type UintArrArrOptionOptionPredicate func(t UintArrArrOptionOption) bool
type Uint8ArrArrOptionOptionPredicate func(t Uint8ArrArrOptionOption) bool
type Uint16ArrArrOptionOptionPredicate func(t Uint16ArrArrOptionOption) bool
type Uint32ArrArrOptionOptionPredicate func(t Uint32ArrArrOptionOption) bool
type Uint64ArrArrOptionOptionPredicate func(t Uint64ArrArrOptionOption) bool
type UintptrArrArrOptionOptionPredicate func(t UintptrArrArrOptionOption) bool
type ByteArrArrOptionOptionPredicate func(t ByteArrArrOptionOption) bool
type RuneArrArrOptionOptionPredicate func(t RuneArrArrOptionOption) bool
type Float32ArrArrOptionOptionPredicate func(t Float32ArrArrOptionOption) bool
type Float64ArrArrOptionOptionPredicate func(t Float64ArrArrOptionOption) bool
type Complex64ArrArrOptionOptionPredicate func(t Complex64ArrArrOptionOption) bool
type Complex128ArrArrOptionOptionPredicate func(t Complex128ArrArrOptionOption) bool
type AnyArrArrOptionOptionPredicate func(t AnyArrArrOptionOption) bool
type BoolOptionArrOptionOptionPredicate func(t BoolOptionArrOptionOption) bool
type StringOptionArrOptionOptionPredicate func(t StringOptionArrOptionOption) bool
type IntOptionArrOptionOptionPredicate func(t IntOptionArrOptionOption) bool
type Int8OptionArrOptionOptionPredicate func(t Int8OptionArrOptionOption) bool
type Int16OptionArrOptionOptionPredicate func(t Int16OptionArrOptionOption) bool
type Int32OptionArrOptionOptionPredicate func(t Int32OptionArrOptionOption) bool
type Int64OptionArrOptionOptionPredicate func(t Int64OptionArrOptionOption) bool
type UintOptionArrOptionOptionPredicate func(t UintOptionArrOptionOption) bool
type Uint8OptionArrOptionOptionPredicate func(t Uint8OptionArrOptionOption) bool
type Uint16OptionArrOptionOptionPredicate func(t Uint16OptionArrOptionOption) bool
type Uint32OptionArrOptionOptionPredicate func(t Uint32OptionArrOptionOption) bool
type Uint64OptionArrOptionOptionPredicate func(t Uint64OptionArrOptionOption) bool
type UintptrOptionArrOptionOptionPredicate func(t UintptrOptionArrOptionOption) bool
type ByteOptionArrOptionOptionPredicate func(t ByteOptionArrOptionOption) bool
type RuneOptionArrOptionOptionPredicate func(t RuneOptionArrOptionOption) bool
type Float32OptionArrOptionOptionPredicate func(t Float32OptionArrOptionOption) bool
type Float64OptionArrOptionOptionPredicate func(t Float64OptionArrOptionOption) bool
type Complex64OptionArrOptionOptionPredicate func(t Complex64OptionArrOptionOption) bool
type Complex128OptionArrOptionOptionPredicate func(t Complex128OptionArrOptionOption) bool
type AnyOptionArrOptionOptionPredicate func(t AnyOptionArrOptionOption) bool
type BoolListOptionOptionPredicate func(t BoolListOptionOption) bool
type StringListOptionOptionPredicate func(t StringListOptionOption) bool
type IntListOptionOptionPredicate func(t IntListOptionOption) bool
type Int8ListOptionOptionPredicate func(t Int8ListOptionOption) bool
type Int16ListOptionOptionPredicate func(t Int16ListOptionOption) bool
type Int32ListOptionOptionPredicate func(t Int32ListOptionOption) bool
type Int64ListOptionOptionPredicate func(t Int64ListOptionOption) bool
type UintListOptionOptionPredicate func(t UintListOptionOption) bool
type Uint8ListOptionOptionPredicate func(t Uint8ListOptionOption) bool
type Uint16ListOptionOptionPredicate func(t Uint16ListOptionOption) bool
type Uint32ListOptionOptionPredicate func(t Uint32ListOptionOption) bool
type Uint64ListOptionOptionPredicate func(t Uint64ListOptionOption) bool
type UintptrListOptionOptionPredicate func(t UintptrListOptionOption) bool
type ByteListOptionOptionPredicate func(t ByteListOptionOption) bool
type RuneListOptionOptionPredicate func(t RuneListOptionOption) bool
type Float32ListOptionOptionPredicate func(t Float32ListOptionOption) bool
type Float64ListOptionOptionPredicate func(t Float64ListOptionOption) bool
type Complex64ListOptionOptionPredicate func(t Complex64ListOptionOption) bool
type Complex128ListOptionOptionPredicate func(t Complex128ListOptionOption) bool
type AnyListOptionOptionPredicate func(t AnyListOptionOption) bool
type BoolListPredicate func(t BoolList) bool
type StringListPredicate func(t StringList) bool
type IntListPredicate func(t IntList) bool
type Int8ListPredicate func(t Int8List) bool
type Int16ListPredicate func(t Int16List) bool
type Int32ListPredicate func(t Int32List) bool
type Int64ListPredicate func(t Int64List) bool
type UintListPredicate func(t UintList) bool
type Uint8ListPredicate func(t Uint8List) bool
type Uint16ListPredicate func(t Uint16List) bool
type Uint32ListPredicate func(t Uint32List) bool
type Uint64ListPredicate func(t Uint64List) bool
type UintptrListPredicate func(t UintptrList) bool
type ByteListPredicate func(t ByteList) bool
type RuneListPredicate func(t RuneList) bool
type Float32ListPredicate func(t Float32List) bool
type Float64ListPredicate func(t Float64List) bool
type Complex64ListPredicate func(t Complex64List) bool
type Complex128ListPredicate func(t Complex128List) bool
type AnyListPredicate func(t AnyList) bool
type BoolOptionListPredicate func(t BoolOptionList) bool
type StringOptionListPredicate func(t StringOptionList) bool
type IntOptionListPredicate func(t IntOptionList) bool
type Int8OptionListPredicate func(t Int8OptionList) bool
type Int16OptionListPredicate func(t Int16OptionList) bool
type Int32OptionListPredicate func(t Int32OptionList) bool
type Int64OptionListPredicate func(t Int64OptionList) bool
type UintOptionListPredicate func(t UintOptionList) bool
type Uint8OptionListPredicate func(t Uint8OptionList) bool
type Uint16OptionListPredicate func(t Uint16OptionList) bool
type Uint32OptionListPredicate func(t Uint32OptionList) bool
type Uint64OptionListPredicate func(t Uint64OptionList) bool
type UintptrOptionListPredicate func(t UintptrOptionList) bool
type ByteOptionListPredicate func(t ByteOptionList) bool
type RuneOptionListPredicate func(t RuneOptionList) bool
type Float32OptionListPredicate func(t Float32OptionList) bool
type Float64OptionListPredicate func(t Float64OptionList) bool
type Complex64OptionListPredicate func(t Complex64OptionList) bool
type Complex128OptionListPredicate func(t Complex128OptionList) bool
type AnyOptionListPredicate func(t AnyOptionList) bool
type BoolArrOptionListPredicate func(t BoolArrOptionList) bool
type StringArrOptionListPredicate func(t StringArrOptionList) bool
type IntArrOptionListPredicate func(t IntArrOptionList) bool
type Int8ArrOptionListPredicate func(t Int8ArrOptionList) bool
type Int16ArrOptionListPredicate func(t Int16ArrOptionList) bool
type Int32ArrOptionListPredicate func(t Int32ArrOptionList) bool
type Int64ArrOptionListPredicate func(t Int64ArrOptionList) bool
type UintArrOptionListPredicate func(t UintArrOptionList) bool
type Uint8ArrOptionListPredicate func(t Uint8ArrOptionList) bool
type Uint16ArrOptionListPredicate func(t Uint16ArrOptionList) bool
type Uint32ArrOptionListPredicate func(t Uint32ArrOptionList) bool
type Uint64ArrOptionListPredicate func(t Uint64ArrOptionList) bool
type UintptrArrOptionListPredicate func(t UintptrArrOptionList) bool
type ByteArrOptionListPredicate func(t ByteArrOptionList) bool
type RuneArrOptionListPredicate func(t RuneArrOptionList) bool
type Float32ArrOptionListPredicate func(t Float32ArrOptionList) bool
type Float64ArrOptionListPredicate func(t Float64ArrOptionList) bool
type Complex64ArrOptionListPredicate func(t Complex64ArrOptionList) bool
type Complex128ArrOptionListPredicate func(t Complex128ArrOptionList) bool
type AnyArrOptionListPredicate func(t AnyArrOptionList) bool
type BoolArrArrOptionListPredicate func(t BoolArrArrOptionList) bool
type StringArrArrOptionListPredicate func(t StringArrArrOptionList) bool
type IntArrArrOptionListPredicate func(t IntArrArrOptionList) bool
type Int8ArrArrOptionListPredicate func(t Int8ArrArrOptionList) bool
type Int16ArrArrOptionListPredicate func(t Int16ArrArrOptionList) bool
type Int32ArrArrOptionListPredicate func(t Int32ArrArrOptionList) bool
type Int64ArrArrOptionListPredicate func(t Int64ArrArrOptionList) bool
type UintArrArrOptionListPredicate func(t UintArrArrOptionList) bool
type Uint8ArrArrOptionListPredicate func(t Uint8ArrArrOptionList) bool
type Uint16ArrArrOptionListPredicate func(t Uint16ArrArrOptionList) bool
type Uint32ArrArrOptionListPredicate func(t Uint32ArrArrOptionList) bool
type Uint64ArrArrOptionListPredicate func(t Uint64ArrArrOptionList) bool
type UintptrArrArrOptionListPredicate func(t UintptrArrArrOptionList) bool
type ByteArrArrOptionListPredicate func(t ByteArrArrOptionList) bool
type RuneArrArrOptionListPredicate func(t RuneArrArrOptionList) bool
type Float32ArrArrOptionListPredicate func(t Float32ArrArrOptionList) bool
type Float64ArrArrOptionListPredicate func(t Float64ArrArrOptionList) bool
type Complex64ArrArrOptionListPredicate func(t Complex64ArrArrOptionList) bool
type Complex128ArrArrOptionListPredicate func(t Complex128ArrArrOptionList) bool
type AnyArrArrOptionListPredicate func(t AnyArrArrOptionList) bool
type BoolOptionArrOptionListPredicate func(t BoolOptionArrOptionList) bool
type StringOptionArrOptionListPredicate func(t StringOptionArrOptionList) bool
type IntOptionArrOptionListPredicate func(t IntOptionArrOptionList) bool
type Int8OptionArrOptionListPredicate func(t Int8OptionArrOptionList) bool
type Int16OptionArrOptionListPredicate func(t Int16OptionArrOptionList) bool
type Int32OptionArrOptionListPredicate func(t Int32OptionArrOptionList) bool
type Int64OptionArrOptionListPredicate func(t Int64OptionArrOptionList) bool
type UintOptionArrOptionListPredicate func(t UintOptionArrOptionList) bool
type Uint8OptionArrOptionListPredicate func(t Uint8OptionArrOptionList) bool
type Uint16OptionArrOptionListPredicate func(t Uint16OptionArrOptionList) bool
type Uint32OptionArrOptionListPredicate func(t Uint32OptionArrOptionList) bool
type Uint64OptionArrOptionListPredicate func(t Uint64OptionArrOptionList) bool
type UintptrOptionArrOptionListPredicate func(t UintptrOptionArrOptionList) bool
type ByteOptionArrOptionListPredicate func(t ByteOptionArrOptionList) bool
type RuneOptionArrOptionListPredicate func(t RuneOptionArrOptionList) bool
type Float32OptionArrOptionListPredicate func(t Float32OptionArrOptionList) bool
type Float64OptionArrOptionListPredicate func(t Float64OptionArrOptionList) bool
type Complex64OptionArrOptionListPredicate func(t Complex64OptionArrOptionList) bool
type Complex128OptionArrOptionListPredicate func(t Complex128OptionArrOptionList) bool
type AnyOptionArrOptionListPredicate func(t AnyOptionArrOptionList) bool
type BoolListOptionListPredicate func(t BoolListOptionList) bool
type StringListOptionListPredicate func(t StringListOptionList) bool
type IntListOptionListPredicate func(t IntListOptionList) bool
type Int8ListOptionListPredicate func(t Int8ListOptionList) bool
type Int16ListOptionListPredicate func(t Int16ListOptionList) bool
type Int32ListOptionListPredicate func(t Int32ListOptionList) bool
type Int64ListOptionListPredicate func(t Int64ListOptionList) bool
type UintListOptionListPredicate func(t UintListOptionList) bool
type Uint8ListOptionListPredicate func(t Uint8ListOptionList) bool
type Uint16ListOptionListPredicate func(t Uint16ListOptionList) bool
type Uint32ListOptionListPredicate func(t Uint32ListOptionList) bool
type Uint64ListOptionListPredicate func(t Uint64ListOptionList) bool
type UintptrListOptionListPredicate func(t UintptrListOptionList) bool
type ByteListOptionListPredicate func(t ByteListOptionList) bool
type RuneListOptionListPredicate func(t RuneListOptionList) bool
type Float32ListOptionListPredicate func(t Float32ListOptionList) bool
type Float64ListOptionListPredicate func(t Float64ListOptionList) bool
type Complex64ListOptionListPredicate func(t Complex64ListOptionList) bool
type Complex128ListOptionListPredicate func(t Complex128ListOptionList) bool
type AnyListOptionListPredicate func(t AnyListOptionList) bool
type BoolOptionOptionListPredicate func(t BoolOptionOptionList) bool
type StringOptionOptionListPredicate func(t StringOptionOptionList) bool
type IntOptionOptionListPredicate func(t IntOptionOptionList) bool
type Int8OptionOptionListPredicate func(t Int8OptionOptionList) bool
type Int16OptionOptionListPredicate func(t Int16OptionOptionList) bool
type Int32OptionOptionListPredicate func(t Int32OptionOptionList) bool
type Int64OptionOptionListPredicate func(t Int64OptionOptionList) bool
type UintOptionOptionListPredicate func(t UintOptionOptionList) bool
type Uint8OptionOptionListPredicate func(t Uint8OptionOptionList) bool
type Uint16OptionOptionListPredicate func(t Uint16OptionOptionList) bool
type Uint32OptionOptionListPredicate func(t Uint32OptionOptionList) bool
type Uint64OptionOptionListPredicate func(t Uint64OptionOptionList) bool
type UintptrOptionOptionListPredicate func(t UintptrOptionOptionList) bool
type ByteOptionOptionListPredicate func(t ByteOptionOptionList) bool
type RuneOptionOptionListPredicate func(t RuneOptionOptionList) bool
type Float32OptionOptionListPredicate func(t Float32OptionOptionList) bool
type Float64OptionOptionListPredicate func(t Float64OptionOptionList) bool
type Complex64OptionOptionListPredicate func(t Complex64OptionOptionList) bool
type Complex128OptionOptionListPredicate func(t Complex128OptionOptionList) bool
type AnyOptionOptionListPredicate func(t AnyOptionOptionList) bool
type BoolArrOptionOptionListPredicate func(t BoolArrOptionOptionList) bool
type StringArrOptionOptionListPredicate func(t StringArrOptionOptionList) bool
type IntArrOptionOptionListPredicate func(t IntArrOptionOptionList) bool
type Int8ArrOptionOptionListPredicate func(t Int8ArrOptionOptionList) bool
type Int16ArrOptionOptionListPredicate func(t Int16ArrOptionOptionList) bool
type Int32ArrOptionOptionListPredicate func(t Int32ArrOptionOptionList) bool
type Int64ArrOptionOptionListPredicate func(t Int64ArrOptionOptionList) bool
type UintArrOptionOptionListPredicate func(t UintArrOptionOptionList) bool
type Uint8ArrOptionOptionListPredicate func(t Uint8ArrOptionOptionList) bool
type Uint16ArrOptionOptionListPredicate func(t Uint16ArrOptionOptionList) bool
type Uint32ArrOptionOptionListPredicate func(t Uint32ArrOptionOptionList) bool
type Uint64ArrOptionOptionListPredicate func(t Uint64ArrOptionOptionList) bool
type UintptrArrOptionOptionListPredicate func(t UintptrArrOptionOptionList) bool
type ByteArrOptionOptionListPredicate func(t ByteArrOptionOptionList) bool
type RuneArrOptionOptionListPredicate func(t RuneArrOptionOptionList) bool
type Float32ArrOptionOptionListPredicate func(t Float32ArrOptionOptionList) bool
type Float64ArrOptionOptionListPredicate func(t Float64ArrOptionOptionList) bool
type Complex64ArrOptionOptionListPredicate func(t Complex64ArrOptionOptionList) bool
type Complex128ArrOptionOptionListPredicate func(t Complex128ArrOptionOptionList) bool
type AnyArrOptionOptionListPredicate func(t AnyArrOptionOptionList) bool
type BoolArrArrOptionOptionListPredicate func(t BoolArrArrOptionOptionList) bool
type StringArrArrOptionOptionListPredicate func(t StringArrArrOptionOptionList) bool
type IntArrArrOptionOptionListPredicate func(t IntArrArrOptionOptionList) bool
type Int8ArrArrOptionOptionListPredicate func(t Int8ArrArrOptionOptionList) bool
type Int16ArrArrOptionOptionListPredicate func(t Int16ArrArrOptionOptionList) bool
type Int32ArrArrOptionOptionListPredicate func(t Int32ArrArrOptionOptionList) bool
type Int64ArrArrOptionOptionListPredicate func(t Int64ArrArrOptionOptionList) bool
type UintArrArrOptionOptionListPredicate func(t UintArrArrOptionOptionList) bool
type Uint8ArrArrOptionOptionListPredicate func(t Uint8ArrArrOptionOptionList) bool
type Uint16ArrArrOptionOptionListPredicate func(t Uint16ArrArrOptionOptionList) bool
type Uint32ArrArrOptionOptionListPredicate func(t Uint32ArrArrOptionOptionList) bool
type Uint64ArrArrOptionOptionListPredicate func(t Uint64ArrArrOptionOptionList) bool
type UintptrArrArrOptionOptionListPredicate func(t UintptrArrArrOptionOptionList) bool
type ByteArrArrOptionOptionListPredicate func(t ByteArrArrOptionOptionList) bool
type RuneArrArrOptionOptionListPredicate func(t RuneArrArrOptionOptionList) bool
type Float32ArrArrOptionOptionListPredicate func(t Float32ArrArrOptionOptionList) bool
type Float64ArrArrOptionOptionListPredicate func(t Float64ArrArrOptionOptionList) bool
type Complex64ArrArrOptionOptionListPredicate func(t Complex64ArrArrOptionOptionList) bool
type Complex128ArrArrOptionOptionListPredicate func(t Complex128ArrArrOptionOptionList) bool
type AnyArrArrOptionOptionListPredicate func(t AnyArrArrOptionOptionList) bool
type BoolOptionArrOptionOptionListPredicate func(t BoolOptionArrOptionOptionList) bool
type StringOptionArrOptionOptionListPredicate func(t StringOptionArrOptionOptionList) bool
type IntOptionArrOptionOptionListPredicate func(t IntOptionArrOptionOptionList) bool
type Int8OptionArrOptionOptionListPredicate func(t Int8OptionArrOptionOptionList) bool
type Int16OptionArrOptionOptionListPredicate func(t Int16OptionArrOptionOptionList) bool
type Int32OptionArrOptionOptionListPredicate func(t Int32OptionArrOptionOptionList) bool
type Int64OptionArrOptionOptionListPredicate func(t Int64OptionArrOptionOptionList) bool
type UintOptionArrOptionOptionListPredicate func(t UintOptionArrOptionOptionList) bool
type Uint8OptionArrOptionOptionListPredicate func(t Uint8OptionArrOptionOptionList) bool
type Uint16OptionArrOptionOptionListPredicate func(t Uint16OptionArrOptionOptionList) bool
type Uint32OptionArrOptionOptionListPredicate func(t Uint32OptionArrOptionOptionList) bool
type Uint64OptionArrOptionOptionListPredicate func(t Uint64OptionArrOptionOptionList) bool
type UintptrOptionArrOptionOptionListPredicate func(t UintptrOptionArrOptionOptionList) bool
type ByteOptionArrOptionOptionListPredicate func(t ByteOptionArrOptionOptionList) bool
type RuneOptionArrOptionOptionListPredicate func(t RuneOptionArrOptionOptionList) bool
type Float32OptionArrOptionOptionListPredicate func(t Float32OptionArrOptionOptionList) bool
type Float64OptionArrOptionOptionListPredicate func(t Float64OptionArrOptionOptionList) bool
type Complex64OptionArrOptionOptionListPredicate func(t Complex64OptionArrOptionOptionList) bool
type Complex128OptionArrOptionOptionListPredicate func(t Complex128OptionArrOptionOptionList) bool
type AnyOptionArrOptionOptionListPredicate func(t AnyOptionArrOptionOptionList) bool
type BoolListOptionOptionListPredicate func(t BoolListOptionOptionList) bool
type StringListOptionOptionListPredicate func(t StringListOptionOptionList) bool
type IntListOptionOptionListPredicate func(t IntListOptionOptionList) bool
type Int8ListOptionOptionListPredicate func(t Int8ListOptionOptionList) bool
type Int16ListOptionOptionListPredicate func(t Int16ListOptionOptionList) bool
type Int32ListOptionOptionListPredicate func(t Int32ListOptionOptionList) bool
type Int64ListOptionOptionListPredicate func(t Int64ListOptionOptionList) bool
type UintListOptionOptionListPredicate func(t UintListOptionOptionList) bool
type Uint8ListOptionOptionListPredicate func(t Uint8ListOptionOptionList) bool
type Uint16ListOptionOptionListPredicate func(t Uint16ListOptionOptionList) bool
type Uint32ListOptionOptionListPredicate func(t Uint32ListOptionOptionList) bool
type Uint64ListOptionOptionListPredicate func(t Uint64ListOptionOptionList) bool
type UintptrListOptionOptionListPredicate func(t UintptrListOptionOptionList) bool
type ByteListOptionOptionListPredicate func(t ByteListOptionOptionList) bool
type RuneListOptionOptionListPredicate func(t RuneListOptionOptionList) bool
type Float32ListOptionOptionListPredicate func(t Float32ListOptionOptionList) bool
type Float64ListOptionOptionListPredicate func(t Float64ListOptionOptionList) bool
type Complex64ListOptionOptionListPredicate func(t Complex64ListOptionOptionList) bool
type Complex128ListOptionOptionListPredicate func(t Complex128ListOptionOptionList) bool
type AnyListOptionOptionListPredicate func(t AnyListOptionOptionList) bool
type BoolArrListPredicate func(t BoolArrList) bool
type StringArrListPredicate func(t StringArrList) bool
type IntArrListPredicate func(t IntArrList) bool
type Int8ArrListPredicate func(t Int8ArrList) bool
type Int16ArrListPredicate func(t Int16ArrList) bool
type Int32ArrListPredicate func(t Int32ArrList) bool
type Int64ArrListPredicate func(t Int64ArrList) bool
type UintArrListPredicate func(t UintArrList) bool
type Uint8ArrListPredicate func(t Uint8ArrList) bool
type Uint16ArrListPredicate func(t Uint16ArrList) bool
type Uint32ArrListPredicate func(t Uint32ArrList) bool
type Uint64ArrListPredicate func(t Uint64ArrList) bool
type UintptrArrListPredicate func(t UintptrArrList) bool
type ByteArrListPredicate func(t ByteArrList) bool
type RuneArrListPredicate func(t RuneArrList) bool
type Float32ArrListPredicate func(t Float32ArrList) bool
type Float64ArrListPredicate func(t Float64ArrList) bool
type Complex64ArrListPredicate func(t Complex64ArrList) bool
type Complex128ArrListPredicate func(t Complex128ArrList) bool
type AnyArrListPredicate func(t AnyArrList) bool
type BoolArrArrListPredicate func(t BoolArrArrList) bool
type StringArrArrListPredicate func(t StringArrArrList) bool
type IntArrArrListPredicate func(t IntArrArrList) bool
type Int8ArrArrListPredicate func(t Int8ArrArrList) bool
type Int16ArrArrListPredicate func(t Int16ArrArrList) bool
type Int32ArrArrListPredicate func(t Int32ArrArrList) bool
type Int64ArrArrListPredicate func(t Int64ArrArrList) bool
type UintArrArrListPredicate func(t UintArrArrList) bool
type Uint8ArrArrListPredicate func(t Uint8ArrArrList) bool
type Uint16ArrArrListPredicate func(t Uint16ArrArrList) bool
type Uint32ArrArrListPredicate func(t Uint32ArrArrList) bool
type Uint64ArrArrListPredicate func(t Uint64ArrArrList) bool
type UintptrArrArrListPredicate func(t UintptrArrArrList) bool
type ByteArrArrListPredicate func(t ByteArrArrList) bool
type RuneArrArrListPredicate func(t RuneArrArrList) bool
type Float32ArrArrListPredicate func(t Float32ArrArrList) bool
type Float64ArrArrListPredicate func(t Float64ArrArrList) bool
type Complex64ArrArrListPredicate func(t Complex64ArrArrList) bool
type Complex128ArrArrListPredicate func(t Complex128ArrArrList) bool
type AnyArrArrListPredicate func(t AnyArrArrList) bool
type BoolOptionArrListPredicate func(t BoolOptionArrList) bool
type StringOptionArrListPredicate func(t StringOptionArrList) bool
type IntOptionArrListPredicate func(t IntOptionArrList) bool
type Int8OptionArrListPredicate func(t Int8OptionArrList) bool
type Int16OptionArrListPredicate func(t Int16OptionArrList) bool
type Int32OptionArrListPredicate func(t Int32OptionArrList) bool
type Int64OptionArrListPredicate func(t Int64OptionArrList) bool
type UintOptionArrListPredicate func(t UintOptionArrList) bool
type Uint8OptionArrListPredicate func(t Uint8OptionArrList) bool
type Uint16OptionArrListPredicate func(t Uint16OptionArrList) bool
type Uint32OptionArrListPredicate func(t Uint32OptionArrList) bool
type Uint64OptionArrListPredicate func(t Uint64OptionArrList) bool
type UintptrOptionArrListPredicate func(t UintptrOptionArrList) bool
type ByteOptionArrListPredicate func(t ByteOptionArrList) bool
type RuneOptionArrListPredicate func(t RuneOptionArrList) bool
type Float32OptionArrListPredicate func(t Float32OptionArrList) bool
type Float64OptionArrListPredicate func(t Float64OptionArrList) bool
type Complex64OptionArrListPredicate func(t Complex64OptionArrList) bool
type Complex128OptionArrListPredicate func(t Complex128OptionArrList) bool
type AnyOptionArrListPredicate func(t AnyOptionArrList) bool
type BoolListListPredicate func(t BoolListList) bool
type StringListListPredicate func(t StringListList) bool
type IntListListPredicate func(t IntListList) bool
type Int8ListListPredicate func(t Int8ListList) bool
type Int16ListListPredicate func(t Int16ListList) bool
type Int32ListListPredicate func(t Int32ListList) bool
type Int64ListListPredicate func(t Int64ListList) bool
type UintListListPredicate func(t UintListList) bool
type Uint8ListListPredicate func(t Uint8ListList) bool
type Uint16ListListPredicate func(t Uint16ListList) bool
type Uint32ListListPredicate func(t Uint32ListList) bool
type Uint64ListListPredicate func(t Uint64ListList) bool
type UintptrListListPredicate func(t UintptrListList) bool
type ByteListListPredicate func(t ByteListList) bool
type RuneListListPredicate func(t RuneListList) bool
type Float32ListListPredicate func(t Float32ListList) bool
type Float64ListListPredicate func(t Float64ListList) bool
type Complex64ListListPredicate func(t Complex64ListList) bool
type Complex128ListListPredicate func(t Complex128ListList) bool
type AnyListListPredicate func(t AnyListList) bool