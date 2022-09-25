package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addData(context *gin.Context) {
	if allowed(context.Param("auth")) {
		context.IndentedJSON(http.StatusNotFound, "ACCESS DENIED")
		return
	}

	id := context.Param("id")
	amount := context.Param("amount")
	stock := context.Param("stock")

	if !contains(games, id) {
		context.IndentedJSON(http.StatusNotFound, "NO PLAYER")
		return
	}

	game, _ := getGameINSTANCE(id)
	i, _ := strconv.Atoi(amount)

	if stock == "sp500" && game.money > i*50 {
		game.money -= i * 50
		game.stocks[0] = "sp500"
	}

	context.IndentedJSON(http.StatusOK, "BOUGHT "+amount+" "+stock)
}
