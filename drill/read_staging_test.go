package drill

import (
	"context"
	"fmt"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/prometheus/common/log"
	"go.opencensus.io/plugin/runmetrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"reflect"
	"testing"
)

func TestReadStaging(t *testing.T) {
	r, err := repo.NewFS("/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotusstorage")
	if err != nil {
		panic(err)
	}
	lockedRepo, err := r.Lock(repo.StorageMiner)
	if err != nil {
		panic(err)
	}
	stagingStore, err := modules.StagingBlockstore(lockedRepo)
	if err != nil {
		panic(err)
	}
	keys, err := stagingStore.AllKeysChan(context.Background())
	if err != nil {
		panic(err)
	}
	for k := range keys {
		fmt.Println(k.String())
	}
	//dataCid, err := cid.Parse("bafkreif5jqx5vkj4bdv3udtjsh3zcp6oorxx6omouzxsfxcrhihagcoipq")
	//if err != nil {
	//	panic(err)
	//}
	//blk, err := stagingStore.Get(dataCid)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(blk.String())
}

func TestReadStagingByDAG(t *testing.T) {
	repoPath := "/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotusstorage"
	nodeType := repo.StorageMiner
	err := runmetrics.Enable(runmetrics.RunMetricOptions{
		EnableCPU:    true,
		EnableMemory: true,
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	r, err := repo.NewFS(repoPath)
	if err != nil {
		panic(fmt.Sprintf("opening fs repo: %s", err))
	}

	if err := r.Init(nodeType); err != nil && err != repo.ErrRepoExists {
		panic(fmt.Sprintf("repo init error: %s", err))
	}

	if err := paramfetch.GetParams(ctx, build.ParametersJSON(), 0); err != nil {
		panic(fmt.Sprintf("fetching proof parameters: %s", err))
	}

	var genBytes []byte
	genBytes = build.MaybeGenesis()

	genesis := node.Options()
	if len(genBytes) > 0 {
		genesis = node.Override(new(modules.Genesis), modules.LoadGenesis(genBytes))
	}

	shutdownChan := make(chan struct{})

	var fullNodeAPI api.FullNode

	stop, err := node.New(ctx,
		node.Override(new(*dtypes.RPCHostVerifier), &dtypes.RPCHostVerifier{
			ValidHosts: []string{"127.0.0.1"},
		}),
		node.FullAPI(&fullNodeAPI),

		node.Override(new(dtypes.Bootstrapper), dtypes.Bootstrapper(false)),
		node.Override(new(dtypes.ShutdownChan), shutdownChan),
		node.Online(),
		node.Repo(r),
		node.Override(node.ExecFunc, func(dag dtypes.StagingDAG) {
				fmt.Println("----- got dag", reflect.TypeOf(dag))

		}),

		genesis,
	)
	if err != nil {
		panic(fmt.Sprintf("initializing node: %s", err))
	}

	// Register all metric views
	if err = view.Register(
		metrics.DefaultViews...,
	); err != nil {
		log.Fatalf("Cannot register the view: %v", err)
	}

	// Set the metric to one so it is published to the exporter
	stats.Record(ctx, metrics.LotusInfo.M(1))

	if err := stop(ctx); err != nil {
		panic(err)
	}
}
