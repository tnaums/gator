package main

import (
	"fmt"
	"gator/internal/config"
	"os"
	"database/sql"
)

import _ "github.com/lib/pq"

func main() {
	fmt.Println("Welcome to the gator!")
	myConfig, err := config.READ()
	if err != nil {
		fmt.Println(err)
	}
	myState := state{configPointer: &myConfig}

	fmt.Printf("Database URL: %s\n", myConfig.DbURL)
	db, err := sql.Open("postgres", myConfig.DbURL)
	dbQueries := database.New(db)	
	fmt.Printf("Current user: %s\n", myConfig.CurrentUserName)
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	command := NewCommand(os.Args)

	myCommands := NewCommands()
	myCommands.register("login", handlerLogin)
	err2 := myCommands.run(&myState, command)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}

}
