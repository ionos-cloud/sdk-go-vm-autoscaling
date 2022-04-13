# ErrorGroupValidateAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** |  | [optional] |
|**Messages** | Pointer to [**[]ErrorGroupValidateMessage**](ErrorGroupValidateMessage.md) |  | [optional] |

## Methods


### GetHttpStatus

`func (o *ErrorGroupValidateAllOf) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ErrorGroupValidateAllOf) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ErrorGroupValidateAllOf) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ErrorGroupValidateAllOf) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorGroupValidateAllOf) GetMessages() []ErrorGroupValidateMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorGroupValidateAllOf) GetMessagesOk() (*[]ErrorGroupValidateMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorGroupValidateAllOf) SetMessages(v []ErrorGroupValidateMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorGroupValidateAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.



