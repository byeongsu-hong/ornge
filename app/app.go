package app

import (
	"github.com/tendermint/tendermint/abci/example/kvstore"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

var _ abcitypes.Application = App{}

type App struct {
	baseApp abcitypes.Application
}

func NewApp() abcitypes.Application {
	return App{
		baseApp: kvstore.NewApplication(),
	}
}

func (a App) Info(info abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return a.baseApp.Info(info)
}

func (a App) Query(query abcitypes.RequestQuery) abcitypes.ResponseQuery {
	return a.baseApp.Query(query)
}

func (a App) CheckTx(tx abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	return a.baseApp.CheckTx(tx)
}

func (a App) InitChain(chain abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return a.baseApp.InitChain(chain)
}

func (a App) BeginBlock(block abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return a.baseApp.BeginBlock(block)
}

func (a App) DeliverTx(tx abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	return a.baseApp.DeliverTx(tx)
}

func (a App) EndBlock(block abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return a.baseApp.EndBlock(block)
}

func (a App) Commit() abcitypes.ResponseCommit {
	return a.baseApp.Commit()
}

func (a App) ListSnapshots(snapshots abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return a.baseApp.ListSnapshots(snapshots)
}

func (a App) OfferSnapshot(snapshot abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return a.baseApp.OfferSnapshot(snapshot)
}

func (a App) LoadSnapshotChunk(chunk abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return a.baseApp.LoadSnapshotChunk(chunk)
}

func (a App) ApplySnapshotChunk(chunk abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return a.baseApp.ApplySnapshotChunk(chunk)
}
