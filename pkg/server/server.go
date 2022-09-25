package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	gameID string   `json:"gameid"`
	money  int      `json:"money"`
	stocks []string `json:"stocks"`
}

var games = []Game{
	{gameID: "abc123", money: 123000},
	{gameID: "bob", money: 123},
}

func allowed(e string) bool {
	if e != "b^H80a60r4a4" {
		return true
	}
	return false
}

func contains(s []Game, str string) bool {
	for _, v := range s {
		if v.gameID == str {
			return true
		}
	}

	return false
}

func addGame(context *gin.Context) {
	var newGame Game

	if err := context.BindJSON(&newGame); err != nil {
		return
	}

	if !contains(games, newGame.gameID) {
		games = append(games, newGame)
	}

	context.IndentedJSON(http.StatusCreated, newGame)

}

func getGameINSTANCE(GameID string) (*Game, error) {
	for i, t := range games {
		if t.gameID == GameID {
			return &games[i], nil
		}
	}

	return nil, errors.New("404: PLAYER NOT FOUND")
}

func getGameByID(context *gin.Context) {
	if allowed(context.Param("auth")) {
		context.IndentedJSON(http.StatusNotFound, "ACCESS DENIED")
		return
	}

	id := context.Param("id")

	player, err := getGameINSTANCE(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, "NOT FOUND")
	}

	context.IndentedJSON(http.StatusOK, player.money)
}

func startGame(context *gin.Context) {
	if allowed(context.Param("auth")) {
		context.IndentedJSON(http.StatusNotFound, "ACCESS DENIED")
		return
	}

	id := context.Param("id")

	newGame := Game{gameID: id, money: 100}

	if !contains(games, newGame.gameID) {
		games = append(games, newGame)
	}

	context.IndentedJSON(http.StatusOK, "CREATED")
}

func Start() {
	router := gin.Default()

	router.GET("/getValue/:id")
	router.GET("/addValue/:id/:value")
	router.GET("/changeValue/:id/:value")

	router.GET("/getGame/:auth/:id", getGameByID)

	router.GET("/startGame/:auth/:id", startGame)

	router.GET("/addMoney/:auth/:id/:amount", addMoney)

	router.GET("/showStocks/:auth/:id", getStocks)

	router.GET("/buyStock/:auth/:stock/:id/:amount", buyStock)

	router.Run("127.0.0.1:9090")

}
