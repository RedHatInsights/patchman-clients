# \TagsApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTagGetTags**](TagsApi.md#ApiTagGetTags) | **Get** /tags | Get the active host tags for a given account



## ApiTagGetTags

> ActiveTags ApiTagGetTags(ctx).Tags(tags).OrderBy(orderBy).OrderHow(orderHow).PerPage(perPage).Page(page).Staleness(staleness).Search(search).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).RegisteredWith(registeredWith).Filter(filter).Execute()

Get the active host tags for a given account



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
    orderBy := "orderBy_example" // string | Ordering field name (optional) (default to "tag")
    orderHow := "orderHow_example" // string | Direction of the ordering. Default to ASC (optional) (default to "ASC")
    perPage := int32(56) // int32 | A number of items to return per page. (optional) (default to 50)
    page := int32(56) // int32 | A page number of the items to return. (optional) (default to 1)
    staleness := []string{"Staleness_example"} // []string | Culling states of the hosts. Default: fresh,stale,unknown (optional) (default to ["fresh","stale","unknown"])
    search := "search_example" // string | Only include tags that match the given search string. The value is matched against namespace, key and value. (optional)
    displayName := "displayName_example" // string | Filter by display_name (optional)
    fqdn := "fqdn_example" // string | Filter by FQDN (optional)
    hostnameOrId := "hostnameOrId_example" // string | Filter by display_name, fqdn, id (optional)
    insightsId := TODO // string | Filter by insights_id (optional)
    providerId := "providerId_example" // string | Filter by provider_id (optional)
    providerType := "providerType_example" // string | Filter by provider_type (optional)
    registeredWith := "registeredWith_example" // string | Filters out any host not registered with the specified service (optional)
    filter := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | Filters hosts based on system_profile fields (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TagsApi.ApiTagGetTags(context.Background()).Tags(tags).OrderBy(orderBy).OrderHow(orderHow).PerPage(perPage).Page(page).Staleness(staleness).Search(search).DisplayName(displayName).Fqdn(fqdn).HostnameOrId(hostnameOrId).InsightsId(insightsId).ProviderId(providerId).ProviderType(providerType).RegisteredWith(registeredWith).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagsApi.ApiTagGetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTagGetTags`: ActiveTags
    fmt.Fprintf(os.Stdout, "Response from `TagsApi.ApiTagGetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTagGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **[]string** | filters out hosts not tagged by the given tags | 
 **orderBy** | **string** | Ordering field name | [default to &quot;tag&quot;]
 **orderHow** | **string** | Direction of the ordering. Default to ASC | [default to &quot;ASC&quot;]
 **perPage** | **int32** | A number of items to return per page. | [default to 50]
 **page** | **int32** | A page number of the items to return. | [default to 1]
 **staleness** | **[]string** | Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **search** | **string** | Only include tags that match the given search string. The value is matched against namespace, key and value. | 
 **displayName** | **string** | Filter by display_name | 
 **fqdn** | **string** | Filter by FQDN | 
 **hostnameOrId** | **string** | Filter by display_name, fqdn, id | 
 **insightsId** | [**string**](string.md) | Filter by insights_id | 
 **providerId** | **string** | Filter by provider_id | 
 **providerType** | **string** | Filter by provider_type | 
 **registeredWith** | **string** | Filters out any host not registered with the specified service | 
 **filter** | **map[string]map[string]interface{}** | Filters hosts based on system_profile fields | 

### Return type

[**ActiveTags**](ActiveTags.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

