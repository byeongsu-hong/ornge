package main

import (
	"github.com/frostornge/ornge/app"
	abciclient "github.com/tendermint/tendermint/abci/client"
	tmconfig "github.com/tendermint/tendermint/config"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmnode "github.com/tendermint/tendermint/node"
	"os"
	"path"
)

func main() {
	logger, err := tmlog.NewDefaultLogger(
		tmlog.LogFormatPlain,
		tmlog.LogLevelInfo,
		true,
	)
	if err != nil {
		panic(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cfg := tmconfig.DefaultConfig()
	cfg.Mode = tmconfig.ModeValidator
	cfg.SetRoot(path.Join(home, ".tendermint"))

	creator := abciclient.NewLocalCreator(app.NewApp())
	node, err := tmnode.New(cfg, logger, creator, nil)
	if err != nil {
		panic(err)
	}

	if err := node.Start(); err != nil {
		panic(err)
	}

	node.Wait()
}
