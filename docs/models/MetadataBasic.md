# MetadataBasic

Metadata about the resource


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | [**time.Time**](time.Time.md) | When the resource was created. | |
|**Etag** | **string** | Resource etag | |
|**LastModifiedDate** | [**time.Time**](time.Time.md) | When the resource was last modified. | |
|**State** | [**MetadataState**](MetadataState.md) |  | |

## Methods


### GetCreatedDate

`func (o *MetadataBasic) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataBasic) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataBasic) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.


### GetEtag

`func (o *MetadataBasic) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *MetadataBasic) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *MetadataBasic) SetEtag(v string)`

SetEtag sets Etag field to given value.


### GetLastModifiedDate

`func (o *MetadataBasic) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataBasic) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataBasic) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.


### GetState

`func (o *MetadataBasic) GetState() MetadataState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataBasic) GetStateOk() (*MetadataState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataBasic) SetState(v MetadataState)`

SetState sets State field to given value.




