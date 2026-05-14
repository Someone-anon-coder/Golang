package main

import "fmt"

// Struct Declaration
type Person struct {
	Name string
	Age  int
}

// Method with value receiver
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Method with pointer receiver
func (p *Person) UpdateInfo(newName string, newAge int) {
	p.Name = newName
	p.Age = newAge
}

// Another struct
type Address struct {
	city string
	country string
}

// Embedded Struct
type Employee struct {
	Person
	Address
	Position string
}

// Method for Employee struct
func (e Employee) DisplayInfo() {
	fmt.Println("Employee Information:")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("City: %s\n", e.city)
	fmt.Printf("Country: %s\n", e.country)
	fmt.Printf("Position: %s\n", e.Position)
}

func main() {
	// Creating Struct
	fmt.Println("=== Creating Struct ===")

	person1 := Person{Name: "Alice", Age: 30}
	fmt.Println(person1)

	// Accessing Struct Fields
	fmt.Println("\n=== Accessing Struct Fields ===")
	
	fmt.Printf("Name: %s, Age: %d\n", person1.Name, person1.Age)

	// Modifying Struct Fields
	fmt.Println("\n=== Modifying Struct Fields ===")
	
	person1.Age = 31
	fmt.Printf("Updated Age: %d\n", person1.Age)

	// Calling Methods
	fmt.Println("\n=== Calling Methods ===")

	person1.Greet()

	// Pointer Receiver Method
	fmt.Println("\n=== Pointer Receiver Method ===")
	
	person1.UpdateInfo("Alice Smith", 32)
	person1.Greet()

	// Anonymous Struct
	fmt.Println("\n=== Anonymous Struct ===")

	anonymous := struct {
		Title string
		Author string
	}{
		Title: "Go Programming",
		Author: "Someone",
	}

	fmt.Printf("Title: %s, Author: %s\n", anonymous.Title, anonymous.Author)

	// Embedded Struct
	fmt.Println("\n=== Embedded Struct ===")

	employee1 := Employee{
		Person: Person{Name: "Bob", Age: 25},
		Address: Address{city: "New York", country: "USA"},
		Position: "Software Engineer",
	}

	employee1.DisplayInfo()

	// Direct access to embedded struct fields
	fmt.Println("\n=== Direct Access to Embedded Struct Fields ===")

	fmt.Printf("Employee Name: %s\n", employee1.Name)
	fmt.Printf("Employee City: %s\n", employee1.city)

	// Struct Comparison
	fmt.Println("\n=== Struct Comparison ===")

	person2 := Person{Name: "Alice Smith", Age: 32}
	fmt.Printf("person1 == person2: %v\n", person1 == person2)

	// Struct pointer
	fmt.Println("\n=== Struct Pointer ===")

	ptr := &person1
	fmt.Printf("Pointer to person1: %v\n", ptr)
	fmt.Printf("Dereferenced pointer: %v\n", *ptr)

	ptr.Age = 33
	fmt.Printf("Updated Age through pointer: %d\n", person1.Age)
}