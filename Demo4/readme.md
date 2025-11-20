# Structs, Methods, Maps, and Composition Challenge: Student Management System

This lesson introduces you to **structs**, **methods** (receivers), **maps**, the **difference between functions and methods**, and **composition** (Go's approach to multiple inheritance through embedding).

## Problem Statement

You are building a **Student Management System** that manages students and teachers. Implement structs with embedded types, methods with different receiver types, and functions that work with maps.

### Key Concepts

#### Structs
- Structs are collections of fields (like classes in other languages)
- Define custom types with named fields
- Example: `type Person struct { Name string; Age int }`

#### Methods vs Functions
- **Functions**: Standalone, not attached to any type
  - Example: `func CalculateAge(p Person) int`
- **Methods**: Attached to a type using receivers
  - Example: `func (s Student) GetInfo() string`
- Methods can have **value receivers** `(s Student)` or **pointer receivers** `(s *Student)`

#### Receivers
- **Value Receiver**: `func (s Student) Method()` - works on a copy, doesn't modify original
- **Pointer Receiver**: `func (s *Student) Method()` - works on original, can modify it
- Use pointer receivers when you need to modify the struct

#### Maps
- Maps are key-value pairs: `map[keyType]valueType`
- Create: `make(map[string]int)` or `map[string]int{}`
- Access: `value, ok := map[key]` (ok indicates if key exists)
- Modify: `map[key] = value`

#### Composition (Multiple Inheritance-like Behavior)
- Go doesn't have traditional multiple inheritance
- Use **struct embedding** (anonymous fields) for composition
- Example: `type Student struct { Person; StudentID string }`
- Embedded fields can be accessed directly: `student.Name` (not `student.Person.Name`)

### Requirements

#### Struct Definitions

##### 1. `Person` Struct
- Fields: `Name` (string), `Age` (int)
- This will be embedded in other structs

##### 2. `Student` Struct
- **Requirement**: Embed `Person` struct using anonymous field
- Additional fields: `StudentID` (string), `Grade` (float64)
- Example: `type Student struct { Person; StudentID string; Grade float64 }`

##### 3. `Teacher` Struct
- **Requirement**: Embed `Person` struct using anonymous field
- Additional fields: `Subject` (string), `Salary` (float64)

#### Functions (Not Methods)

##### Function 1: `CalculateAge`
- **Input**: A `Person` struct
- **Output**: An `int` (the age)
- **Requirement**: This is a **FUNCTION**, not a method (no receiver)
- **Behavior**: Return the age of the person

##### Function 2: `CreateStudentMap`
- **Input**: A slice of `Student` structs (`[]Student`)
- **Output**: A `map[string]Student` where key is StudentID
- **Requirement**: This is a **FUNCTION** that creates a map from a slice
- **Behavior**: Create a map with StudentID as key and Student as value

##### Function 3: `FindStudentByID`
- **Input**: A `map[string]Student` and a student ID (`string`)
- **Output**: A `Student` and a `bool` (found or not)
- **Requirement**: This is a **FUNCTION** that searches in a map
- **Behavior**: Use map lookup with ok idiom: `student, ok := studentMap[id]`

##### Function 4: `UpdateStudentGradeInMap`
- **Input**: A `map[string]Student`, student ID (`string`), and new grade (`float64`)
- **Output**: A `bool` indicating success
- **Requirement**: This is a **FUNCTION** that updates a student in a map
- **Behavior**: Find student, call `UpdateGrade` method, update map entry

#### Methods (With Receivers)

##### Method 1: `GetInfo` (Value Receiver)
- **Receiver**: `Student` (value receiver)
- **Input**: None (uses receiver)
- **Output**: A `string` with formatted student information
- **Requirement**: Use **value receiver** `(s Student)`
- **Behavior**: Return formatted string: `"Student: [Name], Age: [Age], ID: [StudentID]"`
- **Note**: Access embedded Person fields directly: `s.Name`, `s.Age`

##### Method 2: `UpdateGrade` (Pointer Receiver)
- **Receiver**: `*Student` (pointer receiver)
- **Input**: A new grade (`float64`)
- **Output**: None (modifies receiver)
- **Requirement**: Use **pointer receiver** `(s *Student)` to modify struct
- **Behavior**: Update the student's grade field

##### Method 3: `GetFullInfo` (Value Receiver)
- **Receiver**: `Student` (value receiver)
- **Input**: None (uses receiver)
- **Output**: A `map[string]interface{}` with all student information
- **Requirement**: Return a map with keys: `"name"`, `"age"`, `"studentID"`, `"grade"`
- **Behavior**: Create and populate a map with all student fields

##### Method 4: `GetSubjectInfo` (Value Receiver)
- **Receiver**: `Teacher` (value receiver)
- **Input**: None (uses receiver)
- **Output**: A `map[string]string` with teacher information
- **Requirement**: Return a map with keys: `"name"`, `"subject"`, `"salary"` (as string)
- **Behavior**: Create and populate a map with teacher's information

### Sample Input/Output

#### Example 1: Struct Embedding and Method with Value Receiver

**Input:**
```go
student := Student{
    Person: Person{Name: "Alice", Age: 20},
    StudentID: "S001",
    Grade: 85.5,
}
info := student.GetInfo()
```

**Output:**
```go
"Student: Alice, Age: 20, ID: S001"
```

**Note:** Notice how we access `Name` and `Age` directly from `Student` (embedded `Person` fields).

#### Example 2: Method with Pointer Receiver

**Input:**
```go
student := &Student{
    Person: Person{Name: "Bob", Age: 21},
    StudentID: "S002",
    Grade: 75.0,
}
student.UpdateGrade(90.0)
// student.Grade is now 90.0
```

**Output:**
```go
// Grade is modified to 90.0
```

#### Example 3: Method Returning Map

**Input:**
```go
student := Student{
    Person: Person{Name: "Charlie", Age: 19},
    StudentID: "S003",
    Grade: 92.5,
}
info := student.GetFullInfo()
```

**Output:**
```go
map[string]interface{}{
    "name": "Charlie",
    "age": 19,
    "studentID": "S003",
    "grade": 92.5,
}
```

#### Example 4: Function Creating Map from Slice

**Input:**
```go
students := []Student{
    {Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5},
    {Person: Person{Name: "Bob", Age: 21}, StudentID: "S002", Grade: 75.0},
}
studentMap := CreateStudentMap(students)
```

**Output:**
```go
map[string]Student{
    "S001": Student{Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5},
    "S002": Student{Person: Person{Name: "Bob", Age: 21}, StudentID: "S002", Grade: 75.0},
}
```

#### Example 5: Function Searching Map

**Input:**
```go
studentMap := map[string]Student{
    "S001": Student{Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5},
}
student, found := FindStudentByID(studentMap, "S001")
```

**Output:**
```go
student = Student{Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5}
found = true
```

#### Example 6: Function Updating Student in Map

**Input:**
```go
studentMap := map[string]Student{
    "S001": Student{Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5},
}
success := UpdateStudentGradeInMap(studentMap, "S001", 95.0)
// studentMap["S001"].Grade is now 95.0
```

**Output:**
```go
success = true
```

#### Example 7: Function vs Method

**Function (CalculateAge):**
```go
person := Person{Name: "Alice", Age: 20}
age := CalculateAge(person)  // Function call
```

**Method (GetInfo):**
```go
student := Student{Person: Person{Name: "Alice", Age: 20}, StudentID: "S001", Grade: 85.5}
info := student.GetInfo()  // Method call on instance
```

### Key Differences: Functions vs Methods

| Aspect | Function | Method |
|--------|----------|--------|
| Definition | `func FunctionName(params) returnType` | `func (receiver) MethodName(params) returnType` |
| Call | `FunctionName(args)` | `instance.MethodName(args)` |
| Access to data | Through parameters | Through receiver |
| Belongs to | Package | Type (struct) |

### Struct Embedding (Composition)

**Embedding Syntax:**
```go
type Student struct {
    Person      // Anonymous field - embeds Person
    StudentID string
    Grade     float64
}
```

**Accessing Embedded Fields:**
```go
student := Student{
    Person: Person{Name: "Alice", Age: 20},
    StudentID: "S001",
    Grade: 85.5,
}

// Direct access (promoted fields)
name := student.Name  // Not student.Person.Name
age := student.Age    // Not student.Person.Age
```

### Notes

- **Value Receiver**: Use when you don't need to modify the struct
- **Pointer Receiver**: Use when you need to modify the struct
- **Struct Embedding**: Anonymous fields are "promoted" - accessed directly
- **Map Lookup**: Always use `value, ok := map[key]` to check existence
- **Functions vs Methods**: Functions are standalone, methods belong to types
- **Composition over Inheritance**: Go uses embedding for code reuse (composition)

**Requirements:**

- Define `Student` and `Teacher` structs with embedded `Person`
- Implement `CalculateAge` as a **function** (not a method)
- Implement `GetInfo` with **value receiver**
- Implement `UpdateGrade` with **pointer receiver**
- Implement `GetFullInfo` and `GetSubjectInfo` methods returning maps
- Implement `CreateStudentMap` as a **function** creating map from slice
- Implement `FindStudentByID` as a **function** searching map
- Implement `UpdateStudentGradeInMap` as a **function** updating map entry
- Access embedded fields directly (promoted fields)

Try to implement all functions and methods below. Pay attention to the difference between functions and methods, and when to use value vs pointer receivers. Good Luck!

