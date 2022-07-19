package main

import (
	"fmt"
	"github.com/frostornge/ornge/app"
	abciserver "github.com/tendermint/tendermint/abci/server"
)

func main() {
	server, err := abciserver.NewServer(
		fmt.Sprintf("localhost:%d", 26658),
		"grpc",
		app.NewApp(),
	)
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}

	server.Wait()
}
