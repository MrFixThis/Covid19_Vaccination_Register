package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Covid19_Vaccination_Register/pkg/storage"
	"github.com/golang-jwt/jwt/v4"
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
func AthenticateAdministrator(f http.HandlerFunc) http.HandlerFunc {
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
		f(w, r)
	})
}

func IsAuthorized(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rt := r.Header["Token"]

		if rt == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		m := []string{"HS256"}
		t, err := jwt.Parse(rt[0], func(t *jwt.Token) (any, error) {
			return os.Getenv("SECRET"), nil
		}, jwt.WithValidMethods(m))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !t.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		f(w, r)
	})
}

// GenerateJWT generates a new JWT with a specified subject and expiration
// time
func GenerateJWT(sub string, exp int) (string, error) {
	tc := jwt.StandardClaims{
		Subject:   sub,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)

	return t.SignedString(os.Getenv("SECRET"))
}
