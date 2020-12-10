// bootstrap_list_dropwhile.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (l BoolList) DropWhile(p func(bool) bool) BoolList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l StringList) DropWhile(p func(string) bool) StringList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l IntList) DropWhile(p func(int) bool) IntList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Int64List) DropWhile(p func(int64) bool) Int64List {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l ByteList) DropWhile(p func(byte) bool) ByteList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l RuneList) DropWhile(p func(rune) bool) RuneList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float32List) DropWhile(p func(float32) bool) Float32List {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float64List) DropWhile(p func(float64) bool) Float64List {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l AnyList) DropWhile(p func(Any) bool) AnyList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Tuple2List) DropWhile(p func(Tuple2) bool) Tuple2List {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l BoolArrayList) DropWhile(p func([]bool) bool) BoolArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l StringArrayList) DropWhile(p func([]string) bool) StringArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l IntArrayList) DropWhile(p func([]int) bool) IntArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Int64ArrayList) DropWhile(p func([]int64) bool) Int64ArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l ByteArrayList) DropWhile(p func([]byte) bool) ByteArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l RuneArrayList) DropWhile(p func([]rune) bool) RuneArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float32ArrayList) DropWhile(p func([]float32) bool) Float32ArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float64ArrayList) DropWhile(p func([]float64) bool) Float64ArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l AnyArrayList) DropWhile(p func([]Any) bool) AnyArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Tuple2ArrayList) DropWhile(p func([]Tuple2) bool) Tuple2ArrayList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l BoolOptionList) DropWhile(p func(BoolOption) bool) BoolOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l StringOptionList) DropWhile(p func(StringOption) bool) StringOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l IntOptionList) DropWhile(p func(IntOption) bool) IntOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Int64OptionList) DropWhile(p func(Int64Option) bool) Int64OptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l ByteOptionList) DropWhile(p func(ByteOption) bool) ByteOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l RuneOptionList) DropWhile(p func(RuneOption) bool) RuneOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float32OptionList) DropWhile(p func(Float32Option) bool) Float32OptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float64OptionList) DropWhile(p func(Float64Option) bool) Float64OptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l AnyOptionList) DropWhile(p func(AnyOption) bool) AnyOptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Tuple2OptionList) DropWhile(p func(Tuple2Option) bool) Tuple2OptionList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l BoolListList) DropWhile(p func(BoolList) bool) BoolListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l StringListList) DropWhile(p func(StringList) bool) StringListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l IntListList) DropWhile(p func(IntList) bool) IntListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Int64ListList) DropWhile(p func(Int64List) bool) Int64ListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l ByteListList) DropWhile(p func(ByteList) bool) ByteListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l RuneListList) DropWhile(p func(RuneList) bool) RuneListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float32ListList) DropWhile(p func(Float32List) bool) Float32ListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Float64ListList) DropWhile(p func(Float64List) bool) Float64ListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l AnyListList) DropWhile(p func(AnyList) bool) AnyListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}
func (l Tuple2ListList) DropWhile(p func(Tuple2List) bool) Tuple2ListList {
  xs := l
  for xs.NonEmpty() && p(*xs.head) {
    xs = *xs.tail
  }
  return xs
}