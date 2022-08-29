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

	var vc model.VaccineCertificate

	json.NewDecoder(r.Body).Decode(&vc)
	database.DB.Create(&vc)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&vc)
}

// GetVaccineCertificate gets an id-specified VaccineCertificate object from the database
func GetVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vc model.VaccineCertificate
	v := mux.Vars(r)

	tx := database.DB.First(&vc, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vc)
}

// UpdateVaccineCertificate updates an id-specified VaccineCertificate object
func UpdateVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vc model.VaccineCertificate
	v := mux.Vars(r)

	tx := database.DB.First(&vc, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&vc)
	database.DB.Save(&vc)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vc)
}

// DeleteVaccineCertificate deletes an id-specified VaccineCertificate object
func DeleteVaccineCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vc model.VaccineCertificate
	v := mux.Vars(r)

	tx := database.DB.First(&vc, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	database.DB.Delete(&vc)

	w.WriteHeader(http.StatusNoContent)
}

// GetVaccineCertificatees gets all the VaccineCertificate abjects from the database
func GetVaccineCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vcs []model.VaccineCertificate
	database.DB.Find(&vcs)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vcs)
}
