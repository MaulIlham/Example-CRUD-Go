package controllers

import (
	"encoding/json"
	"fmt"
	"gamehsop/entities"
	"gamehsop/services"
	"gamehsop/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Welcome to game shop")
	utils.WriteLog("Request Homepage")
}

func getAllGames(w http.ResponseWriter, r *http.Request)  {
	games, err := services.GetAllDataGame()
	if err != nil {
		utils.WriteLog("Request all data game failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request all data game success")
		log.Println("Request success")
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(games)
}

func searchDataGames(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var game entities.Game
	json.Unmarshal(reqBody,&game)
	fmt.Printf("%s",game.Title)
	games, err := services.SearchDataGame(game.Title)
	if err != nil {
		utils.WriteLog("Request search data game failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request search data game success")
		log.Println("Request success")
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(games)
}

func getDataGameById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	game, err := services.GetDataGameById(id)
	if err != nil {
		utils.WriteLog("Request data game By ID failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request data game By ID success")
		log.Println("Request success")
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(game)
}

func saveDataGame(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var game entities.Game
	json.Unmarshal(reqBody, &game)
	result, err := services.SaveDataGame(game)
	if err != nil {
		utils.WriteLog("Request save data game failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request save data game success")
		log.Println("Request success")
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(result)
}

func updateDataGame(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var game entities.Game
	json.Unmarshal(reqBody,&game)
	err := services.UpdateDataGame(game)
	if err != nil {
		utils.WriteLog("Request update data game failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request update data game success")
		log.Println("Request success")
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(game)
}

func deleteDataGame(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	err := services.DeleteDataGame(id)
	if err != nil {
		utils.WriteLog("Request delete data game failed!")
		log.Println("Request failed!")
	}else {
		utils.WriteLog("Request delete data game success")
		log.Println("Request success")
	}
	json.NewEncoder(w).Encode("Delete Data Succes")
}

func HandleRequest()  {
	log.Println("Starting Server at http://localhost:8844")
	utils.WriteLog("Starting Server in port : 8844")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/games", getAllGames).Methods("GET")
	router.HandleFunc("/games/find", searchDataGames).Methods("GET")
	router.HandleFunc("/game/{id}", getDataGameById).Methods("GET")
	router.HandleFunc("/game", saveDataGame).Methods("POST")
	router.HandleFunc("/game", updateDataGame).Methods("PUT")
	router.HandleFunc("/game/{id}", deleteDataGame).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":8844",router))
}
