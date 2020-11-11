package fp

import "regexp"

func Eq(e Any) Predicate {
	return func(o Any) bool { return e == o }
}

func Even(e Any) bool {
	return e.(int)%2 == 0
}

func Odd(e Any) bool {
	return e.(int)%2 != 0
}

func Positive(e Any) bool {
	return e.(int) >= 0
}

func Negative(e Any) bool {
	return e.(int) < 0
}

func RegexpString(pattern string) Predicate {
	return Regexp(regexp.MustCompile(pattern))
}

func Regexp(r *regexp.Regexp) Predicate {
	return func(e Any) bool {
		return r.MatchString(e.(string))
	}
}
