// bootstrap_option_shortcons.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func Bool(e bool) BoolOption          { return BoolOption{&e} }
func String(e string) StringOption    { return StringOption{&e} }
func Int(e int) IntOption             { return IntOption{&e} }
func Int64(e int64) Int64Option       { return Int64Option{&e} }
func Byte(e byte) ByteOption          { return ByteOption{&e} }
func Rune(e rune) RuneOption          { return RuneOption{&e} }
func Float32(e float32) Float32Option { return Float32Option{&e} }
func Float64(e float64) Float64Option { return Float64Option{&e} }
func AnyOpt(e Any) AnyOption          { return AnyOption{&e} }
