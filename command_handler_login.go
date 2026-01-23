package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0{
		return errors.New("emit macho dwarf: elf header corrupted")
	}
	
	s.configPointer.SetUser(cmd.arguments[0])
	fmt.Printf("User has been set to %s\n", s.configPointer.CurrentUserName)
	return nil
}
