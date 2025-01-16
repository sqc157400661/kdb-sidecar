package output

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
)

// TableOutput is a generic function to print a table from any struct slice based on struct tags
func TableOutput[T any](data []T) {
	// Check if the data slice is empty
	if len(data) == 0 {
		fmt.Println("No data to display.")
		return
	}

	// Create a new table writer
	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)

	// Use reflection to get struct field names from the 'tab' tag
	structType := reflect.TypeOf(data[0])
	var headers []interface{}
	for i := 0; i < structType.NumField(); i++ {
		// Get the tag value for 'tab' (header name)
		tagValue := structType.Field(i).Tag.Get("tab")
		// If the tag is empty, fallback to the field name
		if tagValue == "" {
			tagValue = structType.Field(i).Name
		}
		headers = append(headers, tagValue)
	}

	// Set the table headers based on the struct tags
	t.AppendHeader(headers)

	// Add data rows to the table
	for _, item := range data {
		var row []interface{}
		// Use reflection to get the field values from each struct instance
		for i := 0; i < structType.NumField(); i++ {
			fieldValue := reflect.ValueOf(item).Field(i).Interface()
			row = append(row, fieldValue)
		}
		t.AppendRow(row)
	}

	// Render and display the table
	t.Render()

	// Additional message to simulate the MySQL-like output
	fmt.Printf("Query OK, %d rows affected.\n", len(data))
}