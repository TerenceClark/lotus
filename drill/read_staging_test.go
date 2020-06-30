package drill

import (
	"context"
	"fmt"
	"github.com/filecoin-project/lotus/api"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/urfave/cli/v2"
	"os"
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
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "test.v"},
		&cli.StringFlag{Name: "repo", Value: "/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotus-data"},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("exec action")
		repoPath := "/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotusstorage"
		//nodeType := repo.StorageMiner
		r, err := repo.NewFS(repoPath)
		if err != nil {
			panic(err)
		}

		ok, err := r.Exists()
		if err != nil {
			panic(err)
		}
		if !ok {
			panic(fmt.Sprintf("repo at '%s' is not initialized, run 'lotus-storage-miner init' to set it up", repoPath))
		}
		ctx := context.Background()

		os.Setenv("LOTUS_PATH", "/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotus-data")
		os.Setenv("LOTUS_STORAGE_PATH", "/home/ipfsmain/workspace/filecoin/lotus-dev-env/lotusstorage")
		os.Setenv("FIL_PROOFS_PARAMETER_CACHE", "/home/ipfsmain/workspace/filecoin/lotus-dev-env/filecoin-proof-parameters")
		nodeApi, ncloser, err := lcli.GetFullNodeAPI(c)
		if err != nil {
			panic(err)
		}
		defer ncloser()


		shutdownChan := make(chan struct{})

		var minerapi api.StorageMiner
		stop, err := node.New(ctx,
			node.StorageMiner(&minerapi),
			node.Override(new(dtypes.ShutdownChan), shutdownChan),
			node.Online(),
			node.Repo(r),

			//node.ApplyIf(func(s *node.Settings) bool { return cctx.IsSet("api") },
			//	node.Override(new(dtypes.APIEndpoint), func() (dtypes.APIEndpoint, error) {
			//		return multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/" + cctx.String("api"))
			//	})),
			node.Override(new(api.FullNode), nodeApi),
			node.Override(node.ExecFunc, func() {
				fmt.Println("==== call exec func")
			}),
		)
		if err != nil {
			panic(err)
		}
		if err = stop(ctx); err != nil {
			panic(err)
		}
		return nil
	}
	app.Run(os.Args)
}
