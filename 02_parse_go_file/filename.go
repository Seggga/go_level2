package main

import (
	"fmt"
)

func funcWith5Go() {
	var i int = 10
	go func() { i++ }()
	go func() { i++ }()
	i++
	fmt.Println(i)
	go func() { i++ }()
	go func() { i++ }()
	go func() { i++ }()
	return
}

func funcWithoutGo() {
	var i int = 10
	i++
	fmt.Println(i)
}
