package main

import (
	"fmt"
	"os"
)

func main() {
	//store the path environment variable in a variable
	path, exists := os.LookupEnv("PATH")
	if exists {
		//print the value of path env variable
		fmt.Println("path : %s", path)
	}
}
