# Functional programming in GO

Project and documentation in progress.

## Introduction

This library was inspired by Scala (collection API, functional paradigm and etc).  

## Table of contents

- [Getting started](#getting-started)
- [Collection API](#collection-api)
  * [Lists](#lists)
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
    + [List.Emptiness](#listemptiness)
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
  * [Options](#options)
  * [Arrays](#arrays)
    + [Create](#create-array)
    + [Foreach](#array.foreach)
    + [Heads and tails](#heads-and-tails)
    + [Map](#map)
    + [Filter](#filter)
    + [Find](#find)
    + [Count](#count)
    + [Array drops](#array-drops)
    + [Array takes](#array-takes)
    + [Equals](#equals)
    + [ToString](#tostring)
    + [MkString](#mkstring)
    + [ToList](#tolist)
    + [Supported array types](#supported-array-types)
- [Boxed types](#boxed-types)


## Getting started

Before using functional library import the following package:

```go
import . "github.com/zx80live/gofp/fp" 
```



## Collection API

Current library supports the following collection types: `Arrays`, `Lists` and `Options`. Each type is monad and supports functions such as `Map`, `FlatMap`, `Filter` and etc.



### Lists

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

[🠕](#table-of-contents)

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
```

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)

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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

var sum int = l.FoldLeftInt(0, func(acc, el) int { return acc + el})  // sum = (((((0 + 1) + 2) + 3) + 4) + 5)
str = l.FoldLeftString(">", func(acc, el) string { return fmt.Sprintf("%v|%v", acc, el)}) // >1|2|3|4|5
```

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



#### List.GroupBy

```go
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
  Cons(Shape{"circle", "red", 20}).
  Cons(Shape{"circle", "yellow", 30}).
  Cons(Shape{"triangle", "green", 10}).
  Cons(Shape{"triangle", "red", 20}).
  Cons(Shape{"triangle", "yellow", 30}).
  Cons(Shape{"polygon", "green", 10}).
  Cons(Shape{"polygon", "red", 20}).
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



[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



#### List.Emptiness

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

[🠕](#table-of-contents)



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
res1 := l.MapInt(func(e int) int { return e * 10})                         // List(10,20,30)
res2 := l.MapString(func (e int) string { return fmt.Sprintf("<%v>", e)})  // List("<1>", "<2>", "<3>")
```

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



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

[🠕](#table-of-contents)



#### List.ToString

```go
// Transform list to string
func (l IntList) ToString() string
```

Example:

```go
l1 := MakeIntList(1,2,3)
res1 := l1.ToString()                                              // "List(1,2,3)"
l2 := MakeIntListList(MakeIntList(1,2), NilInt, MakeIntList(3,4))  // "List(List(1,2), List(), List(3,4))"
```

[🠕](#table-of-contents)





#### Supported list types

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

[🠕](#table-of-contents)




### Options
TODO describe


### Arrays

Array is just wrapper for go-arrays which contains convenient functions. All examples of  this section use a `IntArray` type. But this API is supported in the other array types also. (See [Supported array types] section)



#### Create array
```go
type IntArray []int
```
Example:
```go
arr := IntArray([]int{10, 20, 30})
```
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


#### Array.Array drops

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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


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
[🠕](#table-of-contents)


#### Supported array types

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

[🠕](#table-of-contents)



### Optional



## Boxed (rich) types

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
[🠕](#table-of-contents)




## Boxed types



## Concurrency API




## Contributing

TODO describe

## Versioning

TODO describe

## Roadmap

TODO describe

## License

This library is distributed under MIT license found in the LICENSE file.

TODO clarify
[🠕](#table-of-contents)
