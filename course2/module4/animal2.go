package main

// Loading Golang libraries...
import (
	"fmt"
	"strings"
)

// Creating the 'Animal' interface
type Animal interface {
	Eat()
	Speak()
	Move()
}

// Define type Cow
type Cow struct{ name string }

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Speak() { fmt.Println("moo") }
func (c Cow) Move()  { fmt.Println("walk") }

// Define type Bird
type Bird struct{ name string }

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Speak() { fmt.Println("peep") }
func (b Bird) Move()  { fmt.Println("fly") }

// Define type Snake
type Snake struct{ name string }

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Speak() { fmt.Println("slither") }
func (s Snake) Move()  { fmt.Println("hsss") }

// Entry point of the program
func main() {
	var command, name, info string
	newAnimal := map[string]Animal{}
	for true {
		fmt.Printf("> ")
		fmt.Scanln(&command, &name, &info)
		switch command {
		case "newanimal":
			newAnimal[name] = AnimalInit(info)
		case "query":
			AnimalQuery(newAnimal, name, info)
		default:
			fmt.Println("Bad request. Try Again with the commands 'newanimal <animal_name> <animal_type>' or 'query <animal_name> <info>'...")
		}
		command = ""
		name = ""
		info = ""
	}
}

// This function checks an Animal map and returns the requested info on an animal
func AnimalQuery(m map[string]Animal, name, query string) {
	lname := strings.ToLower(name)
	lquery := strings.ToLower(query)
	switch lquery {
	case "eat":
		if animal, err := m[lname]; err {
			animal.Eat()
		} else {
			fmt.Printf("The Animal named '%s' has not been created! Please create it with the 'newanimal' command!\n", lname)
		}
	case "move":
		if animal, err := m[lname]; err {
			animal.Move()
		} else {
			fmt.Printf("The Animal named '%s' has not been created! Please create it with the 'newanimal' command!\n", lname)
		}
	case "speak":
		if animal, err := m[lname]; err {
			animal.Speak()
		} else {
			fmt.Printf("The Animal named '%s' has not been created! Please create it with the 'newanimal' command!\n", lname)
		}
	default:
		fmt.Println("Bad animal info! Animal info can only be 'eat', 'move', or 'speak'. Try Again...")
	}
	name = ""
	query = ""
}

// This function initializes an animal
func AnimalInit(name string) Animal {
	var lname = strings.ToLower(name)
	var animal Animal
	switch lname {
	case "cow":
		animal = Cow{lname}
		fmt.Println("Created it!")
		return animal
	case "bird":
		animal = Bird{lname}
		fmt.Println("Created it!")
		return animal
	case "snake":
		animal = Snake{lname}
		fmt.Println("Created it!")
		return animal
	default:
		fmt.Println("Bad animal type! Animal type can only be 'cow', 'bird' or 'snake'!")
		return nil
	}
}
