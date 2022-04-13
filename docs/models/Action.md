# Action

A background action associated with an autoscaling group.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The resource&#39;s unique identifier | [readonly] |
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Metadata** | Pointer to [**MetadataBasic**](MetadataBasic.md) |  | [optional] |
|**Properties** | Pointer to [**ActionProperties**](ActionProperties.md) |  | [optional] |

## Methods


### GetId

`func (o *Action) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Action) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Action) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Action) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Action) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Action) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Action) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Action) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Action) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Action) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Action) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *Action) GetMetadata() MetadataBasic`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Action) GetMetadataOk() (*MetadataBasic, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Action) SetMetadata(v MetadataBasic)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Action) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Action) GetProperties() ActionProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Action) GetPropertiesOk() (*ActionProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Action) SetProperties(v ActionProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Action) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



