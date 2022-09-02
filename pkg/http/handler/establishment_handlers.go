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

// CreateEstablishment creates a new Establishment object to the database
func CreateEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment

	json.NewDecoder(r.Body).Decode(&e)
	storage.DB.Create(&e)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&e)
}

// GetEstablishment gets an id-specified Establishment object from the database
func GetEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment
	v := mux.Vars(r)

	tx := storage.DB.Preload(clause.Associations).First(&e, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
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

	tx := storage.DB.Preload(clause.Associations).First(&e, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&e)
	storage.DB.Save(&e)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&e)
}

// DeleteEstablishment deletes an id-specified Establishment object
func DeleteEstablishment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var e model.Establishment
	v := mux.Vars(r)

	tx := storage.DB.First(&e, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	storage.DB.Delete(&e)

	w.WriteHeader(http.StatusNoContent)
}

// GetEstablishmentes gets all the Establishment abjects from the database
func GetEstablishments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var es []model.Establishment
	storage.DB.Preload(clause.Associations).Find(&es)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&es)
}
