package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (playerServer *PlayerServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	switch request.Method {
	case http.MethodGet:
		playerServer.showScore(writer, player)
	case http.MethodPost:
		playerServer.processWin(writer, player)
	}
}

func (playerServer *PlayerServer) showScore(writer http.ResponseWriter, player string) {
	score := playerServer.store.GetPlayerScore(player)
	if score == 0 {
		writer.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(writer, score)
}

func (playerServer *PlayerServer) processWin(writer http.ResponseWriter, player string) {
	playerServer.store.RecordWin(player)
	writer.WriteHeader(http.StatusAccepted)
}
