## Lecture 1: Introduction to Go & The Compilation Process

### Core Concepts

Go is a **compiled, statically typed** language designed for speed and developer productivity. Unlike interpreted languages (like Python or JavaScript), Go code must be transformed into machine code before it can run.

### Syntax & Examples

A standard Go program structure includes a package declaration, imports, and a `main` function.

``` go
package main // Tells the compiler this is a standalone executable

import "fmt" // Imports the formatting package for I/O

func main() { // The entry point of every Go program
    fmt.Println("hello world")
}

```

### Deep Explanation

  * **The Compiler's Job:** The Go compiler takes your high-level code and turns it into an executable file (e.g., an `.exe` file on Windows). Once compiled, you can run the program on a machine without needing the original source code or the Go compiler installed.
  * **Speed Comparison:**
      * Go is significantly faster than interpreted languages like Python and Ruby.
      * While slightly slower than C++ or Rust due to its automated memory management, it compiles much faster, making the development cycle very efficient.

### Key Takeaways

  * Go code is compiled directly to machine code.
  * `package main` is required for programs intended to run as standalone executables.
  * Compiled programs are highly portable and performant.

  
-----

## Lecture 2: Variables & Types

### Core Concepts

Go is **statically typed**, meaning variable types are known before the code runs. This allows the compiler to catch type errors early, making development faster and more reliable.

### Syntax & Examples

Go offers multiple ways to declare variables:

**1. The "Standard" Way:**

``` go
var mySkillIssues int = 67

```

**2. The "GOATed" Way (Short Variable Declaration):**

``` go
mySkillIssues := 67 // Uses the "walrus operator"

```

**3. Multiple Declarations:**

``` go
mileage, company := 80276, "Toyota" // Declare multiple variables on one line

```

### Deep Explanation

  * **Type Inference:** When using `:=`, Go automatically determines the type based on the value assigned (e.g., `67` becomes an `int`, `3.14` becomes a `float64`).
  * **The Walrus Operator (`:=`):** This is the idiomatic way to declare variables inside functions. **Constraint:** It cannot be used outside of a function (global/package scope).
  * **Strong Typing:** You cannot perform operations on mismatched types (e.g., adding an `int` to a `float64` or `string`) without explicit conversion.

### Key Takeaways

  * Use `:=` whenever possible inside functions.
  * Go will catch type mismatches during **Compile Time**.
  * You can declare multiple variables of different types on one line.

### Common Mistakes

  * **Redeclaring with `:=`:** You cannot use `:=` to update an existing variable; use `=` instead.
  * **Global Walrus:** Attempting to use `:=` outside of a function will cause a compile error.


-----

## Lecture 3: Advanced Types - Numbers, Runes, & Strings

### Core Concepts

Go provides granular control over numeric types and treats strings as sequences of bytes. It uses **Runes** to handle complex Unicode characters (like Emojis).

These are different type sizes in Go: Signed integers - no decimal (int, int8, int16, int32, int64), Unsigned integers - non-negative numbers/no decimal (uint, uint8, uint16, uint32, uint64, uintptr), Signed decimal numbers (float32, float64), Complex numbers - a complex number has a real and imaginary part (complex64, complex128)

Unless you have a good performance related reason, we'll typically just want to use the "default" types: `bool`, `string`, `int`, `uint`, `byte`, `rune (int 32)`, `float64`, `complex128`

`int32`: Used as a general-purpose 32-bit signed integer for holding raw numbers.
`rune`: Used specifically to signify that the integer represents a single Unicode code point (representing a character or symbol).

### Syntax & Examples

**Type Conversion:**

``` go
var myAge uint16 = 25
myAgeInt := int(myAge) // Explicit conversion required

```

**Runes vs. Bytes:**

``` go
const name = "Bear 🐻"
// Byte length might be different from character (rune) length for Unicode
runeLength := utf8.RuneCountInString(name) 
```

### Deep Explanation

  * **Default Types:** Unless you have a specific performance reason, use `bool`, `string`, `int`, `uint`, `byte`, `rune (int 32)`, `float64`, `complex128`.
  * **Numeric Sizes:** Types like `int8` or `uint32` specify how many bits are used in memory. `int` and `uint` sizes vary (32 or 64-bit) depending on the environment.
  * **Runes:** A `rune` is an alias for `int32`. It represents a single Unicode character, which is essential because characters like emojis take up multiple bytes.
  * **Bytes:** A `byte` is an alias for `uint8` and represents a single 8-bit value (0-255).

### Key Takeaways

  * **Default to `int` and `float64`** for most tasks.
  * **Use `rune`** when you need to process individual characters in a string.
  * **Type conversion** must be explicit; Go does not automatically convert between types.

### Common Mistakes

  * **Ignoring Unicode:** Assuming 1 byte = 1 character. This breaks when using non-English characters or emojis.
  * **Implicit Conversion:** Trying to pass a `uint16` into a function expecting an `int` without converting it first.

  
-----

## Lecture 4: Runtime, Memory, & Comments

### Core Concepts

The **Go Runtime** is a small piece of code included in every executable that manages the program while it's running, most notably through **Garbage Collection**.

### Syntax & Examples

**Comments:**

``` go
// This is a single-line comment

/* 
   This is a 
   multi-line comment 
*/

```

### Deep Explanation

  * **Memory Management:** Go uses a **Garbage Collector** to automatically free up memory that is no longer being used. This makes it safer and easier than C, though slightly heavier in memory than Rust.
  * **Runtime vs. Compile Time:**
      * **Compile Time:** Syntax and type checking occur here.
      * **Run Time:** Value calculations and errors like "division by zero" occur while the program is running.
  * **Efficiency:** Go programs typically use less memory than Java because they don't require a heavy Virtual Machine (JVM).

### Key Takeaways

  * Go handles memory cleanup automatically via the runtime's Garbage Collector.
  * Comments are for documentation and do not execute.
  * Go is more memory-efficient than Java but less so than C/Rust.

-----

## Summary Notes

### Final Overall Golang Cheat Sheet

| Concept           | Syntax / Rule                                                     |
| :---------------- | :---------------------------------------------------------------- |
| **Declaration**   | `name := "Kien"` (inside functions) or `var name string = "Kien"` |
| **Multiple Dec.** | `x, y := 1, 2`                                                    |
| **Comments**      | `//` for line, `/* */` for block                                  |
| **Core Types**    | `int`, `float64`, `bool`, `string`, `rune` (int32)                |
| **Structure**     | `package main` -\> `import` -\> `func main()`                     |
| **Type Conv.**    | Must be explicit: `int(someFloat)`                                |

### Roadmap of Concepts Covered

1.  **Environment:** Compilation, executables, and the Go Runtime.
2.  **Syntax:** Variable declarations (`var` vs `:=`) and comments.
3.  **Data:** Static typing, type inference, and numeric sizes.
4.  **Complex Data:** Byte vs. Rune and UTF-8 string encoding.
