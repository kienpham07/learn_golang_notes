# Golang Lecture Notes: Chapter 4 - Structs

This guide provides a comprehensive overview of Structs in Go, covering everything from basic definitions to memory optimization.

-----

## 1\. Introduction to Structs

### Core Concepts

Structs are used in Go to represent structured data. They allow you to group different types of variables together into a single custom type. In other languages, you might use objects or dictionaries for this purpose.

### Syntax & Examples

To define a struct, use the `type` and `struct` keywords:

``` go
type car struct {
    brand   string
    model   string
    doors   int
    mileage int
}

```

  * **Type Name**: `car` is the name of the new struct type.
  * **Fields**: `brand`, `model`, `doors`, and `mileage` are the fields all `car` instances will have.

### Key Takeaways

  * Structs are the primary way to create custom data structures in Go.
  * They provide a way to model real-world entities cleanly.

-----

## 2\. Anonymous Structs

### Core Concepts

An anonymous struct is a struct defined without a name. Because it lacks a name, it cannot be referenced elsewhere in your code.

### Syntax & Examples

You define and instantiate an anonymous struct simultaneously using a second pair of brackets:

``` go
myCar := struct {
    brand string
    model string
}{
    brand: "Toyota",
    model: "Camry",
}

```

### Deep Explanation

**When to Use:**

  * **One-time use**: Ideal for data structures that aren't needed anywhere else, such as shaping JSON data for an HTTP handler.
  * **Preventing Reuse**: It signals to other developers that this specific structure is not intended to be reused.

**Nesting Anonymous Structs:**
You can also use anonymous structs as fields within named structs:

``` go
type car struct {
    brand string
    wheel struct {
        radius   int
        material string
    }
}

```

### Common Mistakes

  * **Overusing them**: Prefer named structs for readability and reusability unless the "use-once" rule applies.

-----

## 3\. Nested vs. Embedded Structs

### Core Concepts

While Go is not strictly object-oriented, it offers two ways to compose structs: nesting and embedding.

### Syntax & Examples

#### A. Nested Structs (The "Explicit" Way)

In a nested struct, the inner struct is given an explicit field name.

``` go
type Engine struct {
    Horsepower int
}

type Car struct {
    brand string
    eng   Engine // Explicit field name "eng"
}

// Access: Must chain field names
myCar.eng.Horsepower = 300

```

#### B. Embedded Structs (The "Shortcut" Way)

In an embedded struct, you declare the inner struct without a field name—only the type name is used. Go uses **field promotion** to make inner fields accessible at the top level.

``` go
type Car struct {
    brand string
    Engine // Embedded (no field name)
}

// Access: Promoted to the top level
myCar.Horsepower = 300

```

### Comparison Table

| Feature         | Nested Struct                         | Embedded Struct                 |
| :-------------- | :------------------------------------ | :------------------------------ |
| **Declaration** | Has a field name (e.g., `eng Engine`) | No field name (e.g., `Engine`)  |
| **Access**      | Chained: `myCar.eng.Field`            | Direct: `myCar.Field`           |
| **Purpose**     | Clear separation and grouping         | Data-only inheritance/promotion |

-----

## 4\. Struct Methods

### Core Concepts

Go supports methods on structs, which are functions with a **receiver**. While Go doesn't have classes, methods provide similar functionality.

### Syntax & Examples

A receiver is a special parameter that appears before the function name.

``` go
type rect struct {
    width  int
    height int
}

// (r rect) is the receiver
func (r rect) area() int {
    return r.width * r.height
}

r := rect{width: 5, height: 10}
fmt.Println(r.area()) // Prints 50

```

### Deep Explanation

  * **Naming Convention**: By convention, the receiver name is usually the first letter of the struct type (e.g., `r` for `rect`).
  * **Importance**: Receivers allow structs to implement **interfaces**, which is fundamental to Go's design.

-----

## 5\. Memory Layout and Optimization

### Core Concepts

In Go, structs occupy a **contiguous block** of memory, with fields placed in the order they are defined.

### Key Technical Details

  * **Field Alignment**: Go adds "padding" (wasted space) between fields to align them with memory boundaries for faster execution.
  * **Order Matters**: The sequence of fields can significantly impact memory usage.

#### Example: Optimized vs. Unoptimized

**Poorly Designed (6 bytes):**

``` go
type stats struct {
    NumPosts uint8  // 1 byte
    // 1 byte padding (wasted)
    Reach    uint16 // 2 bytes
    NumLikes uint8  // 1 byte
    // 1 byte padding (wasted)
}

```

**Optimized (4 bytes):**

``` go
type stats struct {
    Reach    uint16 // 2 bytes
    NumPosts uint8  // 1 byte
    NumLikes uint8  // 1 byte
}

```

### Key Takeaways

  * To save memory, align fields from **largest to smallest**.
  * Use the `reflect` package to debug memory usage: `reflect.TypeOf(s).Size()`.

-----

## 6\. The Empty Struct

### Core Concepts

The empty struct `struct{}` is the smallest possible type in Go, taking up **zero bytes** of memory.

### Frequently Used Patterns

  * **Unary values**: Used as a placeholder when you only care about the existence of a key (e.g., in maps or channels).
  * **Memory Efficiency**: Smaller than a `bool` (1 byte) or `int` (8 bytes).

-----

## Final Overall Golang Struct Cheat Sheet

| Concept          | Syntax Example                     | Key Fact                          |
| :--------------- | :--------------------------------- | :-------------------------------- |
| **Named Struct** | `type User struct { Name string }` | Reusable and standard.            |
| **Anonymous**    | `var s = struct{ A int }{ A: 1 }`  | Used for one-time logic/JSON.     |
| **Nested**       | `type A struct { b B }`            | Accessed via `a.b.Field`.         |
| **Embedded**     | `type A struct { B }`              | Fields are promoted to `a.Field`. |
| **Methods**      | `func (s S) Run() {}`              | Functions with a receiver.        |
| **Empty Struct** | `struct{}{}`                       | Occupies **0 bytes**.             |
| **Optimization** | `Largest -> Smallest fields`       | Minimizes wasted padding space.   |

-----

### Summary Notes

Structs are the building blocks of data in Go. While Go avoids traditional inheritance, it uses **embedding** for composition and **methods** for behavior. Understanding **memory alignment** and **field promotion** is crucial for writing efficient, idiomatic Go code.
