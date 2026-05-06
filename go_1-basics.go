package main

import "fmt"

// Variable Declaration
var name string = "Someone"
var age int = 100
var height float64 = 6.7
var isStudent bool = false

// Documentation Function
// This function adds two integers and returns the result.
func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	// Print Something
	fmt.Println("Hello, World!")

	// Print Variable Values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	// Operators Arithmetic
	x := 10
	y := 5

	fmt.Println("\nArithmetic Operators:")
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)
	fmt.Println("Division:", x/y)
	fmt.Println("Remainder:", x%y)

	// Operators Comparison
	a := 10
	b := 20

	fmt.Println("\nComparison Operators:")
	fmt.Println("Equal to:", a == b)
	fmt.Println("Not equal to:", a != b)
	fmt.Println("Greater than:", a > b)
	fmt.Println("Less than:", a < b)
	fmt.Println("Greater than or equal to:", a >= b)
	fmt.Println("Less than or equal to:", a <= b)

	// Operators Logical
	isRaining := true
	hasUmbrella := false

	fmt.Println("\nLogical Operators:")
	fmt.Println("AND:", isRaining && hasUmbrella)
	fmt.Println("OR:", isRaining || hasUmbrella)
	fmt.Println("NOT:", !isRaining)

	// Operators Bitwise
	m := 5  // 0101 in binary
	n := 3  // 0011 in binary

	fmt.Println("\nBitwise Operators:")
	fmt.Println("AND:", m&n)   // 0001 in binary
	fmt.Println("OR:", m|n)    // 0111 in binary
	fmt.Println("XOR:", m^n)   // 0110 in binary
	fmt.Println("Left Shift:", m<<1) // 1010 in binary
	fmt.Println("Right Shift:", m>>1) // 0010 in binary

	// Input / Output
	var userName string

	fmt.Print("\nEnter your name: ")
	fmt.Scanln(&userName)

	fmt.Println("Hello,", userName)

	// Input / Output with Multiple Values
	var num1, num2 int

	fmt.Print("Enters Numbers: ")
	fmt.Scanln(&num1, &num2)
	
	fmt.Println("The sum of", num1, "and", num2, "is:", num1 + num2)

	// Function Call
	result := add(5, 10)
	fmt.Println("\nThe sum of 5 and 10 is:", result)
}