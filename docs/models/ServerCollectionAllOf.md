# ServerCollectionAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Items** | Pointer to [**[]Server**](Server.md) |  | [optional] |

## Methods


### GetHref

`func (o *ServerCollectionAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServerCollectionAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServerCollectionAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServerCollectionAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ServerCollectionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerCollectionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerCollectionAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerCollectionAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ServerCollectionAllOf) GetItems() []Server`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServerCollectionAllOf) GetItemsOk() (*[]Server, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServerCollectionAllOf) SetItems(v []Server)`

SetItems sets Items field to given value.

### HasItems

`func (o *ServerCollectionAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.



