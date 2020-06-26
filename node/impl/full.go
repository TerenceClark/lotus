package impl

import (
	x_ctrl "github.com/filecoin-project/lotus/node/impl/x-ctrl"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/node/impl/client"
	"github.com/filecoin-project/lotus/node/impl/common"
	"github.com/filecoin-project/lotus/node/impl/full"
	"github.com/filecoin-project/lotus/node/impl/market"
	"github.com/filecoin-project/lotus/node/impl/paych"
)

var log = logging.Logger("node")

type FullNodeAPI struct {
	common.CommonAPI
	full.ChainAPI
	client.API
	full.MpoolAPI
	market.MarketAPI
	paych.PaychAPI
	full.StateAPI
	full.MsigAPI
	full.WalletAPI
	full.SyncAPI
	x_ctrl.XCtrl
}

var _ api.FullNode = &FullNodeAPI{}
