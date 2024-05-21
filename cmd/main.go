package main

import (
	"context"
	"github.com/Borislavv/go-seo/pkg/shared/logger"
	"github.com/Borislavv/go-seo/pkg/shared/server"
	"github.com/Borislavv/go-seo/pkg/shared/shutdown"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())

	lgr, closeFunc := logger.NewLogger()
	defer closeFunc()

	server.
		NewHTTP(lgr).
		ListenAndServe(ctx, wg)

	shutdown.
		NewGraceful(cancel).
		ListenAndCancel()

	wg.Wait()
}
