package route

import (
	"net/http"

	"github.com/Covid19_Vaccination_Register/route/handler"
	"github.com/gorilla/mux"
)

// InitializeRoutes initializes the server's routes to handle the http requests
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	// Address routes
	as := s.PathPrefix("/addresses").Subrouter()
	as.HandleFunc("", handler.CreateAddress).Methods(http.MethodPost)
	as.HandleFunc("/{id:[0-9]+}", handler.GetAddress).Methods(http.MethodGet)
	as.HandleFunc("/{id:[0-9]+}", handler.UpdateAddress).Methods(http.MethodPut)
	as.HandleFunc("/{id:[0-9]+}", handler.DeleteAddress).Methods(http.MethodDelete)
	as.HandleFunc("", handler.GetAddresses).Methods(http.MethodGet)

	// Vaccine routes
	vi := s.PathPrefix("/vaccines").Subrouter()
	vi.HandleFunc("", handler.CreateVaccine).Methods(http.MethodPost)
	vi.HandleFunc("/{id:[0-9]+}", handler.GetVaccine).Methods(http.MethodGet)
	vi.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccine).Methods(http.MethodPut)
	vi.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccine).Methods(http.MethodDelete)
	vi.HandleFunc("", handler.GetVaccines).Methods(http.MethodGet)

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
	vc := s.PathPrefix("/vaccine_certificate").Subrouter()
	vc.HandleFunc("", handler.CreateVaccineCertificate).Methods(http.MethodPost)
	vc.HandleFunc("/{id:[0-9]+}", handler.GetVaccineCertificate).Methods(http.MethodGet)
	vc.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccineCertificate).Methods(http.MethodPut)
	vc.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccineCertificate).Methods(http.MethodDelete)
	vc.HandleFunc("", handler.GetVaccineCertificates).Methods(http.MethodGet)

	an := s.PathPrefix("/administrators").Subrouter()
	an.HandleFunc("", handler.CreateAdministrator).Methods(http.MethodPost)
	an.HandleFunc("/{id:[0-9]+}", handler.GetAdministrator).Methods(http.MethodGet)
	an.HandleFunc("/{id:[0-9]+}", handler.UpdateAdministrator).Methods(http.MethodPut)
	an.HandleFunc("/{id:[0-9]+}", handler.DeleteAdministrator).Methods(http.MethodDelete)
	an.HandleFunc("", handler.GetAdministrators).Methods(http.MethodGet)

	return r
}
