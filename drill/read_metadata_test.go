package drill

import (
	"fmt"
	"github.com/filecoin-project/lotus/node/modules"
	"github.com/filecoin-project/lotus/node/repo"
	"github.com/ipfs/go-cid"
	"testing"
)

func TestReadMetaData(t *testing.T) {
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
	lockedRepo, err := r.Lock(repo.StorageMiner)
	metaDS, err := modules.Datastore(lockedRepo)
	if err != nil {
		panic(err)
	}
	providerPieceStore := modules.NewProviderPieceStore(metaDS)
	payloadCid, err := cid.Parse("bafkreibhhyqvrp5dh4axmspuvpda43ajmrr672vi5bwz3r4mf24eclf7ca")
	if err != nil {
		panic(err)
	}
	cidInfo, err := providerPieceStore.GetCIDInfo(payloadCid)
	if err != nil {
		panic(err)
	}
	for _, pbl := range cidInfo.PieceBlockLocations {
		fmt.Println("cid info:", cidInfo.CID.String(), pbl.RelOffset)
	}
	pieceCid, err := cid.Parse("bafk4chza5bnhp6y24uchqmvrgjuhe5vrliuqk7cjzodtyi5ss2zllyorv4ea")
	if err != nil {
		panic(err)
	}
	pieceInfo, err := providerPieceStore.GetPieceInfo(pieceCid)
	if err != nil {
		panic(err)
	}
	for _, deal := range pieceInfo.Deals {
		fmt.Println("piece info", deal.DealID, deal.SectorID)
	}
}
