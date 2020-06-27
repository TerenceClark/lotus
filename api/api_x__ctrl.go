package api

import "context"

type XCtrl interface {
	SetValidHosts(ctx context.Context, hosts []string) error
}
