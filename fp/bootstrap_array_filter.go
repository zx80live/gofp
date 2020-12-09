// bootstrap_array_filter.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (m BoolArray) Filter(p BoolPredicate) BoolArray {
	l := len(m)
	acc := make([]bool, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m StringArray) Filter(p StringPredicate) StringArray {
	l := len(m)
	acc := make([]string, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m IntArray) Filter(p IntPredicate) IntArray {
	l := len(m)
	acc := make([]int, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Int64Array) Filter(p Int64Predicate) Int64Array {
	l := len(m)
	acc := make([]int64, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m ByteArray) Filter(p BytePredicate) ByteArray {
	l := len(m)
	acc := make([]byte, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m RuneArray) Filter(p RunePredicate) RuneArray {
	l := len(m)
	acc := make([]rune, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Float32Array) Filter(p Float32Predicate) Float32Array {
	l := len(m)
	acc := make([]float32, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Float64Array) Filter(p Float64Predicate) Float64Array {
	l := len(m)
	acc := make([]float64, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m AnyArray) Filter(p AnyPredicate) AnyArray {
	l := len(m)
	acc := make([]Any, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Tuple2Array) Filter(p Tuple2Predicate) Tuple2Array {
	l := len(m)
	acc := make([]Tuple2, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m BoolArrayArray) Filter(p BoolArrayPredicate) BoolArrayArray {
	l := len(m)
	acc := make([][]bool, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m StringArrayArray) Filter(p StringArrayPredicate) StringArrayArray {
	l := len(m)
	acc := make([][]string, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m IntArrayArray) Filter(p IntArrayPredicate) IntArrayArray {
	l := len(m)
	acc := make([][]int, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Int64ArrayArray) Filter(p Int64ArrayPredicate) Int64ArrayArray {
	l := len(m)
	acc := make([][]int64, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m ByteArrayArray) Filter(p ByteArrayPredicate) ByteArrayArray {
	l := len(m)
	acc := make([][]byte, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m RuneArrayArray) Filter(p RuneArrayPredicate) RuneArrayArray {
	l := len(m)
	acc := make([][]rune, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Float32ArrayArray) Filter(p Float32ArrayPredicate) Float32ArrayArray {
	l := len(m)
	acc := make([][]float32, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Float64ArrayArray) Filter(p Float64ArrayPredicate) Float64ArrayArray {
	l := len(m)
	acc := make([][]float64, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m AnyArrayArray) Filter(p AnyArrayPredicate) AnyArrayArray {
	l := len(m)
	acc := make([][]Any, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
func (m Tuple2ArrayArray) Filter(p Tuple2ArrayPredicate) Tuple2ArrayArray {
	l := len(m)
	acc := make([][]Tuple2, l)
	i := 0
	for _, e := range m {
		if p(e) {
			acc[i] = e
			i++
		}
	}
	return acc[0:i]
}
