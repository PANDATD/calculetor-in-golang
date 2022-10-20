package main


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
