# Go API client for vmaas

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.32.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./vmaas"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/vmaas/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**VmaasWebappAppCVEHandlerGetGet**](docs/DefaultApi.md#vmaaswebappappcvehandlergetget) | **Get** /cves/{cve} | 
*DefaultApi* | [**VmaasWebappAppCVEHandlerPostPost**](docs/DefaultApi.md#vmaaswebappappcvehandlerpostpost) | **Post** /cves | 
*DefaultApi* | [**VmaasWebappAppDBChangeHandlerGet**](docs/DefaultApi.md#vmaaswebappappdbchangehandlerget) | **Get** /dbchange | 
*DefaultApi* | [**VmaasWebappAppErrataHandlerGetGet**](docs/DefaultApi.md#vmaaswebappapperratahandlergetget) | **Get** /errata/{erratum} | 
*DefaultApi* | [**VmaasWebappAppErrataHandlerPostPost**](docs/DefaultApi.md#vmaaswebappapperratahandlerpostpost) | **Post** /errata | 
*DefaultApi* | [**VmaasWebappAppHealthHandlerGet**](docs/DefaultApi.md#vmaaswebappapphealthhandlerget) | **Get** /monitoring/health | Return API liveness status
*DefaultApi* | [**VmaasWebappAppPackagesHandlerGetGet**](docs/DefaultApi.md#vmaaswebappapppackageshandlergetget) | **Get** /packages/{nevra} | 
*DefaultApi* | [**VmaasWebappAppPackagesHandlerPostPost**](docs/DefaultApi.md#vmaaswebappapppackageshandlerpostpost) | **Post** /packages | 
*DefaultApi* | [**VmaasWebappAppPatchesHandlerGetGet**](docs/DefaultApi.md#vmaaswebappapppatcheshandlergetget) | **Get** /patches/{nevra} | 
*DefaultApi* | [**VmaasWebappAppPatchesHandlerPostPost**](docs/DefaultApi.md#vmaaswebappapppatcheshandlerpostpost) | **Post** /patches | 
*DefaultApi* | [**VmaasWebappAppPkgListHandlerPostPost**](docs/DefaultApi.md#vmaaswebappapppkglisthandlerpostpost) | **Post** /pkglist | 
*DefaultApi* | [**VmaasWebappAppPkgtreeHandlerV3GetGet**](docs/DefaultApi.md#vmaaswebappapppkgtreehandlerv3getget) | **Get** /pkgtree/{package_name} | 
*DefaultApi* | [**VmaasWebappAppPkgtreeHandlerV3PostPost**](docs/DefaultApi.md#vmaaswebappapppkgtreehandlerv3postpost) | **Post** /pkgtree | 
*DefaultApi* | [**VmaasWebappAppReadyHandlerGet**](docs/DefaultApi.md#vmaaswebappappreadyhandlerget) | **Get** /monitoring/ready | Return API readiness status
*DefaultApi* | [**VmaasWebappAppReposHandlerGetGet**](docs/DefaultApi.md#vmaaswebappappreposhandlergetget) | **Get** /repos/{repo} | 
*DefaultApi* | [**VmaasWebappAppReposHandlerPostPost**](docs/DefaultApi.md#vmaaswebappappreposhandlerpostpost) | **Post** /repos | 
*DefaultApi* | [**VmaasWebappAppUpdatesHandlerV3GetGet**](docs/DefaultApi.md#vmaaswebappappupdateshandlerv3getget) | **Get** /updates/{nevra} | 
*DefaultApi* | [**VmaasWebappAppUpdatesHandlerV3PostPost**](docs/DefaultApi.md#vmaaswebappappupdateshandlerv3postpost) | **Post** /updates | 
*DefaultApi* | [**VmaasWebappAppVersionHandlerGet**](docs/DefaultApi.md#vmaaswebappappversionhandlerget) | **Get** /version | 
*DefaultApi* | [**VmaasWebappAppVulnerabilitiesHandlerGetGet**](docs/DefaultApi.md#vmaaswebappappvulnerabilitieshandlergetget) | **Get** /vulnerabilities/{nevra} | 
*DefaultApi* | [**VmaasWebappAppVulnerabilitiesHandlerPostPost**](docs/DefaultApi.md#vmaaswebappappvulnerabilitieshandlerpostpost) | **Post** /vulnerabilities | 
*ExperimentalApi* | [**VmaasWebappAppRPMPkgNamesHandlerGetGet**](docs/ExperimentalApi.md#vmaaswebappapprpmpkgnameshandlergetget) | **Get** /package_names/rpms/{rpm} | 
*ExperimentalApi* | [**VmaasWebappAppRPMPkgNamesHandlerPostPost**](docs/ExperimentalApi.md#vmaaswebappapprpmpkgnameshandlerpostpost) | **Post** /package_names/rpms | 
*ExperimentalApi* | [**VmaasWebappAppSRPMPkgNamesHandlerGetGet**](docs/ExperimentalApi.md#vmaaswebappappsrpmpkgnameshandlergetget) | **Get** /package_names/srpms/{srpm} | 
*ExperimentalApi* | [**VmaasWebappAppSRPMPkgNamesHandlerPostPost**](docs/ExperimentalApi.md#vmaaswebappappsrpmpkgnameshandlerpostpost) | **Post** /package_names/srpms | 


## Documentation For Models

 - [CvesRequest](docs/CvesRequest.md)
 - [CvesResponse](docs/CvesResponse.md)
 - [CvesResponseCveList](docs/CvesResponseCveList.md)
 - [DBChangeResponse](docs/DBChangeResponse.md)
 - [DBChangeResponseDbchange](docs/DBChangeResponseDbchange.md)
 - [ErrataRequest](docs/ErrataRequest.md)
 - [ErrataResponse](docs/ErrataResponse.md)
 - [ErrataResponseErrataList](docs/ErrataResponseErrataList.md)
 - [PackagesRequest](docs/PackagesRequest.md)
 - [PackagesResponse](docs/PackagesResponse.md)
 - [PackagesResponsePackageList](docs/PackagesResponsePackageList.md)
 - [PackagesResponseRepositories](docs/PackagesResponseRepositories.md)
 - [PatchesRequest](docs/PatchesRequest.md)
 - [PatchesResponse](docs/PatchesResponse.md)
 - [PkgListItem](docs/PkgListItem.md)
 - [PkgListRequest](docs/PkgListRequest.md)
 - [PkgListResponse](docs/PkgListResponse.md)
 - [PkgTreeItem](docs/PkgTreeItem.md)
 - [PkgTreeItemErrata](docs/PkgTreeItemErrata.md)
 - [PkgTreeItemRepositories](docs/PkgTreeItemRepositories.md)
 - [PkgtreeRequest](docs/PkgtreeRequest.md)
 - [PkgtreeResponse](docs/PkgtreeResponse.md)
 - [RPMPkgNamesRequest](docs/RPMPkgNamesRequest.md)
 - [RPMPkgNamesResponse](docs/RPMPkgNamesResponse.md)
 - [ReposRequest](docs/ReposRequest.md)
 - [ReposResponse](docs/ReposResponse.md)
 - [SRPMPkgNamesRequest](docs/SRPMPkgNamesRequest.md)
 - [SRPMPkgNamesResponse](docs/SRPMPkgNamesResponse.md)
 - [UpdatesV2Response](docs/UpdatesV2Response.md)
 - [UpdatesV2ResponseAvailableUpdates](docs/UpdatesV2ResponseAvailableUpdates.md)
 - [UpdatesV2ResponseUpdateList](docs/UpdatesV2ResponseUpdateList.md)
 - [UpdatesV3Request](docs/UpdatesV3Request.md)
 - [UpdatesV3RequestModulesList](docs/UpdatesV3RequestModulesList.md)
 - [VulnerabilitiesRequest](docs/VulnerabilitiesRequest.md)
 - [VulnerabilitiesResponse](docs/VulnerabilitiesResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



