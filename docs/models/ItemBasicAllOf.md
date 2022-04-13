# ItemBasicAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**MetadataBasic**](MetadataBasic.md) |  | [optional] |
|**Properties** | Pointer to **map[string]interface{}** | Properties of the resource. Contents depend on the resource type. | [optional] |

## Methods


### GetMetadata

`func (o *ItemBasicAllOf) GetMetadata() MetadataBasic`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ItemBasicAllOf) GetMetadataOk() (*MetadataBasic, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ItemBasicAllOf) SetMetadata(v MetadataBasic)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ItemBasicAllOf) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ItemBasicAllOf) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ItemBasicAllOf) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ItemBasicAllOf) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ItemBasicAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



