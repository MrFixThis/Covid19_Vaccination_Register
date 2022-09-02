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

// CreateNurse creates a new Nurse object to the database
func CreateNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse

	json.NewDecoder(r.Body).Decode(&n)
	storage.DB.Create(&n)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&n)
}

// GetNurse gets an id-specified Nurse object from the database
func GetNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse
	v := mux.Vars(r)

	tx := storage.DB.Preload(clause.Associations).First(&n, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
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

	tx := storage.DB.Preload(clause.Associations).First(&n, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&n)
	storage.DB.Save(&n)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&n)
}

// DeleteNurse deletes an id-specified Nurse object
func DeleteNurse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var n model.Nurse
	v := mux.Vars(r)

	tx := storage.DB.First(&n, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	storage.DB.Delete(&n)

	w.WriteHeader(http.StatusNoContent)
}

// GetNursees gets all the Nurse abjects from the database
func GetNurses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ns []model.Nurse
	storage.DB.Preload(clause.Associations).Find(&ns)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ns)
}
