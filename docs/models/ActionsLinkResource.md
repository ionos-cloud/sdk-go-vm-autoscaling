# ActionsLinkResource



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The resource&#39;s unique identifier | [readonly] |
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |

## Methods


### GetId

`func (o *ActionsLinkResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionsLinkResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionsLinkResource) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ActionsLinkResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionsLinkResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionsLinkResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionsLinkResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ActionsLinkResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ActionsLinkResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ActionsLinkResource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ActionsLinkResource) HasHref() bool`

HasHref returns a boolean if a field has been set.



