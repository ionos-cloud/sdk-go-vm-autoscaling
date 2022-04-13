# Resource



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | Unique resource identifier | [readonly] |
|**Type** | Pointer to **string** | Resource type | [optional] [readonly] |
|**Href** | Pointer to **string** | Absolute URL to the resource&#39;s representation. | [optional] [readonly] |

## Methods


### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Resource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Resource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *Resource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Resource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Resource) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Resource) HasHref() bool`

HasHref returns a boolean if a field has been set.



