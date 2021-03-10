# \SapSystemApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSystemProfileGetSapSids**](SapSystemApi.md#ApiSystemProfileGetSapSids) | **Get** /system_profile/sap_sids | get sap system values
[**ApiSystemProfileGetSapSystem**](SapSystemApi.md#ApiSystemProfileGetSapSystem) | **Get** /system_profile/sap_system | get sap system values



## ApiSystemProfileGetSapSids

> InlineResponse2002 ApiSystemProfileGetSapSids(ctx).Search(search).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()

get sap system values



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
    search := "search_example" // string | Only include tags that match the given search string. The value is matched against namespace, key and value. (optional)
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := TODO // map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SapSystemApi.ApiSystemProfileGetSapSids(context.Background()).Search(search).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SapSystemApi.ApiSystemProfileGetSapSids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSystemProfileGetSapSids`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `SapSystemApi.ApiSystemProfileGetSapSids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSystemProfileGetSapSidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Only include tags that match the given search string. The value is matched against namespace, key and value. | 
 **tags** | **[]string** | filters out hosts not tagged by the given tags | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **staleness** | **[]string** | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filters hosts based on system_profile fields | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiSystemProfileGetSapSystem

> InlineResponse2002 ApiSystemProfileGetSapSystem(ctx).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()

get sap system values



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
    tags := []string{"Inner_example"} // []string | filters out hosts not tagged by the given tags (optional)
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := TODO // map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SapSystemApi.ApiSystemProfileGetSapSystem(context.Background()).Tags(tags).PerPage(perPage).Page(page).Staleness(staleness).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SapSystemApi.ApiSystemProfileGetSapSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiSystemProfileGetSapSystem`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `SapSystemApi.ApiSystemProfileGetSapSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSystemProfileGetSapSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **[]string** | filters out hosts not tagged by the given tags | 
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **staleness** | **[]string** | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | [**map[string]interface{}**](map[string]interface{}.md) | Filters hosts based on system_profile fields | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

