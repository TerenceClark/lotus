package node

import (
	"errors"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/lp2p"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-peerstore/pstoremem"
)

func RepoOfStorageMiner() Option {
	return Options(
		ApplyIf(func(s *Settings) bool { return s.Config },
			Error(errors.New("the StorageMiner option must be set before Config option")),
		),
		ApplyIf(func(s *Settings) bool { return s.Online },
			Error(errors.New("the StorageMiner option must be set before Online option")),
		),

		func(s *Settings) error {
			s.nodeType = repo.StorageMiner
			return nil
		},

		//Override(new(api.Common), From(new(common.CommonAPI))),
		//Override(new(sectorstorage.StorageAuth), modules.StorageAuth),
		//
		//Override(new(*stores.Index), stores.NewIndex),
		//Override(new(stores.SectorIndex), From(new(*stores.Index))),
		//Override(new(dtypes.MinerID), modules.MinerID),
		//Override(new(dtypes.MinerAddress), modules.MinerAddress),
		//Override(new(*ffiwrapper.Config), modules.ProofsConfig),
		//Override(new(stores.LocalStorage), From(new(repo.LockedRepo))),
		//Override(new(sealing.SectorIDCounter), modules.SectorIDCounter),
		//Override(new(*sectorstorage.Manager), modules.SectorStorage),
		//Override(new(ffiwrapper.Verifier), ffiwrapper.ProofVerifier),
		//
		//Override(new(sectorstorage.SectorManager), From(new(*sectorstorage.Manager))),
		//Override(new(storage2.Prover), From(new(sectorstorage.SectorManager))),
		//
		//Override(new(*sectorblocks.SectorBlocks), sectorblocks.NewSectorBlocks),
		//Override(new(*storage.Miner), modules.StorageMiner),
		//Override(new(dtypes.NetworkName), modules.StorageNetworkName),
		Override(DefaultTransportsKey, lp2p.DefaultTransports),
		Override(new(host.Host), lp2p.RoutedHost),
		Override(new(routing.Routing), lp2p.Routing),
		Override(new(lp2p.RawHost), lp2p.Host),
		Override(new(lp2p.BaseIpfsRouting), lp2p.DHTRouting(dht.ModeAuto)),
		Override(new(peerstore.Peerstore), pstoremem.NewPeerstore),
		Override(new(dtypes.StagingBlockstore), modules.StagingBlockstore),
		Override(new(dtypes.StagingDAG), modules.StagingDAG),
		Override(new(dtypes.StagingGraphsync), modules.StagingGraphsync),
		//Override(new(retrievalmarket.RetrievalProvider), modules.RetrievalProvider),
		//Override(new(dtypes.ProviderDealStore), modules.NewProviderDealStore),
		//Override(new(dtypes.ProviderDataTransfer), modules.NewProviderDAGServiceDataTransfer),
		//Override(new(dtypes.ProviderRequestValidator), modules.NewProviderRequestValidator),
		//Override(new(dtypes.ProviderPieceStore), modules.NewProviderPieceStore),
		//Override(new(*storedask.StoredAsk), modules.NewStorageAsk),
		//Override(new(storagemarket.StorageProvider), modules.StorageProvider),
		//Override(new(storagemarket.StorageProviderNode), storageadapter.NewProviderNodeAdapter),
		//Override(RegisterProviderValidatorKey, modules.RegisterProviderValidator),
		//Override(HandleRetrievalKey, modules.HandleRetrieval),
		//Override(GetParamsKey, modules.GetParams),
		//Override(HandleDealsKey, modules.HandleDeals),
		//Override(new(gen.WinningPoStProver), storage.NewWinningPoStProver),
		//Override(new(*miner.Miner), modules.SetupBlockProducer),

		//Override(new(dtypes.AcceptingRetrievalDealsConfigFunc), modules.NewAcceptingRetrievalDealsConfigFunc),
		//Override(new(dtypes.SetAcceptingRetrievalDealsConfigFunc), modules.NewSetAcceptingRetrievalDealsConfigFunc),
		//Override(new(dtypes.AcceptingStorageDealsConfigFunc), modules.NewAcceptingStorageDealsConfigFunc),
		//Override(new(dtypes.SetAcceptingStorageDealsConfigFunc), modules.NewSetAcceptingStorageDealsConfigFunc),
		//Override(new(dtypes.StorageDealPieceCidBlocklistConfigFunc), modules.NewStorageDealPieceCidBlocklistConfigFunc),
		//Override(new(dtypes.SetStorageDealPieceCidBlocklistConfigFunc), modules.NewSetStorageDealPieceCidBlocklistConfigFunc),
	)
}
