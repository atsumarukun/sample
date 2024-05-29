package sample

import "fmt"

type Sample struct {
	value string
}

func NewSample(value string) *Sample {
	return &Sample{value: value}
}

func (s Sample) Print() {
	fmt.Println(s.value)
}
