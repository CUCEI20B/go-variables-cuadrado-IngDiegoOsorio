package main

import "fmt"

func main() {
	var lado float32
	fmt.Scanf("%f", &lado)
	area := lado * lado
	fmt.Println(area)
}