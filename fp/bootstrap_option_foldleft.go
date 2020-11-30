// bootstrap_option_foldleft.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package fp

func (o BoolOption) FoldLeftBool(z bool, f func(bool, bool) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftString(z string, f func(string, bool) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftInt(z int, f func(int, bool) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftInt64(z int64, f func(int64, bool) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftByte(z byte, f func(byte, bool) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftRune(z rune, f func(rune, bool) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftFloat32(z float32, f func(float32, bool) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftFloat64(z float64, f func(float64, bool) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o BoolOption) FoldLeftAny(z Any, f func(Any, bool) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftBool(z bool, f func(bool, string) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftString(z string, f func(string, string) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftInt(z int, f func(int, string) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftInt64(z int64, f func(int64, string) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftByte(z byte, f func(byte, string) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftRune(z rune, f func(rune, string) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftFloat32(z float32, f func(float32, string) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftFloat64(z float64, f func(float64, string) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o StringOption) FoldLeftAny(z Any, f func(Any, string) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftBool(z bool, f func(bool, int) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftString(z string, f func(string, int) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftInt(z int, f func(int, int) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftInt64(z int64, f func(int64, int) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftByte(z byte, f func(byte, int) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftRune(z rune, f func(rune, int) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftFloat32(z float32, f func(float32, int) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftFloat64(z float64, f func(float64, int) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o IntOption) FoldLeftAny(z Any, f func(Any, int) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftBool(z bool, f func(bool, int8) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftString(z string, f func(string, int8) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftInt(z int, f func(int, int8) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftInt64(z int64, f func(int64, int8) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftByte(z byte, f func(byte, int8) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftRune(z rune, f func(rune, int8) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftFloat32(z float32, f func(float32, int8) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftFloat64(z float64, f func(float64, int8) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int8Option) FoldLeftAny(z Any, f func(Any, int8) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftBool(z bool, f func(bool, int16) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftString(z string, f func(string, int16) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftInt(z int, f func(int, int16) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftInt64(z int64, f func(int64, int16) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftByte(z byte, f func(byte, int16) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftRune(z rune, f func(rune, int16) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftFloat32(z float32, f func(float32, int16) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftFloat64(z float64, f func(float64, int16) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int16Option) FoldLeftAny(z Any, f func(Any, int16) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftBool(z bool, f func(bool, int32) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftString(z string, f func(string, int32) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftInt(z int, f func(int, int32) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftInt64(z int64, f func(int64, int32) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftByte(z byte, f func(byte, int32) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftRune(z rune, f func(rune, int32) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftFloat32(z float32, f func(float32, int32) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftFloat64(z float64, f func(float64, int32) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int32Option) FoldLeftAny(z Any, f func(Any, int32) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftBool(z bool, f func(bool, int64) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftString(z string, f func(string, int64) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftInt(z int, f func(int, int64) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftInt64(z int64, f func(int64, int64) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftByte(z byte, f func(byte, int64) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftRune(z rune, f func(rune, int64) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftFloat32(z float32, f func(float32, int64) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftFloat64(z float64, f func(float64, int64) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Int64Option) FoldLeftAny(z Any, f func(Any, int64) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftBool(z bool, f func(bool, uint) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftString(z string, f func(string, uint) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftInt(z int, f func(int, uint) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftInt64(z int64, f func(int64, uint) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftByte(z byte, f func(byte, uint) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftRune(z rune, f func(rune, uint) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftFloat32(z float32, f func(float32, uint) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftFloat64(z float64, f func(float64, uint) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintOption) FoldLeftAny(z Any, f func(Any, uint) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftBool(z bool, f func(bool, uint8) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftString(z string, f func(string, uint8) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftInt(z int, f func(int, uint8) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftInt64(z int64, f func(int64, uint8) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftByte(z byte, f func(byte, uint8) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftRune(z rune, f func(rune, uint8) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftFloat32(z float32, f func(float32, uint8) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftFloat64(z float64, f func(float64, uint8) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint8Option) FoldLeftAny(z Any, f func(Any, uint8) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftBool(z bool, f func(bool, uint16) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftString(z string, f func(string, uint16) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftInt(z int, f func(int, uint16) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftInt64(z int64, f func(int64, uint16) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftByte(z byte, f func(byte, uint16) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftRune(z rune, f func(rune, uint16) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftFloat32(z float32, f func(float32, uint16) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftFloat64(z float64, f func(float64, uint16) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint16Option) FoldLeftAny(z Any, f func(Any, uint16) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftBool(z bool, f func(bool, uint32) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftString(z string, f func(string, uint32) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftInt(z int, f func(int, uint32) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftInt64(z int64, f func(int64, uint32) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftByte(z byte, f func(byte, uint32) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftRune(z rune, f func(rune, uint32) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftFloat32(z float32, f func(float32, uint32) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftFloat64(z float64, f func(float64, uint32) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint32Option) FoldLeftAny(z Any, f func(Any, uint32) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftBool(z bool, f func(bool, uint64) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftString(z string, f func(string, uint64) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftInt(z int, f func(int, uint64) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftInt64(z int64, f func(int64, uint64) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftByte(z byte, f func(byte, uint64) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftRune(z rune, f func(rune, uint64) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftFloat32(z float32, f func(float32, uint64) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftFloat64(z float64, f func(float64, uint64) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Uint64Option) FoldLeftAny(z Any, f func(Any, uint64) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftBool(z bool, f func(bool, uintptr) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftString(z string, f func(string, uintptr) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftInt(z int, f func(int, uintptr) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftInt64(z int64, f func(int64, uintptr) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftByte(z byte, f func(byte, uintptr) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftRune(z rune, f func(rune, uintptr) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftFloat32(z float32, f func(float32, uintptr) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftFloat64(z float64, f func(float64, uintptr) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o UintptrOption) FoldLeftAny(z Any, f func(Any, uintptr) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftBool(z bool, f func(bool, byte) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftString(z string, f func(string, byte) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftInt(z int, f func(int, byte) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftInt64(z int64, f func(int64, byte) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftByte(z byte, f func(byte, byte) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftRune(z rune, f func(rune, byte) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftFloat32(z float32, f func(float32, byte) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftFloat64(z float64, f func(float64, byte) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o ByteOption) FoldLeftAny(z Any, f func(Any, byte) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftBool(z bool, f func(bool, rune) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftString(z string, f func(string, rune) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftInt(z int, f func(int, rune) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftInt64(z int64, f func(int64, rune) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftByte(z byte, f func(byte, rune) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftRune(z rune, f func(rune, rune) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftFloat32(z float32, f func(float32, rune) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftFloat64(z float64, f func(float64, rune) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o RuneOption) FoldLeftAny(z Any, f func(Any, rune) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftBool(z bool, f func(bool, float32) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftString(z string, f func(string, float32) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftInt(z int, f func(int, float32) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftInt64(z int64, f func(int64, float32) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftByte(z byte, f func(byte, float32) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftRune(z rune, f func(rune, float32) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftFloat32(z float32, f func(float32, float32) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftFloat64(z float64, f func(float64, float32) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float32Option) FoldLeftAny(z Any, f func(Any, float32) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftBool(z bool, f func(bool, float64) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftString(z string, f func(string, float64) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftInt(z int, f func(int, float64) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftInt64(z int64, f func(int64, float64) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftByte(z byte, f func(byte, float64) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftRune(z rune, f func(rune, float64) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftFloat32(z float32, f func(float32, float64) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftFloat64(z float64, f func(float64, float64) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Float64Option) FoldLeftAny(z Any, f func(Any, float64) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftBool(z bool, f func(bool, complex64) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftString(z string, f func(string, complex64) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftInt(z int, f func(int, complex64) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftInt64(z int64, f func(int64, complex64) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftByte(z byte, f func(byte, complex64) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftRune(z rune, f func(rune, complex64) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftFloat32(z float32, f func(float32, complex64) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftFloat64(z float64, f func(float64, complex64) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex64Option) FoldLeftAny(z Any, f func(Any, complex64) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftBool(z bool, f func(bool, complex128) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftString(z string, f func(string, complex128) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftInt(z int, f func(int, complex128) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftInt64(z int64, f func(int64, complex128) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftByte(z byte, f func(byte, complex128) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftRune(z rune, f func(rune, complex128) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftFloat32(z float32, f func(float32, complex128) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftFloat64(z float64, f func(float64, complex128) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o Complex128Option) FoldLeftAny(z Any, f func(Any, complex128) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftBool(z bool, f func(bool, Any) bool) bool {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftString(z string, f func(string, Any) string) string {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftInt(z int, f func(int, Any) int) int {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftInt64(z int64, f func(int64, Any) int64) int64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftByte(z byte, f func(byte, Any) byte) byte {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftRune(z rune, f func(rune, Any) rune) rune {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftFloat32(z float32, f func(float32, Any) float32) float32 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftFloat64(z float64, f func(float64, Any) float64) float64 {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}
func (o AnyOption) FoldLeftAny(z Any, f func(Any, Any) Any) Any {
	if o.IsDefined() {
		return f(z, *o.value)
	} else {
		return z
	}
}