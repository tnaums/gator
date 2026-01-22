package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)



func main() {
	fmt.Println("Welcome to the gator!")
	myConfig, err := config.READ()
	if err != nil {
		fmt.Println(err)
	}
	
	myState := state{configPointer: &myConfig}

	fmt.Printf("Database URL: %s\n", myConfig.DbURL)
	fmt.Printf("Current user: %s\n", myConfig.CurrentUserName)
	command := NewCommand(os.Args)

	err2 := handlerLogin(&myState, command)
	if err2 != nil {
		fmt.Println(err2)
	}

}
