package api

import "context"

type XCtrl interface {
	SetValidHosts(ctx context.Context, hosts []string) error
	HostVerify(ctx context.Context, host string) error
}
