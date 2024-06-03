package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"exoplanet/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestAddExoplanet(t *testing.T) {
	exoplanet := models.Exoplanet{Name: "Test Planet", Description: "test descpricption", Type: "Terrestrial", Mass: 1.0, Radius: 1.0, Distance: 100}
	body, _ := json.Marshal(exoplanet)

	req, err := http.NewRequest("POST", "/exoplanets", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddExoplanet)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response models.Exoplanet
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, exoplanet.Name, response.Name)
	assert.NotEmpty(t, response.ID)
}

func TestListExoplanets(t *testing.T) {
	req, err := http.NewRequest("GET", "/exoplanets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListExoplanets)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []models.Exoplanet
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
}

func TestGetExoplanetByID(t *testing.T) {
	exoplanet := models.Exoplanet{ID: "12345", Name: "Test Planet",Description: "test descpricption", Type: "Terrestrial", Mass: 1.0, Radius: 1.0, Distance: 100}
	exoplanets[exoplanet.ID] = exoplanet

	req, err := http.NewRequest("GET", "/exoplanets/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets/{id}", GetExoplanetByID)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.Exoplanet
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, exoplanet.Name, response.Name)
}

func TestUpdateExoplanet(t *testing.T) {
	exoplanet := models.Exoplanet{ID: "12345", Name: "Test Planet",Description: "test descpricption", Type: "Terrestrial", Mass: 1.0, Radius: 1.0, Distance: 100}
	exoplanets[exoplanet.ID] = exoplanet

	updatedExoplanet := models.Exoplanet{Name: "Updated Planet",Description: "test descpricption", Type: "Terrestrial", Mass: 2.0, Radius: 1.0, Distance: 150}
	body, _ := json.Marshal(updatedExoplanet)

	req, err := http.NewRequest("PUT", "/exoplanets/12345", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets/{id}", UpdateExoplanet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.Exoplanet
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, updatedExoplanet.Name, response.Name)
}

func TestDeleteExoplanet(t *testing.T) {
	exoplanet := models.Exoplanet{ID: "12345", Name: "Test Planet", Description: "test descpricption",Type: "Terrestrial", Mass: 1.0, Radius: 1.0, Distance: 100}
	exoplanets[exoplanet.ID] = exoplanet

	req, err := http.NewRequest("DELETE", "/exoplanets/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets/{id}", DeleteExoplanet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	_, exists := exoplanets["12345"]
	assert.False(t, exists)
}

func TestCalculateFuelEstimation(t *testing.T) {
	exoplanet := models.Exoplanet{ID: "12345", Name: "Test Planet",Description: "test descpricption", Type: "Terrestrial", Mass: 1.0, Radius: 1.0, Distance: 100}
	exoplanets[exoplanet.ID] = exoplanet

	req, err := http.NewRequest("GET", "/exoplanets/12345/fuel?crew=5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets/{id}/fuel", CalculateFuelEstimation)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response models.FuelEstimation
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.True(t, response.FuelCost > 0)
}
