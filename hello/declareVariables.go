package main

import "fmt"

func add(x,y float64) float64{
	return x + y
}

// notice I have to type every single return value's data type in (string,string).
// there is no shorthand about it
func multiple(a,b string) (string,string){
	return a,b
}

// implicitly have go decide which datatype to assign to num1,num2 when it compiles
func main() {
	num1, num2 := 1.9,2.1
	fmt.Println(add(num1,num2))

	word1,word2 := "Hello", "World"
	fmt.Println(multiple(word1,word2))
}