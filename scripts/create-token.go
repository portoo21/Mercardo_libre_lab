package main

import (
	"fmt"
	"mercado-libre/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Print("Argument role is missing")
		return
	}

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Bad thing happened! %v", err)
		return
	}

	token, err := utils.GenerateToken(1, args[1])

	if err != nil {
		fmt.Printf("Bad thing happened! %v", err)
		return
	}

	fmt.Print(token)

	os.Exit(0)
}
