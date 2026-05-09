package main

import "fmt"

// Simple function
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func greetPerson(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

// Function with return value
func add(num1 int, num2 int) int {
	return num1 + num2
}

// Multiple return values
func divide(a float64, b float64) (float64, string) {
	if b == 0 {
		return 0, "Division by zero is not allowed"
	}
	return a / b, "Division successful"
}

// Variadic function
func total(numbers ...int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	// Calling simple function
	greet()

	// Calling function with parameters
	greetPerson("Alice", 30)

	// Calling function with return value
	num1 := 5
	num2 := 3

	result := add(num1, num2)
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, result)

	// Multiple return values
	quotient, message := divide(10, 2)
	fmt.Printf("Quotient: %f, Message: %s\n", quotient, message)

	quotient, message = divide(10, 0)
	fmt.Printf("Quotient: %f, Message: %s\n", quotient, message)

	// Calling variadic function
	sum := total(1, 2, 3, 4, 5)
	fmt.Printf("The total sum is: %d\n", sum)

	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function!")
	}()

	// Assigning function to a variable
	square := func(x int) int {
		return x * x
	}

	fmt.Printf("The square of 4 is: %d\n", square(4))

	// Closure example
	fmt.Println("Closure Example")

	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}
	increament := counter()

	fmt.Println(increament())
	fmt.Println(increament())
	fmt.Println(increament())

	// Calling recursive function
	fmt.Println("Recursive Function")
	fmt.Printf("Factorial of 5 is: %d\n", factorial(5))
}