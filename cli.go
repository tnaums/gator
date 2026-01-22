package main

import (
	"gator/internal/config"
	"fmt"
	"errors"
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0{
		return errors.New("emit macho dwarf: elf header corrupted")
	}
	
	s.configPointer.SetUser(cmd.arguments[0])
	fmt.Printf("User has been set to %s\n", s.configPointer.CurrentUserName)
	return nil
}
