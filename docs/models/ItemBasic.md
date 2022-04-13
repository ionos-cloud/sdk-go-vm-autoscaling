# ItemBasic



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | Unique resource identifier | [readonly] |
|**Type** | Pointer to **string** | Resource type | [optional] [readonly] |
|**Href** | Pointer to **string** | Absolute URL to the resource&#39;s representation. | [optional] [readonly] |
|**Metadata** | Pointer to [**MetadataBasic**](MetadataBasic.md) |  | [optional] |
|**Properties** | Pointer to **map[string]interface{}** | Properties of the resource. Contents depend on the resource type. | [optional] |

## Methods


### GetId

`func (o *ItemBasic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ItemBasic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ItemBasic) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ItemBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemBasic) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ItemBasic) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ItemBasic) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ItemBasic) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ItemBasic) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ItemBasic) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *ItemBasic) GetMetadata() MetadataBasic`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemBasic) GetMetadataOk() (*MetadataBasic, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemBasic) SetMetadata(v MetadataBasic)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemBasic) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ItemBasic) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ItemBasic) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ItemBasic) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ItemBasic) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



