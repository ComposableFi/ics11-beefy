package types_test

import (
	"testing"

	rpcclienttypes "github.com/ComposableFi/go-substrate-rpc-client/v4/types"

	"github.com/ComposableFi/ics11-beefy/types"
	"github.com/stretchr/testify/suite"
)

type BeefyTestSuite struct {
	suite.Suite
}

func (suite *BeefyTestSuite) SetupTest() {

}

func TestSoloMachineTestSuite(t *testing.T) {
	suite.Run(t, new(BeefyTestSuite))
}

func TestScaleEncodeCustomTypes(t *testing.T) {
	var sb types.SizedByte32 = [32]byte{1, 2}
	_, err := rpcclienttypes.Encode(sb)
	if err != nil {
		t.Error(err)
	}

	var sb2 types.SizedByte2 = [2]byte{1, 2}
	_, err = rpcclienttypes.Encode(sb2)
	if err != nil {
		t.Error(err)
	}

	var u8 types.U8 = 1
	_, err = rpcclienttypes.Encode(u8)
	if err != nil {
		t.Error(err)
	}
}
