# # Pluralsight Course: Go Fundamentals by Mike Van Sickle (NOTES)

## Simple Data Types

- Strings:  
  - Ex.: `"Hello, World!"`
- Numbers: 
  - int: `int`
  - unsigned integers: `uint`
  - floating point numbers: `float32`, `float64`
  - complex numbers: `complex64`, `complex128`
- Booleans: `true` or `false`
- Errors: The error built-in type is the conventional interface for representing an error condition, with the nil value representing no error.
  - Ex.: `errors.New("Something went wrong")`

## When to use `iota`?

- When you want to create a set of related constants that increment in value.
- When you want to create a set of related constants that share the same type.
- When you want to create a set of related constants that share the same type and are initialized with a value that is incremented in value.
- When you need to create a set of related constants that share the same type and are initialized with a value that is incremented in value, but you don't want to have to manually assign the value of each constant.

Ex.:

```go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

const (
    Running = 1 << iota
    Swimming
    Cycling
    Skiing
    Canoeing
)
```


