package main

import "fmt"

func main() {
	a := 15
	b := &a // memory address of a
	fmt.Println(b) // 0xc000016080

	c := *b // the value that's being stored at the memory address b represents (0xc000016080)
	fmt.Println(c)

	// notice here is not *b := 8000 because all items at the left side of := must be pure identifiers and at least one of them must be a new variable name
	// This means container elements (x[i]), struct fields (x.f), pointer dereferences (*p) and qualified identifiers (aPackage.Value) can't appear at the left side of :=
	*b = 8000 // change the value that address b stores to 8000  -->  (b is the address of a) --> a:=8000
	fmt.Println("now the value of a becomes",a)
}


/*summuary,upon seeing &, e.g. &m, m is always an actual value
   		   upon seeing *, e.g. *n, n is always an address
*/