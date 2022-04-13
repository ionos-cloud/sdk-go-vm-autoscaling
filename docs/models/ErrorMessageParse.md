# ErrorMessageParse



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ErrorCode** | Pointer to **int** |  | [optional] |
|**Message** | Pointer to **string** |  | [optional] |

## Methods


### GetErrorCode

`func (o *ErrorMessageParse) GetErrorCode() int`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ErrorMessageParse) GetErrorCodeOk() (*int, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ErrorMessageParse) SetErrorCode(v int)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ErrorMessageParse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorMessageParse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMessageParse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMessageParse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorMessageParse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.



