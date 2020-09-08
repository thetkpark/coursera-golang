package main

import (
	"fmt"
	"strings"
	"bufio"
  	"os"
)

func main () {

	inputReader := bufio.NewReader(os.Stdin)
    input, _ := inputReader.ReadString('\n')
	var lower = strings.ToLower(input)

	if strings.Index(lower, "i") == 0 && strings.LastIndex(lower, "n") == len(lower)-2 && strings.Contains(lower, "a") {
		fmt.Printf("Found!")
		return
	}
	fmt.Printf("Not Found!")
	
}