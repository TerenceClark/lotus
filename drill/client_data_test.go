package drill

import (
	"context"
	"fmt"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/filecoin-project/lotus/node/impl"
	"github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"
	unixfile "github.com/ipfs/go-unixfs/file"
	"github.com/prometheus/common/log"
	"go.opencensus.io/plugin/runmetrics"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/metrics"
	"github.com/filecoin-project/lotus/node"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/repo"
	"testing"
)

func TestClientImportData(t *testing.T) {
	err := runmetrics.Enable(runmetrics.RunMetricOptions{
		EnableCPU:    true,
		EnableMemory: true,
	})
	if err != nil {
		panic(err)
	}
	//if prof := cctx.String("pprof"); prof != "" {
	//	profile, err := os.Create(prof)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if err := pprof.StartCPUProfile(profile); err != nil {
	//		return err
	//	}
	//	defer pprof.StopCPUProfile()
	//}

	//var isBootstrapper dtypes.Bootstrapper
	//switch profile := cctx.String("profile"); profile {
	//case "bootstrapper":
	//	isBootstrapper = true
	//case "":
	//	// do nothing
	//default:
	//	return fmt.Errorf("unrecognized profile type: %q", profile)
	//}

	//ctx, _ := tag.New(context.Background(), tag.Insert(metrics.Version, build.BuildVersion), tag.Insert(metrics.Commit, build.CurrentCommit))
	//{
	//	dir, err := homedir.Expand(cctx.String("repo"))
	//	if err != nil {
	//		log.Warnw("could not expand repo location", "error", err)
	//	} else {
	//		log.Infof("lotus repo: %s", dir)
	//	}
	//}
	ctx := context.Background()

	r, err := repo.NewFS("~/.lotus")
	if err != nil {
		panic(fmt.Sprintf("opening fs repo: %s", err))
	}

	if err := r.Init(repo.FullNode); err != nil && err != repo.ErrRepoExists {
		panic(fmt.Sprintf("repo init error: %s", err))
	}

	if err := paramfetch.GetParams(ctx, build.ParametersJSON(), 0); err != nil {
		panic(fmt.Sprintf("fetching proof parameters: %s", err))
	}

	var genBytes []byte
	//if cctx.String("genesis") != "" {
	//	genBytes, err = ioutil.ReadFile(cctx.String("genesis"))
	//	if err != nil {
	//		return xerrors.Errorf("reading genesis: %w", err)
	//	}
	//} else {
		genBytes = build.MaybeGenesis()
	//}

	//chainfile := cctx.String("import-chain")
	//if chainfile != "" {
	//	chainfile, err := homedir.Expand(chainfile)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if err := ImportChain(r, chainfile); err != nil {
	//		return err
	//	}
	//	if cctx.Bool("halt-after-import") {
	//		fmt.Println("Chain import complete, halting as requested...")
	//		return nil
	//	}
	//}

	genesis := node.Options()
	if len(genBytes) > 0 {
		genesis = node.Override(new(modules.Genesis), modules.LoadGenesis(genBytes))
	}
	//if cctx.String(makeGenFlag) != "" {
	//	if cctx.String(preTemplateFlag) == "" {
	//		return xerrors.Errorf("must also pass file with genesis template to `--%s`", preTemplateFlag)
	//	}
	//	genesis = node.Override(new(modules.Genesis), testing.MakeGenesis(cctx.String(makeGenFlag), cctx.String(preTemplateFlag)))
	//}

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

		genesis,

		//node.ApplyIf(func(s *node.Settings) bool { return false },
		//	node.Override(node.SetApiEndpointKey, func(lr repo.LockedRepo) error {
		//		apima, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/" +
		//			cctx.String("api"))
		//		if err != nil {
		//			return err
		//		}
		//		return lr.SetAPIEndpoint(apima)
		//	})),
		//node.ApplyIf(func(s *node.Settings) bool { return !cctx.Bool("bootstrap") },
		//	node.Unset(node.RunPeerMgrKey),
		//	node.Unset(new(*peermgr.PeerMgr)
		//),
	)
	if err != nil {
		panic(fmt.Sprintf("initializing node: %s", err))
	}

	//if cctx.String("import-key") != "" {
	//	if err := importKey(ctx, api, cctx.String("import-key")); err != nil {
	//		log.Errorf("importing key failed: %+v", err)
	//	}
	//}

	// Register all metric views
	if err = view.Register(
		metrics.DefaultViews...,
	); err != nil {
		log.Fatalf("Cannot register the view: %v", err)
	}

	// Set the metric to one so it is published to the exporter
	stats.Record(ctx, metrics.LotusInfo.M(1))

	//dataCid, err := fullNodeAPI.ClientImport(ctx, api.FileRef{
	//	Path:  "/home/ipfsmain/tmp/Git-2.27.0-64-bit.exe.1",
	//	IsCAR: false,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("import data cid:", dataCid.String())

	dataCid, err := cid.Parse("QmZ8dRwcRQWPPcFkud4rF19bVVSdxkm3TurkARJJWkmQYP")
	if err != nil {
		panic(err)
	}
	//fullNodeAPI.(*client.API).LocalDAG.Get(ctx, dataCid)
	localDAG := fullNodeAPI.(*impl.FullNodeAPI).LocalDAG
	dataNode, err := localDAG.Get(ctx, dataCid)
	file, err := unixfile.NewUnixfsFile(ctx, localDAG, dataNode)
	if err != nil {
		panic(err)
	}
	if err = files.WriteTo(file, "/home/ipfsmain/tmp/Git-2.27.0-64-bit.exe.1_from_ipfs.tgz"); err != nil {
		panic(err)
	}

	if err := stop(ctx); err != nil {
		panic(err)
	}
	//endpoint, err := r.APIEndpoint()
	//if err != nil {
	//	panic(fmt.Sprintf("getting api endpoint: %s", err))
	//}

	// TODO: properly parse api endpoint (or make it a URL)
	//return serveRPC(api, stop, endpoint, shutdownChan)
}
