// bootstrap_array_takewhile.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolArray) TakeWhile(p func(bool) bool) BoolArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]bool, n)
	copy(acc, l)
	return acc
}
func (l StringArray) TakeWhile(p func(string) bool) StringArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]string, n)
	copy(acc, l)
	return acc
}
func (l IntArray) TakeWhile(p func(int) bool) IntArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]int, n)
	copy(acc, l)
	return acc
}
func (l Int64Array) TakeWhile(p func(int64) bool) Int64Array {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]int64, n)
	copy(acc, l)
	return acc
}
func (l ByteArray) TakeWhile(p func(byte) bool) ByteArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]byte, n)
	copy(acc, l)
	return acc
}
func (l RuneArray) TakeWhile(p func(rune) bool) RuneArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]rune, n)
	copy(acc, l)
	return acc
}
func (l Float32Array) TakeWhile(p func(float32) bool) Float32Array {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]float32, n)
	copy(acc, l)
	return acc
}
func (l Float64Array) TakeWhile(p func(float64) bool) Float64Array {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]float64, n)
	copy(acc, l)
	return acc
}
func (l AnyArray) TakeWhile(p func(Any) bool) AnyArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([]Any, n)
	copy(acc, l)
	return acc
}
func (l BoolArrayArray) TakeWhile(p func([]bool) bool) BoolArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]bool, n)
	copy(acc, l)
	return acc
}
func (l StringArrayArray) TakeWhile(p func([]string) bool) StringArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]string, n)
	copy(acc, l)
	return acc
}
func (l IntArrayArray) TakeWhile(p func([]int) bool) IntArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]int, n)
	copy(acc, l)
	return acc
}
func (l Int64ArrayArray) TakeWhile(p func([]int64) bool) Int64ArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]int64, n)
	copy(acc, l)
	return acc
}
func (l ByteArrayArray) TakeWhile(p func([]byte) bool) ByteArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]byte, n)
	copy(acc, l)
	return acc
}
func (l RuneArrayArray) TakeWhile(p func([]rune) bool) RuneArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]rune, n)
	copy(acc, l)
	return acc
}
func (l Float32ArrayArray) TakeWhile(p func([]float32) bool) Float32ArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]float32, n)
	copy(acc, l)
	return acc
}
func (l Float64ArrayArray) TakeWhile(p func([]float64) bool) Float64ArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]float64, n)
	copy(acc, l)
	return acc
}
func (l AnyArrayArray) TakeWhile(p func([]Any) bool) AnyArrayArray {
	var n int
	size := len(l)
	for n = 0; n < size && p(l[n]); n++ {
	}
	acc := make([][]Any, n)
	copy(acc, l)
	return acc
}
