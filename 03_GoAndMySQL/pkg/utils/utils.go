package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)


func ParseBody(r *http.Request, x interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Błąd przy odczycie ciała żądania:", err)
		return err // Zwracamy błąd
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Println("Błąd przy parsowaniu JSON:", err)
		return err // Zwracamy błąd
	}

	return nil
}
