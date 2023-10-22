package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Hello World
	fmt.Println("Hello, World!")

	// Variables and Data Types
	var name string = "John"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Control Structures
	score := 85
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Arrays and Slices
	var fruits [3]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Cherry"

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println("Fruits:", fruits)
	fmt.Println("Colors:", colors)

	// Functions
	message := greet("Alice")
	fmt.Println(message)

	// Structs
	person := Person{Name: "Bob", Age: 25}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)

	// Maps
	grades := make(map[string]int)
	grades["Alice"] = 95
	grades["Bob"] = 87
	grades["Charlie"] = 92

	aliceGrade := grades["Alice"]
	fmt.Println("Alice's Grade:", aliceGrade)

	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}

func greet(name string) string {
	return "Hello, " + name + "!"
}
