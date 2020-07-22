package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func startHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var resp HTTPResponse
	resp.Code = 200
	var ng Game
	ng.PlayerTurn = 1
	ID := generateGameID()
	for gameExists(ID) {
		ID = generateGameID()
	}
	resp.Response = []string{strconv.FormatInt(ID, 10)}
	addGame(ID, ng)
	by, _ := json.Marshal(resp)
	fmt.Fprintln(w, string(by))
}

func moveHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tableID, _ := strconv.ParseInt(r.FormValue("gid"), 10, 64)
	player, _ := strconv.ParseInt(r.FormValue("player"), 10, 64)
	move, _ := strconv.ParseInt(r.FormValue("move"), 10, 64)
	game, err := getGameByID(tableID)
	var hresp HTTPResponse
	if move < 1 || move > 7 || err != nil || game.Table[0][move] != 0 || int(player) != game.PlayerTurn {
		hresp = HTTPResponse{400, []string{"Erorr"}}
	} else {
		line := 0
		for line < 6 && game.Table[line][move-1] == 0 {
			line++
		}
		line--
		game.Table[line][move-1] = game.PlayerTurn
		if game.PlayerTurn == 1 {
			game.PlayerTurn = 2
		} else {
			game.PlayerTurn = 1
		}
		updateTable(tableID, game)
		hresp = HTTPResponse{200, []string{"Ok"}}
		fmt.Println(resultFromGame(game))
	}
	by, _ := json.Marshal(hresp)
	fmt.Fprintln(w, string(by))
}

func readHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	gID, _ := strconv.ParseInt(r.FormValue("gid"), 10, 64)
	game, _ := getGameByID(gID)
	bt, _ := json.Marshal(game)
	fmt.Fprintln(w, string(bt))
}
