# CreateHostOutAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | Pointer to **string** | The ansible host name for remediations | [optional] 
**Created** | **string** | A timestamp when the entry was created. | [optional] 
**CulledTimestamp** | Pointer to **string** | Timestamp from which the host is considered deleted. | [optional] 
**DisplayName** | Pointer to **string** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**Facts** | **[]map[string]interface{}** | A set of facts belonging to the host. | [optional] 
**Id** | **string** | A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts. | [optional] 
**Reporter** | Pointer to **string** | Reporting source of the host. Used when updating the stale_timestamp. | [optional] 
**StaleTimestamp** | Pointer to **string** | Timestamp from which the host is considered stale. | [optional] 
**StaleWarningTimestamp** | Pointer to **string** | Timestamp from which the host is considered too stale to be listed without an explicit toggle. | [optional] 
**Updated** | **string** | A timestamp when the entry was last updated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


