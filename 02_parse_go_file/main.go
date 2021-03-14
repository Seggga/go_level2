package main

import (
	"fmt"

	"github.com/Seggga/go_level2/02_parse_go_file/finder"
)

func main() {
	fileName := "filename.go"
	funcName := "funcWith5Go"

	count, err := finder.FindInGo(fileName, funcName)
	if err != nil {
		fmt.Println("there was a problem during function call", err)
		return
	}
	fmt.Println(count)

}
