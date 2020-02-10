# \PolicyApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicies**](PolicyApi.md#CreatePolicies) | **Post** /policies/ | Create a policy in a tenant
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /policies/{uuid}/ | Delete a policy in the tenant
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /policies/{uuid}/ | Get a policy in the tenant
[**ListPolicies**](PolicyApi.md#ListPolicies) | **Get** /policies/ | List the policies in the tenant
[**UpdatePolicy**](PolicyApi.md#UpdatePolicy) | **Put** /policies/{uuid}/ | Update a policy in the tenant



## CreatePolicies

> PolicyExtended CreatePolicies(ctx, policyIn)

Create a policy in a tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyIn** | [**PolicyIn**](PolicyIn.md)| Policy to create | 

### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, uuid)

Delete a policy in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of policy to delete | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> PolicyExtended GetPolicy(ctx, uuid)

Get a policy in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of policy to get | 

### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> PolicyPagination ListPolicies(ctx, optional)

List the policies in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **name** | **optional.String**| Parameter for filtering resource by name using string contains search. | 
 **scope** | **optional.String**| Parameter for filtering resource by scope. | [default to account]
 **groupName** | **optional.String**| Parameter for filtering resource by group name using string contains search. | 
 **groupUuid** | [**optional.Interface of string**](.md)| Parameter for filtering resource by group uuid using UUID exact match. | 
 **orderBy** | **optional.String**| Parameter for ordering resource by value. | 

### Return type

[**PolicyPagination**](PolicyPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> PolicyExtended UpdatePolicy(ctx, uuid, policyIn)

Update a policy in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of policy to update | 
**policyIn** | [**PolicyIn**](PolicyIn.md)| Policy to update | 

### Return type

[**PolicyExtended**](PolicyExtended.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

