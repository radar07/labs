package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Value struct {
	value  string
	expiry int
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := []byte(`
		{
			"name": "pranav",
			"age": 23
		}
	`)
	var p person
	err := json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)

	// Integers
	x, y := 2, 3
	fmt.Println(x, y)
	x, y = y, x

	var n = new(int)
	*n = 5
	fmt.Println(*n)

	var bd = 0b11000
	fmt.Println(bd)

	var i int8 = 127
	fmt.Println(i)

	// Arrays
	var arr1 = [10]int{1, 3: 24, 9: 10, 4: 5}

	fmt.Println(arr1)

	// Strings
	var world = "🌎"

	var runeWorld = []rune(world)
	fmt.Println(len(runeWorld))
	fmt.Println(cap(runeWorld))

	// Concurrency
	_ = sync.WaitGroup{}
}
