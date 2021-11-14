package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func main() {
	var cow, bird, snake Animal
	var name, info string
	cow.AnimalInit("grass", "walk", "moo")
	bird.AnimalInit("worms", "fly", "peep")
	snake.AnimalInit("mice", "slither", "hsss")

	commands := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	for true {
		fmt.Printf("> ")
		fmt.Scanln(&name, &info)
		switch info {
		case "eat":
			commands[name].Eat()
		case "move":
			commands[name].Move()
		case "speak":
			commands[name].Speak()
		default:
			fmt.Println("Bad request. Try Again...")
		}
		name = ""
		info = ""
	}
}

func (a *Animal) AnimalInit(eat, move, speak string) {
	a.food = eat
	a.locomotion = move
	a.noise = speak
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
