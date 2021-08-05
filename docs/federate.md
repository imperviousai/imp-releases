# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/federate/federate.proto](#proto/imp/api/federate/federate.proto)
    - [LeaveFederationRequest](#federate.LeaveFederationRequest)
    - [LeaveFederationResponse](#federate.LeaveFederationResponse)
    - [RequestFederateRequest](#federate.RequestFederateRequest)
    - [RequestFederateResponse](#federate.RequestFederateResponse)

    - [Federate](#federate.Federate)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/federate/federate.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/federate/federate.proto
Allows for p2p federation between Impervious nodes



<a name="federate.Federate"></a>

### Federate
Federate service allows for P2P federation between Impervious nodes.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RequestFederate | RequestFederateRequest | RequestFederateResponse | RequestFederation performs the federation request to a specific peer. |
| LeaveFederation | LeaveFederationRequest | LeaveFederationResponse | LeaveFederation performs the removal of a federated peer (upon message receipt). |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `RequestFederate` | `POST` | `/v1/federate/request`
| `LeaveFederation` | `POST` | `/v1/federate/leave` <!-- end services -->



<a name="federate.LeaveFederationRequest"></a>

### LeaveFederationRequest
Represents a request to leave a federation from a far end node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |






<a name="federate.LeaveFederationResponse"></a>

### LeaveFederationResponse
Represents a response back from a Leave Federation Request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | returned message ID |






<a name="federate.RequestFederateRequest"></a>

### RequestFederateRequest
Represents a request to federate with a far end node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |






<a name="federate.RequestFederateResponse"></a>

### RequestFederateResponse
Represents a response back from a Federation Request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | returned message ID |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

## Scalar Value Types

| .proto Type | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double | double | double | float | float64 | double | float | Float |
| <a name="float" /> float | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |
