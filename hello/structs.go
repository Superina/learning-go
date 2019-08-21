package main

import "fmt"



// struct is like a python class
type employee struct {
	first string
	last string
	salary int
	bonus int
	manager string
}


/* This function is a VALUE RECEIVER in that in only receives the value passed through by the employee instance. It does NOT modify the value of the instance */

func (e employee) total_comp() int {
	return e.salary + e.bonus
}

/* This function is a POINTER RECEIVER in that in changes the value of the struct instance it receives.
   Fun fact: try taking out the *, it will not report an error, but when you print out emp1.salary, it won't change either
*/
func (e *employee) raise(raise_amount int) {
	e.salary = e.salary + raise_amount
}


func main() {

	emp1 := employee{
		first: "Alice",
		last: "InWonderland",
		salary: 1000000,
		bonus: 3000000,
		manager: "Ed Stark",
	}
	
	// way to access struct value
	fmt.Println(emp1.manager)

	// way to pass an instance of the struct to a function
	fmt.Println(emp1.total_comp())

	// way to modify the value of a struct instance using a function
	emp1.raise(888888)
	fmt.Println(emp1.salary)

}