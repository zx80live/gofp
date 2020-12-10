// bootstrap_option_foreach.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (o BoolOption) Foreach(f func(bool)) { if o.IsDefined() { f(*o.value) } }
func (o StringOption) Foreach(f func(string)) { if o.IsDefined() { f(*o.value) } }
func (o IntOption) Foreach(f func(int)) { if o.IsDefined() { f(*o.value) } }
func (o Int64Option) Foreach(f func(int64)) { if o.IsDefined() { f(*o.value) } }
func (o ByteOption) Foreach(f func(byte)) { if o.IsDefined() { f(*o.value) } }
func (o RuneOption) Foreach(f func(rune)) { if o.IsDefined() { f(*o.value) } }
func (o Float32Option) Foreach(f func(float32)) { if o.IsDefined() { f(*o.value) } }
func (o Float64Option) Foreach(f func(float64)) { if o.IsDefined() { f(*o.value) } }
func (o AnyOption) Foreach(f func(Any)) { if o.IsDefined() { f(*o.value) } }
func (o Tuple2Option) Foreach(f func(Tuple2)) { if o.IsDefined() { f(*o.value) } }
func (o BoolOptionOption) Foreach(f func(BoolOption)) { if o.IsDefined() { f(*o.value) } }
func (o StringOptionOption) Foreach(f func(StringOption)) { if o.IsDefined() { f(*o.value) } }
func (o IntOptionOption) Foreach(f func(IntOption)) { if o.IsDefined() { f(*o.value) } }
func (o Int64OptionOption) Foreach(f func(Int64Option)) { if o.IsDefined() { f(*o.value) } }
func (o ByteOptionOption) Foreach(f func(ByteOption)) { if o.IsDefined() { f(*o.value) } }
func (o RuneOptionOption) Foreach(f func(RuneOption)) { if o.IsDefined() { f(*o.value) } }
func (o Float32OptionOption) Foreach(f func(Float32Option)) { if o.IsDefined() { f(*o.value) } }
func (o Float64OptionOption) Foreach(f func(Float64Option)) { if o.IsDefined() { f(*o.value) } }
func (o AnyOptionOption) Foreach(f func(AnyOption)) { if o.IsDefined() { f(*o.value) } }
func (o Tuple2OptionOption) Foreach(f func(Tuple2Option)) { if o.IsDefined() { f(*o.value) } }
func (o BoolArrayOption) Foreach(f func([]bool)) { if o.IsDefined() { f(*o.value) } }
func (o StringArrayOption) Foreach(f func([]string)) { if o.IsDefined() { f(*o.value) } }
func (o IntArrayOption) Foreach(f func([]int)) { if o.IsDefined() { f(*o.value) } }
func (o Int64ArrayOption) Foreach(f func([]int64)) { if o.IsDefined() { f(*o.value) } }
func (o ByteArrayOption) Foreach(f func([]byte)) { if o.IsDefined() { f(*o.value) } }
func (o RuneArrayOption) Foreach(f func([]rune)) { if o.IsDefined() { f(*o.value) } }
func (o Float32ArrayOption) Foreach(f func([]float32)) { if o.IsDefined() { f(*o.value) } }
func (o Float64ArrayOption) Foreach(f func([]float64)) { if o.IsDefined() { f(*o.value) } }
func (o AnyArrayOption) Foreach(f func([]Any)) { if o.IsDefined() { f(*o.value) } }
func (o Tuple2ArrayOption) Foreach(f func([]Tuple2)) { if o.IsDefined() { f(*o.value) } }
func (o BoolListOption) Foreach(f func(BoolList)) { if o.IsDefined() { f(*o.value) } }
func (o StringListOption) Foreach(f func(StringList)) { if o.IsDefined() { f(*o.value) } }
func (o IntListOption) Foreach(f func(IntList)) { if o.IsDefined() { f(*o.value) } }
func (o Int64ListOption) Foreach(f func(Int64List)) { if o.IsDefined() { f(*o.value) } }
func (o ByteListOption) Foreach(f func(ByteList)) { if o.IsDefined() { f(*o.value) } }
func (o RuneListOption) Foreach(f func(RuneList)) { if o.IsDefined() { f(*o.value) } }
func (o Float32ListOption) Foreach(f func(Float32List)) { if o.IsDefined() { f(*o.value) } }
func (o Float64ListOption) Foreach(f func(Float64List)) { if o.IsDefined() { f(*o.value) } }
func (o AnyListOption) Foreach(f func(AnyList)) { if o.IsDefined() { f(*o.value) } }
func (o Tuple2ListOption) Foreach(f func(Tuple2List)) { if o.IsDefined() { f(*o.value) } }