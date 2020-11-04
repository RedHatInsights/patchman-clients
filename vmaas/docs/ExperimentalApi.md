# \ExperimentalApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#AppRPMPkgNamesHandlerGetGet) | **Get** /v1/package_names/rpms/{rpm} | 
[**AppRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#AppRPMPkgNamesHandlerPostPost) | **Post** /v1/package_names/rpms | 
[**AppSRPMPkgNamesHandlerGetGet**](ExperimentalApi.md#AppSRPMPkgNamesHandlerGetGet) | **Get** /v1/package_names/srpms/{srpm} | 
[**AppSRPMPkgNamesHandlerPostPost**](ExperimentalApi.md#AppSRPMPkgNamesHandlerPostPost) | **Post** /v1/package_names/srpms | 



## AppRPMPkgNamesHandlerGetGet

> RpmPkgNamesResponse AppRPMPkgNamesHandlerGetGet(ctx, rpm)



List of content sets by given rpm name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpm** | **string**| Package name | 

### Return type

[**RpmPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppRPMPkgNamesHandlerPostPost

> RpmPkgNamesResponse AppRPMPkgNamesHandlerPostPost(ctx, optional)



List of content sets by given rpm name and content set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppRPMPkgNamesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppRPMPkgNamesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rpmPkgNamesRequest** | [**optional.Interface of RpmPkgNamesRequest**](RpmPkgNamesRequest.md)|  | 

### Return type

[**RpmPkgNamesResponse**](RPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSRPMPkgNamesHandlerGetGet

> SrpmPkgNamesResponse AppSRPMPkgNamesHandlerGetGet(ctx, srpm)



List of content sets with associated rpm names by given srpm.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**srpm** | **string**| Source package name | 

### Return type

[**SrpmPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppSRPMPkgNamesHandlerPostPost

> SrpmPkgNamesResponse AppSRPMPkgNamesHandlerPostPost(ctx, optional)



List of content sets with associated rpm names by given srpm and content set.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppSRPMPkgNamesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppSRPMPkgNamesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **srpmPkgNamesRequest** | [**optional.Interface of SrpmPkgNamesRequest**](SrpmPkgNamesRequest.md)|  | 

### Return type

[**SrpmPkgNamesResponse**](SRPMPkgNamesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

