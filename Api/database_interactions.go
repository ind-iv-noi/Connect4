package main

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func getGameByID(gID int64) (Game, error) {
	gameQuery, err := db.Query(fmt.Sprintf("SELECT PlayerTurn,GameTable FROM connect4 WHERE GID=%v", gID))
	if err != nil {
		return Game{}, err
	}
	if gameQuery.Next() {
		var game Game
		table := pq.Int64Array{}
		gameQuery.Scan(&game.PlayerTurn, &table)
		game.Table = arrayToMatrix(table)
		return game, nil
	}
	return Game{}, errors.New("Game does not exist")
}

func gameExists(gID int64) bool {
	_, err := getGameByID(gID)
	if err != nil {
		return false
	}
	return true
}

func addGame(gID int64, game Game) {
	sqlStatement := `INSERT INTO connect4 (PlayerTurn,GameTable,Gid) VALUES($1,$2,$3)`
	_, err := db.Exec(sqlStatement, 1, pq.Array(matrixToArray(game.Table)), gID)
	if err != nil {
		panic(err)
	}
}

func updateTable(gID int64, game Game) {
	db.Exec(`UPDATE connect4 SET PlayerTurn = $1 WHERE gid=$2`, game.PlayerTurn, gID)
	db.Exec(`UPDATE connect4 SET GameTable = $1 WHERE gid=$2`, pq.Array(matrixToArray(game.Table)), gID)
}

func deleteGameByID(gID int64) {

}
