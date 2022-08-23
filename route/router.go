package route

import (
	"net/http"

	"github.com/Covid19_Vaccination_Register/route/handler"
	"github.com/gorilla/mux"
)

// InitializeRoutes initializes the server's routes to handle the http requests
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("api/v1").Subrouter()

	// Address routes
	a := s.PathPrefix("/addresses").Subrouter()
	a.HandleFunc("", handler.CreateAddress).Methods(http.MethodPost)
	a.HandleFunc("/addresses/{id:[0-9]+}", handler.GetAddress).Methods(http.MethodGet)
	a.HandleFunc("/addresses/{id:[0-9]+}", handler.UpdateAddress).Methods(http.MethodPut)
	a.HandleFunc("/addresses/{id:[0-9]+}", handler.DeleteAddress).Methods(http.MethodDelete)
	a.HandleFunc("/addresses/", handler.GetAddresses).Methods(http.MethodGet)

	// Vaccine routes
	v := s.PathPrefix("/vaccines").Subrouter()
	v.HandleFunc("", handler.CreateVaccine).Methods(http.MethodPost)
	v.HandleFunc("/{id:[0-9]+}", handler.GetVaccine).Methods(http.MethodGet)
	v.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccine).Methods(http.MethodPut)
	v.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccine).Methods(http.MethodDelete)
	v.HandleFunc("", handler.GetVaccines).Methods(http.MethodGet)

	// Patient routes
	p := s.PathPrefix("/patients").Subrouter()
	p.HandleFunc("", handler.CreatePatient).Methods(http.MethodPost)
	p.HandleFunc("/{id:[0-9]+}", handler.GetPatient).Methods(http.MethodGet)
	p.HandleFunc("/{id:[0-9]+}", handler.UpdatePatient).Methods(http.MethodPut)
	p.HandleFunc("/{id:[0-9]+}", handler.DeletePatient).Methods(http.MethodDelete)
	p.HandleFunc("", handler.GetPatients).Methods(http.MethodGet)

	// Nurse routes
	n := s.PathPrefix("/nurses").Subrouter()
	n.HandleFunc("", handler.CreateNurse).Methods(http.MethodPost)
	n.HandleFunc("/{id:[0-9]+}", handler.GetNurse).Methods(http.MethodGet)
	n.HandleFunc("/{id:[0-9]+}", handler.UpdateNurse).Methods(http.MethodPut)
	n.HandleFunc("/{id:[0-9]+}", handler.DeleteNurse).Methods(http.MethodDelete)
	n.HandleFunc("", handler.GetNurses).Methods(http.MethodGet)

	// Establishment routes
	e := s.PathPrefix("/establishments").Subrouter()
	e.HandleFunc("", handler.CreateEstablishment).Methods(http.MethodPost)
	e.HandleFunc("/{id:[0-9]+}", handler.GetEstablishment).Methods(http.MethodGet)
	e.HandleFunc("/{id:[0-9]+}", handler.UpdateEstablishment).Methods(http.MethodPut)
	e.HandleFunc("/{id:[0-9]+}", handler.DeleteEstablishment).Methods(http.MethodDelete)
	e.HandleFunc("", handler.GetEstablishments).Methods(http.MethodGet)

	// Vaccine certificate routes
	c := s.PathPrefix("/vaccine_certificate").Subrouter()
	c.HandleFunc("", handler.CreateVaccineCertificate).Methods(http.MethodPost)
	c.HandleFunc("/{id:[0-9]+}", handler.GetVaccineCertificate).Methods(http.MethodGet)
	c.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccineCertificate).Methods(http.MethodPut)
	c.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccineCertificate).Methods(http.MethodDelete)
	c.HandleFunc("", handler.GetVaccineCertificates).Methods(http.MethodGet)

	//TODO: See if it's necesary to implement the same with the administrators

	return r
}
