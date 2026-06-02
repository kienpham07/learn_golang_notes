# Golang Lecture Notes: Chapter 6 - Errors

This comprehensive guide transforms the Chapter 6 lecture slides into a structured study resource. It covers Go's unique approach to error handling, the built-in error interface, custom error types, and the crucial "panic vs. error" distinction.

-----

## 1\. The Error Interface

### Core Concepts

In Go, error handling is fundamentally different from languages like Python (exceptions) or JavaScript (try/catch). Go expresses errors as **values**. An "error" is simply any type that implements the built-in `error` interface.

### Syntax & Examples

The `error` interface is one of the simplest and most powerful in Go:

``` go
type error interface {
    Error() string
}

```

**Standard Usage Pattern:**
When a function can fail, it should return an `error` as its **last** return value. The caller then checks if that error is `nil`.

``` go
// Example using strconv.Atoi
i, err := strconv.Atoi("42b")
if err != nil {
    // Handle the error: "42b" is not a valid integer
    fmt.Println("couldn't convert: ", err)
    return
}
// If err is nil, 'i' is safe to use

```

### Deep Explanation

  * **Unique Philosophy:** Go doesn't have special keywords like `throw` or `catch`. Errors are just another piece of data you pass around and inspect.
  * **Underlying Type:** The underlying type of an error is always an **Interface**.
  * **Multiple Interfaces:** A type can be an error and also fulfill other interfaces simultaneously.

### Key Takeaways

  * Functions return `(value, error)`.
  * Always check `if err != nil` before proceeding.
  * The `.Error()` method provides the string description of the failure.

-----

## 2\. Custom Error Types

### Core Concepts

Because `error` is just an interface, you can create your own structs to represent specific failure states. This allows you to carry more metadata than a simple string.

### Syntax & Examples

``` go
type userError struct {
    name string
}

// Fulfilling the error interface
func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        // Return your custom struct as an error
        return userError{name: userName}
    }
    return nil
}

```

### Deep Explanation

Using custom structs for errors is powerful because you can use **Type Assertions** (covered in previous chapters) to check for specific error types and access their fields, such as the `name` field in the example above.

-----

## 3\. The "errors" Package & fmt.Errorf

### Core Concepts

Go provides two primary ways to create "on-the-fly" errors without defining a custom struct: `errors.New()` and `fmt.Errorf()`.

### Syntax & Examples

**The `errors` Package:**
Used for simple, static error messages.

``` go
import "errors"

var err error = errors.New("no dividing by 0")

```

**The `fmt.Errorf` Function:**
The "GOATed" (Greatest of All Time) way to create errors because it supports dynamic formatting and error wrapping.

``` go
// Wrapping an error with more context
return fmt.Errorf("failed to get user: %w", err)

```

### Comparison Table

| Feature            | `errors.New()`         | `fmt.Errorf()`                  |
| :----------------- | :--------------------- | :------------------------------ |
| **Primary Use**    | Static sentinel errors | Dynamic or wrapped errors       |
| **Formatting**     | No specifiers allowed  | Supports `%s`, `%d`, `%v`, etc. |
| **Error Wrapping** | Cannot wrap            | Wraps natively using `%w`       |
| **Package**        | `errors`               | `fmt`                           |

-----

## 4\. Panic, Recover, and log.Fatal

### Core Concepts

Go provides a `panic` function that crashes the program and prints a stack trace. However, it is **not** the intended way to handle normal errors.

### Deep Explanation

  * **Panic:** It "yeets" control out of the current function and up the call stack until it hits a `recover` or crashes the entire program.
  * **Recover:** Used inside a `defer` block to stop a panic and regain control.
  * **The "Anti-Pattern" Warning:** New developers often mistake `panic/recover` for `try/catch`. **Don't do this**. It is considered a "truly astonishingly bad way" to handle errors in Go.

### Key Takeaways

  * **Rule:** Use error values for all "normal" error handling.
  * **Panic Usage:** Basically **never**.
  * **Better Alternative:** If an error is truly unrecoverable and the program must stop, use `log.Fatal()` to print a message and exit cleanly.

-----

## Final Overall Golang Errors Cheat Sheet

| Concept          | Syntax Example                            | When to Use                                      |
| :--------------- | :---------------------------------------- | :----------------------------------------------- |
| **Interface**    | `type error interface { Error() string }` | The foundation of all Go errors.                 |
| **Basic Error**  | `errors.New("message")`                   | For static, simple errors.                       |
| **Formatted**    | `fmt.Errorf("user %d: %w", id, err)`      | Adding context or wrapping errors.               |
| **Custom Error** | `type myErr struct { code int }`          | When you need extra metadata.                    |
| **Panic**        | `panic("boom")`                           | Almost never; only for truly fatal logic errors. |
| **Clean Exit**   | `log.Fatal(err)`                          | When you want to log and stop the program.       |

### Common Interview Questions

  * **Q: Why doesn't Go use try/catch?**
      * *A:* Go treats errors as values to encourage developers to handle failures explicitly as part of the normal control flow, making code more predictable.
  * **Q: What is the significance of the `%w` verb in `fmt.Errorf`?**
      * *A:* It "wraps" an underlying error, allowing the original error to be extracted later while providing new context.

### Summary Notes

Error handling in Go is about **dependability**. By returning errors as values and avoiding the "lazy" trap of `panic`, you create systems that are easier to debug and more resilient. Always treat errors as first-class citizens in your code logic.
