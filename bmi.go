package main

import "fmt"

func main() {
	var height, weight, bmi float32
	fmt.Scanf("%f", &height)
	fmt.Scanf("%f", &weight)
	bmi = weight / (height * height)
	fmt.Println(bmi)
}
