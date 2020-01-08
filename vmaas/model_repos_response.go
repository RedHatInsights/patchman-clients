/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// ReposResponse struct for ReposResponse
type ReposResponse struct {
	Page float32 `json:"page,omitempty"`
	PageSize float32 `json:"page_size,omitempty"`
	Pages float32 `json:"pages,omitempty"`
	RepositoryList map[string][]map[string]interface{} `json:"repository_list,omitempty"`
}
