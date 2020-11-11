package fp

//TODO test
func Eq(e Any) Predicate {
	return func(o Any) bool { return e == o }
}

//TODO test
func Even(e Any) bool {
	return e.(int)%2 == 0
}

//TODO test
func Odd(e Any) bool {
	return e.(int)%2 != 0
}

//TODO test
func Positive(e Any) bool {
	return e.(int) >= 0
}

//TODO test
func Negative(e Any) bool {
	return e.(int) < 0
}
