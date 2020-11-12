// flist_isempty.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY gofp-bootstrap

package fp

func (l ListBool) IsEmpty() bool             { return l.head == nil && l.tail == nil }
func (l ListBool) IsNotEmpty() bool          { return !l.IsEmpty() }
func (l ListRune) IsEmpty() bool             { return l.head == nil && l.tail == nil }
func (l ListRune) IsNotEmpty() bool          { return !l.IsEmpty() }
func (l ListByte) IsEmpty() bool             { return l.head == nil && l.tail == nil }
func (l ListByte) IsNotEmpty() bool          { return !l.IsEmpty() }
func (l ListInt) IsEmpty() bool              { return l.head == nil && l.tail == nil }
func (l ListInt) IsNotEmpty() bool           { return !l.IsEmpty() }
func (l ListInt8) IsEmpty() bool             { return l.head == nil && l.tail == nil }
func (l ListInt8) IsNotEmpty() bool          { return !l.IsEmpty() }
func (l ListInt16) IsEmpty() bool            { return l.head == nil && l.tail == nil }
func (l ListInt16) IsNotEmpty() bool         { return !l.IsEmpty() }
func (l ListInt32) IsEmpty() bool            { return l.head == nil && l.tail == nil }
func (l ListInt32) IsNotEmpty() bool         { return !l.IsEmpty() }
func (l ListInt64) IsEmpty() bool            { return l.head == nil && l.tail == nil }
func (l ListInt64) IsNotEmpty() bool         { return !l.IsEmpty() }
func (l ListUint) IsEmpty() bool             { return l.head == nil && l.tail == nil }
func (l ListUint) IsNotEmpty() bool          { return !l.IsEmpty() }
func (l ListUint8) IsEmpty() bool            { return l.head == nil && l.tail == nil }
func (l ListUint8) IsNotEmpty() bool         { return !l.IsEmpty() }
func (l ListUint16) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListUint16) IsNotEmpty() bool        { return !l.IsEmpty() }
func (l ListUint32) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListUint32) IsNotEmpty() bool        { return !l.IsEmpty() }
func (l ListUint64) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListUint64) IsNotEmpty() bool        { return !l.IsEmpty() }
func (l ListUintptr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListUintptr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListFloat32) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListFloat32) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListFloat64) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListFloat64) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListComplex64) IsEmpty() bool        { return l.head == nil && l.tail == nil }
func (l ListComplex64) IsNotEmpty() bool     { return !l.IsEmpty() }
func (l ListComplex128) IsEmpty() bool       { return l.head == nil && l.tail == nil }
func (l ListComplex128) IsNotEmpty() bool    { return !l.IsEmpty() }
func (l ListString) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListString) IsNotEmpty() bool        { return !l.IsEmpty() }
func (l ListAny) IsEmpty() bool              { return l.head == nil && l.tail == nil }
func (l ListAny) IsNotEmpty() bool           { return !l.IsEmpty() }
func (l ListBoolArr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListBoolArr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListRuneArr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListRuneArr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListByteArr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListByteArr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListIntArr) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListIntArr) IsNotEmpty() bool        { return !l.IsEmpty() }
func (l ListInt8Arr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListInt8Arr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListInt16Arr) IsEmpty() bool         { return l.head == nil && l.tail == nil }
func (l ListInt16Arr) IsNotEmpty() bool      { return !l.IsEmpty() }
func (l ListInt32Arr) IsEmpty() bool         { return l.head == nil && l.tail == nil }
func (l ListInt32Arr) IsNotEmpty() bool      { return !l.IsEmpty() }
func (l ListInt64Arr) IsEmpty() bool         { return l.head == nil && l.tail == nil }
func (l ListInt64Arr) IsNotEmpty() bool      { return !l.IsEmpty() }
func (l ListUintArr) IsEmpty() bool          { return l.head == nil && l.tail == nil }
func (l ListUintArr) IsNotEmpty() bool       { return !l.IsEmpty() }
func (l ListUint8Arr) IsEmpty() bool         { return l.head == nil && l.tail == nil }
func (l ListUint8Arr) IsNotEmpty() bool      { return !l.IsEmpty() }
func (l ListUint16Arr) IsEmpty() bool        { return l.head == nil && l.tail == nil }
func (l ListUint16Arr) IsNotEmpty() bool     { return !l.IsEmpty() }
func (l ListUint32Arr) IsEmpty() bool        { return l.head == nil && l.tail == nil }
func (l ListUint32Arr) IsNotEmpty() bool     { return !l.IsEmpty() }
func (l ListUint64Arr) IsEmpty() bool        { return l.head == nil && l.tail == nil }
func (l ListUint64Arr) IsNotEmpty() bool     { return !l.IsEmpty() }
func (l ListUintptrArr) IsEmpty() bool       { return l.head == nil && l.tail == nil }
func (l ListUintptrArr) IsNotEmpty() bool    { return !l.IsEmpty() }
func (l ListFloat32Arr) IsEmpty() bool       { return l.head == nil && l.tail == nil }
func (l ListFloat32Arr) IsNotEmpty() bool    { return !l.IsEmpty() }
func (l ListFloat64Arr) IsEmpty() bool       { return l.head == nil && l.tail == nil }
func (l ListFloat64Arr) IsNotEmpty() bool    { return !l.IsEmpty() }
func (l ListComplex64Arr) IsEmpty() bool     { return l.head == nil && l.tail == nil }
func (l ListComplex64Arr) IsNotEmpty() bool  { return !l.IsEmpty() }
func (l ListComplex128Arr) IsEmpty() bool    { return l.head == nil && l.tail == nil }
func (l ListComplex128Arr) IsNotEmpty() bool { return !l.IsEmpty() }
func (l ListStringArr) IsEmpty() bool        { return l.head == nil && l.tail == nil }
func (l ListStringArr) IsNotEmpty() bool     { return !l.IsEmpty() }
func (l ListAnyArr) IsEmpty() bool           { return l.head == nil && l.tail == nil }
func (l ListAnyArr) IsNotEmpty() bool        { return !l.IsEmpty() }
