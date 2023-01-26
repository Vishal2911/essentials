package main

import "fmt"

var a string

func init() {
	a = "coding concept"
	abc()
	fmt.Println("2 running init ...")
}

func init() {
	a = "coding concept"
	fmt.Println("1 running init ...")
}


func main() {

	xyz()
	fmt.Println("running main ..." , a )
}

func xyz() {
	fmt.Println("running xyz ...")
}

func abc() {
	fmt.Println("running abc ... coding ***********************")
}

