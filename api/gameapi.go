package api

import (
	"log"
)

func AllGames() interface{} {
	body := make(map[string]interface{})
	response := Post(GameList, body)
	log.Println(5, response.Code)
	if response.IsSuccess() {
		log.Println(6, response.Data)
		return response.Data
	}
	return ""
}
