package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//type Response struct {
//	Message string `json:"message"`
//}
//
//type Jwks struct {
//	Keys []JSONWebKeys `json:"keys"`
//}
//
//type JSONWebKeys struct {
//	Kty string   `json:"kty"`
//	Kid string   `json:"kid"`
//	Use string   `json:"use"`
//	N   string   `json:"n"`
//	E   string   `json:"e"`
//	X5c []string `json:"x5c"`
//}

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

//func getPemCert(token *jwt.Token) (string, error) {
//	cert := ""
//	resp, err := http.Get("https://alfred-sh.us.auth0.com/.well-known/jwks.json")
//
//	if err != nil {
//		return cert, err
//	}
//	defer resp.Body.Close()
//
//	var jwks = Jwks{}
//	err = json.NewDecoder(resp.Body).Decode(&jwks)
//
//	if err != nil {
//		return cert, err
//	}
//
//	for k, _ := range jwks.Keys {
//		if token.Header["kid"] == jwks.Keys[k].Kid {
//			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
//		}
//	}
//
//	if cert == "" {
//		err := errors.New("Unable to find appropriate key.")
//		return cert, err
//	}
//
//	return cert, nil
//}

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (manager *JWTManager) Generate() (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		Username: "",
		Role:     "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
