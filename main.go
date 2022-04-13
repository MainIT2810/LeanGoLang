package main

import (
	"fmt"
	// "bufio" //thư viện chuẩn của Google
	// "os"
	
	)

func main() {
	// fmt.Println("Hello, world.")

	/*
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city:")
	city, _ := reader.ReadString('\n')
	fmt.Println("You live in", city)
	*/

	height := 1.68
	weight := 80.0
	bmi := CalculateBMI(height, weight)
	fmt.Println("Your BMI is", bmi)
}

func CalculateBMI(height, weight float64) float64 {
	return weight / (height * height)
}