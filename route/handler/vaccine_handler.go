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

// CreateVaccine creates a new Vaccine object to the database
func CreateVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccine model.Vaccine

	json.NewDecoder(r.Body).Decode(&vaccine)
	database.DB.Create(&vaccine)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&vaccine)
}

// GetVaccine gets an id-specified Vaccine object from the database
func GetVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccine model.Vaccine
	v := mux.Vars(r)

	tx := database.DB.First(&vaccine, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccine)
}

// UpdateVaccine updates an id-specified Vaccine object
func UpdateVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccine model.Vaccine
	v := mux.Vars(r)

	tx := database.DB.First(&vaccine, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&vaccine)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccine)
}

// DeleteVaccine deletes an id-specified Vaccine object
func DeleteVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccine model.Vaccine
	v := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&vaccine)
	err := database.DB.Delete(&vaccine, v["id"])
	if err != nil {
		http.Error(w, err.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("vaccine deleted successfully")
}

// GetVaccinees gets all the Vaccine abjects from the database
func GetVaccines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccines []model.Vaccine
	database.DB.Find(&vaccines)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccines)
}
