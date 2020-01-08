/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// UpdatesV3Request struct for UpdatesV3Request
type UpdatesV3Request struct {
	PackageList []string `json:"package_list"`
	RepositoryList []string `json:"repository_list,omitempty"`
	ModulesList []UpdatesRequestModulesList `json:"modules_list,omitempty"`
	Releasever string `json:"releasever,omitempty"`
	Basearch string `json:"basearch,omitempty"`
	SecurityOnly bool `json:"security_only,omitempty"`
}
