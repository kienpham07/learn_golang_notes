# Golang Interfaces: A Comprehensive Study Guide

This guide provides a structured overview of Interfaces in Golang, based on the CH5 lecture slides. It is designed for beginners to understand how to leverage abstractions for cleaner, more flexible code.

-----

## At a Glance: Golang Interfaces

| Feature             | Description                                |
| :------------------ | :----------------------------------------- |
| **Definition**      | Collections of method signatures.          |
| **Implementation**  | Implicit (no `implements` keyword).        |
| **Empty Interface** | `interface{}` is implemented by all types. |
| **Best Practice**   | Keep interfaces small and focused.         |

-----

## Lecture 1: Core Concepts of Interfaces

### Lecture Title: Understanding Interface Basics

Interfaces allow you to focus on **what a type does** rather than how it is built. They define shared behaviors that different types can implement, making code modular and easy to update.

### Syntax & Examples

An interface is defined using the `type` and `interface` keywords.

``` go
// Define an interface with method signatures
type shape interface {
    area() float64
    perimeter() float64
}

// Structs that will implement the interface
type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

// Implementation for rect
func (r rect) area() float64 {
    return r.width * r.height
}

// Implementation for circle
func (c circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

```

### Deep Explanation

  * **Method Signatures:** Interfaces only contain the names, parameters, and return types of methods. They do not contain any logic.
  * **Implicit Implementation:** Unlike Java or C\#, Go has no `implements` keyword. If a type has the required methods, it *automatically* satisfies the interface.
  * **The Empty Interface (`interface{}`):** Because it has no required methods, every single type in Go satisfies the empty interface. This is often used to handle values of unknown types.

### Key Takeaways

  * Interfaces decouple definitions from implementations.
  * A type can implement multiple interfaces simultaneously.
  * To check if a struct implements an interface, try passing it into a function that requires that interface.

-----

## Lecture 2: Interface Best Practices

### Lecture Title: Writing Clean Interfaces

Writing abstractions is difficult; keeping them simple is the key to maintainable code.

### Deep Explanation

  * **Keep Interfaces Small:** Interfaces should define the *minimal* behavior necessary. A good example is `io.Reader`, which only has one method.
  * **Name Your Parameters:** While `Copy(string, string) int` is valid, it is ambiguous. Using `Copy(sourceFile string, destinationFile string) (bytesCopied int)` provides immediate clarity to the developer.
  * **Avoid Knowledge of Satisfying Types:** An interface should not know about the structs that implement it. For example, a `car` interface should not have an `IsFiretruck()` method.
  * **Interface Embedding:** You can create "supersets" of interfaces by embedding one into another.

<!-- end list -->

``` go
type car interface {
    Color() string
}

// firetruck 'inherits' Color() from car and adds HoseLength()
type firetruck interface {
    car
    HoseLength() int
}

```

### Summary Notes

  * **Small interfaces** reduce implementation complexity and increase flexibility.
  * **Naming parameters** improves readability and documentation without affecting execution speed.
  * **Interfaces are not classes:** They don't have constructors, destructors, or hierarchical data.

-----

## Lecture 3: Type Assertions & Switches

### Lecture Title: Extracting Underlying Types

Sometimes you need to access the concrete type hidden behind an interface value. This is done via type assertions and type switches.

### Syntax & Examples

**Type Assertion**

``` go
// s is a shape interface
c, ok := s.(circle)
if ok {
    // Now you can access circle-specific fields
    fmt.Println(c.radius)
}

```

**Type Switch**

``` go
func printNumericValue(num interface{}) {
    switch v := num.(type) {
    case int:
        fmt.Printf("Integer: %v\n", v)
    case string:
        fmt.Printf("String: %v\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}

```

### Common Mistakes

  * **Panic on Assertion:** Performing a type assertion like `c := s.(circle)` without the `ok` boolean will cause a panic if the assertion fails. Always use the "comma, ok" idiom.
  * **Bloated Interfaces:** Adding specific check methods (like `IsFiretruck()`) instead of using a type assertion is considered an anti-pattern.

-----

## Final Golang Interface Cheat Sheet

  * **Definition:** `type <Name> interface { Method() <Type> }`
  * **Implementation:** Just define the methods on a struct. No keyword needed.
  * **Empty Interface:** `interface{}` accepts any value.
  * **Type Assertion:** `value, ok := element.(ConcreteType)`
  * **Type Switch:** `switch v := element.(type) { case T: ... }`
  * **Embedding:** Interfaces can be combined by listing them inside a new interface.
  * **Rule of Thumb:** "The bigger the interface, the weaker the abstraction." Keep them small.
