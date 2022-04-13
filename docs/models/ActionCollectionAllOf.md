# ActionCollectionAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Items** | Pointer to [**[]Action**](Action.md) |  | [optional] |

## Methods


### GetHref

`func (o *ActionCollectionAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ActionCollectionAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ActionCollectionAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ActionCollectionAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ActionCollectionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionCollectionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionCollectionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionCollectionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ActionCollectionAllOf) GetItems() []Action`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ActionCollectionAllOf) GetItemsOk() (*[]Action, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ActionCollectionAllOf) SetItems(v []Action)`

SetItems sets Items field to given value.

### HasItems

`func (o *ActionCollectionAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.



