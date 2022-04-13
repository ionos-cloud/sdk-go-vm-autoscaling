# \DocumentationApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**AutoscalingOpenapiJsonGet**](DocumentationApi.md#AutoscalingOpenapiJsonGet) | **Get** /cloudapi/autoscaling/openapi.json | Retrieve VM autoscaling OpenAPI spec (JSON)|
|[**AutoscalingOpenapiYamlGet**](DocumentationApi.md#AutoscalingOpenapiYamlGet) | **Get** /cloudapi/autoscaling/openapi.yaml | Retrieve VM autoscaling OpenAPI spec (YAML)|



## AutoscalingOpenapiJsonGet

```go
var result  = AutoscalingOpenapiJsonGet(ctx)
                      .Execute()
```

Retrieve VM autoscaling OpenAPI spec (JSON)



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

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.DocumentationApi.AutoscalingOpenapiJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentationApi.AutoscalingOpenapiJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingOpenapiJsonGetRequest struct via the builder pattern


### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined



## AutoscalingOpenapiYamlGet

```go
var result  = AutoscalingOpenapiYamlGet(ctx)
                      .Execute()
```

Retrieve VM autoscaling OpenAPI spec (YAML)



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

    configuration := ionoscloud_vm_autoscaling.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud_vm_autoscaling.NewAPIClient(configuration)
    resource, resp, err := apiClient.DocumentationApi.AutoscalingOpenapiYamlGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentationApi.AutoscalingOpenapiYamlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiAutoscalingOpenapiYamlGetRequest struct via the builder pattern


### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


