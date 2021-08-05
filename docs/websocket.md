# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/websocket/websocket.proto](#proto/imp/api/websocket/websocket.proto)
    - [SubscribeRequest](#websocket.SubscribeRequest)
    - [SubscribeResponse](#websocket.SubscribeResponse)

    - [Websocket](#websocket.Websocket)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/websocket/websocket.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/websocket/websocket.proto
Allows for receiving messages from your IMP node



<a name="websocket.Websocket"></a>

### Websocket
Websocket service allows for receiving messages received into your Impervious node.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Subscribe | SubscribeRequest | SubscribeResponse stream | Subscribe opens up a stream/websocket to receive all messages received on your Impervious node. |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `Subscribe` | `GET` | `/v1/subscribe` <!-- end services -->



<a name="websocket.SubscribeRequest"></a>

### SubscribeRequest
Represents a request to subscribe to the event websocket






<a name="websocket.SubscribeResponse"></a>

### SubscribeResponse
Represents a response back from the websocket containing event information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | The ID of the message |
| reply_to_id | string |  | Optional ID of the message the sender is replying to |
| from_pubkey | string |  | The node that sent the message |
| data | string |  | The data the node is sending over |
| service_type | string |  | The type of service the message is meant for |
| amount | int64 |  | The amount sent as part of the message |





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
