// bootstrap_list_reverse.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (l BoolList) Reverse() BoolList {
  acc := NilBool
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l StringList) Reverse() StringList {
  acc := NilString
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l IntList) Reverse() IntList {
  acc := NilInt
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Int64List) Reverse() Int64List {
  acc := NilInt64
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l ByteList) Reverse() ByteList {
  acc := NilByte
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l RuneList) Reverse() RuneList {
  acc := NilRune
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float32List) Reverse() Float32List {
  acc := NilFloat32
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float64List) Reverse() Float64List {
  acc := NilFloat64
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l AnyList) Reverse() AnyList {
  acc := NilAny
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Tuple2List) Reverse() Tuple2List {
  acc := NilTuple2
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l BoolArrayList) Reverse() BoolArrayList {
  acc := NilBoolArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l StringArrayList) Reverse() StringArrayList {
  acc := NilStringArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l IntArrayList) Reverse() IntArrayList {
  acc := NilIntArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Int64ArrayList) Reverse() Int64ArrayList {
  acc := NilInt64Array
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l ByteArrayList) Reverse() ByteArrayList {
  acc := NilByteArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l RuneArrayList) Reverse() RuneArrayList {
  acc := NilRuneArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float32ArrayList) Reverse() Float32ArrayList {
  acc := NilFloat32Array
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float64ArrayList) Reverse() Float64ArrayList {
  acc := NilFloat64Array
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l AnyArrayList) Reverse() AnyArrayList {
  acc := NilAnyArray
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Tuple2ArrayList) Reverse() Tuple2ArrayList {
  acc := NilTuple2Array
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l BoolOptionList) Reverse() BoolOptionList {
  acc := NilBoolOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l StringOptionList) Reverse() StringOptionList {
  acc := NilStringOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l IntOptionList) Reverse() IntOptionList {
  acc := NilIntOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Int64OptionList) Reverse() Int64OptionList {
  acc := NilInt64Option
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l ByteOptionList) Reverse() ByteOptionList {
  acc := NilByteOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l RuneOptionList) Reverse() RuneOptionList {
  acc := NilRuneOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float32OptionList) Reverse() Float32OptionList {
  acc := NilFloat32Option
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float64OptionList) Reverse() Float64OptionList {
  acc := NilFloat64Option
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l AnyOptionList) Reverse() AnyOptionList {
  acc := NilAnyOption
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Tuple2OptionList) Reverse() Tuple2OptionList {
  acc := NilTuple2Option
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l BoolListList) Reverse() BoolListList {
  acc := NilBoolList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l StringListList) Reverse() StringListList {
  acc := NilStringList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l IntListList) Reverse() IntListList {
  acc := NilIntList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Int64ListList) Reverse() Int64ListList {
  acc := NilInt64List
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l ByteListList) Reverse() ByteListList {
  acc := NilByteList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l RuneListList) Reverse() RuneListList {
  acc := NilRuneList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float32ListList) Reverse() Float32ListList {
  acc := NilFloat32List
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Float64ListList) Reverse() Float64ListList {
  acc := NilFloat64List
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l AnyListList) Reverse() AnyListList {
  acc := NilAnyList
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}
func (l Tuple2ListList) Reverse() Tuple2ListList {
  acc := NilTuple2List
  xs := l
  for xs.NonEmpty() {
    acc = acc.Cons(*xs.head)
    xs = *xs.tail
  }
  return acc
}