# \DefaultApi

All URIs are relative to *http://localhost/api/vmaas/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCVEHandlerGetGet**](DefaultApi.md#AppCVEHandlerGetGet) | **Get** /cves/{cve} | 
[**AppCVEHandlerPostPost**](DefaultApi.md#AppCVEHandlerPostPost) | **Post** /cves | 
[**AppDBChangeHandlerGet**](DefaultApi.md#AppDBChangeHandlerGet) | **Get** /dbchange | 
[**AppErrataHandlerGetGet**](DefaultApi.md#AppErrataHandlerGetGet) | **Get** /errata/{erratum} | 
[**AppErrataHandlerPostPost**](DefaultApi.md#AppErrataHandlerPostPost) | **Post** /errata | 
[**AppHealthHandlerGet**](DefaultApi.md#AppHealthHandlerGet) | **Get** /monitoring/health | Return API liveness status
[**AppPackagesHandlerGetGet**](DefaultApi.md#AppPackagesHandlerGetGet) | **Get** /packages/{nevra} | 
[**AppPackagesHandlerPostPost**](DefaultApi.md#AppPackagesHandlerPostPost) | **Post** /packages | 
[**AppPatchesHandlerGetGet**](DefaultApi.md#AppPatchesHandlerGetGet) | **Get** /patches/{nevra} | 
[**AppPatchesHandlerPostPost**](DefaultApi.md#AppPatchesHandlerPostPost) | **Post** /patches | 
[**AppPkgtreeHandlerGetGet**](DefaultApi.md#AppPkgtreeHandlerGetGet) | **Get** /pkgtree/{package_name} | 
[**AppPkgtreeHandlerPostPost**](DefaultApi.md#AppPkgtreeHandlerPostPost) | **Post** /pkgtree | 
[**AppReadyHandlerGet**](DefaultApi.md#AppReadyHandlerGet) | **Get** /monitoring/ready | Return API readiness status
[**AppReposHandlerGetGet**](DefaultApi.md#AppReposHandlerGetGet) | **Get** /repos/{repo} | 
[**AppReposHandlerPostPost**](DefaultApi.md#AppReposHandlerPostPost) | **Post** /repos | 
[**AppUpdatesHandlerV3GetGet**](DefaultApi.md#AppUpdatesHandlerV3GetGet) | **Get** /updates/{nevra} | 
[**AppUpdatesHandlerV3PostPost**](DefaultApi.md#AppUpdatesHandlerV3PostPost) | **Post** /updates | 
[**AppVersionHandlerGet**](DefaultApi.md#AppVersionHandlerGet) | **Get** /version | 
[**AppVulnerabilitiesHandlerGetGet**](DefaultApi.md#AppVulnerabilitiesHandlerGetGet) | **Get** /vulnerabilities/{nevra} | 
[**AppVulnerabilitiesHandlerPostPost**](DefaultApi.md#AppVulnerabilitiesHandlerPostPost) | **Post** /vulnerabilities | 



## AppCVEHandlerGetGet

> CvesResponse AppCVEHandlerGetGet(ctx, cve).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cve := ""CVE-2017-5715, CVE-2017-571[1-5], CVE-2017-5.*"" // string | CVE name or POSIX regular expression pattern

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppCVEHandlerGetGet(context.Background(), cve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppCVEHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCVEHandlerGetGet`: CvesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppCVEHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cve** | **string** | CVE name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppCVEHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> CvesResponse AppCVEHandlerPostPost(ctx).CvesRequest(cvesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cvesRequest := *openapiclient.NewCvesRequest([]string{"CVE-2017-57.*"}) // CvesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppCVEHandlerPostPost(context.Background()).CvesRequest(cvesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppCVEHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppCVEHandlerPostPost`: CvesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppCVEHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCVEHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cvesRequest** | [**CvesRequest**](CvesRequest.md) |  | 

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

> DBChangeResponse AppDBChangeHandlerGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppDBChangeHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppDBChangeHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppDBChangeHandlerGet`: DBChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppDBChangeHandlerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppDBChangeHandlerGetRequest struct via the builder pattern


### Return type

[**DBChangeResponse**](DBChangeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppErrataHandlerGetGet

> ErrataResponse AppErrataHandlerGetGet(ctx, erratum).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    erratum := ""RHSA-2018:0512, RHSA-2018:051[1-5], RH.*"" // string | Errata advisory name or POSIX regular expression pattern

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppErrataHandlerGetGet(context.Background(), erratum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppErrataHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppErrataHandlerGetGet`: ErrataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppErrataHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**erratum** | **string** | Errata advisory name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppErrataHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ErrataResponse AppErrataHandlerPostPost(ctx).ErrataRequest(errataRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    errataRequest := *openapiclient.NewErrataRequest([]string{"RHSA-2018:05.*"}) // ErrataRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppErrataHandlerPostPost(context.Background()).ErrataRequest(errataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppErrataHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppErrataHandlerPostPost`: ErrataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppErrataHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppErrataHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **errataRequest** | [**ErrataRequest**](ErrataRequest.md) |  | 

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

> AppHealthHandlerGet(ctx).Execute()

Return API liveness status

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppHealthHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppHealthHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppHealthHandlerGetRequest struct via the builder pattern


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

> PackagesResponse AppPackagesHandlerGetGet(ctx, nevra).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nevra := ""kernel-2.6.32-696.20.1.el6.x86_64"" // string | Package NEVRA

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPackagesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPackagesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPackagesHandlerGetGet`: PackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPackagesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPackagesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PackagesResponse AppPackagesHandlerPostPost(ctx).PackagesRequest(packagesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    packagesRequest := *openapiclient.NewPackagesRequest([]string{"kernel-2.6.32-696.20.1.el6.x86_64"}) // PackagesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPackagesHandlerPostPost(context.Background()).PackagesRequest(packagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPackagesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPackagesHandlerPostPost`: PackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPackagesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPackagesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packagesRequest** | [**PackagesRequest**](PackagesRequest.md) |  | 

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

> PatchesResponse AppPatchesHandlerGetGet(ctx, nevra).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nevra := ""kernel-2.6.32-696.20.1.el6.x86_64"" // string | Package NEVRA

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPatchesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPatchesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPatchesHandlerGetGet`: PatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPatchesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPatchesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PatchesResponse AppPatchesHandlerPostPost(ctx).PatchesRequest(patchesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    patchesRequest := *openapiclient.NewPatchesRequest([]string{"kernel-2.6.32-696.20.1.el6.x86_64"}) // PatchesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPatchesHandlerPostPost(context.Background()).PatchesRequest(patchesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPatchesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPatchesHandlerPostPost`: PatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPatchesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPatchesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchesRequest** | [**PatchesRequest**](PatchesRequest.md) |  | 

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

> PkgtreeResponse AppPkgtreeHandlerGetGet(ctx, packageName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    packageName := ""kernel-rt"" // string | Package name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPkgtreeHandlerGetGet(context.Background(), packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPkgtreeHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPkgtreeHandlerGetGet`: PkgtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPkgtreeHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageName** | **string** | Package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPkgtreeHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PkgtreeResponse AppPkgtreeHandlerPostPost(ctx).PkgtreeRequest(pkgtreeRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pkgtreeRequest := *openapiclient.NewPkgtreeRequest([]string{"kernel-rt"}) // PkgtreeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppPkgtreeHandlerPostPost(context.Background()).PkgtreeRequest(pkgtreeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppPkgtreeHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPkgtreeHandlerPostPost`: PkgtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppPkgtreeHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPkgtreeHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pkgtreeRequest** | [**PkgtreeRequest**](PkgtreeRequest.md) |  | 

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

> AppReadyHandlerGet(ctx).Execute()

Return API readiness status

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppReadyHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppReadyHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppReadyHandlerGetRequest struct via the builder pattern


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

> ReposResponse AppReposHandlerGetGet(ctx, repo).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    repo := ""rhel-6-server-rpms OR rhel-[4567]-.*-rpms OR rhel-\\d-server-rpms"" // string | Repository name or POSIX regular expression pattern

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppReposHandlerGetGet(context.Background(), repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppReposHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppReposHandlerGetGet`: ReposResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppReposHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repository name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppReposHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> ReposResponse AppReposHandlerPostPost(ctx).ReposRequest(reposRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reposRequest := *openapiclient.NewReposRequest([]string{"rhel-6-server-rpms"}) // ReposRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppReposHandlerPostPost(context.Background()).ReposRequest(reposRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppReposHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppReposHandlerPostPost`: ReposResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppReposHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppReposHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reposRequest** | [**ReposRequest**](ReposRequest.md) |  | 

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


## AppUpdatesHandlerV3GetGet

> UpdatesV2Response AppUpdatesHandlerV3GetGet(ctx, nevra).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nevra := ""kernel-2.6.32-696.20.1.el6.x86_64"" // string | Package NEVRA

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppUpdatesHandlerV3GetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppUpdatesHandlerV3GetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppUpdatesHandlerV3GetGet`: UpdatesV2Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppUpdatesHandlerV3GetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppUpdatesHandlerV3GetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> UpdatesV2Response AppUpdatesHandlerV3PostPost(ctx).UpdatesV3Request(updatesV3Request).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updatesV3Request := *openapiclient.NewUpdatesV3Request([]string{"kernel-2.6.32-696.20.1.el6.x86_64"}) // UpdatesV3Request | Input json (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppUpdatesHandlerV3PostPost(context.Background()).UpdatesV3Request(updatesV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppUpdatesHandlerV3PostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppUpdatesHandlerV3PostPost`: UpdatesV2Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppUpdatesHandlerV3PostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppUpdatesHandlerV3PostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatesV3Request** | [**UpdatesV3Request**](UpdatesV3Request.md) | Input json | 

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

> AppVersionHandlerGet(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppVersionHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppVersionHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppVersionHandlerGetRequest struct via the builder pattern


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

> VulnerabilitiesResponse AppVulnerabilitiesHandlerGetGet(ctx, nevra).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nevra := ""kernel-2.6.32-696.20.1.el6.x86_64"" // string | Package NEVRA

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppVulnerabilitiesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppVulnerabilitiesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVulnerabilitiesHandlerGetGet`: VulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppVulnerabilitiesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppVulnerabilitiesHandlerGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> VulnerabilitiesResponse AppVulnerabilitiesHandlerPostPost(ctx).VulnerabilitiesRequest(vulnerabilitiesRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    vulnerabilitiesRequest := *openapiclient.NewVulnerabilitiesRequest([]string{"kernel-2.6.32-696.20.1.el6.x86_64"}) // VulnerabilitiesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AppVulnerabilitiesHandlerPostPost(context.Background()).VulnerabilitiesRequest(vulnerabilitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AppVulnerabilitiesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppVulnerabilitiesHandlerPostPost`: VulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AppVulnerabilitiesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppVulnerabilitiesHandlerPostPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vulnerabilitiesRequest** | [**VulnerabilitiesRequest**](VulnerabilitiesRequest.md) |  | 

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

