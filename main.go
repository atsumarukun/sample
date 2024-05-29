package sample

import "fmt"

type Sample struct {
	value string
}

func NewSample(value string) *Sample {
	return &Sample{value: value}
}

(_ Sample) func Print() {
	fmt.Println(value)
}
