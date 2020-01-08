/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// ErrataResponseErrataList struct for ErrataResponseErrataList
type ErrataResponseErrataList struct {
	Updated string `json:"updated,omitempty"`
	Severity string `json:"severity,omitempty"`
	ReferenceList []string `json:"reference_list,omitempty"`
	Issued string `json:"issued,omitempty"`
	Description string `json:"description,omitempty"`
	Solution string `json:"solution,omitempty"`
	Summary string `json:"summary,omitempty"`
	Url string `json:"url,omitempty"`
	Synopsis string `json:"synopsis,omitempty"`
	CveList []string `json:"cve_list,omitempty"`
	BugzillaList []string `json:"bugzilla_list,omitempty"`
	PackageList []string `json:"package_list,omitempty"`
	SourcePackageList []string `json:"source_package_list,omitempty"`
	Type string `json:"type,omitempty"`
}
