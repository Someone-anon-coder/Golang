package main

import "fmt"

// Function using pointer
func updateValue(num *int) {
	*num = 100
}

// Struct
type Person struct {
	name string
	age int
}

// Function using pointer to struct
func updatePerson(p *Person) {
	p.name = "Updated Someone"
	p.age = 9999
}

func main(){
	// Basic pointer
	fmt.Println("=== Basic Pointer ===")

	number := 10

	// Pointer stores memory address of variable
	var ptr *int = &number

	fmt.Printf("Value of number: %d\n", number)
	fmt.Printf("Memory address of number: %p\n", &number)
	fmt.Printf("Value stored in pointer (memory address): %p\n", ptr)

	// Dereferencing 
	fmt.Println("\n=== Dereferencing ===")
	fmt.Printf("Value pointed to by pointer: %d\n", *ptr)

	// Modifying value through pointer
	*ptr = 20
	fmt.Printf("Updated value of number: %d\n", number)

	// Passing pointer to function
	fmt.Println("\n=== Passing Pointer to Function ===")

	value := 50
	fmt.Printf("Value before function call: %d\n", value)
	updateValue(&value)
	fmt.Printf("Value after function call: %d\n", value)

	// Nil pointer
	fmt.Println("\n=== Nil Pointer ===")
	var nilPtr *int
	fmt.Printf("Nil pointer value: %v\n", nilPtr)

	// new() function
	fmt.Println("\n=== Using new() Function ===")

	data := new(int)
	fmt.Printf("Memory address of data: %p\n", data)
	fmt.Printf("Default Value of data: %d\n", *data)

	*data = 42
	fmt.Printf("Updated value of data: %d\n", *data)

	// Pointers with Structs
	fmt.Println("\n=== Pointers with Structs ===")

	person := Person{name: "John", age: 30}
	fmt.Printf("Person before update: %+v\n", person)

	updatePerson(&person)
	fmt.Printf("Person after update: %+v\n", person)

	// Struct pointer direct access
	fmt.Println("\n=== Struct Pointer Direct Access ===")

	personPtr := &person
	fmt.Printf("Person pointer: %+v\n", personPtr)

	personPtr.name = "Directly Updated Name"
	fmt.Printf("Person after direct update: %+v\n", *personPtr)

	// Pointer to pointer
	fmt.Println("\n=== Pointer to Pointer ===")

	x := 5

	ptr1 := &x
	ptr2 := &ptr1
	
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value stored in ptr1: %p\n", ptr1)
	fmt.Printf("Value stored in ptr2: %p\n", ptr2)
	fmt.Printf("Value pointed to by ptr1: %d\n", *ptr1)
	fmt.Printf("Value pointed to by ptr2: %d\n", **ptr2)
}