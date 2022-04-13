# ParseError



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int** |  | [optional] |
|**Messages** | Pointer to [**[]ErrorMessageParse**](ErrorMessageParse.md) |  | [optional] |
|**ErrorUuid** | Pointer to **string** |  | [optional] |

## Methods


### GetHttpStatus

`func (o *ParseError) GetHttpStatus() int`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ParseError) GetHttpStatusOk() (*int, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ParseError) SetHttpStatus(v int)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ParseError) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ParseError) GetMessages() []ErrorMessageParse`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ParseError) GetMessagesOk() (*[]ErrorMessageParse, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ParseError) SetMessages(v []ErrorMessageParse)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ParseError) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetErrorUuid

`func (o *ParseError) GetErrorUuid() string`

GetErrorUuid returns the ErrorUuid field if non-nil, zero value otherwise.

### GetErrorUuidOk

`func (o *ParseError) GetErrorUuidOk() (*string, bool)`

GetErrorUuidOk returns a tuple with the ErrorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorUuid

`func (o *ParseError) SetErrorUuid(v string)`

SetErrorUuid sets ErrorUuid field to given value.

### HasErrorUuid

`func (o *ParseError) HasErrorUuid() bool`

HasErrorUuid returns a boolean if a field has been set.



