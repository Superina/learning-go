package main

import "fmt"

func add(x,y float64) float64{
	return x + y
}

// explicitly decalre var num1, num2 as type float 64
func main() {
	var num1, num2 float64 = 1.9,2.1
	fmt.Println(add(num1,num2))
}


// implicitly have go decide which datatype to assign to num1,num2 when it compiles
func main() {
	num1, num2 := 1.9,2.1
	fmt.Println(add(num1,num2))
}


