package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"mercado-libre/controllers"
	"mercado-libre/inputs"
	"mercado-libre/repositories"
	"mercado-libre/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Print("Argument token is missing")
		os.Exit(1)
		return
	}

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Fail to load env vars %v", err)
		os.Exit(1)
		return
	}

	token := args[1]

	role, err := utils.TokenValid(token)

	if err != nil {
		fmt.Printf("Fail to validate token %v", err)
		os.Exit(1)
		return
	}

	if role != "sensitive" {
		fmt.Println("The role does not have permission")
		os.Exit(1)
		return
	}

	clientRepository, err := repositories.InitClientRepository()

	if err != nil {
		fmt.Printf("Fail to get db connection! %v", err)
		os.Exit(1)
		return
	}

	controller := controllers.InitClientController(clientRepository)

	res, err := http.Get("https://62433a7fd126926d0c5d296b.mockapi.io/api/v1/usuarios")
	if err != nil {
		fmt.Printf("Fail to retrieve client list %v", err)
		os.Exit(1)
		return
	}
	data, _ := ioutil.ReadAll(res.Body)
	// due to no support for array directly
	supportedJson := "{ \"value\": " + string(data) + "}"

	read := io.NopCloser(strings.NewReader(supportedJson))

	req, _ := http.NewRequest(http.MethodGet, "https://62433a7fd126926d0c5d296b.mockapi.io/api/v1/usuarios", read)

	inputs := inputs.BatchClientInput{}

	err1 := binding.JSON.Bind(req, &inputs)

	if err1 != nil {
		fmt.Printf("Fail to format clients %v", err1)
		os.Exit(1)
		return
	}

	count, err := controller.BatchCreate(inputs.Value)

	if err != nil {
		fmt.Printf("Fail to store clients %v", err)
		os.Exit(1)
		return
	}

	fmt.Print("Saved clients: ", count)

	os.Exit(0)
}
