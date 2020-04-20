# \PkgtreeApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPkgtreeHandlerGetGet**](PkgtreeApi.md#AppPkgtreeHandlerGetGet) | **Get** /v1/pkgtree/{package_name} | 
[**AppPkgtreeHandlerPostPost**](PkgtreeApi.md#AppPkgtreeHandlerPostPost) | **Post** /v1/pkgtree | 



## AppPkgtreeHandlerGetGet

> PkgtreeResponse AppPkgtreeHandlerGetGet(ctx, packageName)



Get package NEVRAs tree for a single package name.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageName** | **string**| Package name | 

### Return type

[**PkgtreeResponse**](PkgtreeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPkgtreeHandlerPostPost

> PkgtreeResponse AppPkgtreeHandlerPostPost(ctx, optional)



Get package NEVRAs trees for package names. \"package_name_list\" must be a list of package names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPkgtreeHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPkgtreeHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkgtreeRequest** | [**optional.Interface of PkgtreeRequest**](PkgtreeRequest.md)|  | 

### Return type

[**PkgtreeResponse**](PkgtreeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

