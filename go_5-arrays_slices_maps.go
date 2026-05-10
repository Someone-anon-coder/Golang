package main

import "fmt"

func main() {
	// Arrays
	fmt.Println("=== Arrays ===")

	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	
	fmt.Println(numbers)
	fmt.Printf("First element: %d\n", numbers[0])
	fmt.Printf("Length of array: %d\n", len(numbers))

	// Short declaration of array
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// Loop through array
	for i := 0; i < len(colors); i++ {
		fmt.Printf("Color at index %d: %s\n", i, colors[i])
	}

	// Slices
	fmt.Println("\n=== Slices ===")

	scores := []int{90, 80, 70}
	fmt.Printf("Scores: %v\n", scores)

	// Append to slice
	scores = append(scores, 60)
	scores = append(scores, 50, 40)
	fmt.Printf("Scores after appending: %v\n", scores)

	// Slice Operations
	subScores := scores[1:4] // Slicing from index 1 to 3
	fmt.Printf("Sub-slice of scores: %v\n", subScores)

	fmt.Printf("First two scores: %v\n", scores[:2])
	fmt.Printf("Last two scores: %v\n", scores[len(scores)-2:])

	// make function to create a slice with specific length and capacity
	data := make([]int, 5, 10)
	fmt.Printf("Data slice: %v, Length: %d, Capacity: %d\n", data, len(data), cap(data))

	// Copy slices
	fmt.Println("\n=== Copying Slices ===")

	original := []string{"Go", "Python", "Java"}
	copied := make([]string, len(original))

	copy(copied, original)
	fmt.Printf("Original slice: %v\n", original)
	fmt.Printf("Copied slice: %v\n", copied)

	// Modifying the copied slice does not affect the original
	copied[0] = "C++"
	original[1] = "JavaScript"

	fmt.Printf("After modification:\n")
	fmt.Printf("Original slice: %v\n", original)
	fmt.Printf("Copied slice: %v\n", copied)

	// Multi dimensional slices
	fmt.Println("\n=== Multi-dimensional Slices ===")

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Printf("Matrix: %v\n", matrix)
	fmt.Printf("Element at (1,2): %d\n", matrix[1][2])

	// Maps
	fmt.Println("\n=== Maps ===")

	people := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Eve":   28,
	}

	fmt.Printf("People: %v\n", people)

	// Accessing map values
	fmt.Printf("Alice's age: %d\n", people["Alice"])

	// Adding new key-value pair
	people["Charlie"] = 35
	fmt.Printf("People after adding Charlie: %v\n", people)

	// Deleting a key-value pair
	delete(people, "Bob")
	fmt.Printf("People after deleting Bob: %v\n", people)

	// Updating a value
	people["Eve"] = 29
	fmt.Printf("People after updating Eve's age: %v\n", people)

	// Checking if a key exists
	if age, exists := people["Alice"]; exists {
		fmt.Printf("Alice's age: %d\n", age)
	} else {
		fmt.Println("Alice not found in the map")
	}

	// Loop through map
	fmt.Println("\nLooping through people map:")
	for name, age := range people {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Nested maps
	fmt.Println("\n=== Nested Maps ===")

	countries := map[string]map[string]string{
		"USA": {
			"Capital": "Washington D.C.",
			"Currency": "USD",
		},
		"Japan": {
			"Capital": "Tokyo",
			"Currency": "JPY",
		},
	}

	fmt.Printf("Countries: %v\n", countries)
	fmt.Printf("Capital of USA: %s\n", countries["USA"]["Capital"])
	fmt.Printf("Currency of Japan: %s\n", countries["Japan"]["Currency"])


	// Simulating a set using a map
	fmt.Println("\n=== Simulating a Set using a Map ===")

	set := map[string]bool{
		"Go":     true,
		"Python": false,
		"Java":   true,
	}

	fmt.Printf("Set: %v\n", set)

	// Adding an element to the set
	set["C++"] = true
	fmt.Printf("Set after adding C++: %v\n", set)

	// Removing an element from the set
	delete(set, "Python")
	fmt.Printf("Set after removing Python: %v\n", set)

	// Checking if an element exists in the set
	if _, exists := set["Go"]; exists {
		fmt.Println("Go is in the set")
	} else {
		fmt.Println("Go is not in the set")
	}
}