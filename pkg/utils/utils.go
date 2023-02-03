package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
		if err := f(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				WriteJSON(w, nil, e.Status, e)
				return
			}
			WriteJSON(w, nil, http.StatusInternalServerError, ApiError{Err: "Internal server", Status: 500})
		}

	}
}

func WriteJSON(w http.ResponseWriter, res []byte, status int, v any) error {
	w.WriteHeader(status)
	_, err := w.Write(res)
	w.Header().Set("Content-type", "application/json")
	return err
}
