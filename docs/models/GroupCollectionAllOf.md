# GroupCollectionAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Items** | Pointer to [**[]Group**](Group.md) |  | [optional] |

## Methods


### GetHref

`func (o *GroupCollectionAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupCollectionAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupCollectionAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupCollectionAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *GroupCollectionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupCollectionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupCollectionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupCollectionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *GroupCollectionAllOf) GetItems() []Group`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GroupCollectionAllOf) GetItemsOk() (*[]Group, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GroupCollectionAllOf) SetItems(v []Group)`

SetItems sets Items field to given value.

### HasItems

`func (o *GroupCollectionAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.



