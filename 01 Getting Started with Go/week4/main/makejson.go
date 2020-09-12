package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	scanner.Scan()
	address := scanner.Text()

	user := &User{Name: name, Address: address}
	str, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(str))
}