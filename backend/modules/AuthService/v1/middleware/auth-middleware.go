package middleware

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"

	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

//func Validate() error {
//	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options {
//		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//			// Verify 'aud' claim
//			aud := "https://alfred-sh.us.auth0.com/api/v2/"
//			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
//			if !checkAud {
//				return token, errors.New("Invalid audience.")
//			}
//			// Verify 'iss' claim
//			iss := "alfred-sh.us.auth0.com"
//			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
//			if !checkIss {
//				return token, errors.New("Invalid issuer.")
//			}
//
//			cert, err := getPemCert(token)
//			if err != nil {
//				panic(err.Error())
//			}
//
//			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
//			return result, nil
//		},
//		SigningMethod: jwt.SigningMethodRS256,
//	})
//	return nil
//}

func getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://alfred-sh.us.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
