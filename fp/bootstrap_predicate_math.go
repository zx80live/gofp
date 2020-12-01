// bootstrap_predicate_math.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

var EvenInt IntPredicate = func(t int) bool { return t%2 == 0 }
var EvenInt8 Int8Predicate = func(t int8) bool { return t%2 == 0 }
var EvenInt16 Int16Predicate = func(t int16) bool { return t%2 == 0 }
var EvenInt32 Int32Predicate = func(t int32) bool { return t%2 == 0 }
var EvenInt64 Int64Predicate = func(t int64) bool { return t%2 == 0 }
var EvenUint UintPredicate = func(t uint) bool { return t%2 == 0 }
var EvenUint8 Uint8Predicate = func(t uint8) bool { return t%2 == 0 }
var EvenUint16 Uint16Predicate = func(t uint16) bool { return t%2 == 0 }
var EvenUint32 Uint32Predicate = func(t uint32) bool { return t%2 == 0 }
var EvenUint64 Uint64Predicate = func(t uint64) bool { return t%2 == 0 }
var EvenUintptr UintptrPredicate = func(t uintptr) bool { return t%2 == 0 }
var EvenByte BytePredicate = func(t byte) bool { return t%2 == 0 }
var EvenRune RunePredicate = func(t rune) bool { return t%2 == 0 }
var OddInt IntPredicate = func(t int) bool { return t%2 != 0 }
var OddInt8 Int8Predicate = func(t int8) bool { return t%2 != 0 }
var OddInt16 Int16Predicate = func(t int16) bool { return t%2 != 0 }
var OddInt32 Int32Predicate = func(t int32) bool { return t%2 != 0 }
var OddInt64 Int64Predicate = func(t int64) bool { return t%2 != 0 }
var OddUint UintPredicate = func(t uint) bool { return t%2 != 0 }
var OddUint8 Uint8Predicate = func(t uint8) bool { return t%2 != 0 }
var OddUint16 Uint16Predicate = func(t uint16) bool { return t%2 != 0 }
var OddUint32 Uint32Predicate = func(t uint32) bool { return t%2 != 0 }
var OddUint64 Uint64Predicate = func(t uint64) bool { return t%2 != 0 }
var OddUintptr UintptrPredicate = func(t uintptr) bool { return t%2 != 0 }
var OddByte BytePredicate = func(t byte) bool { return t%2 != 0 }
var OddRune RunePredicate = func(t rune) bool { return t%2 != 0 }
var NegInt IntPredicate = func(t int) bool { return t < 0 }
var NegInt8 Int8Predicate = func(t int8) bool { return t < 0 }
var NegInt16 Int16Predicate = func(t int16) bool { return t < 0 }
var NegInt32 Int32Predicate = func(t int32) bool { return t < 0 }
var NegInt64 Int64Predicate = func(t int64) bool { return t < 0 }
var NegUint UintPredicate = func(t uint) bool { return t < 0 }
var NegUint8 Uint8Predicate = func(t uint8) bool { return t < 0 }
var NegUint16 Uint16Predicate = func(t uint16) bool { return t < 0 }
var NegUint32 Uint32Predicate = func(t uint32) bool { return t < 0 }
var NegUint64 Uint64Predicate = func(t uint64) bool { return t < 0 }
var NegUintptr UintptrPredicate = func(t uintptr) bool { return t < 0 }
var NegByte BytePredicate = func(t byte) bool { return t < 0 }
var NegRune RunePredicate = func(t rune) bool { return t < 0 }
var PosInt IntPredicate = func(t int) bool { return t >= 0 }
var PosInt8 Int8Predicate = func(t int8) bool { return t >= 0 }
var PosInt16 Int16Predicate = func(t int16) bool { return t >= 0 }
var PosInt32 Int32Predicate = func(t int32) bool { return t >= 0 }
var PosInt64 Int64Predicate = func(t int64) bool { return t >= 0 }
var PosUint UintPredicate = func(t uint) bool { return t >= 0 }
var PosUint8 Uint8Predicate = func(t uint8) bool { return t >= 0 }
var PosUint16 Uint16Predicate = func(t uint16) bool { return t >= 0 }
var PosUint32 Uint32Predicate = func(t uint32) bool { return t >= 0 }
var PosUint64 Uint64Predicate = func(t uint64) bool { return t >= 0 }
var PosUintptr UintptrPredicate = func(t uintptr) bool { return t >= 0 }
var PosByte BytePredicate = func(t byte) bool { return t >= 0 }
var PosRune RunePredicate = func(t rune) bool { return t >= 0 }
var ZeroInt IntPredicate = func(t int) bool { return t == 0 }
var ZeroInt8 Int8Predicate = func(t int8) bool { return t == 0 }
var ZeroInt16 Int16Predicate = func(t int16) bool { return t == 0 }
var ZeroInt32 Int32Predicate = func(t int32) bool { return t == 0 }
var ZeroInt64 Int64Predicate = func(t int64) bool { return t == 0 }
var ZeroUint UintPredicate = func(t uint) bool { return t == 0 }
var ZeroUint8 Uint8Predicate = func(t uint8) bool { return t == 0 }
var ZeroUint16 Uint16Predicate = func(t uint16) bool { return t == 0 }
var ZeroUint32 Uint32Predicate = func(t uint32) bool { return t == 0 }
var ZeroUint64 Uint64Predicate = func(t uint64) bool { return t == 0 }
var ZeroUintptr UintptrPredicate = func(t uintptr) bool { return t == 0 }
var ZeroByte BytePredicate = func(t byte) bool { return t == 0 }
var ZeroRune RunePredicate = func(t rune) bool { return t == 0 }
var OneInt IntPredicate = func(t int) bool { return t == 1 }
var OneInt8 Int8Predicate = func(t int8) bool { return t == 1 }
var OneInt16 Int16Predicate = func(t int16) bool { return t == 1 }
var OneInt32 Int32Predicate = func(t int32) bool { return t == 1 }
var OneInt64 Int64Predicate = func(t int64) bool { return t == 1 }
var OneUint UintPredicate = func(t uint) bool { return t == 1 }
var OneUint8 Uint8Predicate = func(t uint8) bool { return t == 1 }
var OneUint16 Uint16Predicate = func(t uint16) bool { return t == 1 }
var OneUint32 Uint32Predicate = func(t uint32) bool { return t == 1 }
var OneUint64 Uint64Predicate = func(t uint64) bool { return t == 1 }
var OneUintptr UintptrPredicate = func(t uintptr) bool { return t == 1 }
var OneByte BytePredicate = func(t byte) bool { return t == 1 }
var OneRune RunePredicate = func(t rune) bool { return t == 1 }
