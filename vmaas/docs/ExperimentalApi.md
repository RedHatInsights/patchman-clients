# \ExperimentalApi

All URIs are relative to *http://localhost/api/vmaas/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmaasWebappAppRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#VmaasWebappAppRPMPkgNamesHandlerGetGet) | **Get** /package_names/rpms/{rpm} | 
[**VmaasWebappAppRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#VmaasWebappAppRPMPkgNamesHandlerPostPost) | **Post** /package_names/rpms | 
[**VmaasWebappAppSRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#VmaasWebappAppSRPMPkgNamesHandlerGetGet) | **Get** /package_names/srpms/{srpm} | 
[**VmaasWebappAppSRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#VmaasWebappAppSRPMPkgNamesHandlerPostPost) | **Post** /package_names/srpms | 



## VmaasWebappAppRPMPkgNamesHandlerGetGet

> RPMPkgNamesResponse VmaasWebappAppRPMPkgNamesHandlerGetGet(ctx, rpm).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rpm := ""openssl-libs"" // string | Package name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerGetGet(context.Background(), rpm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppRPMPkgNamesHandlerGetGet`: RPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpm** | **string** | Package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppRPMPkgNamesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RPMPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmaasWebappAppRPMPkgNamesHandlerPostPost

> RPMPkgNamesResponse VmaasWebappAppRPMPkgNamesHandlerPostPost(ctx).RPMPkgNamesRequest(rPMPkgNamesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rPMPkgNamesRequest := *openapiclient.NewRPMPkgNamesRequest([]string{"openssl-libs"}) // RPMPkgNamesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerPostPost(context.Background()).RPMPkgNamesRequest(rPMPkgNamesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppRPMPkgNamesHandlerPostPost`: RPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.VmaasWebappAppRPMPkgNamesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppRPMPkgNamesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rPMPkgNamesRequest** | [**RPMPkgNamesRequest**](RPMPkgNamesRequest.md) |  | 

### Return type

[**RPMPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmaasWebappAppSRPMPkgNamesHandlerGetGet

> SRPMPkgNamesResponse VmaasWebappAppSRPMPkgNamesHandlerGetGet(ctx, srpm).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    srpm := ""openssl"" // string | Source package name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerGetGet(context.Background(), srpm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppSRPMPkgNamesHandlerGetGet`: SRPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srpm** | **string** | Source package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppSRPMPkgNamesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SRPMPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VmaasWebappAppSRPMPkgNamesHandlerPostPost

> SRPMPkgNamesResponse VmaasWebappAppSRPMPkgNamesHandlerPostPost(ctx).SRPMPkgNamesRequest(sRPMPkgNamesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sRPMPkgNamesRequest := *openapiclient.NewSRPMPkgNamesRequest([]string{"openssl"}) // SRPMPkgNamesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerPostPost(context.Background()).SRPMPkgNamesRequest(sRPMPkgNamesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppSRPMPkgNamesHandlerPostPost`: SRPMPkgNamesResponse
    fmt.Fprintf(os.Stdout, "Response from `ExperimentalApi.VmaasWebappAppSRPMPkgNamesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppSRPMPkgNamesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sRPMPkgNamesRequest** | [**SRPMPkgNamesRequest**](SRPMPkgNamesRequest.md) |  | 

### Return type

[**SRPMPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

