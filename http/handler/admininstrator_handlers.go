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

// CreateAdministrator creates a new Administrator object to the database
func CreateAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ar model.Administrator

	json.NewDecoder(r.Body).Decode(&ar)
	database.DB.Create(&ar); ar.Password = ""

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ar)
}

// GetAdministrator gets an id-specified Administrator object from the database
func GetAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ar model.Administrator
	v := mux.Vars(r)

	tx := database.DB.Omit("password").First(&ar, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ar)
}

// UpdateAdministrator updates an id-specified Administrator object
func UpdateAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ar model.Administrator
	v := mux.Vars(r)

	tx := database.DB.First(&ar, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	json.NewDecoder(r.Body).Decode(&ar)
	database.DB.Save(&ar); ar.Password = ""

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ar)
}

// DeleteAdministrator deletes an id-specified Administrator object
func DeleteAdministrator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ar model.Administrator
	v := mux.Vars(r)

	tx := database.DB.First(&ar, v["id"])
	if tx.Error != nil {
		s := http.StatusInternalServerError
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			s = http.StatusNotFound
		}
		http.Error(w, tx.Error.Error(), s)
		return
	}
	database.DB.Delete(&ar)

	w.WriteHeader(http.StatusNoContent)
}

// GetAdministratores gets all the Administrator abjects from the database
func GetAdministrators(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ars []model.Administrator
	database.DB.Omit("password").Find(&ars)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ars)
}
