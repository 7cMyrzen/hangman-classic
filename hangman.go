package main

import (
	"fmt"
	"math/rand"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	words := string(content)