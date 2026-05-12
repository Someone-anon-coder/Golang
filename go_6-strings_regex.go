package main

import (
	"fmt" 
	"regexp"
	"strings"
)

func main() {
	// String Declaration
	fmt.Println("=== String Declaration ===")

	name := "Golang"
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Length: %d\n", len(name))

	// String Concatenation
	fmt.Println("\n=== String Concatenation ===")

	firstName := "Someone"
	lastName := "Coder"

	fullName := firstName + " " + lastName
	fmt.Printf("Full Name: %s\n", fullName)

	// String formatting
	fmt.Println("\n=== String Formatting ===")

	age := 30
	fmt.Printf("%s is %d years old\n", fullName, age)

	// Escaping characters
	fmt.Println("\n=== Escaping Characters ===")

	fmt.Println("Hello\nWorld")
	fmt.Println("Go\tLanguage")
	fmt.Println("She said, \"Hello!\"")

	// Strings Package
	fmt.Println("\n=== Strings Package ===")

	text := "  hello golang world  "
	
	fmt.Printf("Original Text: %s\n", text)
	fmt.Printf("Trimmed Text: '%s'\n", strings.TrimSpace(text))
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(text))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(text))
	fmt.Printf("Contains 'golang': %t\n", strings.Contains(text, "golang"))
	fmt.Printf("Replace 'golang' with 'Go': %s\n", strings.ReplaceAll(text, "golang", "Go"))
	fmt.Printf("Count of 'o': %d\n", strings.Count(text, "o"))
	fmt.Printf("Index of 'golang': %d\n", strings.Index(text, "golang"))
	fmt.Printf("Split by space: %v\n", strings.Split(text, " "))

	// Split and Join
	fmt.Println("\n=== Split and Join ===")

	data := "apple,banana,orange"
	fruits := strings.Split(data, ",")
	fmt.Printf("Fruits: %v\n", fruits)
	fmt.Printf("Joined: %s\n", strings.Join(fruits, " - "))

	// String Slicing
	fmt.Println("\n=== String Slicing ===")

	word := "Golang"
	fmt.Printf("First 3 characters: %s\n", word[:3])
	fmt.Printf("Last 3 characters: %s\n", word[len(word)-3:])
	fmt.Printf("Middle characters: %s\n", word[2:5])

	// Rune and Unicode
	fmt.Println("\n=== Rune and Unicode ===")

	message := "Hello, 世界"
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Length in bytes: %d\n", len(message))
	fmt.Printf("Length in runes: %d\n", len([]rune(message)))

	for index, char := range message {
		fmt.Printf("Character %d: %c (Unicode: U+%04X)\n", index, char, char)
	}

	// String Builder
	fmt.Println("\n=== String Builder ===")

	var builder strings.Builder
	builder.WriteString("Hello ")
	builder.WriteString("Go ")
	builder.WriteString("World!")

	fmt.Printf("Built String: %s\n", builder.String())

	// Regular Expressions
	fmt.Println("\n=== Regular Expressions ===")

	// Regex Meanings
	fmt.Println("\n=== Regex Meanings ===")

	fmt.Println("^ - Start of string")
	fmt.Println("$ - End of string")
	fmt.Println(". - Any character")
	fmt.Println("* - Zero or more of the preceding character")
	fmt.Println("+ - One or more of the preceding character")
	fmt.Println("? - Zero or one of the preceding character")
	fmt.Println("[abc] - Any character inside the brackets")
	fmt.Println("[^abc] - Any character not inside the brackets")
	fmt.Println("[a-z] - Any lowercase letter")
	fmt.Println("[A-Z] - Any uppercase letter")
	fmt.Println("[0-9] - Any digit")
	fmt.Println("\\ - Escape character")
	fmt.Println("| - Alternation (OR)")
	fmt.Println("(...) - Grouping")

	// Regex Example: Validate Email
	fmt.Println("\n=== Regex Example: Validate Email ===")

	email := "someone@example.com"
	pattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(pattern, email)
	if err != nil {
		fmt.Println("Error:", err)
	} else if matched {
		fmt.Println("Valid email address")
	} else {
		fmt.Println("Invalid email address")
	}

	// Find all numbers
	fmt.Println("\n=== Find All Numbers ===")

	textWithNumbers := "The price is 100 dollars and 50 cents."
	numberPattern := `\d+`

	re := regexp.MustCompile(numberPattern)
	numbers := re.FindAllString(textWithNumbers, -1)

	fmt.Printf("Numbers found: %v\n", numbers)

	// Replace using regex
	fmt.Println("\n=== Replace Using Regex ===")

	textToReplace := "The quick brown fox jumps over the lazy dog."
	replacePattern := `\b\w{4}\b` // Replace all 4-letter words // \b - word boundary, \w{4} - exactly 4 word characters, \b - word boundary

	reReplace := regexp.MustCompile(replacePattern)
	replacedText := reReplace.ReplaceAllString(textToReplace, "[REPLACED]")

	fmt.Printf("Original Text: %s\n", textToReplace)
	fmt.Printf("Replaced Text: %s\n", replacedText)
}