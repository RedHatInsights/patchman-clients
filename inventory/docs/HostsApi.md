# \HostsApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHostAddHostList**](HostsApi.md#ApiHostAddHostList) | **Post** /hosts | Create/update multiple host and add them to the host list
[**ApiHostDeleteById**](HostsApi.md#ApiHostDeleteById) | **Delete** /hosts/{host_id_list} | Delete hosts by IDs
[**ApiHostGetHostById**](HostsApi.md#ApiHostGetHostById) | **Get** /hosts/{host_id_list} | Find hosts by their IDs
[**ApiHostGetHostList**](HostsApi.md#ApiHostGetHostList) | **Get** /hosts | Read the entire list of hosts
[**ApiHostGetHostSystemProfileById**](HostsApi.md#ApiHostGetHostSystemProfileById) | **Get** /hosts/{host_id_list}/system_profile | Return one or more hosts system profile
[**ApiHostGetHostTagCount**](HostsApi.md#ApiHostGetHostTagCount) | **Get** /hosts/{host_id_list}/tags/count | Get the number of tags on a host
[**ApiHostGetHostTags**](HostsApi.md#ApiHostGetHostTags) | **Get** /hosts/{host_id_list}/tags | Get the tags on a host
[**ApiHostHostCheckin**](HostsApi.md#ApiHostHostCheckin) | **Put** /hosts/checkin | Update staleness timestamps for a host matching the provided facts
[**ApiHostMergeFacts**](HostsApi.md#ApiHostMergeFacts) | **Patch** /hosts/{host_id_list}/facts/{namespace} | Merge facts under a namespace
[**ApiHostPatchById**](HostsApi.md#ApiHostPatchById) | **Patch** /hosts/{host_id_list} | Update a host
[**ApiHostReplaceFacts**](HostsApi.md#ApiHostReplaceFacts) | **Put** /hosts/{host_id_list}/facts/{namespace} | Replace facts under a namespace



## ApiHostAddHostList

> BulkHostOut ApiHostAddHostList(ctx).CreateHostIn(createHostIn).Execute()

Create/update multiple host and add them to the host list



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
    createHostIn := []CreateHostIn{*openapiclient.NewCreateHostIn("Account_example", "Reporter_example", time.Now())} // []CreateHostIn | A list of host objects to be added to the host list

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostAddHostList(context.Background()).CreateHostIn(createHostIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostAddHostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostAddHostList`: BulkHostOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostAddHostList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostAddHostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHostIn** | [**[]CreateHostIn**](CreateHostIn.md) | A list of host objects to be added to the host list | 

### Return type

[**BulkHostOut**](BulkHostOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostDeleteById

> ApiHostDeleteById(ctx, hostIdList).BranchId(branchId).Execute()

Delete hosts by IDs



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    branchId := "branchId_example" // string | Filter by branch_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostDeleteById(context.Background(), hostIdList).BranchId(branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostDeleteById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostDeleteByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **string** | Filter by branch_id | 

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


## ApiHostGetHostById

> HostQueryOutput ApiHostGetHostById(ctx, hostIdList).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Execute()

Find hosts by their IDs



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    branchId := "branchId_example" // string | Filter by branch_id (optional)
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering, defaults to ASC for display_name and to DESC for updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostById(context.Background(), hostIdList).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostGetHostById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostGetHostById`: HostQueryOutput
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostGetHostById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetHostByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **string** | Filter by branch_id | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**HostQueryOutput**](HostQueryOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostList

> HostQueryOutput ApiHostGetHostList(ctx).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Staleness(staleness).Tags(tags).RegisteredWith(registeredWith).Filter(filter).Execute()

Read the entire list of hosts



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
    displayName := "displayName_example" // string | A part of a searched host’s display name. (optional)
    fqdn := "fqdn_example" // string | Filter by a host's FQDN (optional)
    hostnameOrId := "hostnameOrId_example" // string | Search for a host by display_name, fqdn, id (optional)
    insightsId := TODO // string | Search for a host by insights_id (optional)
    branchId := "branchId_example" // string | Filter by branch_id (optional)
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering, defaults to ASC for display_name and to DESC for updated (optional)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := TODO // map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostList(context.Background()).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Staleness(staleness).Tags(tags).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostGetHostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostGetHostList`: HostQueryOutput
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostGetHostList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetHostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **string** | A part of a searched host’s display name. | 
 **fqdn** | **string** | Filter by a host&#39;s FQDN | 
 **hostnameOrId** | **string** | Search for a host by display_name, fqdn, id | 
 **insightsId** | [**string**](.md) | Search for a host by insights_id | 
 **branchId** | **string** | Filter by branch_id | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **staleness** | [**[]string**](string.md) | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **tags** | [**[]string**](string.md) | filters out hosts not tagged by the given tags | 
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]interface{}**](.md) | Filters hosts based on system_profile fields | 

### Return type

[**HostQueryOutput**](HostQueryOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostSystemProfileById

> SystemProfileByHostOut ApiHostGetHostSystemProfileById(ctx, hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).BranchId(branchId).Execute()

Return one or more hosts system profile



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering, defaults to ASC for display_name and to DESC for updated (optional)
    branchId := "branchId_example" // string | Filter by branch_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostSystemProfileById(context.Background(), hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).BranchId(branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostGetHostSystemProfileById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostGetHostSystemProfileById`: SystemProfileByHostOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostGetHostSystemProfileById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetHostSystemProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **branchId** | **string** | Filter by branch_id | 

### Return type

[**SystemProfileByHostOut**](SystemProfileByHostOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTagCount

> TagCountOut ApiHostGetHostTagCount(ctx, hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Execute()

Get the number of tags on a host



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering, defaults to ASC for display_name and to DESC for updated (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostTagCount(context.Background(), hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostGetHostTagCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostGetHostTagCount`: TagCountOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostGetHostTagCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetHostTagCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**TagCountOut**](TagCountOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTags

> TagsOut ApiHostGetHostTags(ctx, hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Search(search).Execute()

Get the tags on a host



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    perPage := 987 // int32 | A number of items to return per page. (optional) (default to 50)
    page := 987 // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering, defaults to ASC for display_name and to DESC for updated (optional)
    search := "search_example" // string | Only include tags that match the given search string. The value is matched against namespace, key and value. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostTags(context.Background(), hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostGetHostTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostGetHostTags`: TagsOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostGetHostTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostGetHostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **search** | **string** | Only include tags that match the given search string. The value is matched against namespace, key and value. | 

### Return type

[**TagsOut**](TagsOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostCheckin

> BulkHostOut ApiHostHostCheckin(ctx).CreateCheckIn(createCheckIn).Execute()

Update staleness timestamps for a host matching the provided facts



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
    createCheckIn := *openapiclient.NewCreateCheckIn(123) // CreateCheckIn | A list of host objects to be added to the host list

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostHostCheckin(context.Background()).CreateCheckIn(createCheckIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostHostCheckin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostHostCheckin`: BulkHostOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostHostCheckin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostCheckinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckIn** | [**CreateCheckIn**](CreateCheckIn.md) | A list of host objects to be added to the host list | 

### Return type

[**BulkHostOut**](BulkHostOut.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostMergeFacts

> ApiHostMergeFacts(ctx, hostIdList, namespace).Body(body).BranchId(branchId).Execute()

Merge facts under a namespace



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    namespace := "namespace_example" // string | A namespace of the merged facts.
    body := 987 // map[string]interface{} | A dictionary with the new facts to merge with the original ones.
    branchId := "branchId_example" // string | Filter by branch_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostMergeFacts(context.Background(), hostIdList, namespace).Body(body).BranchId(branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostMergeFacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 
**namespace** | **string** | A namespace of the merged facts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostMergeFactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | A dictionary with the new facts to merge with the original ones. | 
 **branchId** | **string** | Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostPatchById

> ApiHostPatchById(ctx, hostIdList).PatchHostIn(patchHostIn).BranchId(branchId).Execute()

Update a host



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    patchHostIn := *openapiclient.NewPatchHostIn() // PatchHostIn | A group of fields to be updated on the host
    branchId := "branchId_example" // string | Filter by branch_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostPatchById(context.Background(), hostIdList).PatchHostIn(patchHostIn).BranchId(branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostPatchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostPatchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchHostIn** | [**PatchHostIn**](PatchHostIn.md) | A group of fields to be updated on the host | 
 **branchId** | **string** | Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostReplaceFacts

> ApiHostReplaceFacts(ctx, hostIdList, namespace).Body(body).BranchId(branchId).Execute()

Replace facts under a namespace



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
    hostIdList := []string{"Inner_example"} // []string | A comma separated list of host IDs.
    namespace := "namespace_example" // string | A namespace of the merged facts.
    body := 987 // map[string]interface{} | A dictionary with the new facts to replace the original ones.
    branchId := "branchId_example" // string | Filter by branch_id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostReplaceFacts(context.Background(), hostIdList, namespace).Body(body).BranchId(branchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostReplaceFacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md) | A comma separated list of host IDs. | 
**namespace** | **string** | A namespace of the merged facts. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiHostReplaceFactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | A dictionary with the new facts to replace the original ones. | 
 **branchId** | **string** | Filter by branch_id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

