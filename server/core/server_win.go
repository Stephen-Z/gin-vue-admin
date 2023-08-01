//go:build windows
// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    6000 * time.Second,
		WriteTimeout:   6000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
