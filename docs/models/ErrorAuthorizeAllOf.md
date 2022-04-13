# ErrorAuthorizeAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** |  | [optional] |
|**Messages** | Pointer to [**[]ErrorAuthorizeMessage**](ErrorAuthorizeMessage.md) |  | [optional] |

## Methods


### GetHttpStatus

`func (o *ErrorAuthorizeAllOf) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ErrorAuthorizeAllOf) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ErrorAuthorizeAllOf) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ErrorAuthorizeAllOf) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ErrorAuthorizeAllOf) GetMessages() []ErrorAuthorizeMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ErrorAuthorizeAllOf) GetMessagesOk() (*[]ErrorAuthorizeMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ErrorAuthorizeAllOf) SetMessages(v []ErrorAuthorizeMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ErrorAuthorizeAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.



