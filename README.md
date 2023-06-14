# Go API client for ionoscloud_vm_autoscaling 

VM Auto Scaling service enables IONOS clients to horizontally scale the number of VM instances, based on configured rules. Use Auto Scaling to ensure you will have a sufficient number of instances to handle your application loads at all times.

Create an Auto Scaling group that contains the server instances; Auto Scaling service will ensure that the number of instances in the group is always within these limits.

When target replica count is specified, Auto Scaling will maintain the set number on instances.

When scaling policies are specified, Auto Scaling will create or delete instances based on the demands of your applications. For each policy, specified scale-in and scale-out actions are performed whenever the corresponding thresholds are met.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.ionos.com/faq/contact](https://docs.ionos.com/faq/contact)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./ionoscloud_vm_autoscaling"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DocumentationApi* | [**AutoscalingOpenapiJsonGet**](docs/api/DocumentationApi.md#autoscalingopenapijsonget) | **Get** /cloudapi/autoscaling/openapi.json | Retrieve VM autoscaling OpenAPI spec (JSON)
*DocumentationApi* | [**AutoscalingOpenapiYamlGet**](docs/api/DocumentationApi.md#autoscalingopenapiyamlget) | **Get** /cloudapi/autoscaling/openapi.yaml | Retrieve VM autoscaling OpenAPI spec (YAML)
*GroupsApi* | [**AutoscalingGroupsActionsFindById**](docs/api/GroupsApi.md#autoscalinggroupsactionsfindbyid) | **Get** /cloudapi/autoscaling/groups/{groupId}/actions/{actionId} | Retrieve action details
*GroupsApi* | [**AutoscalingGroupsActionsGet**](docs/api/GroupsApi.md#autoscalinggroupsactionsget) | **Get** /cloudapi/autoscaling/groups/{groupId}/actions | Retrieve last ten actions
*GroupsApi* | [**AutoscalingGroupsDelete**](docs/api/GroupsApi.md#autoscalinggroupsdelete) | **Delete** /cloudapi/autoscaling/groups/{groupId} | Delete autoscaling groups.
*GroupsApi* | [**AutoscalingGroupsFindById**](docs/api/GroupsApi.md#autoscalinggroupsfindbyid) | **Get** /cloudapi/autoscaling/groups/{groupId} | Retrieve autoscaling groups by UUID
*GroupsApi* | [**AutoscalingGroupsGet**](docs/api/GroupsApi.md#autoscalinggroupsget) | **Get** /cloudapi/autoscaling/groups | List autoscaling groups
*GroupsApi* | [**AutoscalingGroupsPost**](docs/api/GroupsApi.md#autoscalinggroupspost) | **Post** /cloudapi/autoscaling/groups | Create autoscaling groups
*GroupsApi* | [**AutoscalingGroupsPut**](docs/api/GroupsApi.md#autoscalinggroupsput) | **Put** /cloudapi/autoscaling/groups/{groupId} | Update autoscaling groups
*GroupsApi* | [**AutoscalingGroupsServersFindById**](docs/api/GroupsApi.md#autoscalinggroupsserversfindbyid) | **Get** /cloudapi/autoscaling/groups/{groupId}/servers/{serverId} | Retrieve group servers by UUID
*GroupsApi* | [**AutoscalingGroupsServersGet**](docs/api/GroupsApi.md#autoscalinggroupsserversget) | **Get** /cloudapi/autoscaling/groups/{groupId}/servers | Retrieve autoscaling group servers


## Documentation For Models

 - [Action](docs/models/Action.md)
 - [ActionAllOf](docs/models/ActionAllOf.md)
 - [ActionAmount](docs/models/ActionAmount.md)
 - [ActionCollection](docs/models/ActionCollection.md)
 - [ActionCollectionAllOf](docs/models/ActionCollectionAllOf.md)
 - [ActionProperties](docs/models/ActionProperties.md)
 - [ActionResource](docs/models/ActionResource.md)
 - [ActionResourceAllOf](docs/models/ActionResourceAllOf.md)
 - [ActionStatus](docs/models/ActionStatus.md)
 - [ActionType](docs/models/ActionType.md)
 - [ActionsLinkResource](docs/models/ActionsLinkResource.md)
 - [ActionsLinkResourceAllOf](docs/models/ActionsLinkResourceAllOf.md)
 - [AvailabilityZone](docs/models/AvailabilityZone.md)
 - [BusType](docs/models/BusType.md)
 - [Collection](docs/models/Collection.md)
 - [CollectionAllOf](docs/models/CollectionAllOf.md)
 - [CpuFamily](docs/models/CpuFamily.md)
 - [DatacenterServer](docs/models/DatacenterServer.md)
 - [DatacenterServerAllOf](docs/models/DatacenterServerAllOf.md)
 - [Error](docs/models/Error.md)
 - [Error401](docs/models/Error401.md)
 - [Error401AllOf](docs/models/Error401AllOf.md)
 - [Error401Message](docs/models/Error401Message.md)
 - [Error401MessageAllOf](docs/models/Error401MessageAllOf.md)
 - [Error404](docs/models/Error404.md)
 - [Error404AllOf](docs/models/Error404AllOf.md)
 - [Error404Message](docs/models/Error404Message.md)
 - [Error404MessageAllOf](docs/models/Error404MessageAllOf.md)
 - [ErrorAuthorize](docs/models/ErrorAuthorize.md)
 - [ErrorAuthorizeAllOf](docs/models/ErrorAuthorizeAllOf.md)
 - [ErrorAuthorizeMessage](docs/models/ErrorAuthorizeMessage.md)
 - [ErrorAuthorizeMessageAllOf](docs/models/ErrorAuthorizeMessageAllOf.md)
 - [ErrorGroupValidate](docs/models/ErrorGroupValidate.md)
 - [ErrorGroupValidateAllOf](docs/models/ErrorGroupValidateAllOf.md)
 - [ErrorGroupValidateMessage](docs/models/ErrorGroupValidateMessage.md)
 - [ErrorGroupValidateMessageAllOf](docs/models/ErrorGroupValidateMessageAllOf.md)
 - [ErrorMessage](docs/models/ErrorMessage.md)
 - [ErrorMessageParse](docs/models/ErrorMessageParse.md)
 - [ErrorMessageParseAllOf](docs/models/ErrorMessageParseAllOf.md)
 - [ErrorReplicaValidateMessage](docs/models/ErrorReplicaValidateMessage.md)
 - [ErrorReplicaValidateMessageAllOf](docs/models/ErrorReplicaValidateMessageAllOf.md)
 - [ErrorUuid](docs/models/ErrorUuid.md)
 - [Group](docs/models/Group.md)
 - [GroupAllOf](docs/models/GroupAllOf.md)
 - [GroupCollection](docs/models/GroupCollection.md)
 - [GroupCollectionAllOf](docs/models/GroupCollectionAllOf.md)
 - [GroupEntities](docs/models/GroupEntities.md)
 - [GroupImmutableProperties](docs/models/GroupImmutableProperties.md)
 - [GroupPolicy](docs/models/GroupPolicy.md)
 - [GroupPolicyScaleInAction](docs/models/GroupPolicyScaleInAction.md)
 - [GroupPolicyScaleOutAction](docs/models/GroupPolicyScaleOutAction.md)
 - [GroupPostResponse](docs/models/GroupPostResponse.md)
 - [GroupProperties](docs/models/GroupProperties.md)
 - [GroupResource](docs/models/GroupResource.md)
 - [GroupResourceAllOf](docs/models/GroupResourceAllOf.md)
 - [GroupUpdatableProperties](docs/models/GroupUpdatableProperties.md)
 - [GroupUpdate](docs/models/GroupUpdate.md)
 - [Item](docs/models/Item.md)
 - [ItemAllOf](docs/models/ItemAllOf.md)
 - [ItemBasic](docs/models/ItemBasic.md)
 - [ItemBasicAllOf](docs/models/ItemBasicAllOf.md)
 - [Metadata](docs/models/Metadata.md)
 - [MetadataBasic](docs/models/MetadataBasic.md)
 - [MetadataState](docs/models/MetadataState.md)
 - [Metric](docs/models/Metric.md)
 - [ParseError](docs/models/ParseError.md)
 - [ParseErrorAllOf](docs/models/ParseErrorAllOf.md)
 - [QueryUnit](docs/models/QueryUnit.md)
 - [ReplicaNic](docs/models/ReplicaNic.md)
 - [ReplicaProperties](docs/models/ReplicaProperties.md)
 - [ReplicaPropertiesGet](docs/models/ReplicaPropertiesGet.md)
 - [ReplicaPropertiesGetAllOf](docs/models/ReplicaPropertiesGetAllOf.md)
 - [ReplicaPropertiesPost](docs/models/ReplicaPropertiesPost.md)
 - [ReplicaPropertiesPostAllOf](docs/models/ReplicaPropertiesPostAllOf.md)
 - [ReplicaVolumeGet](docs/models/ReplicaVolumeGet.md)
 - [ReplicaVolumePost](docs/models/ReplicaVolumePost.md)
 - [ReplicaVolumePostAllOf](docs/models/ReplicaVolumePostAllOf.md)
 - [Resource](docs/models/Resource.md)
 - [Server](docs/models/Server.md)
 - [ServerAllOf](docs/models/ServerAllOf.md)
 - [ServerCollection](docs/models/ServerCollection.md)
 - [ServerCollectionAllOf](docs/models/ServerCollectionAllOf.md)
 - [ServerProperties](docs/models/ServerProperties.md)
 - [ServerResource](docs/models/ServerResource.md)
 - [ServerResourceAllOf](docs/models/ServerResourceAllOf.md)
 - [ServersLinkResource](docs/models/ServersLinkResource.md)
 - [ServersLinkResourceAllOf](docs/models/ServersLinkResourceAllOf.md)
 - [StartedActions](docs/models/StartedActions.md)
 - [TerminationPolicyType](docs/models/TerminationPolicyType.md)
 - [VolumeHwType](docs/models/VolumeHwType.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@cloud.ionos.com

