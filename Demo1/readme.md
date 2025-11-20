# Control Structure Challenge: Season of a Month

This lesson brings you a challenge to solve.

## Problem Statement

Write **two functions** that determine the season based on a month number:

1. **`SeasonSwitch`**: Implement using the **switch-case** construct
2. **`SeasonIfElse`**: Implement using the **if-else** construct

Both functions should have an input parameter of a month number and return the name of the season to which this month belongs (disregard the day in the month). Make sure you follow the following criteria of month names and their values:

### Month Names and Values

- January: 1
- February: 2
- March: 3
- April: 4
- May: 5
- June: 6
- July: 7
- August: 8
- September: 9
- October: 10
- November: 11
- December: 12

### Season Mappings

- **Winter**: 1, 2, and 12
- **Spring**: 3, 4, and 5
- **Summer**: 6, 7, and 8
- **Autumn**: 9, 10, and 11

### Input/Output

- **Input**: An `int` variable
- **Output**: A `string` variable

### Sample Input

```text
3
```

### Sample Output

```text
Spring
```

### Notes

If the user enters the wrong month value (less than 1 and greater than 12), then simply return "Season unknown".

**Requirements:**

- Implement `SeasonSwitch` using switch-case construct
- Implement `SeasonIfElse` using if-else construct
- Both functions should handle invalid input (month < 1 or month > 12) by returning "Season unknown"

Try to implement both functions below. Feel free to view the solution, after giving it a few shots. Good Luck!
