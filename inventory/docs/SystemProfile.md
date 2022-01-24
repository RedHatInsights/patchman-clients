# SystemProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ansible** | Pointer to [**SystemProfileAnsible**](SystemProfile_ansible.md) |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**BiosReleaseDate** | Pointer to **string** |  | [optional] 
**BiosVendor** | Pointer to **string** |  | [optional] 
**BiosVersion** | Pointer to **string** |  | [optional] 
**CapturedDate** | Pointer to **string** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 
**CoresPerSocket** | Pointer to **int32** |  | [optional] 
**CpuFlags** | Pointer to **[]string** |  | [optional] 
**CpuModel** | Pointer to **string** | The cpu model name | [optional] 
**DiskDevices** | Pointer to [**[]SystemProfileDiskDevice**](SystemProfileDiskDevice.md) |  | [optional] 
**DnfModules** | Pointer to [**[]SystemProfileDnfModule**](SystemProfileDnfModule.md) |  | [optional] 
**EnabledServices** | Pointer to **[]string** |  | [optional] 
**GpgPubkeys** | Pointer to **[]string** |  | [optional] 
**GreenbootFallbackDetected** | Pointer to **bool** | Indicates whether greenboot detected a rolled back update on an edge device. | [optional] 
**GreenbootStatus** | Pointer to **string** | Indicates the greenboot status of an edge device. | [optional] 
**HostType** | Pointer to **string** | Indicates the type of host. | [optional] 
**InfrastructureType** | Pointer to **string** |  | [optional] 
**InfrastructureVendor** | Pointer to **string** |  | [optional] 
**InsightsClientVersion** | Pointer to **string** | The version number of insights client. supports wildcards | [optional] 
**InsightsEggVersion** | Pointer to **string** |  | [optional] 
**InstalledPackages** | Pointer to **[]string** |  | [optional] 
**InstalledPackagesDelta** | Pointer to **[]string** |  | [optional] 
**InstalledProducts** | Pointer to [**[]SystemProfileInstalledProduct**](SystemProfileInstalledProduct.md) |  | [optional] 
**InstalledServices** | Pointer to **[]string** |  | [optional] 
**IsMarketplace** | Pointer to **bool** | Indicates whether the host is part of a marketplace install from AWS, Azure, etc. | [optional] 
**KatelloAgentRunning** | Pointer to **bool** |  | [optional] 
**KernelModules** | Pointer to **[]string** |  | [optional] 
**LastBootTime** | Pointer to **string** |  | [optional] 
**Mssql** | Pointer to [**SystemProfileMssql**](SystemProfile_mssql.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]SystemProfileNetworkInterface**](SystemProfileNetworkInterface.md) |  | [optional] 
**NumberOfCpus** | Pointer to **int32** |  | [optional] 
**NumberOfSockets** | Pointer to **int32** |  | [optional] 
**OperatingSystem** | Pointer to [**SystemProfileOperatingSystem**](SystemProfile_operating_system.md) |  | [optional] 
**OsKernelVersion** | Pointer to **string** | The kernel version represented with a three, optionally four, number scheme. | [optional] 
**OsRelease** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** | A UUID associated with the host&#39;s RHSM certificate | [optional] 
**RhcClientId** | Pointer to **string** | A UUID associated with a cloud_connector | [optional] 
**RhcConfigState** | Pointer to **string** | A UUID associated with the config manager state | [optional] 
**Rhsm** | Pointer to [**SystemProfileRhsm**](SystemProfile_rhsm.md) |  | [optional] 
**RpmOstreeDeployments** | Pointer to [**[]SystemProfileRpmOstreeDeployments**](SystemProfileRpmOstreeDeployments.md) | The list of deployments on the system as reported by rpm-ostree status --json | [optional] 
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
**SystemPurpose** | Pointer to [**SystemProfileSystemPurpose**](SystemProfile_system_purpose.md) |  | [optional] 
**TunedProfile** | Pointer to **string** | Current profile resulting from command tuned-adm active | [optional] 
**YumRepos** | Pointer to [**[]SystemProfileYumRepo**](SystemProfileYumRepo.md) |  | [optional] 

## Methods

### NewSystemProfile

`func NewSystemProfile() *SystemProfile`

NewSystemProfile instantiates a new SystemProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemProfileWithDefaults

`func NewSystemProfileWithDefaults() *SystemProfile`

NewSystemProfileWithDefaults instantiates a new SystemProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsible

`func (o *SystemProfile) GetAnsible() SystemProfileAnsible`

GetAnsible returns the Ansible field if non-nil, zero value otherwise.

### GetAnsibleOk

`func (o *SystemProfile) GetAnsibleOk() (*SystemProfileAnsible, bool)`

GetAnsibleOk returns a tuple with the Ansible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsible

`func (o *SystemProfile) SetAnsible(v SystemProfileAnsible)`

SetAnsible sets Ansible field to given value.

### HasAnsible

`func (o *SystemProfile) HasAnsible() bool`

HasAnsible returns a boolean if a field has been set.

### GetArch

`func (o *SystemProfile) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *SystemProfile) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *SystemProfile) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *SystemProfile) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBiosReleaseDate

`func (o *SystemProfile) GetBiosReleaseDate() string`

GetBiosReleaseDate returns the BiosReleaseDate field if non-nil, zero value otherwise.

### GetBiosReleaseDateOk

`func (o *SystemProfile) GetBiosReleaseDateOk() (*string, bool)`

GetBiosReleaseDateOk returns a tuple with the BiosReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosReleaseDate

`func (o *SystemProfile) SetBiosReleaseDate(v string)`

SetBiosReleaseDate sets BiosReleaseDate field to given value.

### HasBiosReleaseDate

`func (o *SystemProfile) HasBiosReleaseDate() bool`

HasBiosReleaseDate returns a boolean if a field has been set.

### GetBiosVendor

`func (o *SystemProfile) GetBiosVendor() string`

GetBiosVendor returns the BiosVendor field if non-nil, zero value otherwise.

### GetBiosVendorOk

`func (o *SystemProfile) GetBiosVendorOk() (*string, bool)`

GetBiosVendorOk returns a tuple with the BiosVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVendor

`func (o *SystemProfile) SetBiosVendor(v string)`

SetBiosVendor sets BiosVendor field to given value.

### HasBiosVendor

`func (o *SystemProfile) HasBiosVendor() bool`

HasBiosVendor returns a boolean if a field has been set.

### GetBiosVersion

`func (o *SystemProfile) GetBiosVersion() string`

GetBiosVersion returns the BiosVersion field if non-nil, zero value otherwise.

### GetBiosVersionOk

`func (o *SystemProfile) GetBiosVersionOk() (*string, bool)`

GetBiosVersionOk returns a tuple with the BiosVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosVersion

`func (o *SystemProfile) SetBiosVersion(v string)`

SetBiosVersion sets BiosVersion field to given value.

### HasBiosVersion

`func (o *SystemProfile) HasBiosVersion() bool`

HasBiosVersion returns a boolean if a field has been set.

### GetCapturedDate

`func (o *SystemProfile) GetCapturedDate() string`

GetCapturedDate returns the CapturedDate field if non-nil, zero value otherwise.

### GetCapturedDateOk

`func (o *SystemProfile) GetCapturedDateOk() (*string, bool)`

GetCapturedDateOk returns a tuple with the CapturedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedDate

`func (o *SystemProfile) SetCapturedDate(v string)`

SetCapturedDate sets CapturedDate field to given value.

### HasCapturedDate

`func (o *SystemProfile) HasCapturedDate() bool`

HasCapturedDate returns a boolean if a field has been set.

### GetCloudProvider

`func (o *SystemProfile) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *SystemProfile) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *SystemProfile) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *SystemProfile) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *SystemProfile) GetCoresPerSocket() int32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *SystemProfile) GetCoresPerSocketOk() (*int32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *SystemProfile) SetCoresPerSocket(v int32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *SystemProfile) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCpuFlags

`func (o *SystemProfile) GetCpuFlags() []string`

GetCpuFlags returns the CpuFlags field if non-nil, zero value otherwise.

### GetCpuFlagsOk

`func (o *SystemProfile) GetCpuFlagsOk() (*[]string, bool)`

GetCpuFlagsOk returns a tuple with the CpuFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuFlags

`func (o *SystemProfile) SetCpuFlags(v []string)`

SetCpuFlags sets CpuFlags field to given value.

### HasCpuFlags

`func (o *SystemProfile) HasCpuFlags() bool`

HasCpuFlags returns a boolean if a field has been set.

### GetCpuModel

`func (o *SystemProfile) GetCpuModel() string`

GetCpuModel returns the CpuModel field if non-nil, zero value otherwise.

### GetCpuModelOk

`func (o *SystemProfile) GetCpuModelOk() (*string, bool)`

GetCpuModelOk returns a tuple with the CpuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuModel

`func (o *SystemProfile) SetCpuModel(v string)`

SetCpuModel sets CpuModel field to given value.

### HasCpuModel

`func (o *SystemProfile) HasCpuModel() bool`

HasCpuModel returns a boolean if a field has been set.

### GetDiskDevices

`func (o *SystemProfile) GetDiskDevices() []SystemProfileDiskDevice`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *SystemProfile) GetDiskDevicesOk() (*[]SystemProfileDiskDevice, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *SystemProfile) SetDiskDevices(v []SystemProfileDiskDevice)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *SystemProfile) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetDnfModules

`func (o *SystemProfile) GetDnfModules() []SystemProfileDnfModule`

GetDnfModules returns the DnfModules field if non-nil, zero value otherwise.

### GetDnfModulesOk

`func (o *SystemProfile) GetDnfModulesOk() (*[]SystemProfileDnfModule, bool)`

GetDnfModulesOk returns a tuple with the DnfModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnfModules

`func (o *SystemProfile) SetDnfModules(v []SystemProfileDnfModule)`

SetDnfModules sets DnfModules field to given value.

### HasDnfModules

`func (o *SystemProfile) HasDnfModules() bool`

HasDnfModules returns a boolean if a field has been set.

### GetEnabledServices

`func (o *SystemProfile) GetEnabledServices() []string`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *SystemProfile) GetEnabledServicesOk() (*[]string, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *SystemProfile) SetEnabledServices(v []string)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *SystemProfile) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### GetGpgPubkeys

`func (o *SystemProfile) GetGpgPubkeys() []string`

GetGpgPubkeys returns the GpgPubkeys field if non-nil, zero value otherwise.

### GetGpgPubkeysOk

`func (o *SystemProfile) GetGpgPubkeysOk() (*[]string, bool)`

GetGpgPubkeysOk returns a tuple with the GpgPubkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgPubkeys

`func (o *SystemProfile) SetGpgPubkeys(v []string)`

SetGpgPubkeys sets GpgPubkeys field to given value.

### HasGpgPubkeys

`func (o *SystemProfile) HasGpgPubkeys() bool`

HasGpgPubkeys returns a boolean if a field has been set.

### GetGreenbootFallbackDetected

`func (o *SystemProfile) GetGreenbootFallbackDetected() bool`

GetGreenbootFallbackDetected returns the GreenbootFallbackDetected field if non-nil, zero value otherwise.

### GetGreenbootFallbackDetectedOk

`func (o *SystemProfile) GetGreenbootFallbackDetectedOk() (*bool, bool)`

GetGreenbootFallbackDetectedOk returns a tuple with the GreenbootFallbackDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreenbootFallbackDetected

`func (o *SystemProfile) SetGreenbootFallbackDetected(v bool)`

SetGreenbootFallbackDetected sets GreenbootFallbackDetected field to given value.

### HasGreenbootFallbackDetected

`func (o *SystemProfile) HasGreenbootFallbackDetected() bool`

HasGreenbootFallbackDetected returns a boolean if a field has been set.

### GetGreenbootStatus

`func (o *SystemProfile) GetGreenbootStatus() string`

GetGreenbootStatus returns the GreenbootStatus field if non-nil, zero value otherwise.

### GetGreenbootStatusOk

`func (o *SystemProfile) GetGreenbootStatusOk() (*string, bool)`

GetGreenbootStatusOk returns a tuple with the GreenbootStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreenbootStatus

`func (o *SystemProfile) SetGreenbootStatus(v string)`

SetGreenbootStatus sets GreenbootStatus field to given value.

### HasGreenbootStatus

`func (o *SystemProfile) HasGreenbootStatus() bool`

HasGreenbootStatus returns a boolean if a field has been set.

### GetHostType

`func (o *SystemProfile) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *SystemProfile) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *SystemProfile) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *SystemProfile) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetInfrastructureType

`func (o *SystemProfile) GetInfrastructureType() string`

GetInfrastructureType returns the InfrastructureType field if non-nil, zero value otherwise.

### GetInfrastructureTypeOk

`func (o *SystemProfile) GetInfrastructureTypeOk() (*string, bool)`

GetInfrastructureTypeOk returns a tuple with the InfrastructureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureType

`func (o *SystemProfile) SetInfrastructureType(v string)`

SetInfrastructureType sets InfrastructureType field to given value.

### HasInfrastructureType

`func (o *SystemProfile) HasInfrastructureType() bool`

HasInfrastructureType returns a boolean if a field has been set.

### GetInfrastructureVendor

`func (o *SystemProfile) GetInfrastructureVendor() string`

GetInfrastructureVendor returns the InfrastructureVendor field if non-nil, zero value otherwise.

### GetInfrastructureVendorOk

`func (o *SystemProfile) GetInfrastructureVendorOk() (*string, bool)`

GetInfrastructureVendorOk returns a tuple with the InfrastructureVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrastructureVendor

`func (o *SystemProfile) SetInfrastructureVendor(v string)`

SetInfrastructureVendor sets InfrastructureVendor field to given value.

### HasInfrastructureVendor

`func (o *SystemProfile) HasInfrastructureVendor() bool`

HasInfrastructureVendor returns a boolean if a field has been set.

### GetInsightsClientVersion

`func (o *SystemProfile) GetInsightsClientVersion() string`

GetInsightsClientVersion returns the InsightsClientVersion field if non-nil, zero value otherwise.

### GetInsightsClientVersionOk

`func (o *SystemProfile) GetInsightsClientVersionOk() (*string, bool)`

GetInsightsClientVersionOk returns a tuple with the InsightsClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsClientVersion

`func (o *SystemProfile) SetInsightsClientVersion(v string)`

SetInsightsClientVersion sets InsightsClientVersion field to given value.

### HasInsightsClientVersion

`func (o *SystemProfile) HasInsightsClientVersion() bool`

HasInsightsClientVersion returns a boolean if a field has been set.

### GetInsightsEggVersion

`func (o *SystemProfile) GetInsightsEggVersion() string`

GetInsightsEggVersion returns the InsightsEggVersion field if non-nil, zero value otherwise.

### GetInsightsEggVersionOk

`func (o *SystemProfile) GetInsightsEggVersionOk() (*string, bool)`

GetInsightsEggVersionOk returns a tuple with the InsightsEggVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsEggVersion

`func (o *SystemProfile) SetInsightsEggVersion(v string)`

SetInsightsEggVersion sets InsightsEggVersion field to given value.

### HasInsightsEggVersion

`func (o *SystemProfile) HasInsightsEggVersion() bool`

HasInsightsEggVersion returns a boolean if a field has been set.

### GetInstalledPackages

`func (o *SystemProfile) GetInstalledPackages() []string`

GetInstalledPackages returns the InstalledPackages field if non-nil, zero value otherwise.

### GetInstalledPackagesOk

`func (o *SystemProfile) GetInstalledPackagesOk() (*[]string, bool)`

GetInstalledPackagesOk returns a tuple with the InstalledPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackages

`func (o *SystemProfile) SetInstalledPackages(v []string)`

SetInstalledPackages sets InstalledPackages field to given value.

### HasInstalledPackages

`func (o *SystemProfile) HasInstalledPackages() bool`

HasInstalledPackages returns a boolean if a field has been set.

### GetInstalledPackagesDelta

`func (o *SystemProfile) GetInstalledPackagesDelta() []string`

GetInstalledPackagesDelta returns the InstalledPackagesDelta field if non-nil, zero value otherwise.

### GetInstalledPackagesDeltaOk

`func (o *SystemProfile) GetInstalledPackagesDeltaOk() (*[]string, bool)`

GetInstalledPackagesDeltaOk returns a tuple with the InstalledPackagesDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledPackagesDelta

`func (o *SystemProfile) SetInstalledPackagesDelta(v []string)`

SetInstalledPackagesDelta sets InstalledPackagesDelta field to given value.

### HasInstalledPackagesDelta

`func (o *SystemProfile) HasInstalledPackagesDelta() bool`

HasInstalledPackagesDelta returns a boolean if a field has been set.

### GetInstalledProducts

`func (o *SystemProfile) GetInstalledProducts() []SystemProfileInstalledProduct`

GetInstalledProducts returns the InstalledProducts field if non-nil, zero value otherwise.

### GetInstalledProductsOk

`func (o *SystemProfile) GetInstalledProductsOk() (*[]SystemProfileInstalledProduct, bool)`

GetInstalledProductsOk returns a tuple with the InstalledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledProducts

`func (o *SystemProfile) SetInstalledProducts(v []SystemProfileInstalledProduct)`

SetInstalledProducts sets InstalledProducts field to given value.

### HasInstalledProducts

`func (o *SystemProfile) HasInstalledProducts() bool`

HasInstalledProducts returns a boolean if a field has been set.

### GetInstalledServices

`func (o *SystemProfile) GetInstalledServices() []string`

GetInstalledServices returns the InstalledServices field if non-nil, zero value otherwise.

### GetInstalledServicesOk

`func (o *SystemProfile) GetInstalledServicesOk() (*[]string, bool)`

GetInstalledServicesOk returns a tuple with the InstalledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledServices

`func (o *SystemProfile) SetInstalledServices(v []string)`

SetInstalledServices sets InstalledServices field to given value.

### HasInstalledServices

`func (o *SystemProfile) HasInstalledServices() bool`

HasInstalledServices returns a boolean if a field has been set.

### GetIsMarketplace

`func (o *SystemProfile) GetIsMarketplace() bool`

GetIsMarketplace returns the IsMarketplace field if non-nil, zero value otherwise.

### GetIsMarketplaceOk

`func (o *SystemProfile) GetIsMarketplaceOk() (*bool, bool)`

GetIsMarketplaceOk returns a tuple with the IsMarketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarketplace

`func (o *SystemProfile) SetIsMarketplace(v bool)`

SetIsMarketplace sets IsMarketplace field to given value.

### HasIsMarketplace

`func (o *SystemProfile) HasIsMarketplace() bool`

HasIsMarketplace returns a boolean if a field has been set.

### GetKatelloAgentRunning

`func (o *SystemProfile) GetKatelloAgentRunning() bool`

GetKatelloAgentRunning returns the KatelloAgentRunning field if non-nil, zero value otherwise.

### GetKatelloAgentRunningOk

`func (o *SystemProfile) GetKatelloAgentRunningOk() (*bool, bool)`

GetKatelloAgentRunningOk returns a tuple with the KatelloAgentRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKatelloAgentRunning

`func (o *SystemProfile) SetKatelloAgentRunning(v bool)`

SetKatelloAgentRunning sets KatelloAgentRunning field to given value.

### HasKatelloAgentRunning

`func (o *SystemProfile) HasKatelloAgentRunning() bool`

HasKatelloAgentRunning returns a boolean if a field has been set.

### GetKernelModules

`func (o *SystemProfile) GetKernelModules() []string`

GetKernelModules returns the KernelModules field if non-nil, zero value otherwise.

### GetKernelModulesOk

`func (o *SystemProfile) GetKernelModulesOk() (*[]string, bool)`

GetKernelModulesOk returns a tuple with the KernelModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelModules

`func (o *SystemProfile) SetKernelModules(v []string)`

SetKernelModules sets KernelModules field to given value.

### HasKernelModules

`func (o *SystemProfile) HasKernelModules() bool`

HasKernelModules returns a boolean if a field has been set.

### GetLastBootTime

`func (o *SystemProfile) GetLastBootTime() string`

GetLastBootTime returns the LastBootTime field if non-nil, zero value otherwise.

### GetLastBootTimeOk

`func (o *SystemProfile) GetLastBootTimeOk() (*string, bool)`

GetLastBootTimeOk returns a tuple with the LastBootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBootTime

`func (o *SystemProfile) SetLastBootTime(v string)`

SetLastBootTime sets LastBootTime field to given value.

### HasLastBootTime

`func (o *SystemProfile) HasLastBootTime() bool`

HasLastBootTime returns a boolean if a field has been set.

### GetMssql

`func (o *SystemProfile) GetMssql() SystemProfileMssql`

GetMssql returns the Mssql field if non-nil, zero value otherwise.

### GetMssqlOk

`func (o *SystemProfile) GetMssqlOk() (*SystemProfileMssql, bool)`

GetMssqlOk returns a tuple with the Mssql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMssql

`func (o *SystemProfile) SetMssql(v SystemProfileMssql)`

SetMssql sets Mssql field to given value.

### HasMssql

`func (o *SystemProfile) HasMssql() bool`

HasMssql returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *SystemProfile) GetNetworkInterfaces() []SystemProfileNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *SystemProfile) GetNetworkInterfacesOk() (*[]SystemProfileNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *SystemProfile) SetNetworkInterfaces(v []SystemProfileNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *SystemProfile) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetNumberOfCpus

`func (o *SystemProfile) GetNumberOfCpus() int32`

GetNumberOfCpus returns the NumberOfCpus field if non-nil, zero value otherwise.

### GetNumberOfCpusOk

`func (o *SystemProfile) GetNumberOfCpusOk() (*int32, bool)`

GetNumberOfCpusOk returns a tuple with the NumberOfCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCpus

`func (o *SystemProfile) SetNumberOfCpus(v int32)`

SetNumberOfCpus sets NumberOfCpus field to given value.

### HasNumberOfCpus

`func (o *SystemProfile) HasNumberOfCpus() bool`

HasNumberOfCpus returns a boolean if a field has been set.

### GetNumberOfSockets

`func (o *SystemProfile) GetNumberOfSockets() int32`

GetNumberOfSockets returns the NumberOfSockets field if non-nil, zero value otherwise.

### GetNumberOfSocketsOk

`func (o *SystemProfile) GetNumberOfSocketsOk() (*int32, bool)`

GetNumberOfSocketsOk returns a tuple with the NumberOfSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSockets

`func (o *SystemProfile) SetNumberOfSockets(v int32)`

SetNumberOfSockets sets NumberOfSockets field to given value.

### HasNumberOfSockets

`func (o *SystemProfile) HasNumberOfSockets() bool`

HasNumberOfSockets returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *SystemProfile) GetOperatingSystem() SystemProfileOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *SystemProfile) GetOperatingSystemOk() (*SystemProfileOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *SystemProfile) SetOperatingSystem(v SystemProfileOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *SystemProfile) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOsKernelVersion

`func (o *SystemProfile) GetOsKernelVersion() string`

GetOsKernelVersion returns the OsKernelVersion field if non-nil, zero value otherwise.

### GetOsKernelVersionOk

`func (o *SystemProfile) GetOsKernelVersionOk() (*string, bool)`

GetOsKernelVersionOk returns a tuple with the OsKernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsKernelVersion

`func (o *SystemProfile) SetOsKernelVersion(v string)`

SetOsKernelVersion sets OsKernelVersion field to given value.

### HasOsKernelVersion

`func (o *SystemProfile) HasOsKernelVersion() bool`

HasOsKernelVersion returns a boolean if a field has been set.

### GetOsRelease

`func (o *SystemProfile) GetOsRelease() string`

GetOsRelease returns the OsRelease field if non-nil, zero value otherwise.

### GetOsReleaseOk

`func (o *SystemProfile) GetOsReleaseOk() (*string, bool)`

GetOsReleaseOk returns a tuple with the OsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRelease

`func (o *SystemProfile) SetOsRelease(v string)`

SetOsRelease sets OsRelease field to given value.

### HasOsRelease

`func (o *SystemProfile) HasOsRelease() bool`

HasOsRelease returns a boolean if a field has been set.

### GetOwnerId

`func (o *SystemProfile) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SystemProfile) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SystemProfile) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SystemProfile) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRhcClientId

`func (o *SystemProfile) GetRhcClientId() string`

GetRhcClientId returns the RhcClientId field if non-nil, zero value otherwise.

### GetRhcClientIdOk

`func (o *SystemProfile) GetRhcClientIdOk() (*string, bool)`

GetRhcClientIdOk returns a tuple with the RhcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhcClientId

`func (o *SystemProfile) SetRhcClientId(v string)`

SetRhcClientId sets RhcClientId field to given value.

### HasRhcClientId

`func (o *SystemProfile) HasRhcClientId() bool`

HasRhcClientId returns a boolean if a field has been set.

### GetRhcConfigState

`func (o *SystemProfile) GetRhcConfigState() string`

GetRhcConfigState returns the RhcConfigState field if non-nil, zero value otherwise.

### GetRhcConfigStateOk

`func (o *SystemProfile) GetRhcConfigStateOk() (*string, bool)`

GetRhcConfigStateOk returns a tuple with the RhcConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhcConfigState

`func (o *SystemProfile) SetRhcConfigState(v string)`

SetRhcConfigState sets RhcConfigState field to given value.

### HasRhcConfigState

`func (o *SystemProfile) HasRhcConfigState() bool`

HasRhcConfigState returns a boolean if a field has been set.

### GetRhsm

`func (o *SystemProfile) GetRhsm() SystemProfileRhsm`

GetRhsm returns the Rhsm field if non-nil, zero value otherwise.

### GetRhsmOk

`func (o *SystemProfile) GetRhsmOk() (*SystemProfileRhsm, bool)`

GetRhsmOk returns a tuple with the Rhsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRhsm

`func (o *SystemProfile) SetRhsm(v SystemProfileRhsm)`

SetRhsm sets Rhsm field to given value.

### HasRhsm

`func (o *SystemProfile) HasRhsm() bool`

HasRhsm returns a boolean if a field has been set.

### GetRpmOstreeDeployments

`func (o *SystemProfile) GetRpmOstreeDeployments() []SystemProfileRpmOstreeDeployments`

GetRpmOstreeDeployments returns the RpmOstreeDeployments field if non-nil, zero value otherwise.

### GetRpmOstreeDeploymentsOk

`func (o *SystemProfile) GetRpmOstreeDeploymentsOk() (*[]SystemProfileRpmOstreeDeployments, bool)`

GetRpmOstreeDeploymentsOk returns a tuple with the RpmOstreeDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpmOstreeDeployments

`func (o *SystemProfile) SetRpmOstreeDeployments(v []SystemProfileRpmOstreeDeployments)`

SetRpmOstreeDeployments sets RpmOstreeDeployments field to given value.

### HasRpmOstreeDeployments

`func (o *SystemProfile) HasRpmOstreeDeployments() bool`

HasRpmOstreeDeployments returns a boolean if a field has been set.

### GetRunningProcesses

`func (o *SystemProfile) GetRunningProcesses() []string`

GetRunningProcesses returns the RunningProcesses field if non-nil, zero value otherwise.

### GetRunningProcessesOk

`func (o *SystemProfile) GetRunningProcessesOk() (*[]string, bool)`

GetRunningProcessesOk returns a tuple with the RunningProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningProcesses

`func (o *SystemProfile) SetRunningProcesses(v []string)`

SetRunningProcesses sets RunningProcesses field to given value.

### HasRunningProcesses

`func (o *SystemProfile) HasRunningProcesses() bool`

HasRunningProcesses returns a boolean if a field has been set.

### GetSapInstanceNumber

`func (o *SystemProfile) GetSapInstanceNumber() string`

GetSapInstanceNumber returns the SapInstanceNumber field if non-nil, zero value otherwise.

### GetSapInstanceNumberOk

`func (o *SystemProfile) GetSapInstanceNumberOk() (*string, bool)`

GetSapInstanceNumberOk returns a tuple with the SapInstanceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapInstanceNumber

`func (o *SystemProfile) SetSapInstanceNumber(v string)`

SetSapInstanceNumber sets SapInstanceNumber field to given value.

### HasSapInstanceNumber

`func (o *SystemProfile) HasSapInstanceNumber() bool`

HasSapInstanceNumber returns a boolean if a field has been set.

### GetSapSids

`func (o *SystemProfile) GetSapSids() []string`

GetSapSids returns the SapSids field if non-nil, zero value otherwise.

### GetSapSidsOk

`func (o *SystemProfile) GetSapSidsOk() (*[]string, bool)`

GetSapSidsOk returns a tuple with the SapSids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSids

`func (o *SystemProfile) SetSapSids(v []string)`

SetSapSids sets SapSids field to given value.

### HasSapSids

`func (o *SystemProfile) HasSapSids() bool`

HasSapSids returns a boolean if a field has been set.

### GetSapSystem

`func (o *SystemProfile) GetSapSystem() bool`

GetSapSystem returns the SapSystem field if non-nil, zero value otherwise.

### GetSapSystemOk

`func (o *SystemProfile) GetSapSystemOk() (*bool, bool)`

GetSapSystemOk returns a tuple with the SapSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapSystem

`func (o *SystemProfile) SetSapSystem(v bool)`

SetSapSystem sets SapSystem field to given value.

### HasSapSystem

`func (o *SystemProfile) HasSapSystem() bool`

HasSapSystem returns a boolean if a field has been set.

### GetSapVersion

`func (o *SystemProfile) GetSapVersion() string`

GetSapVersion returns the SapVersion field if non-nil, zero value otherwise.

### GetSapVersionOk

`func (o *SystemProfile) GetSapVersionOk() (*string, bool)`

GetSapVersionOk returns a tuple with the SapVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSapVersion

`func (o *SystemProfile) SetSapVersion(v string)`

SetSapVersion sets SapVersion field to given value.

### HasSapVersion

`func (o *SystemProfile) HasSapVersion() bool`

HasSapVersion returns a boolean if a field has been set.

### GetSatelliteManaged

`func (o *SystemProfile) GetSatelliteManaged() bool`

GetSatelliteManaged returns the SatelliteManaged field if non-nil, zero value otherwise.

### GetSatelliteManagedOk

`func (o *SystemProfile) GetSatelliteManagedOk() (*bool, bool)`

GetSatelliteManagedOk returns a tuple with the SatelliteManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteManaged

`func (o *SystemProfile) SetSatelliteManaged(v bool)`

SetSatelliteManaged sets SatelliteManaged field to given value.

### HasSatelliteManaged

`func (o *SystemProfile) HasSatelliteManaged() bool`

HasSatelliteManaged returns a boolean if a field has been set.

### GetSelinuxConfigFile

`func (o *SystemProfile) GetSelinuxConfigFile() string`

GetSelinuxConfigFile returns the SelinuxConfigFile field if non-nil, zero value otherwise.

### GetSelinuxConfigFileOk

`func (o *SystemProfile) GetSelinuxConfigFileOk() (*string, bool)`

GetSelinuxConfigFileOk returns a tuple with the SelinuxConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxConfigFile

`func (o *SystemProfile) SetSelinuxConfigFile(v string)`

SetSelinuxConfigFile sets SelinuxConfigFile field to given value.

### HasSelinuxConfigFile

`func (o *SystemProfile) HasSelinuxConfigFile() bool`

HasSelinuxConfigFile returns a boolean if a field has been set.

### GetSelinuxCurrentMode

`func (o *SystemProfile) GetSelinuxCurrentMode() string`

GetSelinuxCurrentMode returns the SelinuxCurrentMode field if non-nil, zero value otherwise.

### GetSelinuxCurrentModeOk

`func (o *SystemProfile) GetSelinuxCurrentModeOk() (*string, bool)`

GetSelinuxCurrentModeOk returns a tuple with the SelinuxCurrentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxCurrentMode

`func (o *SystemProfile) SetSelinuxCurrentMode(v string)`

SetSelinuxCurrentMode sets SelinuxCurrentMode field to given value.

### HasSelinuxCurrentMode

`func (o *SystemProfile) HasSelinuxCurrentMode() bool`

HasSelinuxCurrentMode returns a boolean if a field has been set.

### GetSubscriptionAutoAttach

`func (o *SystemProfile) GetSubscriptionAutoAttach() string`

GetSubscriptionAutoAttach returns the SubscriptionAutoAttach field if non-nil, zero value otherwise.

### GetSubscriptionAutoAttachOk

`func (o *SystemProfile) GetSubscriptionAutoAttachOk() (*string, bool)`

GetSubscriptionAutoAttachOk returns a tuple with the SubscriptionAutoAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAutoAttach

`func (o *SystemProfile) SetSubscriptionAutoAttach(v string)`

SetSubscriptionAutoAttach sets SubscriptionAutoAttach field to given value.

### HasSubscriptionAutoAttach

`func (o *SystemProfile) HasSubscriptionAutoAttach() bool`

HasSubscriptionAutoAttach returns a boolean if a field has been set.

### GetSubscriptionStatus

`func (o *SystemProfile) GetSubscriptionStatus() string`

GetSubscriptionStatus returns the SubscriptionStatus field if non-nil, zero value otherwise.

### GetSubscriptionStatusOk

`func (o *SystemProfile) GetSubscriptionStatusOk() (*string, bool)`

GetSubscriptionStatusOk returns a tuple with the SubscriptionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStatus

`func (o *SystemProfile) SetSubscriptionStatus(v string)`

SetSubscriptionStatus sets SubscriptionStatus field to given value.

### HasSubscriptionStatus

`func (o *SystemProfile) HasSubscriptionStatus() bool`

HasSubscriptionStatus returns a boolean if a field has been set.

### GetSystemMemoryBytes

`func (o *SystemProfile) GetSystemMemoryBytes() int64`

GetSystemMemoryBytes returns the SystemMemoryBytes field if non-nil, zero value otherwise.

### GetSystemMemoryBytesOk

`func (o *SystemProfile) GetSystemMemoryBytesOk() (*int64, bool)`

GetSystemMemoryBytesOk returns a tuple with the SystemMemoryBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMemoryBytes

`func (o *SystemProfile) SetSystemMemoryBytes(v int64)`

SetSystemMemoryBytes sets SystemMemoryBytes field to given value.

### HasSystemMemoryBytes

`func (o *SystemProfile) HasSystemMemoryBytes() bool`

HasSystemMemoryBytes returns a boolean if a field has been set.

### GetSystemPurpose

`func (o *SystemProfile) GetSystemPurpose() SystemProfileSystemPurpose`

GetSystemPurpose returns the SystemPurpose field if non-nil, zero value otherwise.

### GetSystemPurposeOk

`func (o *SystemProfile) GetSystemPurposeOk() (*SystemProfileSystemPurpose, bool)`

GetSystemPurposeOk returns a tuple with the SystemPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPurpose

`func (o *SystemProfile) SetSystemPurpose(v SystemProfileSystemPurpose)`

SetSystemPurpose sets SystemPurpose field to given value.

### HasSystemPurpose

`func (o *SystemProfile) HasSystemPurpose() bool`

HasSystemPurpose returns a boolean if a field has been set.

### GetTunedProfile

`func (o *SystemProfile) GetTunedProfile() string`

GetTunedProfile returns the TunedProfile field if non-nil, zero value otherwise.

### GetTunedProfileOk

`func (o *SystemProfile) GetTunedProfileOk() (*string, bool)`

GetTunedProfileOk returns a tuple with the TunedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunedProfile

`func (o *SystemProfile) SetTunedProfile(v string)`

SetTunedProfile sets TunedProfile field to given value.

### HasTunedProfile

`func (o *SystemProfile) HasTunedProfile() bool`

HasTunedProfile returns a boolean if a field has been set.

### GetYumRepos

`func (o *SystemProfile) GetYumRepos() []SystemProfileYumRepo`

GetYumRepos returns the YumRepos field if non-nil, zero value otherwise.

### GetYumReposOk

`func (o *SystemProfile) GetYumReposOk() (*[]SystemProfileYumRepo, bool)`

GetYumReposOk returns a tuple with the YumRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumRepos

`func (o *SystemProfile) SetYumRepos(v []SystemProfileYumRepo)`

SetYumRepos sets YumRepos field to given value.

### HasYumRepos

`func (o *SystemProfile) HasYumRepos() bool`

HasYumRepos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


