package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Covid19_Vaccination_Register/pkg/model"
	"github.com/Covid19_Vaccination_Register/pkg/storage"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func AthenticateAdministrator(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ar model.Administrator
		var ac model.Administrator

		json.NewDecoder(r.Body).Decode(&ac)
		storage.DB.Where("admin_name = ?", ac.AdminName).
			Select("password").First(&ar)

		err := bcrypt.CompareHashAndPassword([]byte(ar.Password),
			[]byte(ac.Password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		h(w, r)
	})
}

// func ValidateToken(h http.HandlerFunc) http.HandlerFunc {
//
// }

// getSecret retrives the secret from the .env file in the project's root
func getSecret() string {
	 godotenv.Load(".env")
	 return os.Getenv("SECRET")
}
