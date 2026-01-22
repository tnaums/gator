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

func handlerLogin(s *state, cmd command) error {
	return nil
}
