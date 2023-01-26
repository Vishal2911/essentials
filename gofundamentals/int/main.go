package main

import "fmt"

func addOne(a int , b int) (int , int , string) {
	return a + 1 , b+1 , "coding concept"
}

func main() {
	var a int
	a = 5

	fmt.Println("value of a = ", a)
	a , b , _ := addOne(a , 6)
	
	fmt.Printf("\nafter addOne , value of a = %d , b = %d, name = \n",a , b )
	a , b , _ = addOne(a , b)
	fmt.Printf("\nafter addtwo , value of a = %d , b = %d, name = \n",a , b )
	
}