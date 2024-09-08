package main

import (
	"fmt"
	
)

type person struct{
	firstName string
	lastName string
	contactInfo
} 

type contactInfo struct{
	email string
	zipCode int

}
func main(){
	var alex person// when var is declared without assigning values, go assigns a zero value according to the type of it
	alex.firstName = "alex" 
	alex.lastName = "anderson"
	fmt.Println(&alex)
	

	jim :=person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email:"jim@gmail.com",
			zipCode: 94000,
		},
	}
	

	jim.updateName("jiggy")
	jim.print()
	alex.print()


	
}
	// fmt.Println(jim)
	//fmt.Println(alex)
	// fmt.Printf("%+v",jim)
	
func (p person) print(){
	fmt.Printf("%+v",p)
}

func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}
