# SystemProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | **string** |  | [optional] 
**BiosReleaseDate** | **string** |  | [optional] 
**BiosVendor** | **string** |  | [optional] 
**BiosVersion** | **string** |  | [optional] 
**CapturedDate** | **string** |  | [optional] 
**CloudProvider** | **string** |  | [optional] 
**CoresPerSocket** | **int32** |  | [optional] 
**CpuFlags** | **[]string** |  | [optional] 
**DiskDevices** | [**[]SystemProfileDiskDevices**](SystemProfile_disk_devices.md) |  | [optional] 
**DnfModules** | [**[]SystemProfileDnfModules**](SystemProfile_dnf_modules.md) |  | [optional] 
**EnabledServices** | **[]string** |  | [optional] 
**InfrastructureType** | **string** |  | [optional] 
**InfrastructureVendor** | **string** |  | [optional] 
**InsightsClientVersion** | **string** |  | [optional] 
**InsightsEggVersion** | **string** |  | [optional] 
**InstalledPackages** | **[]string** |  | [optional] 
**InstalledProducts** | [**[]SystemProfileInstalledProducts**](SystemProfile_installed_products.md) |  | [optional] 
**InstalledServices** | **[]string** |  | [optional] 
**KatelloAgentRunning** | **bool** |  | [optional] 
**KernelModules** | **[]string** |  | [optional] 
**LastBootTime** | **string** |  | [optional] 
**NetworkInterfaces** | [**[]SystemProfileNetworkInterfaces**](SystemProfile_network_interfaces.md) |  | [optional] 
**NumberOfCpus** | **int32** |  | [optional] 
**NumberOfSockets** | **int32** |  | [optional] 
**OsKernelVersion** | **string** |  | [optional] 
**OsRelease** | **string** |  | [optional] 
**RunningProcesses** | **[]string** |  | [optional] 
**SapInstanceNumber** | **string** | The instance number of the SAP HANA system | [optional] 
**SapSids** | **[]string** | List of SAP SIDs | [optional] 
**SapSystem** | **bool** | Indicates if SAP is installed on the system | [optional] 
**SapVersion** | **string** | The version of the SAP HANA lifecycle management program | [optional] 
**SatelliteManaged** | **bool** |  | [optional] 
**SelinuxConfigFile** | **string** | The SELinux mode provided in the config file | [optional] 
**SelinuxCurrentMode** | **string** | The current SELinux mode, either enforcing, permissive, or disabled | [optional] 
**SubscriptionAutoAttach** | **string** |  | [optional] 
**SubscriptionStatus** | **string** |  | [optional] 
**SystemMemoryBytes** | **int64** |  | [optional] 
**TunedProfile** | **string** | Current profile resulting from command tuned-adm active | [optional] 
**YumRepos** | [**[]SystemProfileYumRepos**](SystemProfile_yum_repos.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


