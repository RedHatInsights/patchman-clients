# \PermissionApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPermissionOptions**](PermissionApi.md#ListPermissionOptions) | **Get** /permissions/options/ | List the available options for fields of permissions for a tenant
[**ListPermissions**](PermissionApi.md#ListPermissions) | **Get** /permissions/ | List the permissions for a tenant



## ListPermissionOptions

> PermissionOptionsPagination ListPermissionOptions(ctx, field, optional)

List the available options for fields of permissions for a tenant

By default, options of application is returned. And could be resource_type or verb on demand.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**field** | **string**| specify which fields of permission to display | 
 **optional** | ***ListPermissionOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPermissionOptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **application** | **optional.String**| Filter returned options based on application. You may also use a comma-separated list to filter on multiple applications. | 
 **resourceType** | **optional.String**| Filter returned options based on resource_type. You may also use a comma-separated list to filter on multiple resource_types. | 
 **verb** | **optional.String**| Filter returned options based on verb. You may also use a comma-separated list to filter on multiple verbs. | 

### Return type

[**PermissionOptionsPagination**](PermissionOptionsPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPermissions

> PermissionPagination ListPermissions(ctx, optional)

List the permissions for a tenant

By default, responses are sorted in ascending order by permission application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPermissionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPermissionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **orderBy** | **optional.String**| Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 
 **application** | **optional.String**| Exact match for the application name of a permission. You may also use a comma-separated list to match on multiple applications. | 
 **resourceType** | **optional.String**| Exact match for the resource type name of a permission. You may also use a comma-separated list to match on multiple resource_types. | 
 **verb** | **optional.String**| Exact match for the operation verb name of a permission You may also use a comma-separated list to match on multiple verbs. | 
 **permission** | **optional.String**| Partial match for the aggregate permission value name of a permission object. | 
 **excludeGlobals** | **optional.String**| If set to &#39;true&#39;, this will exclude any permission with a global allowance on either &#39;application&#39;, &#39;resource_type&#39; or &#39;verb&#39;. The default is &#39;false&#39;. | [default to false]

### Return type

[**PermissionPagination**](PermissionPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

