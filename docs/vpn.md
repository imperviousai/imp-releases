# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/vpn/vpn.proto](#proto/imp/api/vpn/vpn.proto)
    - [AcceptContractRequest](#vpn.AcceptContractRequest)
    - [AcceptContractResponse](#vpn.AcceptContractResponse)
    - [RefreshContractRequest](#vpn.RefreshContractRequest)
    - [RefreshContractResponse](#vpn.RefreshContractResponse)
    - [RequestQuoteRequest](#vpn.RequestQuoteRequest)
    - [RequestQuoteResponse](#vpn.RequestQuoteResponse)

    - [VPN](#vpn.VPN)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/vpn/vpn.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/vpn/vpn.proto
Allows for an Encrypted Wireguard VPN between Impervious nodes



<a name="vpn.VPN"></a>

### VPN
VPN service allows for an Encrypted Wireguard VPN between Impervious nodes.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RequestQuote | RequestQuoteRequest | RequestQuoteResponse | RequestQuote requests a quote from another node. |
| AcceptContract | AcceptContractRequest | AcceptContractResponse | AcceptContract accepts and pays for a specific contract. |
| RefreshContract | RefreshContractRequest | RefreshContractResponse | RefreshContract accepts and pays for a specific contract refresh on an existing VPN. |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `RequestQuote` | `POST` | `/v1/vpn/quote`
| `AcceptContract` | `POST` | `/v1/vpn/contract`
| `RefreshContract` | `POST` | `/v1/vpn/refresh` <!-- end services -->



<a name="vpn.AcceptContractRequest"></a>

### AcceptContractRequest
Represents a request to Accept (Pay For) a VPN Quote


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |
| nonce | string |  | the identifier from a requested VPN quote |
| price | string |  | the agreed upon price from the VPN quote |






<a name="vpn.AcceptContractResponse"></a>

### AcceptContractResponse
Represents a response back from an accepted VPN Quote


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | returned message ID |






<a name="vpn.RefreshContractRequest"></a>

### RefreshContractRequest
Represents a request to extend/refresh an expiring VPN Connection (i.e. purchase more time)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |
| nonce | string |  | the identifier of the VPN connection |
| price | string |  | the agreed upon price from the VPN quote |






<a name="vpn.RefreshContractResponse"></a>

### RefreshContractResponse
Represents a reponse back from a refreshed VPN connection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | returned message ID |






<a name="vpn.RequestQuoteRequest"></a>

### RequestQuoteRequest
Represents a request to receive a VPN quote from a far end node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |






<a name="vpn.RequestQuoteResponse"></a>

### RequestQuoteResponse
Represents a response back from a VPN Quote Reqeust


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
