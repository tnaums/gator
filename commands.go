package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
// package main

// import (
// 	"gator/internal/config"
// 	"errors"
// 	"gator/internal/database"
// )

// // type state struct {
// // 	db  *database.Queries
// // 	cfg  *config.Config
// // }

// type command struct{
// 	name string
// 	arguments []string
// }

// // NewCommand
// func NewCommand(clargs []string) command {
// 	return command{
// 	 	name: clargs[1],
// 		arguments: clargs[2:len(clargs)],
//  	}
// }

// type commands struct {
// 	handlers    map[string]func(*state, command) error
// }
// func NewCommands() *commands {
// 	var c commands
// 	c.handlers = make(map[string]func(*state, command) error)
// 	return &c
// }


// func (c *commands) run(s *state, cmd command) error {
// 	myFunc, ok := c.handlers[cmd.name]
// 	if !ok {
// 		return errors.New("Command not found")
// 	}
// 	return myFunc(s, cmd)
// }

// func (c *commands) register(name string, f func(*state, command) error) {
// 	c.handlers[name] = f
// }

