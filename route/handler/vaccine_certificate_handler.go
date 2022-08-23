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

// CreateVaccineCertificate creates a new VaccineCertificate object to the database
func CreateVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccineCertificate model.VaccineCertificate

	json.NewDecoder(r.Body).Decode(&vaccineCertificate)
	database.DB.Create(&vaccineCertificate)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&vaccineCertificate)
}

// GetVaccineCertificate gets an id-specified VaccineCertificate object from the database
func GetVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccineCertificate model.VaccineCertificate
	v := mux.Vars(r)

	tx := database.DB.First(&vaccineCertificate, v["id"])
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			http.Error(w, tx.Error.Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccineCertificate)
}

// UpdateVaccineCertificate updates an id-specified VaccineCertificate object
func UpdateVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccineCertificate model.VaccineCertificate
	v := mux.Vars(r)

	tx := database.DB.First(&vaccineCertificate, v["id"])
	if tx.Error != nil {
		status := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), status)
		return
	}
	json.NewDecoder(r.Body).Decode(&vaccineCertificate)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccineCertificate)
}

// DeleteVaccineCertificate deletes an id-specified VaccineCertificate object
func DeleteVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccineCertificate model.VaccineCertificate
	v := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&vaccineCertificate)
	err := database.DB.Delete(&vaccineCertificate, v["id"])
	if err != nil {
		http.Error(w, err.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("vaccineCertificate deleted successfully")
}

// GetVaccineCertificatees gets all the VaccineCertificate abjects from the database
func GetVaccineCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vaccineCertificates []model.VaccineCertificate
	database.DB.Find(&vaccineCertificates)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vaccineCertificates)
}
