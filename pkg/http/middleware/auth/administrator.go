package auth

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Covid19_Vaccination_Register/pkg/storage"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


// adminCredentials struct is used by the authentication middleware
// to store the administrator's credentials coming in the requet's body
type adminCredentials struct {
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
}

// AthenticateAdministrator middleware checks the administrator's credetials,
// validating and authenticating if it is registered in the system. In response,
// an authorization token (JWT) is returned to the client.
func AthenticateAdministrator(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ar adminCredentials
		var ac adminCredentials

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

// func generateToken(h http.HandlerFunc) http.HandlerFunc {
//
// }

// getSecret retrives the secret from the .env file in the project's root
func getSecret() string {
	 godotenv.Load(".env")
	 return os.Getenv("SECRET")
}
