package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("iran.txt")
	if err != nil {
		fmt.Print("error")

	}
	file.WriteString("123456789")
}
