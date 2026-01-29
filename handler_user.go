package main

import (
	"context"
	"fmt"
	"time"
	"html"

	"gator/internal/database"
	"github.com/google/uuid"
)

func handlerListFeeds(s *state, cmd command) error {
	fmt.Println("Listing Feeds...")

	Feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get list of feeds: %w", err)
	}

	fmt.Println()
	for _, feed := range Feeds {
		fmt.Println("----------------------------------------")
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		u, err := s.db.GetByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user for feed: %w", err)
		}
		
		fmt.Printf("* %s\n", u.Name)
		fmt.Println("----------------------------------------")
	}

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) !=2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}
	cu, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get user %s: %v", s.cfg.CurrentUserName, err)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	f, err := s.db.CreateFeeds(context.Background(), database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    cu.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed %s: %w", name, err)
	}
	fmt.Printf("ID: %s\n", f.ID)
	fmt.Printf("Created At: %s\n", f.CreatedAt)
	fmt.Printf("Updated At: %s\n", f.UpdatedAt)
	fmt.Printf("Name: %s\n", f.Name)
	fmt.Printf("Url: %s\n", f.Url)
	fmt.Printf("UserID: %s\n", f.UserID)

	return nil
}

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %v <name>", cmd.Name)
	// }
	
	ctx, _ := context.WithCancel(context.Background())
	//	r, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	r, err := fetchFeed(ctx, "https://hnrss.org/newest")
	//	r, err := fetchFeed(ctx, cmd.Args[0])	
	if err != nil {
		return fmt.Errorf("failed to retrieve rss feed: %w", err)
	}

	fmt.Printf("Title: %#v\n", html.UnescapeString(r.Channel.Title))
	fmt.Printf("Link: %#v\n", r.Channel.Link)
	fmt.Printf("Description: %#v\n", r.Channel.Description)
	for _, item := range r.Channel.Item {
		fmt.Printf("Title: %#v\n", html.UnescapeString(item.Title))
		fmt.Printf("Title: %#v\n", html.UnescapeString(item.Description))
	}
	return nil
}

func handlerUsers(s *state, cmd command) error {
	Users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get list of users: %w", err)
	}

	fmt.Println()
	for _, user := range Users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf(" (current)\n")
		} else {
			fmt.Printf("\n")
		}
	}
	fmt.Println()
	
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset database: %w", err)
	}
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
