package controllers

import (
	"mercado-libre/inputs"
	"mercado-libre/models"
	"mercado-libre/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientRepository *repositories.ClientRepository
}

func InitClientController(repository *repositories.ClientRepository) *ClientController {
	controller := ClientController{
		clientRepository: repository,
	}

	return &controller
}

func (controller *ClientController) BatchCreate(inputs []inputs.ClientInput) (int, error) {
	var clients []models.Client

	for _, input := range inputs {
		client := models.Client{}
		client.FromInput(input)
		clients = append(clients, client)
	}

	return controller.clientRepository.BatchUpdate(clients)
}

func (controller *ClientController) GetClients(c *gin.Context) {
	clients := controller.clientRepository.GetClients()

	if len(clients) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No clients available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func (controller *ClientController) GetClientSensitive(c *gin.Context) {
	role := c.GetString("role")
	id, exists := c.GetQuery("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are missing the id parameter"})
		return
	}

	if role != "sensitive" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Your role does not allow see sensitive data"})
		return
	}
	client, err := controller.clientRepository.GetClientSensitiveData(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}

func (r *ClientController) Close() {
	r.clientRepository.Close()
}
