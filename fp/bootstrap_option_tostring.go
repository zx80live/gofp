// bootstrap_option_tostring.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import "fmt"
func (o BoolOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", BoolToString(Bool(*o.value)))) } else { return "None" } }
func (o StringOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", StringToString(String(*o.value)))) } else { return "None" } }
func (o IntOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", IntToString(Int(*o.value)))) } else { return "None" } }
func (o Int64Option) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Int64ToString(Int64(*o.value)))) } else { return "None" } }
func (o ByteOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", ByteToString(Byte(*o.value)))) } else { return "None" } }
func (o RuneOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", RuneToString(Rune(*o.value)))) } else { return "None" } }
func (o Float32Option) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float32ToString(Float32(*o.value)))) } else { return "None" } }
func (o Float64Option) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float64ToString(Float64(*o.value)))) } else { return "None" } }
func (o AnyOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", AnyToString(Any(*o.value)))) } else { return "None" } }
func (o Tuple2Option) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Tuple2ToString(Tuple2(*o.value)))) } else { return "None" } }
func (o BoolOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", BoolOptionToString(BoolOption(*o.value)))) } else { return "None" } }
func (o StringOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", StringOptionToString(StringOption(*o.value)))) } else { return "None" } }
func (o IntOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", IntOptionToString(IntOption(*o.value)))) } else { return "None" } }
func (o Int64OptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Int64OptionToString(Int64Option(*o.value)))) } else { return "None" } }
func (o ByteOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", ByteOptionToString(ByteOption(*o.value)))) } else { return "None" } }
func (o RuneOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", RuneOptionToString(RuneOption(*o.value)))) } else { return "None" } }
func (o Float32OptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float32OptionToString(Float32Option(*o.value)))) } else { return "None" } }
func (o Float64OptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float64OptionToString(Float64Option(*o.value)))) } else { return "None" } }
func (o AnyOptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", AnyOptionToString(AnyOption(*o.value)))) } else { return "None" } }
func (o Tuple2OptionOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Tuple2OptionToString(Tuple2Option(*o.value)))) } else { return "None" } }
func (o BoolArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", BoolArrayToString(BoolArray(*o.value)))) } else { return "None" } }
func (o StringArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", StringArrayToString(StringArray(*o.value)))) } else { return "None" } }
func (o IntArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", IntArrayToString(IntArray(*o.value)))) } else { return "None" } }
func (o Int64ArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Int64ArrayToString(Int64Array(*o.value)))) } else { return "None" } }
func (o ByteArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", ByteArrayToString(ByteArray(*o.value)))) } else { return "None" } }
func (o RuneArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", RuneArrayToString(RuneArray(*o.value)))) } else { return "None" } }
func (o Float32ArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float32ArrayToString(Float32Array(*o.value)))) } else { return "None" } }
func (o Float64ArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float64ArrayToString(Float64Array(*o.value)))) } else { return "None" } }
func (o AnyArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", AnyArrayToString(AnyArray(*o.value)))) } else { return "None" } }
func (o Tuple2ArrayOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Tuple2ArrayToString(Tuple2Array(*o.value)))) } else { return "None" } }
func (o BoolListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", BoolListToString(BoolList(*o.value)))) } else { return "None" } }
func (o StringListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", StringListToString(StringList(*o.value)))) } else { return "None" } }
func (o IntListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", IntListToString(IntList(*o.value)))) } else { return "None" } }
func (o Int64ListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Int64ListToString(Int64List(*o.value)))) } else { return "None" } }
func (o ByteListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", ByteListToString(ByteList(*o.value)))) } else { return "None" } }
func (o RuneListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", RuneListToString(RuneList(*o.value)))) } else { return "None" } }
func (o Float32ListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float32ListToString(Float32List(*o.value)))) } else { return "None" } }
func (o Float64ListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Float64ListToString(Float64List(*o.value)))) } else { return "None" } }
func (o AnyListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", AnyListToString(AnyList(*o.value)))) } else { return "None" } }
func (o Tuple2ListOption) ToString() String { if o.IsDefined() { return String(fmt.Sprintf("Some(%v)", Tuple2ListToString(Tuple2List(*o.value)))) } else { return "None" } }