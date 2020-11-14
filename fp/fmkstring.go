// fmkstring.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

import "fmt"

func MkStringBoolArr(arr []bool, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringRuneArr(arr []rune, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringByteArr(arr []byte, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringIntArr(arr []int, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt8Arr(arr []int8, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt16Arr(arr []int16, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt32Arr(arr []int32, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt64Arr(arr []int64, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUintArr(arr []uint, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint8Arr(arr []uint8, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint16Arr(arr []uint16, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint32Arr(arr []uint32, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint64Arr(arr []uint64, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUintptrArr(arr []uintptr, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringFloat32Arr(arr []float32, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringFloat64Arr(arr []float64, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringComplex64Arr(arr []complex64, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringComplex128Arr(arr []complex128, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringStringArr(arr []string, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringAnyArr(arr []Any, left, sep, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v%v", content, e, sep)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}
