package main

import "fmt"

func main(){
	// Basics for loop
	fmt.Println("Basics for loop")

	for i := 0; i < 5; i++ {
		fmt.Println("Value: %d", i)
	}

	// While-style loop
	fmt.Println("While-style loop")

	number := 1
	for number <= 5 {
		fmt.Println("Value: %d", number)
		number++
	}

	// Infinite loop
	fmt.Println("Infinite loop")

	/*
	for {
		fmt.Println("This loop will run forever")
	}
	*/

	// break statement
	fmt.Println("Break statement")

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Value: %d", i)
	}

	// continue statement
	fmt.Println("Continue statement")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Value: %d", i)
	}

	// Nested loop
	fmt.Println("Nested loop")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

	// range with slices
	fmt.Println("Range with slices")

	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// range with strings
	fmt.Println("Range with strings")

	str := "Hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}

	// Looping through maps
	fmt.Println("Looping through maps")

	person := map[string]string{
		"name": "Alice",
		"age":  "30",
		"city": "New York",
	}

	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Pattern Printing
	fmt.Println("Pattern Printing")
	
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	fmt.Println()
}
