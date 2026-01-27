package main

import (
	"database/sql"
	"log"
	"os"

	"gator/internal/config"
	"gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"gator/internal/config"
// 	"gator/internal/database"
// 	_ "github.com/lib/pq"
// 	"log"
// 	"os"
// )

// func main() {
// 	fmt.Println("Welcome to the gator!")
// 	myConfig, err := config.READ()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Printf("Database URL: %s\n", myConfig.DbURL)
// 	db, err := sql.Open("postgres", myConfig.DbURL)
// 	if err != nil {
// 		log.Fatalf("error connecting to db: %v", err)
// 	}
// 	defer db.Close()

// 	dbQueries := database.New(db)
// 	myState := &state{
// 		db:  dbQueries,
// 		cfg: &myConfig,
// 	}

// 	fmt.Printf("Current user: %s\n", myConfig.CurrentUserName)
// 	if len(os.Args) < 2 {
// 		fmt.Println("not enough arguments")
// 		os.Exit(1)
// 	}
// 	command := NewCommand(os.Args)

// 	myCommands := NewCommands()
// 	myCommands.register("login", handlerLogin)
// 	myCommands.register("run", handlerRegister)
// 	err2 := myCommands.run(myState, command)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		os.Exit(1)
// 	}

// }
