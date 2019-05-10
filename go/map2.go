package main

import (
	"fmt"
)

func main() {
	fmt.Println(new)
}

func new() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key", key, "value", value)

	}

}
