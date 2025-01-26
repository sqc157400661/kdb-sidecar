package output

import (
	"testing"
)

// User struct with custom 'tab' tags for header mapping
type User struct {
	ID   int    `tab:"ID"`
	Name string `tab:"Name"`
	Age  int    `tab:"Age"`
}

func TestPrintTable(t *testing.T) {
	// Test case 1: Valid data
	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}
	TableOutput(users)
	TableOutput([]User{})
}

func TestFormatOutToStdout(t *testing.T) {
	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}
	FormatOutToStdout(users, "json")
}
