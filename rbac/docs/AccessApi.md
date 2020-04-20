# \AccessApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrincipalAccess**](AccessApi.md#GetPrincipalAccess) | **Get** /access/ | Get the permitted access for a principal in the tenant (defaults to principal from the identity header)



## GetPrincipalAccess

> AccessPagination GetPrincipalAccess(ctx, application, optional)

Get the permitted access for a principal in the tenant (defaults to principal from the identity header)

Access responses are sorted in ascending order by an ID internal to the database

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| The application name(s) to obtain access for the principal. This is an exact match. When no application is supplied, all permissions for the principal are returned. You may also use a comma-separated list to match on multiple applications. | 
 **optional** | ***GetPrincipalAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPrincipalAccessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Unique username of the principal to obtain access for (only available for admins, and if supplied, takes precedence over the identity header). | 
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]

### Return type

[**AccessPagination**](AccessPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

