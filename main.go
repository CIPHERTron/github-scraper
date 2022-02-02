package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Details struct {
	Name        string
	Description string
	Stars       int
	Forks       int
	License     github.License
}

func main() {
	// Loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file")
	}

	var accessKey string = os.Getenv("ACCESS_TOKEN")

	ctx := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessKey},
	)

	tokenClient := oauth2.NewClient(ctx, tokenService)

	// go-github client
	client := github.NewClient(tokenClient)

	repo, _, err := client.Repositories.Get(ctx, "CIPHERTron", "kaagaz")

	if err != nil {
		log.Fatalf("Error getting organizations list")
	}

	data := &Details{
		Name:        *repo.FullName,
		Description: *repo.Description,
		Stars:       *repo.StargazersCount,
		Forks:       *repo.ForksCount,
		License:     *repo.GetLicense(),
	}

	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(dataJSON))

}
