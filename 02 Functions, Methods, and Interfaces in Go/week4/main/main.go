package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct {
	name string
	food string
	locomotion string
	noise string
}

type Bird struct {
	name string
	food string
	locomotion string
	noise string
}

type Snake struct {
	name string
	food string
	locomotion string
	noise string
}

func (b Bird) Eat()  {
	fmt.Println(b.food)
}
func (b Bird) Move()  {
	fmt.Println(b.locomotion)
}
func (b Bird) Speak()  {
	fmt.Println(b.noise)
}
func (b Bird) GetName() string  {
	return b.name
}

func (s Snake) Eat()  {
	fmt.Println(s.food)
}
func (s Snake) Move()  {
	fmt.Println(s.locomotion)
}
func (s Snake) Speak()  {
	fmt.Println(s.noise)
}
func (s Snake) GetName() string  {
	return s.name
}

func (c Cow) Eat()  {
	fmt.Println(c.food)
}
func (c Cow) Move()  {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak()  {
	fmt.Println(c.noise)
}
func (c Cow) GetName() string  {
	return c.name
}


func main() {
	arr := make([]Animal, 0)
	for {
		var command, name, action string
		fmt.Print("> ")
		fmt.Scan(&command)
		fmt.Scan(&name)
		fmt.Scan(&action)

		if command == "newanimal" {
			var temp Animal
			switch action {
			case "cow":
				temp = Cow{
					food: "grass",
					locomotion: "walk",
					noise: "moo",
					name: name,
				}
				break
			case "bird":
				temp = Bird{
					food: "worms",
					locomotion: "fly",
					noise: "peep",
					name: name,
				}
				break
			case "snake":
				temp = Snake{
					food: "mice",
					locomotion: "slither",
					noise: "hsss",
					name: name,
				}
				break
			}
			arr = append(arr, temp)
			fmt.Println("Created it!")
			
		} else {
			for _, animal := range arr{
				if animal.GetName() == name {
					switch action {
					case "eat":
						animal.Eat()
						break
					case "move":
						animal.Move()
						break
					case "speak":
						animal.Speak()
						break
					}
					break
				}
			}
		}
	}
}