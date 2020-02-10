# \GroupApi

All URIs are relative to *http://localhost/api/rbac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrincipalToGroup**](GroupApi.md#AddPrincipalToGroup) | **Post** /groups/{uuid}/principals/ | Add a principal to a group in the tenant
[**AddRoleToGroup**](GroupApi.md#AddRoleToGroup) | **Post** /groups/{uuid}/roles/ | Add a role to a group in the tenant
[**CreateGroup**](GroupApi.md#CreateGroup) | **Post** /groups/ | Create a group in a tenant
[**DeleteGroup**](GroupApi.md#DeleteGroup) | **Delete** /groups/{uuid}/ | Delete a group in the tenant
[**DeletePrincipalFromGroup**](GroupApi.md#DeletePrincipalFromGroup) | **Delete** /groups/{uuid}/principals/ | Remove a principal from a group in the tenant
[**DeleteRoleFromGroup**](GroupApi.md#DeleteRoleFromGroup) | **Delete** /groups/{uuid}/roles/ | Remove a role from a group in the tenant
[**GetGroup**](GroupApi.md#GetGroup) | **Get** /groups/{uuid}/ | Get a group in the tenant
[**ListGroups**](GroupApi.md#ListGroups) | **Get** /groups/ | List the groups for a tenant
[**ListRolesForGroup**](GroupApi.md#ListRolesForGroup) | **Get** /groups/{uuid}/roles/ | List the roles for a group in the tenant
[**UpdateGroup**](GroupApi.md#UpdateGroup) | **Put** /groups/{uuid}/ | Udate a group in the tenant



## AddPrincipalToGroup

> GroupWithPrincipalsAndRoles AddPrincipalToGroup(ctx, uuid, groupPrincipalIn)

Add a principal to a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to update | 
**groupPrincipalIn** | [**GroupPrincipalIn**](GroupPrincipalIn.md)| Principal to add to a group | 

### Return type

[**GroupWithPrincipalsAndRoles**](GroupWithPrincipalsAndRoles.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRoleToGroup

> InlineResponse200 AddRoleToGroup(ctx, uuid, groupRoleIn)

Add a role to a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to update | 
**groupRoleIn** | [**GroupRoleIn**](GroupRoleIn.md)| Role to add to a group | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> GroupOut CreateGroup(ctx, group)

Create a group in a tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**group** | [**Group**](Group.md)| Group to create in tenant | 

### Return type

[**GroupOut**](GroupOut.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> DeleteGroup(ctx, uuid)

Delete a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to delete | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePrincipalFromGroup

> DeletePrincipalFromGroup(ctx, uuid, usernames)

Remove a principal from a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to update | 
**usernames** | **string**| A comma separated list of usernames for principals to remove from the group | 

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


## DeleteRoleFromGroup

> DeleteRoleFromGroup(ctx, uuid, roles)

Remove a role from a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to update | 
**roles** | **string**| A comma separated list of role UUIDs for roles to remove from the group | 

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


## GetGroup

> GroupWithPrincipalsAndRoles GetGroup(ctx, uuid)

Get a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to get | 

### Return type

[**GroupWithPrincipalsAndRoles**](GroupWithPrincipalsAndRoles.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> GroupPagination ListGroups(ctx, optional)

List the groups for a tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGroupsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]
 **name** | **optional.String**| Parameter for filtering resource by name using string contains search. | 
 **scope** | **optional.String**| Parameter for filtering resource by scope. | [default to account]
 **username** | **optional.String**| A username for a principal to filter for groups | 
 **orderBy** | **optional.String**| Parameter for ordering resource by value. | 

### Return type

[**GroupPagination**](GroupPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRolesForGroup

> GroupRolesPagination ListRolesForGroup(ctx, uuid, optional)

List the roles for a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group | 
 **optional** | ***ListRolesForGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRolesForGroupOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exclude** | **optional.Bool**| If this is set to true, the result would be roles excluding the ones in the group | [default to false]
 **limit** | **optional.Int32**| Parameter for selecting the amount of data returned. | [default to 10]
 **offset** | **optional.Int32**| Parameter for selecting the offset of data. | [default to 0]

### Return type

[**GroupRolesPagination**](GroupRolesPagination.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> GroupOut UpdateGroup(ctx, uuid, group)

Udate a group in the tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | [**string**](.md)| ID of group to update | 
**group** | [**Group**](Group.md)| Group to update in tenant | 

### Return type

[**GroupOut**](GroupOut.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

