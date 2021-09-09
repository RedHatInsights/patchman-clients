# \DefaultApi

All URIs are relative to *http://localhost/api/vmaas/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmaasWebappAppCVEHandlerGetGet**](DefaultApi.md#VmaasWebappAppCVEHandlerGetGet) | **Get** /cves/{cve} | 
[**VmaasWebappAppCVEHandlerPostPost**](DefaultApi.md#VmaasWebappAppCVEHandlerPostPost) | **Post** /cves | 
[**VmaasWebappAppDBChangeHandlerGet**](DefaultApi.md#VmaasWebappAppDBChangeHandlerGet) | **Get** /dbchange | 
[**VmaasWebappAppErrataHandlerGetGet**](DefaultApi.md#VmaasWebappAppErrataHandlerGetGet) | **Get** /errata/{erratum} | 
[**VmaasWebappAppErrataHandlerPostPost**](DefaultApi.md#VmaasWebappAppErrataHandlerPostPost) | **Post** /errata | 
[**VmaasWebappAppHealthHandlerGet**](DefaultApi.md#VmaasWebappAppHealthHandlerGet) | **Get** /monitoring/health | Return API liveness status
[**VmaasWebappAppPackagesHandlerGetGet**](DefaultApi.md#VmaasWebappAppPackagesHandlerGetGet) | **Get** /packages/{nevra} | 
[**VmaasWebappAppPackagesHandlerPostPost**](DefaultApi.md#VmaasWebappAppPackagesHandlerPostPost) | **Post** /packages | 
[**VmaasWebappAppPatchesHandlerGetGet**](DefaultApi.md#VmaasWebappAppPatchesHandlerGetGet) | **Get** /patches/{nevra} | 
[**VmaasWebappAppPatchesHandlerPostPost**](DefaultApi.md#VmaasWebappAppPatchesHandlerPostPost) | **Post** /patches | 
[**VmaasWebappAppPkgtreeHandlerV3GetGet**](DefaultApi.md#VmaasWebappAppPkgtreeHandlerV3GetGet) | **Get** /pkgtree/{package_name} | 
[**VmaasWebappAppPkgtreeHandlerV3PostPost**](DefaultApi.md#VmaasWebappAppPkgtreeHandlerV3PostPost) | **Post** /pkgtree | 
[**VmaasWebappAppReadyHandlerGet**](DefaultApi.md#VmaasWebappAppReadyHandlerGet) | **Get** /monitoring/ready | Return API readiness status
[**VmaasWebappAppReposHandlerGetGet**](DefaultApi.md#VmaasWebappAppReposHandlerGetGet) | **Get** /repos/{repo} | 
[**VmaasWebappAppReposHandlerPostPost**](DefaultApi.md#VmaasWebappAppReposHandlerPostPost) | **Post** /repos | 
[**VmaasWebappAppUpdatesHandlerV3GetGet**](DefaultApi.md#VmaasWebappAppUpdatesHandlerV3GetGet) | **Get** /updates/{nevra} | 
[**VmaasWebappAppUpdatesHandlerV3PostPost**](DefaultApi.md#VmaasWebappAppUpdatesHandlerV3PostPost) | **Post** /updates | 
[**VmaasWebappAppVersionHandlerGet**](DefaultApi.md#VmaasWebappAppVersionHandlerGet) | **Get** /version | 
[**VmaasWebappAppVulnerabilitiesHandlerGetGet**](DefaultApi.md#VmaasWebappAppVulnerabilitiesHandlerGetGet) | **Get** /vulnerabilities/{nevra} | 
[**VmaasWebappAppVulnerabilitiesHandlerPostPost**](DefaultApi.md#VmaasWebappAppVulnerabilitiesHandlerPostPost) | **Post** /vulnerabilities | 



## VmaasWebappAppCVEHandlerGetGet

> CvesResponse VmaasWebappAppCVEHandlerGetGet(ctx, cve).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppCVEHandlerGetGet(context.Background(), cve).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppCVEHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppCVEHandlerGetGet`: CvesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppCVEHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cve** | **string** | CVE name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppCVEHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppCVEHandlerPostPost

> CvesResponse VmaasWebappAppCVEHandlerPostPost(ctx).CvesRequest(cvesRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppCVEHandlerPostPost(context.Background()).CvesRequest(cvesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppCVEHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppCVEHandlerPostPost`: CvesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppCVEHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppCVEHandlerPostPostRequest struct via the builder pattern


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


## VmaasWebappAppDBChangeHandlerGet

> DBChangeResponse VmaasWebappAppDBChangeHandlerGet(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppDBChangeHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppDBChangeHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppDBChangeHandlerGet`: DBChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppDBChangeHandlerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppDBChangeHandlerGetRequest struct via the builder pattern


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


## VmaasWebappAppErrataHandlerGetGet

> ErrataResponse VmaasWebappAppErrataHandlerGetGet(ctx, erratum).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppErrataHandlerGetGet(context.Background(), erratum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppErrataHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppErrataHandlerGetGet`: ErrataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppErrataHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**erratum** | **string** | Errata advisory name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppErrataHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppErrataHandlerPostPost

> ErrataResponse VmaasWebappAppErrataHandlerPostPost(ctx).ErrataRequest(errataRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppErrataHandlerPostPost(context.Background()).ErrataRequest(errataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppErrataHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppErrataHandlerPostPost`: ErrataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppErrataHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppErrataHandlerPostPostRequest struct via the builder pattern


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


## VmaasWebappAppHealthHandlerGet

> VmaasWebappAppHealthHandlerGet(ctx).Execute()

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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppHealthHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppHealthHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppHealthHandlerGetRequest struct via the builder pattern


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


## VmaasWebappAppPackagesHandlerGetGet

> PackagesResponse VmaasWebappAppPackagesHandlerGetGet(ctx, nevra).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPackagesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPackagesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPackagesHandlerGetGet`: PackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPackagesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPackagesHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppPackagesHandlerPostPost

> PackagesResponse VmaasWebappAppPackagesHandlerPostPost(ctx).PackagesRequest(packagesRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPackagesHandlerPostPost(context.Background()).PackagesRequest(packagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPackagesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPackagesHandlerPostPost`: PackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPackagesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPackagesHandlerPostPostRequest struct via the builder pattern


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


## VmaasWebappAppPatchesHandlerGetGet

> PatchesResponse VmaasWebappAppPatchesHandlerGetGet(ctx, nevra).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPatchesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPatchesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPatchesHandlerGetGet`: PatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPatchesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPatchesHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppPatchesHandlerPostPost

> PatchesResponse VmaasWebappAppPatchesHandlerPostPost(ctx).PatchesRequest(patchesRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPatchesHandlerPostPost(context.Background()).PatchesRequest(patchesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPatchesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPatchesHandlerPostPost`: PatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPatchesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPatchesHandlerPostPostRequest struct via the builder pattern


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


## VmaasWebappAppPkgtreeHandlerV3GetGet

> PkgtreeResponse VmaasWebappAppPkgtreeHandlerV3GetGet(ctx, packageName).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPkgtreeHandlerV3GetGet(context.Background(), packageName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPkgtreeHandlerV3GetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPkgtreeHandlerV3GetGet`: PkgtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPkgtreeHandlerV3GetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageName** | **string** | Package name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPkgtreeHandlerV3GetGetRequest struct via the builder pattern


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


## VmaasWebappAppPkgtreeHandlerV3PostPost

> PkgtreeResponse VmaasWebappAppPkgtreeHandlerV3PostPost(ctx).PkgtreeRequest(pkgtreeRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppPkgtreeHandlerV3PostPost(context.Background()).PkgtreeRequest(pkgtreeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppPkgtreeHandlerV3PostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppPkgtreeHandlerV3PostPost`: PkgtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppPkgtreeHandlerV3PostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppPkgtreeHandlerV3PostPostRequest struct via the builder pattern


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


## VmaasWebappAppReadyHandlerGet

> VmaasWebappAppReadyHandlerGet(ctx).Execute()

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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppReadyHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppReadyHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppReadyHandlerGetRequest struct via the builder pattern


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


## VmaasWebappAppReposHandlerGetGet

> ReposResponse VmaasWebappAppReposHandlerGetGet(ctx, repo).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppReposHandlerGetGet(context.Background(), repo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppReposHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppReposHandlerGetGet`: ReposResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppReposHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repo** | **string** | Repository name or POSIX regular expression pattern | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppReposHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppReposHandlerPostPost

> ReposResponse VmaasWebappAppReposHandlerPostPost(ctx).ReposRequest(reposRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppReposHandlerPostPost(context.Background()).ReposRequest(reposRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppReposHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppReposHandlerPostPost`: ReposResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppReposHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppReposHandlerPostPostRequest struct via the builder pattern


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


## VmaasWebappAppUpdatesHandlerV3GetGet

> UpdatesV2Response VmaasWebappAppUpdatesHandlerV3GetGet(ctx, nevra).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppUpdatesHandlerV3GetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppUpdatesHandlerV3GetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppUpdatesHandlerV3GetGet`: UpdatesV2Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppUpdatesHandlerV3GetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppUpdatesHandlerV3GetGetRequest struct via the builder pattern


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


## VmaasWebappAppUpdatesHandlerV3PostPost

> UpdatesV2Response VmaasWebappAppUpdatesHandlerV3PostPost(ctx).UpdatesV3Request(updatesV3Request).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppUpdatesHandlerV3PostPost(context.Background()).UpdatesV3Request(updatesV3Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppUpdatesHandlerV3PostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppUpdatesHandlerV3PostPost`: UpdatesV2Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppUpdatesHandlerV3PostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppUpdatesHandlerV3PostPostRequest struct via the builder pattern


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


## VmaasWebappAppVersionHandlerGet

> VmaasWebappAppVersionHandlerGet(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppVersionHandlerGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppVersionHandlerGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppVersionHandlerGetRequest struct via the builder pattern


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


## VmaasWebappAppVulnerabilitiesHandlerGetGet

> VulnerabilitiesResponse VmaasWebappAppVulnerabilitiesHandlerGetGet(ctx, nevra).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppVulnerabilitiesHandlerGetGet(context.Background(), nevra).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppVulnerabilitiesHandlerGetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppVulnerabilitiesHandlerGetGet`: VulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppVulnerabilitiesHandlerGetGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nevra** | **string** | Package NEVRA | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppVulnerabilitiesHandlerGetGetRequest struct via the builder pattern


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


## VmaasWebappAppVulnerabilitiesHandlerPostPost

> VulnerabilitiesResponse VmaasWebappAppVulnerabilitiesHandlerPostPost(ctx).VulnerabilitiesRequest(vulnerabilitiesRequest).Execute()





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
    resp, r, err := api_client.DefaultApi.VmaasWebappAppVulnerabilitiesHandlerPostPost(context.Background()).VulnerabilitiesRequest(vulnerabilitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VmaasWebappAppVulnerabilitiesHandlerPostPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmaasWebappAppVulnerabilitiesHandlerPostPost`: VulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VmaasWebappAppVulnerabilitiesHandlerPostPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVmaasWebappAppVulnerabilitiesHandlerPostPostRequest struct via the builder pattern


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

