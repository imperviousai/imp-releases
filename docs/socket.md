# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/socket/socket.proto](#proto/imp/api/socket/socket.proto)
    - [SendSocketRequest](#socket.SendSocketRequest)
    - [SendSocketResponse](#socket.SendSocketResponse)
    - [StartSocketRequest](#socket.StartSocketRequest)
    - [StartSocketResponse](#socket.StartSocketResponse)
    - [StopSocketRequest](#socket.StopSocketRequest)
    - [StopSocketResponse](#socket.StopSocketResponse)

    - [Socket](#socket.Socket)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/socket/socket.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/socket/socket.proto
Allows for p2p sockets to be established between Impervious nodes



<a name="socket.Socket"></a>

### Socket
Socket service allows for p2p sockets to be established between Impervious nodes.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendSocket | SendSocketRequest | SendSocketResponse | SendSocket sends a socket request to another node. |
| StartSocket | StartSocketRequest | StartSocketResponse | StartSocket starts the socket on your Impervious node. |
| StopSocket | StopSocketRequest | StopSocketResponse | StopSocket stops the socket on your Impervious node. |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `SendSocket` | `POST` | `/v1/socket/sendRequest`
| `StartSocket` | `POST` | `/v1/socket/start`
| `StopSocket` | `POST` | `/v1/socket/stop` <!-- end services -->



<a name="socket.SendSocketRequest"></a>

### SendSocketRequest
Represents a request to send socket connection information to a far end node


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pubkey | string |  | The public key of the far end LND node running IMP |






<a name="socket.SendSocketResponse"></a>

### SendSocketResponse
Represents a response back from the request to establish a socket


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | string |  | returned message ID |






<a name="socket.StartSocketRequest"></a>

### StartSocketRequest
Represents a request to start a socket on an owned IMP node






<a name="socket.StartSocketResponse"></a>

### StartSocketResponse
Represents a response containing socket information from a started socket


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protocol | string |  | the protocol of the socket connection |
| ip | string |  | the ip of the IMP node to connect to |
| port1 | string |  | the first port of the socket connection |
| port2 | string |  | the second port of the socket connection |






<a name="socket.StopSocketRequest"></a>

### StopSocketRequest
Represents a request to stop a socket on an owned IMP node






<a name="socket.StopSocketResponse"></a>

### StopSocketResponse
Represents a response back from a stopped socket





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
