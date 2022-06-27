package main

import (
	"fmt"
)

func main() {
	var x float32
	var y int32
	fmt.Println("Enter a float")
	fmt.Scanln(&x)
	y = int32(x)
	fmt.Println(y)
}
