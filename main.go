package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "Jhon Coltrane", Price: 56.99},
	{ID: "2", Title: "Blue", Artist: "Mcallister", Price: 17.99},
	{ID: "3", Title: "Train", Artist: "Coltrane", Price: 88.15},
}

func getAlbums(c *gin.Context) {
	//Serializamos el slice de album a JSON
	c.IndentedJSON(http.StatusOK, albums)
}
func postAlbum(c *gin.Context) {
	//Serializamos el slice de album a JSON
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	//Serializamos el slice de album a JSON
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album no encontrado"})
}

func putAlbum(c *gin.Context) {
	// Obtenemos el ID de la URL
	id := c.Param("id")

	// Buscamos el álbum en la colección
	var foundAlbum *album
	for i, a := range albums {
		if a.ID == id {
			foundAlbum = &albums[i]
			break
		}
	}

	if foundAlbum == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Álbum no encontrado"})
		return
	}

	// Decodificamos el cuerpo JSON de la solicitud en una estructura auxiliar
	var updatedAlbum album
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Solicitud JSON no válida"})
		return
	}

	// Actualizamos los campos del álbum encontrado con los nuevos valores
	foundAlbum.Title = updatedAlbum.Title
	foundAlbum.Artist = updatedAlbum.Artist
	foundAlbum.Price = updatedAlbum.Price

	c.IndentedJSON(http.StatusOK, foundAlbum)
}

func deleteAlbum(c *gin.Context) {
	//Serializamos el slice de album a JSON
	var id = c.Param("id")
	var foundIndex = -1
	for i, a := range albums {
		if a.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Álbum no encontrado"})
		return
	}
	albums = append(albums[:foundIndex], albums[foundIndex+1:]...)
	c.IndentedJSON(http.StatusAccepted, gin.H{"Message": "Album eliniado de forma correcta"})
}

func main() {
	fmt.Println("Bienvenido a Vinyl-api")
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/album/:id", getAlbumById)
	router.PUT("/album/:id", putAlbum)
	router.DELETE("/album/:id", deleteAlbum)
	router.Run("localhost:8080")
}
