package seed

import (
	"encoding/hex"
	"encoding/json"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/genesis"
	"io/ioutil"
	"path/filepath"
)

type MinerSeedInfo struct {
	MAddr address.Address
	Gm *genesis.Miner
	Key *types.KeyInfo
}

func WriteGenesisMiners(sbroot string, miners []*MinerSeedInfo) error {
	for _, mi := range miners {
		maddr := mi.MAddr
		gm := mi.Gm
		key := mi.Key
		output := map[string]genesis.Miner{
			maddr.String(): *gm,
		}

		out, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return err
		}

		log.Infof("Writing preseal manifest to %s", filepath.Join(sbroot, "pre-seal-"+maddr.String()+".json"))

		if err := ioutil.WriteFile(filepath.Join(sbroot, "pre-seal-"+maddr.String()+".json"), out, 0664); err != nil {
			return err
		}

		if key != nil {
			b, err := json.Marshal(key)
			if err != nil {
				return err
			}

			// TODO: allow providing key
			if err := ioutil.WriteFile(filepath.Join(sbroot, "pre-seal-"+maddr.String()+".key"), []byte(hex.EncodeToString(b)), 0664); err != nil {
				return err
			}
		}
	}

	return nil
}