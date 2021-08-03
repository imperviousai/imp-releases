# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/signing/signing.proto](#proto/imp/api/signing/signing.proto)
    - [SignRequest](#signing.SignRequest)
    - [SignResponse](#signing.SignResponse)
    - [VerifyRequest](#signing.VerifyRequest)
    - [VerifyResponse](#signing.VerifyResponse)

    - [Signing](#signing.Signing)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/signing/signing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/signing/signing.proto
Allows an Imp node to sign and verify messaging with the connected LND



<a name="signing.Signing"></a>

### Signing


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SignMessage | SignRequest | SignResponse |  |
| VerifySignature | VerifyRequest | VerifyResponse |  |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `SignMessage` | `POST` | `/v1/sign`
| `VerifySignature` | `POST` | `/v1/verify` <!-- end services -->



<a name="signing.SignRequest"></a>

### SignRequest
Represents a request to sign a message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | string |  | message to be signed |






<a name="signing.SignResponse"></a>

### SignResponse
Represents a response from a signature request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| signature | string |  | signature of signed message returned from LND |






<a name="signing.VerifyRequest"></a>

### VerifyRequest
Represents a request to verify a signature and message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| msg | string |  | message to be verified |
| signature | string |  | signature of message |






<a name="signing.VerifyResponse"></a>

### VerifyResponse
Represents a response back from a verification request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | bool |  | result of signature verification |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

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
