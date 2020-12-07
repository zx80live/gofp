# Functional programming in GO


## Introduction

This library was inspired by Scala (collection API, functional paradigm and, concurrence API etc). 

## Table of contents

- [Getting started](#getting-started)<a name="toc0"></a>
- [Collection API](#collection-api)
  * [List](#list)<a name="list0"></a>
    + [List structure](#list-structure)
    + [Create list](#create-list)
    + [List.Copy](#listcopy)
    + [List.Count](#listcount)
    + [List.Drops](#listdrops)
    + [List.Equals](#listequals)
    + [List.Filter](#listfilter)
    + [List.Find](#listfind)
    + [List.FlatMap](#listflatmap)
    + [List.Flatten](#listflatten)
    + [List.FoldLeft](#listfoldleft)
    + [List.Foreach](#listforeach)
    + [List.GroupBy](#listgroupby)
    + [List.Heads and tails](#listheads-and-tails)
    + [List.Empty NonEmpty](#listempty-nonempty)
    + [List.Map](#listmap)
    + [List.MkString](#listmkstring)
    + [List.Nil](#listnil)
    + [List.Cons](#listcons)
    + [List.Reduce](#listreduce)
    + [List.Reverse](#listreverse)
    + [List.Size](#listsize)
    + [List.Takes](#listtakes)
    + [List.ToArray](#listtoarray)
    + [List.ToString](#listtostring)
    + [Supported list types](#supported-list-types)
  * [Option](#option)<a name="option0"></a>
    + [Create option](#create-option)
    + [Option.Equals](#optionequals)
    + [Option.Filter](#optionfilter)
    + [Option.FlatMap](#optionflatmap)
    + [Option.Flatten](#optionflatten)
    + [Option.FoldLeft](#optionfoldleft)
    + [Option.Foreach](#optionforeach)
    + [Option.IsDefined](#optionisdefined)
    + [Option.IsEmpty](#optionisempty)
    + [Option.Map](#optionmap)
    + [Option.None](#optionnone)
    + [Option.ToString](#optiontostring)
    + [Supported option types](#supported-option-types)
  * [Array](#array)<a name="array0"></a>
    + [Create array](#create-array)
    + [Array.Count](#arraycount)
    + [Array.Drops](#arraydrops)
    + [Array.Equals](#arrayequals)
    + [Array.Filter](#arrayfilter)
    + [Array.Find](#arrayfind)
    + [Array.Foreach](#arrayforeach)
    + [Array.Heads and tails](#arrayheads-and-tails)
    + [Array.Map](#arraymap)
    + [Array.MkString](#arraymkstring)
    + [Array.Takes](#arraytakes)
    + [Array.ToList](#arraytolist)
    + [Array.ToString](#arraytostring)
    + [Supported array types](#supported-array-types)
- [Boxed types](#boxed-types)<a name="boxed0"></a>
    + [Underlined](#underlined)
    + [Converters](#converters)
    + [List constructors](#list-constructors)
    + [Math and logic operations](#math-and-logic-operations)
    + [Supported boxed types](#supported-boxed-types)
- [Predicates](#predicates)<a name="predicates0"></a>
    + [Empty predicates](#empty-predicates)
    + [Numeric predicates](#numeric-predicates)
    + [String predicates](#string-predicates)
    + [Predicates composition](#predicates-composition)
    + [Supported predicate types](#supported-predicate-types)
- [Transformers](#transformers)<a name="transformers0"></a>
    + [Identity](#identity)
    + [String transformers](#string-transformers)
- [Predef](#predef)<a name="predef0"></a>
    + [Require](#require)
- [Concurrency API](#concurrency-api)<a name="concurrency0"></a>
  * [Future](#future)<a name="future0"></a>
    + [Create future](#create-future)  
    + [Future.FlatMap](#futureflatmap)
    + [Future.FlatMap Blocking composition](#futureflatmapblocking-composition)
    + [Future.FlatMap NON-blocking composition](#futureflatmapnon-blocking-composition)
    + [Future.Map](#futuremap)
    + [Future.Result](#futureresult)
    + [Future.Success](#futuresuccess)
    + [Supported future types](#supported-future-types)
- [Bootstrap](#bootstrap)<a name="bootstrap0"></a>
- [License](#license)<a name="license0"></a>


## Getting started

Before using functional library import the following package in your project:

```go
import . "github.com/zx80live/gofp/fp"
```

Or do the following steps if you want to create an empty project:

1. Create new go-module project in your `$WORK_DIR` directory
```bash
# create your project dir
cd $WORK_DIR
mkdir test-gofp
cd test-gofp

# create your go-module
go mod init example.com/username/test-gofp

# create file in your preferred editor, for example
vim test.go       
```

2. Write the following code in the `$WORK_DIR/test-gofp/test.go` file:
```go
package main

import (
	"fmt"
    . "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello functional programming in GO!")
	
	l := MakeIntList(1,2,3,4,5)
	res1 := l.Filter(EvenInt)
	
	fmt.Println(res1.ToString())
}
```

3. Execute file:
```bash
go run test.go
```
The result of above command should be:
```
Hello functional programming in GO!
List(2,4)
```

4. Explore this library by the examples which are presented in that [documentation](#toc0). Good luck!

## Collection API

Current library supports the following collection types: `Arrays`, `Lists` and `Options`. Each type is monad and supports functions such as `Map`, `FlatMap`, `Filter` and etc.

[[🠕]](#list0)

### List

All examples of  this section use a `IntList` type. But this API is supported in the other array types also. (See [Supported list types](#supported-list-types) section)
[[🠕]](#list0)

#### List structure

List is recursive functional data structure which has `head` and `tail` as sublist. For example, an `IntList(1,2,3)` will be presented in the following structure: 

```go
   1 ──▶ ( 2 ──▶ ( 3 ──▶ Nil ) )
  └─┘    └─────────────────────┘          
   ↑               ↑
  head            tail
```

List is implemented in the following structure:

```go
type IntList struct {
  head *int
  tail *IntList
}
```

[[🠕]](#list0)

#### Create list

```go
// Create list
// O(n)
func MakeIntList(elements ...int) IntList

// Construct list by prepend head
// O(1)
func (l IntList) Cons(e int) IntList
```

Example:

```go
// create list by factory
var l IntList = MakeIntList(1,2,3,4,5)

// create list from Nil
l2 := NilInt.Cons(5).Cons(4).Cons(3).Cons(2).Cons(1)  // equals IntList(1,2,3,4,5)

// create list from boxed (rich) types
l3 := Int(10).Cons(20).Cons(30)               // List(30,20,10)
l4 := String("one").Cons("two").Cons("three") // List("three", "two", "one")

```

[[🠕]](#list0)



#### List.Copy

```go
// Copy references of elements to new list
// O(2*n)
func (l IntList) Copy() IntList
```

Example:

```go
l1 := MakeIntList(1,2,3,4,5)
l2 := l1.Copy()
```

[[🠕]](#list0)

#### List.Count

```go
// Count elements which satisfy a predicate
// O(1..n)
func (l IntList) Count(predicate func(int) bool)
```

Example:

```go
l := MakeIntList(1,2,3,4,5)
l.Count(func(e int) bool { return e % 2 == 0})    // 2
l.Count(EvenInt)                                  // 2
```

[[🠕]](#list0)



#### List.Drops

```go
// Returns new list without n-first elements
// O(1..n)
func (l IntList) Drop(n int) IntList

// Returns new list without n-last elements
// O(1..n)
func (l IntList) DropRight(n int) IntList

// Returns new list without first elements which statisfy a predicate
// O(1..n)
func (l IntList) DropWhile(predicate func(int) bool) IntList
```

Example:

```go
l := MakeIntList(10,20,30,40,50)
res1 := l.Drop(2)                                 // IntList(30,40,50)
res2 := l.DropRight(2)                            // IntList(10,20,30)
res3 := l.DropWhile(func (e int) bool { e < 40 }) // IntList(40,50)                      
```

[[🠕]](#list0)



#### List.Empty NonEmpty

```go
// Returns true if list is empty
// O(n)
func (l IntList) IsEmpty() bool

// Returns true if list is not empty
// O(n)
func (l IntList) NonEmpty() bool
```

Example:

```go
l1 := MakeIntList(1,2,3)
res1 := l1.IsEmpty()      // false
res2 := l1.NoEmpty()      // true

l2 := MakeIntList()    
res3 := l2.IsEmpty()      // true
res4 := l2.NonEmpty()     // fase

res5 := NilInt.IsEmpty()  // true
res6 := NilInt.NonEmpty() // false
```

[[🠕]](#list0)



#### List.Equals

```go
// Returns true if both lists are equal
// O(n)
func (a IntLIst) Equals(b IntList) bool
```

Example:

```go
// compare the simple lists
l1 := MakeIntList(10,20,30,40,50)
l2 := MakeIntList(10,20,30,40,50)
l3 := MakeIntList({10,20,30)

l1.Equals(l2)   // true
l1.Equals(l3)   // false
                   
// compare the nested lists
l4 := MakeIntListList(MakeIntList(1,2), MakeIntList(3,4,5))           
l5 := MakeIntListList(MakeIntList(1,2), MakeIntList(3,4,5)) 
l6 := MakeIntListList(MakeIntList(1,2,3), MakeIntList(4,5))           
                 
l4.Equals(l5)  // true
l4.Equals(l6)  // false
```

[[🠕]](#list0)



#### List.Filter

```go
// Returns new list with elements which statisfy a predicate
// O(n)
func (l IntList) Filter(predicate func(int) bool) IntList
```

Example:

```go
l := MakeIntList(1,-2,-3,4,5,-6,7,8,-9,10)
res := l.
         Filter(func (e int) bool { return e % 2 == 0}).   // filter even numbers
         Filter(func (e int) bool { return e >= 0 })       // filter positive numbers

fmt.Println(res.ToString())    // IntList(4, 8, 10)

res2 := l.Filter(EvenInt).Filter(PosInt)                   // use library predicates O(2*n)
res3 := l.Filter(EvenInt.And(PosInt))                      // compose predicates     O(n)
```

[[🠕]](#list0)



#### List.Find

```go
// Returns first element which statisfy a predicate
// O(1..n)
func (l IntList) Find(predicate func(int) bool)
```

Example:

```go
l := IntList(1,2,3,4,5,6)

var res1 IntOption = l.Find(func(e int) bool { return e == 3 }) // Some(3)
var res2 IntOption = l.Find(EvenInt)                            // Some(2)
var res3 IntOption = l.Find(NegInt)                             // None

```

[[🠕]](#list0)



#### List.FlatMap

```go
// FlatMap gives the ability to chain operations together
func (l IntList) FlatMap<Type>(f func(int) <Type>List) <Type>List

// examples
func (l IntList) FlatMapInt(f func(int) IntList) IntList
func (l IntList) FlatMapString(f func(int) StringList) StringList
...
```

Example:

```go
// Example: flatten nested list
// List(List(1,2), Nil, List(3,4,5), List(6,7))
l := MakeIntListList(MakeIntList(1,2), NilInt, MakeIntList(3,4,5), MakeIntList(6,7)) 
l.FlatMapInt(func(e IntList) IntList { return e })        // List(1,2,3,4,5,6,7)
l.FlatMapInt(func(e IntList) IntList { 
    return e.MapInt(func(e int) int) { return e * 10 }})  // List(10,20,30,40,50,60,70)
```

[[🠕]](#list0)



#### List.Flatten

```go
// Collapse the nested elements of a collection to flat collection
func (l <Type>ListList) Flatten() <Type>List
func (l <Type>OptionList) Flatten() <Type>List

// examples
func (l IntIntList) Flatten() IntList
func (l IntOptionList) Flatten() IntList
```

Example:

```go
// List(List(1,2,3), Nil, List(4,5), List(6,7))
l := MakeIntListList(MakeIntList(1,2,3), Nil, MakeIntList(4,5), MakeIntList(6, 7))
var res IntList = l.Flatten()   // List(1,2,3,4,5,6,7)
```

[[🠕]](#list0)



#### List.FoldLeft

```go
// Applies binary function to each element of list going left to right
// O(n)
// z - initial value
func (l IntList) FoldLeft<Type>(z <Type>, func(<Type>, int) <Type>) <Type>

// examples
func (l IntList) FoldLeftInt(z Int, func(int, int) int) int
func (l IntList) FoldLeftString(z string, func(string, int) string) string

```

Example:

```go
l := MakeIntList(1,2,3,4,5)


sum := l.FoldLeftInt(
          0, 
          func(acc, el) int { return acc + el})   // sum = (((((0 + 1) + 2) + 3) + 4) + 5)

str := l.FoldLeftString(
          ">", 
          func(acc, el) string { return fmt.Sprintf("%v|%v", acc, el)}) // >1|2|3|4|5
```

[[🠕]](#list0)



#### List.Foreach

```go
// Iterate over elements
// O(n)
func (l IntList) Foreach(func(int))
```

Example:

```go
l := MakeIntList(1,2,3)
f.Foreach(func(e int) {
    fmt.Println("> ", e)
})
```

[[🠕]](#list0)



#### List.GroupBy

```go
// Partitions this list into a map of according to some discriminator function
func (l <A>List) GroupBy<K>(func(<A>)<K>) map[<K>]<A>List

// examples
func (l IntList) GroupByInt(func(int) int) map[int]IntList
func (l AnyList) GroupByAny(func(Any)Any) map[Any]AnyList
```

Example: group by identity

```go
l1 := MakeIntList(1,2,1,1,3,2)
res1 := l1.GroupByInt(func(e int) int { return e }) /* Map(1 -> List(1,1,1)
                                                           2 -> List(2,2)
                                                           3 -> List(3))  */
res2 := l1.GroutByInt(IntIdentity)                  // the same as res1
```

Example: group by different attributes

```go
type Shape struct {
  name  string
  color string
  area  int
}

shapes := NilAny.
  Cons(Shape{"circle", "green", 10}).
  Cons(Shape{"triangle", "yellow", 30}).
  Cons(Shape{"polygon", "green", 10}).
  Cons(Shape{"triangle", "green", 10}).
  Cons(Shape{"circle", "red", 20}).  
  Cons(Shape{"polygon", "red", 20}).
  Cons(Shape{"circle", "yellow", 30}).
  Cons(Shape{"triangle", "red", 20}).
  Cons(Shape{"polygon", "yellow", 30})

// Group shapes by name
byShape := func(e Any) Any { return e.(Shape).name }
res1 := shapes.GroupByAny(byShape)
/*
 Map(
   "polygon"  -> List({"polygon" "green" 10}, {"polygon" "red" 20}, {"polygon" "yellow" 30})
   "triangle" -> List({"triangle" "green" 10}, {"triangle" "red" 20}, {"triangle" "yellow" 30})
   "circle"   -> List({"circle" "green" 10}, {"circle" "red" 20}, {"circle" "yellow" 30})
 )
 */

// Group shapes by color
byColor := func(e Any) Any { return e.(Shape).color }
res2 := shapes.GroupByAny(byColor)
/*
 Map(
   "yellow" -> List({"circle" "yellow" 30}, {"triangle" "yellow" 30}, {"polygon" "yellow" 30})
   "red"    -> List({"circle" "red" 20}, {"triangle" "red" 20}, {"polygon" "red" 20})
   "green"  -> List({"circle" "green" 10}, {"triangle" "green" 10}, {"polygon" "green" 10})
 )
*/

byArea := func(e Any) Any { return e.(Shape).area }
res3 := shapes.GroupByAny(byArea)
/*
 Map (
   30 -> List({"circle" "yellow" 30}, {"triangle" "yellow" 30}, {"polygon" "yellow" 30})
   20 -> List({"circle" "red" 20}, {"triangle" "red" 20}, {"polygon" "red" 20})
   10 -> List({"circle" "green" 10}, {"triangle" "green" 10}, {"polygon" "green" 10})
 )
*/
```
[[🠕]](#list0)



#### List.Heads and tails

```go
// Returns head of list or throws exeption if list is empty
// O(1)
func (l IntList) Head() int

// Returns optional head from list
// O(1)
func (l IntList) HeadOption() IntOption

// Returns tail as list without first element
// O(1)
func (l IntList) Tail() IntList
```

Example:

```go
l := MakeIntList(1,2,3)
h1 := l.Head()                      // 1
t1 := l.Tail()                      // List(2,3)
h1Opt := l.HeadOption()             // Some(1)
h2Opt := l.Tail().Tail().Tail()     // NoneInt
l.Tail().Tail().Tail().Head()       // panic("there is no heads")
```

[[🠕]](#list0)




#### List.Map

```go
// Transform each element of list to other type
// O(n)
func (l IntList) Map<Type>(func (int) <Type>) <Type>List

// examples
func(l IntList) MapInt(func (int) int) IntList
func(l IntList) MapString(func (int) string) StringList

```

Example:

```go
l := MakeIntList(1,2,3)
r1 := l.MapInt(func(e int) int { return e * 10})                        // List(10,20,30)
r2 := l.MapString(func (e int) string { return fmt.Sprintf("<%v>", e)}) // List("<1>","<2>","<3>")
```

[[🠕]](#list0)



#### List.MkString

```go
// Make string representation of that list with decorated elements and separator
// O(n)
func (l IntList) MkString(start, sep, end string) string
```

Example:

```go
l := MakeIntList(1,2,3)
str := l.MkString("<", "|", ">")      // "<1|2|3>"
```

[[🠕]](#list0)



#### List.Nil

```go
// Empty list, can be used as initial tail before creating
var Nil<Type> <Type>List

// examples
var NilInt IntList = ...
var NilString StringList = ...
...
```

Example:

```go
// create list from empty tail
l := NilInt.Cons(3).Cons(2).Cons(1)   // List(1,2,3)
```

[[🠕]](#list0)



#### List.Cons

```go
// Create new list with new head and tail as current list
// O(1)
func (l IntList) Cons(e int) IntList
```

Example:

```go
// create list from empty tail
l1 := NilInt.Cons(3).Cons(2).Cons(1)   // List(1,2,3)

l2 := MakeIntList(1,2,3)
l3 := l2.Cons(4)                       // List(4,1,2,3)
```

[[🠕]](#list0)



#### List.Reduce

```go
// Reduces elements of list using associative binary operator
// O(n)
func (l IntList) Reduce(func (int, int) int) int
```

Example:

```go
l := MakeIntList(1,2,3)
sum := l.Reduce(func(acc, el int) int { return acc + el })  // sum = 1 + 2 + 3
```

[[🠕]](#list0)



#### List.Reverse

```go
// Reverse order of list elements
// O(n)
func (l IntList) Reverse() IntList
```

Example:

```go
l1 := MakeIntList(1,2,3)
l2 := l1.Reverse()                       // List(3,2,1)
```

[[🠕]](#list0)



#### List.Size

```go
// Count elements of list
// O(n)
func (l IntList) Size() int
```

Example:

```go
l1 := MakeIntList(1,2,3)
l2 := MakeIntList()

res1 := l1.Size()         // 3
res2 := l2.Size()         // 0
res3 := NilInt            // 0
```

[[🠕]](#list0)



#### List.Takes

```go
// Returns first n-elements
// O(1..n)
func (l IntList) Take(n int) IntList

// Returns last n-elements
// O(1..n)
func (l IntList) TakeRight(n int) IntList

// Returns first elements which statisfy a predicate
// O(1..n)
func (l IntList) TakeWhile(func(int) bool) IntList
```

Example:

```go
l := IntList(10,20,30,40,50)
res1 := l.Take(2)                                 // List(10,20)
res2 := l.TakeRight(2)                            // List(40,50)
res3 := l.TakeWhile(func (e int) bool { e < 40})  // List(10,20,30)
```

[[🠕]](#list0)



#### List.ToArray

```go
// Returns elements of list as array
// O(n)
func (l IntList) ToArray() []int
```

Example:

```go
l := MakeIntList(1,2,3)
a := l.ToArray()          // []int{1,2,3}
```

[[🠕]](#list0)



#### List.ToString

```go
// Transform list to string
func (l IntList) ToString() string
```

Example:

```go
l1 := MakeIntList(1,2,3)
res1 := l1.ToString()                                              // "List(1,2,3)"
l2 := MakeIntListList(
        MakeIntList(1,2), 
        NilInt, 
        MakeIntList(3,4))  // "List(List(1,2), List(), List(3,4))"
```

[[🠕]](#list0)





#### Supported list types

<details><summary>Supported list types (click to expand)</summary>

| Array type        | Scala analogue        | Go analogue |
| ----------------- | --------------------- | ----------- |
| BoolList          | List[Boolean]         |             |
| StringList        | List[String]          |             |
| IntList           | List[Int]             |             |
| Int64List         | List[Long]            |             |
| ByteList          | List[Byte]            |             |
| RuneList          | List[Char]            |             |
| Float32List       | List[Float]           |             |
| Float64List       | List[Double]          |             |
| AnyList           | List[Any]             |             |
| BoolArrayList     | List[Array[Bool]]     |             |
| StringArrayList   | List[Array[String]]   |             |
| IntArrayList      | List[Array[Int]]      |             |
| Int64ArrayList    | List[Array[Long]]     |             |
| ByteArrayList     | List[Array[Byte]]     |             |
| RuneArrayList     | List[Array[Char]]     |             |
| Float32ArrayList  | List[Array[Float]]    |             |
| Float64ArrayList  | List[Array[Double]]   |             |
| AnyArrayList      | List[Array[Any]]      |             |
| BoolOptionList    | List[Option[Boolean]] |             |
| StringOptionList  | List[Option[String]]  |             |
| IntOptionList     | List[Option[Int]]     |             |
| Int64OptionList   | List[Option[Long]]    |             |
| ByteOptionList    | List[Option[Byte]]    |             |
| RuneOptionList    | List[Option[Char]]    |             |
| Float32OptionList | List[Option[Float]]   |             |
| Float64OptionList | List[Option[Double]]  |             |
| AnyOptionList     | List[Option[Any]]     |             |
| BoolListList      | List[List[Boolean]]   |             |
| StringListList    | List[List[String]]    |             |
| IntListList       | List[List[Int]]       |             |
| Int64ListList     | List[List[Long]]      |             |
| ByteListList      | List[List[Byte]]      |             |
| RuneListList      | List[List[Char]]      |             |
| Float32ListList   | List[List[Float]]     |             |
| Float64ListList   | List[List[Double]]    |             |
| AnyListList       | List[List[Any]]       |             |
</details>

[[🠕]](#list0)



### Option

All examples of  this section use a `IntOption` type. But this API is supported in the other array types also. (See [Supported option types](#supported-option-types) section)

#### Create option
```go
func MakeIntOption(e int) IntOption
func IntOpt(e int) IntOption
```
Example:
```go
o1 := MakeIntOption(10)     // Some(10)
o2 := IntOpt(20)            // Some(20)
```
[[🠕]](#option0)

#### Option.Equals
```go
// Returns true if both options are equal
func (a IntOption) Equals(b IntOption) bool
```
Example:
```go
o1 := IntOpt(10)
o2 := IntOpt(20)
o3 := IntOpt(10)

o1.Equals(o2)            // false
o1.Equals(o3)            // true
o1.Equals(NoneInt)       // false
NoneInt.Equals(NoneInt)  // true
```
[[🠕]](#option0)


#### Option.Filter
```go
// Filter content of option by predicate
func (o IntOption) Filter(predicate func(int)bool) IntOption
```
Example:
```go
o := IntOption(10)
res1 := o.Filter(func(e int) bool { return e % 2 == 0})  // Some(10)
res2 := o.Filter(EvenInt)                                // Some(10)
res3 := o.Filter(func(e int) bool { return e == 20 })    // None
res4 := o.Filter(NegInt)                                 // None
```
[[🠕]](#option0)


#### Option.FlatMap
```go
// FlatMap gives the ability to chain operations together
func (o IntOption) FlatMap<Type>(f func(int) <Type>Option) <Type>Option

// examples
func (o IntOption) FlatMapInt(f func(int) IntOption) IntOption
func (o IntOption) FlatMapString(f func(int) StringOption) StringOption
...
```
Example:
```go
// Task: calculate the sum of elements
//       if all elements are defined then the sum will be defined too
//       if one or more elements are not defined then the sum will not be defined too

// implementation of the task
sumFunc := func(a, b, c IntOption) IntOption {
    sum := a.FlatMapInt(func (av int) IntOption {
        return b.FlatMapInt(func bv int) IntOption {
            return c.MapInt(func cv int) IntOption {
                return av + bv + cv
            } 
        }
    })
    return sum
}

// Case when a, b, c are defined
a := IntOpt(10)
b := IntOpt(20)
c := IntOpt(30)

sumFunc(a, b, c)                                          // Some(60)


// Cases when one or more of a, b, c are not defined
sumFunc(IntOpt(10), NoneInt, IntOpt(30))                  // None
sumFunc(IntOpt(10), IntOpt(20), NoneInt)                  // None
sumFunc(NoneInt, IntOpt(20), IntOpt(30))                  // None
sumFunc(NoneInt, NoneInt, IntOpt(30))                     // None
...
```
[[🠕]](#option0)


#### Option.Flatten
```go 
// Collapse the nested option to flat option
func (o IntIntOption) IntOption
```
Example:
```go
o := MakeIntOptionOption(MakeIntOption(10))    // Some(Some(10))
flatten := o.Flatten()                         // Some(10)
```
[[🠕]](#option0)


#### Option.FoldLeft
```go
// Applies binary function to element of option
// z - initial value
func (l IntOption) FoldLeft<Type>(z <Type>, func(<Type>, int) <Type>) <Type>

// examples
func (l IntOption) FoldLeftInt(z Int, func(int, int) int) int
func (l IntOption) FoldLeftString(z string, func(string, int) string) string
...
```
Example:
```go
o := IntOpt(10)
res1 := l.FoldLeftInt(
            0, 
            func(acc, el) int { return acc + el})   // sum = 0 + 10
```
[[🠕]](#option0)


#### Option.Foreach
```go
// Iterate over content
func (l IntOption) Foreach(func(int))
```
Example:
```go
l := IntOpt(10)
f.Foreach(func(e int) {
    fmt.Println("> ", e)
})
```
[[🠕]](#option0)


#### Option.IsDefined
```go
// Returns true if option contains a value
func (o IntOption) IsDefined() bool
```
Example:
```go
IntOption(10).IsDefined()     // true
NoneInt.IsDefined()           // false
```
[[🠕]](#option0)


#### Option.IsEmpty
```go
// Returns true if option is not defined
func (o IntOption) IsEmpty() bool
```
Example:
```go
NoneInt.IsEmpty()             // true
IntOpt(10).IsEmpty()          // false
```
[[🠕]](#option0)


#### Option.Map
```go
// Transform element of option to other type
func (l IntOption) Map<Type>(func (int) <Type>) <Type>Option

// examples
func(l IntOption) MapInt(func (int) int) IntOption
func(l IntOption) MapString(func (int) string) StringOption
```
Example:
```go
o := IntOpt(10)
res1 := o.MapInt(func (e int) int { return e * 10 })                       // Some(100)
res2 := o.MapString(func (e int) string { return fmt.Sprintf("<%v>", e) }) // Some("<10>") 
```
[[🠕]](#option0)

#### Option.None

There are empty option constants for each supported optional type

```go
var NoneInt IntOption = ...
var NoneString IntOption = ...
...
```
Example:
```go
o1 := NoneInt            // None
o2 := NoneString         // None

func parseIp(str string) StringOption {
    if str != "" { 
        return StringOpt(str) 
    } else {
        return NoneString
    } 
}

ip1 := parseIp("")           // None
ip2 := parseIp("127.0.0.1")  // Some("127.0.0.1")
```
[[🠕]](#option0)


#### Option.ToString
```go
// Transform option to string
func (o IntOption) string
```
Example:
```go
var o1 IntOption = IntOption(10)
var o2 IntOption = NoneInt

var str1 string = o1.ToString()      // "Some(10)"
var str2 string = o2.TOString()      // "None"
```
[[🠕]](#option0)

#### Supported option types
<details><summary>Supported option types (click to expand)</summary>

| Option type         | Scala analogue          |
| ------------------- | ----------------------- |
| BoolOption          | Option[Boolean]         |
| StringOption        | Option[String]          |
| IntOption           | Option[Int]             |
| Int64Option         | Option[Long]            |
| ByteOption          | Option[Byte]            |
| RuneOption          | Option[Char]            |
| Float32Option       | Option[Float]           |
| Float64Option       | Option[Double]          |
| AnyOption           | Option[Any]             |
| BoolOptionOption    | Option[Option[Boolean]] |
| StringOptionOption  | Option[Option[String]]  |
| IntOptionOption     | Option[Option[Int]]     |
| Int64OptionOption   | Option[Option[Long]]    |
| ByteOptionOption    | Option[Option[Byte]]    |
| RuneOptionOption    | Option[Option[Char]]    |
| Float32OptionOption | Option[Option[Float]]   |
| AnyOptionOption     | Option[Option[Double]]  |
| BoolArrayOption     | Option[Array[Boolean]]  |
| StringArrayOption   | Option[Array[String]]   |
| IntArrayOption      | Option[Array[Int]]      |
| Int64ArrayOption    | Option[Array[Long]]     |
| ByteArrayOption     | Option[Array[Byte]]     |
| RuneArrayOption     | Option[Array[Char]]     |
| Float32ArrayOption  | Option[Array[Float]]    |
| Float64ArrayOption  | Option[Array[Double]]   |
| AnyArrayOption      | Option[Array[Any]]      |
| BoolListOption      | Option[List[Boolean]]   |
| StringListOption    | Option[List[String]]    |
| IntListOption       | Option[List[Int]]       |
| Int64ListOption     | Option[List[Long]]      |
| ByteListOption      | Option[List[Byte]]      |
| RuneListOption      | Option[List[Char]]      |
| Float32ListOption   | Option[List[Float]]     |
| Float64ListOption   | Option[List[Double]]    |
| AnyListOption       | Option[List[Any]]       |
</details>

[[🠕]](#option0)


### Array

Array is just wrapper for go-arrays which contains convenient functions. All examples of  this section use a `IntArray` type. But this API is supported in the other array types also. (See [Supported array types](#supported-array-types) section)
Warning, current implementation of the array wrapper is mutable data structure.



#### Create array
```go
type IntArray []int
```
Example:
```go
arr := IntArray([]int{10, 20, 30})
```
[[🠕]](#array0)


#### Array.Count

```go
// Returns count of elements which satisfy a predicate
// O(n)
func (a IntArray) Count(predicate func(int) bool) int
```
Example:
```go
arr := IntArray([]int{1,2,3,4,5})
arr.Count(func (e int) bool { return e % 2 == 0})         // 2
arr.Count(EvenInt)                                        // 2
arr.Count(PosInt)                                         // 5
arr.Count(NegInt)                                         // 0
```
[[🠕]](#array0)


#### Array.Drops

```go
// Returns new array without n-first elements
// O(1..n)
func (a IntArray) Drop(n int) IntArray

// Returns new array without n-last elements
// O(1..n)
func (a IntArray) DropRight(n int) IntArray

// Returns new array without first elements which statisfy a predicate
// O(1..n)
func (a IntArray) DropWhile(predicate func(int) bool) IntArray
```

Example:

```go
arr := IntArray([]int{10,20,30,40,50})
res1 := arr.Drop(2)                                 // Array(30,40,50)
res2 := arr.DropRight(2)                            // Array(10,20,30)
res3 := arr.DropWhile(func (e int) bool { e < 40 }) // Array(40,50)
```
[[🠕]](#array0)


#### Array.Equals

```go
// Returns true if both arrays are equal
// O(n)
func (a IntArray) Equals(b IntArray) bool
```

Example:

```go
arr1 := IntArray([]int{10,20,30,40,50})
arr2 := IntArray([]int{10,20,30,40,50})
arr3 := IntArray([]int{10,20,30})

arr1.Equals(arr2)   // true
arr1.Equals(arr3)   // false
```
[[🠕]](#array0)



#### Array.Filter

```go
// Returns new array with elements which statisfy a predicate
// O(n)
func (a IntArray) Filter(predicate func(int) bool) IntArray
```

Example:

```go
arr := IntArray([]int {1,-2,-3,4,5,-6,7,8,-9,10})
res := arr.
         Filter(func (e int) bool { return e % 2 == 0}).   // filter even numbers
         Filter(func (e int) bool { return e >= 0 })       // filter positive numbers

fmt.Println(res.ToString())    // [4, 8, 10]

res2 := arr.Filter(EvenInt).Filter(PosInt)                 // use library predicates
res3 := arr.Filter(EvenInt.And(PosInt))                    // compose predicates

```
[[🠕]](#array0)



#### Array.Find

```go
// Returns first element which statisfy a predicate
// O(1..n)
func (a IntArray) Find(predicate func(int) bool) IntOption
```

Example:

```go
arr := IntArray([]int {1,2,3,4,5,6})

var res1 IntOption = arr.Find(func(e int) bool { return e == 3 }) // Some(3)
var res2 IntOption = arr.Find(EvenInt)                            // Some(2)
var res3 IntOption = arr.Find(NegInt)                             // None
```
[[🠕]](#array0)



#### Array.Foreach

```go
// Applies function `f` to each element of array. 
// O(n)
func (a IntArray) Foreach(f func(int))
```
Example:

```go
arr := IntArray([]int{10, 20, 30})

// functional stype
arr.Foreach(func (e int) {              
    fmt.Println(e)           // print each element
})

// functional style (shortest)
arr.Foreach(PrintlnInt)

// imperative style
for _, e := range arr {
    fmt.Println(e)
}
```
[[🠕]](#array0)


#### Array.Heads and tails

```go
// Returns head of array or throws exeption if array is empty
// O(1)
func (a IntArray) Head() int

// Returns optional head from array
// O(1)
func (a IntArray) HeadOption() IntOption

// Returns array without first element
// O(c)
func (a IntArray) Tail() IntArray
```
Example:
```go
arr := IntArray([]int{10, 20, 30})

var h = arr.Head()                      // 10
var hopt IntOption = arr.HeadOption()   // Some(10)
var tail IntArray = arr.Tail()          // IntArray(20, 30)
var tail2 = tail.Tail()                 // IntArray(30)
var tail3 = tail2.Tail()                // empty array

var hopt3 = tail3.HeadOption()          // NoneInt
tail3.Head()                            // panic("there is no heads")
```
[[🠕]](#array0)


#### Array.Map

```go
// Transform each element of array to new element
// O(n)
func (a IntArray) Map<Type>(transformer func(int) <Type>) <Type>Array

func (a IntArray) MapInt(transformer func(int) int) IntArray
func (a IntArray) MapString(transformer func(int) string) StringArray
...
```

Example:

```go
arr := IntArray([]int{10, 20, 30})

var res StringArray = arr.
                        MapInt(func(e int) int { return e * 10 }).
                        MapString(func (e int) string) { return fmt.Sprintf("~%v~", e) }

res.Foreach(PrintString)
```

```
Output:
~100~
~200~
~300~
```
[[🠕]](#array0)



#### Array.MkString

```go
// Make string representation of that array with decorated elements and separator
// O(n)
func (a IntArray) MkString(start, sep, end string) string
```

Example:

```go
arr := IntArray([]int{1,2,3})
fmt.Println(arr.MkString("(", "|", ")"))  // (1|2|3)
```
[[🠕]](#array0)



#### Array.takes

```go
// Returns first n-elements
// O(1..n)
func (a IntArray) Take(n int) IntArray

// Returns last n-elements
// O(1..n)
func (a IntArray) TakeRight(n int) IntArray

// Returns first elements which statisfy a predicate
// O(1..n)
func (a IntArray) TakeWhile(predicate func(int) bool) IntArray
```

Example:

```go
arr := IntArray([]int{10,20,30,40,50})
res1 := arr.Take(2)                                 // Array(10,20)
res2 := arr.TakeRight(2)                            // Array(40,50)
res3 := arr.TakeWhile(func (e int) bool { e < 40})  // Array(10,20,30)
```
[[🠕]](#array0)



#### Array.ToList

```go
// Transform array to recursive functional data structure
// O(n)
func (a IntArray) ToList() IntList
```

Example:

```go
arr := IntArray([]int{1,2,3})
var l IntList = arr.ToList()
```
[[🠕]](#array0)


#### Array.ToString

```go
// Make string representation of that array
// O(n)
func (a IntArray) ToString() string
```

Example:

```go
arr := IntArray([]int{1,2,3})
fmt.Println(arr.ToString())                        // [1,2,3]

arr2 := IntIntArray([]int{1,2}, []int[]{3,4,5})
fmt.Println(arr2.ToString())                       // [[1,2], [3,4,5]]          
```
[[🠕]](#array0)


#### Supported array types
<details><summary>Supported array types (click to expand)</summary>

| Array type          | Scala analogue        | Go analogue   |
| ------------------- | --------------------- | ------------- |
| `BoolArray`         | Array[Bool]           | `[]bool`      |
| `StringArray`       | Array[String]         | `[]string`    |
| `IntArray`          | Array[Int]            | `[]int`       |
| `Int64Array`        | Array[Long]           | `[]int64`     |
| `ByteArray`         | Array[Byte]           | `[]byte`      |
| `RuneArray`         | Array[Rune]           | `[]rune`      |
| `Float32Array`      | Array[Float]          | `[]float32`   |
| `Float64Array`      | Array[Long]           | `[]float64`   |
| `AnyArray`          | Array[Any]            | `[]Any`       |
| `BoolArrayArray`    | Array[Array[Boolean]] | `[][]bool`    |
| `StringArrayArray`  | Array[Array[String]]  | `[][]string`  |
| `IntArrayArray`     | Array[Array[Int]]     | `[][]int`     |
| `Int64ArrayArray`   | Array[Array[Long]]    | `[][]int64`   |
| `ByteArrayArray`    | Array[Array[Byte]]    | `[][]byte`    |
| `RuneArrayArray`    | Array[Array[Char]]    | `[][]rune`    |
| `Float32ArrayArray` | Array[Array[Float]]   | `[][]float32` |
| `Float64ArrayArray` | Array[Array[Double]]  | `[][]float64` |
| `AnyArrayArray`     | Array[Array[Any]]     | `[][]Any`     |
</details>

[[🠕]](#array0)



## Boxed types
Boxed types are just wrappers under primitive go-types. These wrappers support additional operations which are desrcibed in this section.

[[🠕]](#boxed0)

### Underlined
```go
// Returns underlined value from boxed
func (e Int) Underlined() int
func (e String) Underlined() string
...
```
Example:
```go
var a Int = Int(10)   // explicit boxing
var b Int = 20        // implicit boxing

var v1 int = a.Underlined()   // returns underlined value 10
var v2 int = b.Underlined()   // returns underlined value 20

```

[[🠕]](#boxed0)


### Converters

#### String.ToArray
```go
// Converts string to array of runes (character codes)
func (s String) ToArray() RuneArray
```
Example:
```go
s := String("Hello")
arr := s.ToArray()      // RuneArray(72,101,108,108,111)
```
[[🠕]](#boxed0)

#### String.ToLetterArray
```go
// Coverts string to letter array
func (s String) ToLetterArray() StringArray
```
Example:
```go
s := String("Hello")
arr := s.ToLetterArray()  // StringArray("H", "e", "l", "l", "o")
```
[[🠕]](#boxed0)

#### String.ToInt
```go
// Converts string to integer or throws exception for wrong format 
func (s String) ToInt Int
```
Example:
```go
var n1 Int = String("10").ToInt()     // Int(10)
var n2 Int = String("zz").ToInt()     // panic("wrong format")
```
[[🠕]](#boxed0)

#### String.ToIntOption
```go
// Converts string to Some(int) for correct format and None for wrong format
func (s String) ToIntOption() IntOption
```
Example:
```go
var n1 IntOption = String("10").ToIntOption()  // Some(10)
var n2 IntOption = String("zz").ToIntOption()  // None
```
[[🠕]](#boxed0)

### List constructors
```go
// Create list from two elements `a` and `b`
func (a Int) Cons(b Int) IntList
func (a String) Cons(b String) StringList
...
```
Example:
```go
l1 := Int(10).Cons(20)                         // List(20, 10)
l2 := String("one").Cons("two").Cons("three")  // List("three", "two", "one")
```
[[🠕]](#boxed0)

### Range constructor

#### To
```go
// Creates list with values from range from=n, to=t inclusive
func (n Int) To(t Int) IntList
```
Example:
```go
l := Int(0).To(5)     // List(0,1,2,3,4,5)
```
[[🠕]](#boxed0)

#### Until
```go
// Creates list with values from range from=n, to=(t-1)
func (n Int) Until(t Int) IntList
```
Example:
```go
l := Int(0).To(5)     // List(0,1,2,3,4)
```
[[🠕]](#boxed0)

### Math and logic operations

#### IsBetween
```go
// check if a number is > left and < right (exclusive)
func (a Int) IsBetween(left, right int) bool
```
Example:
```go
Int(10).IsBetween(9, 15)      // true
Int(10).IsBetween(0, 5)       // false
Int(10).IsBetween(10, 11)     // false
```
[[🠕]](#boxed0)

#### IsBetweenInclusive
```go
// check if a number is >= left and <= right (inclusive)
func (a Int) IsBetweenInclusive(left, right int) bool
```
Example:
```go
Int(10).IsBetween(10, 15)      // true
Int(10).IsBetween(0, 5)        // false
```
[[🠕]](#boxed0)

#### Min/Max numeric
```go
// Returns min value of `a` and `b`
func (a Int) Min(b Int) Int
func (a Byte) Min(b Byte) Byte
...

// Returns max value of `a` and `b`
func (a Int) Max(b Int) Int
func (a Byte) Max(b Byte) Byte
...
```
Example:
```go
res1 := Int(10).Min(20)     // 10
res2 := Int(10).Max(20)     // 20
```
[[🠕]](#boxed0)

### Supported boxed types
<details><summary>Supported boxed types (click to expand)</summary>

| Boxed type | Go type      |
| ---------- | ------------ |
| Bool       | bool         |
| String     | string       |
| Int        | int          |
| Int8       | int8         |
| Int16      | int16        |
| Int32      | int32        |
| Int64      | int64        |
| Uint       | uint         |
| Uint16     | uint16       |
| Uint32     | uint32       |
| Uint64     | uint64       |
| Uintptr    | uintptr      |
| Byte       | byte         |
| Rune       | rune         |
| Float32    | float32      |
| Float64    | float64      |
| Complex64  | complex64    |
| Complex128 | complex128   |
| Any        | interface {} |
</details>

[[🠕]](#boxed0)

## Predicates
Predicates are used as function's arguments in the following operations on monads: `Filter`, `Find`, `Count`, `Exist` ...
[[🠕]](#predicates0)

### Empty predicates
```go
// Always returns true for all input values
var EmptyIntPredicate IntPredicate = func(t int) bool { return true }
...
```
Example:
```go

func getFilter() IntPredicate {
  var filterType string = config.getProperty("filter")
  
  if filterType == "even" {
     return EvenInt  // func (e int) bool { return e % 2 == 0 }
  } else if filterType == "odd" {
     return OddInt   // func (e int) bool { return e % 2 != 0 }
  } else {
    return EmptyIntPredicate   // default filter is no-filter
  }
}

l := MakeIntList(1,2,3)

res1 := l.Filter(getFilter())  // if config propery does not contain
                               // filter type then filtering will not use

```

[[🠕]](#predicates0)

### Numeric predicates
```go
// Check if number is even
var EvenInt IntPredicate = ...
var EvenInt64 Int64Predicate = ...
var EvenByte BytePredicate = ...
var EvenRune RunePredicate = ...

// Check if number is odd
var OddInt IntPredicate = ...
...

// Check if number is negative
var NegInt IntPredicate = ...
...

// Check if number is positive
var PosInt IntPredicate = ...
...

// Check if number is zero
var ZeroInt IntPredicate = ...
...

// Check if number is equal to 1
var OneInt IntPredicate = ...
...
```
Examples:
```go
l := MakeIntList(1,2,-3,-4,5,6,-7,8,9,0)

res1 := l.Filter(EvenInt)                // List(2,-4,6,8,0)
res2 := l.Filter(NegInt)                 // List(-3,-4,-7)
res3 := l.Filter(EvenInt).Filter(NegInt) // List(-4)
```
[[🠕]](#predicates0)

### String predicates
```go
// Build string predicate which check if string is matched to regexp
func MatchRegexp(r *regexp.Regexp) StringPredicate

// Build string predicate which check if string is matched to pattern
func MatchRegexpString(pattern string) StringPredicate
```
Examples:
```go
l := MakeStringList("Hello", "abc", "127.0.0.1", "255.255.255.0", "world", "127", "255")

var matchIp StringPredicate = 
        MatchRegexpString("[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}\\.[\\d]{1,3}")
   
ipList := l.Filter(matchIp)           // List("127.0.0.1", "255.255.255.0")

var onlyWords = MatchRegexpString("[a-zA-Z]+")
words := l.Filter(onlyWords)          // List("Hello", "abc", "world")
```

[[🠕]](#predicates0)

### Predicates composition
```go
// Apply boolean operator to two predicates
func (p1 IntPredicate) And(p2 IntPredicate) IntPredicate
func (p1 IntPredicate) Or(p2 IntPredicate) IntPredicate
func (p1 IntPredicate) Neg(p2 IntPredicate) IntPredicate
func (p1 IntPredicate) Xor(p2 IntPredicate) IntPredicate
...
```
Examples:
```go
l := MakeIntList(1,-2,3,4,5,-6,7,8,9,0)

var p1 IntPredicate = func(e int) bool { return e % 2 == 0 }  // even numbers
var p2 IntPredicate = func(e int) bool { return e < 0 }       // neg numbers

p3 := p1.And(p2)

l.Filter(p3)     // List(-2,-6)

```
[[🠕]](#predicates0)

### Supported predicate types
<details><summary>Supported predicate types (click to expand)</summary>

| Predicate type               |      |
| ---------------------------- | ---- |
| BoolPredicate                |      |
| StringPredicate              |      |
| IntPredicate                 |      |
| Int64Predicate               |      |
| BytePredicate                |      |
| RunePredicate                |      |
| Float32Predicate             |      |
| Float64Predicate             |      |
| AnyPredicate                 |      |
| BoolArrayPredicate           |      |
| StringArrayPredicate         |      |
| IntArrayPredicate            |      |
| Int64ArrayPredicate          |      |
| ByteArrayPredicate           |      |
| RuneArrayPredicate           |      |
| Float32ArrayPredicate        |      |
| Float64ArrayPredicate        |      |
| AnyArrayPredicate            |      |
| BoolArrayArrayPredicate      |      |
| StringArrayArrayPredicate    |      |
| IntArrayArrayPredicate       |      |
| Int64ArrayArrayPredicate     |      |
| ByteArrayArrayPredicate      |      |
| RuneArrayArrayPredicate      |      |
| Float32ArrayArrayPredicate   |      |
| Float64ArrayArrayPredicate   |      |
| AnyArrayArrayPredicate       |      |
| BoolOptionPredicate          |      |
| StringOptionPredicate        |      |
| IntOptionPredicate           |      |
| Int64OptionPredicate         |      |
| ByteOptionPredicate          |      |
| RuneOptionPredicate          |      |
| Float64OptionPredicate       |      |
| AnyOptionPredicate           |      |
| BoolOptionOptionPredicate    |      |
| StringOptionOptionPredicate  |      |
| IntOptionOptionPredicate     |      |
| Int64OptionOptionPredicate   |      |
| RuneOptionOptionPredicate    |      |
| Float32OptionOptionPredicate |      |
| Float64OptionOptionPredicate |      |
| AnyOptionOptionPredicate     |      |
| BoolArrayOptionPredicate     |      |
| StringArrayOptionPredicate   |      |
| IntArrayOptionPredicate      |      |
| Int64ArrayOptionPredicate    |      |
| ByteArrayOptionPredicate     |      |
| RuneArrayOptionPredicate     |      |
| Float32ArrayOptionPredicate  |      |
| Float64ArrayOptionPredicate  |      |
| AnyArrayOptionPredicate      |      |
| BoolListOptionPredicate      |      |
| StringListOptionPredicate    |      |
| IntListOptionPredicate       |      |
| Int64ListOptionPredicate     |      |
| ByteListOptionPredicate      |      |
| RuneListOptionPredicate      |      |
| Float32ListOptionPredicate   |      |
| Float64ListOptionPredicate   |      |
| AnyListOptionPredicate       |      |
| BoolListPredicate            |      |
| StringListPredicate          |      |
| IntListPredicate             |      |
| Int64ListPredicate           |      |
| ByteListPredicate            |      |
| RuneListPredicate            |      |
| Float32ListPredicate         |      |
| Float64ListPredicate         |      |
| AnyListPredicate             |      |
| BoolArrayListPredicate       |      |
| StringArrayListPredicate     |      |
| IntArrayListPredicate        |      |
| Int64ArrayListPredicate      |      |
| ByteArrayListPredicate       |      |
| RuneArrayListPredicate       |      |
| Float32ArrayListPredicate    |      |
| Float64ArrayListPredicate    |      |
| AnyArrayListPredicate        |      |
| BoolOptionListPredicate      |      |
| StringOptionListPredicate    |      |
| IntOptionListPredicate       |      |
| Int64OptionListPredicate     |      |
| ByteOptionListPredicate      |      |
| RuneOptionListPredicate      |      |
| Float32OptionListPredicate   |      |
| Float64OptionListPredicate   |      |
| AnyOptionListPredicate       |      |
| BoolListListPredicate        |      |
| StringListListPredicate      |      |
| IntListListPredicate         |      |
| Int64ListListPredicate       |      |
| ByteListListPredicate        |      |
| RuneListListPredicate        |      |
| Float32ListListPredicate     |      |
| Float64ListListPredicate     |      |
| AnyListListPredicate         |      |

</details>

[[🠕]](#predicates0)

## Transformers
Transformers are functions which transform one type to other type `A => B`. Transformers are used as function's arguments in the following operations on monads: `Map`, `FlatMap`, `GroupBy`
[[🠕]](#transformers0)

### Identity
```go
// Just return value without transformation
var IntIdentity func(int) int = func(i int) int { return i }
...
```
Examples:
```go
l := MakeIntList(1,2,3,1,3,3,4)

groups := l.GroupBy(IntIdentity)
/*
   Map(1 -> List(1,1),
       2 -> List(2),
       3 -> List(3,3,3),
       4 -> List(4))
*/
```

[[🠕]](#transformers0)

### String transformers
```go
// Build transformer which transform string to matched regex groups
func RegexGroups(r *regexp.Regexp) func(string) []string
func StringRegexGroups(pattern string) func(string) []string
```
Example:
```go
l := MakeStringList("10014-dav", "10015-pav", "10016-ant", "10017-din")


regex := StringRegexGroups("([0-9]+)\\-([a-z]+)") // string => []string transformer

res1 := l.MapStringArray(idNameRegex) 
/* StringArrayList (Array("10015-dav", "10014", "dav"),
                    Array("10015-pav", "10015", "pav"),
                    Array("10016-ant", "10016", "ant"),
                    Array("10017-din", "10017", "din")) */
```
[[🠕]](#transformers0)

## Predef
Contains some convenient utils.
[[🠕]](#predef0)
### Require
```go
// Throws exception with message if e-condition is false
// Can be used as validation of public functions argument
func Require(e bool, msg string)
```
Examples:
```go
func Div(a, b Int) Float32 {
  Require(a != 0, "Divide by zero!")
  ...
}
```
[[🠕]](#predef0)

## Concurrency API

Concurrence API implements a `Future` as abstraction over go routines and channels. It allows to combine async calculations in blocking and non blocking manner.
All examples of this section use a `IntFuture` type. But this API is supported in the other future types also. (See [Supported future types section](#supported-future-types))

Before using concurrence API library import the following packages:
```go
import (
  . "github.com/zx80live/gofp/fp"
  . "github.com/zx80live/gofp/fp/concurrent"
)
```
[[🠕]](#concurrency0)

### Create future
```go
// Execute go-routine for calculate function
func MakeIntFuture(f func() Int) IntFuture
```
Example:
```go
t1 := time.Now().Unix()
// create future
f1 := MakeIntFuture(func() Int {
        time.Sleep(2 * time.Second)
        return Int(10)
      })
t2 := time.Now().Unix()
fmt.Println("at ", t2 - t1)      // at 0 milliseconds
                                 // microbenchmark demonstrates async running of IntFuture

// If you want to wait and get result of future use Result method
// Warning! But this method blocks current thread (routine)
t3 := time.Now().Unix()
res1 := f1.Result()
t4 := time.Now().Unix()

fmt.Println("result ", res1, "at ", t4 - t3)  // result 10 at 2 seconds 

```
[[🠕]](#future0)

### Future.FlatMap
```go
// Gives the ability to compose async operations together
func (f IntFuture) FlatMap<Type>(t func(Int) <Type>Future) <Type>Future

// examples
func (f IntFuture) FlatMapInt(t func(Int) IntFuture) IntFuture
func (f IntFuture) FlatMapString(t func(Int) StringFuture) StringFuture
...
```
[[🠕]](#future0)

#### Future.FlatMap.Blocking composition
`FlatMap` allows to compose two async calculations (which are presented as two futures) in the blocking manner. In other words, the resulting future will be result of sequential calculations of two futures. So the inner future which is mapped to outer future will be started when outer future will be completed.
This example uses `Future.Result` invocation which is described in [Future.Result](#futureresult) section.

```go
// Task: implement two async functions.
// Each function takes some time (2 seconds) and return some number.
// And then we should multiple results of these functions in blocking manner.
// Full time of execution both futures will be ~4000 milliseconds.
// Therefore these futures will be executed SEQUENTIALLY

t1 := time.Now().Unix()
f := MakeIntFuture(func () Int {
    time.Sleep(2 * time.Second)                             // some payload emulation
    return Int(10)                             
}).FlatMapInt(func (a Int) Int {
    return MakeIntFuture(func () Int {
        time.Sleep(2 * time.Second)                         // some payload emulation
        return Int(e * 20)                     
    })
})
t2 := time.Now().Unix()
fmt.Println("create and compose futures at", t2 - t1)       // at 0 milliseconds


t3 := time.Now().Unix()
// block current main-thread and wait and get result of future
var res1 Int = f.Result()                                   // 10 * 20
t4 := time.Now().Unix()
fmt.Println("get result of composing futures at", t4 - t3)  // at 4000 milliseconds
                                                            // SEQUENTIAL execution

```
[[🠕]](#future0)

#### Future.FlatMap.NON-blocking composition
`FlatMap` allows to compose two async calculations (which are presented as two futures) in the non-blocking manner. In other words, the resulting future will be result of parallel calculations of two futures. So the both futures should be defined before than they will be composed via `FlatMap`. In this case they will be executed parallely.
This example uses `Future.Result` invocation which is described in [Future.Result](#futureresult) section.

```go

// Task: implement two async functions.
// Each function takes some time (2 seconds) and return some number.
// And then we should multiple results of these functions in NON-blocking manner.
// Full time of executions both futures will be ~2000 milliseconds.
// Therefore these futures will be executed PARALLEL

t1 := time.Now().Unix()
f1 := MakeIntFuture(func () Int {
        time.Sleep(2 * time.Second)                         // some payload emulation
        return Int(10)
      })

f2 := MakeIntFuture(func () Int {
        time.Sleep(2 * time.Second)                         // some payload emulation
        return Int(20)
      })

// compose futures in NON-blocking manner
var futureResult IntFuture = f1.FlatMapInt(func (a Int) IntFuture {
                                 return f2.MapInt(func (b Int) Int {
                                   return Int(a * b)
                                 })
                             })

t2 := time.Now().Unix()
fmt.Println("create and compose futures at", t2 - t1)       // at 0 milliseconds


t3 := time.Now().Unix()
// block current main-thread and wait and get result of future
var res1 Int = futureResult.Result()                        // 10 * 20
t4 := time.Now().Unix()
fmt.Println("get result of composing futures at", t4 - t3)  // at 2000 milliseconds
                                                            // PARALLEL execution
                                                            
```
[[🠕]](#future0)

### Future.Map
```go
// Transform content of the future
func (f IntFuture) Map<Type>(t func(Int) <Type>) <Type>Future

// examples
func (f IntFuture) MapInt(t func(Int) Int) IntFuture
func (f IntFuture) MapString(t func(Int) String) StringFuture
...
```
Example:
```go
f1 := MakeIntFuture(func() Int { return Int(10) })             // Future(10)

res1 := f1.MapInt(func(e Int) Int { return Int(e * 10) })      // Future(100)
res2 := f1.MapString(func(e Int) String { 
    return String(fmt.Sprintf("<%v>", e)) })                   // Future("<10>")
```
[[🠕]](#future0)

### Future.Result
```go
// Await and return the result of future
// This method blocks the current thread (routine) from which it was invoked
func (f IntFuture) Result() Int
```
Example:
```go
t1 := time.Now().Unix()
f1 := MakeIntFuture(func() Int { return Int(10) })  
f2 := MakeIntFuture(func() Int { 
	time.sleep(2 * time.Second)
	return Int(20) 
})                                
t2 := time.Now().Unix()

fmt.Println("create futures at ", t2 - t1)          // at ~0 milliseconds


t3 := time.Now().Unix()
res1 := f1.Result()                
t4 := time.Now().Unix()
fmt.Printl("blocks current thread and return result of f1 at", t4 - t3)   // at ~0 millis

t5 := time.Now().Unix()
res2 := f2.Result()                
t5 := time.Now().Unix()
fmt.Printl("blocks current thread and return result of f2 at", t6 - t5)   // at ~2000 millis

```
[[🠕]](#future0)

### Future.Success
```go
// Create already completed future with specified result
// Thus, it can be used when known result should be returned as future
func SuccessIntFuture(v Int) IntFuture
```
Example:
```go
var f1 IntFuture = SuccessIntFuture(10)      // Future(10)
```
[[🠕]](#future0)

### Supported future types
<details><summary>Supported future types (click to expand)</summary>

| Future type              | Scala analogue                |
| ------------------------ | ----------------------------- |
| BoolFuture               | Future[Boolean]               |
| StringFuture             | Future[String]                |
| IntFuture                | Future[Int]                   |
| Int64Future              | Future[Long]                  |
| ByteFuture               | Future[Byte]                  |
| RuneFuture               | Future[Char]                  |
| Float32Future            | Future[Float]                 |
| Float64Future            | Future[Double]                |
| AnyFuture                | Future[Any]                   |
| BoolOptionFuture         | Future[Option[Boolean]]       |
| StringOptionFuture       | Future[Option[String]]        |
| IntOptionFuture          | Future[Option[Int]]           |
| Int64OptionFuture        | Future[Option[Long]]          |
| ByteOptionFuture         | Future[Option[Byte]]          |
| RuneOptionFuture         | Future[Option[Char]]          |
| Float32OptionFuture      | Future[Option[Float]]         |
| Float64OptionFuture      | Future[Option[Double]]        |
| AnyOptionFuture          | Future[Option[Any]]           |
| BoolListOptionFuture     | Future[Option[Boolean]]       |
| StringListOptionFuture   | Future[Option[List[String]]]  |
| IntListOptionFuture      | Future[Option[List[Int]]]     |
| Int64ListOptionFuture    | Future[Option[List[Long]]]    |
| ByteListOptionFuture     | Future[Option[List[Byte]]]    |
| RuneListOptionFuture     | Future[Option[List[Char]]]    |
| Float32ListOptionFuture  | Future[Option[List[Float]]]   |
| Float64ListOptionFuture  | Future[Option[List[Double]]]  |
| AnyListOptionFuture      | Future[Option[List[Any]]]     |
| BoolArrayOptionFuture    | Future[Option[Array[Bool]]]   |
| StringArrayOptionFuture  | Future[Option[Array[String]]] |
| IntArrayOptionFuture     | Future[Option[Array[Int]]]    |
| Int64ArrayOptionFuture   | Future[Option[Array[Long]]]   |
| ByteArrayOptionFuture    | Future[Option[Array[Byte]]]   |
| RuneArrayOptionFuture    | Future[Option[Array[Char]]]   |
| Float32ArrayOptionFuture | Future[Option[Array[Float]]]  |
| Float64ArrayOptionFuture | Future[Option[Array[Double]]] |
| AnyArrayOptionFuture     | Future[Option[Array[Any]]]    |
| BoolListFuture           | Future[List[Boolean]]         |
| StringListFuture         | Future[List[String]]          |
| IntListFuture            | Future[List[Int]]             |
| Int64ListFuture          | Future[List[Long]]            |
| ByteListFuture           | Future[List[Byte]]            |
| RuneListFuture           | Future[List[Char]]            |
| Float32ListFuture        | Future[List[Float]]           |
| Float64ListFuture        | Future[List[Long]]            |
| AnyListFuture            | Future[List[Any]]             |
| BoolArrayFuture          | Future[Array[Boolean]]        |
| StringArrayFuture        | Future[Array[String]]         |
| IntArrayFuture           | Future[Array[Int]]            |
| Int64ArrayFuture         | Future[Array[Long]]           |
| ByteArrayFuture          | Future[Array[Byte]]           |
| RuneArrayFuture          | Future[Array[Char]]           |
| Float32ArrayFuture       | Future[Array[Float]]          |
| Float64ArrayFuture       | Future[Array[Double]]         |
| AnyArrayFuture           | Future[Array[Any]]            |
</details>

[[🠕]](#future0)

## Bootstrap
Almost all code base of this library was generated by code-generator from the project [gofp-bootstrap](https://github.com/zx80live/gofp-bootstrap). That code generator is implemented in Scala language for the following reasons:
 - string interpolations
 - functional programming, etc
 
But the next generation of code generator can be implemented by the current [gofp](https://github.com/zx80live/gofp) library.
[[🠕]](#bootstrap0)

## License

This library is distributed under MIT license found in the LICENSE file.

[[🠕]](#table-of-contents)
