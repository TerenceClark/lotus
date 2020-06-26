package x_ctrl

import (
	"context"
	"github.com/filecoin-project/lotus/tools/dlog/drpclog"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type XCtrl struct {
	fx.In
}

func (a *XCtrl) SetValidHosts(ctx context.Context, hosts []string) error {
	drpclog.L.Debug("SetValidHosts", zap.Strings("hosts", hosts))
	return nil
}