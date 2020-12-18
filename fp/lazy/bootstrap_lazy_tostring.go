// bootstrap_lazy_tostring.go
// DO NOT EDIT THIS FILE WAS GENERATED AUTOMATICALLY BY https://github.com/zx80live/gofp-bootstrap

package lazy

import "fmt"

func (n LazyBool) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyBool(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyBool(?)")
	}
}
func (n LazyString) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyString(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyString(?)")
	}
}
func (n LazyInt) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyInt(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyInt(?)")
	}
}
func (n LazyInt64) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyInt64(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyInt64(?)")
	}
}
func (n LazyByte) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyByte(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyByte(?)")
	}
}
func (n LazyRune) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyRune(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyRune(?)")
	}
}
func (n LazyFloat32) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyFloat32(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyFloat32(?)")
	}
}
func (n LazyFloat64) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyFloat64(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyFloat64(?)")
	}
}
func (n LazyAny) ToString() string {
	if n.cached != nil {
		return fmt.Sprintf("LazyAny(%v)", *n.cached)
	} else {
		return fmt.Sprintf("LazyAny(?)")
	}
}
