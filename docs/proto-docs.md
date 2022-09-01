<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [v1/beefy.proto](#v1/beefy.proto)
    - [BeefyAuthoritySet](#beefy.v1.BeefyAuthoritySet)
    - [BeefyMmrLeaf](#beefy.v1.BeefyMmrLeaf)
    - [BeefyMmrLeafPartial](#beefy.v1.BeefyMmrLeafPartial)
    - [ClientState](#beefy.v1.ClientState)
    - [ClientStateUpdateProof](#beefy.v1.ClientStateUpdateProof)
    - [Commitment](#beefy.v1.Commitment)
    - [CommitmentSignature](#beefy.v1.CommitmentSignature)
    - [ConsensusState](#beefy.v1.ConsensusState)
    - [ConsensusStateUpdateProof](#beefy.v1.ConsensusStateUpdateProof)
    - [Header](#beefy.v1.Header)
    - [Misbehaviour](#beefy.v1.Misbehaviour)
    - [ParachainHeader](#beefy.v1.ParachainHeader)
    - [PayloadItem](#beefy.v1.PayloadItem)
    - [SignedCommitment](#beefy.v1.SignedCommitment)
  
    - [RelayChain](#beefy.v1.RelayChain)
  
- [Scalar Value Types](#scalar-value-types)



<a name="v1/beefy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/beefy.proto



<a name="beefy.v1.BeefyAuthoritySet"></a>

### BeefyAuthoritySet
Beefy Authority Info


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | Id of the authority set, it should be strictly increasing |
| `len` | [uint32](#uint32) |  | size of the authority set |
| `authority_root` | [bytes](#bytes) |  | merkle root of the sorted authority public keys. |






<a name="beefy.v1.BeefyMmrLeaf"></a>

### BeefyMmrLeaf
BeefyMmrLeaf leaf data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [uint32](#uint32) |  | leaf version |
| `parent_number` | [uint32](#uint32) |  | parent block for this leaf |
| `parent_hash` | [bytes](#bytes) |  | parent hash for this leaf |
| `beefy_next_authority_set` | [BeefyAuthoritySet](#beefy.v1.BeefyAuthoritySet) |  | beefy next authority set. |
| `parachain_heads` | [bytes](#bytes) |  | merkle root hash of parachain heads included in the leaf. |






<a name="beefy.v1.BeefyMmrLeafPartial"></a>

### BeefyMmrLeafPartial
Partial data for MmrLeaf


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [uint32](#uint32) |  | leaf version |
| `parent_number` | [uint32](#uint32) |  | parent block for this leaf |
| `parent_hash` | [bytes](#bytes) |  | parent hash for this leaf |
| `beefy_next_authority_set` | [BeefyAuthoritySet](#beefy.v1.BeefyAuthoritySet) |  | next authority set. |






<a name="beefy.v1.ClientState"></a>

### ClientState
ClientState from Beefy tracks the current validator set, latest height,
and a possible frozen height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mmr_root_hash` | [bytes](#bytes) |  | Latest mmr root hash |
| `latest_beefy_height` | [uint32](#uint32) |  | block number for the latest mmr_root_hash |
| `frozen_height` | [uint64](#uint64) |  | Block height when the client was frozen due to a misbehaviour |
| `relay_chain` | [RelayChain](#beefy.v1.RelayChain) |  | Known relay chains |
| `para_id` | [uint32](#uint32) |  | ParaId of associated parachain |
| `latest_para_height` | [uint32](#uint32) |  | latest parachain height |
| `beefy_activation_block` | [uint32](#uint32) |  | block number that the beefy protocol was activated on the relay chain. This should be the first block in the merkle-mountain-range tree. |
| `authority` | [BeefyAuthoritySet](#beefy.v1.BeefyAuthoritySet) |  | authorities for the current round |
| `next_authority_set` | [BeefyAuthoritySet](#beefy.v1.BeefyAuthoritySet) |  | authorities for the next round |






<a name="beefy.v1.ClientStateUpdateProof"></a>

### ClientStateUpdateProof
data needed to update the client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mmr_leaf` | [BeefyMmrLeaf](#beefy.v1.BeefyMmrLeaf) |  | the new mmr leaf SCALE encoded. |
| `mmr_leaf_index` | [uint64](#uint64) |  | leaf index for the mmr_leaf |
| `mmr_proof` | [bytes](#bytes) | repeated | proof that this mmr_leaf index is valid. |
| `signed_commitment` | [SignedCommitment](#beefy.v1.SignedCommitment) |  | signed commitment data |
| `authorities_proof` | [bytes](#bytes) | repeated | generated using full authority list from runtime |






<a name="beefy.v1.Commitment"></a>

### Commitment
Commitment message signed by beefy validators


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payload` | [PayloadItem](#beefy.v1.PayloadItem) | repeated | array of payload items signed by Beefy validators |
| `block_numer` | [uint32](#uint32) |  | block number for this commitment |
| `validator_set_id` | [uint64](#uint64) |  | validator set that signed this commitment |






<a name="beefy.v1.CommitmentSignature"></a>

### CommitmentSignature
Signature belonging to a single validator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signature` | [bytes](#bytes) |  | actual signature bytes |
| `authority_index` | [uint32](#uint32) |  | authority leaf index in the merkle tree. |






<a name="beefy.v1.ConsensusState"></a>

### ConsensusState
ConsensusState defines the consensus state from Tendermint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp that corresponds to the block height in which the ConsensusState was stored. |
| `root` | [bytes](#bytes) |  | packet commitment root |






<a name="beefy.v1.ConsensusStateUpdateProof"></a>

### ConsensusStateUpdateProof
Parachain headers and their mmr proofs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `parachain_headers` | [ParachainHeader](#beefy.v1.ParachainHeader) | repeated | parachain headers needed for proofs and ConsensusState |
| `mmr_proofs` | [bytes](#bytes) | repeated | mmr proofs for the headers gotten from rpc "mmr_generateProofs" |
| `mmr_size` | [uint64](#uint64) |  | size of the mmr for the given proof |






<a name="beefy.v1.Header"></a>

### Header
Header contains the neccessary data to prove finality about IBC commitments


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `consensus_state_update` | [ConsensusStateUpdateProof](#beefy.v1.ConsensusStateUpdateProof) |  | optional payload to update ConsensusState |
| `client_state` | [ClientStateUpdateProof](#beefy.v1.ClientStateUpdateProof) |  | optional payload to update the ClientState. |






<a name="beefy.v1.Misbehaviour"></a>

### Misbehaviour
Misbehaviour is a wrapper over two conflicting Headers
that implements Misbehaviour interface expected by ICS-02


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header_1` | [Header](#beefy.v1.Header) |  |  |
| `header_2` | [Header](#beefy.v1.Header) |  |  |






<a name="beefy.v1.ParachainHeader"></a>

### ParachainHeader
data needed to prove parachain header inclusion in mmr.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `parachain_header` | [bytes](#bytes) |  | scale-encoded parachain header bytes |
| `mmr_leaf_partial` | [BeefyMmrLeafPartial](#beefy.v1.BeefyMmrLeafPartial) |  | see beefy spec |
| `parachain_heads_proof` | [bytes](#bytes) | repeated | proofs for our header in the parachain heads root |
| `heads_leaf_index` | [uint32](#uint32) |  | leaf index for parachain heads proof |
| `heads_total_count` | [uint32](#uint32) |  | total number of para heads in parachain_heads_root |
| `extrinsic_proof` | [bytes](#bytes) | repeated | trie merkle proof of inclusion in header.extrinsic_root |
| `timestamp_extrinsic` | [bytes](#bytes) |  | the actual timestamp extrinsic |






<a name="beefy.v1.PayloadItem"></a>

### PayloadItem
Actual payload items


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payload_id` | [bytes](#bytes) |  | 2-byte payload id |
| `payload_data` | [bytes](#bytes) |  | arbitrary length payload data., eg mmr_root_hash |






<a name="beefy.v1.SignedCommitment"></a>

### SignedCommitment
signed commitment data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commitment` | [Commitment](#beefy.v1.Commitment) |  | commitment data being signed |
| `signatures` | [CommitmentSignature](#beefy.v1.CommitmentSignature) | repeated | gotten from rpc subscription |





 <!-- end messages -->


<a name="beefy.v1.RelayChain"></a>

### RelayChain


| Name | Number | Description |
| ---- | ------ | ----------- |
| POLKADOT | 0 |  |
| KUSAMA | 1 |  |
| ROCOCO | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

