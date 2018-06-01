/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Object containing node sync status
type GetSyncResponse struct {

	// Current sync status
	Status string `json:"status,omitempty"`

	// Current blockchain height
	BlockChainHeight float32 `json:"blockChainHeight,omitempty"`

	// Current sync percentage
	SyncPercentage float32 `json:"syncPercentage,omitempty"`

	// Height node is synced to
	Height float32 `json:"height,omitempty"`

	// Recent sync error messages
	Error_ string `json:"error,omitempty"`

	// Node type
	Type_ string `json:"type,omitempty"`
}