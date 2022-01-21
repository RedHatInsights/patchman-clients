# \HostsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHostDeleteById**](HostsApi.md#ApiHostDeleteById) | **Delete** /hosts/{host_id_list} | Delete hosts by IDs
[**ApiHostDeleteHostList**](HostsApi.md#ApiHostDeleteHostList) | **Delete** /hosts | Delete the entire list of hosts filtered by the given parameters
[**ApiHostGetHostById**](HostsApi.md#ApiHostGetHostById) | **Get** /hosts/{host_id_list} | Find hosts by their IDs
[**ApiHostGetHostList**](HostsApi.md#ApiHostGetHostList) | **Get** /hosts | Read the entire list of hosts
[**ApiHostGetHostSystemProfileById**](HostsApi.md#ApiHostGetHostSystemProfileById) | **Get** /hosts/{host_id_list}/system_profile | Return one or more hosts system profile
[**ApiHostGetHostTagCount**](HostsApi.md#ApiHostGetHostTagCount) | **Get** /hosts/{host_id_list}/tags/count | Get the number of tags on a host
[**ApiHostGetHostTags**](HostsApi.md#ApiHostGetHostTags) | **Get** /hosts/{host_id_list}/tags | Get the tags on a host
[**ApiHostHostCheckin**](HostsApi.md#ApiHostHostCheckin) | **Post** /hosts/checkin | Update staleness timestamps for a host matching the provided facts
[**ApiHostMergeFacts**](HostsApi.md#ApiHostMergeFacts) | **Patch** /hosts/{host_id_list}/facts/{namespace} | Merge facts under a namespace
[**ApiHostPatchById**](HostsApi.md#ApiHostPatchById) | **Patch** /hosts/{host_id_list} | Update a host
[**ApiHostReplaceFacts**](HostsApi.md#ApiHostReplaceFacts) | **Put** /hosts/{host_id_list}/facts/{namespace} | Replace facts under a namespace



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


## ApiHostDeleteHostList

> ApiHostDeleteHostList(ctx).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).RegisteredWith(registeredWith).Staleness(staleness).Tags(tags).Filter(filter).Execute()

Delete the entire list of hosts filtered by the given parameters



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
    displayName := "displayName_example" // string | Filter by display_name (optional)
    fqdn := "fqdn_example" // string | Filter by FQDN (optional)
    hostnameOrId := "hostnameOrId_example" // string | Filter by display_name, fqdn, id (optional)
    insightsId := TODO // string | Filter by insights_id (optional)
    providerId := "providerId_example" // string | Filter by provider_id (optional)
    providerType := "providerType_example" // string | Filter by provider_type (optional)
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    filter := map[string]OneOfmapobject{"key": "TODO"} // map[string]OneOfmapobject | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostDeleteHostList(context.Background()).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).RegisteredWith(registeredWith).Staleness(staleness).Tags(tags).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostDeleteHostList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostDeleteHostListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **string** | Filter by display_name | 
 **fqdn** | **string** | Filter by FQDN | 
 **hostnameOrId** | **string** | Filter by display_name, fqdn, id | 
 **insightsId** | [**string**](string.md) | Filter by insights_id | 
 **providerId** | **string** | Filter by provider_id | 
 **providerType** | **string** | Filter by provider_type | 
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **staleness** | **[]string** | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **tags** | **[]string** | filters out hosts not tagged by the given tags | 
 **filter** | [**map[string]OneOfmapobject**](oneOf&lt;map,object&gt;.md) | Filters hosts based on system_profile fields | 

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
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system (optional)

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
 **orderHow** | **string** | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system | 

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

> HostQueryOutput ApiHostGetHostList(ctx).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Staleness(staleness).Tags(tags).RegisteredWith(registeredWith).Filter(filter).Fields(fields).Execute()

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
    displayName := "displayName_example" // string | Filter by display_name (optional)
    fqdn := "fqdn_example" // string | Filter by FQDN (optional)
    hostnameOrId := "hostnameOrId_example" // string | Filter by display_name, fqdn, id (optional)
    insightsId := TODO // string | Filter by insights_id (optional)
    providerId := "providerId_example" // string | Filter by provider_id (optional)
    providerType := "providerType_example" // string | Filter by provider_type (optional)
    branchId := "branchId_example" // string | Filter by branch_id (optional)
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system (optional)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := map[string]OneOfmapobject{"key": "TODO"} // map[string]OneOfmapobject | Filters hosts based on system_profile fields (optional)
    fields := map[string]OneOfmapobject{"key": "TODO"} // map[string]OneOfmapobject | Fetches only mentioned system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostList(context.Background()).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).BranchId(branchId).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).Staleness(staleness).Tags(tags).RegisteredWith(registeredWith).Filter(filter).Fields(fields).Execute()
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
 **displayName** | **string** | Filter by display_name | 
 **fqdn** | **string** | Filter by FQDN | 
 **hostnameOrId** | **string** | Filter by display_name, fqdn, id | 
 **insightsId** | [**string**](string.md) | Filter by insights_id | 
 **providerId** | **string** | Filter by provider_id | 
 **providerType** | **string** | Filter by provider_type | 
 **branchId** | **string** | Filter by branch_id | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **orderBy** | **string** | Ordering field name | 
 **orderHow** | **string** | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system | 
 **staleness** | **[]string** | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **tags** | **[]string** | filters out hosts not tagged by the given tags | 
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]OneOfmapobject**](oneOf&lt;map,object&gt;.md) | Filters hosts based on system_profile fields | 
 **fields** | [**map[string]OneOfmapobject**](oneOf&lt;map,object&gt;.md) | Fetches only mentioned system_profile fields | 

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

> SystemProfileByHostOut ApiHostGetHostSystemProfileById(ctx, hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).BranchId(branchId).Fields(fields).Execute()

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
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system (optional)
    branchId := "branchId_example" // string | Filter by branch_id (optional)
    fields := map[string]OneOfmapobject{"key": "TODO"} // map[string]OneOfmapobject | Fetches only mentioned system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostGetHostSystemProfileById(context.Background(), hostIdList).PerPage(perPage).Page(page).OrderBy(orderBy).OrderHow(orderHow).BranchId(branchId).Fields(fields).Execute()
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
 **orderHow** | **string** | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system | 
 **branchId** | **string** | Filter by branch_id | 
 **fields** | [**map[string]OneOfmapobject**](oneOf&lt;map,object&gt;.md) | Fetches only mentioned system_profile fields | 

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
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system (optional)

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
 **orderHow** | **string** | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system | 

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
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    orderBy := "orderBy_example" // string | Ordering field name (optional)
    orderHow := "orderHow_example" // string | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system (optional)
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
 **orderHow** | **string** | Direction of the ordering; defaults to ASC for display_name, and to DESC for updated and operating_system | 
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

> CreateHostOut ApiHostHostCheckin(ctx).CreateCheckIn(createCheckIn).Execute()

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
    createCheckIn := *openapiclient.NewCreateCheckIn(string(123), string(123), string(123), string(123), []string{"IpAddresses_example"}, "Fqdn_example", []string{"MacAddresses_example"}, "ProviderId_example", "ProviderType_example") // CreateCheckIn | Data required to create a check-in record for a host.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HostsApi.ApiHostHostCheckin(context.Background()).CreateCheckIn(createCheckIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ApiHostHostCheckin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiHostHostCheckin`: CreateHostOut
    fmt.Fprintf(os.Stdout, "Response from `HostsApi.ApiHostHostCheckin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiHostHostCheckinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheckIn** | [**CreateCheckIn**](CreateCheckIn.md) | Data required to create a check-in record for a host. | 

### Return type

[**CreateHostOut**](CreateHostOut.md)

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
    body := map[string]interface{}(Object) // map[string]interface{} | A dictionary with the new facts to merge with the original ones.
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
    body := map[string]interface{}(Object) // map[string]interface{} | A dictionary with the new facts to replace the original ones.
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

