package fp_test

import (
	. "github.com/zx80live/gofp/fp"
	. "github.com/zx80live/gofp/gotest"
	"regexp"
	"testing"
)

func TestRegexpStringGroups(t *testing.T) {
	str := "{'shape': 'circle', 'area': 100, 'color': 'green'}"
	f := RegexpStringGroups("{('(\\w+)': '(\\w+)'), ('(\\w+)': (\\d+)), ('(\\w+)': '(\\w+)')}")

	groups := f(str).([]string)
	AssertEqual(len(groups), 10, t)
	AssertEqual(groups[0], str, t)
	AssertEqual(groups[1], "'shape': 'circle'", t)
	AssertEqual(groups[2], "shape", t)
	AssertEqual(groups[3], "circle", t)
	AssertEqual(groups[4], "'area': 100", t)
	AssertEqual(groups[5], "area", t)
	AssertEqual(groups[6], "100", t)
	AssertEqual(groups[7], "'color': 'green'", t)
	AssertEqual(groups[8], "color", t)
	AssertEqual(groups[9], "green", t)
}

func TestRegexpGroups(t *testing.T) {
	str := "{'shape': 'circle', 'area': 100, 'color': 'green'}"
	f := RegexpGroups(regexp.MustCompile("{('(\\w+)': '(\\w+)'), ('(\\w+)': (\\d+)), ('(\\w+)': '(\\w+)')}"))

	groups := f(str).([]string)
	AssertEqual(len(groups), 10, t)
	AssertEqual(groups[0], str, t)
	AssertEqual(groups[1], "'shape': 'circle'", t)
	AssertEqual(groups[2], "shape", t)
	AssertEqual(groups[3], "circle", t)
	AssertEqual(groups[4], "'area': 100", t)
	AssertEqual(groups[5], "area", t)
	AssertEqual(groups[6], "100", t)
	AssertEqual(groups[7], "'color': 'green'", t)
	AssertEqual(groups[8], "color", t)
	AssertEqual(groups[9], "green", t)
}
