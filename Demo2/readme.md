# Collection Processing Challenge: Season Statistics Analyzer

This lesson builds upon Demo 1 and introduces you to working with collections using **for loops**, **for-range**, **maps**, and **arrays/slices**.

## Problem Statement

You are building a **Season Statistics Analyzer** that processes multiple months and provides statistical insights about seasons. Implement the following functions:

1. **`CountSeasons`**: Count how many months belong to each season
2. **`GetSeasonMonths`**: Get all months that belong to a specific season
3. **`ProcessMonthlyData`**: Process an array of month numbers and return season statistics

### Requirements

#### Function 1: `CountSeasons`
- **Input**: A slice of month numbers (`[]int`)
- **Output**: A `map[string]int` where keys are season names and values are counts
- **Implementation**: Use a **for-range** loop to iterate through the input slice
- **Behavior**: Count how many months in the input belong to each season (Winter, Spring, Summer, Autumn)
- **Invalid months**: Ignore any month number that is less than 1 or greater than 12

#### Function 2: `GetSeasonMonths`
- **Input**: A season name (`string`)
- **Output**: A slice of month numbers (`[]int`) that belong to that season
- **Implementation**: Use a **for-range** loop to iterate through a map
- **Behavior**: Return all month numbers that belong to the given season
- **Invalid season**: Return an empty slice if the season name is invalid

#### Function 3: `ProcessMonthlyData`
- **Input**: A slice of month numbers (`[]int`)
- **Output**: A `map[string][]int` where keys are season names and values are slices of month numbers
- **Implementation**: Use a **for loop** (traditional indexed loop) to iterate through the input
- **Behavior**: Group all months by their respective seasons
- **Invalid months**: Ignore any month number that is less than 1 or greater than 12

### Season Mappings

- **Winter**: 1, 2, and 12
- **Spring**: 3, 4, and 5
- **Summer**: 6, 7, and 8
- **Autumn**: 9, 10, and 11

### Sample Input/Output

#### Example 1: `CountSeasons`

**Input:**
```go
months := []int{1, 3, 6, 9, 2, 4, 7, 10, 12, 5}
```

**Output:**
```go
map[string]int{
    "Winter": 3,
    "Spring": 3,
    "Summer": 2,
    "Autumn": 2,
}
```

#### Example 2: `GetSeasonMonths`

**Input:**
```go
season := "Spring"
```

**Output:**
```go
[]int{3, 4, 5}
```

#### Example 3: `ProcessMonthlyData`

**Input:**
```go
months := []int{1, 3, 6, 9, 2, 4, 7, 10}
```

**Output:**
```go
map[string][]int{
    "Winter": []int{1, 2},
    "Spring": []int{3, 4},
    "Summer": []int{6, 7},
    "Autumn": []int{9, 10},
}
```

### Helper Function

You may want to create a helper function `getSeason(month int) string` that returns the season name for a given month number. This can reuse logic from Demo 1.

### Notes

- Use **for-range** for `CountSeasons` and `GetSeasonMonths`
- Use a traditional **for loop** (with index) for `ProcessMonthlyData`
- Create a **map** to store season-to-months mapping for `GetSeasonMonths`
- Handle invalid inputs gracefully (return empty slices/maps or ignore invalid months)
- All season names should be capitalized: "Winter", "Spring", "Summer", "Autumn"

**Requirements:**

- Implement `CountSeasons` using for-range loop
- Implement `GetSeasonMonths` using for-range loop with a map
- Implement `ProcessMonthlyData` using traditional for loop
- Handle invalid month numbers (ignore them)
- Handle invalid season names (return empty slice)

Try to implement all three functions below. Good Luck!

