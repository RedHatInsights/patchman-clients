/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// SystemProfileInstalledProducts Representation of one installed product
type SystemProfileInstalledProducts struct {
	// The product ID
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// Subscription status for product
	Status string `json:"status,omitempty"`
}
