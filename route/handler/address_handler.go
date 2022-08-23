package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Covid19_Vaccination_Register/database"
	"github.com/Covid19_Vaccination_Register/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateAddress creates a new Address object to the database
func CreateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var address model.Address

	json.NewDecoder(r.Body).Decode(&address)
	database.DB.Create(&address)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&address)
}

// GetAddress gets an id-specified Address object from the database
func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var address model.Address
	v := mux.Vars(r)

	tx := database.DB.First(&address, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&address)
}

// UpdateAddress updates an id-specified Address object
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var address model.Address
	v := mux.Vars(r)

	tx := database.DB.First(&address, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&address)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&address)
}

// DeleteAddress deletes an id-specified Address object
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var address model.Address
	v := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&address)
	err := database.DB.Delete(&address, v["id"])
	if err != nil {
		http.Error(w, err.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("address deleted successfully")
}

// GetAddresses gets all the Address abjects from the database
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var addresses []model.Address
	database.DB.Find(&addresses)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&addresses)
}
