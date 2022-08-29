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

	var ad model.Address

	json.NewDecoder(r.Body).Decode(&ad)
	database.DB.Create(&ad)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ad)
}

// GetAddress gets an id-specified Address object from the database
func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ad model.Address
	v := mux.Vars(r)

	tx := database.DB.First(&ad, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ad)
}

// UpdateAddress updates an id-specified Address object
func UpdateAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ad model.Address
	v := mux.Vars(r)

	tx := database.DB.First(&ad, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&ad)
	database.DB.Save(&ad)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ad)
}

// DeleteAddress deletes an id-specified Address object
func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ad model.Address
	v := mux.Vars(r)

	tx := database.DB.First(&ad, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	database.DB.Delete(&ad)

	w.WriteHeader(http.StatusNoContent)
}

// GetAddresses gets all the Address abjects from the database
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ads []model.Address
	database.DB.Find(&ads)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ads)
}
