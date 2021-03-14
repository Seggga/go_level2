package main

import (
	"fmt"

	"github.com/Seggga/go_level2/01_reflection/change"
)

type SomeStruct struct {
	Field1 string
	Field2 int64
	Field3 bool
}

func main() {

	st := SomeStruct{
		Field1: "one",
		Field2: 10,
		Field3: true,
	}

	values := map[string]interface{}{
		"Field1": "new",
		"Field2": int64(20),
		"Field3": false,
	}

	fmt.Println(st)
	err := change.SetValues(&st, values)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	fmt.Println(st)
}
