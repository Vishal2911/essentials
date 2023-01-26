package main

import "fmt"


func main() {

	var a bool 
	if a {
		fmt.Println("value of a is true")
	}else {
		fmt.Println("value of a is false")
	}

	name := "coding-concept"
	if name == "Coding concept" {
		fmt.Println("string matched , value = ", name)
	}else if name == "coding concept" {
		fmt.Println("string  matched , value = ", name)
	}else {
		fmt.Println("string not matched , value = ", name)
	}


	number1 := 5
	number2 := 10
	if number1 > number2 {
		fmt.Printf("%d grater %d \n",number1, number2)
	}else if number1 <= number2 {
		fmt.Printf("%d less or equal %d \n",number1, number2)
	}else {
		fmt.Printf("%d equal %d \n",number1, number2)
	}


	if number1 > number2 || number1 < number2 {
		fmt.Printf("%d not equal %d \n",number1, number2)
	}

	if number1 < number2 && number1 > 3 {
		fmt.Printf("%d less than  %d  and %d grater than 3\n",number1, number2 , number1)
	}





}