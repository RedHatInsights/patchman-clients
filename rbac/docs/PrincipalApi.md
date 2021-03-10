# \PrincipalApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrincipals**](PrincipalApi.md#ListPrincipals) | **Get** /principals/ | List the principals for a tenant



## ListPrincipals

> PrincipalPagination ListPrincipals(ctx).Limit(limit).Offset(offset).MatchCriteria(matchCriteria).Usernames(usernames).SortOrder(sortOrder).Email(email).Status(status).AdminOnly(adminOnly).OrderBy(orderBy).Execute()

List the principals for a tenant



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
    matchCriteria := "matchCriteria_example" // string | Parameter for specifying the matching criteria for an object's name or email. (optional) (default to "exact")
    usernames := "userA,userB" // string | Comma separated usernames of principals to get. If match_criteria is specified, only the first username will be picked up for search. (optional)
    sortOrder := "sortOrder_example" // string | The sort order of the query, either ascending or descending. Defaults to ascending. (optional)
    email := "email_example" // string | E-mail address of principal to search for. Could be combined with match_criteria for searching. (optional)
    status := "status_example" // string | Set the status of users to get back. Could not be used with: usernames, email, admin_only (optional) (default to "enabled")
    adminOnly := "adminOnly_example" // string | Get only admin users within an account. Setting this would ignore the parameters: usernames, email (optional) (default to "false")
    orderBy := "orderBy_example" // string | Parameter for ordering principals by value. For inverse ordering, supply '-' before the param value, such as: ?order_by=-username (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PrincipalApi.ListPrincipals(context.Background()).Limit(limit).Offset(offset).MatchCriteria(matchCriteria).Usernames(usernames).SortOrder(sortOrder).Email(email).Status(status).AdminOnly(adminOnly).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrincipalApi.ListPrincipals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrincipals`: PrincipalPagination
    fmt.Fprintf(os.Stdout, "Response from `PrincipalApi.ListPrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **int32** | Parameter for selecting the offset of data. | [default to 0]
 **matchCriteria** | **string** | Parameter for specifying the matching criteria for an object&#39;s name or email. | [default to &quot;exact&quot;]
 **usernames** | **string** | Comma separated usernames of principals to get. If match_criteria is specified, only the first username will be picked up for search. | 
 **sortOrder** | **string** | The sort order of the query, either ascending or descending. Defaults to ascending. | 
 **email** | **string** | E-mail address of principal to search for. Could be combined with match_criteria for searching. | 
 **status** | **string** | Set the status of users to get back. Could not be used with: usernames, email, admin_only | [default to &quot;enabled&quot;]
 **adminOnly** | **string** | Get only admin users within an account. Setting this would ignore the parameters: usernames, email | [default to &quot;false&quot;]
 **orderBy** | **string** | Parameter for ordering principals by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-username | 

### Return type

[**PrincipalPagination**](PrincipalPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

