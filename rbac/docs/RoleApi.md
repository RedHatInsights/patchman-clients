# \RoleApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoles**](RoleApi.md#CreateRoles) | **Post** /roles/ | Create a roles for a tenant
[**DeleteRole**](RoleApi.md#DeleteRole) | **Delete** /roles/{uuid}/ | Delete a role in the tenant
[**GetRole**](RoleApi.md#GetRole) | **Get** /roles/{uuid}/ | Get a role in the tenant
[**GetRoleAccess**](RoleApi.md#GetRoleAccess) | **Get** /roles/{uuid}/access/ | Get access for a role in the tenant
[**ListRoles**](RoleApi.md#ListRoles) | **Get** /roles/ | List the roles for a tenant
[**UpdateRole**](RoleApi.md#UpdateRole) | **Put** /roles/{uuid}/ | Update a Role in the tenant



## CreateRoles

> RoleWithAccess CreateRoles(ctx, roleIn)

Create a roles for a tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleIn** | [**RoleIn**](RoleIn.md)| Role to create | 

### Return type

[**RoleWithAccess**](RoleWithAccess.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, uuid)

Delete a role in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of role to delete | 

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


## GetRole

> RoleWithAccess GetRole(ctx, uuid)

Get a role in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of role to get | 

### Return type

[**RoleWithAccess**](RoleWithAccess.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAccess

> AccessPagination GetRoleAccess(ctx, uuid, optional)

Get access for a role in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of the role | 
 **optional** | ***GetRoleAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRoleAccessOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ListRoles

> RolePaginationDynamic ListRoles(ctx, optional)

List the roles for a tenant

By default, responses are sorted in ascending order by role name

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRolesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **name** | **optional.String**| Parameter for filtering resource by name using string contains search. | 
 **scope** | **optional.String**| Parameter for filtering resource by scope. | [default to account]
 **orderBy** | **optional.String**| Parameter for ordering resource by value. For inverse ordering, supply &#39;-&#39; before the param value, such as: ?order_by&#x3D;-name | 
 **addFields** | [**optional.Interface of []string**](string.md)| Parameter for add list of fields to display for roles. | 
 **username** | **optional.String**| Unique username of the principal to obtain roles for (only available for admins, and if supplied, takes precedence over the identity header). | 

### Return type

[**RolePaginationDynamic**](RolePaginationDynamic.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> UpdateRole(ctx, uuid, roleWithAccess)

Update a Role in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of role to update | 
**roleWithAccess** | [**RoleWithAccess**](RoleWithAccess.md)| Update to a Role | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

