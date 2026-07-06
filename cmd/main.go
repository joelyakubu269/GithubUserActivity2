package main

import (
	"flag"
	"fmt"
	"githubUserActivity/github"
)

func main() {
	var user string
	flag.StringVar(&user, "username", "kamranahmedse", "username of github user")
	flag.Parse()
	events, err := github.FetchUserActivity(user)
	if err != nil {
		fmt.Printf("had error %v\n", err)
	}
	val, err := github.Formatter(events)
	if err != nil {
		fmt.Printf("had error %v\n", err)

	}
	fmt.Println(val)
}
