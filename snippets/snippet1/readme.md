# Methods vs Functions: Understanding Receivers in Go

This example demonstrates the key differences between **methods** and **functions** in Go, and how **receivers** work.

## Overview

In Go, you can define operations on types in two ways:
1. **Methods** - Functions attached to a type using receivers
2. **Functions** - Standalone functions that are not attached to any type

## Code Structure

This example uses a `TwoInts` struct with two integers (`a` and `b`) and demonstrates both methods and functions that operate on this struct.

## Methods (With Receivers)

### Definition
Methods are functions that have a **receiver** - they belong to a specific type.

**Syntax:**
```go
func (receiver Type) MethodName(parameters) returnType {
    // method body
}
```

### Examples in This Code

#### Method 1: `AddThem()`
```go
func (tn *TwoInts) AddThem() int {
    return tn.a + tn.b
}
```

- **Receiver**: `(tn *TwoInts)` - pointer receiver
- **Call**: `two1.AddThem()` - called on an instance
- **Returns**: Sum of `a` and `b`

#### Method 2: `AddToParam()`
```go
func (tn *TwoInts) AddToParam(param int) int {
    return tn.a + tn.b + param
}
```

- **Receiver**: `(tn *TwoInts)` - pointer receiver
- **Call**: `two1.AddToParam(20)` - called on an instance with parameter
- **Returns**: Sum of `a`, `b`, and the parameter

## Functions (Without Receivers)

### Definition
Functions are standalone operations that are not attached to any type. They take parameters and return values.

**Syntax:**
```go
func FunctionName(parameters) returnType {
    // function body
}
```

### Examples in This Code

#### Function 1: `CalculateSum()`
```go
func CalculateSum(tn TwoInts) int {
    return tn.a + tn.b
}
```

- **No Receiver**: This is a standalone function
- **Call**: `CalculateSum(two2)` - called directly, passing struct as argument
- **Returns**: Sum of `a` and `b`

#### Function 2: `Multiply()`
```go
func Multiply(tn TwoInts) int {
    return tn.a * tn.b
}
```

- **No Receiver**: This is a standalone function
- **Call**: `Multiply(two2)` - called directly, passing struct as argument
- **Returns**: Product of `a` and `b`

## Receivers Explained

### What is a Receiver?

A receiver is a parameter that appears between the `func` keyword and the function name. It "attaches" the function to a type, making it a method.

### Types of Receivers

#### 1. Pointer Receiver `(*Type)`
```go
func (tn *TwoInts) AddThem() int
```

**Characteristics:**
- Uses `*Type` syntax
- Works on the original instance (can modify it)
- More efficient for large structs (doesn't copy)
- Can modify the struct's fields
- Example: `(tn *TwoInts)`

**When to Use:**
- When you need to modify the struct
- When working with large structs (performance)
- When you want to ensure you're working with the same instance

#### 2. Value Receiver `(Type)`
```go
func (tn TwoInts) SomeMethod() int
```

**Characteristics:**
- Uses `Type` syntax (no asterisk)
- Works on a copy of the instance
- Cannot modify the original struct
- Creates a copy each time called
- Example: `(tn TwoInts)`

**When to Use:**
- When you don't need to modify the struct
- When working with small structs
- When you want to ensure immutability

## Key Differences: Methods vs Functions

| Aspect | Method | Function |
|--------|--------|----------|
| **Definition** | `func (receiver) MethodName()` | `func FunctionName(params)` |
| **Receiver** | Yes - attached to type | No - standalone |
| **Call Syntax** | `instance.Method()` | `FunctionName(args)` |
| **Belongs To** | A specific type | Package |
| **Access to Data** | Through receiver | Through parameters |
| **Example** | `two1.AddThem()` | `CalculateSum(two2)` |

## Usage Examples

### Calling Methods
```go
two1 := new(TwoInts)
two1.a = 12
two1.b = 10

// Method calls - called on instance
sum := two1.AddThem()              // Returns 22
sumWithParam := two1.AddToParam(20) // Returns 42
```

### Calling Functions
```go
two2 := TwoInts{3, 4}

// Function calls - pass struct as argument
sum := CalculateSum(two2)    // Returns 7
product := Multiply(two2)    // Returns 12
```

## When to Use Methods vs Functions

### Use Methods When:
- ✅ The operation is closely related to the type
- ✅ You want object-oriented behavior
- ✅ The operation naturally belongs to the type
- ✅ You want to chain operations: `obj.Method1().Method2()`
- ✅ You want to implement interfaces

### Use Functions When:
- ✅ The operation is general-purpose
- ✅ The operation doesn't naturally belong to a specific type
- ✅ You want to work with multiple types
- ✅ The operation is a utility function
- ✅ You want to keep the code more functional

## Important Notes

### Pointer Receiver Behavior
- Both value and pointer types can call methods with pointer receivers
- Go automatically handles the conversion:
  ```go
  two1 := TwoInts{12, 10}        // value
  two2 := &TwoInts{12, 10}       // pointer
  
  two1.AddThem()  // Works! Go converts to pointer automatically
  two2.AddThem()  // Works! Already a pointer
  ```

### Value Receiver Behavior
- Methods with value receivers work on copies
- Modifications won't affect the original:
  ```go
  func (tn TwoInts) SetA(newA int) {
      tn.a = newA  // This won't modify the original!
  }
  ```

### Function Calls
- Functions require explicit passing of the struct:
  ```go
  two := TwoInts{12, 10}
  result := CalculateSum(two)  // Must pass struct explicitly
  ```

## Summary

- **Methods** have receivers and are called on instances: `instance.Method()`
- **Functions** are standalone and called directly: `FunctionName(args)`
- **Pointer receivers** `(*Type)` can modify the struct and are more efficient
- **Value receivers** `(Type)` work on copies and cannot modify the original
- Choose methods for type-specific operations, functions for general-purpose operations

## Code Output

When you run this code, you'll see:
```
The sum is: 22
Add them to the param: 42
The sum is: 7
Function call - Sum: 7
Function call - Multiply: 12
```

This demonstrates both methods and functions working with the same `TwoInts` struct, showing how they can achieve similar results with different approaches.

