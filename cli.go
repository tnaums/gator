package main

import (
	"gator/internal/config"
)
type state struct {
	configPointer  *config.Config
}

type command struct{
	name string
	arguments []string
}

// NewCommand
func NewCommand(clargs []string) command {
	//	fmt.Println(clargs[1])
	//	fmt.Println(clargs[2:len(clargs)])
	return command{
	 	name: clargs[1],
		arguments: clargs[2:len(clargs)],
 	}
}

func handlerLogin(s *state, cmd command) error {
	return nil
}
