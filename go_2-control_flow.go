package main
import "fmt"

func main(){
	// Even or Odd
	number := 45

	if number % 2 == 0 {
		fmt.Printf("%d is an even number\n", number)
	} else {
		fmt.Printf("%d is an odd number\n", number)
	}

	// Positive, Negative or Zero
	number2 := -10

	if number2 > 0 {
		fmt.Printf("%d is a positive number\n", number2)
	} else if number2 < 0 {
		fmt.Printf("%d is a negative number\n", number2)
	} else {
		fmt.Printf("%d is zero\n", number2)
	}

	// Nested if-else
	age := 25
	gender := "male"

	if age >= 18 {
		if gender == "male" {
			fmt.Println("You are an adult male.")
		} else if gender == "female" {
			fmt.Println("You are an adult female.")
		} else {
			fmt.Println("You are an adult.")
		}
	} else {
		fmt.Println("You are not an adult.")
	}

	// Short statement in if
	if score := 85; score >= 50 {
		fmt.Println("You passed the exam")
	} else {
		fmt.Println("You failed the exam")
	}

	// Logical Condition
	isStudent := true
	isEmployed := false

	if isStudent && !isEmployed {
		fmt.Println("You are a student and not employed.")
	} else if isStudent && isEmployed {
		fmt.Println("You are a student and employed.")
	} else if !isStudent && isEmployed {
		fmt.Println("You are not a student but employed.")
	} else {
		fmt.Println("You are neither a student nor employed.")
	}

	// Switch statement
	day := "Saturday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday.")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("Invalid day.")
	}

	// Switch without expression
	grade := 85

	switch {
	case grade >= 90:
		fmt.Println("Excellent")
	case grade >= 80:
		fmt.Println("Good")
	case grade >= 70:
		fmt.Println("Fair")
	default:
		fmt.Println("Needs Improvement")
	}
}