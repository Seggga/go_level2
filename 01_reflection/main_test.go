package main

import (
	"testing"

	"github.com/Seggga/go_level2/01_reflection/change"
)

func TestSetValues(t *testing.T) {
	// Initialize data
	goodStruct := SomeStruct{Field1: "old", Field2: 10, Field3: true}
	badStruct := struct{ Field1, Field2 int64 }{Field1: 1, Field2: 10}
	emptyStruct := struct{}{}

	values := map[string]interface{}{
		"Field1": "new",
		"Field2": int64(20),
		"Field3": false,
	}

	// different struct for testing
	errGood := change.SetValues(&goodStruct, values)
	errBad := change.SetValues(&badStruct, values)
	errEmpty := change.SetValues(&emptyStruct, values)

	// GOOD
	if errGood != nil {
		t.Errorf("Error in good struct: %v", errGood)
	}
	if goodStruct.Field1 != "new" || goodStruct.Field2 != 20 || goodStruct.Field3 {
		t.Errorf("Error in good struct: expected %v, got %v", values, goodStruct)
	}
	// BAD
	if errBad == nil {
		t.Errorf("Error in bad struct: expected an error, got nil.\nBadsSruct %v\n mab %v", badStruct, values)
	}
	// EMPTY
	if errEmpty != nil {
		t.Error("Error in empty struct: expected an error, got nil")
	}
}
