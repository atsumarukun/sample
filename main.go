package sample

import "fmt"

type Sample struct {
	value string
}

(_ Sample) func Print() {
	fmt.Println(value)
}
