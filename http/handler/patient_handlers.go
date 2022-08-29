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

// CreatePatient creates a new Patient object to the database
func CreatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p model.Patient

	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&p)
}

// GetPatient gets an id-specified Patient object from the database
func GetPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p model.Patient
	v := mux.Vars(r)

	tx := database.DB.First(&p, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p)
}

// UpdatePatient updates an id-specified Patient object
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p model.Patient
	v := mux.Vars(r)

	tx := database.DB.First(&p, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p)
}

// DeletePatient deletes an id-specified Patient object
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p model.Patient
	v := mux.Vars(r)

	tx := database.DB.First(&p, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	database.DB.Delete(&p)

	w.WriteHeader(http.StatusNoContent)
}

// GetPatientes gets all the Patient abjects from the database
func GetPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ps []model.Patient
	database.DB.Find(&ps)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ps)
}
