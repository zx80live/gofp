// fmkstring.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

import "fmt"

func MkStringBoolArr(arr []bool, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringRuneArr(arr []rune, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringByteArr(arr []byte, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringIntArr(arr []int, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt8Arr(arr []int8, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt16Arr(arr []int16, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt32Arr(arr []int32, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInt64Arr(arr []int64, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUintArr(arr []uint, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint8Arr(arr []uint8, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint16Arr(arr []uint16, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint32Arr(arr []uint32, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUint64Arr(arr []uint64, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringUintptrArr(arr []uintptr, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringFloat32Arr(arr []float32, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringFloat64Arr(arr []float64, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringComplex64Arr(arr []complex64, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringComplex128Arr(arr []complex128, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringStringArr(arr []string, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}

func MkStringInterfaceArr(arr []interface{}, left, delim, right string) string {
	content := ""
	for _, e := range arr {
		content = fmt.Sprintf("%v%v,", content, e)
	}
	l := len(content)
	if l > 0 {
		content = content[:l-1]
	}
	return fmt.Sprintf("%v%v%v", left, content, right)
}