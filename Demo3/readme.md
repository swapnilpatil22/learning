# Advanced Slice Operations Challenge: Data Manipulation Toolkit

This lesson introduces you to advanced slice operations including **copying**, **appending**, **multidimensional slices**, **inserting slices**, and **magnifying slices**.

## Problem Statement

You are building a **Data Manipulation Toolkit** that provides various operations for working with slices. Implement the following functions:

1. **`CopySlice`**: Create a deep copy of a slice
2. **`AppendToSlice`**: Append multiple elements to a slice
3. **`CreateMatrix`**: Create a 2D matrix (multidimensional slice)
4. **`InsertSlice`**: Insert a slice into another slice at a specific position
5. **`MagnifySlice`**: Expand a slice by duplicating elements

### Requirements

#### Function 1: `CopySlice`
- **Input**: A slice of integers (`[]int`)
- **Output**: A new slice (`[]int`) containing the same elements as the input
- **Implementation**: Use `copy()` function or manual copying to create an independent copy
- **Behavior**: Create a deep copy that is independent of the original slice
- **Note**: Changes to the returned slice should not affect the original slice

#### Function 2: `AppendToSlice`
- **Input**: A slice of integers (`[]int`) and a variadic list of integers (`...int`)
- **Output**: A new slice (`[]int`) with all elements appended
- **Implementation**: Use `append()` function to add elements
- **Behavior**: Append all provided elements to the end of the slice
- **Example**: `AppendToSlice([]int{1, 2}, 3, 4, 5)` returns `[]int{1, 2, 3, 4, 5}`

#### Function 3: `CreateMatrix`
- **Input**: Number of rows (`int`) and columns (`int`)
- **Output**: A 2D slice (`[][]int`) initialized with zeros
- **Implementation**: Create a multidimensional slice (slice of slices)
- **Behavior**: Create a matrix with specified dimensions, all elements initialized to 0
- **Example**: `CreateMatrix(2, 3)` returns `[][]int{{0, 0, 0}, {0, 0, 0}}`

#### Function 4: `InsertSlice`
- **Input**: Original slice (`[]int`), index position (`int`), and slice to insert (`[]int`)
- **Output**: A new slice (`[]int`) with the inserted slice at the specified position
- **Implementation**: Use `append()` and slicing to insert a slice into another slice
- **Behavior**: Insert the `toInsert` slice into `original` at the specified `index`
- **Edge cases**: 
  - If `index < 0`, insert at the beginning
  - If `index > len(original)`, append at the end
- **Example**: `InsertSlice([]int{1, 2, 3}, 1, []int{10, 20})` returns `[]int{1, 10, 20, 2, 3}`

#### Function 5: `MagnifySlice`
- **Input**: A slice of integers (`[]int`) and a multiplier (`int`)
- **Output**: A new slice (`[]int`) where each element appears 'multiplier' times
- **Implementation**: Use `append()` in a loop to expand the slice
- **Behavior**: Duplicate each element in the slice the specified number of times
- **Example**: `MagnifySlice([]int{1, 2, 3}, 2)` returns `[]int{1, 1, 2, 2, 3, 3}`
- **Edge case**: If multiplier is 0 or negative, return an empty slice

### Sample Input/Output

#### Example 1: `CopySlice`

**Input:**
```go
original := []int{1, 2, 3, 4, 5}
```

**Output:**
```go
[]int{1, 2, 3, 4, 5}
```

**Note:** The returned slice should be independent - modifying it should not affect the original.

#### Example 2: `AppendToSlice`

**Input:**
```go
slice := []int{10, 20}
elements := 30, 40, 50
AppendToSlice(slice, elements...)
```

**Output:**
```go
[]int{10, 20, 30, 40, 50}
```

#### Example 3: `CreateMatrix`

**Input:**
```go
rows := 3
cols := 4
CreateMatrix(rows, cols)
```

**Output:**
```go
[][]int{
    {0, 0, 0, 0},
    {0, 0, 0, 0},
    {0, 0, 0, 0},
}
```

#### Example 4: `InsertSlice`

**Input:**
```go
original := []int{1, 2, 3, 4}
index := 2
toInsert := []int{10, 20, 30}
InsertSlice(original, index, toInsert)
```

**Output:**
```go
[]int{1, 2, 10, 20, 30, 3, 4}
```

**Edge Case Examples:**
- `InsertSlice([]int{1, 2}, -1, []int{10})` → `[]int{10, 1, 2}` (insert at beginning)
- `InsertSlice([]int{1, 2}, 5, []int{10})` → `[]int{1, 2, 10}` (append at end)

#### Example 5: `MagnifySlice`

**Input:**
```go
slice := []int{1, 2, 3}
multiplier := 3
MagnifySlice(slice, multiplier)
```

**Output:**
```go
[]int{1, 1, 1, 2, 2, 2, 3, 3, 3}
```

**Edge Case:**
- `MagnifySlice([]int{1, 2}, 0)` → `[]int{}` (empty slice)

### Key Concepts

#### Copying Slices
- Use `copy(dst, src)` to copy elements from source to destination
- Create destination slice with `make([]int, len(src))` first
- Manual copying: iterate and assign each element

#### Appending to Slices
- Use `append(slice, elements...)` to add elements
- `append()` returns a new slice (may create a new underlying array)
- Variadic functions accept variable number of arguments using `...`

#### Multidimensional Slices
- A slice of slices: `[][]int`
- Create outer slice: `make([][]int, rows)`
- Create inner slices: `make([]int, cols)` for each row
- Access elements: `matrix[row][col]`

#### Inserting Slices
- Split original slice at index: `original[:index]` and `original[index:]`
- Combine: `append(append(original[:index], toInsert...), original[index:]...)`
- Handle edge cases for invalid indices

#### Magnifying Slices
- Use nested loops: outer loop for elements, inner loop for multiplier
- Append each element multiple times using `append()`

### Notes

- **Slice Independence**: When copying, ensure the new slice is independent of the original
- **Capacity**: `append()` may allocate a new underlying array if capacity is exceeded
- **Multidimensional Initialization**: Initialize each row separately in a loop
- **Index Bounds**: Always validate indices when inserting to avoid panics
- **Variadic Parameters**: Use `...` to accept variable number of arguments
- **Empty Slices**: Return `[]int{}` for edge cases (empty input, invalid multiplier, etc.)

**Requirements:**

- Implement `CopySlice` using `copy()` function
- Implement `AppendToSlice` using `append()` with variadic parameters
- Implement `CreateMatrix` creating a multidimensional slice
- Implement `InsertSlice` using `append()` and slicing operations
- Implement `MagnifySlice` using `append()` in nested loops
- Handle edge cases gracefully (invalid indices, empty inputs, etc.)

Try to implement all five functions below. Good Luck!

