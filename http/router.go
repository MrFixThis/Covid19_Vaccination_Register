package route

import (
	"net/http"
	"os"

	"github.com/Covid19_Vaccination_Register/http/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InitializeRoutes initializes the server's routes to handle the http requests
func InitializeRoutes() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	// Address routes
	ad := s.PathPrefix("/addresses").Subrouter()
	ad.HandleFunc("", handler.CreateAddress).Methods(http.MethodPost)
	ad.HandleFunc("/{id:[0-9]+}", handler.GetAddress).Methods(http.MethodGet)
	ad.HandleFunc("/{id:[0-9]+}", handler.UpdateAddress).Methods(http.MethodPut)
	ad.HandleFunc("/{id:[0-9]+}", handler.DeleteAddress).Methods(http.MethodDelete)
	ad.HandleFunc("", handler.GetAddresses).Methods(http.MethodGet)

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
	vc := s.PathPrefix("/vaccine_certificates").Subrouter()
	vc.HandleFunc("", handler.CreateVaccineCertificate).Methods(http.MethodPost)
	vc.HandleFunc("/{id:[0-9]+}", handler.GetVaccineCertificate).Methods(http.MethodGet)
	vc.HandleFunc("/{id:[0-9]+}", handler.UpdateVaccineCertificate).Methods(http.MethodPut)
	vc.HandleFunc("/{id:[0-9]+}", handler.DeleteVaccineCertificate).Methods(http.MethodDelete)
	vc.HandleFunc("", handler.GetVaccineCertificates).Methods(http.MethodGet)

	ar := s.PathPrefix("/administrators").Subrouter()
	ar.HandleFunc("", handler.CreateAdministrator).Methods(http.MethodPost)
	ar.HandleFunc("/{id:[0-9]+}", handler.GetAdministrator).Methods(http.MethodGet)
	ar.HandleFunc("/{id:[0-9]+}", handler.UpdateAdministrator).Methods(http.MethodPut)
	ar.HandleFunc("/{id:[0-9]+}", handler.DeleteAdministrator).Methods(http.MethodDelete)
	ar.HandleFunc("", handler.GetAdministrators).Methods(http.MethodGet)

	return handlers.LoggingHandler(os.Stdout, r)
}
