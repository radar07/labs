package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v64/github"
)

func main() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByUserOptions{Sort: "full_name"}

	repos, response, err := client.Repositories.ListByUser(context.Background(), "radar07", opt)

	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(repos[0])
	fmt.Println(response)
}
