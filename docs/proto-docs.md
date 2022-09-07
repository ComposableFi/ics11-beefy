<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [v1/beefy.proto](#v1/beefy.proto)
    - [BeefyAuthoritySet](#ibc.lightclients.beefy.v1.BeefyAuthoritySet)
    - [BeefyMMRLeaf](#ibc.lightclients.beefy.v1.BeefyMMRLeaf)
    - [ClientState](#ibc.lightclients.beefy.v1.ClientState)
    - [Commitment](#ibc.lightclients.beefy.v1.Commitment)
    - [CommitmentSignature](#ibc.lightclients.beefy.v1.CommitmentSignature)
    - [ConsensusState](#ibc.lightclients.beefy.v1.ConsensusState)
    - [Header](#ibc.lightclients.beefy.v1.Header)
    - [MMRUpdateProof](#ibc.lightclients.beefy.v1.MMRUpdateProof)
    - [Misbehaviour](#ibc.lightclients.beefy.v1.Misbehaviour)
    - [ParachainHeader](#ibc.lightclients.beefy.v1.ParachainHeader)
    - [ParachainHeadersWithProof](#ibc.lightclients.beefy.v1.ParachainHeadersWithProof)
    - [PartialMMRLeaf](#ibc.lightclients.beefy.v1.PartialMMRLeaf)
    - [Payload](#ibc.lightclients.beefy.v1.Payload)
    - [SignedCommitment](#ibc.lightclients.beefy.v1.SignedCommitment)
  
    - [RelayChain](#ibc.lightclients.beefy.v1.RelayChain)
  
- [Scalar Value Types](#scalar-value-types)



<a name="v1/beefy.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/beefy.proto



<a name="ibc.lightclients.beefy.v1.BeefyAuthoritySet"></a>

### BeefyAuthoritySet
Beefy Authority Info


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | Id of the authority set, it should be strictly increasing |
| `len` | [uint32](#uint32) |  | size of the authority set |
| `authority_root` | [bytes](#bytes) |  | merkle root of the sorted authority public keys. |






<a name="ibc.lightclients.beefy.v1.BeefyMMRLeaf"></a>

### BeefyMMRLeaf
BeefyMMRLeaf leaf data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [uint32](#uint32) |  | leaf version |
| `parent_number` | [uint32](#uint32) |  | parent block for this leaf |
| `parent_hash` | [bytes](#bytes) |  | parent hash for this leaf |
| `beefy_next_authority_set` | [BeefyAuthoritySet](#ibc.lightclients.beefy.v1.BeefyAuthoritySet) |  | beefy next authority set. |
| `parachain_heads` | [bytes](#bytes) |  | merkle root hash of parachain heads included in the leaf. |






<a name="ibc.lightclients.beefy.v1.ClientState"></a>

### ClientState
ClientState from Beefy tracks the current validator set, latest height,
and a possible frozen height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mmr_root_hash` | [bytes](#bytes) |  | Latest mmr root hash |
| `latest_beefy_height` | [uint32](#uint32) |  | block number for the latest mmr_root_hash |
| `frozen_height` | [uint64](#uint64) |  | Block height when the client was frozen due to a misbehaviour |
| `relay_chain` | [RelayChain](#ibc.lightclients.beefy.v1.RelayChain) |  | Known relay chains |
| `para_id` | [uint32](#uint32) |  | ParaId of associated parachain |
| `latest_para_height` | [uint32](#uint32) |  | latest parachain height |
| `beefy_activation_block` | [uint32](#uint32) |  | block number that the beefy protocol was activated on the relay chain. This should be the first block in the merkle-mountain-range tree. |
| `authority` | [BeefyAuthoritySet](#ibc.lightclients.beefy.v1.BeefyAuthoritySet) |  | authorities for the current round |
| `next_authority_set` | [BeefyAuthoritySet](#ibc.lightclients.beefy.v1.BeefyAuthoritySet) |  | authorities for the next round |






<a name="ibc.lightclients.beefy.v1.Commitment"></a>

### Commitment
Commitment message signed by beefy validators


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payload` | [Payload](#ibc.lightclients.beefy.v1.Payload) | repeated | array of payload items signed by Beefy validators |
| `block_numer` | [uint32](#uint32) |  | block number for this commitment |
| `validator_set_id` | [uint64](#uint64) |  | validator set that signed this commitment |






<a name="ibc.lightclients.beefy.v1.CommitmentSignature"></a>

### CommitmentSignature
Signature belonging to a single validator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signature` | [bytes](#bytes) |  | actual signature bytes |
| `authority_index` | [uint32](#uint32) |  | authority leaf index in the merkle tree. |






<a name="ibc.lightclients.beefy.v1.ConsensusState"></a>

### ConsensusState
ConsensusState defines the consensus state from Tendermint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timestamp that corresponds to the block height in which the ConsensusState was stored. |
| `root` | [bytes](#bytes) |  | packet commitment root |






<a name="ibc.lightclients.beefy.v1.Header"></a>

### Header
Header contains the neccessary data to prove finality about IBC commitments


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `headers_with_proof` | [ParachainHeadersWithProof](#ibc.lightclients.beefy.v1.ParachainHeadersWithProof) |  | optional payload to update ConsensusState |
| `mmr_update_proof` | [MMRUpdateProof](#ibc.lightclients.beefy.v1.MMRUpdateProof) |  | optional payload to update the ClientState. |






<a name="ibc.lightclients.beefy.v1.MMRUpdateProof"></a>

### MMRUpdateProof
data needed to update the client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signed_commitment` | [SignedCommitment](#ibc.lightclients.beefy.v1.SignedCommitment) |  | signed commitment data |
| `latest_mmr_leaf_index` | [uint64](#uint64) |  | leaf index for the mmr_leaf |
| `latest_mmr_leaf` | [BeefyMMRLeaf](#ibc.lightclients.beefy.v1.BeefyMMRLeaf) |  | the new mmr leaf SCALE encoded. |
| `mmr_proof` | [bytes](#bytes) | repeated | proof that this mmr_leaf index is valid. |
| `authorities_proof` | [bytes](#bytes) | repeated | generated using full authority list from runtime |






<a name="ibc.lightclients.beefy.v1.Misbehaviour"></a>

### Misbehaviour
Misbehaviour is a wrapper over two conflicting Headers
that implements Misbehaviour interface expected by ICS-02


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header1` | [Header](#ibc.lightclients.beefy.v1.Header) |  |  |
| `header2` | [Header](#ibc.lightclients.beefy.v1.Header) |  |  |






<a name="ibc.lightclients.beefy.v1.ParachainHeader"></a>

### ParachainHeader
data needed to prove parachain header inclusion in mmr.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `parachain_header` | [bytes](#bytes) |  | scale-encoded parachain header bytes |
| `partial_mmr_leaf` | [PartialMMRLeaf](#ibc.lightclients.beefy.v1.PartialMMRLeaf) |  | see beefy spec |
| `parachain_heads_proof` | [bytes](#bytes) | repeated | proofs for our header in the parachain heads root |
| `heads_leaf_index` | [uint32](#uint32) |  | leaf index for parachain heads proof |
| `heads_total_count` | [uint32](#uint32) |  | total number of para heads in parachain_heads_root |
| `extrinsic_proof` | [bytes](#bytes) | repeated | trie merkle proof of inclusion in header.extrinsic_root |
| `timestamp_extrinsic` | [bytes](#bytes) |  | the actual timestamp extrinsic |






<a name="ibc.lightclients.beefy.v1.ParachainHeadersWithProof"></a>

### ParachainHeadersWithProof
Parachain headers and their mmr proofs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `headers` | [ParachainHeader](#ibc.lightclients.beefy.v1.ParachainHeader) | repeated | parachain headers needed for proofs and ConsensusState |
| `mmr_proofs` | [bytes](#bytes) | repeated | mmr proofs for the headers gotten from rpc "mmr_generateProofs" |
| `mmr_size` | [uint64](#uint64) |  | size of the mmr for the given proof |






<a name="ibc.lightclients.beefy.v1.PartialMMRLeaf"></a>

### PartialMMRLeaf
Partial data for MMRLeaf


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [uint32](#uint32) |  | leaf version |
| `parent_number` | [uint32](#uint32) |  | parent block for this leaf |
| `parent_hash` | [bytes](#bytes) |  | parent hash for this leaf |
| `beefy_next_authority_set` | [BeefyAuthoritySet](#ibc.lightclients.beefy.v1.BeefyAuthoritySet) |  | next authority set. |






<a name="ibc.lightclients.beefy.v1.Payload"></a>

### Payload
Actual payload items


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payload_id` | [bytes](#bytes) |  | 2-byte payload id |
| `payload_data` | [bytes](#bytes) |  | arbitrary length payload data., eg mmr_root_hash |






<a name="ibc.lightclients.beefy.v1.SignedCommitment"></a>

### SignedCommitment
signed commitment data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commitment` | [Commitment](#ibc.lightclients.beefy.v1.Commitment) |  | commitment data being signed |
| `signatures` | [CommitmentSignature](#ibc.lightclients.beefy.v1.CommitmentSignature) | repeated | gotten from rpc subscription |





 <!-- end messages -->


<a name="ibc.lightclients.beefy.v1.RelayChain"></a>

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

