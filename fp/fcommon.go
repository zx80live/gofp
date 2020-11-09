package fp

type PredicateBool func(e bool) bool
type PredicateString func(e string) bool
type PredicateInt func(e int) bool
type PredicateInt8 func(e int8) bool
type PredicateInt16 func(e int16) bool
type PredicateInt32 func(e int32) bool
type PredicateUInt func(e uint) bool
type PredicateUInt8 func(e uint8) bool
type PredicateUInt16 func(e uint16) bool
type PredicateUInt32 func(e uint32) bool
type PredicateUInt64 func(e uint64) bool
type PredicateUIntPtr func(e uintptr) bool
type PredicateByte func(e byte) bool
type PredicateRune func(e rune) bool
type PredicateFloat32 func(e float32) bool
type PredicateFloat64 func(e float64) bool
type PredicateComplex64 func(e complex64) bool
type PredicateComplex128 func(e complex128) bool

type FunctorBool func(e bool) bool



type FunctorString func(e string) string
type FunctorStringInt func(e string) int



type FunctorInt func(e int) int
type FunctorIntString func(e int) string



type FunctorInt8 func(e int8) int8



type FunctorInt16 func(e int16) int16



type FunctorInt32 func(e int32) int32



type FunctorUInt func(e uint) uint


type FunctorUInt8 func(e uint8) uint8



type FunctorUInt16 func(e uint16) uint16



type FunctorUInt32 func(e uint32) uint32



type FunctorUInt64 func(e uint64) uint64



type FunctorUIntPtr func(e uintptr) uintptr



type FunctorByte func(e byte) byte



type FunctorRune func(e rune) rune



type FunctorFloat32 func(e float32) float32



type FunctorFloat64 func(e float64) float64



type FunctorComplex64 func(e complex64) complex64



type FunctorComplex128 func(e complex128) complex128

