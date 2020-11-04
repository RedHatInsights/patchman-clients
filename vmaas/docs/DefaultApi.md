# \DefaultApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCVEHandlerGetGet**](DefaultApi.md#AppCVEHandlerGetGet) | **Get** /v1/cves/{cve} | 
[**AppCVEHandlerPostPost**](DefaultApi.md#AppCVEHandlerPostPost) | **Post** /v1/cves | 
[**AppDBChangeHandlerGet**](DefaultApi.md#AppDBChangeHandlerGet) | **Get** /v1/dbchange | 
[**AppErrataHandlerGetGet**](DefaultApi.md#AppErrataHandlerGetGet) | **Get** /v1/errata/{erratum} | 
[**AppErrataHandlerPostPost**](DefaultApi.md#AppErrataHandlerPostPost) | **Post** /v1/errata | 
[**AppHealthHandlerGet**](DefaultApi.md#AppHealthHandlerGet) | **Get** /v1/monitoring/health | Return API liveness status
[**AppPackagesHandlerGetGet**](DefaultApi.md#AppPackagesHandlerGetGet) | **Get** /v1/packages/{nevra} | 
[**AppPackagesHandlerPostPost**](DefaultApi.md#AppPackagesHandlerPostPost) | **Post** /v1/packages | 
[**AppPatchesHandlerGetGet**](DefaultApi.md#AppPatchesHandlerGetGet) | **Get** /v1/patches/{nevra} | 
[**AppPatchesHandlerPostPost**](DefaultApi.md#AppPatchesHandlerPostPost) | **Post** /v1/patches | 
[**AppPkgtreeHandlerGetGet**](DefaultApi.md#AppPkgtreeHandlerGetGet) | **Get** /v1/pkgtree/{package_name} | 
[**AppPkgtreeHandlerPostPost**](DefaultApi.md#AppPkgtreeHandlerPostPost) | **Post** /v1/pkgtree | 
[**AppReadyHandlerGet**](DefaultApi.md#AppReadyHandlerGet) | **Get** /v1/monitoring/ready | Return API readiness status
[**AppReposHandlerGetGet**](DefaultApi.md#AppReposHandlerGetGet) | **Get** /v1/repos/{repo} | 
[**AppReposHandlerPostPost**](DefaultApi.md#AppReposHandlerPostPost) | **Post** /v1/repos | 
[**AppUpdatesHandlerGetGet**](DefaultApi.md#AppUpdatesHandlerGetGet) | **Get** /v1/updates/{nevra} | 
[**AppUpdatesHandlerPostPost**](DefaultApi.md#AppUpdatesHandlerPostPost) | **Post** /v1/updates | 
[**AppUpdatesHandlerV2GetGet**](DefaultApi.md#AppUpdatesHandlerV2GetGet) | **Get** /v2/updates/{nevra} | 
[**AppUpdatesHandlerV2PostPost**](DefaultApi.md#AppUpdatesHandlerV2PostPost) | **Post** /v2/updates | 
[**AppUpdatesHandlerV3GetGet**](DefaultApi.md#AppUpdatesHandlerV3GetGet) | **Get** /v3/updates/{nevra} | 
[**AppUpdatesHandlerV3PostPost**](DefaultApi.md#AppUpdatesHandlerV3PostPost) | **Post** /v3/updates | 
[**AppVersionHandlerGet**](DefaultApi.md#AppVersionHandlerGet) | **Get** /v1/version | 
[**AppVulnerabilitiesHandlerGetGet**](DefaultApi.md#AppVulnerabilitiesHandlerGetGet) | **Get** /v1/vulnerabilities/{nevra} | 
[**AppVulnerabilitiesHandlerPostPost**](DefaultApi.md#AppVulnerabilitiesHandlerPostPost) | **Post** /v1/vulnerabilities | 



## AppCVEHandlerGetGet

> CvesResponse AppCVEHandlerGetGet(ctx, cve)



Get details about CVEs. It is possible to use POSIX regular expression as a pattern for CVE names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cve** | **string**| CVE name or POSIX regular expression pattern | 

### Return type

[**CvesResponse**](CvesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppCVEHandlerPostPost

> CvesResponse AppCVEHandlerPostPost(ctx, optional)



Get details about CVEs with additional parameters. As a \"cve_list\" parameter a complete list of CVE names can be provided OR one POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppCVEHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppCVEHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cvesRequest** | [**optional.Interface of CvesRequest**](CvesRequest.md)|  | 

### Return type

[**CvesResponse**](CvesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDBChangeHandlerGet

> DbChangeResponse AppDBChangeHandlerGet(ctx, )



Get last-updated-times for VMaaS DB

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**DbChangeResponse**](DBChangeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppErrataHandlerGetGet

> ErrataResponse AppErrataHandlerGetGet(ctx, erratum)



Get details about errata. It is possible to use POSIX regular expression as a pattern for errata names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**erratum** | **string**| Errata advisory name or POSIX regular expression pattern | 

### Return type

[**ErrataResponse**](ErrataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppErrataHandlerPostPost

> ErrataResponse AppErrataHandlerPostPost(ctx, optional)



Get details about errata with additional parameters. \"errata_list\" parameter can be either a list of errata names OR a single POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppErrataHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppErrataHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **errataRequest** | [**optional.Interface of ErrataRequest**](ErrataRequest.md)|  | 

### Return type

[**ErrataResponse**](ErrataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppHealthHandlerGet

> AppHealthHandlerGet(ctx, )

Return API liveness status

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPackagesHandlerGetGet

> PackagesResponse AppPackagesHandlerGetGet(ctx, nevra)



Get details about packages.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**PackagesResponse**](PackagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPackagesHandlerPostPost

> PackagesResponse AppPackagesHandlerPostPost(ctx, optional)



Get details about packages. \"package_list\" must be a list of package NEVRAs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppPackagesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppPackagesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packagesRequest** | [**optional.Interface of PackagesRequest**](PackagesRequest.md)|  | 

### Return type

[**PackagesResponse**](PackagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## AppReadyHandlerGet

> AppReadyHandlerGet(ctx, )

Return API readiness status

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppReposHandlerGetGet

> ReposResponse AppReposHandlerGetGet(ctx, repo)



Get details about a repository or repository-expression. It is allowed to use POSIX regular expression as a pattern for repository names.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string**| Repository name or POSIX regular expression pattern | 

### Return type

[**ReposResponse**](ReposResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppReposHandlerPostPost

> ReposResponse AppReposHandlerPostPost(ctx, optional)



Get details about list of repositories. \"repository_list\" can be either a list of repository names, OR a single POSIX regular expression.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppReposHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppReposHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reposRequest** | [**optional.Interface of ReposRequest**](ReposRequest.md)|  | 

### Return type

[**ReposResponse**](ReposResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerGetGet

> UpdatesResponse AppUpdatesHandlerGetGet(ctx, nevra)



List security updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesResponse**](UpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerPostPost

> UpdatesResponse AppUpdatesHandlerPostPost(ctx, optional)



List security updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesRequest** | [**optional.Interface of UpdatesRequest**](UpdatesRequest.md)| Input json | 

### Return type

[**UpdatesResponse**](UpdatesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV2GetGet

> UpdatesV2Response AppUpdatesHandlerV2GetGet(ctx, nevra)



List security updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV2PostPost

> UpdatesV2Response AppUpdatesHandlerV2PostPost(ctx, optional)



List security updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerV2PostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerV2PostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesRequest** | [**optional.Interface of UpdatesRequest**](UpdatesRequest.md)| Input json | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV3GetGet

> UpdatesV2Response AppUpdatesHandlerV3GetGet(ctx, nevra)



List all updates for single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdatesHandlerV3PostPost

> UpdatesV2Response AppUpdatesHandlerV3PostPost(ctx, optional)



List all updates for list of package NEVRAs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppUpdatesHandlerV3PostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppUpdatesHandlerV3PostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesV3Request** | [**optional.Interface of UpdatesV3Request**](UpdatesV3Request.md)| Input json | 

### Return type

[**UpdatesV2Response**](UpdatesV2Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVersionHandlerGet

> AppVersionHandlerGet(ctx, )



Get version of application

### Required Parameters

This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVulnerabilitiesHandlerGetGet

> VulnerabilitiesResponse AppVulnerabilitiesHandlerGetGet(ctx, nevra)



List of applicable CVEs for a single package NEVRA

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string**| Package NEVRA | 

### Return type

[**VulnerabilitiesResponse**](VulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppVulnerabilitiesHandlerPostPost

> VulnerabilitiesResponse AppVulnerabilitiesHandlerPostPost(ctx, optional)



List of applicable CVEs to a package list.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppVulnerabilitiesHandlerPostPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AppVulnerabilitiesHandlerPostPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilitiesRequest** | [**optional.Interface of VulnerabilitiesRequest**](VulnerabilitiesRequest.md)|  | 

### Return type

[**VulnerabilitiesResponse**](VulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

