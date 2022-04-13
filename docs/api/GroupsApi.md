# \GroupsApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**AutoscalingGroupsActionsFindById**](GroupsApi.md#AutoscalingGroupsActionsFindById) | **Get** /cloudapi/autoscaling/groups/{groupId}/actions/{actionId} | Retrieve action details|
|[**AutoscalingGroupsActionsGet**](GroupsApi.md#AutoscalingGroupsActionsGet) | **Get** /cloudapi/autoscaling/groups/{groupId}/actions | Retrieve last ten actions|
|[**AutoscalingGroupsDelete**](GroupsApi.md#AutoscalingGroupsDelete) | **Delete** /cloudapi/autoscaling/groups/{groupId} | Delete autoscaling groups.|
|[**AutoscalingGroupsFindById**](GroupsApi.md#AutoscalingGroupsFindById) | **Get** /cloudapi/autoscaling/groups/{groupId} | Retrieve autoscaling groups by UUID|
|[**AutoscalingGroupsGet**](GroupsApi.md#AutoscalingGroupsGet) | **Get** /cloudapi/autoscaling/groups | List autoscaling groups|
|[**AutoscalingGroupsPost**](GroupsApi.md#AutoscalingGroupsPost) | **Post** /cloudapi/autoscaling/groups | Create autoscaling groups|
|[**AutoscalingGroupsPut**](GroupsApi.md#AutoscalingGroupsPut) | **Put** /cloudapi/autoscaling/groups/{groupId} | Update autoscaling groups|
|[**AutoscalingGroupsServersFindById**](GroupsApi.md#AutoscalingGroupsServersFindById) | **Get** /cloudapi/autoscaling/groups/{groupId}/servers/{serverId} | Retrieve group servers by UUID|
|[**AutoscalingGroupsServersGet**](GroupsApi.md#AutoscalingGroupsServersGet) | **Get** /cloudapi/autoscaling/groups/{groupId}/servers | Retrieve autoscaling group servers|



## AutoscalingGroupsActionsFindById

```go
var result Action = AutoscalingGroupsActionsFindById(ctx, actionId, groupId)
                      .Depth(depth)
                      .Execute()
```

Retrieve action details



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsActionsFindById(context.Background(), actionId, groupId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsActionsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsActionsFindById`: Action
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsActionsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**actionId** | [**string**](.md) |  | |
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsActionsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |

### Return type

[**Action**](Action.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsActionsGet

```go
var result ActionCollection = AutoscalingGroupsActionsGet(ctx, groupId)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve last ten actions



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)
    orderBy := "orderBy_example" // string | Define the property to be used for ordering the returned list; valid values are 'createdDate' and 'lastModifiedDate'. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsActionsGet(context.Background(), groupId).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsActionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsActionsGet`: ActionCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsActionsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsActionsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |
| **orderBy** | **string** | Define the property to be used for ordering the returned list; valid values are &#39;createdDate&#39; and &#39;lastModifiedDate&#39;. | [default to &quot;createdDate&quot;]|

### Return type

[**ActionCollection**](ActionCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsDelete

```go
var result  = AutoscalingGroupsDelete(ctx, groupId)
                      .Execute()
```

Delete autoscaling groups.



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
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsDelete(context.Background(), groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | [**string**](.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsFindById

```go
var result Group = AutoscalingGroupsFindById(ctx, groupId)
                      .Depth(depth)
                      .Execute()
```

Retrieve autoscaling groups by UUID



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsFindById(context.Background(), groupId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsFindById`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |

### Return type

[**Group**](Group.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsGet

```go
var result GroupCollection = AutoscalingGroupsGet(ctx)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

List autoscaling groups



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)
    orderBy := "orderBy_example" // string | Define the property to be used for ordering the returned list; valid values are 'createdDate' and 'lastModifiedDate'. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsGet(context.Background()).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsGet`: GroupCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |
| **orderBy** | **string** | Define the property to be used for ordering the returned list; valid values are &#39;createdDate&#39; and &#39;lastModifiedDate&#39;. | [default to &quot;createdDate&quot;]|

### Return type

[**GroupCollection**](GroupCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsPost

```go
var result GroupPostResponse = AutoscalingGroupsPost(ctx)
                      .Group(group)
                      .Execute()
```

Create autoscaling groups



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
    group := *openapiclient.NewGroup("1d67ca27-d4c0-419d-9a64-9ea42dfdd036", *openapiclient.NewGroupProperties("de/txl")) // Group | 

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsPost(context.Background()).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsPost`: GroupPostResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **group** | [**Group**](Group.md) |  | |

### Return type

[**GroupPostResponse**](GroupPostResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## AutoscalingGroupsPut

```go
var result Group = AutoscalingGroupsPut(ctx, groupId)
                      .GroupUpdate(groupUpdate)
                      .Execute()
```

Update autoscaling groups



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
    groupUpdate := *openapiclient.NewGroupUpdate(*openapiclient.NewGroupUpdatableProperties()) // GroupUpdate | 

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsPut(context.Background(), groupId).GroupUpdate(groupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsPut`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | [**string**](.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **groupUpdate** | [**GroupUpdate**](GroupUpdate.md) |  | |

### Return type

[**Group**](Group.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## AutoscalingGroupsServersFindById

```go
var result Server = AutoscalingGroupsServersFindById(ctx, groupId, serverId)
                      .Depth(depth)
                      .Execute()
```

Retrieve group servers by UUID



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsServersFindById(context.Background(), groupId, serverId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsServersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsServersFindById`: Server
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsServersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |
|**serverId** | [**string**](.md) |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsServersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |

### Return type

[**Server**](Server.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## AutoscalingGroupsServersGet

```go
var result ServerCollection = AutoscalingGroupsServersGet(ctx, groupId)
                      .Depth(depth)
                      .OrderBy(orderBy)
                      .Execute()
```

Retrieve autoscaling group servers



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
    depth := float32(8.14) // float32 | Controls the detail depth of the response objects.    - depth=0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth=1: Direct properties and children references are included.    - depth=2: Direct properties and children properties are included.    - depth=3: Direct properties and children properties and children's children are included.    - depth=... and so on   (optional)
    orderBy := "orderBy_example" // string | Define the property to be used for ordering the returned list; valid values are 'createdDate' and 'lastModifiedDate'. (optional) (default to "createdDate")

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.GroupsApi.AutoscalingGroupsServersGet(context.Background(), groupId).Depth(depth).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.AutoscalingGroupsServersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `AutoscalingGroupsServersGet`: ServerCollection
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.AutoscalingGroupsServersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**groupId** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingGroupsServersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **float32** | Controls the detail depth of the response objects.    - depth&#x3D;0: Only direct properties are included; children (such as executions or transitions) are not included.    - depth&#x3D;1: Direct properties and children references are included.    - depth&#x3D;2: Direct properties and children properties are included.    - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.    - depth&#x3D;... and so on   | |
| **orderBy** | **string** | Define the property to be used for ordering the returned list; valid values are &#39;createdDate&#39; and &#39;lastModifiedDate&#39;. | [default to &quot;createdDate&quot;]|

### Return type

[**ServerCollection**](ServerCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


