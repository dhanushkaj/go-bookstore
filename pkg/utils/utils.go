package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func ParseBody(r *http.Request, X interface{}) {

	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}

}

type apiFunc func(http.ResponseWriter, *http.Request) error

func MakeHttpHandler(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//	if err := validateJWT(w, r); err != nil {
		//		if e, ok := err.(ApiError); ok {
		//			WriteJSON(w, []byte(e.Err), e.Status, e)
		//			return
		//		}
		//	} else {
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				WriteJSON(w, []byte(e.Err), e.Status, e)
				return
			}
			WriteJSON(w, nil, http.StatusInternalServerError, ApiError{Err: "Internal server", Status: 500})
		}
		//}
	}
}

func WriteJSON(w http.ResponseWriter, res []byte, status int, v any) error {
	w.WriteHeader(status)
	_, err := w.Write(res)
	w.Header().Set("Content-type", "application/json")
	return err
}

type KeycloakKeys struct {
	Keys []struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
		Kty string `json:"kty"`
	} `json:"keys"`
}

func validateJWT(w http.ResponseWriter, r *http.Request) error {

	// Get the JWT from the Authorization header
	jwtToken := r.Header.Get("Authorization")

	if jwtToken == "" {
		//WriteJSON(w, []byte("Authorization header is missing"), http.StatusBadRequest, nil)
		return ApiError{Err: "Authorization header is missing", Status: http.StatusBadRequest}
	}
	resp, _ := http.Get("https://stg-auth.kaleris.net/realms/shipxpress/protocol/openid-connect/certs")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Parse the JSON data into a Go data structure
	var keys KeycloakKeys
	err = json.Unmarshal(body, &keys)
	if err != nil {
		log.Fatalf("Failed to parse Keycloak keys: %v", err)
	}
	for _, key := range keys.Keys {

		rsaPublicKey := &rsa.PublicKey{
			E: int(key.E[0]),
			N: new(big.Int).SetBytes([]byte(key.N)),
		}

		publicKeyBytes := x509.MarshalPKCS1PublicKey(rsaPublicKey)

		publicKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicKeyBytes,
		})

		//	block, _ := pem.Decode(publicKeyPEM)

		//publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

		key, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyPEM)

		fmt.Println(string(publicKeyPEM))

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return &key, nil
		})

		if err != nil {
			return ApiError{Err: "Error while parsing JWT", Status: http.StatusUnauthorized}
		}

		if !token.Valid {
			return ApiError{Err: "Invalid Token", Status: http.StatusUnauthorized}
		}

	}
	return nil
}
