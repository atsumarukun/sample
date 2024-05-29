package sample

import "fmt"

type Test struct {
	value string
}

func NewTest(value string) *Test {
	return &Test{value: value}
}

func (t Test) Print() {
	fmt.Println(t.value)
}
