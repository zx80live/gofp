// bootstrap_list_mkstring.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

import "fmt"
func (l BoolList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, BoolToString(Bool(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l StringList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, StringToString(String(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l IntList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, IntToString(Int(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Int64List) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Int64ToString(Int64(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l ByteList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, ByteToString(Byte(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l RuneList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, RuneToString(Rune(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float32List) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float32ToString(Float32(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float64List) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float64ToString(Float64(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l AnyList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, AnyToString(Any(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Tuple2List) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Tuple2ToString(Tuple2(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l BoolArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, BoolArrayToString(BoolArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l StringArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, StringArrayToString(StringArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l IntArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, IntArrayToString(IntArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Int64ArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Int64ArrayToString(Int64Array(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l ByteArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, ByteArrayToString(ByteArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l RuneArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, RuneArrayToString(RuneArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float32ArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float32ArrayToString(Float32Array(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float64ArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float64ArrayToString(Float64Array(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l AnyArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, AnyArrayToString(AnyArray(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Tuple2ArrayList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Tuple2ArrayToString(Tuple2Array(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l BoolOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, BoolOptionToString(BoolOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l StringOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, StringOptionToString(StringOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l IntOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, IntOptionToString(IntOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Int64OptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Int64OptionToString(Int64Option(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l ByteOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, ByteOptionToString(ByteOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l RuneOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, RuneOptionToString(RuneOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float32OptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float32OptionToString(Float32Option(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float64OptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float64OptionToString(Float64Option(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l AnyOptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, AnyOptionToString(AnyOption(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Tuple2OptionList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Tuple2OptionToString(Tuple2Option(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l BoolListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, BoolListToString(BoolList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l StringListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, StringListToString(StringList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l IntListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, IntListToString(IntList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Int64ListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Int64ListToString(Int64List(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l ByteListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, ByteListToString(ByteList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l RuneListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, RuneListToString(RuneList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float32ListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float32ListToString(Float32List(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Float64ListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Float64ListToString(Float64List(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l AnyListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, AnyListToString(AnyList(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}
func (l Tuple2ListList) MkString(start, sep, end string) String {
   content := ""
   xs := l
   for xs.NonEmpty() {
     content = fmt.Sprintf("%v%v%v", content, Tuple2ListToString(Tuple2List(*xs.head)), sep)
     xs = *xs.tail
   }
	 s := len(content)
	 if s > 0 {
		 content = content[:s-1]
	 }
	 return String(fmt.Sprintf("%v%v%v", start, content, end))
}