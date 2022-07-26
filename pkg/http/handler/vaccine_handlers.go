package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Covid19_Vaccination_Register/pkg/model"
	"github.com/Covid19_Vaccination_Register/pkg/storage"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateVaccine creates a new Vaccine object to the database
func CreateVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vn model.Vaccine

	json.NewDecoder(r.Body).Decode(&vn)
	storage.DB.Create(&vn)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&vn)
}

// GetVaccine gets an id-specified Vaccine object from the database
func GetVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vn model.Vaccine
	v := mux.Vars(r)

	tx := storage.DB.Preload(clause.Associations).First(&vn, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vn)
}

// UpdateVaccine updates an id-specified Vaccine object
func UpdateVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vn model.Vaccine
	v := mux.Vars(r)

	tx := storage.DB.Preload(clause.Associations).First(&vn, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&vn)
	storage.DB.Save(&vn)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vn)
}

// DeleteVaccine deletes an id-specified Vaccine object
func DeleteVaccine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vn model.Vaccine
	v := mux.Vars(r)

	tx := storage.DB.First(&vn, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	storage.DB.Delete(&vn)

	w.WriteHeader(http.StatusNoContent)
}

// GetVaccinees gets all the Vaccine abjects from the database
func GetVaccines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vns []model.Vaccine
	storage.DB.Preload(clause.Associations).First(&vns)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&vns)
}
