package main

import (
	"fmt"
	"strconv"
)

func main() {

	// traditional way
	var a int32
	var c = 5

	// short-hand
	b := int(a) + c
	d := a + int32(c)

	//string to int type casting
	convertedvalue , _ := strconv.Atoi("5")
	e := d + int32(convertedvalue)

	fmt.Println("a = ", a, ",b = ", b, ", c = ", c, ", d = ", d , ", e = ",e)

	// traditional way
	var p string
	var q = "Coding Concept"

	// short-hand
	r := "hello team"

	s := r + fmt.Sprintf(" %v", b)

	fmt.Println("p = ", p, ",q = ", q, ", r = ", r, ", s = ", s)

	// traditional way
	var x bool
	var y = true

	// short-hand
	z := true

	fmt.Println("x = ", x, ",y = ", y, ", z = ", z)

}
