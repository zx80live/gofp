// bootstrap_list_headoption.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (l BoolList) HeadOption() BoolOption { if l.NonEmpty() { return MkBoolOption(*l.head) } else { return NoneBool } }
func (l StringList) HeadOption() StringOption { if l.NonEmpty() { return MkStringOption(*l.head) } else { return NoneString } }
func (l IntList) HeadOption() IntOption { if l.NonEmpty() { return MkIntOption(*l.head) } else { return NoneInt } }
func (l Int64List) HeadOption() Int64Option { if l.NonEmpty() { return MkInt64Option(*l.head) } else { return NoneInt64 } }
func (l ByteList) HeadOption() ByteOption { if l.NonEmpty() { return MkByteOption(*l.head) } else { return NoneByte } }
func (l RuneList) HeadOption() RuneOption { if l.NonEmpty() { return MkRuneOption(*l.head) } else { return NoneRune } }
func (l Float32List) HeadOption() Float32Option { if l.NonEmpty() { return MkFloat32Option(*l.head) } else { return NoneFloat32 } }
func (l Float64List) HeadOption() Float64Option { if l.NonEmpty() { return MkFloat64Option(*l.head) } else { return NoneFloat64 } }
func (l AnyList) HeadOption() AnyOption { if l.NonEmpty() { return MkAnyOption(*l.head) } else { return NoneAny } }
func (l Tuple2List) HeadOption() Tuple2Option { if l.NonEmpty() { return MkTuple2Option(*l.head) } else { return NoneTuple2 } }
func (l BoolArrayList) HeadOption() BoolArrayOption { if l.NonEmpty() { return MkBoolArrayOption(*l.head) } else { return NoneBoolArray } }
func (l StringArrayList) HeadOption() StringArrayOption { if l.NonEmpty() { return MkStringArrayOption(*l.head) } else { return NoneStringArray } }
func (l IntArrayList) HeadOption() IntArrayOption { if l.NonEmpty() { return MkIntArrayOption(*l.head) } else { return NoneIntArray } }
func (l Int64ArrayList) HeadOption() Int64ArrayOption { if l.NonEmpty() { return MkInt64ArrayOption(*l.head) } else { return NoneInt64Array } }
func (l ByteArrayList) HeadOption() ByteArrayOption { if l.NonEmpty() { return MkByteArrayOption(*l.head) } else { return NoneByteArray } }
func (l RuneArrayList) HeadOption() RuneArrayOption { if l.NonEmpty() { return MkRuneArrayOption(*l.head) } else { return NoneRuneArray } }
func (l Float32ArrayList) HeadOption() Float32ArrayOption { if l.NonEmpty() { return MkFloat32ArrayOption(*l.head) } else { return NoneFloat32Array } }
func (l Float64ArrayList) HeadOption() Float64ArrayOption { if l.NonEmpty() { return MkFloat64ArrayOption(*l.head) } else { return NoneFloat64Array } }
func (l AnyArrayList) HeadOption() AnyArrayOption { if l.NonEmpty() { return MkAnyArrayOption(*l.head) } else { return NoneAnyArray } }
func (l Tuple2ArrayList) HeadOption() Tuple2ArrayOption { if l.NonEmpty() { return MkTuple2ArrayOption(*l.head) } else { return NoneTuple2Array } }
func (l BoolOptionList) HeadOption() BoolOptionOption { if l.NonEmpty() { return MkBoolOptionOption(*l.head) } else { return NoneBoolOption } }
func (l StringOptionList) HeadOption() StringOptionOption { if l.NonEmpty() { return MkStringOptionOption(*l.head) } else { return NoneStringOption } }
func (l IntOptionList) HeadOption() IntOptionOption { if l.NonEmpty() { return MkIntOptionOption(*l.head) } else { return NoneIntOption } }
func (l Int64OptionList) HeadOption() Int64OptionOption { if l.NonEmpty() { return MkInt64OptionOption(*l.head) } else { return NoneInt64Option } }
func (l ByteOptionList) HeadOption() ByteOptionOption { if l.NonEmpty() { return MkByteOptionOption(*l.head) } else { return NoneByteOption } }
func (l RuneOptionList) HeadOption() RuneOptionOption { if l.NonEmpty() { return MkRuneOptionOption(*l.head) } else { return NoneRuneOption } }
func (l Float32OptionList) HeadOption() Float32OptionOption { if l.NonEmpty() { return MkFloat32OptionOption(*l.head) } else { return NoneFloat32Option } }
func (l Float64OptionList) HeadOption() Float64OptionOption { if l.NonEmpty() { return MkFloat64OptionOption(*l.head) } else { return NoneFloat64Option } }
func (l AnyOptionList) HeadOption() AnyOptionOption { if l.NonEmpty() { return MkAnyOptionOption(*l.head) } else { return NoneAnyOption } }
func (l Tuple2OptionList) HeadOption() Tuple2OptionOption { if l.NonEmpty() { return MkTuple2OptionOption(*l.head) } else { return NoneTuple2Option } }
func (l BoolListList) HeadOption() BoolListOption { if l.NonEmpty() { return MkBoolListOption(*l.head) } else { return NoneBoolList } }
func (l StringListList) HeadOption() StringListOption { if l.NonEmpty() { return MkStringListOption(*l.head) } else { return NoneStringList } }
func (l IntListList) HeadOption() IntListOption { if l.NonEmpty() { return MkIntListOption(*l.head) } else { return NoneIntList } }
func (l Int64ListList) HeadOption() Int64ListOption { if l.NonEmpty() { return MkInt64ListOption(*l.head) } else { return NoneInt64List } }
func (l ByteListList) HeadOption() ByteListOption { if l.NonEmpty() { return MkByteListOption(*l.head) } else { return NoneByteList } }
func (l RuneListList) HeadOption() RuneListOption { if l.NonEmpty() { return MkRuneListOption(*l.head) } else { return NoneRuneList } }
func (l Float32ListList) HeadOption() Float32ListOption { if l.NonEmpty() { return MkFloat32ListOption(*l.head) } else { return NoneFloat32List } }
func (l Float64ListList) HeadOption() Float64ListOption { if l.NonEmpty() { return MkFloat64ListOption(*l.head) } else { return NoneFloat64List } }
func (l AnyListList) HeadOption() AnyListOption { if l.NonEmpty() { return MkAnyListOption(*l.head) } else { return NoneAnyList } }
func (l Tuple2ListList) HeadOption() Tuple2ListOption { if l.NonEmpty() { return MkTuple2ListOption(*l.head) } else { return NoneTuple2List } }