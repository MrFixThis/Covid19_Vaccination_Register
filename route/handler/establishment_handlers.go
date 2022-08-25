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

// CreateEstablishment creates a new Establishment object to the database
func CreateEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment

	json.NewDecoder(r.Body).Decode(&e)
	database.DB.Create(&e)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&e)
}

// GetEstablishment gets an id-specified Establishment object from the database
func GetEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment
	v := mux.Vars(r)

	tx := database.DB.First(&e, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&e)
}

// UpdateEstablishment updates an id-specified Establishment object
func UpdateEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment
	v := mux.Vars(r)

	tx := database.DB.First(&e, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&e)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&e)
}

// DeleteEstablishment deletes an id-specified Establishment object
func DeleteEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment
	v := mux.Vars(r)

	tx := database.DB.First(&e, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusOK
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	database.DB.Delete(&e)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("n deleted successfully")
}

// GetEstablishmentes gets all the Establishment abjects from the database
func GetEstablishments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var es []model.Establishment
	database.DB.Find(&es)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&es)
}
