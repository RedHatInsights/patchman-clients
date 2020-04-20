# \PatchesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPatchesHandlerGetGet**](PatchesApi.md#AppPatchesHandlerGetGet) | **Get** /v1/patches/{nevra} | 
[**AppPatchesHandlerPostPost**](PatchesApi.md#AppPatchesHandlerPostPost) | **Post** /v1/patches | 



## AppPatchesHandlerGetGet

> PatchesResponse AppPatchesHandlerGetGet(ctx, nevra)



List of applicable CVEs for a single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**PatchesResponse**](PatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPatchesHandlerPostPost

> PatchesResponse AppPatchesHandlerPostPost(ctx, optional)



List of applicable errata to a package list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPatchesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPatchesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchesRequest** | [**optional.Interface of PatchesRequest**](PatchesRequest.md)|  | 

### Return type

[**PatchesResponse**](PatchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

