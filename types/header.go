package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"

	"github.com/ChainSafe/gossamer/lib/trie"
	"github.com/ComposableFi/go-substrate-rpc-client/v4/scale"
	rpcclienttypes "github.com/ComposableFi/go-substrate-rpc-client/v4/types"
	codec "github.com/ComposableFi/go-substrate-rpc-client/v4/types/codec"
	"github.com/ComposableFi/ics11-beefy/exported"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	ics02 "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

var _ ibcexported.Header = &Header{}

const revisionNumber = 0

type Head []byte

type HeadData struct {
	Head
}

// DecodeParachainHeader decodes an encoded substrate header to a concrete Header type. It takes encoded bytes
// as an argument and returns a concrete substrate Header type.
func DecodeParachainHeader(hb []byte) (rpcclienttypes.Header, error) {
	var headData HeadData
	err := codec.Decode(hb, &headData)
	if err != nil {
		return rpcclienttypes.Header{}, err
	}

	var h rpcclienttypes.Header
	err = codec.Decode(headData.Head, &h)
	if err != nil {
		return rpcclienttypes.Header{}, err
	}
	return h, nil
}

// DecodeExtrinsicTimestamp decodes a scale encoded timestamp to a time.Time type
func DecodeExtrinsicTimestamp(encodedExtrinsic []byte) (time.Time, error) {
	var extrinsic rpcclienttypes.Extrinsic
	decodeErr := codec.Decode(encodedExtrinsic, &extrinsic)
	if decodeErr != nil {
		return time.Time{}, decodeErr
	}

	unix, unixDecodeErr := scale.NewDecoder(bytes.NewReader(extrinsic.Method.Args[:])).DecodeUintCompact()
	if unixDecodeErr != nil {
		return time.Time{}, unixDecodeErr
	}
	t := time.UnixMilli(unix.Int64())

	return t, nil
}

// ConsensusState returns the updated consensus state associated with the header
func (h Header) ConsensusState() *ConsensusState {
	parachainHeader, err := DecodeParachainHeader(h.HeadersWithProof.Headers[0].ParachainHeader)
	if err != nil {
		log.Fatal(err)
	}

	rootHash := parachainHeader.StateRoot[:]

	return &ConsensusState{
		Root: rootHash,
	}
}

// ClientType defines that the Header is a Beefy consensus algorithm
func (h Header) ClientType() string {
	return exported.Beefy
}

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetHeight() ibcexported.Height {
	parachainHeader, err := DecodeParachainHeader(h.HeadersWithProof.Headers[0].ParachainHeader)
	if err != nil {
		log.Fatal(err)
	}
	return ics02.NewHeight(revisionNumber, uint64(parachainHeader.Number))
}

// ValidateBasic calls the SignedHeader ValidateBasic function and checks
// that validatorsets are not nil.
// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
// with MsgCreateClient
func (h Header) ValidateBasic() error {
	for _, header := range h.HeadersWithProof.Headers {
		decHeader, err := DecodeParachainHeader(header.ParachainHeader)
		if err != nil {
			return err
		}

		rootHash := decHeader.ExtrinsicsRoot[:]
		extrinsicsProof := header.ExtrinsicProof

		key := make([]byte, 4)
		binary.LittleEndian.PutUint32(key, 0)
		t := trie.NewEmptyTrie()
		if err := t.LoadFromProof(extrinsicsProof, rootHash); err != nil {
			return err
		}

		if ext := t.Get(key); len(ext) == 0 {
			return fmt.Errorf("invalid key length")
		}

		// todo: decode extrinsic.
	}

	return nil
}

// GetPubKey unmarshals the new public key into a cryptotypes.PubKey type.
// An error is returned if the new public key is nil or the cached value
// is not a PubKey.
func (h Header) GetPubKey() (cryptotypes.PubKey, error) {
	return nil, nil
}
