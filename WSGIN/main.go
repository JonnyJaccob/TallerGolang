package main

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)
type digimons struct{
	ID int
	Name string
}
var digimonList []digimons
func main(){

	num := 42
    strNum := strconv.Itoa(num)
    fmt.Printf("Número convertido a cadena: %s\n", strNum)

	router := gin.Default()

	//router.GET("/dig")

	router.GET("/dig",getDigimons)

	router.GET("/dig/:id",getDigimonsID)

	router.POST("/dig", postDigimons)

	digimonList = []digimons{
		{ID: 1, Name: "Agumon"},
		{ID: 2, Name: "Gabumon"},
	}

	router.Run("localhost:8083")
}

func getDigimons(c *gin.Context){
	c.IndentedJSON(http.StatusOK, digimonList)
}
func getDigimonsID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Convierte el parámetro de la URL a un entero
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID no válido"})
        return
    }

    digimon, err := findDigimonByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    } else {
        c.IndentedJSON(http.StatusOK, digimon)
    }
}
func postDigimons(c *gin.Context){ 
	var newDigimon digimons
	if err := c.BindJSON(&newDigimon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Agrega el nuevo Digimon al slice digimons.
	// Puedes generar un nuevo ID de alguna manera o usar uno incremental.
	newDigimon.ID = len(digimonList) + 1
	digimonList = append(digimonList, newDigimon)

	c.IndentedJSON(http.StatusCreated, newDigimon)
}


func findDigimonByID(id int) (digimons, error) {
    for _, digimon := range digimonList {
        if digimon.ID == id {
            return digimon, nil // Se encontró el Digimon
        }
    }
    // Si no se encuentra, devuelve un error
    return digimons{}, fmt.Errorf("Digimon con ID %d no encontrado", id)
}