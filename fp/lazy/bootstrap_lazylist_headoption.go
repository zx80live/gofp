// bootstrap_lazylist_headoption.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package lazy

import . "github.com/zx80live/gofp/fp"

func (l BoolLazyList) HeadOption() (BoolOption, BoolLazyList) {
	if l.IsEmpty() {
		return NoneBool, NilBoolLazyList
	} else {
		s := (*l.state)()
		return MkBoolOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l StringLazyList) HeadOption() (StringOption, StringLazyList) {
	if l.IsEmpty() {
		return NoneString, NilStringLazyList
	} else {
		s := (*l.state)()
		return MkStringOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l IntLazyList) HeadOption() (IntOption, IntLazyList) {
	if l.IsEmpty() {
		return NoneInt, NilIntLazyList
	} else {
		s := (*l.state)()
		return MkIntOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Int64LazyList) HeadOption() (Int64Option, Int64LazyList) {
	if l.IsEmpty() {
		return NoneInt64, NilInt64LazyList
	} else {
		s := (*l.state)()
		return MkInt64Option(s.head.Eval().Cached()), *s.tail
	}
}

func (l ByteLazyList) HeadOption() (ByteOption, ByteLazyList) {
	if l.IsEmpty() {
		return NoneByte, NilByteLazyList
	} else {
		s := (*l.state)()
		return MkByteOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l RuneLazyList) HeadOption() (RuneOption, RuneLazyList) {
	if l.IsEmpty() {
		return NoneRune, NilRuneLazyList
	} else {
		s := (*l.state)()
		return MkRuneOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float32LazyList) HeadOption() (Float32Option, Float32LazyList) {
	if l.IsEmpty() {
		return NoneFloat32, NilFloat32LazyList
	} else {
		s := (*l.state)()
		return MkFloat32Option(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float64LazyList) HeadOption() (Float64Option, Float64LazyList) {
	if l.IsEmpty() {
		return NoneFloat64, NilFloat64LazyList
	} else {
		s := (*l.state)()
		return MkFloat64Option(s.head.Eval().Cached()), *s.tail
	}
}

func (l AnyLazyList) HeadOption() (AnyOption, AnyLazyList) {
	if l.IsEmpty() {
		return NoneAny, NilAnyLazyList
	} else {
		s := (*l.state)()
		return MkAnyOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Tuple2LazyList) HeadOption() (Tuple2Option, Tuple2LazyList) {
	if l.IsEmpty() {
		return NoneTuple2, NilTuple2LazyList
	} else {
		s := (*l.state)()
		return MkTuple2Option(s.head.Eval().Cached()), *s.tail
	}
}

func (l BoolArrayLazyList) HeadOption() (BoolArrayOption, BoolArrayLazyList) {
	if l.IsEmpty() {
		return NoneBoolArray, NilBoolArrayLazyList
	} else {
		s := (*l.state)()
		return MkBoolArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l StringArrayLazyList) HeadOption() (StringArrayOption, StringArrayLazyList) {
	if l.IsEmpty() {
		return NoneStringArray, NilStringArrayLazyList
	} else {
		s := (*l.state)()
		return MkStringArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l IntArrayLazyList) HeadOption() (IntArrayOption, IntArrayLazyList) {
	if l.IsEmpty() {
		return NoneIntArray, NilIntArrayLazyList
	} else {
		s := (*l.state)()
		return MkIntArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Int64ArrayLazyList) HeadOption() (Int64ArrayOption, Int64ArrayLazyList) {
	if l.IsEmpty() {
		return NoneInt64Array, NilInt64ArrayLazyList
	} else {
		s := (*l.state)()
		return MkInt64ArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l ByteArrayLazyList) HeadOption() (ByteArrayOption, ByteArrayLazyList) {
	if l.IsEmpty() {
		return NoneByteArray, NilByteArrayLazyList
	} else {
		s := (*l.state)()
		return MkByteArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l RuneArrayLazyList) HeadOption() (RuneArrayOption, RuneArrayLazyList) {
	if l.IsEmpty() {
		return NoneRuneArray, NilRuneArrayLazyList
	} else {
		s := (*l.state)()
		return MkRuneArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float32ArrayLazyList) HeadOption() (Float32ArrayOption, Float32ArrayLazyList) {
	if l.IsEmpty() {
		return NoneFloat32Array, NilFloat32ArrayLazyList
	} else {
		s := (*l.state)()
		return MkFloat32ArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float64ArrayLazyList) HeadOption() (Float64ArrayOption, Float64ArrayLazyList) {
	if l.IsEmpty() {
		return NoneFloat64Array, NilFloat64ArrayLazyList
	} else {
		s := (*l.state)()
		return MkFloat64ArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l AnyArrayLazyList) HeadOption() (AnyArrayOption, AnyArrayLazyList) {
	if l.IsEmpty() {
		return NoneAnyArray, NilAnyArrayLazyList
	} else {
		s := (*l.state)()
		return MkAnyArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Tuple2ArrayLazyList) HeadOption() (Tuple2ArrayOption, Tuple2ArrayLazyList) {
	if l.IsEmpty() {
		return NoneTuple2Array, NilTuple2ArrayLazyList
	} else {
		s := (*l.state)()
		return MkTuple2ArrayOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l BoolOptionLazyList) HeadOption() (BoolOptionOption, BoolOptionLazyList) {
	if l.IsEmpty() {
		return NoneBoolOption, NilBoolOptionLazyList
	} else {
		s := (*l.state)()
		return MkBoolOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l StringOptionLazyList) HeadOption() (StringOptionOption, StringOptionLazyList) {
	if l.IsEmpty() {
		return NoneStringOption, NilStringOptionLazyList
	} else {
		s := (*l.state)()
		return MkStringOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l IntOptionLazyList) HeadOption() (IntOptionOption, IntOptionLazyList) {
	if l.IsEmpty() {
		return NoneIntOption, NilIntOptionLazyList
	} else {
		s := (*l.state)()
		return MkIntOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Int64OptionLazyList) HeadOption() (Int64OptionOption, Int64OptionLazyList) {
	if l.IsEmpty() {
		return NoneInt64Option, NilInt64OptionLazyList
	} else {
		s := (*l.state)()
		return MkInt64OptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l ByteOptionLazyList) HeadOption() (ByteOptionOption, ByteOptionLazyList) {
	if l.IsEmpty() {
		return NoneByteOption, NilByteOptionLazyList
	} else {
		s := (*l.state)()
		return MkByteOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l RuneOptionLazyList) HeadOption() (RuneOptionOption, RuneOptionLazyList) {
	if l.IsEmpty() {
		return NoneRuneOption, NilRuneOptionLazyList
	} else {
		s := (*l.state)()
		return MkRuneOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float32OptionLazyList) HeadOption() (Float32OptionOption, Float32OptionLazyList) {
	if l.IsEmpty() {
		return NoneFloat32Option, NilFloat32OptionLazyList
	} else {
		s := (*l.state)()
		return MkFloat32OptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float64OptionLazyList) HeadOption() (Float64OptionOption, Float64OptionLazyList) {
	if l.IsEmpty() {
		return NoneFloat64Option, NilFloat64OptionLazyList
	} else {
		s := (*l.state)()
		return MkFloat64OptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l AnyOptionLazyList) HeadOption() (AnyOptionOption, AnyOptionLazyList) {
	if l.IsEmpty() {
		return NoneAnyOption, NilAnyOptionLazyList
	} else {
		s := (*l.state)()
		return MkAnyOptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Tuple2OptionLazyList) HeadOption() (Tuple2OptionOption, Tuple2OptionLazyList) {
	if l.IsEmpty() {
		return NoneTuple2Option, NilTuple2OptionLazyList
	} else {
		s := (*l.state)()
		return MkTuple2OptionOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l BoolListLazyList) HeadOption() (BoolListOption, BoolListLazyList) {
	if l.IsEmpty() {
		return NoneBoolList, NilBoolListLazyList
	} else {
		s := (*l.state)()
		return MkBoolListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l StringListLazyList) HeadOption() (StringListOption, StringListLazyList) {
	if l.IsEmpty() {
		return NoneStringList, NilStringListLazyList
	} else {
		s := (*l.state)()
		return MkStringListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l IntListLazyList) HeadOption() (IntListOption, IntListLazyList) {
	if l.IsEmpty() {
		return NoneIntList, NilIntListLazyList
	} else {
		s := (*l.state)()
		return MkIntListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Int64ListLazyList) HeadOption() (Int64ListOption, Int64ListLazyList) {
	if l.IsEmpty() {
		return NoneInt64List, NilInt64ListLazyList
	} else {
		s := (*l.state)()
		return MkInt64ListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l ByteListLazyList) HeadOption() (ByteListOption, ByteListLazyList) {
	if l.IsEmpty() {
		return NoneByteList, NilByteListLazyList
	} else {
		s := (*l.state)()
		return MkByteListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l RuneListLazyList) HeadOption() (RuneListOption, RuneListLazyList) {
	if l.IsEmpty() {
		return NoneRuneList, NilRuneListLazyList
	} else {
		s := (*l.state)()
		return MkRuneListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float32ListLazyList) HeadOption() (Float32ListOption, Float32ListLazyList) {
	if l.IsEmpty() {
		return NoneFloat32List, NilFloat32ListLazyList
	} else {
		s := (*l.state)()
		return MkFloat32ListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Float64ListLazyList) HeadOption() (Float64ListOption, Float64ListLazyList) {
	if l.IsEmpty() {
		return NoneFloat64List, NilFloat64ListLazyList
	} else {
		s := (*l.state)()
		return MkFloat64ListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l AnyListLazyList) HeadOption() (AnyListOption, AnyListLazyList) {
	if l.IsEmpty() {
		return NoneAnyList, NilAnyListLazyList
	} else {
		s := (*l.state)()
		return MkAnyListOption(s.head.Eval().Cached()), *s.tail
	}
}

func (l Tuple2ListLazyList) HeadOption() (Tuple2ListOption, Tuple2ListLazyList) {
	if l.IsEmpty() {
		return NoneTuple2List, NilTuple2ListLazyList
	} else {
		s := (*l.state)()
		return MkTuple2ListOption(s.head.Eval().Cached()), *s.tail
	}
}