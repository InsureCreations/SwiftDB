package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDataValue(key string) (*Game, error) {
	for i, t := range games {
		if t.gameID == key {
			return &games[i], nil
		}
	}

	return nil, errors.New("404: VALUE NOT FOUND")
}

func getData(context *gin.Context) {
	id := context.Param("id")

	game, _ := getDataValue(id)

	context.IndentedJSON(http.StatusOK, game.stocks)
}
