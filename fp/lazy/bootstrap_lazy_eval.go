// bootstrap_lazy_eval.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package lazy

func (n LazyBool) Eval() LazyBool {
	if n.cached != nil {
		//fmt.Println(" Bool.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Bool.eval", cached)
		return LazyBool{n.eval, &cached}
	}
}
func (n LazyString) Eval() LazyString {
	if n.cached != nil {
		//fmt.Println(" String.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*String.eval", cached)
		return LazyString{n.eval, &cached}
	}
}
func (n LazyInt) Eval() LazyInt {
	if n.cached != nil {
		//fmt.Println(" Int.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Int.eval", cached)
		return LazyInt{n.eval, &cached}
	}
}
func (n LazyInt64) Eval() LazyInt64 {
	if n.cached != nil {
		//fmt.Println(" Int64.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Int64.eval", cached)
		return LazyInt64{n.eval, &cached}
	}
}
func (n LazyByte) Eval() LazyByte {
	if n.cached != nil {
		//fmt.Println(" Byte.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Byte.eval", cached)
		return LazyByte{n.eval, &cached}
	}
}
func (n LazyRune) Eval() LazyRune {
	if n.cached != nil {
		//fmt.Println(" Rune.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Rune.eval", cached)
		return LazyRune{n.eval, &cached}
	}
}
func (n LazyFloat32) Eval() LazyFloat32 {
	if n.cached != nil {
		//fmt.Println(" Float32.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float32.eval", cached)
		return LazyFloat32{n.eval, &cached}
	}
}
func (n LazyFloat64) Eval() LazyFloat64 {
	if n.cached != nil {
		//fmt.Println(" Float64.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float64.eval", cached)
		return LazyFloat64{n.eval, &cached}
	}
}
func (n LazyAny) Eval() LazyAny {
	if n.cached != nil {
		//fmt.Println(" Any.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Any.eval", cached)
		return LazyAny{n.eval, &cached}
	}
}
func (n LazyTuple2) Eval() LazyTuple2 {
	if n.cached != nil {
		//fmt.Println(" Tuple2.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Tuple2.eval", cached)
		return LazyTuple2{n.eval, &cached}
	}
}
func (n LazyBoolArray) Eval() LazyBoolArray {
	if n.cached != nil {
		//fmt.Println(" BoolArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*BoolArray.eval", cached)
		return LazyBoolArray{n.eval, &cached}
	}
}
func (n LazyStringArray) Eval() LazyStringArray {
	if n.cached != nil {
		//fmt.Println(" StringArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*StringArray.eval", cached)
		return LazyStringArray{n.eval, &cached}
	}
}
func (n LazyIntArray) Eval() LazyIntArray {
	if n.cached != nil {
		//fmt.Println(" IntArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*IntArray.eval", cached)
		return LazyIntArray{n.eval, &cached}
	}
}
func (n LazyInt64Array) Eval() LazyInt64Array {
	if n.cached != nil {
		//fmt.Println(" Int64Array.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Int64Array.eval", cached)
		return LazyInt64Array{n.eval, &cached}
	}
}
func (n LazyByteArray) Eval() LazyByteArray {
	if n.cached != nil {
		//fmt.Println(" ByteArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*ByteArray.eval", cached)
		return LazyByteArray{n.eval, &cached}
	}
}
func (n LazyRuneArray) Eval() LazyRuneArray {
	if n.cached != nil {
		//fmt.Println(" RuneArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*RuneArray.eval", cached)
		return LazyRuneArray{n.eval, &cached}
	}
}
func (n LazyFloat32Array) Eval() LazyFloat32Array {
	if n.cached != nil {
		//fmt.Println(" Float32Array.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float32Array.eval", cached)
		return LazyFloat32Array{n.eval, &cached}
	}
}
func (n LazyFloat64Array) Eval() LazyFloat64Array {
	if n.cached != nil {
		//fmt.Println(" Float64Array.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float64Array.eval", cached)
		return LazyFloat64Array{n.eval, &cached}
	}
}
func (n LazyAnyArray) Eval() LazyAnyArray {
	if n.cached != nil {
		//fmt.Println(" AnyArray.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*AnyArray.eval", cached)
		return LazyAnyArray{n.eval, &cached}
	}
}
func (n LazyTuple2Array) Eval() LazyTuple2Array {
	if n.cached != nil {
		//fmt.Println(" Tuple2Array.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Tuple2Array.eval", cached)
		return LazyTuple2Array{n.eval, &cached}
	}
}
func (n LazyBoolOption) Eval() LazyBoolOption {
	if n.cached != nil {
		//fmt.Println(" BoolOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*BoolOption.eval", cached)
		return LazyBoolOption{n.eval, &cached}
	}
}
func (n LazyStringOption) Eval() LazyStringOption {
	if n.cached != nil {
		//fmt.Println(" StringOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*StringOption.eval", cached)
		return LazyStringOption{n.eval, &cached}
	}
}
func (n LazyIntOption) Eval() LazyIntOption {
	if n.cached != nil {
		//fmt.Println(" IntOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*IntOption.eval", cached)
		return LazyIntOption{n.eval, &cached}
	}
}
func (n LazyInt64Option) Eval() LazyInt64Option {
	if n.cached != nil {
		//fmt.Println(" Int64Option.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Int64Option.eval", cached)
		return LazyInt64Option{n.eval, &cached}
	}
}
func (n LazyByteOption) Eval() LazyByteOption {
	if n.cached != nil {
		//fmt.Println(" ByteOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*ByteOption.eval", cached)
		return LazyByteOption{n.eval, &cached}
	}
}
func (n LazyRuneOption) Eval() LazyRuneOption {
	if n.cached != nil {
		//fmt.Println(" RuneOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*RuneOption.eval", cached)
		return LazyRuneOption{n.eval, &cached}
	}
}
func (n LazyFloat32Option) Eval() LazyFloat32Option {
	if n.cached != nil {
		//fmt.Println(" Float32Option.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float32Option.eval", cached)
		return LazyFloat32Option{n.eval, &cached}
	}
}
func (n LazyFloat64Option) Eval() LazyFloat64Option {
	if n.cached != nil {
		//fmt.Println(" Float64Option.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float64Option.eval", cached)
		return LazyFloat64Option{n.eval, &cached}
	}
}
func (n LazyAnyOption) Eval() LazyAnyOption {
	if n.cached != nil {
		//fmt.Println(" AnyOption.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*AnyOption.eval", cached)
		return LazyAnyOption{n.eval, &cached}
	}
}
func (n LazyTuple2Option) Eval() LazyTuple2Option {
	if n.cached != nil {
		//fmt.Println(" Tuple2Option.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Tuple2Option.eval", cached)
		return LazyTuple2Option{n.eval, &cached}
	}
}
func (n LazyBoolList) Eval() LazyBoolList {
	if n.cached != nil {
		//fmt.Println(" BoolList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*BoolList.eval", cached)
		return LazyBoolList{n.eval, &cached}
	}
}
func (n LazyStringList) Eval() LazyStringList {
	if n.cached != nil {
		//fmt.Println(" StringList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*StringList.eval", cached)
		return LazyStringList{n.eval, &cached}
	}
}
func (n LazyIntList) Eval() LazyIntList {
	if n.cached != nil {
		//fmt.Println(" IntList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*IntList.eval", cached)
		return LazyIntList{n.eval, &cached}
	}
}
func (n LazyInt64List) Eval() LazyInt64List {
	if n.cached != nil {
		//fmt.Println(" Int64List.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Int64List.eval", cached)
		return LazyInt64List{n.eval, &cached}
	}
}
func (n LazyByteList) Eval() LazyByteList {
	if n.cached != nil {
		//fmt.Println(" ByteList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*ByteList.eval", cached)
		return LazyByteList{n.eval, &cached}
	}
}
func (n LazyRuneList) Eval() LazyRuneList {
	if n.cached != nil {
		//fmt.Println(" RuneList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*RuneList.eval", cached)
		return LazyRuneList{n.eval, &cached}
	}
}
func (n LazyFloat32List) Eval() LazyFloat32List {
	if n.cached != nil {
		//fmt.Println(" Float32List.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float32List.eval", cached)
		return LazyFloat32List{n.eval, &cached}
	}
}
func (n LazyFloat64List) Eval() LazyFloat64List {
	if n.cached != nil {
		//fmt.Println(" Float64List.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Float64List.eval", cached)
		return LazyFloat64List{n.eval, &cached}
	}
}
func (n LazyAnyList) Eval() LazyAnyList {
	if n.cached != nil {
		//fmt.Println(" AnyList.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*AnyList.eval", cached)
		return LazyAnyList{n.eval, &cached}
	}
}
func (n LazyTuple2List) Eval() LazyTuple2List {
	if n.cached != nil {
		//fmt.Println(" Tuple2List.cached", *n.cached)
		return n
	} else {
		cached := n.eval()
		//fmt.Println("*Tuple2List.eval", cached)
		return LazyTuple2List{n.eval, &cached}
	}
}
