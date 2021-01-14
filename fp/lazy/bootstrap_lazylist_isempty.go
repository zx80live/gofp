// bootstrap_lazylist_isempty.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package lazy

func (l BoolLazyList) IsEmpty() bool {
	if l == NilBoolLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyBoolState
}
func (l StringLazyList) IsEmpty() bool {
	if l == NilStringLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyStringState
}
func (l IntLazyList) IsEmpty() bool {
	if l == NilIntLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyIntState
}
func (l Int64LazyList) IsEmpty() bool {
	if l == NilInt64LazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyInt64State
}
func (l ByteLazyList) IsEmpty() bool {
	if l == NilByteLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyByteState
}
func (l RuneLazyList) IsEmpty() bool {
	if l == NilRuneLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyRuneState
}
func (l Float32LazyList) IsEmpty() bool {
	if l == NilFloat32LazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat32State
}
func (l Float64LazyList) IsEmpty() bool {
	if l == NilFloat64LazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat64State
}
func (l AnyLazyList) IsEmpty() bool {
	if l == NilAnyLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyAnyState
}
func (l Tuple2LazyList) IsEmpty() bool {
	if l == NilTuple2LazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyTuple2State
}
func (l BoolArrayLazyList) IsEmpty() bool {
	if l == NilBoolArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyBoolArrayState
}
func (l StringArrayLazyList) IsEmpty() bool {
	if l == NilStringArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyStringArrayState
}
func (l IntArrayLazyList) IsEmpty() bool {
	if l == NilIntArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyIntArrayState
}
func (l Int64ArrayLazyList) IsEmpty() bool {
	if l == NilInt64ArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyInt64ArrayState
}
func (l ByteArrayLazyList) IsEmpty() bool {
	if l == NilByteArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyByteArrayState
}
func (l RuneArrayLazyList) IsEmpty() bool {
	if l == NilRuneArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyRuneArrayState
}
func (l Float32ArrayLazyList) IsEmpty() bool {
	if l == NilFloat32ArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat32ArrayState
}
func (l Float64ArrayLazyList) IsEmpty() bool {
	if l == NilFloat64ArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat64ArrayState
}
func (l AnyArrayLazyList) IsEmpty() bool {
	if l == NilAnyArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyAnyArrayState
}
func (l Tuple2ArrayLazyList) IsEmpty() bool {
	if l == NilTuple2ArrayLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyTuple2ArrayState
}
func (l BoolOptionLazyList) IsEmpty() bool {
	if l == NilBoolOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyBoolOptionState
}
func (l StringOptionLazyList) IsEmpty() bool {
	if l == NilStringOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyStringOptionState
}
func (l IntOptionLazyList) IsEmpty() bool {
	if l == NilIntOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyIntOptionState
}
func (l Int64OptionLazyList) IsEmpty() bool {
	if l == NilInt64OptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyInt64OptionState
}
func (l ByteOptionLazyList) IsEmpty() bool {
	if l == NilByteOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyByteOptionState
}
func (l RuneOptionLazyList) IsEmpty() bool {
	if l == NilRuneOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyRuneOptionState
}
func (l Float32OptionLazyList) IsEmpty() bool {
	if l == NilFloat32OptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat32OptionState
}
func (l Float64OptionLazyList) IsEmpty() bool {
	if l == NilFloat64OptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat64OptionState
}
func (l AnyOptionLazyList) IsEmpty() bool {
	if l == NilAnyOptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyAnyOptionState
}
func (l Tuple2OptionLazyList) IsEmpty() bool {
	if l == NilTuple2OptionLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyTuple2OptionState
}
func (l BoolListLazyList) IsEmpty() bool {
	if l == NilBoolListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyBoolListState
}
func (l StringListLazyList) IsEmpty() bool {
	if l == NilStringListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyStringListState
}
func (l IntListLazyList) IsEmpty() bool {
	if l == NilIntListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyIntListState
}
func (l Int64ListLazyList) IsEmpty() bool {
	if l == NilInt64ListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyInt64ListState
}
func (l ByteListLazyList) IsEmpty() bool {
	if l == NilByteListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyByteListState
}
func (l RuneListLazyList) IsEmpty() bool {
	if l == NilRuneListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyRuneListState
}
func (l Float32ListLazyList) IsEmpty() bool {
	if l == NilFloat32ListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat32ListState
}
func (l Float64ListLazyList) IsEmpty() bool {
	if l == NilFloat64ListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyFloat64ListState
}
func (l AnyListLazyList) IsEmpty() bool {
	if l == NilAnyListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyAnyListState
}
func (l Tuple2ListLazyList) IsEmpty() bool {
	if l == NilTuple2ListLazyList {
		return true
	}
	s := (*l.state)()
	return s == EmptyTuple2ListState
}