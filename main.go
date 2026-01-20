package main

import (
	"fmt"
	"gator/internal/config"
)


func main() {
	fmt.Println("Welcome to the gator!")
	myConfig, err := config.READ()
	if err != nil {
		fmt.Println(err)
	}

	myConfig.SetUser("tnaums")

	myConfig, err = config.READ()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Database URL: %s\n", myConfig.DbURL)
	fmt.Printf("Current user: %s\n", myConfig.CurrentUserName)
}
