package apistruct

import (
	"context"
)

type XCtrlStruct struct {
	Internal struct {
		SetValidHosts func(ctx context.Context, hosts []string) error `perm:"admin"`
		HostVerify func(ctx context.Context, host string) error `perm:"read"`
	}
}

func (s *XCtrlStruct) SetValidHosts(ctx context.Context, hosts []string) error {
	//drpclog.L.Debug("call set valid host", zap.Any("internal", reflect.TypeOf(s.Internal)))
	return s.Internal.SetValidHosts(ctx, hosts)
}
func (s *XCtrlStruct) HostVerify(ctx context.Context, host string) error {
	return s.Internal.HostVerify(ctx, host)
}