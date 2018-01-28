package controllers

import (
	"net/http"
)

func ObtainPublicFilesHandler(publicDir string) http.Handler  {
	return http.FileServer(http.Dir(publicDir))
}

