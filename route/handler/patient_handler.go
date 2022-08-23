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

	var patient model.Patient

	json.NewDecoder(r.Body).Decode(&patient)
	database.DB.Create(&patient)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&patient)
}

// GetPatient gets an id-specified Patient object from the database
func GetPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patient model.Patient
	v := mux.Vars(r)

	tx := database.DB.First(&patient, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&patient)
}

// UpdatePatient updates an id-specified Patient object
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patient model.Patient
	v := mux.Vars(r)

	tx := database.DB.First(&patient, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&patient)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&patient)
}

// DeletePatient deletes an id-specified Patient object
func DeletePatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patient model.Patient
	v := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&patient)
	err := database.DB.Delete(&patient, v["id"])
	if err != nil {
		http.Error(w, err.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("patient deleted successfully")
}

// GetPatientes gets all the Patient abjects from the database
func GetPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var patients []model.Patient
	database.DB.Find(&patients)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&patients)
}
