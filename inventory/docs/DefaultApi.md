# \DefaultApi

All URIs are relative to *http://localhost/api/inventory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSystemProfileValidateSchema**](DefaultApi.md#ApiSystemProfileValidateSchema) | **Post** /system_profile/validate_schema | validate system profile schema



## ApiSystemProfileValidateSchema

> ApiSystemProfileValidateSchema(ctx, repoBranch, optional)

validate system profile schema

Required permissions: inventory:hosts:read

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoBranch** | **string**| The branch of the inventory-schemas repo to use | 
 **optional** | ***ApiSystemProfileValidateSchemaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiSystemProfileValidateSchemaOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repoFork** | **optional.String**| The fork of the inventory-schemas repo to use | 
 **days** | **optional.Int32**| How many days worth of kafka messages to validate | 

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

