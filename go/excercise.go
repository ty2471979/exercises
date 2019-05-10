package main

import (
	"fmt"
)

func main() {
	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 20

	fmt.Println(vertices)

}
