# \PrincipalApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPrincipals**](PrincipalApi.md#ListPrincipals) | **Get** /principals/ | List the principals for a tenant



## ListPrincipals

> PrincipalPagination ListPrincipals(ctx, optional)

List the principals for a tenant

By default, responses are sorted in ascending order by username

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPrincipalsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPrincipalsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **usernames** | **optional.String**| Usernames of principals to get | 
 **sortOrder** | **optional.String**| The sort order of the query, either ascending or descending | 
 **email** | **optional.String**| Exact e-mail address of principal to search for | 
 **status** | **optional.String**| Set the status of users to get back. Could not be used with: usernames, email, admin_only | [default to enabled]
 **adminOnly** | **optional.String**| Get only admin users within an account. Setting this would ignore the parameters: usernames, email | [default to false]

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

