# \CrossAccountRequestApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCrossAccountRequests**](CrossAccountRequestApi.md#CreateCrossAccountRequests) | **Post** /cross-account-requests/ | Create a cross account request
[**GetCrossAccountRequest**](CrossAccountRequestApi.md#GetCrossAccountRequest) | **Get** /cross-account-requests/{uuid}/ | Get a cross account request
[**ListCrossAccountRequests**](CrossAccountRequestApi.md#ListCrossAccountRequests) | **Get** /cross-account-requests/ | List the cross account requests for a user or account
[**PatchCrossAccountRequest**](CrossAccountRequestApi.md#PatchCrossAccountRequest) | **Patch** /cross-account-requests/{uuid}/ | Update a cross account request
[**PutCrossAccountRequest**](CrossAccountRequestApi.md#PutCrossAccountRequest) | **Put** /cross-account-requests/{uuid}/ | Update a cross account request



## CreateCrossAccountRequests

> CrossAccountRequestOut CreateCrossAccountRequests(ctx).CrossAccountRequestIn(crossAccountRequestIn).Execute()

Create a cross account request

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
    crossAccountRequestIn := *openapiclient.NewCrossAccountRequestIn("12345", "01/01/2021", "01/01/2021", []string{"Role Name"}) // CrossAccountRequestIn | CrossAccountRequest to create

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CrossAccountRequestApi.CreateCrossAccountRequests(context.Background()).CrossAccountRequestIn(crossAccountRequestIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrossAccountRequestApi.CreateCrossAccountRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCrossAccountRequests`: CrossAccountRequestOut
    fmt.Fprintf(os.Stdout, "Response from `CrossAccountRequestApi.CreateCrossAccountRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCrossAccountRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crossAccountRequestIn** | [**CrossAccountRequestIn**](CrossAccountRequestIn.md) | CrossAccountRequest to create | 

### Return type

[**CrossAccountRequestOut**](CrossAccountRequestOut.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossAccountRequest

> CrossAccountRequestDetailByAccount GetCrossAccountRequest(ctx, uuid).QueryBy(queryBy).Account(account).ApprovedOnly(approvedOnly).Execute()

Get a cross account request

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
    uuid := TODO // string | ID of cross account request to get
    queryBy := "queryBy_example" // string | Parameter for filtering resource by either a user's ID, or a client's account number. The default value is target_account. (optional)
    account := "account_example" // string | Parameter for filtering resource by an account number. Value can be a comma-separated list of ids. To be used in tandem with ?query_by=user_id to further filter a user's requests by account number. (optional)
    approvedOnly := "approvedOnly_example" // string | Parameter for filtering resource which have been approved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CrossAccountRequestApi.GetCrossAccountRequest(context.Background(), uuid).QueryBy(queryBy).Account(account).ApprovedOnly(approvedOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrossAccountRequestApi.GetCrossAccountRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrossAccountRequest`: CrossAccountRequestDetailByAccount
    fmt.Fprintf(os.Stdout, "Response from `CrossAccountRequestApi.GetCrossAccountRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of cross account request to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossAccountRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **queryBy** | **string** | Parameter for filtering resource by either a user&#39;s ID, or a client&#39;s account number. The default value is target_account. | 
 **account** | **string** | Parameter for filtering resource by an account number. Value can be a comma-separated list of ids. To be used in tandem with ?query_by&#x3D;user_id to further filter a user&#39;s requests by account number. | 
 **approvedOnly** | **string** | Parameter for filtering resource which have been approved. | 

### Return type

[**CrossAccountRequestDetailByAccount**](CrossAccountRequestDetailByAccount.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCrossAccountRequests

> CrossAccountRequestPagination ListCrossAccountRequests(ctx).Limit(limit).Offset(offset).QueryBy(queryBy).Account(account).ApprovedOnly(approvedOnly).Status(status).OrderBy(orderBy).Execute()

List the cross account requests for a user or account



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
    limit := int32(56) // int32 | Parameter for selecting the amount of data returned. (optional) (default to 10)
    offset := int32(56) // int32 | Parameter for selecting the offset of data. (optional) (default to 0)
    queryBy := "queryBy_example" // string | Parameter for filtering resource by either a user's ID, or a client's account number. The default value is target_account. (optional)
    account := "account_example" // string | Parameter for filtering resource by an account number. Value can be a comma-separated list of ids. To be used in tandem with ?query_by=user_id to further filter a user's requests by account number. (optional)
    approvedOnly := "approvedOnly_example" // string | Parameter for filtering resource which have been approved. (optional)
    status := "status_example" // string | Parameter for filtering resource based on status. (optional)
    orderBy := "orderBy_example" // string | Parameter for ordering by field. For inverse ordering, use '-', e.g. ?order_by=-start_date. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CrossAccountRequestApi.ListCrossAccountRequests(context.Background()).Limit(limit).Offset(offset).QueryBy(queryBy).Account(account).ApprovedOnly(approvedOnly).Status(status).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrossAccountRequestApi.ListCrossAccountRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCrossAccountRequests`: CrossAccountRequestPagination
    fmt.Fprintf(os.Stdout, "Response from `CrossAccountRequestApi.ListCrossAccountRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCrossAccountRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **queryBy** | **string** | Parameter for filtering resource by either a user&#39;s ID, or a client&#39;s account number. The default value is target_account. | 
 **account** | **string** | Parameter for filtering resource by an account number. Value can be a comma-separated list of ids. To be used in tandem with ?query_by&#x3D;user_id to further filter a user&#39;s requests by account number. | 
 **approvedOnly** | **string** | Parameter for filtering resource which have been approved. | 
 **status** | **string** | Parameter for filtering resource based on status. | 
 **orderBy** | **string** | Parameter for ordering by field. For inverse ordering, use &#39;-&#39;, e.g. ?order_by&#x3D;-start_date. | 

### Return type

[**CrossAccountRequestPagination**](CrossAccountRequestPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrossAccountRequest

> CrossAccountRequestDetailByAccount PatchCrossAccountRequest(ctx, uuid).CrossAccountRequestPatch(crossAccountRequestPatch).Execute()

Update a cross account request



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
    uuid := TODO // string | ID of cross account request to get
    crossAccountRequestPatch := *openapiclient.NewCrossAccountRequestPatch() // CrossAccountRequestPatch | Updates to CrossAccountRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CrossAccountRequestApi.PatchCrossAccountRequest(context.Background(), uuid).CrossAccountRequestPatch(crossAccountRequestPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrossAccountRequestApi.PatchCrossAccountRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCrossAccountRequest`: CrossAccountRequestDetailByAccount
    fmt.Fprintf(os.Stdout, "Response from `CrossAccountRequestApi.PatchCrossAccountRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of cross account request to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrossAccountRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossAccountRequestPatch** | [**CrossAccountRequestPatch**](CrossAccountRequestPatch.md) | Updates to CrossAccountRequest | 

### Return type

[**CrossAccountRequestDetailByAccount**](CrossAccountRequestDetailByAccount.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrossAccountRequest

> CrossAccountRequestDetailByAccount PutCrossAccountRequest(ctx, uuid).CrossAccountRequestUpdateIn(crossAccountRequestUpdateIn).Execute()

Update a cross account request



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
    uuid := TODO // string | ID of cross account request to get
    crossAccountRequestUpdateIn := *openapiclient.NewCrossAccountRequestUpdateIn("01/01/2021", "01/01/2021", []string{"Role Name"}) // CrossAccountRequestUpdateIn | Updates to CrossAccountRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CrossAccountRequestApi.PutCrossAccountRequest(context.Background(), uuid).CrossAccountRequestUpdateIn(crossAccountRequestUpdateIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrossAccountRequestApi.PutCrossAccountRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutCrossAccountRequest`: CrossAccountRequestDetailByAccount
    fmt.Fprintf(os.Stdout, "Response from `CrossAccountRequestApi.PutCrossAccountRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md) | ID of cross account request to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrossAccountRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crossAccountRequestUpdateIn** | [**CrossAccountRequestUpdateIn**](CrossAccountRequestUpdateIn.md) | Updates to CrossAccountRequest | 

### Return type

[**CrossAccountRequestDetailByAccount**](CrossAccountRequestDetailByAccount.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

