package main

import (
	"github.com/Borislavv/go-seo/cmd/app/seo"
	"github.com/Borislavv/go-seo/pkg/shared/config"
)

func init() {
	config.Load()
}

func main() {
	new(seo.App).Run()
}
