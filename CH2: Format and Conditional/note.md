# Go Programming Lecture Notes: Chapter 1 & 2

## Lecture 1: Variables, Runes, and String Encoding

### Core Concepts

Go treats strings as sequences of bytes but introduces the `rune` type to handle modern, multi-byte character sets like Unicode and emojis.

  * **String**: A read-only slice of bytes. It can hold any data, but usually holds UTF-8 encoded text.
  * **Rune**: An alias for `int32`. It represents a single Unicode code point.
  * **UTF-8 Encoding**: A variable-length encoding where a character can take between 1 and 4 bytes.

### Syntax & Examples

``` go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    // A string containing an emoji (multi-byte character)
    const emoji = "🐻" 
    
    // len() returns the number of bytes
    fmt.Printf("Byte length: %d\n", len(emoji)) // Output: 4
    
    // RuneCountInString returns the actual character count
    fmt.Printf("Rune length: %d\n", utf8.RuneCountInString(emoji)) // Output: 1
}

```

### Deep Explanation: Bytes vs. Runes

| Character | Type  | UTF-8 Hex Bytes | Byte Count (`len`) | Rune Count |
| :-------- | :---- | :-------------- | :----------------- | :--------- |
| `b`       | ASCII | `62`            | 1                  | 1          |
| `🐻`       | Emoji | `F0 9F 90 BB`   | 4                  | 1          |

**Analogy**: Think of a string as a shelf of Lego bricks. Some "characters" (like standard letters) are just 1 brick wide. Others (like emojis) are 4 bricks wide. `len()` tells you how many bricks are on the shelf, while `RuneCountInString` tells you how many actual "toys" were built.

### Key Takeaways

  * Use `len()` for memory/byte size.
  * Use `utf8.RuneCountInString()` when you need to know the number of characters for UI or logic.

-----

## Lecture 2: Formatting Strings

### Core Concepts

Go uses the `fmt` package, following the `printf` tradition from C, to interpolate values into strings.

### Syntax & Examples

| Verb   | Description                | Example                                   |
| :----- | :------------------------- | :---------------------------------------- |
| `%v`   | Default format (catch-all) | `fmt.Sprintf("%v", 10)` -\> "10"          |
| `%s`   | String                     | `fmt.Sprintf("%s", "Go")` -\> "Go"        |
| `%d`   | Integer                    | `fmt.Sprintf("%d", 10)` -\> "10"          |
| `%f`   | Float                      | `fmt.Sprintf("%f", 10.5)` -\> "10.500000" |
| `%.2f` | Float (2 decimal places)   | `fmt.Sprintf("%.2f", 10.523)` -\> "10.52" |

  * **`fmt.Printf()`**: Prints directly to standard output.
  * **`fmt.Sprintf()`**: Returns the formatted string as a variable.

-----

## Lecture 3: Constants

### Core Concepts

Constants in Go must be known at **compile time**. They cannot be assigned values that are only determined when the program is running.

### Syntax & Examples

``` go
const firstName = "Lane"
const lastName = "Wagner"

// Valid: Computed at compile time
const fullName = firstName + " " + lastName 

// INVALID: time.Now() is a runtime function
// const currentTime = time.Now() 
```

### Common Mistakes

Trying to use functions that interact with the system or network inside a `const` declaration will cause a compiler error. Use `var` for those instead.

-----

## Lecture 4: Conditionals

### Core Concepts

Go's `if` statements are designed for readability, removing unnecessary parentheses and introducing "initial statements" to manage scope.

### Syntax & Examples

**Standard If/Else:**

``` go
if height > 6 {
    fmt.Println("Super tall")
} else if height > 4 {
    fmt.Println("Tall enough")
} else {
    fmt.Println("Too short")
}

```

  * **Note**: The opening brace `{` **must** be on the same line as the condition.

**Initial Statement (The "Go Way"):**
You can declare a variable inside the `if` block. Its scope is limited only to that block.

``` go
// length is only accessible inside this if/else block
if length := getLength(email); length < 10 {
    fmt.Printf("Email too short: %d chars", length)
}

```

### Deep Explanation: Why use Initial Statements?

1.  **Cleaner Code**: It's more concise.
2.  **Scope Limitation**: It prevents "variable pollution," ensuring you don't accidentally use that temporary variable elsewhere in your function.

-----

## Lecture 5: Switch and Fallthrough

### Core Concepts

Unlike C or Java, Go's `switch` statements **do not** fall through by default. Once a match is found, it executes and exits.

### Syntax & Examples

If you *want* to continue to the next case, you must use the `fallthrough` keyword.

``` go
switch os {
case "macOS":
    fallthrough
case "mac":
    creator = "Apple" // Both macOS and mac will land here
default:
    creator = "Unknown"
}

```

-----

## Final Summary & Roadmap

### Most Important Concepts to Master

1.  **UTF-8 Awareness**: Always remember that `len(string)` is not character count.
2.  **Scope Control**: Master the `if INITIAL; CONDITION` pattern for cleaner functions.
3.  **Compile-time Constants**: Understand the difference between what the compiler knows and what the machine knows at runtime.

### Golang Cheat Sheet

  * **Print**: `fmt.Printf("%v", val)`
  * **Format String**: `s := fmt.Sprintf("%s", str)`
  * **Character Count**: `utf8.RuneCountInString(s)`
  * **Conditionals**: No parentheses; brace on the same line.
  * **Switch**: No default fallthrough; use `fallthrough` if needed.
