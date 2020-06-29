package drill

import (
	"context"
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetParams(t *testing.T) {
	sectorSizeInt, err := units.RAMInBytes("2KiB")
	if err != nil {
		panic(err)
	}
	os.Setenv("IPFS_GATEWAY", "https://proof-parameters.s3.cn-south-1.jdcloud-oss.com/ipfs/")
	ssize := abi.SectorSize(sectorSizeInt)
	paramsConf, err := ioutil.ReadFile("/home/ipfsmain/go/src/github.com/filecoin-project/lotus/build/proof-params/parameters.json")
	if err := paramfetch.GetParams(context.Background(), paramsConf, uint64(ssize)); err != nil {
		panic(err)
	}
}
