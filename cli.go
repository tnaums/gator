package main

import (
	"gator/internal/config"
	"fmt"
	//	"errors"
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
	return command{
	 	name: clargs[1],
		arguments: clargs[2:len(clargs)],
 	}
}

type commands struct {
	handlers    map[string]func(*state, command) error
}


func (c *commands) run(s *state, cmd command) error {
	fmt.Println("Running...")
	return nil
}

func (c *commands) register(s *state, cmd command) error {
	fmt.Println("Registering...")
	c.handlers["login"] = handlerLogin
	return nil
}

