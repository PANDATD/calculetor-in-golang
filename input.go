package main

import "fmt"

var MyMessage string ="\n"
var value1 float64 
var value2 float64


func main(){
	fmt.Println(MyMessage)

for{
	ShowChoice()
	UserChoice := GetUserChoice()
	if UserChoice == 1{
		
		value1, value2 := GetIntInput()
		result := Add(value1,value2)
		fmt.Printf("Addition of %v + %v = %v\n\n", value1, value2,result)
		
	}else if UserChoice == 2{
		
		value1, value2 := GetIntInput()
		result := Sub(value1,value2)
		fmt.Printf("Substraction of %v - %v = %v\n\n", value1, value2,result)

	}else if UserChoice == 3{
		
		value1,value2 := GetIntInput()
		result := Mul(value1,value2)
		fmt.Printf("Multiplication of %v * %v = %v\n\n", value1, value2,result)

	}else if UserChoice == 4{
		
		value1,value2 := GetIntInput()
		result := Div(value1,value2)
		fmt.Printf("Division of %v / %v = %v", value1, value2,result)

	}else{
		break
	}
}
	
}


/////////////////////////////////////////////////////////////////////////



func ShowChoice(){
	var ShowChoice string="\n1.Addition\n2.Substraction\n3.Mulatiplication\n4.Division\n5.any integer key than 1 to 4 \n"
	fmt.Println(ShowChoice)
}

func GetUserChoice()int{
	var UserChoice int
	fmt.Println("Select Opretion no which ever you want to use")
	fmt.Scanln(&UserChoice)
	return UserChoice
}

func GetIntInput()(float64, float64){
	fmt.Println("Enter the First value:	")
	fmt.Scanln(&value1)
	fmt.Println("Enter the Second value: ")
	fmt.Scanln(&value2)
	return value1 ,value2
}

////////////////////////////////////////////////////////////



func Add(value1 float64, value2 float64)float64{
	return value1+value2
}

func Sub(value1 float64, value2 float64)float64{
	return value1-value2
}

func Mul(value1 float64, value2 float64)float64{
	return value1*value2
}

func Div(value1 float64, value2 float64)float64{
	return value1/value2
}
