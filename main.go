package main

import (
	"fmt"
	"gator/internal/config"
)


func main() {
	fmt.Println("Welcome to the gator!")
	myConfig := config.READ()
	fmt.Println(myConfig.DbURL)
}
