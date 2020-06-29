package drill

import (
	"context"
	"fmt"
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"io/ioutil"
	"testing"
)

func TestGetParams(t *testing.T) {
	// gate way 可能下不到最新的 params
	//os.Setenv("IPFS_GATEWAY", "https://proof-parameters.s3.cn-south-1.jdcloud-oss.com/ipfs/")
	//os.Setenv("https_proxy", "http://127.0.0.1:58591")
	sectorSizeInt, err := units.RAMInBytes("2KiB")
	if err != nil {
		panic(err)
	}
	ssize := abi.SectorSize(sectorSizeInt)
	fmt.Println(ssize)
	paramsConf, err := ioutil.ReadFile("/home/ipfsmain/go/src/github.com/filecoin-project/lotus/build/proof-params/parameters.json")
	if err := paramfetch.GetParams(context.Background(), paramsConf, uint64(ssize)); err != nil {
		panic(err)
	}
}
