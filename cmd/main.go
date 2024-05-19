package main

import (
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/Borislavv/go-seo/pkg/shared/server"
)

func main() {
	lgr, closeFunc := logger.NewLogger()
	defer closeFunc()

	lgr.Info("Hello world from logger by interface")

	server.
		NewHTTP(lgr).
		ListenAndServe()
}
