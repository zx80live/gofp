// bootstrap_option_equals.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (a BoolOption) Equals(b BoolOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return BoolEquals(Bool(*a.value), Bool(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a StringOption) Equals(b StringOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return StringEquals(String(*a.value), String(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a IntOption) Equals(b IntOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return IntEquals(Int(*a.value), Int(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Int64Option) Equals(b Int64Option) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Int64Equals(Int64(*a.value), Int64(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a ByteOption) Equals(b ByteOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return ByteEquals(Byte(*a.value), Byte(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a RuneOption) Equals(b RuneOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return RuneEquals(Rune(*a.value), Rune(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float32Option) Equals(b Float32Option) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float32Equals(Float32(*a.value), Float32(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float64Option) Equals(b Float64Option) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float64Equals(Float64(*a.value), Float64(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a AnyOption) Equals(b AnyOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return AnyEquals(Any(*a.value), Any(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Tuple2Option) Equals(b Tuple2Option) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Tuple2Equals(Tuple2(*a.value), Tuple2(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a BoolOptionOption) Equals(b BoolOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return BoolOptionEquals(BoolOption(*a.value), BoolOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a StringOptionOption) Equals(b StringOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return StringOptionEquals(StringOption(*a.value), StringOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a IntOptionOption) Equals(b IntOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return IntOptionEquals(IntOption(*a.value), IntOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Int64OptionOption) Equals(b Int64OptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Int64OptionEquals(Int64Option(*a.value), Int64Option(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a ByteOptionOption) Equals(b ByteOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return ByteOptionEquals(ByteOption(*a.value), ByteOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a RuneOptionOption) Equals(b RuneOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return RuneOptionEquals(RuneOption(*a.value), RuneOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float32OptionOption) Equals(b Float32OptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float32OptionEquals(Float32Option(*a.value), Float32Option(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float64OptionOption) Equals(b Float64OptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float64OptionEquals(Float64Option(*a.value), Float64Option(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a AnyOptionOption) Equals(b AnyOptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return AnyOptionEquals(AnyOption(*a.value), AnyOption(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Tuple2OptionOption) Equals(b Tuple2OptionOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Tuple2OptionEquals(Tuple2Option(*a.value), Tuple2Option(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a BoolArrayOption) Equals(b BoolArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return BoolArrayEquals(BoolArray(*a.value), BoolArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a StringArrayOption) Equals(b StringArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return StringArrayEquals(StringArray(*a.value), StringArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a IntArrayOption) Equals(b IntArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return IntArrayEquals(IntArray(*a.value), IntArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Int64ArrayOption) Equals(b Int64ArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Int64ArrayEquals(Int64Array(*a.value), Int64Array(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a ByteArrayOption) Equals(b ByteArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return ByteArrayEquals(ByteArray(*a.value), ByteArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a RuneArrayOption) Equals(b RuneArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return RuneArrayEquals(RuneArray(*a.value), RuneArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float32ArrayOption) Equals(b Float32ArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float32ArrayEquals(Float32Array(*a.value), Float32Array(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float64ArrayOption) Equals(b Float64ArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float64ArrayEquals(Float64Array(*a.value), Float64Array(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a AnyArrayOption) Equals(b AnyArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return AnyArrayEquals(AnyArray(*a.value), AnyArray(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Tuple2ArrayOption) Equals(b Tuple2ArrayOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Tuple2ArrayEquals(Tuple2Array(*a.value), Tuple2Array(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a BoolListOption) Equals(b BoolListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return BoolListEquals(BoolList(*a.value), BoolList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a StringListOption) Equals(b StringListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return StringListEquals(StringList(*a.value), StringList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a IntListOption) Equals(b IntListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return IntListEquals(IntList(*a.value), IntList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Int64ListOption) Equals(b Int64ListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Int64ListEquals(Int64List(*a.value), Int64List(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a ByteListOption) Equals(b ByteListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return ByteListEquals(ByteList(*a.value), ByteList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a RuneListOption) Equals(b RuneListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return RuneListEquals(RuneList(*a.value), RuneList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float32ListOption) Equals(b Float32ListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float32ListEquals(Float32List(*a.value), Float32List(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Float64ListOption) Equals(b Float64ListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Float64ListEquals(Float64List(*a.value), Float64List(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a AnyListOption) Equals(b AnyListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return AnyListEquals(AnyList(*a.value), AnyList(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }
func (a Tuple2ListOption) Equals(b Tuple2ListOption) bool {
  if a.IsDefined() {
    if b.IsDefined() {
      return Tuple2ListEquals(Tuple2List(*a.value), Tuple2List(*b.value))
    } else { return false}
  } else if b.IsDefined() { return false } else { return true } }