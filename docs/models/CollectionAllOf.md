# CollectionAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Items** | Pointer to [**[]Resource**](Resource.md) |  | [optional] |

## Methods


### GetType

`func (o *CollectionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CollectionAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItems

`func (o *CollectionAllOf) GetItems() []Resource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CollectionAllOf) GetItemsOk() (*[]Resource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CollectionAllOf) SetItems(v []Resource)`

SetItems sets Items field to given value.

### HasItems

`func (o *CollectionAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.



