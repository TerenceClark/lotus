package drpclog

import (
	"github.com/filecoin-project/lotus/tools/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("rpc")
}