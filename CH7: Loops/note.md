# Golang Lecture Notes: Chapter 7 - Loops

This study guide transforms the Chapter 7 lecture slides into a professional markdown resource. It covers Go's streamlined approach to iteration, where the `for` keyword handles all looping logic.

-----

## 1\. The Standard For Loop

### Lecture Title: Basic Loop Syntax

Go uses a single keyword, `for`, to handle all types of iteration. The basic syntax is inherited from the C family but without the parentheses.

### Syntax & Examples

``` go
for INITIAL; CONDITION; AFTER {
    // do something
}

```

  * **INITIAL**: Executed once at the beginning. You can declare variables here that are scoped only to the loop.
  * **CONDITION**: Evaluated before every iteration. If it evaluates to `false`, the loop terminates.
  * **AFTER**: Executed at the end of every iteration.

**Example: Printing 0 through 9**

``` go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

```

### Key Takeaways

  * Variables declared in the `INITIAL` statement are not accessible outside the loop.
  * Go does not use parentheses around the loop components.

-----

## 2\. Flexible Loop Structures

### Lecture Title: Omitting Conditions & "While" Loops

Go is unique because it lacks a `while` keyword. Instead, you create different loop behaviors by omitting parts of the `for` statement.

### Syntax & Examples

**The "While" Loop Equivalent**
If you only provide a `CONDITION`, the loop behaves exactly like a `while` loop in other languages.

``` go
plantHeight := 1
for plantHeight < 5 {
    fmt.Println("still growing! current height:", plantHeight)
    plantHeight++
}

```

**Infinite Loops**
Omitting all components (or just the condition) creates a loop that runs forever until manually stopped or broken.

``` go
for {
    // do something forever
}

```

### Deep Explanation

By using only one keyword for all iteration types, Go simplifies the language grammar. Whether you need a counter-based loop, a condition-based loop, or an infinite background worker, `for` is always the answer.

-----

## 3\. Loop Control: Break & Continue

### Lecture Title: Changing Control Flow

To manage complex logic within a loop, Go provides `break` and `continue` to exit or skip iterations.

### Syntax & Examples

**Continue**
The `continue` keyword stops the current iteration and jumps immediately to the `AFTER` statement and then the next `CONDITION` check.

``` go
for i := 0; i < 10; i++ {
    if i % 2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Prints: 1, 3, 5, 7, 9
}

```

**Break**
The `break` keyword immediately terminates the loop and moves execution to the line following the loop's closing brace.

``` go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Stop the loop entirely when i is 5
    }
    fmt.Println(i) // Prints: 0, 1, 2, 3, 4
}

```

### Key Takeaways

  * **Guard Clause Pattern**: Use `continue` to skip invalid data early in a loop to avoid deeply nested `if` statements.

-----

## 4\. Logical Operators in Loops

### Lecture Title: FizzBuzz Logic

Complex loops often require combining multiple conditions using logical operators.

### Syntax & Examples

| Operator | Meaning            | Example                    |
| :------- | :----------------- | :------------------------- |
| `%`      | Modulo (Remainder) | `7 % 3` is `1`             |
| `&&`     | Logical AND        | `true && false` is `false` |

| `||` | Logical OR | `true || false` is `true` |

-----

## Practical Application: Connection Accumulator

### Lecture Title: Exercise - Counting Connections

This exercise demonstrates using a loop to accumulate a value over time, simulating how connections grow as more people join a group chat.

``` go
func countConnections(groupSize int) int {
    res := 0
    // We start from 1 because the first person adds 0 connections
    for i := 1; i <= groupSize; i++ {
        // Every new person 'i' creates (i-1) new connections
        temp := i - 1
        res += temp
    }
    return res
}

```

### Summary Notes

Loops in Go are intentionally simple. By mastering the `for` keyword and its various configurations (standard, condition-only, and infinite), you can handle any iterative task in the language.

-----

## Final Overall Golang Loops Cheat Sheet

| Feature           | Syntax                                |
| :---------------- | :------------------------------------ |
| **Standard Loop** | `for i := 0; i < 10; i++ { ... }`     |
| **While-style**   | `for condition { ... }`               |
| **Infinite**      | `for { ... }`                         |
| **Skip Current**  | `continue`                            |
| **Exit Loop**     | `break`                               |
| **Modulo**        | `i % n == 0` (check for divisibility) |

### Common Mistakes

  * **Infinite Loops**: Forgetting to update the variable in a "while-style" loop (e.g., forgetting `i++`), causing the program to hang.
  * **Condition Errors**: Off-by-one errors in the `CONDITION` (using `<` when you mean `<=`).
  * **Panic on Break**: Trying to `break` outside of a loop context.
