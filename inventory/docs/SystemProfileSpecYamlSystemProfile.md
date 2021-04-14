# SystemProfileSpecYamlSystemProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** |  | [optional] 
**BiosReleaseDate** | Pointer to **string** |  | [optional] 
**BiosVendor** | Pointer to **string** |  | [optional] 
**BiosVersion** | Pointer to **string** |  | [optional] 
**CapturedDate** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CoresPerSocket** | Pointer to **int32** |  | [optional] 
**CpuFlags** | Pointer to **[]string** |  | [optional] 
**CpuModel** | Pointer to **string** | The cpu model name | [optional] 
**DiskDevices** | Pointer to [**[]SystemProfileSpecYamlDiskDevice**](SystemProfileSpecYamlDiskDevice.md) |  | [optional] 
**DnfModules** | Pointer to [**[]SystemProfileSpecYamlDnfModule**](SystemProfileSpecYamlDnfModule.md) |  | [optional] 
**EnabledServices** | Pointer to **[]string** |  | [optional] 
**GpgPubkeys** | Pointer to **[]string** |  | [optional] 
**InfrastructureType** | Pointer to **string** |  | [optional] 
**InfrastructureVendor** | Pointer to **string** |  | [optional] 
**InsightsClientVersion** | Pointer to **string** |  | [optional] 
**InsightsEggVersion** | Pointer to **string** |  | [optional] 
**InstalledPackages** | Pointer to **[]string** |  | [optional] 
**InstalledPackagesDelta** | Pointer to **[]string** |  | [optional] 
**InstalledProducts** | Pointer to [**[]SystemProfileSpecYamlInstalledProduct**](SystemProfileSpecYamlInstalledProduct.md) |  | [optional] 
**InstalledServices** | Pointer to **[]string** |  | [optional] 
**IsMarketplace** | Pointer to **bool** | Indicates whether the host is part of a marketplace install from AWS, Azure, etc. | [optional] 
**KatelloAgentRunning** | Pointer to **bool** |  | [optional] 
**KernelModules** | Pointer to **[]string** |  | [optional] 
**LastBootTime** | Pointer to **time.Time** |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]SystemProfileSpecYamlNetworkInterface**](SystemProfileSpecYamlNetworkInterface.md) |  | [optional] 
**NumberOfCpus** | Pointer to **int32** |  | [optional] 
**NumberOfSockets** | Pointer to **int32** |  | [optional] 
**OperatingSystem** | Pointer to [**SystemProfileSpecYamlSystemProfileOperatingSystem**](system_profile_spec_yaml_SystemProfile_operating_system.md) |  | [optional] 
**OsKernelVersion** | Pointer to **string** | The kernel version represented with a three, optionally four, number scheme. | [optional] 
**OsRelease** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** | A UUID associated with the host&#39;s RHSM certificate | [optional] 
**RhcClientId** | Pointer to **string** | A UUID associated with a cloud_connector | [optional] 
**RhcConfigState** | Pointer to **string** | A UUID associated with the config manager state | [optional] 
**Rhsm** | Pointer to [**SystemProfileSpecYamlSystemProfileRhsm**](system_profile_spec_yaml_SystemProfile_rhsm.md) |  | [optional] 
**RunningProcesses** | Pointer to **[]string** |  | [optional] 
**SapInstanceNumber** | Pointer to **string** | The instance number of the SAP HANA system (a two-digit number between 00 and 99) | [optional] 
**SapSids** | Pointer to **[]string** |  | [optional] 
**SapSystem** | Pointer to **bool** | Indicates if SAP is installed on the system | [optional] 
**SapVersion** | Pointer to **string** | The version of the SAP HANA lifecycle management program | [optional] 
**SatelliteManaged** | Pointer to **bool** |  | [optional] 
**SelinuxConfigFile** | Pointer to **string** | The SELinux mode provided in the config file | [optional] 
**SelinuxCurrentMode** | Pointer to **string** | The current SELinux mode, either enforcing, permissive, or disabled | [optional] 
**SubscriptionAutoAttach** | Pointer to **string** |  | [optional] 
**SubscriptionStatus** | Pointer to **string** |  | [optional] 
**SystemMemoryBytes** | Pointer to **int64** |  | [optional] 
**TunedProfile** | Pointer to **string** | Current profile resulting from command tuned-adm active | [optional] 
**YumRepos** | Pointer to [**[]SystemProfileSpecYamlYumRepo**](SystemProfileSpecYamlYumRepo.md) |  | [optional] 

## Methods

### NewSystemProfileSpecYamlSystemProfile

`func NewSystemProfileSpecYamlSystemProfile() *SystemProfileSpecYamlSystemProfile`

NewSystemProfileSpecYamlSystemProfile instantiates a new SystemProfileSpecYamlSystemProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileSpecYamlSystemProfileWithDefaults

`func NewSystemProfileSpecYamlSystemProfileWithDefaults() *SystemProfileSpecYamlSystemProfile`

NewSystemProfileSpecYamlSystemProfileWithDefaults instantiates a new SystemProfileSpecYamlSystemProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *SystemProfileSpecYamlSystemProfile) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *SystemProfileSpecYamlSystemProfile) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *SystemProfileSpecYamlSystemProfile) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *SystemProfileSpecYamlSystemProfile) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBiosReleaseDate

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosReleaseDate() string`

GetBiosReleaseDate returns the BiosReleaseDate field if non-nil, zero value otherwise.

### GetBiosReleaseDateOk

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosReleaseDateOk() (*string, bool)`

GetBiosReleaseDateOk returns a tuple with the BiosReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosReleaseDate

`func (o *SystemProfileSpecYamlSystemProfile) SetBiosReleaseDate(v string)`

SetBiosReleaseDate sets BiosReleaseDate field to given value.

### HasBiosReleaseDate

`func (o *SystemProfileSpecYamlSystemProfile) HasBiosReleaseDate() bool`

HasBiosReleaseDate returns a boolean if a field has been set.

### GetBiosVendor

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosVendor() string`

GetBiosVendor returns the BiosVendor field if non-nil, zero value otherwise.

### GetBiosVendorOk

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosVendorOk() (*string, bool)`

GetBiosVendorOk returns a tuple with the BiosVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVendor

`func (o *SystemProfileSpecYamlSystemProfile) SetBiosVendor(v string)`

SetBiosVendor sets BiosVendor field to given value.

### HasBiosVendor

`func (o *SystemProfileSpecYamlSystemProfile) HasBiosVendor() bool`

HasBiosVendor returns a boolean if a field has been set.

### GetBiosVersion

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosVersion() string`

GetBiosVersion returns the BiosVersion field if non-nil, zero value otherwise.

### GetBiosVersionOk

`func (o *SystemProfileSpecYamlSystemProfile) GetBiosVersionOk() (*string, bool)`

GetBiosVersionOk returns a tuple with the BiosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVersion

`func (o *SystemProfileSpecYamlSystemProfile) SetBiosVersion(v string)`

SetBiosVersion sets BiosVersion field to given value.

### HasBiosVersion

`func (o *SystemProfileSpecYamlSystemProfile) HasBiosVersion() bool`

HasBiosVersion returns a boolean if a field has been set.

### GetCapturedDate

`func (o *SystemProfileSpecYamlSystemProfile) GetCapturedDate() string`

GetCapturedDate returns the CapturedDate field if non-nil, zero value otherwise.

### GetCapturedDateOk

`func (o *SystemProfileSpecYamlSystemProfile) GetCapturedDateOk() (*string, bool)`

GetCapturedDateOk returns a tuple with the CapturedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedDate

`func (o *SystemProfileSpecYamlSystemProfile) SetCapturedDate(v string)`

SetCapturedDate sets CapturedDate field to given value.

### HasCapturedDate

`func (o *SystemProfileSpecYamlSystemProfile) HasCapturedDate() bool`

HasCapturedDate returns a boolean if a field has been set.

### GetCloudProvider

`func (o *SystemProfileSpecYamlSystemProfile) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SystemProfileSpecYamlSystemProfile) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SystemProfileSpecYamlSystemProfile) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *SystemProfileSpecYamlSystemProfile) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *SystemProfileSpecYamlSystemProfile) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *SystemProfileSpecYamlSystemProfile) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *SystemProfileSpecYamlSystemProfile) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *SystemProfileSpecYamlSystemProfile) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCpuFlags

`func (o *SystemProfileSpecYamlSystemProfile) GetCpuFlags() []string`

GetCpuFlags returns the CpuFlags field if non-nil, zero value otherwise.

### GetCpuFlagsOk

`func (o *SystemProfileSpecYamlSystemProfile) GetCpuFlagsOk() (*[]string, bool)`

GetCpuFlagsOk returns a tuple with the CpuFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFlags

`func (o *SystemProfileSpecYamlSystemProfile) SetCpuFlags(v []string)`

SetCpuFlags sets CpuFlags field to given value.

### HasCpuFlags

`func (o *SystemProfileSpecYamlSystemProfile) HasCpuFlags() bool`

HasCpuFlags returns a boolean if a field has been set.

### GetCpuModel

`func (o *SystemProfileSpecYamlSystemProfile) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *SystemProfileSpecYamlSystemProfile) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *SystemProfileSpecYamlSystemProfile) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *SystemProfileSpecYamlSystemProfile) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDiskDevices

`func (o *SystemProfileSpecYamlSystemProfile) GetDiskDevices() []SystemProfileSpecYamlDiskDevice`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetDiskDevicesOk() (*[]SystemProfileSpecYamlDiskDevice, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *SystemProfileSpecYamlSystemProfile) SetDiskDevices(v []SystemProfileSpecYamlDiskDevice)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *SystemProfileSpecYamlSystemProfile) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetDnfModules

`func (o *SystemProfileSpecYamlSystemProfile) GetDnfModules() []SystemProfileSpecYamlDnfModule`

GetDnfModules returns the DnfModules field if non-nil, zero value otherwise.

### GetDnfModulesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetDnfModulesOk() (*[]SystemProfileSpecYamlDnfModule, bool)`

GetDnfModulesOk returns a tuple with the DnfModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnfModules

`func (o *SystemProfileSpecYamlSystemProfile) SetDnfModules(v []SystemProfileSpecYamlDnfModule)`

SetDnfModules sets DnfModules field to given value.

### HasDnfModules

`func (o *SystemProfileSpecYamlSystemProfile) HasDnfModules() bool`

HasDnfModules returns a boolean if a field has been set.

### GetEnabledServices

`func (o *SystemProfileSpecYamlSystemProfile) GetEnabledServices() []string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetEnabledServicesOk() (*[]string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *SystemProfileSpecYamlSystemProfile) SetEnabledServices(v []string)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *SystemProfileSpecYamlSystemProfile) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### GetGpgPubkeys

`func (o *SystemProfileSpecYamlSystemProfile) GetGpgPubkeys() []string`

GetGpgPubkeys returns the GpgPubkeys field if non-nil, zero value otherwise.

### GetGpgPubkeysOk

`func (o *SystemProfileSpecYamlSystemProfile) GetGpgPubkeysOk() (*[]string, bool)`

GetGpgPubkeysOk returns a tuple with the GpgPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPubkeys

`func (o *SystemProfileSpecYamlSystemProfile) SetGpgPubkeys(v []string)`

SetGpgPubkeys sets GpgPubkeys field to given value.

### HasGpgPubkeys

`func (o *SystemProfileSpecYamlSystemProfile) HasGpgPubkeys() bool`

HasGpgPubkeys returns a boolean if a field has been set.

### GetInfrastructureType

`func (o *SystemProfileSpecYamlSystemProfile) GetInfrastructureType() string`

GetInfrastructureType returns the InfrastructureType field if non-nil, zero value otherwise.

### GetInfrastructureTypeOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInfrastructureTypeOk() (*string, bool)`

GetInfrastructureTypeOk returns a tuple with the InfrastructureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureType

`func (o *SystemProfileSpecYamlSystemProfile) SetInfrastructureType(v string)`

SetInfrastructureType sets InfrastructureType field to given value.

### HasInfrastructureType

`func (o *SystemProfileSpecYamlSystemProfile) HasInfrastructureType() bool`

HasInfrastructureType returns a boolean if a field has been set.

### GetInfrastructureVendor

`func (o *SystemProfileSpecYamlSystemProfile) GetInfrastructureVendor() string`

GetInfrastructureVendor returns the InfrastructureVendor field if non-nil, zero value otherwise.

### GetInfrastructureVendorOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInfrastructureVendorOk() (*string, bool)`

GetInfrastructureVendorOk returns a tuple with the InfrastructureVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVendor

`func (o *SystemProfileSpecYamlSystemProfile) SetInfrastructureVendor(v string)`

SetInfrastructureVendor sets InfrastructureVendor field to given value.

### HasInfrastructureVendor

`func (o *SystemProfileSpecYamlSystemProfile) HasInfrastructureVendor() bool`

HasInfrastructureVendor returns a boolean if a field has been set.

### GetInsightsClientVersion

`func (o *SystemProfileSpecYamlSystemProfile) GetInsightsClientVersion() string`

GetInsightsClientVersion returns the InsightsClientVersion field if non-nil, zero value otherwise.

### GetInsightsClientVersionOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInsightsClientVersionOk() (*string, bool)`

GetInsightsClientVersionOk returns a tuple with the InsightsClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsClientVersion

`func (o *SystemProfileSpecYamlSystemProfile) SetInsightsClientVersion(v string)`

SetInsightsClientVersion sets InsightsClientVersion field to given value.

### HasInsightsClientVersion

`func (o *SystemProfileSpecYamlSystemProfile) HasInsightsClientVersion() bool`

HasInsightsClientVersion returns a boolean if a field has been set.

### GetInsightsEggVersion

`func (o *SystemProfileSpecYamlSystemProfile) GetInsightsEggVersion() string`

GetInsightsEggVersion returns the InsightsEggVersion field if non-nil, zero value otherwise.

### GetInsightsEggVersionOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInsightsEggVersionOk() (*string, bool)`

GetInsightsEggVersionOk returns a tuple with the InsightsEggVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsEggVersion

`func (o *SystemProfileSpecYamlSystemProfile) SetInsightsEggVersion(v string)`

SetInsightsEggVersion sets InsightsEggVersion field to given value.

### HasInsightsEggVersion

`func (o *SystemProfileSpecYamlSystemProfile) HasInsightsEggVersion() bool`

HasInsightsEggVersion returns a boolean if a field has been set.

### GetInstalledPackages

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledPackages() []string`

GetInstalledPackages returns the InstalledPackages field if non-nil, zero value otherwise.

### GetInstalledPackagesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledPackagesOk() (*[]string, bool)`

GetInstalledPackagesOk returns a tuple with the InstalledPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackages

`func (o *SystemProfileSpecYamlSystemProfile) SetInstalledPackages(v []string)`

SetInstalledPackages sets InstalledPackages field to given value.

### HasInstalledPackages

`func (o *SystemProfileSpecYamlSystemProfile) HasInstalledPackages() bool`

HasInstalledPackages returns a boolean if a field has been set.

### GetInstalledPackagesDelta

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledPackagesDelta() []string`

GetInstalledPackagesDelta returns the InstalledPackagesDelta field if non-nil, zero value otherwise.

### GetInstalledPackagesDeltaOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledPackagesDeltaOk() (*[]string, bool)`

GetInstalledPackagesDeltaOk returns a tuple with the InstalledPackagesDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackagesDelta

`func (o *SystemProfileSpecYamlSystemProfile) SetInstalledPackagesDelta(v []string)`

SetInstalledPackagesDelta sets InstalledPackagesDelta field to given value.

### HasInstalledPackagesDelta

`func (o *SystemProfileSpecYamlSystemProfile) HasInstalledPackagesDelta() bool`

HasInstalledPackagesDelta returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledProducts() []SystemProfileSpecYamlInstalledProduct`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledProductsOk() (*[]SystemProfileSpecYamlInstalledProduct, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *SystemProfileSpecYamlSystemProfile) SetInstalledProducts(v []SystemProfileSpecYamlInstalledProduct)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *SystemProfileSpecYamlSystemProfile) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetInstalledServices

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledServices() []string`

GetInstalledServices returns the InstalledServices field if non-nil, zero value otherwise.

### GetInstalledServicesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetInstalledServicesOk() (*[]string, bool)`

GetInstalledServicesOk returns a tuple with the InstalledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledServices

`func (o *SystemProfileSpecYamlSystemProfile) SetInstalledServices(v []string)`

SetInstalledServices sets InstalledServices field to given value.

### HasInstalledServices

`func (o *SystemProfileSpecYamlSystemProfile) HasInstalledServices() bool`

HasInstalledServices returns a boolean if a field has been set.

### GetIsMarketplace

`func (o *SystemProfileSpecYamlSystemProfile) GetIsMarketplace() bool`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *SystemProfileSpecYamlSystemProfile) GetIsMarketplaceOk() (*bool, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *SystemProfileSpecYamlSystemProfile) SetIsMarketplace(v bool)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *SystemProfileSpecYamlSystemProfile) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### GetKatelloAgentRunning

`func (o *SystemProfileSpecYamlSystemProfile) GetKatelloAgentRunning() bool`

GetKatelloAgentRunning returns the KatelloAgentRunning field if non-nil, zero value otherwise.

### GetKatelloAgentRunningOk

`func (o *SystemProfileSpecYamlSystemProfile) GetKatelloAgentRunningOk() (*bool, bool)`

GetKatelloAgentRunningOk returns a tuple with the KatelloAgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKatelloAgentRunning

`func (o *SystemProfileSpecYamlSystemProfile) SetKatelloAgentRunning(v bool)`

SetKatelloAgentRunning sets KatelloAgentRunning field to given value.

### HasKatelloAgentRunning

`func (o *SystemProfileSpecYamlSystemProfile) HasKatelloAgentRunning() bool`

HasKatelloAgentRunning returns a boolean if a field has been set.

### GetKernelModules

`func (o *SystemProfileSpecYamlSystemProfile) GetKernelModules() []string`

GetKernelModules returns the KernelModules field if non-nil, zero value otherwise.

### GetKernelModulesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetKernelModulesOk() (*[]string, bool)`

GetKernelModulesOk returns a tuple with the KernelModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelModules

`func (o *SystemProfileSpecYamlSystemProfile) SetKernelModules(v []string)`

SetKernelModules sets KernelModules field to given value.

### HasKernelModules

`func (o *SystemProfileSpecYamlSystemProfile) HasKernelModules() bool`

HasKernelModules returns a boolean if a field has been set.

### GetLastBootTime

`func (o *SystemProfileSpecYamlSystemProfile) GetLastBootTime() time.Time`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *SystemProfileSpecYamlSystemProfile) GetLastBootTimeOk() (*time.Time, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *SystemProfileSpecYamlSystemProfile) SetLastBootTime(v time.Time)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *SystemProfileSpecYamlSystemProfile) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *SystemProfileSpecYamlSystemProfile) GetNetworkInterfaces() []SystemProfileSpecYamlNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetNetworkInterfacesOk() (*[]SystemProfileSpecYamlNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *SystemProfileSpecYamlSystemProfile) SetNetworkInterfaces(v []SystemProfileSpecYamlNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *SystemProfileSpecYamlSystemProfile) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetNumberOfCpus

`func (o *SystemProfileSpecYamlSystemProfile) GetNumberOfCpus() int32`

GetNumberOfCpus returns the NumberOfCpus field if non-nil, zero value otherwise.

### GetNumberOfCpusOk

`func (o *SystemProfileSpecYamlSystemProfile) GetNumberOfCpusOk() (*int32, bool)`

GetNumberOfCpusOk returns a tuple with the NumberOfCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpus

`func (o *SystemProfileSpecYamlSystemProfile) SetNumberOfCpus(v int32)`

SetNumberOfCpus sets NumberOfCpus field to given value.

### HasNumberOfCpus

`func (o *SystemProfileSpecYamlSystemProfile) HasNumberOfCpus() bool`

HasNumberOfCpus returns a boolean if a field has been set.

### GetNumberOfSockets

`func (o *SystemProfileSpecYamlSystemProfile) GetNumberOfSockets() int32`

GetNumberOfSockets returns the NumberOfSockets field if non-nil, zero value otherwise.

### GetNumberOfSocketsOk

`func (o *SystemProfileSpecYamlSystemProfile) GetNumberOfSocketsOk() (*int32, bool)`

GetNumberOfSocketsOk returns a tuple with the NumberOfSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSockets

`func (o *SystemProfileSpecYamlSystemProfile) SetNumberOfSockets(v int32)`

SetNumberOfSockets sets NumberOfSockets field to given value.

### HasNumberOfSockets

`func (o *SystemProfileSpecYamlSystemProfile) HasNumberOfSockets() bool`

HasNumberOfSockets returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SystemProfileSpecYamlSystemProfile) GetOperatingSystem() SystemProfileSpecYamlSystemProfileOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SystemProfileSpecYamlSystemProfile) GetOperatingSystemOk() (*SystemProfileSpecYamlSystemProfileOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SystemProfileSpecYamlSystemProfile) SetOperatingSystem(v SystemProfileSpecYamlSystemProfileOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SystemProfileSpecYamlSystemProfile) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOsKernelVersion

`func (o *SystemProfileSpecYamlSystemProfile) GetOsKernelVersion() string`

GetOsKernelVersion returns the OsKernelVersion field if non-nil, zero value otherwise.

### GetOsKernelVersionOk

`func (o *SystemProfileSpecYamlSystemProfile) GetOsKernelVersionOk() (*string, bool)`

GetOsKernelVersionOk returns a tuple with the OsKernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsKernelVersion

`func (o *SystemProfileSpecYamlSystemProfile) SetOsKernelVersion(v string)`

SetOsKernelVersion sets OsKernelVersion field to given value.

### HasOsKernelVersion

`func (o *SystemProfileSpecYamlSystemProfile) HasOsKernelVersion() bool`

HasOsKernelVersion returns a boolean if a field has been set.

### GetOsRelease

`func (o *SystemProfileSpecYamlSystemProfile) GetOsRelease() string`

GetOsRelease returns the OsRelease field if non-nil, zero value otherwise.

### GetOsReleaseOk

`func (o *SystemProfileSpecYamlSystemProfile) GetOsReleaseOk() (*string, bool)`

GetOsReleaseOk returns a tuple with the OsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRelease

`func (o *SystemProfileSpecYamlSystemProfile) SetOsRelease(v string)`

SetOsRelease sets OsRelease field to given value.

### HasOsRelease

`func (o *SystemProfileSpecYamlSystemProfile) HasOsRelease() bool`

HasOsRelease returns a boolean if a field has been set.

### GetOwnerId

`func (o *SystemProfileSpecYamlSystemProfile) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SystemProfileSpecYamlSystemProfile) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SystemProfileSpecYamlSystemProfile) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SystemProfileSpecYamlSystemProfile) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRhcClientId

`func (o *SystemProfileSpecYamlSystemProfile) GetRhcClientId() string`

GetRhcClientId returns the RhcClientId field if non-nil, zero value otherwise.

### GetRhcClientIdOk

`func (o *SystemProfileSpecYamlSystemProfile) GetRhcClientIdOk() (*string, bool)`

GetRhcClientIdOk returns a tuple with the RhcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhcClientId

`func (o *SystemProfileSpecYamlSystemProfile) SetRhcClientId(v string)`

SetRhcClientId sets RhcClientId field to given value.

### HasRhcClientId

`func (o *SystemProfileSpecYamlSystemProfile) HasRhcClientId() bool`

HasRhcClientId returns a boolean if a field has been set.

### GetRhcConfigState

`func (o *SystemProfileSpecYamlSystemProfile) GetRhcConfigState() string`

GetRhcConfigState returns the RhcConfigState field if non-nil, zero value otherwise.

### GetRhcConfigStateOk

`func (o *SystemProfileSpecYamlSystemProfile) GetRhcConfigStateOk() (*string, bool)`

GetRhcConfigStateOk returns a tuple with the RhcConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhcConfigState

`func (o *SystemProfileSpecYamlSystemProfile) SetRhcConfigState(v string)`

SetRhcConfigState sets RhcConfigState field to given value.

### HasRhcConfigState

`func (o *SystemProfileSpecYamlSystemProfile) HasRhcConfigState() bool`

HasRhcConfigState returns a boolean if a field has been set.

### GetRhsm

`func (o *SystemProfileSpecYamlSystemProfile) GetRhsm() SystemProfileSpecYamlSystemProfileRhsm`

GetRhsm returns the Rhsm field if non-nil, zero value otherwise.

### GetRhsmOk

`func (o *SystemProfileSpecYamlSystemProfile) GetRhsmOk() (*SystemProfileSpecYamlSystemProfileRhsm, bool)`

GetRhsmOk returns a tuple with the Rhsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhsm

`func (o *SystemProfileSpecYamlSystemProfile) SetRhsm(v SystemProfileSpecYamlSystemProfileRhsm)`

SetRhsm sets Rhsm field to given value.

### HasRhsm

`func (o *SystemProfileSpecYamlSystemProfile) HasRhsm() bool`

HasRhsm returns a boolean if a field has been set.

### GetRunningProcesses

`func (o *SystemProfileSpecYamlSystemProfile) GetRunningProcesses() []string`

GetRunningProcesses returns the RunningProcesses field if non-nil, zero value otherwise.

### GetRunningProcessesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetRunningProcessesOk() (*[]string, bool)`

GetRunningProcessesOk returns a tuple with the RunningProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningProcesses

`func (o *SystemProfileSpecYamlSystemProfile) SetRunningProcesses(v []string)`

SetRunningProcesses sets RunningProcesses field to given value.

### HasRunningProcesses

`func (o *SystemProfileSpecYamlSystemProfile) HasRunningProcesses() bool`

HasRunningProcesses returns a boolean if a field has been set.

### GetSapInstanceNumber

`func (o *SystemProfileSpecYamlSystemProfile) GetSapInstanceNumber() string`

GetSapInstanceNumber returns the SapInstanceNumber field if non-nil, zero value otherwise.

### GetSapInstanceNumberOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSapInstanceNumberOk() (*string, bool)`

GetSapInstanceNumberOk returns a tuple with the SapInstanceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapInstanceNumber

`func (o *SystemProfileSpecYamlSystemProfile) SetSapInstanceNumber(v string)`

SetSapInstanceNumber sets SapInstanceNumber field to given value.

### HasSapInstanceNumber

`func (o *SystemProfileSpecYamlSystemProfile) HasSapInstanceNumber() bool`

HasSapInstanceNumber returns a boolean if a field has been set.

### GetSapSids

`func (o *SystemProfileSpecYamlSystemProfile) GetSapSids() []string`

GetSapSids returns the SapSids field if non-nil, zero value otherwise.

### GetSapSidsOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSapSidsOk() (*[]string, bool)`

GetSapSidsOk returns a tuple with the SapSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSids

`func (o *SystemProfileSpecYamlSystemProfile) SetSapSids(v []string)`

SetSapSids sets SapSids field to given value.

### HasSapSids

`func (o *SystemProfileSpecYamlSystemProfile) HasSapSids() bool`

HasSapSids returns a boolean if a field has been set.

### GetSapSystem

`func (o *SystemProfileSpecYamlSystemProfile) GetSapSystem() bool`

GetSapSystem returns the SapSystem field if non-nil, zero value otherwise.

### GetSapSystemOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSapSystemOk() (*bool, bool)`

GetSapSystemOk returns a tuple with the SapSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSystem

`func (o *SystemProfileSpecYamlSystemProfile) SetSapSystem(v bool)`

SetSapSystem sets SapSystem field to given value.

### HasSapSystem

`func (o *SystemProfileSpecYamlSystemProfile) HasSapSystem() bool`

HasSapSystem returns a boolean if a field has been set.

### GetSapVersion

`func (o *SystemProfileSpecYamlSystemProfile) GetSapVersion() string`

GetSapVersion returns the SapVersion field if non-nil, zero value otherwise.

### GetSapVersionOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSapVersionOk() (*string, bool)`

GetSapVersionOk returns a tuple with the SapVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapVersion

`func (o *SystemProfileSpecYamlSystemProfile) SetSapVersion(v string)`

SetSapVersion sets SapVersion field to given value.

### HasSapVersion

`func (o *SystemProfileSpecYamlSystemProfile) HasSapVersion() bool`

HasSapVersion returns a boolean if a field has been set.

### GetSatelliteManaged

`func (o *SystemProfileSpecYamlSystemProfile) GetSatelliteManaged() bool`

GetSatelliteManaged returns the SatelliteManaged field if non-nil, zero value otherwise.

### GetSatelliteManagedOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSatelliteManagedOk() (*bool, bool)`

GetSatelliteManagedOk returns a tuple with the SatelliteManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteManaged

`func (o *SystemProfileSpecYamlSystemProfile) SetSatelliteManaged(v bool)`

SetSatelliteManaged sets SatelliteManaged field to given value.

### HasSatelliteManaged

`func (o *SystemProfileSpecYamlSystemProfile) HasSatelliteManaged() bool`

HasSatelliteManaged returns a boolean if a field has been set.

### GetSelinuxConfigFile

`func (o *SystemProfileSpecYamlSystemProfile) GetSelinuxConfigFile() string`

GetSelinuxConfigFile returns the SelinuxConfigFile field if non-nil, zero value otherwise.

### GetSelinuxConfigFileOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSelinuxConfigFileOk() (*string, bool)`

GetSelinuxConfigFileOk returns a tuple with the SelinuxConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxConfigFile

`func (o *SystemProfileSpecYamlSystemProfile) SetSelinuxConfigFile(v string)`

SetSelinuxConfigFile sets SelinuxConfigFile field to given value.

### HasSelinuxConfigFile

`func (o *SystemProfileSpecYamlSystemProfile) HasSelinuxConfigFile() bool`

HasSelinuxConfigFile returns a boolean if a field has been set.

### GetSelinuxCurrentMode

`func (o *SystemProfileSpecYamlSystemProfile) GetSelinuxCurrentMode() string`

GetSelinuxCurrentMode returns the SelinuxCurrentMode field if non-nil, zero value otherwise.

### GetSelinuxCurrentModeOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSelinuxCurrentModeOk() (*string, bool)`

GetSelinuxCurrentModeOk returns a tuple with the SelinuxCurrentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxCurrentMode

`func (o *SystemProfileSpecYamlSystemProfile) SetSelinuxCurrentMode(v string)`

SetSelinuxCurrentMode sets SelinuxCurrentMode field to given value.

### HasSelinuxCurrentMode

`func (o *SystemProfileSpecYamlSystemProfile) HasSelinuxCurrentMode() bool`

HasSelinuxCurrentMode returns a boolean if a field has been set.

### GetSubscriptionAutoAttach

`func (o *SystemProfileSpecYamlSystemProfile) GetSubscriptionAutoAttach() string`

GetSubscriptionAutoAttach returns the SubscriptionAutoAttach field if non-nil, zero value otherwise.

### GetSubscriptionAutoAttachOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSubscriptionAutoAttachOk() (*string, bool)`

GetSubscriptionAutoAttachOk returns a tuple with the SubscriptionAutoAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAutoAttach

`func (o *SystemProfileSpecYamlSystemProfile) SetSubscriptionAutoAttach(v string)`

SetSubscriptionAutoAttach sets SubscriptionAutoAttach field to given value.

### HasSubscriptionAutoAttach

`func (o *SystemProfileSpecYamlSystemProfile) HasSubscriptionAutoAttach() bool`

HasSubscriptionAutoAttach returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *SystemProfileSpecYamlSystemProfile) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *SystemProfileSpecYamlSystemProfile) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *SystemProfileSpecYamlSystemProfile) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetSystemMemoryBytes

`func (o *SystemProfileSpecYamlSystemProfile) GetSystemMemoryBytes() int64`

GetSystemMemoryBytes returns the SystemMemoryBytes field if non-nil, zero value otherwise.

### GetSystemMemoryBytesOk

`func (o *SystemProfileSpecYamlSystemProfile) GetSystemMemoryBytesOk() (*int64, bool)`

GetSystemMemoryBytesOk returns a tuple with the SystemMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemoryBytes

`func (o *SystemProfileSpecYamlSystemProfile) SetSystemMemoryBytes(v int64)`

SetSystemMemoryBytes sets SystemMemoryBytes field to given value.

### HasSystemMemoryBytes

`func (o *SystemProfileSpecYamlSystemProfile) HasSystemMemoryBytes() bool`

HasSystemMemoryBytes returns a boolean if a field has been set.

### GetTunedProfile

`func (o *SystemProfileSpecYamlSystemProfile) GetTunedProfile() string`

GetTunedProfile returns the TunedProfile field if non-nil, zero value otherwise.

### GetTunedProfileOk

`func (o *SystemProfileSpecYamlSystemProfile) GetTunedProfileOk() (*string, bool)`

GetTunedProfileOk returns a tuple with the TunedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunedProfile

`func (o *SystemProfileSpecYamlSystemProfile) SetTunedProfile(v string)`

SetTunedProfile sets TunedProfile field to given value.

### HasTunedProfile

`func (o *SystemProfileSpecYamlSystemProfile) HasTunedProfile() bool`

HasTunedProfile returns a boolean if a field has been set.

### GetYumRepos

`func (o *SystemProfileSpecYamlSystemProfile) GetYumRepos() []SystemProfileSpecYamlYumRepo`

GetYumRepos returns the YumRepos field if non-nil, zero value otherwise.

### GetYumReposOk

`func (o *SystemProfileSpecYamlSystemProfile) GetYumReposOk() (*[]SystemProfileSpecYamlYumRepo, bool)`

GetYumReposOk returns a tuple with the YumRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumRepos

`func (o *SystemProfileSpecYamlSystemProfile) SetYumRepos(v []SystemProfileSpecYamlYumRepo)`

SetYumRepos sets YumRepos field to given value.

### HasYumRepos

`func (o *SystemProfileSpecYamlSystemProfile) HasYumRepos() bool`

HasYumRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


