package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func main() {
	for {
		var name, action string
		fmt.Scan(&name)
		fmt.Scan(&action)
		var a Animal
		switch name {
		case "cow":
			a = Animal{"grass", "walk", "moo"}
		case "bird":
			a = Animal{"worms", "fly", "peep"}
		case "snake":
			a = Animal{"mice", "slither", "hsss"}
		}

		switch action {
		case "eat":
			a.Eat()
			continue
		case "move":
			a.Move()
			continue
		case "speak":
			a.Speak()
			continue
		}
	}
}
