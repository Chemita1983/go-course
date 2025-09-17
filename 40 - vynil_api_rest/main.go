// Utilizaremos Gin Framework que es un framework web para desarrollar aplicaciones webs y apis

// go get -u github.com/gin-gonic/gin

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []album{
	{ID: "1", Title: "Abbey Road", Artist: "The Beatles", Price: 29.99},
	{ID: "2", Title: "Dark Side of the Moon", Artist: "Pink Floyd", Price: 34.95},
	{ID: "3", Title: "Rumours", Artist: "Fleetwood Mac", Price: 27.50},
	{ID: "4", Title: "Thriller", Artist: "Michael Jackson", Price: 31.00},
	{ID: "5", Title: "Back to Black", Artist: "Amy Winehouse", Price: 25.75},
}

func main() {

	router := gin.Default() // Trae configuraciones por defecto para trabajar con gin

	// Peticion GET, que recibe el endpoint y unos handlers, donde devolvemos un json con
	// mensaje y un http status code 200
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Bienvenido a Vinyl-API",
		})
	})

	router.GET("/album/:id", getAlbumById)

	router.GET("/albums", getAlbums)

	router.POST("/album", postAlbum)

	router.PUT("/album/:id", editAlbum)

	router.DELETE("/album/:id", deleteAlbum)

	// Le indicamos donde queremos levantar nuestra aplicacion
	router.Run("localhost:8080")
}

// Funcion que devuelve un album por id (GET)
func getAlbumById(ctx *gin.Context) {

	id := ctx.Param("id")

	for _, value := range albums {
		if value.ID == id {
			ctx.IndentedJSON(http.StatusOK, value)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Funcion que devuelve todos los albums en tipo json (GET)
func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

// Funcion que ingresa un nuevo album (POST)
func postAlbum(ctx *gin.Context) {
	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, albums)
}

// Funcion para editar un album (PUT)
func editAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, value := range albums {
		if value.ID == id {
			var newAlbum album
			if err := ctx.BindJSON(&newAlbum); err != nil {
				return
			}
			albums[index] = newAlbum
			ctx.IndentedJSON(http.StatusOK, albums)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Funcion para borrar un album (DELETE)
func deleteAlbum(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, value := range albums {
		if value.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			ctx.IndentedJSON(http.StatusNoContent, "")
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}
