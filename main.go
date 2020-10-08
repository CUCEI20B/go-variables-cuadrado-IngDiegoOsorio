package main

import "fmt"

func main() {

	var lado float32

	fmt.Printf("Dime cuanto mide un lado del cuadrado: ")

	fmt.Scanf("%f", &lado)
	total := lado * lado
	fmt.Println("El area del cuadrado es: ", total)
}