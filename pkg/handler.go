package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type General struct {
	Name      string `json:"name"`
	Country   string `json:"country"`
	Agein1939 int    `json:"age"`
}

var generals = []General{
	{"Bernard Montgomery", "Great Britain", 51},
	{"Arthur Harris", "Great Britain", 46},
	{"George Patton", "USA", 53},
	{"Dwight Eisenhower", "USA", 48},
	{"Georgy Zhukov","USSR", 43},
	{"Konstantin Rokossovsky", "USSR", 43},
	{"Erwin Rommel", "Germany", 47},
	{"Heinz Guderian" ,"Germany", 51},
}

func GetGenerals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generals)
}

func GetGeneralByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, general := range generals {
		if general.Name == params["name"] {
			json.NewEncoder(w).Encode(general)
			return
		}
	}
	http.NotFound(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("WW2 Generals App is healthy!"))
}