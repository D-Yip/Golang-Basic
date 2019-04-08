package main

import (
	"fmt"
)

func tryRecover() {

	defer func() {
		r := recover()
		if error, ok := r.(error); ok {
			fmt.Println("Error occurred:", error)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is an error"))
	panic(123)
}

func main() {
	tryRecover()
}
