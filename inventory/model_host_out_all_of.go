/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// HostOutAllOf struct for HostOutAllOf
type HostOutAllOf struct {
	// A set of facts belonging to the host.
	Facts []map[string]interface{} `json:"facts,omitempty"`
}
