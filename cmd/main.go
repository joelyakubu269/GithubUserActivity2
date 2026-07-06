package main

import (
	"flag"
	"fmt"
	"githubUserActivity/github"
)

func main() {
	flag.Usage = func() {
		fmt.Println("GitHub User Activity CLI")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  go run ./cmd -username=<github-username>")
		fmt.Println()
		fmt.Println("Example:")
		fmt.Println("  go run ./cmd -username=joelyakubu269")
		fmt.Println()
		flag.PrintDefaults()
	}
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
