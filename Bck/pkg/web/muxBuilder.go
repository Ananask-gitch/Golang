package web

import (
	"net/http"
)

func MakeMux() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
