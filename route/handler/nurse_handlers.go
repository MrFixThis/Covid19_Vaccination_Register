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

// CreateNurse creates a new Nurse object to the database
func CreateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse

	json.NewDecoder(r.Body).Decode(&n)
	database.DB.Create(&n)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&n)
}

// GetNurse gets an id-specified Nurse object from the database
func GetNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse
	v := mux.Vars(r)

	tx := database.DB.First(&n, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&n)
}

// UpdateNurse updates an id-specified Nurse object
func UpdateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse
	v := mux.Vars(r)

	tx := database.DB.First(&n, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&n)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&n)
}

// DeleteNurse deletes an id-specified Nurse object
func DeleteNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse
	v := mux.Vars(r)

	tx := database.DB.First(&n, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusOK
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	database.DB.Delete(&n)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("n deleted successfully")
}

// GetNursees gets all the Nurse abjects from the database
func GetNurses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ns []model.Nurse
	database.DB.Find(&ns)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ns)
}
