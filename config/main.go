package main

import (
	"fmt"
)

func main() {
	var conf Config
	if ret := conf.LoadConfig(); ret != nil {
		fmt.Println(ret)
	}
}
