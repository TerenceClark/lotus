package drill

import (
	"fmt"
	"github.com/filecoin-project/specs-actors/actors/builtin"
	"testing"
)

func TestPrintAddr(t *testing.T) {
	fmt.Println(builtin.SystemActorAddr)
}
