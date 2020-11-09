package fp

type FList struct {
  head Element
  tail *FList
  functor Functor
}

type FListInt struct {
  head int
  tail *FListInt
  functor Functor
}

type FListString struct {
  head string
  tail *FListString
  functor Functor
}
