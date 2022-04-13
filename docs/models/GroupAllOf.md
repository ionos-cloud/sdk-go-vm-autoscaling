# GroupAllOf



## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Entities** | Pointer to [**GroupEntities**](GroupEntities.md) |  | [optional] |
|**Properties** | Pointer to [**GroupProperties**](GroupProperties.md) |  | [optional] |

## Methods


### GetHref

`func (o *GroupAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GroupAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GroupAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GroupAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *GroupAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GroupAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEntities

`func (o *GroupAllOf) GetEntities() GroupEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *GroupAllOf) GetEntitiesOk() (*GroupEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *GroupAllOf) SetEntities(v GroupEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *GroupAllOf) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetProperties

`func (o *GroupAllOf) GetProperties() GroupProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GroupAllOf) GetPropertiesOk() (*GroupProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GroupAllOf) SetProperties(v GroupProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GroupAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



