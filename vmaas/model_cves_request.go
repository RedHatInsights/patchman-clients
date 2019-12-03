/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// CvesRequest struct for CvesRequest
type CvesRequest struct {
	Page float32 `json:"page,omitempty"`
	PageSize float32 `json:"page_size,omitempty"`
	CveList []string `json:"cve_list"`
	ModifiedSince string `json:"modified_since,omitempty"`
	PublishedSince string `json:"published_since,omitempty"`
	RhOnly bool `json:"rh_only,omitempty"`
}