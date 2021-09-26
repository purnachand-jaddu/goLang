package main

import "fmt"

type Pair struct {
	first  int
	second int
}

func main() {
	pair := Pair{}
	pair.first = 1
	pair.second = 2
	duplicate := pair
	fmt.Println(&pair == &duplicate)

	//Change something
	duplicate.second = 3
	fmt.Println(pair == duplicate)

}
