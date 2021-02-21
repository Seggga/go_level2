package main

import (
	"fmt"
	"time"
)

//ErrorWithTimeStamp includes a timestamp and a message fields
type ErrorWithTimeStamp struct {
	timeStamp time.Time
	message   string
}

func (m *ErrorWithTimeStamp) Error() string {
	return fmt.Sprintf("%s : %s", m.timeStamp.String(), m.message)
}

func newError(message string) error {
	return &ErrorWithTimeStamp{
		timeStamp: time.Now(),
		message:   message,
	}
}

func main() {
	slice := []int32{1, 2, 3, 4, 5}

	defer func() {
		if v := recover(); v != nil {
			message := fmt.Sprintf("index out of range [%d] with length %d", len(slice), len(slice))
			err := newError(message)
			fmt.Printf("%v", err)
		}
	}()

	fmt.Println(slice[len(slice)]) //попытка обраиться к элементу за пределами capacity слайса

}
