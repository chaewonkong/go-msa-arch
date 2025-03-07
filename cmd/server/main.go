package main

import (
	"github.com/chaewonkong/go-msa-arch/internal/app"
	"github.com/chaewonkong/go-msa-arch/internal/app/server"
)

func main() {
	app.RunApplication("server", &server.Factory{})
}
