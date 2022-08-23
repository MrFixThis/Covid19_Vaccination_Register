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

	var nurse model.Nurse

	json.NewDecoder(r.Body).Decode(&nurse)
	database.DB.Create(&nurse)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&nurse)
}

// GetNurse gets an id-specified Nurse object from the database
func GetNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nurse model.Nurse
	v := mux.Vars(r)

	tx := database.DB.First(&nurse, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&nurse)
}

// UpdateNurse updates an id-specified Nurse object
func UpdateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nurse model.Nurse
	v := mux.Vars(r)

	tx := database.DB.First(&nurse, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&nurse)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&nurse)
}

// DeleteNurse deletes an id-specified Nurse object
func DeleteNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nurse model.Nurse
	v := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&nurse)
	err := database.DB.Delete(&nurse, v["id"])
	if err != nil {
		http.Error(w, err.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("nurse deleted successfully")
}

// GetNursees gets all the Nurse abjects from the database
func GetNurses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nurses []model.Nurse
	database.DB.Find(&nurses)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&nurses)
}
