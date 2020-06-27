package main

import (
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/tools/util"
	"net/http"
)

// 给请求包一层IP验证
func WrapServeHTTP(a api.FullNode, rpcServer *jsonrpc.RPCServer) func(resp http.ResponseWriter, req *http.Request) {
	return func(resp http.ResponseWriter, req *http.Request) {
		rAddr := util.RemoteIPFromReqAddr(req.RemoteAddr)
		if err := a.HostVerify(req.Context(), rAddr); err != nil {
			log.Warnf("Host Verification failed: %s", rAddr)
			resp.WriteHeader(401)
			return
		}
		rpcServer.ServeHTTP(resp, req)
	}
}
