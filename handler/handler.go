package handler

import (
	"encoding/json"
	"exoplanet/models"
	"exoplanet/validator"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	exoplanets = make(map[string]models.Exoplanet)
	mu         sync.Mutex
)

func generateID() string {
	return strconv.Itoa(rand.Intn(1000000))
}

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validator.ValidateExoplanet(exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exoplanet.ID = generateID()

	mu.Lock()
	defer mu.Unlock()
	exoplanets[exoplanet.ID] = exoplanet

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
	
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	planets := make([]models.Exoplanet, 0, len(exoplanets))
	for _, exoplanet := range exoplanets {
		planets = append(planets, exoplanet)
	}

	json.NewEncoder(w).Encode(planets)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	defer mu.Unlock()
	exoplanet, exists := exoplanets[id]

	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedExoplanet models.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	var exoplanet models.Exoplanet
	var exists bool
	if exoplanet, exists = exoplanets[id]; !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
		if updatedExoplanet.Name != "" {
			exoplanet.Name = updatedExoplanet.Name
		}
		if updatedExoplanet.Type != "" {
			exoplanet.Type = updatedExoplanet.Type
		}
		if updatedExoplanet.Mass != 0 {
			exoplanet.Mass = updatedExoplanet.Mass
		}
		if updatedExoplanet.Radius != 0 {
			exoplanet.Radius = updatedExoplanet.Radius
		}
		if updatedExoplanet.Distance != 0 {
			exoplanet.Distance = updatedExoplanet.Distance
		}
	

	
	exoplanets[id] = exoplanet

	json.NewEncoder(w).Encode(exoplanet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := exoplanets[id]; !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	delete(exoplanets, id)
	w.WriteHeader(http.StatusNoContent)
}

func CalculateFuelEstimation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	crew, err := strconv.Atoi(r.URL.Query().Get("crew"))
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	mu.Lock()
	exoplanet, exists := exoplanets[id]
	defer mu.Unlock()

	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	var gravity float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else if exoplanet.Type == "Terrestrial" {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}
	fuelCost := float64(exoplanet.Distance) / (gravity * gravity) * float64(crew)
	estimation := models.FuelEstimation{FuelCost: fuelCost}

	json.NewEncoder(w).Encode(estimation)
}
