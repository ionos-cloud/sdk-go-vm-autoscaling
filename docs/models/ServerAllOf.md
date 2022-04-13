# ServerAllOf

Links a data center server to an autoscaling group. Please note that this entities UUID is different from that of the data center server, whose UUID is stored in the `datacenterServer` property.


## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | URL to the object representation (absolute path) | [optional] [readonly] |
|**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] |
|**Type** | Pointer to **string** | The type of object that has been created | [optional] [readonly] |
|**Properties** | Pointer to [**ServerProperties**](ServerProperties.md) |  | [optional] |

## Methods


### GetHref

`func (o *ServerAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServerAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServerAllOf) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServerAllOf) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *ServerAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ServerAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProperties

`func (o *ServerAllOf) GetProperties() ServerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ServerAllOf) GetPropertiesOk() (*ServerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ServerAllOf) SetProperties(v ServerProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ServerAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.



