package main

import "fmt"

var MyMessage string = "Welcome to Golang Based calculetor\n"
var value1 float64
var value2 float64

func main() {
	fmt.Println(MyMessage)

	for {
		ShowChoice()
		UserChoice := GetUserChoice()
		if UserChoice == 1 {

			value1, value2 := GetIntInput()
			result := Add(value1, value2)
			fmt.Printf("Addition of %v + %v = %v\n\n", value1, value2, result)

		} else if UserChoice == 2 {

			value1, value2 := GetIntInput()
			result := Sub(value1, value2)
			fmt.Printf("Substraction of %v - %v = %v\n\n", value1, value2, result)

		} else if UserChoice == 3 {

			value1, value2 := GetIntInput()
			result := Mul(value1, value2)
			fmt.Printf("Multiplication of %v * %v = %v\n\n", value1, value2, result)

		} else if UserChoice == 4 {

			value1, value2 := GetIntInput()
			result := Div(value1, value2)
			fmt.Printf("Division of %v / %v = %v", value1, value2, result)

		} else {
			break
		}
	}

}

/////////////////////////////////////////////////////////////////////////
