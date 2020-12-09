// bootstrap_array_count.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (l BoolArray) Count(p func(bool) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l StringArray) Count(p func(string) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l IntArray) Count(p func(int) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Int64Array) Count(p func(int64) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l ByteArray) Count(p func(byte) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l RuneArray) Count(p func(rune) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Float32Array) Count(p func(float32) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Float64Array) Count(p func(float64) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l AnyArray) Count(p func(Any) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Tuple2Array) Count(p func(Tuple2) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l BoolArrayArray) Count(p func([]bool) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l StringArrayArray) Count(p func([]string) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l IntArrayArray) Count(p func([]int) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Int64ArrayArray) Count(p func([]int64) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l ByteArrayArray) Count(p func([]byte) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l RuneArrayArray) Count(p func([]rune) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Float32ArrayArray) Count(p func([]float32) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Float64ArrayArray) Count(p func([]float64) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l AnyArrayArray) Count(p func([]Any) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
func (l Tuple2ArrayArray) Count(p func([]Tuple2) bool) int {
	c := 0
	for _, e := range l {
		if p(e) {
			c++
		}
	}
	return c
}
