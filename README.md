# Functional programming in GO

This library was inspired by Scala (collection API, functional paradigm and etc). 



## Introduction



## First steps

Before using functional library import the following package:

```go
import . "github.com/zx80live/gofp/fp"
```



## Collection API

Current library supports the following collection types: `Arrays`, `Lists` and `Options`. Each type is monad and supports functions such as `Map`, `FlatMap`, `Filter` and etc.



### Lists



### Options



### Arrays

Array is just wrapper for go-arrays which contains convenient functions. All examples of  this section use a `IntArray` type. But this API is supported in the other array types also. (See [Supported array types] section)



##### Create array

```go
type IntArray []int
```

Example:

```go
arr := IntArray([]int{10, 20, 30})
```



##### Traverse array

```go
// apply function `f` to each element of array
// complexity: O(n)
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



##### Heads and tails

```go
// returns head of array or throws exeption if array is empty
// complexity: O(1)
func (a IntArray) Head() int

// returns optional head from array
// complexity: O(1)
func (a IntArray) HeadOption() IntOption

// returns array without first element
// complexity: O(c)
func (a IntArray) Tail() IntArray
```

Examples:

```go
arr := IntArray([]int{10, 20, 30})

var h = arr.Head()                      // returns 10
var hopt IntOption = arr.HeadOption()   // returns Some(10)
var tail IntArray = arr.Tail()          // returns IntArray(20, 30)
var tail2 = tail.Tail()                 // returns IntArray(30)
var tail3 = tail2.Tail()                // returns empty array

var hopt3 = tail3.HeadOption()          // return NoneInt
tail3.Head()                            // panic("there is no heads")
```



##### Array Map

```go
// transform each element of array to new element
// complexity: O(n)
func (a IntArray) Map<Type>(transformer func(int) <Type>) <Type>Array

// examples:
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



##### Array Filter

```go
// returns new array with elements which statisfy a predicate
// complexity: O(n)
func (a IntArray) Filter(predicate func(int) bool) IntArray
```

Examples:

```go
arr := IntArray([]int {1,-2,-3,4,5,-6,7,8,-9,10})
res := arr.
         Filter(func (e int) bool { return e % 2 == 0}).   // filter even numbers
         Filter(func (e int) bool { return e >= 0 })       // filter positive numbers

fmt.Println(res.ToString())    // [4, 8, 10]

res2 := arr.Filter(EvenInt).Filter(PosInt)                 // use library predicates
res3 := arr.Filter(EvenInt.And(PosInt))                    // compose predicates

```



##### Array Find

```go
// returns first element which statisfy a predicate
// complexity: O(1..n)
func (a IntArray) Find(predicate func(int) bool) IntOption
```

Examples:

```go
arr := IntArray([]int {1,2,3,4,5,6})

var res1 IntOption = arr.Find(func(e int) bool { return e == 3 }) // Some(3)
var res2 IntOption = arr.Find(EvenInt)                            // Some(2)
var res3 IntOption = arr.Find(NegInt)                             // None
```



##### Array Count

```go
// returns count of elements which satisfy a predicate
// complexity: O(n)
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



##### Array drops

```go
// returns new array without n-first elements
// complexity: O(1..n)
func (a IntArray) Drop(n int) IntArray

// returns new array without n-last elements
// complexity: O(1..n)
func (a IntArray) DropRight(n int) IntArray

// returns new array without first elements which statisfy a predicate
// complexity: O(1..n)
func (a IntArray) DropWhile(predicate func(int) bool) IntArray
```

Examples:

```go
arr := IntArray([]int{10,20,30,40,50})
res1 := arr.drop(2)                                 // Array(30,40,50)
res2 := arr.dropRight(2)                            // Array(10,20,30)
res3 := arr.dropWhile(func (e int) bool { e < 40 }) // Array(40,50)
```



##### Array takes

```go
// returns first n-elements
// complexity: O(1..n)
func (a IntArray) Take(n int) IntArray

// returns last n-elements
// complexity: O(1..n)
func (a IntArray) TakeRight(n int) IntArray

// returns first elements which statisfy a predicate
// complexity: O(1..n)
func (a IntArray) TakeWhile(predicate func(int) bool) IntArray
```

Examples:

```go
arr := IntArray([]int{10,20,30,40,50})
res1 := arr.take(2)                                 // Array(10,20)
res2 := arr.takeRight(2)                            // Array(40,50)
res3 := arr.takeWhile(func (e int) bool { e < 40})  // Array(10,20,30)
```



##### Array Equals

```go
// returns true if both arrays are equal
// complexity: O(n)
func (a IntArray) Equals(b IntArray) bool
```

Examples:

```go
arr1 := IntArray([]int{10,20,30,40,50})
arr2 := IntArray([]int{10,20,30,40,50})
arr3 := IntArray([]int{10,20,30})

arr1.Equals(arr2)   // true
arr1.Equals(arr3)   // false
```



##### Array ToString

```go
// make string representation of that array
// complexity: O(n)
func (a IntArray) ToString() string
```

Examples:

```go
arr := IntArray([]int{1,2,3})
fmt.Println(arr.ToString())                        // [1,2,3]

arr2 := IntIntArray([]int{1,2}, []int[]{3,4,5})
fmt.Println(arr2.ToString())                       // [[1,2], [3,4,5]]          
```



##### Array MkString

```go
// make string representation of that array with decorated elements and separatorArray takes
// complexity: O(n)
func (a IntArray) MkString(start, sep, end string) string
```

Examples:

```go
arr := IntArray([]int{1,2,3})
fmt.Println(arr.MkString("(", "|", ")"))  // (1|2|3)
```



##### Array ToList

```go
// transform array to recursive functional data structure
// complexity: O(n)
func (a IntArray) ToList() IntList
```

Examples:

```go
arr := IntArray([]int{1,2,3})
var l IntList = arr.ToList()
```



##### Supported array types

| Array type                       | Go analogue   |
| -------------------------------- | ------------- |
| `BoolArray`                      | `[]bool`      |
| `StringArray`                    | `[]string`    |
| `IntArray`                       | `[]int`       |
| `Int64Array`                     | `[]int64`     |
| `ByteArray`                      | `[]byte`      |
| `RuneArray`                      | `[]rune`      |
| `Float32Array`                   | `[]float32`   |
| `Float64Array`                   | `[]float64`   |
| `AnyArray`                       | `[]Any`       |
| `BoolArrayArray`                 | `[][]bool`    |
| `StringArrayArrayRuneArrayArray` | `[][]string`  |
| `IntArrayArray`                  | `[][]int`     |
| `Int64ArrayArray`                | `[][]int64`   |
| `ByteArrayArray`                 | `[][]byte`    |
| `RuneArrayArray`                 | `[][]rune`    |
| `Float32ArrayArray`              | `[][]float32` |
| `Float64ArrayArray`              | `[][]float64` |
| `AnyArrayArray`                  | `[][]Any`     |



### Lists

List is recursive functional data structure which has `head` and `tail` as sublist.





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