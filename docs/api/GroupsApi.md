# \GroupsApi

All URIs are relative to *https://api.ionos.com/cloudapi/autoscaling*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GroupsActionsFindById**](GroupsApi.md#GroupsActionsFindById) | **Get** /groups/{groupId}/actions/{actionId} | Get Scaling Action Details by ID|
|[**GroupsActionsGet**](GroupsApi.md#GroupsActionsGet) | **Get** /groups/{groupId}/actions | Get Scaling Actions|
|[**GroupsDelete**](GroupsApi.md#GroupsDelete) | **Delete** /groups/{groupId} | Delete an Auto Scaling Group by ID|
|[**GroupsFindById**](GroupsApi.md#GroupsFindById) | **Get** /groups/{groupId} | Get an Auto Scaling by ID|
|[**GroupsGet**](GroupsApi.md#GroupsGet) | **Get** /groups | Get Auto Scaling Groups|
|[**GroupsPost**](GroupsApi.md#GroupsPost) | **Post** /groups | Create an Auto Scaling Group|
|[**GroupsPut**](GroupsApi.md#GroupsPut) | **Put** /groups/{groupId} | Update an Auto Scaling Group by ID|
|[**GroupsServersFindById**](GroupsApi.md#GroupsServersFindById) | **Get** /groups/{groupId}/servers/{serverId} | Get Auto Scaling Group Server by ID|
|[**GroupsServersGet**](GroupsApi.md#GroupsServersGet) | **Get** /groups/{groupId}/servers | Get Auto Scaling Group Servers|



## GroupsActionsFindById

```go
var result Action = GroupsActionsFindById(ctx, actionId, groupId)
                      .Depth(depth)
                      .Execute()
```

Get Scaling Action Details by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    actionId := TODO // string | 
    groupId := "groupId_example" // string | 
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsActionsFindById(context.Background(), actionId, groupId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsActionsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsActionsFindById`: Action
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsActionsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**actionId** | [**string**](../models/.md) |  | |
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsActionsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |

### Return type

[**Action**](../models/Action.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsActionsGet

```go
var result ActionCollection = GroupsActionsGet(ctx, groupId)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

Get Scaling Actions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := "groupId_example" // string | 
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)
    orderBy := "orderBy_example" // string | Use this parameter to specify by which the returned list should be sorted. Valid values are: ``createdDate`` and ``lastModifiedDate``. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsActionsGet(context.Background(), groupId).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsActionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsActionsGet`: ActionCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsActionsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsActionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |
| **orderBy** | **string** | Use this parameter to specify by which the returned list should be sorted. Valid values are: &#x60;&#x60;createdDate&#x60;&#x60; and &#x60;&#x60;lastModifiedDate&#x60;&#x60;. | [default to &quot;createdDate&quot;]|

### Return type

[**ActionCollection**](../models/ActionCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsDelete

```go
var result  = GroupsDelete(ctx, groupId)
                      .Execute()
```

Delete an Auto Scaling Group by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := TODO // string | 

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsDelete(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | [**string**](../models/.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsFindById

```go
var result Group = GroupsFindById(ctx, groupId)
                      .Depth(depth)
                      .Execute()
```

Get an Auto Scaling by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := "groupId_example" // string | 
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsFindById(context.Background(), groupId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsFindById`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |

### Return type

[**Group**](../models/Group.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsGet

```go
var result GroupCollection = GroupsGet(ctx)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

Get Auto Scaling Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)
    orderBy := "orderBy_example" // string | Use this parameter to specify by which the returned list should be sorted. Valid values are: ``createdDate`` and ``lastModifiedDate``. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsGet(context.Background()).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsGet`: GroupCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiGroupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |
| **orderBy** | **string** | Use this parameter to specify by which the returned list should be sorted. Valid values are: &#x60;&#x60;createdDate&#x60;&#x60; and &#x60;&#x60;lastModifiedDate&#x60;&#x60;. | [default to &quot;createdDate&quot;]|

### Return type

[**GroupCollection**](../models/GroupCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsPost

```go
var result GroupPostResponse = GroupsPost(ctx)
                      .GroupPost(groupPost)
                      .Execute()
```

Create an Auto Scaling Group



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupPost := *openapiclient.NewGroupPost("1d67ca27-d4c0-419d-9a64-9ea42dfdd036", *openapiclient.NewGroupProperties("de/txl")) // GroupPost | 

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsPost(context.Background()).GroupPost(groupPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsPost`: GroupPostResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiGroupsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **groupPost** | [**GroupPost**](../models/GroupPost.md) |  | |

### Return type

[**GroupPostResponse**](../models/GroupPostResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## GroupsPut

```go
var result Group = GroupsPut(ctx, groupId)
                      .GroupPut(groupPut)
                      .Execute()
```

Update an Auto Scaling Group by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := TODO // string | 
    groupPut := *openapiclient.NewGroupPut(*openapiclient.NewGroupPutProperties("de/txl", int64(10), int64(1), "Autoscaling Group #1", *openapiclient.NewGroupPolicy(openapiclient.Metric("INSTANCE_CPU_UTILIZATION_AVERAGE"), *openapiclient.NewGroupPolicyScaleInAction(float32(1), openapiclient.ActionAmount("ABSOLUTE"), true), float32(33), *openapiclient.NewGroupPolicyScaleOutAction(float32(1), openapiclient.ActionAmount("ABSOLUTE")), float32(77), openapiclient.QueryUnit("PER_HOUR")), *openapiclient.NewReplicaPropertiesPost(int32(2), int32(2048)))) // GroupPut | 

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsPut(context.Background(), groupId).GroupPut(groupPut).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsPut`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | [**string**](../models/.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **groupPut** | [**GroupPut**](../models/GroupPut.md) |  | |

### Return type

[**Group**](../models/Group.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## GroupsServersFindById

```go
var result Server = GroupsServersFindById(ctx, groupId, serverId)
                      .Depth(depth)
                      .Execute()
```

Get Auto Scaling Group Server by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := "groupId_example" // string | 
    serverId := TODO // string | 
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsServersFindById(context.Background(), groupId, serverId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsServersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsServersFindById`: Server
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsServersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |
|**serverId** | [**string**](../models/.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsServersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |

### Return type

[**Server**](../models/Server.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GroupsServersGet

```go
var result ServerCollection = GroupsServersGet(ctx, groupId)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

Get Auto Scaling Group Servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud_vm_autoscaling "github.com/ionos-cloud/sdk-go-vm-autoscaling"
)

func main() {
    groupId := "groupId_example" // string | 
    depth := float32(8.14) // float32 | With this parameter, you control the level of detail of the response objects:    - ``0``: Only direct properties are included; children (such as executions or transitions) are not considered.    - ``1``: Direct properties and children references are included.    - ``2``: Direct properties and children properties are included.    - ``3``: Direct properties and children properties and children's children are included.    - etc.   (optional)
    orderBy := "orderBy_example" // string | Use this parameter to specify by which the returned list should be sorted. Valid values are: ``createdDate`` and ``lastModifiedDate``. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.GroupsServersGet(context.Background(), groupId).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GroupsServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GroupsServersGet`: ServerCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GroupsServersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGroupsServersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | With this parameter, you control the level of detail of the response objects:    - &#x60;&#x60;0&#x60;&#x60;: Only direct properties are included; children (such as executions or transitions) are not considered.    - &#x60;&#x60;1&#x60;&#x60;: Direct properties and children references are included.    - &#x60;&#x60;2&#x60;&#x60;: Direct properties and children properties are included.    - &#x60;&#x60;3&#x60;&#x60;: Direct properties and children properties and children&#39;s children are included.    - etc.   | |
| **orderBy** | **string** | Use this parameter to specify by which the returned list should be sorted. Valid values are: &#x60;&#x60;createdDate&#x60;&#x60; and &#x60;&#x60;lastModifiedDate&#x60;&#x60;. | [default to &quot;createdDate&quot;]|

### Return type

[**ServerCollection**](../models/ServerCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


