// Generated by gen_types. Don't modify manually!

package main

import (
  "fmt"
)

func assertComparable(obj Object, msg string) Comparable {
  switch c := obj.(type) {
  case Comparable:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Comparable", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertVector(obj Object, msg string) *Vector {
  switch c := obj.(type) {
  case *Vector:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Vector", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertChar(obj Object, msg string) Char {
  switch c := obj.(type) {
  case Char:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Char", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertString(obj Object, msg string) String {
  switch c := obj.(type) {
  case String:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "String", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertSymbol(obj Object, msg string) Symbol {
  switch c := obj.(type) {
  case Symbol:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Symbol", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertKeyword(obj Object, msg string) Keyword {
  switch c := obj.(type) {
  case Keyword:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Keyword", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertBool(obj Object, msg string) Bool {
  switch c := obj.(type) {
  case Bool:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Bool", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertNumber(obj Object, msg string) Number {
  switch c := obj.(type) {
  case Number:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Number", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertSeq(obj Object, msg string) Seq {
  switch c := obj.(type) {
  case Seq:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Seq", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}

func assertCallable(obj Object, msg string) Callable {
  switch c := obj.(type) {
  case Callable:
    return c
  default:
    if msg == "" {
      msg = fmt.Sprintf("Expected %s, got %s", "Callable", obj.GetType().ToString(false))
    }
    panic(RT.newError(msg))
  }
}