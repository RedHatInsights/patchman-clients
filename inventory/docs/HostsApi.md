# \HostsApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiHostDeleteById**](HostsApi.md#ApiHostDeleteById) | **Delete** /hosts/{host_id_list} | Delete hosts by IDs
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

> ApiHostDeleteById(ctx, hostIdList, optional)

Delete hosts by IDs

Delete hosts by IDs <br /><br /> Required permissions: inventory:hosts:write

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostDeleteByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostDeleteByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **optional.String**| Filter by branch_id | 

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

> AHostInventoryQueryResult ApiHostGetHostById(ctx, hostIdList, optional)

Find hosts by their IDs

Find one or more hosts by their ID. <br /><br /> Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchId** | **optional.String**| Filter by branch_id | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**AHostInventoryQueryResult**](A_Host_Inventory_query_result.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostList

> AHostInventoryQueryResult ApiHostGetHostList(ctx, optional)

Read the entire list of hosts

Read the entire list of all hosts available to the account. <br /><br /> Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiHostGetHostListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **optional.String**| A part of a searched host’s display name. | 
 **fqdn** | **optional.String**| Filter by a host&#39;s FQDN | 
 **hostnameOrId** | **optional.String**| Search for a host by display_name, fqdn, id | 
 **insightsId** | [**optional.Interface of string**](.md)| Search for a host by insights_id | 
 **branchId** | **optional.String**| Filter by branch_id | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **staleness** | [**optional.Interface of []string**](string.md)| Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **tags** | [**optional.Interface of []string**](string.md)| filters out hosts not tagged by the given tags | 
 **registeredWith** | **optional.String**| Filters out any host not registered with the specified service | 
 **filter** | [**optional.Interface of map[string]interface{}**](.md)| Filters hosts based on system_profile fields | 

### Return type

[**AHostInventoryQueryResult**](A_Host_Inventory_query_result.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostSystemProfileById

> AHostSystemProfileQueryResult ApiHostGetHostSystemProfileById(ctx, hostIdList, optional)

Return one or more hosts system profile

Find one or more hosts by their ID and return the id and system profile <br /><br /> Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostSystemProfileByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostSystemProfileByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **branchId** | **optional.String**| Filter by branch_id | 
 **fields** | [**optional.Interface of map[string]interface{}**](.md)| Fetches only mentioned system_profile fields | 

### Return type

[**AHostSystemProfileQueryResult**](A_host_system_profile_query_result.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTagCount

> InlineResponse2001 ApiHostGetHostTagCount(ctx, hostIdList, optional)

Get the number of tags on a host

Get the number of tags on a host <br /><br /> Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostTagCountOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostTagCountOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostGetHostTags

> InlineResponse200 ApiHostGetHostTags(ctx, hostIdList, optional)

Get the tags on a host

Get the tags on a host <br /><br /> Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
 **optional** | ***ApiHostGetHostTagsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostGetHostTagsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **orderBy** | **optional.String**| Ordering field name | 
 **orderHow** | **optional.String**| Direction of the ordering, defaults to ASC for display_name and to DESC for updated | 
 **search** | **optional.String**| Only include tags that match the given search string. The value is matched against namespace, key and value. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostHostCheckin

> map[string]interface{} ApiHostHostCheckin(ctx, uNKNOWNBASETYPE)

Update staleness timestamps for a host matching the provided facts

Finds a host and updates its staleness timestamps. It uses the supplied canonical facts to determine which host to update. By default, the staleness timestamp is set to 1 hour from when the request is received; however, this can be overridden by supplying the interval. <br /><br /> Required permissions: inventory:hosts:write

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uNKNOWNBASETYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)| Data required to create a check-in record for a host. | 

### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiHostMergeFacts

> ApiHostMergeFacts(ctx, hostIdList, namespace, body, optional)

Merge facts under a namespace

Merge one or multiple hosts facts under a namespace. <br /><br /> Required permissions: inventory:hosts:write

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**namespace** | **string**| A namespace of the merged facts. | 
**body** | **map[string]interface{}**| A dictionary with the new facts to merge with the original ones. | 
 **optional** | ***ApiHostMergeFactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostMergeFactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branchId** | **optional.String**| Filter by branch_id | 

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

> ApiHostPatchById(ctx, hostIdList, hostData, optional)

Update a host

Update a host <br /><br /> Required permissions: inventory:hosts:write

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**hostData** | [**HostData**](HostData.md)|  | 
 **optional** | ***ApiHostPatchByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostPatchByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **branchId** | **optional.String**| Filter by branch_id | 

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

> ApiHostReplaceFacts(ctx, hostIdList, namespace, body, optional)

Replace facts under a namespace

Replace facts under a namespace <br /><br /> Required permissions: inventory:hosts:write

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostIdList** | [**[]string**](string.md)| A comma separated list of host IDs. | 
**namespace** | **string**| A namespace of the merged facts. | 
**body** | **map[string]interface{}**| A dictionary with the new facts to replace the original ones. | 
 **optional** | ***ApiHostReplaceFactsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiHostReplaceFactsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **branchId** | **optional.String**| Filter by branch_id | 

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

