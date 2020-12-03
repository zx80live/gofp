# Functional programming in GO



## Introduction

This library was inspired by Scala (collection API, functional paradigm and etc).  

## Table of contents

- [Getting started](#getting-started)
- [Collection API](#collection-api)
  * [Lists](#lists)
    + [Create list](#create-list)
    + [Copy list](#list.copy)
    + [Count](#list.count)
    + [Drops](#drops)
    + [Equals](#equals)
    + [Filter](#filter)
    + [Find](#find)
    + [FlatMap](#flatmap)
    + [Flatten](#flatten)
    + [FoldLeft](#foldleft)
    + [Foreach](#foreach)
    + [GroupBy](#groupby)
    + [Heads and tails](#heads-and-tails)
    + [Emptiness](#emptiness)
    + [Map](#map)
    + [MkString](#mkstring)
    + [Nil](#nil)
    + [Prepend](#prepend)
    + [Reduce](#reduce)
    + [Reverse](#reverse)
    + [Size](#size)
    + [List takes](#list-takes)
    + [ToArray](#toarray)
    + [ToString](#tostring)
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

List is recursive functional data structure which has `head` and `tail` as sublist. 

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

[arr🠕](#table-of-contents)

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

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.FlatMap

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Flatten

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.FoldLeft

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Foreach

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.GroupBy

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Heads and tails

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Emptiness

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Map

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.MkString

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Nil

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Prepend

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Reduce

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Reverse

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.Size

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.takes

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.ToArray

```go

```

Example:

```go

```

[🠕](#table-of-contents)



#### List.ToString

```go

```

Example:

```go

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
// Make string representation of that array with decorated elements and separatorArray takes
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
