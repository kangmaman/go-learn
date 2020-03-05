package main

import (
	"fmt"
	"os"
)

func main() {
	//set the USERNAME environment vari
	os.Setenv("USERNAME", "kangmaman")

	username, exists := os.LookupEnv("USERNAME")
	if exists {
		//print the value of path env variable
		fmt.Printf("USERNAME: %s", username)
	}
}
