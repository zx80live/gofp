package fp_test

import (
  "fmt"
  "testing"
  . "github.com/zx80live/gofp/fp"
  . "github.com/zx80live/gofp/gotest"
)


func TestHead(t *testing.T) {
  l := Nil.Cons(3).Cons(2).Cons(1)

  AssertEqual(l.Head(), 1, t)
  AssertEqual(l.Tail().Head(), 2, t)
  AssertEqual(l.Tail().Tail().Head(), 3, t)
  AssertEqual(l.Tail().Tail().Tail().Head(), nil, t)
}

func TestTail(t *testing.T) {
  AssertEqualArrays(Nil.ToArray(), []Any{}, t)

  l := Nil.Cons(3).Cons(2).Cons(1)
  AssertEqualArrays(l.Tail().ToArray(), []Any{2,3}, t)
  AssertEqualArrays(l.Tail().Tail().ToArray(), []Any{3}, t)
  AssertEqualArrays(l.Tail().Tail().Tail().ToArray(), []Any{}, t)
}

func TestToArray(t *testing.T) {
  AssertEqualArrays(Nil.ToArray(), []Any{}, t)
  AssertEqualArrays(Nil.Cons(3).Cons(2).Cons(1).ToArray(), []Any{1,2,3}, t)
}

func TestIsEmpty(t *testing.T) {
  AssertTrue(Nil.IsEmpty(), t)
  AssertFalse(Nil.Cons(1).IsEmpty(), t)
  AssertFalse(Nil.Cons(3).Cons(2).Cons(1).IsEmpty(), t)
}

func TestIsNotEmpty(t *testing.T) {
  AssertFalse(Nil.IsNotEmpty(), t)
  AssertTrue(Nil.Cons(1).IsNotEmpty(), t)
  AssertTrue(Nil.Cons(3).Cons(2).Cons(1).IsNotEmpty(), t)
}

func TestCopy(t *testing.T) {
  l1 := Nil.Cons(3).Cons(2).Cons(1)
  l2 := l1.Copy()

  AssertEqualArrays(l1.ToArray(), l2.ToArray(), t)
}

func TestCons(t *testing.T) {
  AssertEqualArrays(Nil.Cons(3).ToArray(), []Any{3}, t)
  AssertEqualArrays(Nil.Cons(3).Cons(2).ToArray(), []Any{2, 3}, t)
  AssertEqualArrays(Nil.Cons(3).Cons(2).Cons(1).ToArray(), []Any{1, 2, 3}, t)
}

func TestConsArr(t *testing.T) {
  AssertEqualArrays(Nil.ConsArr([]Any{1}).ToArray(), []Any{1}, t)
  AssertEqualArrays(Nil.ConsArr([]Any{1,2}).ToArray(), []Any{2,1}, t)
  AssertEqualArrays(Nil.ConsArr([]Any{1,2,3}).ToArray(), []Any{3,2,1}, t)

  l := Nil.Cons(3).Cons(2).Cons(1)
  AssertEqualArrays(l.ConsArr([]Any{10,20,30}).ToArray(), []Any{30, 20, 10, 1,2,3}, t)
}

func TestEquals(t *testing.T) {
   AssertTrue(Nil.Equals(Nil), t)
   AssertFalse(Nil.Equals(Nil.Cons(3).Cons(2).Cons(1)), t)

   AssertTrue(
     Nil.Cons(3).Cons(2).Cons(1).
     Equals(Nil.Cons(3).Cons(2).Cons(1)), t)

   AssertFalse(
     Nil.Cons(3).Cons(2).Cons(1).Equals(
       Nil.Cons(1).Cons(2).Cons(3),
     ), t)

   fmt.Println("")
}

func TestToString(t *testing.T) {
  AssertEqual(Nil.ToString(), "[]", t)
  AssertEqual(Nil.Cons(1).ToString(), "[1]", t)
  AssertEqual(Nil.Cons(3).Cons(2).Cons(1).ToString(), "[1,2,3]", t)
}

func TestMkString(t *testing.T) {
  AssertEqual(Nil.MkString("(", "|", ")"), "()", t)
  AssertEqual(Nil.Cons(1).MkString("(", "|", ")"), "(1)", t)
  AssertEqual(Nil.Cons(3).Cons(2).Cons(1).MkString("(", "|", ")"), "(1|2|3)", t)
}

func TestForeach(t *testing.T) {
   sum := 0
   Nil.Cons(3).Cons(2).Cons(1).Foreach(func(e Any) {
     sum += e.(int)
   })

   AssertEqual(sum, 6, t)
}

func TestSize(t *testing.T) {
  AssertEqual(Nil.Size(), 0, t)
  AssertEqual(Nil.Cons(1).Size(), 1, t)
  AssertEqual(Nil.Cons(3).Cons(2).Cons(1).Size(), 3, t)
}

func TestFoldLeft(t *testing.T) {
  AssertEqual(
    Nil.FoldLeft(100, func(a Accumulator, e Element) Any {
      return a.(int) + e.(int)
    }), 100, t)


  l := Nil.Cons(5).Cons(4).Cons(3).Cons(2).Cons(1)

  result := l.FoldLeft("init", func(acc Accumulator, e Element) Any{
    return fmt.Sprintf("%v_%v",acc, e)
  })

  AssertEqual(result, "init_1_2_3_4_5", t)
}

func TestReduce(t *testing.T) {
  l := Nil.Cons(5).Cons(4).Cons(3).Cons(2).Cons(1)
  AssertEqual(l.Reduce(func(a Any, b Any) Any {
    return a.(int) + b.(int)
  }), 5+4+3+2+1, t)

  AssertEqual(l.Reduce(func(a Any,b Any) Any {
    return fmt.Sprintf("%v_%v", a, b)
  }), "1_2_3_4_5", t)
}

func TestExists(t *testing.T) {

  eq := func(e Any) Predicate {
    return func(a Any) bool {
      return e == a
    }
  }

  even := func(e Any) bool {
    return e.(int) % 2 == 0
  }

  neg := func (e Any) bool {
    return e.(int) < 0
  }

  AssertFalse(Nil.Exist(eq(10)), t)
  l := Nil.Cons(3).Cons(2).Cons(1)
  AssertTrue(l.Exist(eq(1)), t)
  AssertTrue(l.Exist(eq(2)), t)
  AssertTrue(l.Exist(eq(3)), t)
  AssertTrue(l.Exist(eq(3)), t)
  AssertTrue(l.Exist(even), t)
  AssertFalse(l.Exist(eq(10)), t)
  AssertFalse(l.Exist(neg), t)

}

func TestCount(t *testing.T) {
  eq := func(e Any) Predicate {
    return func(a Any) bool {
      return e == a
    }
  }

  odd := func(e Any) bool {
    return e.(int) % 2 != 0
  }

  neg := func (e Any) bool {
    return e.(int) < 0
  }

  AssertEqual(Nil.Count(odd), 0, t)
  l := Nil.ConsArr([]Any{5,4,-3,2,-1})

  AssertEqual(l.Count(eq(5)), 1, t)
  AssertEqual(l.Count(odd), 3, t)
  AssertEqual(l.Count(neg), 2, t)
}

func TestZipWithIndex(t *testing.T) {
  AssertEqualArrays(Nil.ZipWithIndex().ToArray(), []Any{}, t)

  AssertEqualArrays(Nil.Cons(100).ZipWithIndex().ToArray(), []Any { Tuple2 {100, 0}  }, t)
  l := Nil.Cons(30).Cons(20).Cons(10)

  AssertEqualArrays(l.ZipWithIndex().ToArray(), []Any {
    Tuple2 {10, 0},
    Tuple2 {20, 1},
    Tuple2 {30, 2},
  }, t)
}

func TestZip(t *testing.T) {
  AssertEqualArrays(Nil.Zip(Nil).ToArray(), []Any{}, t)
  
  l1 := Nil.ConsArr([]Any{3,2,1})
  l2 := Nil.ConsArr([]Any{"three","two","one"})

  zipped := l1.Zip(l2)
  expected := []Any {
    Tuple2 { 1, "one" },
    Tuple2 { 2, "two" },
    Tuple2 { 3, "three" },
  }

  AssertEqualArrays(zipped.ToArray(), expected, t)

  l3 := Nil.ConsArr([]Any{"four", "three", "two", "one"})
  zipped2 := l1.Zip(l3)
  expected2 := expected

  AssertEqualArrays(zipped2.ToArray(), expected2, t)
}

func TestReverse(t *testing.T) {
  AssertEqualArrays(Nil.Reverse().ToArray(), []Any{}, t)

  l := Nil.ConsArr([]Any{3,2,1})
  AssertEqualArrays(l.Reverse().ToArray(), []Any{3,2,1}, t)
  AssertEqualArrays(l.Reverse().Reverse().ToArray(), l.ToArray(), t)
}


func TestFind(t *testing.T) {
  eq := func(e Any) Predicate {
    return func(a Any) bool {
      return e == a
    }
  }

  even := func(e Any) bool {
    return e.(int) % 2 == 0
  }

  neg := func (e Any) bool {
    return e.(int) < 0
  }

  AssertEqual(Nil.Find(eq(100)), nil, t)
  l := Nil.ConsArr([]Any{5,4,3,2,1})

  AssertEqual(l.Find(eq(3)), 3, t)
  AssertEqual(l.Find(even),2, t)
  AssertEqual(l.Find(neg),nil, t)
  AssertEqual(l.Find(eq(100)),nil, t)

}

func TestGroupBy(t *testing.T) {
  // grouping list of numbers by identity
  identity := func(e Any) Any { return e }
  l := Nil.ConsArr([]Any{5,4,4,3,5,2,1,3,5})
  res1 := l.GroupBy(identity)

  AssertEqualArrays(res1[1].ToArray(), []Any{1}, t)
  AssertEqualArrays(res1[2].ToArray(), []Any{2}, t)
  AssertEqualArrays(res1[3].ToArray(), []Any{3,3}, t)
  AssertEqualArrays(res1[4].ToArray(), []Any{4,4}, t)
  AssertEqualArrays(res1[5].ToArray(), []Any{5,5,5}, t)

  // grouping list of structs by extractor
  type Shape struct {
    name string
    area int
    color string
  }
 
  l = Nil.ConsArr([]Any {
    Shape { "circle", 100, "green" },
    Shape { "square", 101, "white" },
    Shape { "circle", 200, "yellow" },
    Shape { "triangle", 200, "pink" },
    Shape { "square", 10, "blue" },
    Shape { "triangle", 10, "magenta" },
  })

  byName := l.GroupBy(func(e Any) Any { return e.(Shape).name })
  AssertEqualArrays(byName["circle"].ToArray(), []Any{
    Shape { "circle", 100, "green" },
    Shape { "circle", 200, "yellow" },
  } ,t)
  AssertEqualArrays(byName["square"].ToArray(), []Any{
    Shape { "square", 101, "white" },
    Shape { "square", 10, "blue" },
  } ,t)
  AssertEqualArrays(byName["triangle"].ToArray(), []Any{
    Shape { "triangle", 200, "pink" },
    Shape { "triangle", 10, "magenta" },
  } ,t)


  byArea := l.GroupBy(func(e Any) Any { return e.(Shape).area })
  AssertEqualArrays(byArea[200].ToArray(), []Any{
    Shape { "circle", 200, "yellow" },
    Shape { "triangle", 200, "pink" },
  } ,t)
  AssertEqualArrays(byArea[100].ToArray(), []Any{
    Shape { "circle", 100, "green" },
  } ,t)
  AssertEqualArrays(byArea[101].ToArray(), []Any{
    Shape { "square", 101, "white" },
  } ,t)
  AssertEqualArrays(byArea[10].ToArray(), []Any{
    Shape { "square", 10, "blue" },
    Shape { "triangle", 10, "magenta" },
  } ,t)
}

func TestMap(t *testing.T) {
  mult100 := func(e Any) Any {
    return e.(int) * 100
  }

  decorator := func(e Any) Any {
    return fmt.Sprintf("(%v)", e)
  }

  l := Nil.ConsArr([]Any{3,2,1})
  res := l.Map(mult100).Map(decorator)

  AssertEqualArrays(res.ToArray(), []Any{
    "(100)",
    "(200)",
    "(300)",
  }, t)
}
