# ParseErrorAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int** |  | [optional] |
|**Messages** | Pointer to [**[]ErrorMessageParse**](ErrorMessageParse.md) |  | [optional] |

## Methods


### GetHttpStatus

`func (o *ParseErrorAllOf) GetHttpStatus() int`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ParseErrorAllOf) GetHttpStatusOk() (*int, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ParseErrorAllOf) SetHttpStatus(v int)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ParseErrorAllOf) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessages

`func (o *ParseErrorAllOf) GetMessages() []ErrorMessageParse`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ParseErrorAllOf) GetMessagesOk() (*[]ErrorMessageParse, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ParseErrorAllOf) SetMessages(v []ErrorMessageParse)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ParseErrorAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.



