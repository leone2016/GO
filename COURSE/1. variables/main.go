package main

import "fmt"

/*
Variables
* short variable declaration operator
* using the var keyword to declare a variable
* scope


 */
var x int
type person struct{
	Fname string
	Lname string
}
/*
Function
* func(receiver) identifier(params)(return){<CODE>}
* methods
*/

func(p person) speak(){
	fmt.Println(p.Fname, "says, Good morning, James")

}
/*
Composition
* embedded types
* interfaces
* polymorphism
*/

type screatAgent struct {
	person
	licenseToKill bool
}

func (sa screatAgent) speak()  {
	fmt.Println(sa.Fname, sa.Lname, "says, Shaken not stirred")
}

// polymorphism
// tanto como persona y secreatAngent implementa el metodo speak()
type human interface {
	speak()
}
func saySomething(h human){
	h.speak()
}
func main() {
	//x := 7
	fmt.Printf("%T",x)
	fmt.Println(x)
	/*
	Compositive literal
	type {}
	 */
	xi := []int{2,3,4,6,7,8,9}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job": 42,
	}
	fmt.Println(m)
	p1 := person{
		"Leo",
		"Medina",
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := screatAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sa1.speak()
	saySomething(p1)
	saySomething(sa1)
}
