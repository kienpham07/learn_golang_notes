# Golang Lecture Notes: Foundations & Conditionals

The following notes cover essential Golang concepts including variable types, string formatting, and control structures. This guide is designed for beginners to build a solid foundation in Go's unique syntax and behavior.

-----

## Lecture 1: Runes and String Encoding

### Core Concepts

Go handles text differently than languages like C. While a "character" in older languages is often just a single byte (ASCII), Go uses **UTF-8 encoding**, allowing it to support a vast range of symbols and languages like emojis or Chinese characters.

  * **String**: A sequence of bytes.
  * **Rune**: An alias for `int32`. It represents a single **Unicode code point** (a conceptual character).

### Syntax & Examples

``` go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    name := "🐻" // Bear emoji
    
    // len() counts bytes
    fmt.Printf("Byte length: %d\n", len(name)) 
    
    // utf8.RuneCountInString() counts actual characters
    fmt.Printf("Rune length: %d\n", utf8.RuneCountInString(name))
}

```

### Deep Explanation

Go’s `len()` function does not count "letters"; it strictly counts **bytes**.

  * **ASCII characters** (like 'a', 'b', '1') take up **1 byte** each. For these, byte length and rune length are the same.
  * **Complex characters** (like 🐻) can take up multiple bytes; the bear emoji specifically requires **4 bytes** (Hex: `F0 9F 90 BB`).

| String  | Character(s)  | Hex Bytes (UTF-8)  | Total Bytes | Total Runes |
| :------ | :------------ | :----------------- | :---------- | :---------- |
| "🐻"     | 🐻             | F0 9F 90 BB        | 4           | 1           |
| "boots" | b, o, o, t, s | 62, 6F, 6F, 74, 73 | 5           | 5           |

### Key Takeaways

  * Use `len()` when you need to know memory usage (bytes).
  * Use `utf8.RuneCountInString()` when you need to know the actual number of characters.

-----

## Lecture 2: Computed Constants

### Core Concepts

Constants in Go must be known at **compile time**. Unlike JavaScript, you cannot assign a value to a constant that is only determined when the program is running (run-time).

### Syntax & Examples

**Valid (Compile-time):**

``` go
const firstName = "Lane"
const lastName = "Wagner"
// Constants can be computed from other constants
const fullName = firstName + " " + lastName

```

**Invalid (Run-time):**

``` go
// Error: time.Now() is only known when the program runs
const currentTime = time.Now()

```

### Summary Notes

Constants are for values that are truly "constant" and fixed before the code ever starts executing.

-----

## Lecture 3: Formatting Strings

### Core Concepts

Go uses `fmt.Printf` and `fmt.Sprintf` for string interpolation, following the C "printf" tradition.

  * `fmt.Printf()`: Prints the formatted string to standard out.
  * `fmt.Sprintf()`: Returns the formatted string as a variable instead of printing it.

### Syntax & Formatting Verbs

| Verb   | Purpose                                   | Example                                     |
| :----- | :---------------------------------------- | :------------------------------------------ |
| `%v`   | **Default**: Catch-all for any value      | `fmt.Sprintf("%v", 10)` -\> "10"            |
| `%s`   | **String**                                | `fmt.Sprintf("%s", "Hi")` -\> "Hi"          |
| `%d`   | **Integer** (Decimal)                     | `fmt.Sprintf("%d", 10)` -\> "10"            |
| `%f`   | **Float**                                 | `fmt.Sprintf("%f", 10.523)` -\> "10.523000" |
| `%.2f` | **Float (Rounded)**: Rounds to 2 decimals | `fmt.Sprintf("%.2f", 10.523)` -\> "10.52"   |

-----

## Lecture 4: Conditionals (If/Else)

### Core Concepts

Go's `if` statements omit parentheses around the condition for a cleaner syntax.

### Syntax & Examples

``` go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}

```

### Common Mistakes

  * **Parentheses**: Do not put parentheses around your condition (e.g., `if (x > 1)` is incorrect).
  * **Brace Placement**: The opening brace `{` **must** be on the same line as the condition.

-----

## Lecture 5: The Initial Statement

### Core Concepts

An `if` block can include an "initial" statement before the condition. Variables created here are only defined within the scope of the `if` body.

### Syntax & Examples

**Standard Way:**

``` go
length := getLength(email)
if length < 10 {
    fmt.Printf("Email is too short: %d", length)
}
// 'length' is still available here

```

**Go "Initial Statement" Way:**

``` go
// if INITIAL_STATEMENT; CONDITION { ... }
if length := getLength(email); length < 10 {
    fmt.Printf("Email is too short: %d", length)
}
// 'length' is NOT available here (Limited Scope)

```

### Key Takeaways

1.  **Conciseness**: It makes the code shorter.
2.  **Scope Limitation**: It prevents accidental use of variables elsewhere in the function.

-----

## Lecture 6: Switch and Fallthrough

### Core Concepts

By default, Go `switch` statements break automatically after a match. To continue to the next case, use the `fallthrough` keyword.

### Syntax & Examples

``` go
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"
    case "macOS":
        fallthrough // Continues to next case
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve" 
    default:
        creator = "Unknown"
    }
    return creator
}

```

-----

## Lecture 7: The Multi-line Trap (Semicolon Insertion)

### Core Concepts

Go uses a "Semicolon Insertion Rule" where the compiler automatically adds hidden semicolons at the end of lines if they look like complete statements. This can break multi-line conditions if not handled carefully.

### Syntax & Examples

**The Mistake (Syntax Error):**
Putting logical operators at the start of a new line causes Go to insert a semicolon at the end of the previous line, breaking the condition.

``` go
if conditionA
|| conditionB { // ERROR: Go inserts ';' after conditionA
    ...
}

```

**The Fix:**
Always leave the operator (`||`, `&&`) at the end of the line to signal that the statement continues.

``` go
if conditionA ||
conditionB { // CORRECT
    ...
}

```

### Common Mistakes

  * Breaking a line before an operator.
  * Placing the opening brace `{` on a new line (which also triggers automatic semicolon insertion).

### Summary Notes

To avoid syntax errors, ensure that any line intended to continue to the next one ends with an operator or a comma, signaling to the compiler that the statement is incomplete.

-----

## 🚀 Final Golang Cheat Sheet

### Variables & Types

  * `rune`: Alias for `int32`, used for single Unicode characters.
  * `len(str)`: Returns byte count.
  * `utf8.RuneCountInString(str)`: Returns actual character count.

### Constants

  * Must be known at **compile time**.
  * Cannot be set to run-time values like `time.Now()`.

### String Formatting (`fmt`)

  * `%v`: Default format catch-all.
  * `%d`: Integer; `%s`: String; `%f`: Float.
  * `%.2f`: Float rounded to 2 decimal places.

### Control Flow

  * **If**: No parentheses; opening brace `{` must stay on the same line.
  * **Initial Statement**: `if val := getVal(); val > 0 { ... }` limits variable scope.
  * **Switch**: Auto-breaks; use `fallthrough` to chain cases.
  * **Multi-line**: Keep operators (`||`, `&&`) at the end of the line to avoid the "Semicolon Trap".
