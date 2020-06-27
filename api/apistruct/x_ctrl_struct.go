package apistruct

import (
	"context"
)

type XCtrlStruct struct {
	Internal struct {
		SetValidHosts func(ctx context.Context, hosts []string) error `perm:"admin"`
	}
}

func (s *XCtrlStruct) SetValidHosts(ctx context.Context, hosts []string) error {
	//drpclog.L.Debug("call set valid host", zap.Any("internal", reflect.TypeOf(s.Internal)))
	return s.Internal.SetValidHosts(ctx, hosts)
}
