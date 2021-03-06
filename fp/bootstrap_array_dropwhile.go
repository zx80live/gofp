// bootstrap_array_dropwhile.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp


func (l BoolArray) DropWhile(p func(bool) bool) BoolArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]bool, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l StringArray) DropWhile(p func(string) bool) StringArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]string, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l IntArray) DropWhile(p func(int) bool) IntArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]int, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Int64Array) DropWhile(p func(int64) bool) Int64Array {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]int64, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l ByteArray) DropWhile(p func(byte) bool) ByteArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]byte, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l RuneArray) DropWhile(p func(rune) bool) RuneArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]rune, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Float32Array) DropWhile(p func(float32) bool) Float32Array {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]float32, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Float64Array) DropWhile(p func(float64) bool) Float64Array {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]float64, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l AnyArray) DropWhile(p func(Any) bool) AnyArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]Any, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Tuple2Array) DropWhile(p func(Tuple2) bool) Tuple2Array {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([]Tuple2, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l BoolArrayArray) DropWhile(p func([]bool) bool) BoolArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]bool, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l StringArrayArray) DropWhile(p func([]string) bool) StringArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]string, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l IntArrayArray) DropWhile(p func([]int) bool) IntArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]int, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Int64ArrayArray) DropWhile(p func([]int64) bool) Int64ArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]int64, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l ByteArrayArray) DropWhile(p func([]byte) bool) ByteArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]byte, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l RuneArrayArray) DropWhile(p func([]rune) bool) RuneArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]rune, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Float32ArrayArray) DropWhile(p func([]float32) bool) Float32ArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]float32, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Float64ArrayArray) DropWhile(p func([]float64) bool) Float64ArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]float64, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l AnyArrayArray) DropWhile(p func([]Any) bool) AnyArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]Any, size - n)
  copy(acc, l[n: size])
  return acc
}
func (l Tuple2ArrayArray) DropWhile(p func([]Tuple2) bool) Tuple2ArrayArray {
  size := len(l)
  var n int
  for n = 0; n < size && p(l[n]); n ++ {}
  acc := make([][]Tuple2, size - n)
  copy(acc, l[n: size])
  return acc
}