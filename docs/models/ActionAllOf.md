# ActionAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Metadata** | Pointer to [**MetadataBasic**](MetadataBasic.md) |  | [optional] |
|**Properties** | Pointer to [**ActionProperties**](ActionProperties.md) |  | [optional] |

## Methods


### GetHref

`func (o *ActionAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ActionAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ActionAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ActionAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ActionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ActionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadata

`func (o *ActionAllOf) GetMetadata() MetadataBasic`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ActionAllOf) GetMetadataOk() (*MetadataBasic, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ActionAllOf) SetMetadata(v MetadataBasic)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ActionAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ActionAllOf) GetProperties() ActionProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ActionAllOf) GetPropertiesOk() (*ActionProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ActionAllOf) SetProperties(v ActionProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ActionAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



