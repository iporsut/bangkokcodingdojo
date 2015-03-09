package main

import (
	"bangkokcodingdojo/fizzbuzz"
	"fmt"
	"strings"
)

func main() {
	var start, finish int

	fmt.Print("Start : ")
	fmt.Scanf("%d", &start)

	fmt.Print("Finish : ")
	fmt.Scanf("%d", &finish)

	fmt.Println(strings.Join(fizzbuzz.Range(start, finish), ","))
}
