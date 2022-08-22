package route

import (
	"net/http"

	"github.com/Covid19_Vaccination_Register/route/handler"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	p := r.PathPrefix("api/v1").Subrouter()

	// Address routes
	s := p.PathPrefix("/addresss").Subrouter()
	s.HandleFunc("", handler.CreateAddress).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetAddress).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdateAddress).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeleteAddress).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetAddresses).Methods(http.MethodGet)

	// Vaccine routes
	s = p.PathPrefix("/vaccines").Subrouter()
	s.HandleFunc("", handler.CreateVaccine).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetVaccine).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccine).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccine).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetVaccines).Methods(http.MethodGet)

	// Patient routes
	s = p.PathPrefix("/patients").Subrouter()
	s.HandleFunc("", handler.CreatePatient).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetPatient).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdatePatient).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeletePatient).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetPatients).Methods(http.MethodGet)

	// Nurse routes
	s = p.PathPrefix("/nurses").Subrouter()
	s.HandleFunc("", handler.CreateNurse).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetNurse).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdateNurse).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeleteNurse).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetNurses).Methods(http.MethodGet)

	// Establishment routes
	s = p.PathPrefix("/establishments").Subrouter()
	s.HandleFunc("", handler.CreateEstablishment).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetEstablishment).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdateEstablishment).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeleteEstablishment).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetEstablishments).Methods(http.MethodGet)

	// Vaccine certificate routes
	s = p.PathPrefix("/vaccine_certificate").Subrouter()
	s.HandleFunc("", handler.CreateVaccineCertificate).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", handler.GetVaccineCertificate).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccineCertificate).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccineCertificate).Methods(http.MethodDelete)
	s.HandleFunc("", handler.GetVaccineCertificates).Methods(http.MethodGet)

	//TODO: See if it's necesary to implement the same with the administrators

	return r
}
