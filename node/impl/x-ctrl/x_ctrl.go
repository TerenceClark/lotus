package x_ctrl

import (
	"context"
	"errors"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/tools/dlog/drpclog"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type XCtrl struct {
	fx.In

	*dtypes.RPCHostVerifier
}

func (a *XCtrl) SetValidHosts(ctx context.Context, hosts []string) error {
	drpclog.L.Debug("SetValidHosts", zap.Strings("hosts", hosts))
	a.VHostsLock.Lock()
	defer a.VHostsLock.Unlock()
	// todo write to file
	a.ValidHosts = append(hosts, "127.0.0.1")
	return nil
}

func (a *XCtrl) HostVerify(ctx context.Context, host string) error {
	drpclog.L.Debug("do HostVerify", zap.String("host", host), zap.Strings("v hosts", a.ValidHosts))
	a.VHostsLock.RLock()
	defer a.VHostsLock.RUnlock()

	for _, h := range a.ValidHosts {
		if host == h {
			return nil
		}
	}
	return errors.New("invalid client")
}
