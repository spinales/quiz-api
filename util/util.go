package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// respondwithError return error message
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
