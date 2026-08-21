package adapter

import "net/http"

func StatusCode(conclusion string) int {
	if conclusion == "pass" {
		return http.StatusOK
	}
	return http.StatusUnprocessableEntity
}
