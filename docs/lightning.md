# Protocol Documentation
<a name="top"></a>

<!--
## Table of Contents

- [proto/imp/api/lightning/lightning.proto](#proto/imp/api/lightning/lightning.proto)
    - [CheckInvoiceRequest](#lightning.CheckInvoiceRequest)
    - [CheckInvoiceResponse](#lightning.CheckInvoiceResponse)
    - [GenerateInvoiceRequest](#lightning.GenerateInvoiceRequest)
    - [GenerateInvoiceResponse](#lightning.GenerateInvoiceResponse)
    - [PayInvoiceRequest](#lightning.PayInvoiceRequest)
    - [PayInvoiceResponse](#lightning.PayInvoiceResponse)

    - [Lightning](#lightning.Lightning)

- [Scalar Value Types](#scalar-value-types)



<a name="proto/imp/api/lightning/lightning.proto"></a>
<p align="right"><a href="#top">Top</a></p>

-->

## proto/imp/api/lightning/lightning.proto
Allows for p2p messaging between Impervious nodes



<a name="lightning.Lightning"></a>

### Lightning
Lightning service allows lightning actions on your underlying lightning node from the Impervious node.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateInvoice | GenerateInvoiceRequest | GenerateInvoiceResponse | GenerateInvoice allows you to generate an invoice for a specific payment amount from your lightning node. |
| PayInvoice | PayInvoiceRequest | PayInvoiceResponse | PayInvoice allows you to pay a specific invoice with your lightning node. |
| CheckInvoice | CheckInvoiceRequest | CheckInvoiceResponse | CheckInvoice allows you to check a specific invoice to see if it was paid. |


#### HTTP bindings

| Method Name | Method | Pattern |
| ----------- | ------ | ------- |
| `GenerateInvoice` | `POST` | `/v1/lightning/generateinvoice`
| `PayInvoice` | `POST` | `/v1/lightning/payinvoice`
| `CheckInvoice` | `POST` | `/v1/lightning/checkinvoice` <!-- end services -->



<a name="lightning.CheckInvoiceRequest"></a>

### CheckInvoiceRequest
Represents an request to check an invoice.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invoice | string |  | The invoice to check |






<a name="lightning.CheckInvoiceResponse"></a>

### CheckInvoiceResponse
Represents a response back from the invoice check.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paid | bool |  | The boolean result representing whether or not the invoice was paid |






<a name="lightning.GenerateInvoiceRequest"></a>

### GenerateInvoiceRequest
Represents an invoice creation request from your lightning node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| amount | int64 |  | The amount of satoshis you want to receive |
| memo | string |  | The human readable memo you want the sender to see |






<a name="lightning.GenerateInvoiceResponse"></a>

### GenerateInvoiceResponse
Represents a response back from an invoice generation request.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invoice | string |  | The invoice from your lightning node |






<a name="lightning.PayInvoiceRequest"></a>

### PayInvoiceRequest
Represents an invoice that will be paid by your lightning node.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invoice | string |  | The invoice to pay |






<a name="lightning.PayInvoiceResponse"></a>

### PayInvoiceResponse
Represents a response back from the payment result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| preimage | string |  | The preimage from the payment result, if successful |





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
