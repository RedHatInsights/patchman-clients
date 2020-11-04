/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// UpdatesRequest struct for UpdatesRequest
type UpdatesRequest struct {
	PackageList []string `json:"package_list"`
	RepositoryList []string `json:"repository_list,omitempty"`
	ModulesList []UpdatesRequestModulesList `json:"modules_list,omitempty"`
	Releasever string `json:"releasever,omitempty"`
	Basearch string `json:"basearch,omitempty"`
	LatestOnly bool `json:"latest_only,omitempty"`
}
