package main

import (
	"fmt"
	"sync"
)

type InMemoryDB struct {
	Store map[string]map[string]Value
}

type Value struct {
	value  string
	expiry int
}

type person struct {
	name string
	age  int
}

func main() {
	// db := make(map[string]map[string]Value)
	// db["a"]["b"] = Value{value: "pranav"}
	// fmt.Printf("%#v\n", db)

	p := person{name: "Pranav", age: 25}
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

	// Loops - for loops only
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	switch 1 {
	case 2:
		fmt.Println("true")
	case 3:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}
}
