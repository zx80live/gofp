package fp

/*
  func (l FListInt) MapInt(f FunctorInt) FListInt {
    return FListInt {
      head : l.head,
      tail : l.tail,
      functor : func(e Element) Element {
        if processed := l.functor(e); processed != nil {
                                                         return f(processed.(int))
                                                         } else {
                                                                                     return nil
                                                                                     }
      },
    }
  }

  func (l FListInt) MapString(f FunctorIntString) FListString {
    panic("")
  }

  func foo() {
    var l1 FListInt = FListInt { 0, nil, EmptyFunctor }

    l1.MapInt(func(e int) int {
      return e * 10
    }).MapString(func(e int) string {
      return "hello"
    })

    l1.MapString(func(e int) string {
      return "hello"
    })
  }
*/
