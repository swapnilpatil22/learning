package main

// Helper function: getSeason returns the season name for a given month number
// You can reuse logic from Demo 1 here
func getSeason(month int) string {
	// TODO: Implement helper function to get season name from month number
	// Hint: You can use switch-case or if-else from Demo 1
	return ""
}

// CountSeasons counts how many months belong to each season
// Input: A slice of month numbers
// Output: A map with season names as keys and counts as values
// Requirement: Use for-range loop to iterate through the input slice
func CountSeasons(months []int) map[string]int {
	// TODO: Implement using for-range loop
	// Iterate through months slice and count occurrences of each season
	result := make(map[string]int)
	return result
}

// GetSeasonMonths returns all month numbers that belong to a specific season
// Input: A season name (string)
// Output: A slice of month numbers belonging to that season
// Requirement: Use for-range loop to iterate through a map
func GetSeasonMonths(season string) []int {
	// TODO: Implement using for-range loop with a map
	// Create a map that maps season names to their month numbers
	// Then use for-range to iterate through the map and find matching months
	result := []int{}
	return result
}

// ProcessMonthlyData groups months by their respective seasons
// Input: A slice of month numbers
// Output: A map with season names as keys and slices of month numbers as values
// Requirement: Use traditional for loop (with index) to iterate through the input
func ProcessMonthlyData(months []int) map[string][]int {
	// TODO: Implement using traditional for loop (for i := 0; i < len(months); i++)
	// Group all months by their seasons
	result := make(map[string][]int)
	return result
}

