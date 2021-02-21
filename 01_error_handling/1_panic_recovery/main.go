package main

import "fmt"

func main() {
	slice := []int32{1, 2, 3, 4, 5}

	defer func() {
		if v := recover(); v != nil {
			err := fmt.Errorf("index out of range [%d] with length %d", len(slice), len(slice))
			fmt.Printf("%v", err)
		}
	}()

	fmt.Println(slice[len(slice)]) //попытка обраиться к элементу за пределами capacity слайса

}
