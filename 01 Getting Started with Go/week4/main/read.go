package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	fmt.Scan(&fileName)

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(raw))
	str := string(raw)
	line := strings.Split(str, "\n")

	lists := make([]Name, 0)

	for i:= 0; i < len(line); i++ {
		tempArr := strings.Split(line[i], " ")
		temp := Name{
			fname: tempArr[0],
			lname: tempArr[1],
		}
		lists = append(lists, temp)
	}

	for i:= 0; i< len(lists); i++ {
		fmt.Println(lists[i].fname + " " + lists[i].lname)
	}
}