package main

import (
	"fmt"
)

type SampleStruct struct {
	val int
}

func main() {
	first := SampleStruct{}
	second := first
	second.val = 100
	fmt.Println(&first, &second)
}
