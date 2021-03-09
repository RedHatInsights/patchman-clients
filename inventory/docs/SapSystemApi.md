# \SapSystemApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSystemProfileGetSapSids**](SapSystemApi.md#ApiSystemProfileGetSapSids) | **Get** /system_profile/sap_sids | get sap system values
[**ApiSystemProfileGetSapSystem**](SapSystemApi.md#ApiSystemProfileGetSapSystem) | **Get** /system_profile/sap_system | get sap system values



## ApiSystemProfileGetSapSids

> InlineResponse2002 ApiSystemProfileGetSapSids(ctx, optional)

get sap system values

Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiSystemProfileGetSapSidsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiSystemProfileGetSapSidsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Only include tags that match the given search string. The value is matched against namespace, key and value. | 
 **tags** | [**optional.Interface of []string**](string.md)| filters out hosts not tagged by the given tags | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **staleness** | [**optional.Interface of []string**](string.md)| Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **optional.String**| Filters out any host not registered with the specified service | 
 **filter** | [**optional.Interface of map[string]interface{}**](.md)| Filters hosts based on system_profile fields | 

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

> InlineResponse2002 ApiSystemProfileGetSapSystem(ctx, optional)

get sap system values

Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApiSystemProfileGetSapSystemOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiSystemProfileGetSapSystemOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | [**optional.Interface of []string**](string.md)| filters out hosts not tagged by the given tags | 
 **perPage** | **optional.Int32**| A number of items to return per page. | [default to 50]
 **page** | **optional.Int32**| A page number of the items to return. | [default to 1]
 **staleness** | [**optional.Interface of []string**](string.md)| Culling states of the hosts. Default: fresh,stale,unknown | [default to [&quot;fresh&quot;,&quot;stale&quot;,&quot;unknown&quot;]]
 **registeredWith** | **optional.String**| Filters out any host not registered with the specified service | 
 **filter** | [**optional.Interface of map[string]interface{}**](.md)| Filters hosts based on system_profile fields | 

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

