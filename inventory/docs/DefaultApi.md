# \DefaultApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSystemProfileValidateSchema**](DefaultApi.md#ApiSystemProfileValidateSchema) | **Post** /system_profile/validate_schema | validate system profile schema



## ApiSystemProfileValidateSchema

> ApiSystemProfileValidateSchema(ctx).RepoBranch(repoBranch).RepoFork(repoFork).Days(days).MaxMessages(maxMessages).Execute()

validate system profile schema



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
    repoBranch := "repoBranch_example" // string | The branch of the inventory-schemas repo to use
    repoFork := "repoFork_example" // string | The fork of the inventory-schemas repo to use (optional)
    days := int32(56) // int32 | How many days worth of data to validate (optional)
    maxMessages := int32(56) // int32 | Stops polling when this number of messages has been collected (optional) (default to 10000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ApiSystemProfileValidateSchema(context.Background()).RepoBranch(repoBranch).RepoFork(repoFork).Days(days).MaxMessages(maxMessages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiSystemProfileValidateSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSystemProfileValidateSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repoBranch** | **string** | The branch of the inventory-schemas repo to use | 
 **repoFork** | **string** | The fork of the inventory-schemas repo to use | 
 **days** | **int32** | How many days worth of data to validate | 
 **maxMessages** | **int32** | Stops polling when this number of messages has been collected | [default to 10000]

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

