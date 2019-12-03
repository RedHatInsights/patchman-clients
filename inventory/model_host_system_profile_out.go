/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// HostSystemProfileOut Individual host record that contains only the host id and system profile
type HostSystemProfileOut struct {
	Id string `json:"id,omitempty"`
	SystemProfile SystemProfileIn `json:"system_profile,omitempty"`
}