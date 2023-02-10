package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/golang-jwt/jwt"
)

func main() {
	/*r, _ := repository.InitDB()
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	models.InitBook(r.DB)
	routes.RegisterBookstoreRoutes(api)

	log.Fatal(http.ListenAndServe("localhost:3000", router))*/

	n := "5aDN1JYbeufd8d_PfBaW02NpcnVTVJtl4jt0hRR6kR3ZXLm5bn2eWjXn5jfqwoV6PJTIVSnoY_UJIMDTDbYbMroLb5LR8x6UuLXU2X_uS714ComP1gOzi9MvM9A5OeuVqfDsV9ywm-CRvk3pYRcAhbXosXEyHoBGCeNno-1OAJlVyqchL8ehqaL8-Dp0PhxLdkQiDIUG_1cjDc4HGh78mrt0zK9jemLN0fgxKm1PqMGtrWDnH9P69QZfL5kW2-Q-lWVxvCZqVBlN6wt_72Lhd6hdBgJ4hR7JZ3uadxgrZ_nnT_8MDxkghYFCstdIsr-GPie3g9OdxGi6VHbF-EdmdQ"
	e := "AQAB"

	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJXZHdNd2t1ZUtmaXlkSmg2aVU4b20wNGVweEQ4cEpKNTlNVzVuLS1aSlFRIn0.eyJleHAiOjE2OTgxNDUwNTUsImlhdCI6MTY2NjYwOTA1NSwianRpIjoiZWM3ZTNjMjktMGIzOC00MWU3LThlMDctNGJmODNlMWU1MTYzIiwiaXNzIjoiaHR0cHM6Ly9zdGctYXV0aC5rYWxlcmlzLm5ldC9yZWFsbXMvc2hpcHhwcmVzcyIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJkNDI2ZmZmYy01YWYzLTQ4ZTYtOTE2OC02ZmRlNWZjZDNiNjAiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ0cmF4LWNsbS1zZXJ2aWNlIiwic2Vzc2lvbl9zdGF0ZSI6ImFkMzI3YWQxLWQ2YTctNGI2Ny1iOGUwLTU0M2U1ODA2YWJiMSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJkZWZhdWx0LXJvbGVzLXNoaXB4cHJlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiZW1haWwga2FsZXJpcyBwcm9maWxlIiwic2lkIjoiYWQzMjdhZDEtZDZhNy00YjY3LWI4ZTAtNTQzZTU4MDZhYmIxIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJzeC11c2VyLWlkIjoieG1sY2xlYW5oYXJiZXIiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJ4bWxjbGVhbmhhcmJlcl85NTJiNDZjYjI1Mjc5MzcyZGFmNDVjZGQ0YzZjY2FmYyIsImVtYWlsIjoiZGpheWF0aGlsYWthQGthbGVyaXMuY29tIn0.F5j2fkyKm9JNktGEiEOAzbbqz3t_F1VY5ONBeoXsZLz7fMWslPyStQJTrs66Cjs2VF48K8CSH9EP02IdvpK6F7Ng2CuOT00EqyVS7SG5FPFzQuYV9sWKOrW8_l7y6YcRxIxejpLtl3GBBAlnXefHk6aIkd-nUo5FUUvHrqmrRYy7I8bq9gBIRS9kBy3vjN0OI-ja-RVbUEmSrwxrLEA9tFpYK2_RXYQOLDez9Vlth6fHPp0fmwHdtQDjB9R4_0PjELRiNCk1iH24n2CbWSpkH1PbTJrFnE2QPstLW0jz8gC3tiFIA6x2LZzabvx-uQHq9sNStQMAc-rOi5dezQx-Wg"

	rsaPublicKey := &rsa.PublicKey{
		E: int(e[0]),
		N: new(big.Int).SetBytes([]byte(n)),
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

	tokena, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return &key, nil
	})

	if err != nil {
		fmt.Println("Error while parsing JWT2:", err)
		//continue
	}

	println(tokena.Valid)
}
