/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.12.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// UpdatesV2Response struct for UpdatesV2Response
type UpdatesV2Response struct {
	UpdateList map[string]UpdatesV2ResponseUpdateList `json:"update_list,omitempty"`
	RepositoryList []string `json:"repository_list,omitempty"`
	ModulesList []UpdatesRequestModulesList `json:"modules_list,omitempty"`
	Releasever string `json:"releasever,omitempty"`
	Basearch string `json:"basearch,omitempty"`
}
