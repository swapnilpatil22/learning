package main

// Person represents a basic person with name and age
// This struct will be embedded in other structs (composition)
type Person struct {
	Name string
	Age  int
}

// Student represents a student with person information and student ID
// Requirement: Embed Person struct (composition for multiple inheritance-like behavior)
type Student struct {
	// TODO: Embed Person struct here
	// Hint: Use anonymous field embedding: Person
	StudentID string
	Grade     float64
}

// Teacher represents a teacher with person information and subject
// Requirement: Embed Person struct (composition)
type Teacher struct {
	// TODO: Embed Person struct here
	Subject string
	Salary  float64
}

// CalculateAge is a FUNCTION (not a method) that calculates age
// Input: A Person struct
// Output: The age of the person
// Requirement: This should be a standalone function, not a method
func CalculateAge(p Person) int {
	// TODO: Implement as a function (not a method)
	// Return the age of the person
	return 0
}

// GetInfo is a METHOD with value receiver
// Input: None (uses receiver)
// Output: A string containing student information
// Requirement: Use value receiver (not pointer receiver)
func (s Student) GetInfo() string {
	// TODO: Implement method with value receiver
	// Return a formatted string with student's name, age, and student ID
	// Format: "Student: [Name], Age: [Age], ID: [StudentID]"
	return ""
}

// UpdateGrade is a METHOD with pointer receiver
// Input: A new grade (float64)
// Output: None (modifies the receiver)
// Requirement: Use pointer receiver to modify the struct
func (s *Student) UpdateGrade(newGrade float64) {
	// TODO: Implement method with pointer receiver
	// Update the student's grade
	// Hint: Use pointer receiver (*Student) to modify the struct
}

// GetFullInfo is a METHOD that returns detailed information
// Input: None (uses receiver)
// Output: A map[string]interface{} containing all student information
// Requirement: Return a map with keys: "name", "age", "studentID", "grade"
func (s Student) GetFullInfo() map[string]interface{} {
	// TODO: Implement method that returns a map
	// Create and return a map containing all student fields
	result := make(map[string]interface{})
	return result
}

// GetSubjectInfo is a METHOD for Teacher
// Input: None (uses receiver)
// Output: A map[string]string with teacher information
// Requirement: Return a map with keys: "name", "subject", "salary" (as string)
func (t Teacher) GetSubjectInfo() map[string]string {
	// TODO: Implement method that returns a map
	// Create and return a map containing teacher's name, subject, and salary
	result := make(map[string]string)
	return result
}

// CreateStudentMap creates a map of students
// Input: A slice of Student structs
// Output: A map[string]Student where key is StudentID and value is Student
// Requirement: This is a FUNCTION (not a method) that works with maps
func CreateStudentMap(students []Student) map[string]Student {
	// TODO: Implement function that creates a map from slice
	// Iterate through students slice and create a map
	// Key: StudentID, Value: Student struct
	result := make(map[string]Student)
	return result
}

// FindStudentByID finds a student in a map by ID
// Input: A map[string]Student and a student ID (string)
// Output: The Student struct and a boolean indicating if found
// Requirement: This is a FUNCTION that searches in a map
func FindStudentByID(studentMap map[string]Student, id string) (Student, bool) {
	// TODO: Implement function that searches map
	// Use map lookup with ok idiom: value, ok := map[key]
	var result Student
	found := false
	return result, found
}

// UpdateStudentGradeInMap updates a student's grade in a map
// Input: A map[string]Student, student ID, and new grade
// Output: A boolean indicating success
// Requirement: Update the student's grade in the map using pointer receiver method
func UpdateStudentGradeInMap(studentMap map[string]Student, id string, newGrade float64) bool {
	// TODO: Implement function that updates student in map
	// Find the student, update their grade using UpdateGrade method
	// Return true if student found and updated, false otherwise
	// Hint: Get student from map, call UpdateGrade method, put back in map
	return false
}
